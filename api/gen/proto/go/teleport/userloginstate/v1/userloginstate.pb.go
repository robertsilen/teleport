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
// 	protoc-gen-go v1.34.0
// 	protoc        (unknown)
// source: teleport/userloginstate/v1/userloginstate.proto

package userloginstatev1

import (
	v1 "github.com/gravitational/teleport/api/gen/proto/go/teleport/header/v1"
	v11 "github.com/gravitational/teleport/api/gen/proto/go/teleport/trait/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// UserLoginState describes the ephemeral user login state for a user.
type UserLoginState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// header is the header for the resource.
	Header *v1.ResourceHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// spec is the specification for the user login state.
	Spec *Spec `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
}

func (x *UserLoginState) Reset() {
	*x = UserLoginState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_userloginstate_v1_userloginstate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserLoginState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserLoginState) ProtoMessage() {}

func (x *UserLoginState) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_userloginstate_v1_userloginstate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserLoginState.ProtoReflect.Descriptor instead.
func (*UserLoginState) Descriptor() ([]byte, []int) {
	return file_teleport_userloginstate_v1_userloginstate_proto_rawDescGZIP(), []int{0}
}

func (x *UserLoginState) GetHeader() *v1.ResourceHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *UserLoginState) GetSpec() *Spec {
	if x != nil {
		return x.Spec
	}
	return nil
}

// Spec is the specification for a user login state.
type Spec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// roles are the user roles attached to the user.
	Roles []string `protobuf:"bytes,1,rep,name=roles,proto3" json:"roles,omitempty"`
	// traits are the traits attached to the user.
	Traits []*v11.Trait `protobuf:"bytes,2,rep,name=traits,proto3" json:"traits,omitempty"`
	// user_type is the type of user this state represents.
	UserType string `protobuf:"bytes,3,opt,name=user_type,json=userType,proto3" json:"user_type,omitempty"`
	// original_roles are the user roles that are part of the user's static definition. These roles are
	// not affected by access granted by access lists and are obtained prior to granting access list access.
	OriginalRoles []string `protobuf:"bytes,4,rep,name=original_roles,json=originalRoles,proto3" json:"original_roles,omitempty"`
	// original_traits are the user traits that are part of the user's static definition. These traits are
	// not affected by access granted by access lists and are obtained prior to granting access list access.
	OriginalTraits []*v11.Trait `protobuf:"bytes,5,rep,name=original_traits,json=originalTraits,proto3" json:"original_traits,omitempty"`
}

func (x *Spec) Reset() {
	*x = Spec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_userloginstate_v1_userloginstate_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Spec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Spec) ProtoMessage() {}

func (x *Spec) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_userloginstate_v1_userloginstate_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Spec.ProtoReflect.Descriptor instead.
func (*Spec) Descriptor() ([]byte, []int) {
	return file_teleport_userloginstate_v1_userloginstate_proto_rawDescGZIP(), []int{1}
}

func (x *Spec) GetRoles() []string {
	if x != nil {
		return x.Roles
	}
	return nil
}

func (x *Spec) GetTraits() []*v11.Trait {
	if x != nil {
		return x.Traits
	}
	return nil
}

func (x *Spec) GetUserType() string {
	if x != nil {
		return x.UserType
	}
	return ""
}

func (x *Spec) GetOriginalRoles() []string {
	if x != nil {
		return x.OriginalRoles
	}
	return nil
}

func (x *Spec) GetOriginalTraits() []*v11.Trait {
	if x != nil {
		return x.OriginalTraits
	}
	return nil
}

var File_teleport_userloginstate_v1_userloginstate_proto protoreflect.FileDescriptor

var file_teleport_userloginstate_v1_userloginstate_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x6c,
	0x6f, 0x67, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1a, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x27, 0x74,
	0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2f, 0x76,
	0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x2f, 0x74, 0x72, 0x61, 0x69, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x72, 0x61, 0x69, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x82, 0x01, 0x0a, 0x0e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x3a, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x12, 0x34, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x20, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x70, 0x65, 0x63, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x22, 0xd5, 0x01, 0x0a, 0x04, 0x53,
	0x70, 0x65, 0x63, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0x30, 0x0a, 0x06, 0x74, 0x72, 0x61,
	0x69, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74, 0x65, 0x6c, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72,
	0x61, 0x69, 0x74, 0x52, 0x06, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x6f, 0x72, 0x69, 0x67,
	0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0d, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x12,
	0x41, 0x0a, 0x0f, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x74, 0x72, 0x61, 0x69,
	0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61,
	0x69, 0x74, 0x52, 0x0e, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x54, 0x72, 0x61, 0x69,
	0x74, 0x73, 0x42, 0x60, 0x5a, 0x5e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x67, 0x72, 0x61, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x2f, 0x74,
	0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x2f, 0x76, 0x31, 0x3b, 0x75, 0x73, 0x65, 0x72, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_teleport_userloginstate_v1_userloginstate_proto_rawDescOnce sync.Once
	file_teleport_userloginstate_v1_userloginstate_proto_rawDescData = file_teleport_userloginstate_v1_userloginstate_proto_rawDesc
)

func file_teleport_userloginstate_v1_userloginstate_proto_rawDescGZIP() []byte {
	file_teleport_userloginstate_v1_userloginstate_proto_rawDescOnce.Do(func() {
		file_teleport_userloginstate_v1_userloginstate_proto_rawDescData = protoimpl.X.CompressGZIP(file_teleport_userloginstate_v1_userloginstate_proto_rawDescData)
	})
	return file_teleport_userloginstate_v1_userloginstate_proto_rawDescData
}

var file_teleport_userloginstate_v1_userloginstate_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_teleport_userloginstate_v1_userloginstate_proto_goTypes = []interface{}{
	(*UserLoginState)(nil),    // 0: teleport.userloginstate.v1.UserLoginState
	(*Spec)(nil),              // 1: teleport.userloginstate.v1.Spec
	(*v1.ResourceHeader)(nil), // 2: teleport.header.v1.ResourceHeader
	(*v11.Trait)(nil),         // 3: teleport.trait.v1.Trait
}
var file_teleport_userloginstate_v1_userloginstate_proto_depIdxs = []int32{
	2, // 0: teleport.userloginstate.v1.UserLoginState.header:type_name -> teleport.header.v1.ResourceHeader
	1, // 1: teleport.userloginstate.v1.UserLoginState.spec:type_name -> teleport.userloginstate.v1.Spec
	3, // 2: teleport.userloginstate.v1.Spec.traits:type_name -> teleport.trait.v1.Trait
	3, // 3: teleport.userloginstate.v1.Spec.original_traits:type_name -> teleport.trait.v1.Trait
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_teleport_userloginstate_v1_userloginstate_proto_init() }
func file_teleport_userloginstate_v1_userloginstate_proto_init() {
	if File_teleport_userloginstate_v1_userloginstate_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_teleport_userloginstate_v1_userloginstate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserLoginState); i {
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
		file_teleport_userloginstate_v1_userloginstate_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Spec); i {
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
			RawDescriptor: file_teleport_userloginstate_v1_userloginstate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_teleport_userloginstate_v1_userloginstate_proto_goTypes,
		DependencyIndexes: file_teleport_userloginstate_v1_userloginstate_proto_depIdxs,
		MessageInfos:      file_teleport_userloginstate_v1_userloginstate_proto_msgTypes,
	}.Build()
	File_teleport_userloginstate_v1_userloginstate_proto = out.File
	file_teleport_userloginstate_v1_userloginstate_proto_rawDesc = nil
	file_teleport_userloginstate_v1_userloginstate_proto_goTypes = nil
	file_teleport_userloginstate_v1_userloginstate_proto_depIdxs = nil
}
