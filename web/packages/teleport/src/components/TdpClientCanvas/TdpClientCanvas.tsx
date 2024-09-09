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

import React, { memo, useEffect } from 'react';
import { DebouncedFunc } from 'shared/utils/highbar';

import type { CSSProperties, MutableRefObject } from 'react';

function TdpClientCanvas(props: Props) {
  const {
    canvasRef,
    onKeyDown,
    onKeyUp,
    onFocusOut,
    onMouseMove,
    onMouseDown,
    onMouseUp,
    onMouseWheelScroll,
    windowOnResize,
    onContextMenu,
    style,
  } = props;

  useEffect(() => {
    // Empty dependency array ensures this runs only once after initial render.
    // This code will run after the component has been mounted and the canvasRef has been assigned.
    const canvas = canvasRef.current;
    if (canvas) {
      // Make the canvas a focusable keyboard listener
      // https://stackoverflow.com/a/51267699/6277051
      // https://stackoverflow.com/a/16492878/6277051
      canvas.tabIndex = -1;
      canvas.style.outline = 'none';
      canvas.focus();
    }
  }, [canvasRef]);

  useEffect(() => {
    const canvas = canvasRef.current;
    if (!canvas) {
      return;
    }
    window.addEventListener('resize', windowOnResize);
    canvas.addEventListener('onmousemove', onMouseMove);
    canvas.addEventListener('oncontextmenu', onContextMenu);
    canvas.addEventListener('onmousedown', onMouseDown);
    canvas.addEventListener('onmouseup', onMouseUp);
    canvas.addEventListener('onwheel', onMouseWheelScroll);
    canvas.addEventListener('onkeydown', onKeyDown);
    canvas.addEventListener('onkeyup', onKeyUp);
    canvas.addEventListener('focusout', onFocusOut);

    return () => {
      window.removeEventListener('resize', windowOnResize);
      canvas.removeEventListener('mousemove', onMouseMove);
      canvas.removeEventListener('contextmenu', onContextMenu);
      canvas.removeEventListener('mousedown', onMouseDown);
      canvas.removeEventListener('mouseup', onMouseUp);
      canvas.removeEventListener('wheel', onMouseWheelScroll);
      canvas.removeEventListener('keydown', onKeyDown);
      canvas.removeEventListener('keyup', onKeyUp);
      canvas.removeEventListener('focusout', onFocusOut);
    };
  }, [canvasRef]);

  // useEffect(() => {
  //   if (client) {
  //     const canvas = canvasRef.current;
  //     const _clearCanvas = () => {
  //       const ctx = canvas.getContext('2d');
  //       ctx.clearRect(0, 0, canvas.width, canvas.height);
  //     };
  //     client.on(TdpClientEvent.RESET, _clearCanvas);

  //     return () => {
  //       client.removeListener(TdpClientEvent.RESET, _clearCanvas);
  //     };
  //   }
  // }, [client]);

  return <canvas style={{ ...style }} ref={canvasRef} />;
}

export type Props = {
  canvasRef: MutableRefObject<HTMLCanvasElement>;
  onKeyDown?: (e: KeyboardEvent) => any;
  onKeyUp?: (e: KeyboardEvent) => any;
  onFocusOut?: () => any;
  onMouseMove?: (e: MouseEvent) => any;
  onMouseDown?: (e: MouseEvent) => any;
  onMouseUp?: (e: MouseEvent) => any;
  onMouseWheelScroll?: (e: WheelEvent) => any;
  onContextMenu?: () => boolean;
  windowOnResize?: DebouncedFunc<() => void>;
  style?: CSSProperties;
  updatePointer?: boolean;
};

export default memo(TdpClientCanvas);
