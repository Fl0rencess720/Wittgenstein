// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: gateway/seminar/v1/seminar.proto

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

type TopicMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid          string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Content      string   `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	Participants []string `protobuf:"bytes,3,rep,name=participants,proto3" json:"participants,omitempty"`
}

func (x *TopicMetadata) Reset() {
	*x = TopicMetadata{}
	mi := &file_gateway_seminar_v1_seminar_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TopicMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TopicMetadata) ProtoMessage() {}

func (x *TopicMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_seminar_v1_seminar_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TopicMetadata.ProtoReflect.Descriptor instead.
func (*TopicMetadata) Descriptor() ([]byte, []int) {
	return file_gateway_seminar_v1_seminar_proto_rawDescGZIP(), []int{0}
}

func (x *TopicMetadata) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *TopicMetadata) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *TopicMetadata) GetParticipants() []string {
	if x != nil {
		return x.Participants
	}
	return nil
}

type Speech struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid     string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	RoleUid string `protobuf:"bytes,2,opt,name=role_uid,json=roleUid,proto3" json:"role_uid,omitempty"`
	Content string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *Speech) Reset() {
	*x = Speech{}
	mi := &file_gateway_seminar_v1_seminar_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Speech) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Speech) ProtoMessage() {}

func (x *Speech) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_seminar_v1_seminar_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Speech.ProtoReflect.Descriptor instead.
func (*Speech) Descriptor() ([]byte, []int) {
	return file_gateway_seminar_v1_seminar_proto_rawDescGZIP(), []int{1}
}

func (x *Speech) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *Speech) GetRoleUid() string {
	if x != nil {
		return x.RoleUid
	}
	return ""
}

func (x *Speech) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type Topic struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid          string    `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Participants []string  `protobuf:"bytes,2,rep,name=participants,proto3" json:"participants,omitempty"`
	Speeches     []*Speech `protobuf:"bytes,3,rep,name=speeches,proto3" json:"speeches,omitempty"`
	Title        string    `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	TitleImage   string    `protobuf:"bytes,5,opt,name=title_image,json=titleImage,proto3" json:"title_image,omitempty"`
	Content      string    `protobuf:"bytes,6,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *Topic) Reset() {
	*x = Topic{}
	mi := &file_gateway_seminar_v1_seminar_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Topic) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Topic) ProtoMessage() {}

func (x *Topic) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_seminar_v1_seminar_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Topic.ProtoReflect.Descriptor instead.
func (*Topic) Descriptor() ([]byte, []int) {
	return file_gateway_seminar_v1_seminar_proto_rawDescGZIP(), []int{2}
}

func (x *Topic) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *Topic) GetParticipants() []string {
	if x != nil {
		return x.Participants
	}
	return nil
}

func (x *Topic) GetSpeeches() []*Speech {
	if x != nil {
		return x.Speeches
	}
	return nil
}

func (x *Topic) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Topic) GetTitleImage() string {
	if x != nil {
		return x.TitleImage
	}
	return ""
}

func (x *Topic) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type CreateTopicRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Phone        string   `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`
	Content      string   `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	Participants []string `protobuf:"bytes,3,rep,name=participants,proto3" json:"participants,omitempty"`
}

func (x *CreateTopicRequest) Reset() {
	*x = CreateTopicRequest{}
	mi := &file_gateway_seminar_v1_seminar_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateTopicRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTopicRequest) ProtoMessage() {}

func (x *CreateTopicRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_seminar_v1_seminar_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTopicRequest.ProtoReflect.Descriptor instead.
func (*CreateTopicRequest) Descriptor() ([]byte, []int) {
	return file_gateway_seminar_v1_seminar_proto_rawDescGZIP(), []int{3}
}

func (x *CreateTopicRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *CreateTopicRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *CreateTopicRequest) GetParticipants() []string {
	if x != nil {
		return x.Participants
	}
	return nil
}

type CreateTopicReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *CreateTopicReply) Reset() {
	*x = CreateTopicReply{}
	mi := &file_gateway_seminar_v1_seminar_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateTopicReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTopicReply) ProtoMessage() {}

func (x *CreateTopicReply) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_seminar_v1_seminar_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTopicReply.ProtoReflect.Descriptor instead.
func (*CreateTopicReply) Descriptor() ([]byte, []int) {
	return file_gateway_seminar_v1_seminar_proto_rawDescGZIP(), []int{4}
}

func (x *CreateTopicReply) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

type DeleteTopicRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *DeleteTopicRequest) Reset() {
	*x = DeleteTopicRequest{}
	mi := &file_gateway_seminar_v1_seminar_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteTopicRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTopicRequest) ProtoMessage() {}

func (x *DeleteTopicRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_seminar_v1_seminar_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTopicRequest.ProtoReflect.Descriptor instead.
func (*DeleteTopicRequest) Descriptor() ([]byte, []int) {
	return file_gateway_seminar_v1_seminar_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteTopicRequest) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

type DeleteTopicReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DeleteTopicReply) Reset() {
	*x = DeleteTopicReply{}
	mi := &file_gateway_seminar_v1_seminar_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteTopicReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTopicReply) ProtoMessage() {}

func (x *DeleteTopicReply) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_seminar_v1_seminar_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTopicReply.ProtoReflect.Descriptor instead.
func (*DeleteTopicReply) Descriptor() ([]byte, []int) {
	return file_gateway_seminar_v1_seminar_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteTopicReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type StartTopicRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TopicId string `protobuf:"bytes,1,opt,name=topicId,proto3" json:"topicId,omitempty"`
}

func (x *StartTopicRequest) Reset() {
	*x = StartTopicRequest{}
	mi := &file_gateway_seminar_v1_seminar_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StartTopicRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartTopicRequest) ProtoMessage() {}

func (x *StartTopicRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_seminar_v1_seminar_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartTopicRequest.ProtoReflect.Descriptor instead.
func (*StartTopicRequest) Descriptor() ([]byte, []int) {
	return file_gateway_seminar_v1_seminar_proto_rawDescGZIP(), []int{7}
}

func (x *StartTopicRequest) GetTopicId() string {
	if x != nil {
		return x.TopicId
	}
	return ""
}

type StartTopicReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Content:
	//
	//	*StartTopicReply_Reasoning
	//	*StartTopicReply_Text
	Content isStartTopicReply_Content `protobuf_oneof:"content"`
}

func (x *StartTopicReply) Reset() {
	*x = StartTopicReply{}
	mi := &file_gateway_seminar_v1_seminar_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StartTopicReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartTopicReply) ProtoMessage() {}

func (x *StartTopicReply) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_seminar_v1_seminar_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartTopicReply.ProtoReflect.Descriptor instead.
func (*StartTopicReply) Descriptor() ([]byte, []int) {
	return file_gateway_seminar_v1_seminar_proto_rawDescGZIP(), []int{8}
}

func (m *StartTopicReply) GetContent() isStartTopicReply_Content {
	if m != nil {
		return m.Content
	}
	return nil
}

func (x *StartTopicReply) GetReasoning() string {
	if x, ok := x.GetContent().(*StartTopicReply_Reasoning); ok {
		return x.Reasoning
	}
	return ""
}

func (x *StartTopicReply) GetText() string {
	if x, ok := x.GetContent().(*StartTopicReply_Text); ok {
		return x.Text
	}
	return ""
}

type isStartTopicReply_Content interface {
	isStartTopicReply_Content()
}

type StartTopicReply_Reasoning struct {
	Reasoning string `protobuf:"bytes,1,opt,name=reasoning,proto3,oneof"`
}

type StartTopicReply_Text struct {
	Text string `protobuf:"bytes,2,opt,name=text,proto3,oneof"`
}

func (*StartTopicReply_Reasoning) isStartTopicReply_Content() {}

func (*StartTopicReply_Text) isStartTopicReply_Content() {}

type StopTopicRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TopicId string `protobuf:"bytes,1,opt,name=topicId,proto3" json:"topicId,omitempty"`
}

func (x *StopTopicRequest) Reset() {
	*x = StopTopicRequest{}
	mi := &file_gateway_seminar_v1_seminar_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StopTopicRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopTopicRequest) ProtoMessage() {}

func (x *StopTopicRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_seminar_v1_seminar_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopTopicRequest.ProtoReflect.Descriptor instead.
func (*StopTopicRequest) Descriptor() ([]byte, []int) {
	return file_gateway_seminar_v1_seminar_proto_rawDescGZIP(), []int{9}
}

func (x *StopTopicRequest) GetTopicId() string {
	if x != nil {
		return x.TopicId
	}
	return ""
}

type StopTopicReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *StopTopicReply) Reset() {
	*x = StopTopicReply{}
	mi := &file_gateway_seminar_v1_seminar_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StopTopicReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopTopicReply) ProtoMessage() {}

func (x *StopTopicReply) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_seminar_v1_seminar_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopTopicReply.ProtoReflect.Descriptor instead.
func (*StopTopicReply) Descriptor() ([]byte, []int) {
	return file_gateway_seminar_v1_seminar_proto_rawDescGZIP(), []int{10}
}

func (x *StopTopicReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GetTopicsMetadataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Phone string `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`
}

func (x *GetTopicsMetadataRequest) Reset() {
	*x = GetTopicsMetadataRequest{}
	mi := &file_gateway_seminar_v1_seminar_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTopicsMetadataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTopicsMetadataRequest) ProtoMessage() {}

func (x *GetTopicsMetadataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_seminar_v1_seminar_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTopicsMetadataRequest.ProtoReflect.Descriptor instead.
func (*GetTopicsMetadataRequest) Descriptor() ([]byte, []int) {
	return file_gateway_seminar_v1_seminar_proto_rawDescGZIP(), []int{11}
}

func (x *GetTopicsMetadataRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

type GetTopicsMetadataReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Topics []*TopicMetadata `protobuf:"bytes,1,rep,name=topics,proto3" json:"topics,omitempty"`
}

func (x *GetTopicsMetadataReply) Reset() {
	*x = GetTopicsMetadataReply{}
	mi := &file_gateway_seminar_v1_seminar_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTopicsMetadataReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTopicsMetadataReply) ProtoMessage() {}

func (x *GetTopicsMetadataReply) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_seminar_v1_seminar_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTopicsMetadataReply.ProtoReflect.Descriptor instead.
func (*GetTopicsMetadataReply) Descriptor() ([]byte, []int) {
	return file_gateway_seminar_v1_seminar_proto_rawDescGZIP(), []int{12}
}

func (x *GetTopicsMetadataReply) GetTopics() []*TopicMetadata {
	if x != nil {
		return x.Topics
	}
	return nil
}

// 获取讨论主题的详细信息，进入讨论时加载
type GetTopicRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *GetTopicRequest) Reset() {
	*x = GetTopicRequest{}
	mi := &file_gateway_seminar_v1_seminar_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTopicRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTopicRequest) ProtoMessage() {}

func (x *GetTopicRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_seminar_v1_seminar_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTopicRequest.ProtoReflect.Descriptor instead.
func (*GetTopicRequest) Descriptor() ([]byte, []int) {
	return file_gateway_seminar_v1_seminar_proto_rawDescGZIP(), []int{13}
}

func (x *GetTopicRequest) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

type GetTopicReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Topic *Topic `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
}

func (x *GetTopicReply) Reset() {
	*x = GetTopicReply{}
	mi := &file_gateway_seminar_v1_seminar_proto_msgTypes[14]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTopicReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTopicReply) ProtoMessage() {}

func (x *GetTopicReply) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_seminar_v1_seminar_proto_msgTypes[14]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTopicReply.ProtoReflect.Descriptor instead.
func (*GetTopicReply) Descriptor() ([]byte, []int) {
	return file_gateway_seminar_v1_seminar_proto_rawDescGZIP(), []int{14}
}

func (x *GetTopicReply) GetTopic() *Topic {
	if x != nil {
		return x.Topic
	}
	return nil
}

var File_gateway_seminar_v1_seminar_proto protoreflect.FileDescriptor

var file_gateway_seminar_v1_seminar_proto_rawDesc = []byte{
	0x0a, 0x20, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x73, 0x65, 0x6d, 0x69, 0x6e, 0x61,
	0x72, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x6d, 0x69, 0x6e, 0x61, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0f, 0x57, 0x69, 0x74, 0x74, 0x67, 0x65, 0x6e, 0x73, 0x74, 0x65, 0x69, 0x6e,
	0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x5f, 0x0a, 0x0d, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x75, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x22,
	0x0a, 0x0c, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e,
	0x74, 0x73, 0x22, 0x4f, 0x0a, 0x06, 0x53, 0x70, 0x65, 0x65, 0x63, 0x68, 0x12, 0x10, 0x0a, 0x03,
	0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x19,
	0x0a, 0x08, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x72, 0x6f, 0x6c, 0x65, 0x55, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x22, 0xc3, 0x01, 0x0a, 0x05, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12,
	0x22, 0x0a, 0x0c, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61,
	0x6e, 0x74, 0x73, 0x12, 0x33, 0x0a, 0x08, 0x73, 0x70, 0x65, 0x65, 0x63, 0x68, 0x65, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x57, 0x69, 0x74, 0x74, 0x67, 0x65, 0x6e, 0x73,
	0x74, 0x65, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x70, 0x65, 0x65, 0x63, 0x68, 0x52, 0x08,
	0x73, 0x70, 0x65, 0x65, 0x63, 0x68, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1f,
	0x0a, 0x0b, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x68, 0x0a, 0x12, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12,
	0x22, 0x0a, 0x0c, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61,
	0x6e, 0x74, 0x73, 0x22, 0x24, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x70,
	0x69, 0x63, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x22, 0x26, 0x0a, 0x12, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69,
	0x64, 0x22, 0x2c, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x6f, 0x70, 0x69, 0x63,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x2d, 0x0a, 0x11, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x49, 0x64, 0x22, 0x52,
	0x0a, 0x0f, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x1e, 0x0a, 0x09, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x69, 0x6e,
	0x67, 0x12, 0x14, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x42, 0x09, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x22, 0x2c, 0x0a, 0x10, 0x53, 0x74, 0x6f, 0x70, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x49, 0x64,
	0x22, 0x2a, 0x0a, 0x0e, 0x53, 0x74, 0x6f, 0x70, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x30, 0x0a, 0x18,
	0x47, 0x65, 0x74, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x22, 0x50,
	0x0a, 0x16, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x36, 0x0a, 0x06, 0x74, 0x6f, 0x70, 0x69,
	0x63, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x57, 0x69, 0x74, 0x74, 0x67,
	0x65, 0x6e, 0x73, 0x74, 0x65, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x70, 0x69, 0x63,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x06, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x73,
	0x22, 0x23, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x75, 0x69, 0x64, 0x22, 0x3d, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x70, 0x69,
	0x63, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2c, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x57, 0x69, 0x74, 0x74, 0x67, 0x65, 0x6e, 0x73,
	0x74, 0x65, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x05, 0x74,
	0x6f, 0x70, 0x69, 0x63, 0x32, 0xcb, 0x05, 0x0a, 0x07, 0x53, 0x65, 0x6d, 0x69, 0x6e, 0x61, 0x72,
	0x12, 0x79, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x12,
	0x23, 0x2e, 0x57, 0x69, 0x74, 0x74, 0x67, 0x65, 0x6e, 0x73, 0x74, 0x65, 0x69, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x57, 0x69, 0x74, 0x74, 0x67, 0x65, 0x6e, 0x73, 0x74,
	0x65, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x70,
	0x69, 0x63, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x3a,
	0x01, 0x2a, 0x22, 0x17, 0x2f, 0x73, 0x65, 0x6d, 0x69, 0x6e, 0x61, 0x72, 0x2f, 0x74, 0x6f, 0x70,
	0x69, 0x63, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x8b, 0x01, 0x0a, 0x11,
	0x47, 0x65, 0x74, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x29, 0x2e, 0x57, 0x69, 0x74, 0x74, 0x67, 0x65, 0x6e, 0x73, 0x74, 0x65, 0x69, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x57,
	0x69, 0x74, 0x74, 0x67, 0x65, 0x6e, 0x73, 0x74, 0x65, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x3a, 0x01, 0x2a,
	0x22, 0x17, 0x2f, 0x73, 0x65, 0x6d, 0x69, 0x6e, 0x61, 0x72, 0x2f, 0x74, 0x6f, 0x70, 0x69, 0x63,
	0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x6f, 0x0a, 0x08, 0x47, 0x65, 0x74,
	0x54, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x20, 0x2e, 0x57, 0x69, 0x74, 0x74, 0x67, 0x65, 0x6e, 0x73,
	0x74, 0x65, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x70, 0x69, 0x63,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x57, 0x69, 0x74, 0x74, 0x67, 0x65,
	0x6e, 0x73, 0x74, 0x65, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x70,
	0x69, 0x63, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x3a,
	0x01, 0x2a, 0x22, 0x16, 0x2f, 0x73, 0x65, 0x6d, 0x69, 0x6e, 0x61, 0x72, 0x2f, 0x74, 0x6f, 0x70,
	0x69, 0x63, 0x2f, 0x67, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x79, 0x0a, 0x0b, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x23, 0x2e, 0x57, 0x69, 0x74, 0x74,
	0x67, 0x65, 0x6e, 0x73, 0x74, 0x65, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21,
	0x2e, 0x57, 0x69, 0x74, 0x74, 0x67, 0x65, 0x6e, 0x73, 0x74, 0x65, 0x69, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x3a, 0x01, 0x2a, 0x22, 0x17, 0x2f, 0x73,
	0x65, 0x6d, 0x69, 0x6e, 0x61, 0x72, 0x2f, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x2f, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x56, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x6f,
	0x70, 0x69, 0x63, 0x12, 0x22, 0x2e, 0x57, 0x69, 0x74, 0x74, 0x67, 0x65, 0x6e, 0x73, 0x74, 0x65,
	0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x6f, 0x70, 0x69, 0x63,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x57, 0x69, 0x74, 0x74, 0x67, 0x65,
	0x6e, 0x73, 0x74, 0x65, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54,
	0x6f, 0x70, 0x69, 0x63, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x30, 0x01, 0x12, 0x73, 0x0a,
	0x09, 0x53, 0x74, 0x6f, 0x70, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x21, 0x2e, 0x57, 0x69, 0x74,
	0x74, 0x67, 0x65, 0x6e, 0x73, 0x74, 0x65, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x6f,
	0x70, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e,
	0x57, 0x69, 0x74, 0x74, 0x67, 0x65, 0x6e, 0x73, 0x74, 0x65, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x74, 0x6f, 0x70, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x22,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x3a, 0x01, 0x2a, 0x22, 0x17, 0x2f, 0x73, 0x65, 0x6d, 0x69,
	0x6e, 0x61, 0x72, 0x2f, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x2f, 0x73, 0x74, 0x6f, 0x70, 0x70, 0x69,
	0x6e, 0x67, 0x42, 0x28, 0x5a, 0x26, 0x57, 0x69, 0x74, 0x74, 0x67, 0x65, 0x6e, 0x73, 0x74, 0x65,
	0x69, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x73,
	0x65, 0x6d, 0x69, 0x6e, 0x61, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gateway_seminar_v1_seminar_proto_rawDescOnce sync.Once
	file_gateway_seminar_v1_seminar_proto_rawDescData = file_gateway_seminar_v1_seminar_proto_rawDesc
)

func file_gateway_seminar_v1_seminar_proto_rawDescGZIP() []byte {
	file_gateway_seminar_v1_seminar_proto_rawDescOnce.Do(func() {
		file_gateway_seminar_v1_seminar_proto_rawDescData = protoimpl.X.CompressGZIP(file_gateway_seminar_v1_seminar_proto_rawDescData)
	})
	return file_gateway_seminar_v1_seminar_proto_rawDescData
}

var file_gateway_seminar_v1_seminar_proto_msgTypes = make([]protoimpl.MessageInfo, 15)
var file_gateway_seminar_v1_seminar_proto_goTypes = []any{
	(*TopicMetadata)(nil),            // 0: Wittgenstein.v1.TopicMetadata
	(*Speech)(nil),                   // 1: Wittgenstein.v1.Speech
	(*Topic)(nil),                    // 2: Wittgenstein.v1.Topic
	(*CreateTopicRequest)(nil),       // 3: Wittgenstein.v1.CreateTopicRequest
	(*CreateTopicReply)(nil),         // 4: Wittgenstein.v1.CreateTopicReply
	(*DeleteTopicRequest)(nil),       // 5: Wittgenstein.v1.DeleteTopicRequest
	(*DeleteTopicReply)(nil),         // 6: Wittgenstein.v1.DeleteTopicReply
	(*StartTopicRequest)(nil),        // 7: Wittgenstein.v1.StartTopicRequest
	(*StartTopicReply)(nil),          // 8: Wittgenstein.v1.StartTopicReply
	(*StopTopicRequest)(nil),         // 9: Wittgenstein.v1.StopTopicRequest
	(*StopTopicReply)(nil),           // 10: Wittgenstein.v1.StopTopicReply
	(*GetTopicsMetadataRequest)(nil), // 11: Wittgenstein.v1.GetTopicsMetadataRequest
	(*GetTopicsMetadataReply)(nil),   // 12: Wittgenstein.v1.GetTopicsMetadataReply
	(*GetTopicRequest)(nil),          // 13: Wittgenstein.v1.GetTopicRequest
	(*GetTopicReply)(nil),            // 14: Wittgenstein.v1.GetTopicReply
}
var file_gateway_seminar_v1_seminar_proto_depIdxs = []int32{
	1,  // 0: Wittgenstein.v1.Topic.speeches:type_name -> Wittgenstein.v1.Speech
	0,  // 1: Wittgenstein.v1.GetTopicsMetadataReply.topics:type_name -> Wittgenstein.v1.TopicMetadata
	2,  // 2: Wittgenstein.v1.GetTopicReply.topic:type_name -> Wittgenstein.v1.Topic
	3,  // 3: Wittgenstein.v1.Seminar.CreateTopic:input_type -> Wittgenstein.v1.CreateTopicRequest
	11, // 4: Wittgenstein.v1.Seminar.GetTopicsMetadata:input_type -> Wittgenstein.v1.GetTopicsMetadataRequest
	13, // 5: Wittgenstein.v1.Seminar.GetTopic:input_type -> Wittgenstein.v1.GetTopicRequest
	5,  // 6: Wittgenstein.v1.Seminar.DeleteTopic:input_type -> Wittgenstein.v1.DeleteTopicRequest
	7,  // 7: Wittgenstein.v1.Seminar.StartTopic:input_type -> Wittgenstein.v1.StartTopicRequest
	9,  // 8: Wittgenstein.v1.Seminar.StopTopic:input_type -> Wittgenstein.v1.StopTopicRequest
	4,  // 9: Wittgenstein.v1.Seminar.CreateTopic:output_type -> Wittgenstein.v1.CreateTopicReply
	12, // 10: Wittgenstein.v1.Seminar.GetTopicsMetadata:output_type -> Wittgenstein.v1.GetTopicsMetadataReply
	14, // 11: Wittgenstein.v1.Seminar.GetTopic:output_type -> Wittgenstein.v1.GetTopicReply
	6,  // 12: Wittgenstein.v1.Seminar.DeleteTopic:output_type -> Wittgenstein.v1.DeleteTopicReply
	8,  // 13: Wittgenstein.v1.Seminar.StartTopic:output_type -> Wittgenstein.v1.StartTopicReply
	10, // 14: Wittgenstein.v1.Seminar.StopTopic:output_type -> Wittgenstein.v1.StopTopicReply
	9,  // [9:15] is the sub-list for method output_type
	3,  // [3:9] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_gateway_seminar_v1_seminar_proto_init() }
func file_gateway_seminar_v1_seminar_proto_init() {
	if File_gateway_seminar_v1_seminar_proto != nil {
		return
	}
	file_gateway_seminar_v1_seminar_proto_msgTypes[8].OneofWrappers = []any{
		(*StartTopicReply_Reasoning)(nil),
		(*StartTopicReply_Text)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_gateway_seminar_v1_seminar_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   15,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gateway_seminar_v1_seminar_proto_goTypes,
		DependencyIndexes: file_gateway_seminar_v1_seminar_proto_depIdxs,
		MessageInfos:      file_gateway_seminar_v1_seminar_proto_msgTypes,
	}.Build()
	File_gateway_seminar_v1_seminar_proto = out.File
	file_gateway_seminar_v1_seminar_proto_rawDesc = nil
	file_gateway_seminar_v1_seminar_proto_goTypes = nil
	file_gateway_seminar_v1_seminar_proto_depIdxs = nil
}
