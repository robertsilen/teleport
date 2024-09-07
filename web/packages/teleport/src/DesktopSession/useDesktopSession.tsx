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

import useAttempt from 'shared/hooks/useAttemptNext';

import { ButtonState, TdpClient } from 'teleport/lib/tdp';
import useWebAuthn from 'teleport/lib/useWebAuthn';
import desktopService from 'teleport/services/desktops';
import userService from 'teleport/services/user';
import { getHostName } from 'teleport/services/api';
import { ClipboardData } from 'teleport/lib/tdp/codec';
import { Sha256Digest } from 'teleport/lib/util';
import cfg from 'teleport/config';

import useTdpClientCanvas from './useTdpClientCanvas';

import type { UrlDesktopParams } from 'teleport/config';
import type { NotificationItem } from 'shared/components/Notification';

export default function useDesktopSession() {
  const { attempt: fetchAttempt, run } = useAttempt('processing');
  const latestClipboardDigest = useRef('');
  const encoder = useRef(new TextEncoder());
  const clientCanvasProps = useTdpClientCanvas();

  // // tdpConnection tracks the state of the tdpClient's TDP connection
  // // - 'processing' at first
  // // - 'success' once the first TdpClientEvent.IMAGE_FRAGMENT is seen
  // // - 'failed' if a fatal error is encountered, should have a statusText
  // // - '' if the connection closed gracefully by the server, should have a statusText
  // const { attempt: tdpConnection, setAttempt: setTdpConnection } =
  //   useAttempt('processing');
  const [tdpClient, setTdpClient] = useState<TdpClient>(null);

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

  // Default TdpClientEvent.TDP_ERROR and TdpClientEvent.CLIENT_ERROR handler
  const onError = (error: Error) => {
    setDirectorySharingState(defaultDirectorySharingState);
    setClipboardSharingState(defaultClipboardSharingState);
    // should merge this + wsStatus into 1 connection var
    // setTdpConnection(prevState => {
    //   // Sometimes when a connection closes due to an error, we get a cascade of
    //   // errors. Here we update the status only if it's not already 'failed', so
    //   // that the first error message (which is usually the most informative) is
    //   // displayed to the user.
    //   if (prevState.status !== 'failed') {
    //     return {
    //       status: 'failed',
    //       statusText: error.message || error.toString(),
    //     };
    //   }
    //   return prevState;
    // });
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
    setDirectorySharingState(defaultDirectorySharingState);
    setClipboardSharingState(defaultClipboardSharingState);
    // setTdpConnection({
    //   status: '', // gracefully disconnecting
    //   statusText: info,
    // });
  };

  setTdpClient(
    new TdpClient(addr, { onClipboardData, onError, onWarning, onInfo })
  );

  const webauthn = useWebAuthn(tdpClient);

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

  return {
    hostname,
    username,
    clipboardSharingState,
    setClipboardSharingState,
    directorySharingState,
    setDirectorySharingState,
    fetchAttempt,
    // tdpConnection,
    webauthn,
    // setTdpConnection,
    showAnotherSessionActiveDialog,
    setShowAnotherSessionActiveDialog,
    onShareDirectory,
    onCtrlAltDel,
    alerts,
    onRemoveAlert,
    // this shouldn't be spread, but passed as its own object and
    // _then_ spread as props into `TdpCanvas`
    ...clientCanvasProps,
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
