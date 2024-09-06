// Copyright 2023 Gravitational, Inc
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
// source: teleport/machineid/v1/workload_identity_service.proto

package machineidv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The request for an individual x509 SVID.
type SVIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A PKIX, ASN.1 DER encoded public key that should be included in the x509
	// SVID.
	// Required.
	PublicKey []byte `protobuf:"bytes,1,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	// The path that should be included in the SPIFFE ID.
	// This should have a preceding slash and should not have a trailing slash.
	// Required.
	SpiffeIdPath string `protobuf:"bytes,2,opt,name=spiffe_id_path,json=spiffeIdPath,proto3" json:"spiffe_id_path,omitempty"`
	// The DNS SANs that should be included in the x509 SVID.
	// Optional.
	DnsSans []string `protobuf:"bytes,3,rep,name=dns_sans,json=dnsSans,proto3" json:"dns_sans,omitempty"`
	// The IP SANs that should be included in the x509 SVID.
	// Optional.
	IpSans []string `protobuf:"bytes,4,rep,name=ip_sans,json=ipSans,proto3" json:"ip_sans,omitempty"`
	// A hint that provides a way of distinguishing between SVIDs. These are
	// user configured and are sent back to the actual workload.
	// Optional.
	Hint string `protobuf:"bytes,5,opt,name=hint,proto3" json:"hint,omitempty"`
	// The TTL to use for the x509 SVID. A maximum value is enforced on this
	// field. Callers should inspect the returned cert to determine if their
	// requested TTL has been met, and if not, adjust their behaviour. If not
	// supplied, the default TTL will be the maximum value.
	Ttl *durationpb.Duration `protobuf:"bytes,6,opt,name=ttl,proto3" json:"ttl,omitempty"`
}

func (x *SVIDRequest) Reset() {
	*x = SVIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_machineid_v1_workload_identity_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SVIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SVIDRequest) ProtoMessage() {}

func (x *SVIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_machineid_v1_workload_identity_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SVIDRequest.ProtoReflect.Descriptor instead.
func (*SVIDRequest) Descriptor() ([]byte, []int) {
	return file_teleport_machineid_v1_workload_identity_service_proto_rawDescGZIP(), []int{0}
}

func (x *SVIDRequest) GetPublicKey() []byte {
	if x != nil {
		return x.PublicKey
	}
	return nil
}

func (x *SVIDRequest) GetSpiffeIdPath() string {
	if x != nil {
		return x.SpiffeIdPath
	}
	return ""
}

func (x *SVIDRequest) GetDnsSans() []string {
	if x != nil {
		return x.DnsSans
	}
	return nil
}

func (x *SVIDRequest) GetIpSans() []string {
	if x != nil {
		return x.IpSans
	}
	return nil
}

func (x *SVIDRequest) GetHint() string {
	if x != nil {
		return x.Hint
	}
	return ""
}

func (x *SVIDRequest) GetTtl() *durationpb.Duration {
	if x != nil {
		return x.Ttl
	}
	return nil
}

// The generated x509 SVID.
type SVIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A ASN.1 DER encoded x509 SVID.
	Certificate []byte `protobuf:"bytes,1,opt,name=certificate,proto3" json:"certificate,omitempty"`
	// The full SPIFFE ID that was included in the x509 SVID.
	SpiffeId string `protobuf:"bytes,2,opt,name=spiffe_id,json=spiffeId,proto3" json:"spiffe_id,omitempty"`
	// The hint that was included in SVIDRequest in order to allow a workload to
	// distinguish an individual SVID.
	Hint string `protobuf:"bytes,3,opt,name=hint,proto3" json:"hint,omitempty"`
}

func (x *SVIDResponse) Reset() {
	*x = SVIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_machineid_v1_workload_identity_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SVIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SVIDResponse) ProtoMessage() {}

func (x *SVIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_machineid_v1_workload_identity_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SVIDResponse.ProtoReflect.Descriptor instead.
func (*SVIDResponse) Descriptor() ([]byte, []int) {
	return file_teleport_machineid_v1_workload_identity_service_proto_rawDescGZIP(), []int{1}
}

func (x *SVIDResponse) GetCertificate() []byte {
	if x != nil {
		return x.Certificate
	}
	return nil
}

func (x *SVIDResponse) GetSpiffeId() string {
	if x != nil {
		return x.SpiffeId
	}
	return ""
}

func (x *SVIDResponse) GetHint() string {
	if x != nil {
		return x.Hint
	}
	return ""
}

// The request for SignX509SVIDs.
type SignX509SVIDsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The SVIDs that should be generated. This is repeated to allow a bot to
	// request multiple SVIDs at once and reduce the number of round trips.
	// Must be non-zero length.
	Svids []*SVIDRequest `protobuf:"bytes,1,rep,name=svids,proto3" json:"svids,omitempty"`
}

func (x *SignX509SVIDsRequest) Reset() {
	*x = SignX509SVIDsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_machineid_v1_workload_identity_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignX509SVIDsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignX509SVIDsRequest) ProtoMessage() {}

func (x *SignX509SVIDsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_machineid_v1_workload_identity_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignX509SVIDsRequest.ProtoReflect.Descriptor instead.
func (*SignX509SVIDsRequest) Descriptor() ([]byte, []int) {
	return file_teleport_machineid_v1_workload_identity_service_proto_rawDescGZIP(), []int{2}
}

func (x *SignX509SVIDsRequest) GetSvids() []*SVIDRequest {
	if x != nil {
		return x.Svids
	}
	return nil
}

// The response for SignX509SVIDs.
type SignX509SVIDsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The generated SVIDs.
	Svids []*SVIDResponse `protobuf:"bytes,1,rep,name=svids,proto3" json:"svids,omitempty"`
}

func (x *SignX509SVIDsResponse) Reset() {
	*x = SignX509SVIDsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_machineid_v1_workload_identity_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignX509SVIDsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignX509SVIDsResponse) ProtoMessage() {}

func (x *SignX509SVIDsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_machineid_v1_workload_identity_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignX509SVIDsResponse.ProtoReflect.Descriptor instead.
func (*SignX509SVIDsResponse) Descriptor() ([]byte, []int) {
	return file_teleport_machineid_v1_workload_identity_service_proto_rawDescGZIP(), []int{3}
}

func (x *SignX509SVIDsResponse) GetSvids() []*SVIDResponse {
	if x != nil {
		return x.Svids
	}
	return nil
}

// The request for an individual JWT SVID.
type JWTSVIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The path that should be included in the SPIFFE ID.
	// This should have a preceding slash and should not have a trailing slash.
	// Required.
	SpiffeIdPath string `protobuf:"bytes,1,opt,name=spiffe_id_path,json=spiffeIdPath,proto3" json:"spiffe_id_path,omitempty"`
	// The value that should be included in the JWT SVID as the `aud` claim.
	// Required.
	Audiences []string `protobuf:"bytes,2,rep,name=audiences,proto3" json:"audiences,omitempty"`
	// The TTL to use for the x509 SVID. A maximum value is enforced on this
	// field. Callers should inspect the returned cert to determine if their
	// requested TTL has been met, and if not, adjust their behaviour. If not
	// supplied, the default TTL will be the maximum value.
	Ttl *durationpb.Duration `protobuf:"bytes,3,opt,name=ttl,proto3" json:"ttl,omitempty"`
	// A hint that provides a way of distinguishing between SVIDs. These are
	// user configured and are sent back to the actual workload.
	// Optional.
	Hint string `protobuf:"bytes,4,opt,name=hint,proto3" json:"hint,omitempty"`
}

func (x *JWTSVIDRequest) Reset() {
	*x = JWTSVIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_machineid_v1_workload_identity_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JWTSVIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JWTSVIDRequest) ProtoMessage() {}

func (x *JWTSVIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_machineid_v1_workload_identity_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JWTSVIDRequest.ProtoReflect.Descriptor instead.
func (*JWTSVIDRequest) Descriptor() ([]byte, []int) {
	return file_teleport_machineid_v1_workload_identity_service_proto_rawDescGZIP(), []int{4}
}

func (x *JWTSVIDRequest) GetSpiffeIdPath() string {
	if x != nil {
		return x.SpiffeIdPath
	}
	return ""
}

func (x *JWTSVIDRequest) GetAudiences() []string {
	if x != nil {
		return x.Audiences
	}
	return nil
}

func (x *JWTSVIDRequest) GetTtl() *durationpb.Duration {
	if x != nil {
		return x.Ttl
	}
	return nil
}

func (x *JWTSVIDRequest) GetHint() string {
	if x != nil {
		return x.Hint
	}
	return ""
}

// The generated JWT SVID.
type JWTSVIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The JWT SVID.
	Jwt string `protobuf:"bytes,1,opt,name=jwt,proto3" json:"jwt,omitempty"`
	// The JTI that was included in the JWT.
	Jti string `protobuf:"bytes,2,opt,name=jti,proto3" json:"jti,omitempty"`
	// The full SPIFFE ID that was included in the x509 SVID.
	SpiffeId string `protobuf:"bytes,3,opt,name=spiffe_id,json=spiffeId,proto3" json:"spiffe_id,omitempty"`
	// The audience that was included in the JWT.
	Audience []string `protobuf:"bytes,4,rep,name=audience,proto3" json:"audience,omitempty"`
	// The hint that was included in SVIDRequest in order to allow a workload to
	// distinguish an individual SVID.
	Hint string `protobuf:"bytes,5,opt,name=hint,proto3" json:"hint,omitempty"`
}

func (x *JWTSVIDResponse) Reset() {
	*x = JWTSVIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_machineid_v1_workload_identity_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JWTSVIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JWTSVIDResponse) ProtoMessage() {}

func (x *JWTSVIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_machineid_v1_workload_identity_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JWTSVIDResponse.ProtoReflect.Descriptor instead.
func (*JWTSVIDResponse) Descriptor() ([]byte, []int) {
	return file_teleport_machineid_v1_workload_identity_service_proto_rawDescGZIP(), []int{5}
}

func (x *JWTSVIDResponse) GetJwt() string {
	if x != nil {
		return x.Jwt
	}
	return ""
}

func (x *JWTSVIDResponse) GetJti() string {
	if x != nil {
		return x.Jti
	}
	return ""
}

func (x *JWTSVIDResponse) GetSpiffeId() string {
	if x != nil {
		return x.SpiffeId
	}
	return ""
}

func (x *JWTSVIDResponse) GetAudience() []string {
	if x != nil {
		return x.Audience
	}
	return nil
}

func (x *JWTSVIDResponse) GetHint() string {
	if x != nil {
		return x.Hint
	}
	return ""
}

// The request for SignJWTSVIDs.
type SignJWTSVIDsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Svids []*JWTSVIDRequest `protobuf:"bytes,1,rep,name=svids,proto3" json:"svids,omitempty"`
}

func (x *SignJWTSVIDsRequest) Reset() {
	*x = SignJWTSVIDsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_machineid_v1_workload_identity_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignJWTSVIDsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignJWTSVIDsRequest) ProtoMessage() {}

func (x *SignJWTSVIDsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_machineid_v1_workload_identity_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignJWTSVIDsRequest.ProtoReflect.Descriptor instead.
func (*SignJWTSVIDsRequest) Descriptor() ([]byte, []int) {
	return file_teleport_machineid_v1_workload_identity_service_proto_rawDescGZIP(), []int{6}
}

func (x *SignJWTSVIDsRequest) GetSvids() []*JWTSVIDRequest {
	if x != nil {
		return x.Svids
	}
	return nil
}

// The response for SignJWTSVIDs.
type SignJWTSVIDsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Svids []*JWTSVIDResponse `protobuf:"bytes,1,rep,name=svids,proto3" json:"svids,omitempty"`
}

func (x *SignJWTSVIDsResponse) Reset() {
	*x = SignJWTSVIDsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_machineid_v1_workload_identity_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignJWTSVIDsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignJWTSVIDsResponse) ProtoMessage() {}

func (x *SignJWTSVIDsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_machineid_v1_workload_identity_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignJWTSVIDsResponse.ProtoReflect.Descriptor instead.
func (*SignJWTSVIDsResponse) Descriptor() ([]byte, []int) {
	return file_teleport_machineid_v1_workload_identity_service_proto_rawDescGZIP(), []int{7}
}

func (x *SignJWTSVIDsResponse) GetSvids() []*JWTSVIDResponse {
	if x != nil {
		return x.Svids
	}
	return nil
}

var File_teleport_machineid_v1_workload_identity_service_proto protoreflect.FileDescriptor

var file_teleport_machineid_v1_workload_identity_service_proto_rawDesc = []byte{
	0x0a, 0x35, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x6d, 0x61, 0x63, 0x68, 0x69,
	0x6e, 0x65, 0x69, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64,
	0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x2e, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x69, 0x64, 0x2e, 0x76, 0x31, 0x1a, 0x1e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc7,
	0x01, 0x0a, 0x0b, 0x53, 0x56, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x24, 0x0a,
	0x0e, 0x73, 0x70, 0x69, 0x66, 0x66, 0x65, 0x5f, 0x69, 0x64, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x70, 0x69, 0x66, 0x66, 0x65, 0x49, 0x64, 0x50,
	0x61, 0x74, 0x68, 0x12, 0x19, 0x0a, 0x08, 0x64, 0x6e, 0x73, 0x5f, 0x73, 0x61, 0x6e, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x64, 0x6e, 0x73, 0x53, 0x61, 0x6e, 0x73, 0x12, 0x17,
	0x0a, 0x07, 0x69, 0x70, 0x5f, 0x73, 0x61, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x06, 0x69, 0x70, 0x53, 0x61, 0x6e, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x69, 0x6e, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x69, 0x6e, 0x74, 0x12, 0x2b, 0x0a, 0x03, 0x74,
	0x74, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x03, 0x74, 0x74, 0x6c, 0x22, 0x61, 0x0a, 0x0c, 0x53, 0x56, 0x49, 0x44,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x65, 0x72, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x63,
	0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x70,
	0x69, 0x66, 0x66, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73,
	0x70, 0x69, 0x66, 0x66, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x69, 0x6e, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x69, 0x6e, 0x74, 0x22, 0x50, 0x0a, 0x14, 0x53,
	0x69, 0x67, 0x6e, 0x58, 0x35, 0x30, 0x39, 0x53, 0x56, 0x49, 0x44, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x05, 0x73, 0x76, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x22, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x6d, 0x61,
	0x63, 0x68, 0x69, 0x6e, 0x65, 0x69, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x56, 0x49, 0x44, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x05, 0x73, 0x76, 0x69, 0x64, 0x73, 0x22, 0x52, 0x0a,
	0x15, 0x53, 0x69, 0x67, 0x6e, 0x58, 0x35, 0x30, 0x39, 0x53, 0x56, 0x49, 0x44, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x05, 0x73, 0x76, 0x69, 0x64, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x2e, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x69, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x56,
	0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x05, 0x73, 0x76, 0x69, 0x64,
	0x73, 0x22, 0x95, 0x01, 0x0a, 0x0e, 0x4a, 0x57, 0x54, 0x53, 0x56, 0x49, 0x44, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0e, 0x73, 0x70, 0x69, 0x66, 0x66, 0x65, 0x5f, 0x69,
	0x64, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x70,
	0x69, 0x66, 0x66, 0x65, 0x49, 0x64, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x75,
	0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x61,
	0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x2b, 0x0a, 0x03, 0x74, 0x74, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x03, 0x74, 0x74, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x69, 0x6e, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x69, 0x6e, 0x74, 0x22, 0x82, 0x01, 0x0a, 0x0f, 0x4a, 0x57,
	0x54, 0x53, 0x56, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x6a, 0x77, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6a, 0x77, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x6a, 0x74, 0x69, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6a, 0x74,
	0x69, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x70, 0x69, 0x66, 0x66, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x70, 0x69, 0x66, 0x66, 0x65, 0x49, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x61, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x08, 0x61, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x69,
	0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x69, 0x6e, 0x74, 0x22, 0x52,
	0x0a, 0x13, 0x53, 0x69, 0x67, 0x6e, 0x4a, 0x57, 0x54, 0x53, 0x56, 0x49, 0x44, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3b, 0x0a, 0x05, 0x73, 0x76, 0x69, 0x64, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e,
	0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x69, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x57, 0x54,
	0x53, 0x56, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x05, 0x73, 0x76, 0x69,
	0x64, 0x73, 0x22, 0x54, 0x0a, 0x14, 0x53, 0x69, 0x67, 0x6e, 0x4a, 0x57, 0x54, 0x53, 0x56, 0x49,
	0x44, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x05, 0x73, 0x76,
	0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x74, 0x65, 0x6c, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x2e, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x69, 0x64, 0x2e, 0x76,
	0x31, 0x2e, 0x4a, 0x57, 0x54, 0x53, 0x56, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x52, 0x05, 0x73, 0x76, 0x69, 0x64, 0x73, 0x32, 0xf2, 0x01, 0x0a, 0x17, 0x57, 0x6f, 0x72,
	0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x6c, 0x0a, 0x0d, 0x53, 0x69, 0x67, 0x6e, 0x58, 0x35, 0x30, 0x39,
	0x53, 0x56, 0x49, 0x44, 0x73, 0x12, 0x2b, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x2e, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x69, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69,
	0x67, 0x6e, 0x58, 0x35, 0x30, 0x39, 0x53, 0x56, 0x49, 0x44, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x6d, 0x61,
	0x63, 0x68, 0x69, 0x6e, 0x65, 0x69, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x58,
	0x35, 0x30, 0x39, 0x53, 0x56, 0x49, 0x44, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x69, 0x0a, 0x0c, 0x53, 0x69, 0x67, 0x6e, 0x4a, 0x57, 0x54, 0x53, 0x56, 0x49,
	0x44, 0x73, 0x12, 0x2a, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x6d, 0x61,
	0x63, 0x68, 0x69, 0x6e, 0x65, 0x69, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x4a,
	0x57, 0x54, 0x53, 0x56, 0x49, 0x44, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b,
	0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e,
	0x65, 0x69, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x4a, 0x57, 0x54, 0x53, 0x56,
	0x49, 0x44, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x56, 0x5a,
	0x54, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x61, 0x76,
	0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x6d, 0x61, 0x63,
	0x68, 0x69, 0x6e, 0x65, 0x69, 0x64, 0x2f, 0x76, 0x31, 0x3b, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e,
	0x65, 0x69, 0x64, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_teleport_machineid_v1_workload_identity_service_proto_rawDescOnce sync.Once
	file_teleport_machineid_v1_workload_identity_service_proto_rawDescData = file_teleport_machineid_v1_workload_identity_service_proto_rawDesc
)

func file_teleport_machineid_v1_workload_identity_service_proto_rawDescGZIP() []byte {
	file_teleport_machineid_v1_workload_identity_service_proto_rawDescOnce.Do(func() {
		file_teleport_machineid_v1_workload_identity_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_teleport_machineid_v1_workload_identity_service_proto_rawDescData)
	})
	return file_teleport_machineid_v1_workload_identity_service_proto_rawDescData
}

var file_teleport_machineid_v1_workload_identity_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_teleport_machineid_v1_workload_identity_service_proto_goTypes = []any{
	(*SVIDRequest)(nil),           // 0: teleport.machineid.v1.SVIDRequest
	(*SVIDResponse)(nil),          // 1: teleport.machineid.v1.SVIDResponse
	(*SignX509SVIDsRequest)(nil),  // 2: teleport.machineid.v1.SignX509SVIDsRequest
	(*SignX509SVIDsResponse)(nil), // 3: teleport.machineid.v1.SignX509SVIDsResponse
	(*JWTSVIDRequest)(nil),        // 4: teleport.machineid.v1.JWTSVIDRequest
	(*JWTSVIDResponse)(nil),       // 5: teleport.machineid.v1.JWTSVIDResponse
	(*SignJWTSVIDsRequest)(nil),   // 6: teleport.machineid.v1.SignJWTSVIDsRequest
	(*SignJWTSVIDsResponse)(nil),  // 7: teleport.machineid.v1.SignJWTSVIDsResponse
	(*durationpb.Duration)(nil),   // 8: google.protobuf.Duration
}
var file_teleport_machineid_v1_workload_identity_service_proto_depIdxs = []int32{
	8, // 0: teleport.machineid.v1.SVIDRequest.ttl:type_name -> google.protobuf.Duration
	0, // 1: teleport.machineid.v1.SignX509SVIDsRequest.svids:type_name -> teleport.machineid.v1.SVIDRequest
	1, // 2: teleport.machineid.v1.SignX509SVIDsResponse.svids:type_name -> teleport.machineid.v1.SVIDResponse
	8, // 3: teleport.machineid.v1.JWTSVIDRequest.ttl:type_name -> google.protobuf.Duration
	4, // 4: teleport.machineid.v1.SignJWTSVIDsRequest.svids:type_name -> teleport.machineid.v1.JWTSVIDRequest
	5, // 5: teleport.machineid.v1.SignJWTSVIDsResponse.svids:type_name -> teleport.machineid.v1.JWTSVIDResponse
	2, // 6: teleport.machineid.v1.WorkloadIdentityService.SignX509SVIDs:input_type -> teleport.machineid.v1.SignX509SVIDsRequest
	6, // 7: teleport.machineid.v1.WorkloadIdentityService.SignJWTSVIDs:input_type -> teleport.machineid.v1.SignJWTSVIDsRequest
	3, // 8: teleport.machineid.v1.WorkloadIdentityService.SignX509SVIDs:output_type -> teleport.machineid.v1.SignX509SVIDsResponse
	7, // 9: teleport.machineid.v1.WorkloadIdentityService.SignJWTSVIDs:output_type -> teleport.machineid.v1.SignJWTSVIDsResponse
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_teleport_machineid_v1_workload_identity_service_proto_init() }
func file_teleport_machineid_v1_workload_identity_service_proto_init() {
	if File_teleport_machineid_v1_workload_identity_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_teleport_machineid_v1_workload_identity_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*SVIDRequest); i {
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
		file_teleport_machineid_v1_workload_identity_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*SVIDResponse); i {
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
		file_teleport_machineid_v1_workload_identity_service_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*SignX509SVIDsRequest); i {
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
		file_teleport_machineid_v1_workload_identity_service_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*SignX509SVIDsResponse); i {
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
		file_teleport_machineid_v1_workload_identity_service_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*JWTSVIDRequest); i {
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
		file_teleport_machineid_v1_workload_identity_service_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*JWTSVIDResponse); i {
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
		file_teleport_machineid_v1_workload_identity_service_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*SignJWTSVIDsRequest); i {
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
		file_teleport_machineid_v1_workload_identity_service_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*SignJWTSVIDsResponse); i {
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
			RawDescriptor: file_teleport_machineid_v1_workload_identity_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_teleport_machineid_v1_workload_identity_service_proto_goTypes,
		DependencyIndexes: file_teleport_machineid_v1_workload_identity_service_proto_depIdxs,
		MessageInfos:      file_teleport_machineid_v1_workload_identity_service_proto_msgTypes,
	}.Build()
	File_teleport_machineid_v1_workload_identity_service_proto = out.File
	file_teleport_machineid_v1_workload_identity_service_proto_rawDesc = nil
	file_teleport_machineid_v1_workload_identity_service_proto_goTypes = nil
	file_teleport_machineid_v1_workload_identity_service_proto_depIdxs = nil
}
