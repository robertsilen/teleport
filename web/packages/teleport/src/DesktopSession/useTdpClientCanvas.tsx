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

import { useState, useEffect, useRef } from 'react';
import { Attempt } from 'shared/hooks/useAttemptNext';
import { NotificationItem } from 'shared/components/Notification';
import { debounce } from 'shared/utils/highbar';

import { TdpClient, ButtonState, ScrollAxis } from 'teleport/lib/tdp';
import {
  ClientScreenSpec,
  ClipboardData,
  PngFrame,
} from 'teleport/lib/tdp/codec';
import { Sha256Digest } from 'teleport/lib/util';

import { TopBarHeight } from './TopBar';
import {
  ClipboardSharingState,
  DirectorySharingState,
  Setter,
  clipboardSharingPossible,
  defaultClipboardSharingState,
  defaultDirectorySharingState,
  isSharingClipboard,
} from './useDesktopSession';
import { KeyboardHandler } from './KeyboardHandler';

import type { BitmapFrame } from 'teleport/lib/tdp/client';

declare global {
  interface Navigator {
    userAgentData?: { platform: any };
  }
}

export default function useTdpClientCanvas() {
  // const {
  //   username,
  //   desktopName,
  //   clusterId,
  //   setTdpConnection,
  //   clipboardSharingState,
  //   setClipboardSharingState,
  //   setDirectorySharingState,
  //   setAlerts,
  // } = props;

  // this should be moved into part of wsStatus probably.
  // really, the only thing its doing is tracking when we've received
  // the first frame to know "hey im connected", but perhaps we should
  // rename it/move it to better track what we are trying to do
  const initialTdpConnectionSucceeded = useRef(false);
  const keyboardHandler = useRef(new KeyboardHandler());

  useEffect(() => {
    keyboardHandler.current = new KeyboardHandler();
    // On unmount, clear all the timeouts on the keyboardHandler.
    return () => {
      // eslint-disable-next-line react-hooks/exhaustive-deps
      keyboardHandler.current.dispose();
    };
  }, []);

  /**
   * Synchronize the canvas resolution and display size with the
   * given ClientScreenSpec.
   */
  const syncCanvas = (canvas: HTMLCanvasElement, spec: ClientScreenSpec) => {
    const { width, height } = spec;
    canvas.width = width;
    canvas.height = height;
    canvas.style.width = `${width}px`;
    canvas.style.height = `${height}px`;
  };

  // Default TdpClientEvent.TDP_PNG_FRAME handler (buffered)
  const clientOnPngFrame = (
    ctx: CanvasRenderingContext2D,
    pngFrame: PngFrame
  ) => {
    // The first image fragment we see signals a successful TDP connection.
    if (!initialTdpConnectionSucceeded.current) {
      syncCanvas(ctx.canvas, getDisplaySize());
      // setTdpConnection({ status: 'success' });
      initialTdpConnectionSucceeded.current = true;
    }
    ctx.drawImage(pngFrame.data, pngFrame.left, pngFrame.top);
  };

  // Default TdpClientEvent.TDP_BMP_FRAME handler (buffered)
  const clientOnBitmapFrame = (
    ctx: CanvasRenderingContext2D,
    bmpFrame: BitmapFrame
  ) => {
    // The first image fragment we see signals a successful TDP connection.
    if (!initialTdpConnectionSucceeded.current) {
      // setTdpConnection({ status: 'success' });
      initialTdpConnectionSucceeded.current = true;
    }
    ctx.putImageData(bmpFrame.image_data, bmpFrame.left, bmpFrame.top);
  };

  // Default TdpClientEvent.TDP_CLIENT_SCREEN_SPEC handler.
  const clientOnClientScreenSpec = (
    cli: TdpClient,
    canvas: HTMLCanvasElement,
    spec: ClientScreenSpec
  ) => {
    syncCanvas(canvas, spec);
  };

  const canvasOnKeyDown = (cli: TdpClient, e: KeyboardEvent) => {
    keyboardHandler.current.handleKeyboardEvent({
      cli,
      e,
      state: ButtonState.DOWN,
    });

    // The key codes in the if clause below are those that have been empirically determined not
    // to count as transient activation events. According to the documentation, a keydown for
    // the Esc key and any "shortcut key reserved by the user agent" don't count as activation
    // events: https://developer.mozilla.org/en-US/docs/Web/Security/User_activation.
    if (e.key !== 'Meta' && e.key !== 'Alt' && e.key !== 'Escape') {
      // Opportunistically sync local clipboard to remote while
      // transient user activation is in effect.
      // https://developer.mozilla.org/en-US/docs/Web/API/Clipboard/readText#security
      sendLocalClipboardToRemote(cli);
    }
  };

  const canvasOnKeyUp = (cli: TdpClient, e: KeyboardEvent) => {
    keyboardHandler.current.handleKeyboardEvent({
      cli,
      e,
      state: ButtonState.UP,
    });
  };

  const canvasOnFocusOut = () => {
    keyboardHandler.current.onFocusOut();
  };

  const canvasOnMouseMove = (
    cli: TdpClient,
    canvas: HTMLCanvasElement,
    e: MouseEvent
  ) => {
    const rect = canvas.getBoundingClientRect();
    const x = e.clientX - rect.left;
    const y = e.clientY - rect.top;
    cli.sendMouseMove(x, y);
  };

  const canvasOnMouseDown = (cli: TdpClient, e: MouseEvent) => {
    if (e.button === 0 || e.button === 1 || e.button === 2) {
      cli.sendMouseButton(e.button, ButtonState.DOWN);
    }

    // Opportunistically sync local clipboard to remote while
    // transient user activation is in effect.
    // https://developer.mozilla.org/en-US/docs/Web/API/Clipboard/readText#security
    sendLocalClipboardToRemote(cli);
  };

  const canvasOnMouseUp = (cli: TdpClient, e: MouseEvent) => {
    if (e.button === 0 || e.button === 1 || e.button === 2) {
      cli.sendMouseButton(e.button, ButtonState.UP);
    }
  };

  const canvasOnMouseWheelScroll = (cli: TdpClient, e: WheelEvent) => {
    e.preventDefault();
    // We only support pixel scroll events, not line or page events.
    // https://developer.mozilla.org/en-US/docs/Web/API/WheelEvent/deltaMode
    if (e.deltaMode === WheelEvent.DOM_DELTA_PIXEL) {
      if (e.deltaX) {
        cli.sendMouseWheelScroll(ScrollAxis.HORIZONTAL, -e.deltaX);
      }
      if (e.deltaY) {
        cli.sendMouseWheelScroll(ScrollAxis.VERTICAL, -e.deltaY);
      }
    }
  };

  // Block browser context menu so as not to obscure the context menu
  // on the remote machine.
  const canvasOnContextMenu = () => false;

  const windowOnResize = debounce(
    (cli: TdpClient) => {
      const spec = getDisplaySize();
      cli.resize(spec);
    },
    250,
    { trailing: true }
  );

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

  return {
    tdpClient,
    clientScreenSpecToRequest: getDisplaySize(),
    clientOnPngFrame,
    clientOnBitmapFrame,
    clientOnClientScreenSpec,
    clientOnTdpError,
    clientOnClipboardData,
    clientOnTdpWarning,
    clientOnTdpInfo,
    canvasOnKeyDown,
    canvasOnKeyUp,
    canvasOnFocusOut,
    canvasOnMouseMove,
    canvasOnMouseDown,
    canvasOnMouseUp,
    canvasOnMouseWheelScroll,
    canvasOnContextMenu,
    windowOnResize,
  };
}

// Calculates the size (in pixels) of the display.
// Since we want to maximize the display size for the user, this is simply
// the full width of the screen and the full height sans top bar.
function getDisplaySize() {
  return {
    width: window.innerWidth,
    height: window.innerHeight - TopBarHeight,
  };
}

type Props = {
  username: string;
  desktopName: string;
  clusterId: string;
  setTdpConnection: Setter<Attempt>;
  clipboardSharingState: ClipboardSharingState;
  setClipboardSharingState: Setter<ClipboardSharingState>;
  setDirectorySharingState: Setter<DirectorySharingState>;
  setAlerts: Setter<NotificationItem[]>;
};
