/*
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

/*
The MIT License (MIT)

Copyright (c) 2014 Call-Em-All

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

import React, { createRef, forwardRef, MutableRefObject } from 'react';
import styled, { CSSProp, withTheme } from 'styled-components';

import Modal, { BackdropProps, Props as ModalProps } from '../Modal';
import { Transition } from './Transition';
import { color, ResponsiveValue, ThemeValue } from 'styled-system';
import { Theme } from 'design/theme/themes/types';
import Box, { BoxProps } from 'design/Box';
import { CSSObject } from 'styled-components';
import Flex from 'design/Flex';

type Offset = { top: number; left: number };
type Rect = Offset & { bottom: number; right: number };
type Dimensions = { width: number; height: number };

export type Origin = {
  horizontal: HorizontalOrigin;
  vertical: VerticalOrigin;
};

export type HorizontalOrigin = 'left' | 'center' | 'right' | number;
export type VerticalOrigin = 'top' | 'center' | 'bottom' | number;
export type GrowDirections = 'top-left' | 'bottom-right';
export type Position = 'top' | 'right' | 'bottom' | 'left';

type NumericOrigin = {
  horizontal: number;
  vertical: number;
};

function getOffsetTop(rect: Dimensions, vertical: VerticalOrigin): number {
  let offset = 0;

  if (typeof vertical === 'number') {
    offset = vertical;
  } else if (vertical === 'center') {
    offset = rect.height / 2;
  } else if (vertical === 'bottom') {
    offset = rect.height;
  }

  return offset;
}

function getOffsetLeft(rect: Dimensions, horizontal: HorizontalOrigin): number {
  let offset = 0;

  if (typeof horizontal === 'number') {
    offset = horizontal;
  } else if (horizontal === 'center') {
    offset = rect.width / 2;
  } else if (horizontal === 'right') {
    offset = rect.width;
  }

  return offset;
}

/**
 * Returns popover position, relative to the anchor. If unambiguously defined by
 * the transform origin, returns this one. The ambiguous cases (transform origin
 * on one of the popover corners) are resolved by looking into the anchor
 * origin. If still ambiguous (corner touching a corner), prefers a vertical
 * position.
 */
function getPopoverPosition(
  anchorOrigin: Origin,
  transformOrigin: Origin
): Position | null {
  const allowedByTransformOrigin = getAllowedPopoverPositions(transformOrigin);
  switch (allowedByTransformOrigin.length) {
    case 0:
      return null;
    case 1:
      return allowedByTransformOrigin[0];

    default: {
      const preferredByAnchorOrigin =
        getPreferredPopoverPositions(anchorOrigin);
      const resolved = allowedByTransformOrigin.filter(d =>
        preferredByAnchorOrigin.includes(d)
      );
      if (resolved.length === 0) return null;
      return resolved[0];
    }
  }
}

function getAllowedPopoverPositions(transformOrigin: Origin) {
  const allowed: Position[] = [];
  // Note: order matters here. The first one will be preferred when no
  // unambiguous decision is reached, so we arbitrarily prefer vertical over
  // horizontal arrows.
  if (transformOrigin.vertical === 'top') allowed.push('bottom');
  if (transformOrigin.vertical === 'bottom') allowed.push('top');
  if (transformOrigin.horizontal === 'left') allowed.push('right');
  if (transformOrigin.horizontal === 'right') allowed.push('left');
  return allowed;
}

function getPreferredPopoverPositions(anchorOrigin: Origin) {
  const preferred: Position[] = [];
  if (anchorOrigin.vertical === 'top') preferred.push('top');
  if (anchorOrigin.vertical === 'bottom') preferred.push('bottom');
  if (anchorOrigin.horizontal === 'left') preferred.push('left');
  if (anchorOrigin.horizontal === 'right') preferred.push('right');
  return preferred;
}

function getPopoverMarginTop(
  popoverPos: Position | null,
  arrow: boolean,
  popoverMargin: number
): number {
  const margin = arrow ? popoverMargin + arrowLength : popoverMargin;
  if (popoverPos === 'top') return margin;
  if (popoverPos === 'bottom') return -margin;
  return 0;
}

function getPopoverMarginLeft(
  popoverPos: Position | null,
  arrow: boolean,
  popoverMargin: number
): number {
  const margin = arrow ? popoverMargin + arrowLength : popoverMargin;
  if (popoverPos === 'left') return margin;
  if (popoverPos === 'right') return -margin;
  return 0;
}

function getTransformOriginValue(transformOrigin: Origin): string {
  return [transformOrigin.horizontal, transformOrigin.vertical]
    .map(n => {
      return typeof n === 'number' ? `${n}px` : n;
    })
    .join(' ');
}

// Sum the scrollTop between two elements.
function getScrollParent(parent: Element, child: Element): number {
  let element = child;
  let scrollTop = 0;

  while (element && element !== parent) {
    element = element.parentElement;
    scrollTop += element.scrollTop;
  }
  return scrollTop;
}

function getAnchorEl(anchorEl: Element | (() => Element)): Element {
  return typeof anchorEl === 'function' ? anchorEl() : anchorEl;
}

const arrowLength = 8; //pixels
const arrowWidth = 2 * arrowLength;
const arrowArm = Math.SQRT2 * arrowLength;

/**
 * Constant CSS props of arrows, indexed by the popover position (which is
 * opposite to the arrow direction).
 */
const arrowGeometry = {
  top: {
    width: arrowArm,
    height: arrowArm,
    marginTop: -(arrowArm - arrowWidth),
    // transform: 'rotate(45deg)',
    // clipPath: 'polygon(0% 0%, 100% 0%, 50% 100%)',
  },

  right: {
    width: arrowLength,
    height: arrowWidth,
    clipPath: 'polygon(100% 0%, 100% 100%, 0% 50%)',
  },

  bottom: {
    width: arrowWidth,
    height: arrowArm,
  },

  left: {
    width: arrowLength,
    height: arrowWidth,
    clipPath: 'polygon(0% 0%, 100% 50%, 0% 100%)',
  },
};

class PopoverInternal extends React.Component<ThemedProps> {
  paperRef = createRef<HTMLDivElement>();
  arrowRef = createRef<HTMLDivElement>();
  handleResize: () => void;

  static defaultProps = {
    anchorReference: 'anchorEl',
    anchorOrigin: {
      vertical: 'top',
      horizontal: 'left',
    },
    marginThreshold: 16,
    transformOrigin: {
      vertical: 'top',
      horizontal: 'left',
    },
    growDirections: 'bottom-right',
    arrow: false,
    popoverMargin: 0,
    arrowMargin: 4,
    bg: 'levels.elevated',
  };

  constructor(props: ThemedProps) {
    super(props);

    if (typeof window !== 'undefined') {
      this.handleResize = () => {
        // Because we debounce the event, the open property might no longer be true
        // when the callback resolves.
        if (!this.props.open) {
          return;
        }

        this.setPositioningStyles();
      };
    }
  }

  componentDidMount() {
    if (this.props.action) {
      this.props.action({
        updatePosition: this.handleResize,
      });
    }
  }

  setPositioningStyles = () => {
    const paper = this.paperRef.current;
    const arrow = this.arrowRef.current;
    const { top, left, bottom, right, transformOrigin, arrowStyle } =
      this.getPositioningStyle();

    if (this.props.growDirections === 'bottom-right') {
      if (top !== null) {
        paper.style.top = top;
      }
      if (left !== null) {
        paper.style.left = left;
      }
    } else {
      if (bottom !== null) {
        paper.style.bottom = bottom;
      }
      if (right !== null) {
        paper.style.right = right;
      }
    }
    paper.style.transformOrigin = transformOrigin;
    console.log(arrowStyle);

    if (arrow && arrowStyle) {
      for (const [prop, value] of Object.entries(arrowStyle)) {
        arrow.style[prop] = value;
      }
    }
  };

  getPositioningStyle = () => {
    const element = this.paperRef.current;
    const { anchorReference, marginThreshold, arrowMargin, popoverMargin } =
      this.props;

    // Check if the parent has requested anchoring on an inner content node
    const contentAnchorOffset = this.getContentAnchorOffset(element);
    const elemRect = element.getBoundingClientRect();

    // Get the transform origin point on the element itself
    const transformOrigin = this.getTransformOrigin(
      elemRect,
      popoverMargin,
      contentAnchorOffset
    );

    if (anchorReference === 'none') {
      return {
        top: null,
        left: null,
        transformOrigin: getTransformOriginValue(transformOrigin),
      };
    }

    // Get the offset of of the anchoring element
    const anchorOffset = this.getAnchorOffset(contentAnchorOffset);

    // Calculate element positioning
    let top = anchorOffset.top - transformOrigin.vertical;
    let left = anchorOffset.left - transformOrigin.horizontal;

    // bottom and right correspond to the calculated position of the element from the top left, not
    // from the bottom right, meaning they must be inverted before using them as `bottom` and
    // `right` CSS properties.
    let bottom = top + elemRect.height;
    let right = left + elemRect.width;

    // Window thresholds taking required margin into account
    const heightThreshold = window.innerHeight - marginThreshold;
    const widthThreshold = window.innerWidth - marginThreshold;

    // Check if the vertical axis needs shifting
    if (top < marginThreshold) {
      const diff = top - marginThreshold;
      top -= diff;
      transformOrigin.vertical += diff;
    } else if (bottom > heightThreshold) {
      const diff = bottom - heightThreshold;
      top -= diff;
      transformOrigin.vertical += diff;
    }

    // Check if the horizontal axis needs shifting
    if (left < marginThreshold) {
      const diff = left - marginThreshold;
      left -= diff;
      transformOrigin.horizontal += diff;
    } else if (right > widthThreshold) {
      const diff = right - widthThreshold;
      left -= diff;
      transformOrigin.horizontal += diff;
    }

    bottom = top + elemRect.height;
    right = left + elemRect.width;

    const popoverPos = getPopoverPosition(
      this.props.anchorOrigin,
      this.props.transformOrigin
    );

    const arrowStyle = this.getArrowStyle(
      anchorOffset,
      { left, right, top, bottom },
      popoverPos,
      transformOrigin
    );

    return {
      top: `${top}px`,
      left: `${left}px`,
      bottom: `${window.innerHeight - bottom}px`,
      right: `${window.innerWidth - right}px`,
      transformOrigin: getTransformOriginValue(transformOrigin),
      arrowStyle,
    };
  };

  getArrowStyle = (
    anchorOffset: Offset,
    popoverRect: Rect,
    popoverPos: Position,
    transformOrigin: NumericOrigin
  ): React.CSSProperties | null => {
    const { arrowMargin, popoverMargin, arrow } = this.props;
    if (!arrow) {
      return null;
    }

    const { width, height } = arrowGeometry[popoverPos];

    let left = anchorOffset.left;
    switch (popoverPos) {
      case 'top':
      case 'bottom':
        left -= arrowWidth / 2;
        if (left < popoverRect.left + arrowMargin) {
          left = popoverRect.left + arrowMargin;
        } else if (left > popoverRect.right - arrowWidth - arrowMargin) {
          left = popoverRect.right - arrowWidth - arrowMargin;
        }
        break;
      case 'left':
        left -= arrowLength + popoverMargin;
        break;
      case 'right':
        left += popoverMargin;
        break;
      default:
        popoverPos satisfies never;
    }

    let top = anchorOffset.top;
    switch (popoverPos) {
      case 'left':
      case 'right':
        top -= arrowWidth / 2;
        if (top < popoverRect.top + arrowMargin) {
          top = popoverRect.top + arrowMargin;
        } else if (top > popoverRect.bottom - arrowWidth - arrowMargin) {
          top = popoverRect.bottom - arrowWidth - arrowMargin;
        }
        break;
      case 'top':
        top -= arrowLength + popoverMargin;
        break;
      case 'bottom':
        top += popoverMargin;
        break;
      default:
        popoverPos satisfies never;
    }

    let marginLeft = transformOrigin.horizontal;
    if (marginLeft < arrowMargin) {
      marginLeft = arrowMargin;
    }
    if (marginLeft > popoverRect.right - arrowWidth - arrowMargin) {
      marginLeft = popoverRect.right - arrowWidth - arrowMargin;
    }

    return {
      // left: `${left}px`,
      // top: `${top}px`,
      // width: `${width}px`,
      // height: `${height}px`,
      marginLeft: `${marginLeft}px`,
    };
  };

  // Returns the top/left offset of the position
  // to attach to on the anchor element (or body if none is provided)
  getAnchorOffset(contentAnchorOffset: number): Offset {
    const { anchorEl, anchorOrigin } = this.props;

    // If an anchor element wasn't provided, just use the parent body element of this Popover
    const anchorElement = getAnchorEl(anchorEl) || document.body;

    const anchorRect = anchorElement.getBoundingClientRect();

    const anchorVertical =
      contentAnchorOffset === 0 ? anchorOrigin.vertical : 'center';

    return {
      top: anchorRect.top + getOffsetTop(anchorRect, anchorVertical),
      left:
        anchorRect.left + getOffsetLeft(anchorRect, anchorOrigin.horizontal),
    };
  }

  // Returns the vertical offset of inner content to anchor the transform on if provided
  getContentAnchorOffset(element: HTMLElement): number {
    const { getContentAnchorEl, anchorReference } = this.props;
    let contentAnchorOffset = 0;

    if (getContentAnchorEl && anchorReference === 'anchorEl') {
      const contentAnchorEl = getContentAnchorEl(element);

      if (contentAnchorEl && element.contains(contentAnchorEl)) {
        const scrollTop = getScrollParent(element, contentAnchorEl);
        contentAnchorOffset =
          contentAnchorEl.offsetTop +
            contentAnchorEl.clientHeight / 2 -
            scrollTop || 0;
      }
    }

    return contentAnchorOffset;
  }

  // Return the base transform origin using the element
  // and taking the content anchor offset into account if in use
  getTransformOrigin(
    elemRect: Dimensions,
    popoverMargin: number,
    contentAnchorOffset = 0
  ): NumericOrigin {
    const { transformOrigin, anchorOrigin } = this.props;

    const popoverPos = getPopoverPosition(
      this.props.anchorOrigin,
      this.props.transformOrigin
    );

    const vertical =
      getOffsetTop(elemRect, transformOrigin.vertical) +
      getPopoverMarginTop(
        popoverPos,
        this.props.arrow,
        this.props.popoverMargin
      ) +
      contentAnchorOffset;

    const horizontal =
      getOffsetLeft(elemRect, transformOrigin.horizontal) +
      getPopoverMarginLeft(
        popoverPos,
        this.props.arrow,
        this.props.popoverMargin
      );

    return {
      vertical,
      horizontal,
    };
  }

  handleEntering = () => {
    if (this.props.onEntering) {
      this.props.onEntering(this.paperRef.current);
    }

    this.setPositioningStyles();
  };

  render() {
    const { children, open, popoverCss, arrowCss, ...other } = this.props;
    const popoverPos = getPopoverPosition(
      this.props.anchorOrigin,
      this.props.transformOrigin
    );
    let flexDir;
    switch (popoverPos) {
      case 'top':
        flexDir = 'column-reverse';
        break;
      case 'right':
        flexDir = 'row';
        break;
      case 'bottom':
        flexDir = 'column';
        break;
      case 'left':
        flexDir = 'row-reverse';
        break;
      default:
        popoverPos satisfies never;
    }

    return (
      <Modal
        open={open}
        BackdropProps={{ invisible: true, ...this.props.backdropProps }}
        {...other}
      >
        <Transition onEntering={this.handleEntering}>
          <div>
            <Flex
              ref={this.paperRef}
              flexDirection={flexDir}
              shadow={!this.props.arrow}
              // popoverCss={popoverCss}
              data-mui-test="Popover"
              style={{
                position: 'absolute',
                maxWidth: 'calc(100% - 32px)',
                maxHeight: 'calc(100% - 32px)',
              }}
            >
              {this.props.arrow && (
                <Arrow
                  pos={popoverPos}
                  ref={this.arrowRef}
                  /*bg={this.props.bg} css={arrowCss}*/
                />
              )}
              <StyledPopover shadow={!this.props.arrow} bg={this.props.bg}>
                {children}
              </StyledPopover>
            </Flex>
          </div>
        </Transition>
      </Modal>
    );
  }
}

const Popover = withTheme(PopoverInternal) as React.ComponentType<Props>;
export default Popover;

const Arrow = forwardRef(({ pos, ...rest }, ref) => {
  const { width, height, marginTop } = arrowGeometry[pos];
  return (
    <Box
      ref={ref}
      style={{
        width,
        height,
        overflow: 'hidden',
        background: 'green',
        alignSelf: 'start',
      }}
      {...rest}
    >
      <div
        style={{
          background: 'red',
          width: '100%',
          height: `${height}px`,
          marginTop: `${marginTop}px`,
          transformOrigin: 'bottom left',
          transform: 'rotate(45deg)',
        }}
      ></div>
    </Box>
  );
});

interface Props extends Omit<ModalProps, 'children' | 'open'> {
  /**
   * This is callback property. It's called by the component on mount.  This is
   * useful when you want to trigger an action programmatically.  It currently
   * only supports updatePosition() action.
   *
   * @param actions This object contains all possible actions that can be
   * triggered programmatically.
   */
  action?: (actions: { updatePosition: () => void }) => void;

  /**
   * This is the DOM element, or a function that returns the DOM element, that
   * may be used to set the position of the popover.
   */
  anchorEl?: Element | (() => Element);

  /**
   * This is the point on the anchor where the popover's `anchorEl` will attach
   * to.
   */
  anchorOrigin?: Origin;

  /**
   * These are the directions in which `Popover` will grow if its content
   * increases its dimensions after `Popover` is opened.
   */
  growDirections?: GrowDirections;

  /**
   * This determines which anchor prop to refer to to set the position of the
   * popover.
   */
  anchorReference?: 'anchorEl' | 'none';

  /**
   * The content of the component.
   */
  children?: React.ReactNode;

  /**
   * This function is called in order to retrieve the content anchor element.
   * It's the opposite of the `anchorEl` property.  The content anchor element
   * should be an element inside the popover.  It's used to correctly scroll and
   * set the position of the popover.  The positioning strategy tries to make
   * the content anchor element just above the anchor element.
   */
  getContentAnchorEl?: (paperElement: HTMLElement) => HTMLElement;

  /**
   * Specifies how close to the edge of the window the popover can appear.
   */
  marginThreshold?: number;

  /**
   * Callback fired when the component requests to be closed.
   *
   * @param event The event source of the callback.
   * @param reason Can be:`"escapeKeyDown"`, `"backdropClick"`
   */
  onClose?: (
    event: React.MouseEvent | KeyboardEvent,
    reason: 'escapeKeyDown' | 'backdropClick'
  ) => void;

  /**
   * Callback fired when the component is entering.
   */
  onEntering?: (paperElement: HTMLElement) => void;

  /**
   * If `true`, the popover is visible.
   */
  open: boolean;

  /**
   * This is the point on the popover which will attach to the anchor's origin.
   *
   * Options:
   * vertical: [top, center, bottom, x(px)];
   * horizontal: [left, center, right, x(px)].
   */
  transformOrigin?: Origin;

  /** Returns additional styles applied to the internal popover element. */
  popoverCss?: () => CSSProp;

  /** Properties applied to the backdrop element. */
  backdropProps?: BackdropProps;

  arrow?: boolean;

  arrowCss?: () => CSSProp;

  popoverMargin?: number;

  arrowMargin?: number;

  bg?: BoxProps['bg'];
}

interface ThemedProps extends Props {
  theme: Theme;
}

export const StyledPopover = styled(Flex)<{
  shadow: boolean;
  popoverCss?: () => CSSProp;
}>`
  box-shadow: ${props => (props.shadow ? props.theme.boxShadow[1] : 'none')};
  border-radius: 4px;
  min-height: 16px;
  min-width: 16px;
  outline: none;
  overflow-x: hidden;
  overflow-y: auto;
  ${props => props.popoverCss && props.popoverCss()}
`;
