// Copyright 2021-2022 Gravitational, Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: teleport/samlidp/v1/samlidp.proto

package samlidpv1

import (
	proto "github.com/gravitational/teleport/api/client/proto"
	types "github.com/gravitational/teleport/api/types"
	wrappers "github.com/gravitational/teleport/api/types/wrappers"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// ProcessSAMLIdPRequestRequest is a request to create and sign the SAML IdP response
// to a SAML IdP auth request.
type ProcessSAMLIdPRequestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// destination is the destination of the response.
	Destination string `protobuf:"bytes,1,opt,name=destination,proto3" json:"destination,omitempty"`
	// request_id is the request ID.
	RequestId string `protobuf:"bytes,2,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	// request_time is the time the request was made.
	RequestTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=request_time,json=requestTime,proto3" json:"request_time,omitempty"`
	// Metadata_url is the metadata URL of the SAML IdP.
	MetadataUrl string `protobuf:"bytes,4,opt,name=metadata_url,json=metadataUrl,proto3" json:"metadata_url,omitempty"`
	// signature_method is the signature method to use.
	//
	// Deprecated: the server will choose the signature method.
	//
	// Deprecated: Marked as deprecated in teleport/samlidp/v1/samlidp.proto.
	SignatureMethod string `protobuf:"bytes,5,opt,name=signature_method,json=signatureMethod,proto3" json:"signature_method,omitempty"`
	// assertion is the SAML assertion to sign.
	Assertion []byte `protobuf:"bytes,6,opt,name=assertion,proto3" json:"assertion,omitempty"`
	// service_provider_sso_descriptor is the raw bytes of the service provider's SSO descriptor.
	ServiceProviderSsoDescriptor []byte `protobuf:"bytes,7,opt,name=service_provider_sso_descriptor,json=serviceProviderSsoDescriptor,proto3" json:"service_provider_sso_descriptor,omitempty"`
	// mfa_response is an mfa challenge response used to verify the user.
	MfaResponse *proto.MFAAuthenticateResponse `protobuf:"bytes,8,opt,name=mfa_response,json=mfaResponse,proto3" json:"mfa_response,omitempty"`
}

func (x *ProcessSAMLIdPRequestRequest) Reset() {
	*x = ProcessSAMLIdPRequestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_samlidp_v1_samlidp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProcessSAMLIdPRequestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessSAMLIdPRequestRequest) ProtoMessage() {}

func (x *ProcessSAMLIdPRequestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_samlidp_v1_samlidp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessSAMLIdPRequestRequest.ProtoReflect.Descriptor instead.
func (*ProcessSAMLIdPRequestRequest) Descriptor() ([]byte, []int) {
	return file_teleport_samlidp_v1_samlidp_proto_rawDescGZIP(), []int{0}
}

func (x *ProcessSAMLIdPRequestRequest) GetDestination() string {
	if x != nil {
		return x.Destination
	}
	return ""
}

func (x *ProcessSAMLIdPRequestRequest) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *ProcessSAMLIdPRequestRequest) GetRequestTime() *timestamppb.Timestamp {
	if x != nil {
		return x.RequestTime
	}
	return nil
}

func (x *ProcessSAMLIdPRequestRequest) GetMetadataUrl() string {
	if x != nil {
		return x.MetadataUrl
	}
	return ""
}

// Deprecated: Marked as deprecated in teleport/samlidp/v1/samlidp.proto.
func (x *ProcessSAMLIdPRequestRequest) GetSignatureMethod() string {
	if x != nil {
		return x.SignatureMethod
	}
	return ""
}

func (x *ProcessSAMLIdPRequestRequest) GetAssertion() []byte {
	if x != nil {
		return x.Assertion
	}
	return nil
}

func (x *ProcessSAMLIdPRequestRequest) GetServiceProviderSsoDescriptor() []byte {
	if x != nil {
		return x.ServiceProviderSsoDescriptor
	}
	return nil
}

func (x *ProcessSAMLIdPRequestRequest) GetMfaResponse() *proto.MFAAuthenticateResponse {
	if x != nil {
		return x.MfaResponse
	}
	return nil
}

// ProcessSAMLIdPRequestResponse is a response to processing the SAML IdP auth request.
type ProcessSAMLIdPRequestResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// response is the SAML response.
	Response []byte `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *ProcessSAMLIdPRequestResponse) Reset() {
	*x = ProcessSAMLIdPRequestResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_samlidp_v1_samlidp_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProcessSAMLIdPRequestResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessSAMLIdPRequestResponse) ProtoMessage() {}

func (x *ProcessSAMLIdPRequestResponse) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_samlidp_v1_samlidp_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessSAMLIdPRequestResponse.ProtoReflect.Descriptor instead.
func (*ProcessSAMLIdPRequestResponse) Descriptor() ([]byte, []int) {
	return file_teleport_samlidp_v1_samlidp_proto_rawDescGZIP(), []int{1}
}

func (x *ProcessSAMLIdPRequestResponse) GetResponse() []byte {
	if x != nil {
		return x.Response
	}
	return nil
}

// TestSAMLIdPAttributeMappingRequest is a request to test attribute mapping.
type TestSAMLIdPAttributeMappingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// service_provider is a SAML service provider with attribute mapping.
	ServiceProvider *types.SAMLIdPServiceProviderV1 `protobuf:"bytes,1,opt,name=service_provider,json=serviceProvider,proto3" json:"service_provider,omitempty"`
	// users is a list of users whose details will be used
	// to evaluate attribute mapping.
	Users []*types.UserV2 `protobuf:"bytes,2,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *TestSAMLIdPAttributeMappingRequest) Reset() {
	*x = TestSAMLIdPAttributeMappingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_samlidp_v1_samlidp_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestSAMLIdPAttributeMappingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestSAMLIdPAttributeMappingRequest) ProtoMessage() {}

func (x *TestSAMLIdPAttributeMappingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_samlidp_v1_samlidp_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestSAMLIdPAttributeMappingRequest.ProtoReflect.Descriptor instead.
func (*TestSAMLIdPAttributeMappingRequest) Descriptor() ([]byte, []int) {
	return file_teleport_samlidp_v1_samlidp_proto_rawDescGZIP(), []int{2}
}

func (x *TestSAMLIdPAttributeMappingRequest) GetServiceProvider() *types.SAMLIdPServiceProviderV1 {
	if x != nil {
		return x.ServiceProvider
	}
	return nil
}

func (x *TestSAMLIdPAttributeMappingRequest) GetUsers() []*types.UserV2 {
	if x != nil {
		return x.Users
	}
	return nil
}

// TestSAMLIdPAttributeMappingResponse is a response to attribute mapping test request.
type TestSAMLIdPAttributeMappingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// mapped_attributes is the result of attribute mapping evaluation.
	MappedAttributes []*MappedAttribute `protobuf:"bytes,1,rep,name=mapped_attributes,json=mappedAttributes,proto3" json:"mapped_attributes,omitempty"`
}

func (x *TestSAMLIdPAttributeMappingResponse) Reset() {
	*x = TestSAMLIdPAttributeMappingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_samlidp_v1_samlidp_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestSAMLIdPAttributeMappingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestSAMLIdPAttributeMappingResponse) ProtoMessage() {}

func (x *TestSAMLIdPAttributeMappingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_samlidp_v1_samlidp_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestSAMLIdPAttributeMappingResponse.ProtoReflect.Descriptor instead.
func (*TestSAMLIdPAttributeMappingResponse) Descriptor() ([]byte, []int) {
	return file_teleport_samlidp_v1_samlidp_proto_rawDescGZIP(), []int{3}
}

func (x *TestSAMLIdPAttributeMappingResponse) GetMappedAttributes() []*MappedAttribute {
	if x != nil {
		return x.MappedAttributes
	}
	return nil
}

// MappedAttribute is a result of attribute mapping with username
// of a user whose username, role and traits are used for evaluation.
type MappedAttribute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// username is username of user whose detail is used for attribute mapping.
	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	// mapped_values is a result of attribute mapping where key is requested
	// attribute name and value is result of evaluated predicate expression.
	MappedValues map[string]*wrappers.StringValues `protobuf:"bytes,2,rep,name=mapped_values,json=mappedValues,proto3" json:"mapped_values,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *MappedAttribute) Reset() {
	*x = MappedAttribute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_samlidp_v1_samlidp_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MappedAttribute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MappedAttribute) ProtoMessage() {}

func (x *MappedAttribute) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_samlidp_v1_samlidp_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MappedAttribute.ProtoReflect.Descriptor instead.
func (*MappedAttribute) Descriptor() ([]byte, []int) {
	return file_teleport_samlidp_v1_samlidp_proto_rawDescGZIP(), []int{4}
}

func (x *MappedAttribute) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *MappedAttribute) GetMappedValues() map[string]*wrappers.StringValues {
	if x != nil {
		return x.MappedValues
	}
	return nil
}

var File_teleport_samlidp_v1_samlidp_proto protoreflect.FileDescriptor

var file_teleport_samlidp_v1_samlidp_proto_rawDesc = []byte{
	0x0a, 0x21, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x73, 0x61, 0x6d, 0x6c, 0x69,
	0x64, 0x70, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x61, 0x6d, 0x6c, 0x69, 0x64, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x13, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x73, 0x61,
	0x6d, 0x6c, 0x69, 0x64, 0x70, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x2f, 0x6c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x74, 0x65, 0x6c, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x2f, 0x6c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2d, 0x74, 0x65,
	0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x6c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x2f, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2f, 0x77, 0x72, 0x61,
	0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x98, 0x03, 0x0a, 0x1c,
	0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x53, 0x41, 0x4d, 0x4c, 0x49, 0x64, 0x50, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d,
	0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x3d, 0x0a,
	0x0c, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x0b, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x55, 0x72, 0x6c, 0x12,
	0x2d, 0x0a, 0x10, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x6d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x02, 0x18, 0x01, 0x52, 0x0f, 0x73,
	0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x1c,
	0x0a, 0x09, 0x61, 0x73, 0x73, 0x65, 0x72, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x09, 0x61, 0x73, 0x73, 0x65, 0x72, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x45, 0x0a, 0x1f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x5f, 0x73, 0x73, 0x6f, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x1c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x53, 0x73, 0x6f, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x6f, 0x72, 0x12, 0x41, 0x0a, 0x0c, 0x6d, 0x66, 0x61, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x4d, 0x46, 0x41, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0b, 0x6d, 0x66, 0x61, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3b, 0x0a, 0x1d, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x53, 0x41, 0x4d, 0x4c, 0x49, 0x64, 0x50, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x95, 0x01, 0x0a, 0x22, 0x54, 0x65, 0x73, 0x74, 0x53, 0x41, 0x4d, 0x4c,
	0x49, 0x64, 0x50, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x4d, 0x61, 0x70, 0x70,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4a, 0x0a, 0x10, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x53, 0x41, 0x4d,
	0x4c, 0x49, 0x64, 0x50, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x56, 0x31, 0x52, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x56, 0x32, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x22, 0x78, 0x0a, 0x23, 0x54,
	0x65, 0x73, 0x74, 0x53, 0x41, 0x4d, 0x4c, 0x49, 0x64, 0x50, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x51, 0x0a, 0x11, 0x6d, 0x61, 0x70, 0x70, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e,
	0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x73, 0x61, 0x6d, 0x6c, 0x69, 0x64, 0x70,
	0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x70, 0x70, 0x65, 0x64, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x52, 0x10, 0x6d, 0x61, 0x70, 0x70, 0x65, 0x64, 0x41, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x73, 0x22, 0xe3, 0x01, 0x0a, 0x0f, 0x4d, 0x61, 0x70, 0x70, 0x65, 0x64,
	0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x5b, 0x0a, 0x0d, 0x6d, 0x61, 0x70, 0x70, 0x65, 0x64, 0x5f,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x74,
	0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x73, 0x61, 0x6d, 0x6c, 0x69, 0x64, 0x70, 0x2e,
	0x76, 0x31, 0x2e, 0x4d, 0x61, 0x70, 0x70, 0x65, 0x64, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x2e, 0x4d, 0x61, 0x70, 0x70, 0x65, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x6d, 0x61, 0x70, 0x70, 0x65, 0x64, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x73, 0x1a, 0x57, 0x0a, 0x11, 0x4d, 0x61, 0x70, 0x70, 0x65, 0x64, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2c, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x77, 0x72, 0x61, 0x70, 0x70,
	0x65, 0x72, 0x73, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0xa3, 0x02, 0x0a, 0x0e,
	0x53, 0x41, 0x4d, 0x4c, 0x49, 0x64, 0x50, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x7e,
	0x0a, 0x15, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x53, 0x41, 0x4d, 0x4c, 0x49, 0x64, 0x50,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x2e, 0x73, 0x61, 0x6d, 0x6c, 0x69, 0x64, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x53, 0x41, 0x4d, 0x4c, 0x49, 0x64, 0x50, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x74, 0x65, 0x6c,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x73, 0x61, 0x6d, 0x6c, 0x69, 0x64, 0x70, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x53, 0x41, 0x4d, 0x4c, 0x49, 0x64, 0x50, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x90,
	0x01, 0x0a, 0x1b, 0x54, 0x65, 0x73, 0x74, 0x53, 0x41, 0x4d, 0x4c, 0x49, 0x64, 0x50, 0x41, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x12, 0x37,
	0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x73, 0x61, 0x6d, 0x6c, 0x69, 0x64,
	0x70, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x53, 0x41, 0x4d, 0x4c, 0x49, 0x64, 0x50,
	0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x38, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x2e, 0x73, 0x61, 0x6d, 0x6c, 0x69, 0x64, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65,
	0x73, 0x74, 0x53, 0x41, 0x4d, 0x4c, 0x49, 0x64, 0x50, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x52, 0x5a, 0x50, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x67, 0x72, 0x61, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65,
	0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x2f, 0x73, 0x61, 0x6d, 0x6c, 0x69, 0x64, 0x70, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x61, 0x6d, 0x6c,
	0x69, 0x64, 0x70, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_teleport_samlidp_v1_samlidp_proto_rawDescOnce sync.Once
	file_teleport_samlidp_v1_samlidp_proto_rawDescData = file_teleport_samlidp_v1_samlidp_proto_rawDesc
)

func file_teleport_samlidp_v1_samlidp_proto_rawDescGZIP() []byte {
	file_teleport_samlidp_v1_samlidp_proto_rawDescOnce.Do(func() {
		file_teleport_samlidp_v1_samlidp_proto_rawDescData = protoimpl.X.CompressGZIP(file_teleport_samlidp_v1_samlidp_proto_rawDescData)
	})
	return file_teleport_samlidp_v1_samlidp_proto_rawDescData
}

var file_teleport_samlidp_v1_samlidp_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_teleport_samlidp_v1_samlidp_proto_goTypes = []any{
	(*ProcessSAMLIdPRequestRequest)(nil),        // 0: teleport.samlidp.v1.ProcessSAMLIdPRequestRequest
	(*ProcessSAMLIdPRequestResponse)(nil),       // 1: teleport.samlidp.v1.ProcessSAMLIdPRequestResponse
	(*TestSAMLIdPAttributeMappingRequest)(nil),  // 2: teleport.samlidp.v1.TestSAMLIdPAttributeMappingRequest
	(*TestSAMLIdPAttributeMappingResponse)(nil), // 3: teleport.samlidp.v1.TestSAMLIdPAttributeMappingResponse
	(*MappedAttribute)(nil),                     // 4: teleport.samlidp.v1.MappedAttribute
	nil,                                         // 5: teleport.samlidp.v1.MappedAttribute.MappedValuesEntry
	(*timestamppb.Timestamp)(nil),               // 6: google.protobuf.Timestamp
	(*proto.MFAAuthenticateResponse)(nil),       // 7: proto.MFAAuthenticateResponse
	(*types.SAMLIdPServiceProviderV1)(nil),      // 8: types.SAMLIdPServiceProviderV1
	(*types.UserV2)(nil),                        // 9: types.UserV2
	(*wrappers.StringValues)(nil),               // 10: wrappers.StringValues
}
var file_teleport_samlidp_v1_samlidp_proto_depIdxs = []int32{
	6,  // 0: teleport.samlidp.v1.ProcessSAMLIdPRequestRequest.request_time:type_name -> google.protobuf.Timestamp
	7,  // 1: teleport.samlidp.v1.ProcessSAMLIdPRequestRequest.mfa_response:type_name -> proto.MFAAuthenticateResponse
	8,  // 2: teleport.samlidp.v1.TestSAMLIdPAttributeMappingRequest.service_provider:type_name -> types.SAMLIdPServiceProviderV1
	9,  // 3: teleport.samlidp.v1.TestSAMLIdPAttributeMappingRequest.users:type_name -> types.UserV2
	4,  // 4: teleport.samlidp.v1.TestSAMLIdPAttributeMappingResponse.mapped_attributes:type_name -> teleport.samlidp.v1.MappedAttribute
	5,  // 5: teleport.samlidp.v1.MappedAttribute.mapped_values:type_name -> teleport.samlidp.v1.MappedAttribute.MappedValuesEntry
	10, // 6: teleport.samlidp.v1.MappedAttribute.MappedValuesEntry.value:type_name -> wrappers.StringValues
	0,  // 7: teleport.samlidp.v1.SAMLIdPService.ProcessSAMLIdPRequest:input_type -> teleport.samlidp.v1.ProcessSAMLIdPRequestRequest
	2,  // 8: teleport.samlidp.v1.SAMLIdPService.TestSAMLIdPAttributeMapping:input_type -> teleport.samlidp.v1.TestSAMLIdPAttributeMappingRequest
	1,  // 9: teleport.samlidp.v1.SAMLIdPService.ProcessSAMLIdPRequest:output_type -> teleport.samlidp.v1.ProcessSAMLIdPRequestResponse
	3,  // 10: teleport.samlidp.v1.SAMLIdPService.TestSAMLIdPAttributeMapping:output_type -> teleport.samlidp.v1.TestSAMLIdPAttributeMappingResponse
	9,  // [9:11] is the sub-list for method output_type
	7,  // [7:9] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_teleport_samlidp_v1_samlidp_proto_init() }
func file_teleport_samlidp_v1_samlidp_proto_init() {
	if File_teleport_samlidp_v1_samlidp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_teleport_samlidp_v1_samlidp_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ProcessSAMLIdPRequestRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_teleport_samlidp_v1_samlidp_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ProcessSAMLIdPRequestResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_teleport_samlidp_v1_samlidp_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*TestSAMLIdPAttributeMappingRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_teleport_samlidp_v1_samlidp_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*TestSAMLIdPAttributeMappingResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_teleport_samlidp_v1_samlidp_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*MappedAttribute); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_teleport_samlidp_v1_samlidp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_teleport_samlidp_v1_samlidp_proto_goTypes,
		DependencyIndexes: file_teleport_samlidp_v1_samlidp_proto_depIdxs,
		MessageInfos:      file_teleport_samlidp_v1_samlidp_proto_msgTypes,
	}.Build()
	File_teleport_samlidp_v1_samlidp_proto = out.File
	file_teleport_samlidp_v1_samlidp_proto_rawDesc = nil
	file_teleport_samlidp_v1_samlidp_proto_goTypes = nil
	file_teleport_samlidp_v1_samlidp_proto_depIdxs = nil
}
