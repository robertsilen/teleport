/**
 * Teleport
 * Copyright (C) 2023  Gravitational, Inc.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

import {
  useEffect,
  useRef,
  useState,
  useMemo,
  Dispatch,
  SetStateAction,
} from 'react';
import { useParams } from 'react-router';
import { debounce } from 'shared/utils/highbar';

import useAttempt from 'shared/hooks/useAttemptNext';

import { ButtonState, TdpClient } from 'teleport/lib/tdp';
import useWebAuthn from 'teleport/lib/useWebAuthn';
import desktopService from 'teleport/services/desktops';
import userService from 'teleport/services/user';
import { getHostName } from 'teleport/services/api';
import {
  ClientScreenSpec,
  ClipboardData,
  PngFrame,
  PointerData,
} from 'teleport/lib/tdp/codec';
import { Sha256Digest } from 'teleport/lib/util';
import cfg from 'teleport/config';
import { BitmapFrame } from 'teleport/lib/tdp/client';

import useTdpClientCanvas from './useTdpClientCanvas';
import { TopBarHeight } from './TopBar';

import type { UrlDesktopParams } from 'teleport/config';
import type { NotificationItem } from 'shared/components/Notification';

export type TdpConnection = {
  status: '' | 'open' | 'closed';
  receivedFirstFrame?: boolean;
  statusText: string;
};

export default function useDesktopSession() {
  const { attempt: fetchAttempt, run } = useAttempt('processing');
  const latestClipboardDigest = useRef('');
  const encoder = useRef(new TextEncoder());
  const [tdpConnection, setTdpConnection] = useState<TdpConnection>({
    status: '',
    statusText: '',
  });

  // // tdpConnection tracks the state of the tdpClient's TDP connection
  // // - 'processing' at first
  // // - 'success' once the first TdpClientEvent.IMAGE_FRAGMENT is seen
  // // - 'failed' if a fatal error is encountered, should have a statusText
  // // - '' if the connection closed gracefully by the server, should have a statusText
  // const { attempt: tdpConnection, setAttempt: setTdpConnection } =
  //   useAttempt('processing');
  const [tdpClient, setTdpClient] = useState<TdpClient>(null);
  const clientCanvasProps = useTdpClientCanvas(tdpClient);

  const { username, desktopName, clusterId } = useParams<UrlDesktopParams>();

  const [hostname, setHostname] = useState<string>('');

  const [directorySharingState, setDirectorySharingState] =
    useState<DirectorySharingState>(defaultDirectorySharingState);

  const [clipboardSharingState, setClipboardSharingState] =
    useState<ClipboardSharingState>(defaultClipboardSharingState);

  useEffect(() => {
    const clearReadListenerPromise = initClipboardPermissionTracking(
      'clipboard-read',
      setClipboardSharingState
    );
    const clearWriteListenerPromise = initClipboardPermissionTracking(
      'clipboard-write',
      setClipboardSharingState
    );

    return () => {
      clearReadListenerPromise.then(clearReadListener => clearReadListener());
      clearWriteListenerPromise.then(clearWriteListener =>
        clearWriteListener()
      );
    };
  }, []);

  const [showAnotherSessionActiveDialog, setShowAnotherSessionActiveDialog] =
    useState(false);

  document.title = useMemo(
    () => `${username}@${hostname} • ${clusterId}`,
    [clusterId, hostname, username]
  );

  useEffect(() => {
    run(() =>
      Promise.all([
        desktopService
          .fetchDesktop(clusterId, desktopName)
          .then(desktop => setHostname(desktop.name)),
        userService.fetchUserContext().then(user => {
          setClipboardSharingState(prevState => ({
            ...prevState,
            allowedByAcl: user.acl.clipboardSharingEnabled,
          }));
          setDirectorySharingState(prevState => ({
            ...prevState,
            allowedByAcl: user.acl.directorySharingEnabled,
          }));
        }),
        desktopService
          .checkDesktopIsActive(clusterId, desktopName)
          .then(isActive => {
            setShowAnotherSessionActiveDialog(isActive);
          }),
      ])
    );
  }, [clusterId, desktopName, run]);

  const [alerts, setAlerts] = useState<NotificationItem[]>([]);
  const onRemoveAlert = (id: string) => {
    setAlerts(prevState => prevState.filter(alert => alert.id !== id));
  };

  const addr = cfg.api.desktopWsAddr
    .replace(':fqdn', getHostName())
    .replace(':clusterId', clusterId)
    .replace(':desktopName', desktopName)
    .replace(':username', username);

  // Default TdpClientEvent.TDP_CLIPBOARD_DATA handler.
  const onClipboardData = async (clipboardData: ClipboardData) => {
    if (
      clipboardData.data &&
      (await sysClipboardGuard(clipboardSharingState, 'write'))
    ) {
      navigator.clipboard.writeText(clipboardData.data);
      let digest = await Sha256Digest(clipboardData.data, encoder.current);
      latestClipboardDigest.current = digest;
    }
  };

  // example of pulling client and canvas out of tdpclientcanvas
  const onScreenSpec = (spec: ClientScreenSpec) => {
    clientCanvasProps.syncCanvas(spec);
  };

  // Default TdpClientEvent.TDP_ERROR and TdpClientEvent.CLIENT_ERROR handler
  const onError = (error: Error) => {
    // setDirectorySharingState(defaultDirectorySharingState);
    setClipboardSharingState(defaultClipboardSharingState);
    // should merge this + wsStatus into 1 connection var
    setTdpConnection(prevState => {
      // Sometimes when a connection closes due to an error, we get a cascade of
      // errors. Here we update the status only if it's not already 'failed', so
      // that the first error message (which is usually the most informative) is
      // displayed to the user.
      if (prevState.status !== 'closed') {
        return {
          status: 'closed',
          statusText: error.message || error.toString(),
        };
      }
      return prevState;
    });
  };

  // Default TdpClientEvent.TDP_WARNING and TdpClientEvent.CLIENT_WARNING handler
  const onWarning = (warning: string) => {
    setAlerts(prevState => {
      return [
        ...prevState,
        {
          content: warning,
          severity: 'warn',
          id: crypto.randomUUID(),
        },
      ];
    });
  };

  // TODO(zmb3): this is not what an info-level alert should do.
  // rename it to something like onGracefulDisconnect
  const onInfo = (info: string) => {
    // setDirectorySharingState(defaultDirectorySharingState);
    setClipboardSharingState(defaultClipboardSharingState);
    setTdpConnection({
      status: 'closed', // gracefully disconnecting
      statusText: info,
    });
  };

  const onWsOpen = () => {
    setTdpConnection({ status: 'open', statusText: '' });
  };

  const onWsClose = (message: string) => {
    setTdpConnection({ status: 'closed', statusText: message });
  };

  // create a closure to enable rendered buffering and return
  // the "listener" from it
  const onBmpFrame = () => {
    const canvas = clientCanvasProps.canvasRef.current;
    if (!canvas) {
      return;
    }
    const ctx = canvas.getContext('2d');

    // Buffered rendering logic
    var bmpBuffer: BitmapFrame[] = [];
    const renderBuffer = () => {
      if (bmpBuffer.length) {
        for (let i = 0; i < bmpBuffer.length; i++) {
          // not sure why we sync the canvas during first frame when it doesn't seem
          // to care about any of the frame data at all? and we sync canvas
          if (!tdpConnection.receivedFirstFrame) {
            setTdpConnection(prevState => ({
              ...prevState,
              receivedFirstFrame: true,
            }));
          }
          const bmpFrame = bmpBuffer[i];
          if (ctx) {
            ctx.putImageData(bmpFrame.image_data, bmpFrame.left, bmpFrame.top);
          }
        }
        bmpBuffer = [];
      }
      requestAnimationFrame(renderBuffer);
    };
    requestAnimationFrame(renderBuffer);

    const pushToBmpBuffer = (bmpFrame: BitmapFrame) => {
      bmpBuffer.push(bmpFrame);
    };
    return pushToBmpBuffer;
  };

  // create a closure to enable rendered buffering and return
  // the "listener" from it
  const onPngFrame = () => {
    const canvas = clientCanvasProps.canvasRef.current;
    if (!canvas) {
      return;
    }
    const ctx = canvas.getContext('2d');

    // Buffered rendering logic
    var pngBuffer: PngFrame[] = [];
    const renderBuffer = () => {
      if (pngBuffer.length) {
        for (let i = 0; i < pngBuffer.length; i++) {
          // not sure why we sync the canvas during first frame when it doesn't seem
          // to care about any of the frame data at all? and we sync canvas
          if (!tdpConnection.receivedFirstFrame) {
            setTdpConnection(prevState => ({
              ...prevState,
              receivedFirstFrame: true,
            }));
            clientCanvasProps.syncCanvas(getDisplaySize());
          }
          const pngFrame = pngBuffer[i];
          if (ctx) {
            ctx.drawImage(pngFrame.data, pngFrame.left, pngFrame.top);
          }
        }
        pngBuffer = [];
      }
      requestAnimationFrame(renderBuffer);
    };
    requestAnimationFrame(renderBuffer);

    const pushToPngBuffer = (pngFrame: PngFrame) => {
      pngBuffer.push(pngFrame);
    };
    return pushToPngBuffer;
  };

  const onPointer = (pointer: PointerData) => {
    const canvas = clientCanvasProps.canvasRef.current;
    if (!canvas) {
      return;
    }
    if (typeof pointer.data === 'boolean') {
      canvas.style.cursor = pointer.data ? 'default' : 'none';
      return;
    }
    let cursor = document.createElement('canvas');
    cursor.width = pointer.data.width;
    cursor.height = pointer.data.height;
    cursor
      .getContext('2d', { colorSpace: pointer.data.colorSpace })
      .putImageData(pointer.data, 0, 0);
    if (pointer.data.width > 32 || pointer.data.height > 32) {
      // scale the cursor down to at most 32px - max size fully supported by browsers
      const resized = document.createElement('canvas');
      let scale = Math.min(32 / cursor.width, 32 / cursor.height);
      resized.width = cursor.width * scale;
      resized.height = cursor.height * scale;

      let context = resized.getContext('2d', {
        colorSpace: pointer.data.colorSpace,
      });
      context.scale(scale, scale);
      context.drawImage(cursor, 0, 0);
      cursor = resized;
    }
    canvas.style.cursor = `url(${cursor.toDataURL()}) ${
      pointer.hotspot_x
    } ${pointer.hotspot_y}, auto`;
  };

  useEffect(() => {
    if (!tdpClient) {
      setTdpClient(
        new TdpClient(addr, {
          onClipboardData,
          onError,
          onWarning,
          onInfo,
          onWsOpen,
          onPngFrame: onPngFrame(), // for buffered rendering
          onBmpFrame: onBmpFrame(),
          onScreenSpec,
          onPointer,
          onWsClose,
        })
      );
    }
    // TODO (avatus) : fix this
    // eslint-disable-next-line
  }, [tdpClient, addr]);

  const sendLocalClipboardToRemote = async (cli: TdpClient) => {
    if (await sysClipboardGuard(clipboardSharingState, 'read')) {
      navigator.clipboard.readText().then(text => {
        Sha256Digest(text, encoder.current).then(digest => {
          if (text && digest !== latestClipboardDigest.current) {
            cli.sendClipboardData({
              data: text,
            });
            latestClipboardDigest.current = digest;
          }
        });
      });
    }
  };

  const webauthn = useWebAuthn(tdpClient);

  const onDisconnect = () => {
    setClipboardSharingState(prevState => ({
      ...prevState,
      isSharing: false,
    }));
    setDirectorySharingState(prevState => ({
      ...prevState,
      isSharing: false,
    }));
    tdpClient.shutdown();
  };

  const onShareDirectory = () => {
    try {
      window
        .showDirectoryPicker()
        .then(sharedDirHandle => {
          // Permissions granted and/or directory selected
          setDirectorySharingState(prevState => ({
            ...prevState,
            directorySelected: true,
          }));
          tdpClient.addSharedDirectory(sharedDirHandle);
          tdpClient.sendSharedDirectoryAnnounce();
        })
        .catch(e => {
          setDirectorySharingState(prevState => ({
            ...prevState,
            directorySelected: false,
          }));
          setAlerts(prevState => [
            ...prevState,
            {
              id: crypto.randomUUID(),
              severity: 'warn',
              content: 'Failed to open the directory picker: ' + e.message,
            },
          ]);
        });
    } catch (e) {
      setDirectorySharingState(prevState => ({
        ...prevState,
        directorySelected: false,
      }));
      setAlerts(prevState => [
        ...prevState,
        {
          id: crypto.randomUUID(),
          severity: 'warn',
          // This is a gross error message, but should be infrequent enough that its worth just telling
          // the user the likely problem, while also displaying the error message just in case that's not it.
          // In a perfect world, we could check for which error message this is and display
          // context appropriate directions.
          content:
            'Encountered an error while attempting to share a directory: ' +
            e.message +
            '. \n\nYour user role supports directory sharing over desktop access, \
          however this feature is only available by default on some Chromium \
          based browsers like Google Chrome or Microsoft Edge. Brave users can \
          use the feature by navigating to brave://flags/#file-system-access-api \
          and selecting "Enable". If you\'re not already, please switch to a supported browser.',
        },
      ]);
    }
  };

  const onCtrlAltDel = () => {
    if (!tdpClient) {
      return;
    }
    tdpClient.sendKeyboardInput('ControlLeft', ButtonState.DOWN);
    tdpClient.sendKeyboardInput('AltLeft', ButtonState.DOWN);
    tdpClient.sendKeyboardInput('Delete', ButtonState.DOWN);
  };

  const windowOnResize = debounce(
    () => {
      const spec = getDisplaySize();
      tdpClient.resize(spec);
    },
    250,
    { trailing: true }
  );

  return {
    webauthn,
    tdpClient,
    username,
    hostname,
    tdpConnection,
    onCtrlAltDel,
    alerts,
    onRemoveAlert,
    onDisconnect,
    clipboardSharingState,
    directorySharingState,
    clientCanvasProps,
    fetchAttempt,
    windowOnResize,
    onShareDirectory,
    showAnotherSessionActiveDialog,
    setShowAnotherSessionActiveDialog,
  };
}

export type State = ReturnType<typeof useDesktopSession>;

type CommonFeatureState = {
  /**
   * Whether the feature is allowed by the acl.
   *
   * Undefined if it hasn't been queried yet.
   */
  allowedByAcl?: boolean;
  /**
   * Whether the feature is available in the browser.
   */
  browserSupported: boolean;
};

/**
 * The state of the directory sharing feature.
 */
export type DirectorySharingState = CommonFeatureState & {
  /**
   * Whether the user is currently sharing a directory.
   */
  directorySelected: boolean;
};

/**
 * The state of the clipboard sharing feature.
 */
export type ClipboardSharingState = CommonFeatureState & {
  /**
   * The current state of the 'clipboard-read' permission.
   *
   * Undefined if it hasn't been queried yet.
   */
  readState?: PermissionState;
  /**
   * The current state of the 'clipboard-write' permission.
   *
   * Undefined if it hasn't been queried yet.
   */
  writeState?: PermissionState;
};

export type Setter<T> = Dispatch<SetStateAction<T>>;

async function initClipboardPermissionTracking(
  name: 'clipboard-read' | 'clipboard-write',
  setClipboardSharingState: Setter<ClipboardSharingState>
) {
  const handleChange = () => {
    if (name === 'clipboard-read') {
      setClipboardSharingState(prevState => ({
        ...prevState,
        readState: perm.state,
      }));
    } else {
      setClipboardSharingState(prevState => ({
        ...prevState,
        writeState: perm.state,
      }));
    }
  };

  // Query the permission state
  const perm = await navigator.permissions.query({
    name: name as PermissionName,
  });

  // Set its change handler
  perm.onchange = handleChange;
  // Set the initial state
  handleChange();

  // Return a cleanup function that removes the change handler (for use by useEffect)
  return () => {
    perm.onchange = null;
  };
}

/**
 * Determines whether a feature is/should-be possible based on whether it's allowed by the acl
 * and whether it's supported by the browser.
 */
function commonFeaturePossible(
  commonFeatureState: CommonFeatureState
): boolean {
  return commonFeatureState.allowedByAcl && commonFeatureState.browserSupported;
}

/**
 * Determines whether clipboard sharing is/should-be possible based on whether it's allowed by the acl
 * and whether it's supported by the browser.
 */
export function clipboardSharingPossible(
  clipboardSharingState: ClipboardSharingState
): boolean {
  return commonFeaturePossible(clipboardSharingState);
}

/**
 * Returns whether clipboard sharing is active.
 */
export function isSharingClipboard(
  clipboardSharingState: ClipboardSharingState
): boolean {
  return (
    clipboardSharingState.allowedByAcl &&
    clipboardSharingState.browserSupported &&
    clipboardSharingState.readState === 'granted' &&
    clipboardSharingState.writeState === 'granted'
  );
}

/**
 * Provides a user-friendly message indicating whether clipboard sharing is enabled,
 * and the reason it is disabled.
 */
export function clipboardSharingMessage(state: ClipboardSharingState): string {
  if (!state.allowedByAcl) {
    return 'Clipboard Sharing disabled by Teleport RBAC.';
  }
  if (!state.browserSupported) {
    return 'Clipboard Sharing is not supported in this browser.';
  }
  if (state.readState === 'denied' || state.writeState === 'denied') {
    return 'Clipboard Sharing disabled due to browser permissions.';
  }

  return isSharingClipboard(state)
    ? 'Clipboard Sharing enabled.'
    : 'Clipboard Sharing disabled.';
}

/**
 * Determines whether directory sharing is/should-be possible based on whether it's allowed by the acl
 * and whether it's supported by the browser.
 */
export function directorySharingPossible(
  directorySharingState: DirectorySharingState
): boolean {
  return commonFeaturePossible(directorySharingState);
}

/**
 * Returns whether directory sharing is active.
 */
export function isSharingDirectory(
  directorySharingState: DirectorySharingState
): boolean {
  return (
    directorySharingState.allowedByAcl &&
    directorySharingState.browserSupported &&
    directorySharingState.directorySelected
  );
}

export const defaultDirectorySharingState: DirectorySharingState = {
  browserSupported: navigator.userAgent.includes('Chrome'),
  directorySelected: false,
};

export const defaultClipboardSharingState: ClipboardSharingState = {
  browserSupported: navigator.userAgent.includes('Chrome'),
};

/**
 * To be called before any system clipboard read/write operation.
 */
function sysClipboardGuard(
  clipboardSharingState: ClipboardSharingState,
  checkingFor: 'read' | 'write'
): boolean {
  // If we're not allowed to share the clipboard according to the acl
  // or due to the browser we're using, never try to read or write.
  if (!clipboardSharingPossible(clipboardSharingState)) {
    return false;
  }

  // If the relevant state is 'prompt', try the operation so that the
  // user is prompted to allow it.
  const checkingForRead = checkingFor === 'read';
  const checkingForWrite = checkingFor === 'write';
  const relevantStateIsPrompt =
    (checkingForRead && clipboardSharingState.readState === 'prompt') ||
    (checkingForWrite && clipboardSharingState.writeState === 'prompt');
  if (relevantStateIsPrompt) {
    return true;
  }

  // Otherwise try only if both read and write permissions are granted
  // and the document has focus (without focus we get an uncatchable error).
  //
  // Note that there's no situation where only one of read or write is granted,
  // but the other is denied, and we want to try the operation. The feature is
  // either fully enabled or fully disabled.
  return isSharingClipboard(clipboardSharingState) && document.hasFocus();
}

// Calculates the size (in pixels) of the display.
// Since we want to maximize the display size for the user, this is simply
// the full width of the screen and the full height sans top bar.
export function getDisplaySize() {
  return {
    width: window.innerWidth,
    height: window.innerHeight - TopBarHeight,
  };
}
