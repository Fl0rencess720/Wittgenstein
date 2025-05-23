// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: gateway/role/v1/role.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type Role struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid         string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Avatar      string `protobuf:"bytes,3,opt,name=avatar,proto3" json:"avatar,omitempty"`
	ApiPath     string `protobuf:"bytes,4,opt,name=api_path,json=apiPath,proto3" json:"api_path,omitempty"`
	ApiKey      string `protobuf:"bytes,5,opt,name=api_key,json=apiKey,proto3" json:"api_key,omitempty"`
	Model       *Model `protobuf:"bytes,6,opt,name=model,proto3" json:"model,omitempty"`
	Name        string `protobuf:"bytes,7,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Role) Reset() {
	*x = Role{}
	mi := &file_gateway_role_v1_role_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Role) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Role) ProtoMessage() {}

func (x *Role) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_role_v1_role_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Role.ProtoReflect.Descriptor instead.
func (*Role) Descriptor() ([]byte, []int) {
	return file_gateway_role_v1_role_proto_rawDescGZIP(), []int{0}
}

func (x *Role) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *Role) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Role) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *Role) GetApiPath() string {
	if x != nil {
		return x.ApiPath
	}
	return ""
}

func (x *Role) GetApiKey() string {
	if x != nil {
		return x.ApiKey
	}
	return ""
}

func (x *Role) GetModel() *Model {
	if x != nil {
		return x.Model
	}
	return nil
}

func (x *Role) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Model struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Provider string `protobuf:"bytes,1,opt,name=provider,proto3" json:"provider,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Model) Reset() {
	*x = Model{}
	mi := &file_gateway_role_v1_role_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Model) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Model) ProtoMessage() {}

func (x *Model) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_role_v1_role_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Model.ProtoReflect.Descriptor instead.
func (*Model) Descriptor() ([]byte, []int) {
	return file_gateway_role_v1_role_proto_rawDescGZIP(), []int{1}
}

func (x *Model) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (x *Model) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateRoleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Phone string `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`
	Role  *Role  `protobuf:"bytes,2,opt,name=role,proto3" json:"role,omitempty"`
}

func (x *CreateRoleRequest) Reset() {
	*x = CreateRoleRequest{}
	mi := &file_gateway_role_v1_role_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateRoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRoleRequest) ProtoMessage() {}

func (x *CreateRoleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_role_v1_role_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRoleRequest.ProtoReflect.Descriptor instead.
func (*CreateRoleRequest) Descriptor() ([]byte, []int) {
	return file_gateway_role_v1_role_proto_rawDescGZIP(), []int{2}
}

func (x *CreateRoleRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *CreateRoleRequest) GetRole() *Role {
	if x != nil {
		return x.Role
	}
	return nil
}

type CreateRoleReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *CreateRoleReply) Reset() {
	*x = CreateRoleReply{}
	mi := &file_gateway_role_v1_role_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateRoleReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRoleReply) ProtoMessage() {}

func (x *CreateRoleReply) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_role_v1_role_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRoleReply.ProtoReflect.Descriptor instead.
func (*CreateRoleReply) Descriptor() ([]byte, []int) {
	return file_gateway_role_v1_role_proto_rawDescGZIP(), []int{3}
}

func (x *CreateRoleReply) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

type DeleteRoleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Phone string `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`
	Uid   string `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *DeleteRoleRequest) Reset() {
	*x = DeleteRoleRequest{}
	mi := &file_gateway_role_v1_role_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteRoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRoleRequest) ProtoMessage() {}

func (x *DeleteRoleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_role_v1_role_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRoleRequest.ProtoReflect.Descriptor instead.
func (*DeleteRoleRequest) Descriptor() ([]byte, []int) {
	return file_gateway_role_v1_role_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteRoleRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *DeleteRoleRequest) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

type DeleteRoleReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DeleteRoleReply) Reset() {
	*x = DeleteRoleReply{}
	mi := &file_gateway_role_v1_role_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteRoleReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRoleReply) ProtoMessage() {}

func (x *DeleteRoleReply) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_role_v1_role_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRoleReply.ProtoReflect.Descriptor instead.
func (*DeleteRoleReply) Descriptor() ([]byte, []int) {
	return file_gateway_role_v1_role_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteRoleReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GetRolesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Phone string `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`
}

func (x *GetRolesRequest) Reset() {
	*x = GetRolesRequest{}
	mi := &file_gateway_role_v1_role_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRolesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRolesRequest) ProtoMessage() {}

func (x *GetRolesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_role_v1_role_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRolesRequest.ProtoReflect.Descriptor instead.
func (*GetRolesRequest) Descriptor() ([]byte, []int) {
	return file_gateway_role_v1_role_proto_rawDescGZIP(), []int{6}
}

func (x *GetRolesRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

type GetRolesReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Roles []*Role `protobuf:"bytes,1,rep,name=roles,proto3" json:"roles,omitempty"`
}

func (x *GetRolesReply) Reset() {
	*x = GetRolesReply{}
	mi := &file_gateway_role_v1_role_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRolesReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRolesReply) ProtoMessage() {}

func (x *GetRolesReply) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_role_v1_role_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRolesReply.ProtoReflect.Descriptor instead.
func (*GetRolesReply) Descriptor() ([]byte, []int) {
	return file_gateway_role_v1_role_proto_rawDescGZIP(), []int{7}
}

func (x *GetRolesReply) GetRoles() []*Role {
	if x != nil {
		return x.Roles
	}
	return nil
}

type SetRoleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Phone string `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`
	Uid   string `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Role  *Role  `protobuf:"bytes,3,opt,name=role,proto3" json:"role,omitempty"`
}

func (x *SetRoleRequest) Reset() {
	*x = SetRoleRequest{}
	mi := &file_gateway_role_v1_role_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetRoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetRoleRequest) ProtoMessage() {}

func (x *SetRoleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_role_v1_role_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetRoleRequest.ProtoReflect.Descriptor instead.
func (*SetRoleRequest) Descriptor() ([]byte, []int) {
	return file_gateway_role_v1_role_proto_rawDescGZIP(), []int{8}
}

func (x *SetRoleRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *SetRoleRequest) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *SetRoleRequest) GetRole() *Role {
	if x != nil {
		return x.Role
	}
	return nil
}

type SetRoleReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *SetRoleReply) Reset() {
	*x = SetRoleReply{}
	mi := &file_gateway_role_v1_role_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetRoleReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetRoleReply) ProtoMessage() {}

func (x *SetRoleReply) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_role_v1_role_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetRoleReply.ProtoReflect.Descriptor instead.
func (*SetRoleReply) Descriptor() ([]byte, []int) {
	return file_gateway_role_v1_role_proto_rawDescGZIP(), []int{9}
}

func (x *SetRoleReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GetAvailableModelsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAvailableModelsRequest) Reset() {
	*x = GetAvailableModelsRequest{}
	mi := &file_gateway_role_v1_role_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAvailableModelsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAvailableModelsRequest) ProtoMessage() {}

func (x *GetAvailableModelsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_role_v1_role_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAvailableModelsRequest.ProtoReflect.Descriptor instead.
func (*GetAvailableModelsRequest) Descriptor() ([]byte, []int) {
	return file_gateway_role_v1_role_proto_rawDescGZIP(), []int{10}
}

type GetAvailableModelsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Models []*Model `protobuf:"bytes,1,rep,name=models,proto3" json:"models,omitempty"`
}

func (x *GetAvailableModelsReply) Reset() {
	*x = GetAvailableModelsReply{}
	mi := &file_gateway_role_v1_role_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAvailableModelsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAvailableModelsReply) ProtoMessage() {}

func (x *GetAvailableModelsReply) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_role_v1_role_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAvailableModelsReply.ProtoReflect.Descriptor instead.
func (*GetAvailableModelsReply) Descriptor() ([]byte, []int) {
	return file_gateway_role_v1_role_proto_rawDescGZIP(), []int{11}
}

func (x *GetAvailableModelsReply) GetModels() []*Model {
	if x != nil {
		return x.Models
	}
	return nil
}

type GetModeratorAndParticipantsByUIDsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Phone     string   `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`
	Moderator string   `protobuf:"bytes,2,opt,name=moderator,proto3" json:"moderator,omitempty"`
	Uids      []string `protobuf:"bytes,3,rep,name=uids,proto3" json:"uids,omitempty"`
}

func (x *GetModeratorAndParticipantsByUIDsRequest) Reset() {
	*x = GetModeratorAndParticipantsByUIDsRequest{}
	mi := &file_gateway_role_v1_role_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetModeratorAndParticipantsByUIDsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetModeratorAndParticipantsByUIDsRequest) ProtoMessage() {}

func (x *GetModeratorAndParticipantsByUIDsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_role_v1_role_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetModeratorAndParticipantsByUIDsRequest.ProtoReflect.Descriptor instead.
func (*GetModeratorAndParticipantsByUIDsRequest) Descriptor() ([]byte, []int) {
	return file_gateway_role_v1_role_proto_rawDescGZIP(), []int{12}
}

func (x *GetModeratorAndParticipantsByUIDsRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *GetModeratorAndParticipantsByUIDsRequest) GetModerator() string {
	if x != nil {
		return x.Moderator
	}
	return ""
}

func (x *GetModeratorAndParticipantsByUIDsRequest) GetUids() []string {
	if x != nil {
		return x.Uids
	}
	return nil
}

type GetModeratorAndParticipantsByUIDsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Moderator    *Role   `protobuf:"bytes,1,opt,name=moderator,proto3" json:"moderator,omitempty"`
	Participants []*Role `protobuf:"bytes,2,rep,name=participants,proto3" json:"participants,omitempty"`
}

func (x *GetModeratorAndParticipantsByUIDsReply) Reset() {
	*x = GetModeratorAndParticipantsByUIDsReply{}
	mi := &file_gateway_role_v1_role_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetModeratorAndParticipantsByUIDsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetModeratorAndParticipantsByUIDsReply) ProtoMessage() {}

func (x *GetModeratorAndParticipantsByUIDsReply) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_role_v1_role_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetModeratorAndParticipantsByUIDsReply.ProtoReflect.Descriptor instead.
func (*GetModeratorAndParticipantsByUIDsReply) Descriptor() ([]byte, []int) {
	return file_gateway_role_v1_role_proto_rawDescGZIP(), []int{13}
}

func (x *GetModeratorAndParticipantsByUIDsReply) GetModerator() *Role {
	if x != nil {
		return x.Moderator
	}
	return nil
}

func (x *GetModeratorAndParticipantsByUIDsReply) GetParticipants() []*Role {
	if x != nil {
		return x.Participants
	}
	return nil
}

var File_gateway_role_v1_role_proto protoreflect.FileDescriptor

var file_gateway_role_v1_role_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2f, 0x76,
	0x31, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x41, 0x79,
	0x61, 0x6e, 0x61, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc1, 0x01, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x70, 0x69,
	0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x70, 0x69,
	0x50, 0x61, 0x74, 0x68, 0x12, 0x17, 0x0a, 0x07, 0x61, 0x70, 0x69, 0x5f, 0x6b, 0x65, 0x79, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x12, 0x25, 0x0a,
	0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x41,
	0x79, 0x61, 0x6e, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x05, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x37, 0x0a, 0x05, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x4d, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x22, 0x0a, 0x04,
	0x72, 0x6f, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x41, 0x79, 0x61,
	0x6e, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65,
	0x22, 0x23, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x75, 0x69, 0x64, 0x22, 0x3b, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75,
	0x69, 0x64, 0x22, 0x2b, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x27, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x22, 0x35, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x52,
	0x6f, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x24, 0x0a, 0x05, 0x72, 0x6f, 0x6c,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x41, 0x79, 0x61, 0x6e, 0x61,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x22,
	0x5c, 0x0a, 0x0e, 0x53, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x22, 0x0a, 0x04, 0x72, 0x6f, 0x6c,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x41, 0x79, 0x61, 0x6e, 0x61, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x22, 0x28, 0x0a,
	0x0c, 0x53, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x1b, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x41, 0x76,
	0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x42, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x41, 0x76, 0x61, 0x69, 0x6c,
	0x61, 0x62, 0x6c, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x27, 0x0a, 0x06, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x41, 0x79, 0x61, 0x6e, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x52, 0x06, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x22, 0x72, 0x0a, 0x28, 0x47, 0x65, 0x74, 0x4d,
	0x6f, 0x64, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x41, 0x6e, 0x64, 0x50, 0x61, 0x72, 0x74, 0x69,
	0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x73, 0x42, 0x79, 0x55, 0x49, 0x44, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x6f,
	0x64, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d,
	0x6f, 0x64, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x69, 0x64, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x75, 0x69, 0x64, 0x73, 0x22, 0x8a, 0x01, 0x0a,
	0x26, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x41, 0x6e, 0x64,
	0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x73, 0x42, 0x79, 0x55, 0x49,
	0x44, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2c, 0x0a, 0x09, 0x6d, 0x6f, 0x64, 0x65, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x41, 0x79, 0x61,
	0x6e, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x09, 0x6d, 0x6f, 0x64, 0x65,
	0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x32, 0x0a, 0x0c, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69,
	0x70, 0x61, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x41, 0x79,
	0x61, 0x6e, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x0c, 0x70, 0x61, 0x72,
	0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x73, 0x32, 0x8c, 0x05, 0x0a, 0x0b, 0x52, 0x6f,
	0x6c, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x5f, 0x0a, 0x0a, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x1b, 0x2e, 0x41, 0x79, 0x61, 0x6e, 0x61, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x41, 0x79, 0x61, 0x6e, 0x61, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x3a, 0x01, 0x2a, 0x22, 0x0e, 0x2f, 0x72, 0x6f, 0x6c,
	0x65, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x5f, 0x0a, 0x0a, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x1b, 0x2e, 0x41, 0x79, 0x61, 0x6e, 0x61,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x41, 0x79, 0x61, 0x6e, 0x61, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x3a, 0x01, 0x2a, 0x22, 0x0e, 0x2f, 0x72, 0x6f,
	0x6c, 0x65, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x58, 0x0a, 0x08, 0x47,
	0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0x19, 0x2e, 0x41, 0x79, 0x61, 0x6e, 0x61, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x17, 0x2e, 0x41, 0x79, 0x61, 0x6e, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x18, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x12, 0x3a, 0x01, 0x2a, 0x22, 0x0d, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2f, 0x67, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x8b, 0x01, 0x0a, 0x21, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x64,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x41, 0x6e, 0x64, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69,
	0x70, 0x61, 0x6e, 0x74, 0x73, 0x42, 0x79, 0x55, 0x49, 0x44, 0x73, 0x12, 0x32, 0x2e, 0x41, 0x79,
	0x61, 0x6e, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x41, 0x6e, 0x64, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e,
	0x74, 0x73, 0x42, 0x79, 0x55, 0x49, 0x44, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x30, 0x2e, 0x41, 0x79, 0x61, 0x6e, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x6f,
	0x64, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x41, 0x6e, 0x64, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63,
	0x69, 0x70, 0x61, 0x6e, 0x74, 0x73, 0x42, 0x79, 0x55, 0x49, 0x44, 0x73, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x00, 0x12, 0x7c, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61,
	0x62, 0x6c, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x12, 0x23, 0x2e, 0x41, 0x79, 0x61, 0x6e,
	0x61, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c,
	0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21,
	0x2e, 0x41, 0x79, 0x61, 0x6e, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x76, 0x61,
	0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x3a, 0x01, 0x2a, 0x22, 0x13, 0x2f, 0x72,
	0x6f, 0x6c, 0x65, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x67, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x12, 0x55, 0x0a, 0x07, 0x53, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x18, 0x2e, 0x41,
	0x79, 0x61, 0x6e, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x41, 0x79, 0x61, 0x6e, 0x61, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x18,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x3a, 0x01, 0x2a, 0x22, 0x0d, 0x2f, 0x72, 0x6f, 0x6c, 0x65,
	0x2f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x42, 0x1e, 0x5a, 0x1c, 0x41, 0x79, 0x61, 0x6e,
	0x61, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x72, 0x6f,
	0x6c, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gateway_role_v1_role_proto_rawDescOnce sync.Once
	file_gateway_role_v1_role_proto_rawDescData = file_gateway_role_v1_role_proto_rawDesc
)

func file_gateway_role_v1_role_proto_rawDescGZIP() []byte {
	file_gateway_role_v1_role_proto_rawDescOnce.Do(func() {
		file_gateway_role_v1_role_proto_rawDescData = protoimpl.X.CompressGZIP(file_gateway_role_v1_role_proto_rawDescData)
	})
	return file_gateway_role_v1_role_proto_rawDescData
}

var file_gateway_role_v1_role_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_gateway_role_v1_role_proto_goTypes = []any{
	(*Role)(nil),                                     // 0: Ayana.v1.Role
	(*Model)(nil),                                    // 1: Ayana.v1.Model
	(*CreateRoleRequest)(nil),                        // 2: Ayana.v1.CreateRoleRequest
	(*CreateRoleReply)(nil),                          // 3: Ayana.v1.CreateRoleReply
	(*DeleteRoleRequest)(nil),                        // 4: Ayana.v1.DeleteRoleRequest
	(*DeleteRoleReply)(nil),                          // 5: Ayana.v1.DeleteRoleReply
	(*GetRolesRequest)(nil),                          // 6: Ayana.v1.GetRolesRequest
	(*GetRolesReply)(nil),                            // 7: Ayana.v1.GetRolesReply
	(*SetRoleRequest)(nil),                           // 8: Ayana.v1.SetRoleRequest
	(*SetRoleReply)(nil),                             // 9: Ayana.v1.SetRoleReply
	(*GetAvailableModelsRequest)(nil),                // 10: Ayana.v1.GetAvailableModelsRequest
	(*GetAvailableModelsReply)(nil),                  // 11: Ayana.v1.GetAvailableModelsReply
	(*GetModeratorAndParticipantsByUIDsRequest)(nil), // 12: Ayana.v1.GetModeratorAndParticipantsByUIDsRequest
	(*GetModeratorAndParticipantsByUIDsReply)(nil),   // 13: Ayana.v1.GetModeratorAndParticipantsByUIDsReply
}
var file_gateway_role_v1_role_proto_depIdxs = []int32{
	1,  // 0: Ayana.v1.Role.model:type_name -> Ayana.v1.Model
	0,  // 1: Ayana.v1.CreateRoleRequest.role:type_name -> Ayana.v1.Role
	0,  // 2: Ayana.v1.GetRolesReply.roles:type_name -> Ayana.v1.Role
	0,  // 3: Ayana.v1.SetRoleRequest.role:type_name -> Ayana.v1.Role
	1,  // 4: Ayana.v1.GetAvailableModelsReply.models:type_name -> Ayana.v1.Model
	0,  // 5: Ayana.v1.GetModeratorAndParticipantsByUIDsReply.moderator:type_name -> Ayana.v1.Role
	0,  // 6: Ayana.v1.GetModeratorAndParticipantsByUIDsReply.participants:type_name -> Ayana.v1.Role
	2,  // 7: Ayana.v1.RoleManager.CreateRole:input_type -> Ayana.v1.CreateRoleRequest
	4,  // 8: Ayana.v1.RoleManager.DeleteRole:input_type -> Ayana.v1.DeleteRoleRequest
	6,  // 9: Ayana.v1.RoleManager.GetRoles:input_type -> Ayana.v1.GetRolesRequest
	12, // 10: Ayana.v1.RoleManager.GetModeratorAndParticipantsByUIDs:input_type -> Ayana.v1.GetModeratorAndParticipantsByUIDsRequest
	10, // 11: Ayana.v1.RoleManager.GetAvailableModels:input_type -> Ayana.v1.GetAvailableModelsRequest
	8,  // 12: Ayana.v1.RoleManager.SetRole:input_type -> Ayana.v1.SetRoleRequest
	3,  // 13: Ayana.v1.RoleManager.CreateRole:output_type -> Ayana.v1.CreateRoleReply
	5,  // 14: Ayana.v1.RoleManager.DeleteRole:output_type -> Ayana.v1.DeleteRoleReply
	7,  // 15: Ayana.v1.RoleManager.GetRoles:output_type -> Ayana.v1.GetRolesReply
	13, // 16: Ayana.v1.RoleManager.GetModeratorAndParticipantsByUIDs:output_type -> Ayana.v1.GetModeratorAndParticipantsByUIDsReply
	11, // 17: Ayana.v1.RoleManager.GetAvailableModels:output_type -> Ayana.v1.GetAvailableModelsReply
	9,  // 18: Ayana.v1.RoleManager.SetRole:output_type -> Ayana.v1.SetRoleReply
	13, // [13:19] is the sub-list for method output_type
	7,  // [7:13] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_gateway_role_v1_role_proto_init() }
func file_gateway_role_v1_role_proto_init() {
	if File_gateway_role_v1_role_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_gateway_role_v1_role_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gateway_role_v1_role_proto_goTypes,
		DependencyIndexes: file_gateway_role_v1_role_proto_depIdxs,
		MessageInfos:      file_gateway_role_v1_role_proto_msgTypes,
	}.Build()
	File_gateway_role_v1_role_proto = out.File
	file_gateway_role_v1_role_proto_rawDesc = nil
	file_gateway_role_v1_role_proto_goTypes = nil
	file_gateway_role_v1_role_proto_depIdxs = nil
}
