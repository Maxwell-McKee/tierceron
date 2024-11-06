// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.21.12
// source: trcshtalksdk/trcshtalksdk.proto

package trcshtalksdk

import (
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

type Diagnostics int32

const (
	// if a diagnostic deprecates, comment it out
	// if adding a diagnostic, append to the end incrementing integer
	Diagnostics_ALL          Diagnostics = 0 // Default
	Diagnostics_HEALTH_CHECK Diagnostics = 1 // future plugins
)

// Enum value maps for Diagnostics.
var (
	Diagnostics_name = map[int32]string{
		0: "ALL",
		1: "HEALTH_CHECK",
	}
	Diagnostics_value = map[string]int32{
		"ALL":          0,
		"HEALTH_CHECK": 1,
	}
)

func (x Diagnostics) Enum() *Diagnostics {
	p := new(Diagnostics)
	*p = x
	return p
}

func (x Diagnostics) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Diagnostics) Descriptor() protoreflect.EnumDescriptor {
	return file_trcshtalksdk_trcshtalksdk_proto_enumTypes[0].Descriptor()
}

func (Diagnostics) Type() protoreflect.EnumType {
	return &file_trcshtalksdk_trcshtalksdk_proto_enumTypes[0]
}

func (x Diagnostics) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Diagnostics.Descriptor instead.
func (Diagnostics) EnumDescriptor() ([]byte, []int) {
	return file_trcshtalksdk_trcshtalksdk_proto_rawDescGZIP(), []int{0}
}

type PluginQuery int32

const (
	PluginQuery_ACTIVE_COUNT PluginQuery = 0
)

// Enum value maps for PluginQuery.
var (
	PluginQuery_name = map[int32]string{
		0: "ACTIVE_COUNT",
	}
	PluginQuery_value = map[string]int32{
		"ACTIVE_COUNT": 0,
	}
)

func (x PluginQuery) Enum() *PluginQuery {
	p := new(PluginQuery)
	*p = x
	return p
}

func (x PluginQuery) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PluginQuery) Descriptor() protoreflect.EnumDescriptor {
	return file_trcshtalksdk_trcshtalksdk_proto_enumTypes[1].Descriptor()
}

func (PluginQuery) Type() protoreflect.EnumType {
	return &file_trcshtalksdk_trcshtalksdk_proto_enumTypes[1]
}

func (x PluginQuery) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PluginQuery.Descriptor instead.
func (PluginQuery) EnumDescriptor() ([]byte, []int) {
	return file_trcshtalksdk_trcshtalksdk_proto_rawDescGZIP(), []int{1}
}

type DiagnosticRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageId   string        `protobuf:"bytes,1,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	Diagnostics []Diagnostics `protobuf:"varint,2,rep,packed,name=diagnostics,proto3,enum=trcshtalksdk.Diagnostics" json:"diagnostics,omitempty"`
	TenantId    string        `protobuf:"bytes,3,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
	Data        []string      `protobuf:"bytes,4,rep,name=data,proto3" json:"data,omitempty"`
	Queries     []PluginQuery `protobuf:"varint,5,rep,packed,name=queries,proto3,enum=trcshtalksdk.PluginQuery" json:"queries,omitempty"`
}

func (x *DiagnosticRequest) Reset() {
	*x = DiagnosticRequest{}
	mi := &file_trcshtalksdk_trcshtalksdk_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DiagnosticRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiagnosticRequest) ProtoMessage() {}

func (x *DiagnosticRequest) ProtoReflect() protoreflect.Message {
	mi := &file_trcshtalksdk_trcshtalksdk_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiagnosticRequest.ProtoReflect.Descriptor instead.
func (*DiagnosticRequest) Descriptor() ([]byte, []int) {
	return file_trcshtalksdk_trcshtalksdk_proto_rawDescGZIP(), []int{0}
}

func (x *DiagnosticRequest) GetMessageId() string {
	if x != nil {
		return x.MessageId
	}
	return ""
}

func (x *DiagnosticRequest) GetDiagnostics() []Diagnostics {
	if x != nil {
		return x.Diagnostics
	}
	return nil
}

func (x *DiagnosticRequest) GetTenantId() string {
	if x != nil {
		return x.TenantId
	}
	return ""
}

func (x *DiagnosticRequest) GetData() []string {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *DiagnosticRequest) GetQueries() []PluginQuery {
	if x != nil {
		return x.Queries
	}
	return nil
}

type DiagnosticResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageId string `protobuf:"bytes,1,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	Results   string `protobuf:"bytes,2,opt,name=results,proto3" json:"results,omitempty"`
}

func (x *DiagnosticResponse) Reset() {
	*x = DiagnosticResponse{}
	mi := &file_trcshtalksdk_trcshtalksdk_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DiagnosticResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiagnosticResponse) ProtoMessage() {}

func (x *DiagnosticResponse) ProtoReflect() protoreflect.Message {
	mi := &file_trcshtalksdk_trcshtalksdk_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiagnosticResponse.ProtoReflect.Descriptor instead.
func (*DiagnosticResponse) Descriptor() ([]byte, []int) {
	return file_trcshtalksdk_trcshtalksdk_proto_rawDescGZIP(), []int{1}
}

func (x *DiagnosticResponse) GetMessageId() string {
	if x != nil {
		return x.MessageId
	}
	return ""
}

func (x *DiagnosticResponse) GetResults() string {
	if x != nil {
		return x.Results
	}
	return ""
}

var File_trcshtalksdk_trcshtalksdk_proto protoreflect.FileDescriptor

var file_trcshtalksdk_trcshtalksdk_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x74, 0x72, 0x63, 0x73, 0x68, 0x74, 0x61, 0x6c, 0x6b, 0x73, 0x64, 0x6b, 0x2f, 0x74,
	0x72, 0x63, 0x73, 0x68, 0x74, 0x61, 0x6c, 0x6b, 0x73, 0x64, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0c, 0x74, 0x72, 0x63, 0x73, 0x68, 0x74, 0x61, 0x6c, 0x6b, 0x73, 0x64, 0x6b, 0x22,
	0xd5, 0x01, 0x0a, 0x11, 0x44, 0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x49, 0x64, 0x12, 0x3b, 0x0a, 0x0b, 0x64, 0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x74,
	0x69, 0x63, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x74, 0x72, 0x63, 0x73,
	0x68, 0x74, 0x61, 0x6c, 0x6b, 0x73, 0x64, 0x6b, 0x2e, 0x44, 0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73,
	0x74, 0x69, 0x63, 0x73, 0x52, 0x0b, 0x64, 0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63,
	0x73, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x33, 0x0a, 0x07, 0x71, 0x75, 0x65, 0x72, 0x69, 0x65, 0x73, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x74, 0x72, 0x63, 0x73, 0x68, 0x74, 0x61, 0x6c, 0x6b, 0x73,
	0x64, 0x6b, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x07,
	0x71, 0x75, 0x65, 0x72, 0x69, 0x65, 0x73, 0x22, 0x4d, 0x0a, 0x12, 0x44, 0x69, 0x61, 0x67, 0x6e,
	0x6f, 0x73, 0x74, 0x69, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x2a, 0x28, 0x0a, 0x0b, 0x44, 0x69, 0x61, 0x67, 0x6e, 0x6f,
	0x73, 0x74, 0x69, 0x63, 0x73, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x4c, 0x4c, 0x10, 0x00, 0x12, 0x10,
	0x0a, 0x0c, 0x48, 0x45, 0x41, 0x4c, 0x54, 0x48, 0x5f, 0x43, 0x48, 0x45, 0x43, 0x4b, 0x10, 0x01,
	0x2a, 0x1f, 0x0a, 0x0b, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x0c, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x10,
	0x00, 0x32, 0x67, 0x0a, 0x10, 0x54, 0x72, 0x63, 0x73, 0x68, 0x54, 0x61, 0x6c, 0x6b, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x53, 0x0a, 0x0e, 0x52, 0x75, 0x6e, 0x44, 0x69, 0x61, 0x67,
	0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x73, 0x12, 0x1f, 0x2e, 0x74, 0x72, 0x63, 0x73, 0x68, 0x74,
	0x61, 0x6c, 0x6b, 0x73, 0x64, 0x6b, 0x2e, 0x44, 0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69,
	0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x74, 0x72, 0x63, 0x73, 0x68,
	0x74, 0x61, 0x6c, 0x6b, 0x73, 0x64, 0x6b, 0x2e, 0x44, 0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x74,
	0x69, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x57, 0x5a, 0x55, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x72, 0x69, 0x6d, 0x62, 0x6c, 0x65,
	0x2d, 0x6f, 0x73, 0x73, 0x2f, 0x74, 0x69, 0x65, 0x72, 0x63, 0x65, 0x72, 0x6f, 0x6e, 0x2f, 0x69,
	0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x74, 0x72, 0x63, 0x73,
	0x68, 0x68, 0x69, 0x76, 0x65, 0x2f, 0x74, 0x72, 0x63, 0x73, 0x68, 0x6b, 0x2f, 0x74, 0x72, 0x63,
	0x73, 0x68, 0x74, 0x61, 0x6c, 0x6b, 0x2f, 0x74, 0x72, 0x63, 0x73, 0x68, 0x74, 0x61, 0x6c, 0x6b,
	0x73, 0x64, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_trcshtalksdk_trcshtalksdk_proto_rawDescOnce sync.Once
	file_trcshtalksdk_trcshtalksdk_proto_rawDescData = file_trcshtalksdk_trcshtalksdk_proto_rawDesc
)

func file_trcshtalksdk_trcshtalksdk_proto_rawDescGZIP() []byte {
	file_trcshtalksdk_trcshtalksdk_proto_rawDescOnce.Do(func() {
		file_trcshtalksdk_trcshtalksdk_proto_rawDescData = protoimpl.X.CompressGZIP(file_trcshtalksdk_trcshtalksdk_proto_rawDescData)
	})
	return file_trcshtalksdk_trcshtalksdk_proto_rawDescData
}

var file_trcshtalksdk_trcshtalksdk_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_trcshtalksdk_trcshtalksdk_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_trcshtalksdk_trcshtalksdk_proto_goTypes = []any{
	(Diagnostics)(0),           // 0: trcshtalksdk.Diagnostics
	(PluginQuery)(0),           // 1: trcshtalksdk.PluginQuery
	(*DiagnosticRequest)(nil),  // 2: trcshtalksdk.DiagnosticRequest
	(*DiagnosticResponse)(nil), // 3: trcshtalksdk.DiagnosticResponse
}
var file_trcshtalksdk_trcshtalksdk_proto_depIdxs = []int32{
	0, // 0: trcshtalksdk.DiagnosticRequest.diagnostics:type_name -> trcshtalksdk.Diagnostics
	1, // 1: trcshtalksdk.DiagnosticRequest.queries:type_name -> trcshtalksdk.PluginQuery
	2, // 2: trcshtalksdk.TrcshTalkService.RunDiagnostics:input_type -> trcshtalksdk.DiagnosticRequest
	3, // 3: trcshtalksdk.TrcshTalkService.RunDiagnostics:output_type -> trcshtalksdk.DiagnosticResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_trcshtalksdk_trcshtalksdk_proto_init() }
func file_trcshtalksdk_trcshtalksdk_proto_init() {
	if File_trcshtalksdk_trcshtalksdk_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_trcshtalksdk_trcshtalksdk_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_trcshtalksdk_trcshtalksdk_proto_goTypes,
		DependencyIndexes: file_trcshtalksdk_trcshtalksdk_proto_depIdxs,
		EnumInfos:         file_trcshtalksdk_trcshtalksdk_proto_enumTypes,
		MessageInfos:      file_trcshtalksdk_trcshtalksdk_proto_msgTypes,
	}.Build()
	File_trcshtalksdk_trcshtalksdk_proto = out.File
	file_trcshtalksdk_trcshtalksdk_proto_rawDesc = nil
	file_trcshtalksdk_trcshtalksdk_proto_goTypes = nil
	file_trcshtalksdk_trcshtalksdk_proto_depIdxs = nil
}