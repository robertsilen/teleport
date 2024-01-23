/*
Copyright 2019 Gravitational, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

import React from 'react';

import { FailedApp } from 'teleterm/ui/components/App';
import Logger from 'teleterm/logger';

export class CatchError extends React.Component {
  logger = new Logger('components/CatchError');

  static getDerivedStateFromError(error) {
    return { error };
  }

  state = {
    error: null,
  };

  componentDidCatch(err) {
    this.logger.error('render', err);
  }

  render() {
    if (this.state.error) {
      return (
        <FailedApp message={this.state.error?.message || this.state.error} />
      );
    }

    return this.props.children;
  }
}
