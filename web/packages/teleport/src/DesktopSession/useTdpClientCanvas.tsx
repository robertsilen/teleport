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

import { useEffect, useRef } from 'react';

import { TdpClient, ButtonState, ScrollAxis } from 'teleport/lib/tdp';
import { ClientScreenSpec } from 'teleport/lib/tdp/codec';

import { KeyboardHandler } from './KeyboardHandler';

declare global {
  interface Navigator {
    userAgentData?: { platform: any };
  }
}

export default function useTdpClientCanvas(cli: TdpClient) {
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

  const canvasRef = useRef<HTMLCanvasElement>(null);

  // this should be moved into part of wsStatus probably.
  // really, the only thing its doing is tracking when we've received
  // the first frame to know "hey im connected", but perhaps we should
  // rename it/move it to better track what we are trying to do
  // const initialTdpConnectionSucceeded = useRef(false);
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
  const syncCanvas = (spec: ClientScreenSpec) => {
    const canvas = canvasRef.current;
    if (canvas) {
      return;
    }
    const { width, height } = spec;
    canvas.width = width;
    canvas.height = height;
    canvas.style.width = `${width}px`;
    canvas.style.height = `${height}px`;
  };

  const onKeyDown = (e: KeyboardEvent) => {
    if (!cli) {
      return;
    }
    keyboardHandler.current.handleKeyboardEvent({
      cli,
      e,
      state: ButtonState.DOWN,
    });

    // TODO (avatus): figure where to call this in client data

    // // The key codes in the if clause below are those that have been empirically determined not
    // // to count as transient activation events. According to the documentation, a keydown for
    // // the Esc key and any "shortcut key reserved by the user agent" don't count as activation
    // // events: https://developer.mozilla.org/en-US/docs/Web/Security/User_activation.
    // if (e.key !== 'Meta' && e.key !== 'Alt' && e.key !== 'Escape') {
    //   onKeyDown();
    // }
  };

  const onKeyUp = (e: KeyboardEvent) => {
    if (!cli) {
      return;
    }
    keyboardHandler.current.handleKeyboardEvent({
      cli,
      e,
      state: ButtonState.UP,
    });
  };

  const onFocusOut = () => {
    keyboardHandler.current.onFocusOut();
  };

  const onMouseMove = (e: MouseEvent) => {
    const canvas = canvasRef.current;
    if (!cli || !canvas) {
      return;
    }
    const rect = canvas.getBoundingClientRect();
    const x = e.clientX - rect.left;
    const y = e.clientY - rect.top;
    cli.sendMouseMove(x, y);
  };

  const onMouseDown = (e: MouseEvent) => {
    if (!cli) {
      return;
    }
    if (e.button === 0 || e.button === 1 || e.button === 2) {
      cli.sendMouseButton(e.button, ButtonState.DOWN);
    }

    // TODO (avatus) : figure out where to call this in client data
    // // Opportunistically sync local clipboard to remote while
    // // transient user activation is in effect.
    // // https://developer.mozilla.org/en-US/docs/Web/API/Clipboard/readText#security
    // sendLocalClipboardToRemote(cli);
  };

  const onMouseUp = (e: MouseEvent) => {
    if (!cli) {
      return;
    }
    if (e.button === 0 || e.button === 1 || e.button === 2) {
      cli.sendMouseButton(e.button, ButtonState.UP);
    }
  };

  const onMouseWheelScroll = (e: WheelEvent) => {
    if (!cli) {
      return;
    }
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
  const onContextMenu = () => false;

  return {
    syncCanvas,
    canvasRef,
    onKeyDown,
    onKeyUp,
    onFocusOut,
    onMouseMove,
    onMouseDown,
    onMouseUp,
    onMouseWheelScroll,
    onContextMenu,
  };
}

// type Props = {
//   username: string;
//   desktopName: string;
//   clusterId: string;
//   setTdpConnection: Setter<Attempt>;
//   clipboardSharingState: ClipboardSharingState;
//   setClipboardSharingState: Setter<ClipboardSharingState>;
//   setDirectorySharingState: Setter<DirectorySharingState>;
//   setAlerts: Setter<NotificationItem[]>;
// };
