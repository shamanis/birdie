// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: api/birdie.proto

package birdie

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ResponseStatus int32

const (
	ResponseStatus_STATUS_OK    ResponseStatus = 0
	ResponseStatus_STATUS_ERROR ResponseStatus = 1
)

// Enum value maps for ResponseStatus.
var (
	ResponseStatus_name = map[int32]string{
		0: "STATUS_OK",
		1: "STATUS_ERROR",
	}
	ResponseStatus_value = map[string]int32{
		"STATUS_OK":    0,
		"STATUS_ERROR": 1,
	}
)

func (x ResponseStatus) Enum() *ResponseStatus {
	p := new(ResponseStatus)
	*p = x
	return p
}

func (x ResponseStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResponseStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_api_birdie_proto_enumTypes[0].Descriptor()
}

func (ResponseStatus) Type() protoreflect.EnumType {
	return &file_api_birdie_proto_enumTypes[0]
}

func (x ResponseStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ResponseStatus.Descriptor instead.
func (ResponseStatus) EnumDescriptor() ([]byte, []int) {
	return file_api_birdie_proto_rawDescGZIP(), []int{0}
}

type BaseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status ResponseStatus `protobuf:"varint,1,opt,name=status,proto3,enum=ResponseStatus" json:"status,omitempty"`
	Error  string         `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *BaseResponse) Reset() {
	*x = BaseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_birdie_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseResponse) ProtoMessage() {}

func (x *BaseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_birdie_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseResponse.ProtoReflect.Descriptor instead.
func (*BaseResponse) Descriptor() ([]byte, []int) {
	return file_api_birdie_proto_rawDescGZIP(), []int{0}
}

func (x *BaseResponse) GetStatus() ResponseStatus {
	if x != nil {
		return x.Status
	}
	return ResponseStatus_STATUS_OK
}

func (x *BaseResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type Entry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Ttl   uint64 `protobuf:"varint,3,opt,name=ttl,proto3" json:"ttl,omitempty"`
}

func (x *Entry) Reset() {
	*x = Entry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_birdie_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Entry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Entry) ProtoMessage() {}

func (x *Entry) ProtoReflect() protoreflect.Message {
	mi := &file_api_birdie_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Entry.ProtoReflect.Descriptor instead.
func (*Entry) Descriptor() ([]byte, []int) {
	return file_api_birdie_proto_rawDescGZIP(), []int{1}
}

func (x *Entry) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Entry) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *Entry) GetTtl() uint64 {
	if x != nil {
		return x.Ttl
	}
	return 0
}

type Entries struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entry []*Entry `protobuf:"bytes,1,rep,name=entry,proto3" json:"entry,omitempty"`
}

func (x *Entries) Reset() {
	*x = Entries{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_birdie_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Entries) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Entries) ProtoMessage() {}

func (x *Entries) ProtoReflect() protoreflect.Message {
	mi := &file_api_birdie_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Entries.ProtoReflect.Descriptor instead.
func (*Entries) Descriptor() ([]byte, []int) {
	return file_api_birdie_proto_rawDescGZIP(), []int{2}
}

func (x *Entries) GetEntry() []*Entry {
	if x != nil {
		return x.Entry
	}
	return nil
}

type LoadQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *LoadQuery) Reset() {
	*x = LoadQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_birdie_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoadQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoadQuery) ProtoMessage() {}

func (x *LoadQuery) ProtoReflect() protoreflect.Message {
	mi := &file_api_birdie_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoadQuery.ProtoReflect.Descriptor instead.
func (*LoadQuery) Descriptor() ([]byte, []int) {
	return file_api_birdie_proto_rawDescGZIP(), []int{3}
}

func (x *LoadQuery) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type SearchQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pattern string `protobuf:"bytes,1,opt,name=pattern,proto3" json:"pattern,omitempty"`
}

func (x *SearchQuery) Reset() {
	*x = SearchQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_birdie_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchQuery) ProtoMessage() {}

func (x *SearchQuery) ProtoReflect() protoreflect.Message {
	mi := &file_api_birdie_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchQuery.ProtoReflect.Descriptor instead.
func (*SearchQuery) Descriptor() ([]byte, []int) {
	return file_api_birdie_proto_rawDescGZIP(), []int{4}
}

func (x *SearchQuery) GetPattern() string {
	if x != nil {
		return x.Pattern
	}
	return ""
}

var File_api_birdie_proto protoreflect.FileDescriptor

var file_api_birdie_proto_rawDesc = []byte{
	0x0a, 0x10, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x69, 0x72, 0x64, 0x69, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x4d, 0x0a, 0x0c, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x27, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x0f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x41,
	0x0a, 0x05, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x74, 0x74, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x74, 0x74,
	0x6c, 0x22, 0x27, 0x0a, 0x07, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x12, 0x1c, 0x0a, 0x05,
	0x65, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x05, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x22, 0x1d, 0x0a, 0x09, 0x4c, 0x6f,
	0x61, 0x64, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x27, 0x0a, 0x0b, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x74, 0x74,
	0x65, 0x72, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x74, 0x74, 0x65,
	0x72, 0x6e, 0x2a, 0x31, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x4f,
	0x4b, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x45, 0x52,
	0x52, 0x4f, 0x52, 0x10, 0x01, 0x32, 0xc5, 0x01, 0x0a, 0x06, 0x42, 0x69, 0x72, 0x64, 0x69, 0x65,
	0x12, 0x2f, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x0d, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x20, 0x0a, 0x05, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x06, 0x2e, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x1a, 0x0d, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x26, 0x0a, 0x09, 0x42, 0x75, 0x6c, 0x6b, 0x53, 0x74, 0x6f, 0x72, 0x65,
	0x12, 0x08, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x1a, 0x0d, 0x2e, 0x42, 0x61, 0x73,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x1c, 0x0a, 0x04, 0x4c,
	0x6f, 0x61, 0x64, 0x12, 0x0a, 0x2e, 0x4c, 0x6f, 0x61, 0x64, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a,
	0x06, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x22, 0x00, 0x12, 0x22, 0x0a, 0x06, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x12, 0x0c, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x1a, 0x08, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x22, 0x00, 0x42, 0x15, 0x5a,
	0x13, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x62, 0x69,
	0x72, 0x64, 0x69, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_birdie_proto_rawDescOnce sync.Once
	file_api_birdie_proto_rawDescData = file_api_birdie_proto_rawDesc
)

func file_api_birdie_proto_rawDescGZIP() []byte {
	file_api_birdie_proto_rawDescOnce.Do(func() {
		file_api_birdie_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_birdie_proto_rawDescData)
	})
	return file_api_birdie_proto_rawDescData
}

var file_api_birdie_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_birdie_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_birdie_proto_goTypes = []interface{}{
	(ResponseStatus)(0),   // 0: ResponseStatus
	(*BaseResponse)(nil),  // 1: BaseResponse
	(*Entry)(nil),         // 2: Entry
	(*Entries)(nil),       // 3: Entries
	(*LoadQuery)(nil),     // 4: LoadQuery
	(*SearchQuery)(nil),   // 5: SearchQuery
	(*emptypb.Empty)(nil), // 6: google.protobuf.Empty
}
var file_api_birdie_proto_depIdxs = []int32{
	0, // 0: BaseResponse.status:type_name -> ResponseStatus
	2, // 1: Entries.entry:type_name -> Entry
	6, // 2: Birdie.Ping:input_type -> google.protobuf.Empty
	2, // 3: Birdie.Store:input_type -> Entry
	3, // 4: Birdie.BulkStore:input_type -> Entries
	4, // 5: Birdie.Load:input_type -> LoadQuery
	5, // 6: Birdie.Search:input_type -> SearchQuery
	1, // 7: Birdie.Ping:output_type -> BaseResponse
	1, // 8: Birdie.Store:output_type -> BaseResponse
	1, // 9: Birdie.BulkStore:output_type -> BaseResponse
	2, // 10: Birdie.Load:output_type -> Entry
	3, // 11: Birdie.Search:output_type -> Entries
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_birdie_proto_init() }
func file_api_birdie_proto_init() {
	if File_api_birdie_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_birdie_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseResponse); i {
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
		file_api_birdie_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Entry); i {
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
		file_api_birdie_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Entries); i {
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
		file_api_birdie_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoadQuery); i {
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
		file_api_birdie_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchQuery); i {
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
			RawDescriptor: file_api_birdie_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_birdie_proto_goTypes,
		DependencyIndexes: file_api_birdie_proto_depIdxs,
		EnumInfos:         file_api_birdie_proto_enumTypes,
		MessageInfos:      file_api_birdie_proto_msgTypes,
	}.Build()
	File_api_birdie_proto = out.File
	file_api_birdie_proto_rawDesc = nil
	file_api_birdie_proto_goTypes = nil
	file_api_birdie_proto_depIdxs = nil
}
