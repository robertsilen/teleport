/**
 * Copyright 2023 Gravitational, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import styled from 'styled-components';

import { AwsOidc } from '../AwsOidc';

export interface CommonInstructionsProps extends PreviousStepProps {
  onNext: (updatedAwsOidc?: AwsOidc) => void;
  clusterPublicUri: string;
}

export interface PreviousStepProps {
  onPrev: (updatedAwsOidc?: AwsOidc) => void;
  awsOidc?: AwsOidc;
}

export const InstructionsContainer = styled.div`
  flex: 0 0 600px;
  padding-right: 100px;
`;