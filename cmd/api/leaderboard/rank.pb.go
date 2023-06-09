// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: pkg/protos/rank.proto

package leaderboard

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type EndgameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Score float64 `protobuf:"fixed64,1,opt,name=score,proto3" json:"score,omitempty"`
}

func (x *EndgameRequest) Reset() {
	*x = EndgameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_protos_rank_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EndgameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndgameRequest) ProtoMessage() {}

func (x *EndgameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_protos_rank_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndgameRequest.ProtoReflect.Descriptor instead.
func (*EndgameRequest) Descriptor() ([]byte, []int) {
	return file_pkg_protos_rank_proto_rawDescGZIP(), []int{0}
}

func (x *EndgameRequest) GetScore() float64 {
	if x != nil {
		return x.Score
	}
	return 0
}

type EndgameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status    string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Timestamp string `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *EndgameResponse) Reset() {
	*x = EndgameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_protos_rank_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EndgameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndgameResponse) ProtoMessage() {}

func (x *EndgameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_protos_rank_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndgameResponse.ProtoReflect.Descriptor instead.
func (*EndgameResponse) Descriptor() ([]byte, []int) {
	return file_pkg_protos_rank_proto_rawDescGZIP(), []int{1}
}

func (x *EndgameResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *EndgameResponse) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

type LeaderBoardRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LeaderBoardRequest) Reset() {
	*x = LeaderBoardRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_protos_rank_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeaderBoardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeaderBoardRequest) ProtoMessage() {}

func (x *LeaderBoardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_protos_rank_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeaderBoardRequest.ProtoReflect.Descriptor instead.
func (*LeaderBoardRequest) Descriptor() ([]byte, []int) {
	return file_pkg_protos_rank_proto_rawDescGZIP(), []int{2}
}

type LeaderBoardResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status    string    `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Timestamp string    `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Result    []*Result `protobuf:"bytes,3,rep,name=result,proto3" json:"result,omitempty"`
}

func (x *LeaderBoardResponse) Reset() {
	*x = LeaderBoardResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_protos_rank_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeaderBoardResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeaderBoardResponse) ProtoMessage() {}

func (x *LeaderBoardResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_protos_rank_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeaderBoardResponse.ProtoReflect.Descriptor instead.
func (*LeaderBoardResponse) Descriptor() ([]byte, []int) {
	return file_pkg_protos_rank_proto_rawDescGZIP(), []int{3}
}

func (x *LeaderBoardResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *LeaderBoardResponse) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

func (x *LeaderBoardResponse) GetResult() []*Result {
	if x != nil {
		return x.Result
	}
	return nil
}

type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       *wrapperspb.Int64Value `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Username string                 `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Score    float64                `protobuf:"fixed64,3,opt,name=score,proto3" json:"score,omitempty"`
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_protos_rank_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_protos_rank_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_pkg_protos_rank_proto_rawDescGZIP(), []int{4}
}

func (x *Result) GetId() *wrapperspb.Int64Value {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Result) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Result) GetScore() float64 {
	if x != nil {
		return x.Score
	}
	return 0
}

var File_pkg_protos_rank_proto protoreflect.FileDescriptor

var file_pkg_protos_rank_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x72, 0x61, 0x6e,
	0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x26, 0x0a, 0x0e, 0x45, 0x6e, 0x64, 0x67, 0x61, 0x6d, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x22, 0x47, 0x0a,
	0x0f, 0x45, 0x6e, 0x64, 0x67, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x14, 0x0a, 0x12, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x42, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x6c, 0x0a, 0x13,
	0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x1f, 0x0a, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x67, 0x0a, 0x06, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x12, 0x2b, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x73, 0x63,
	0x6f, 0x72, 0x65, 0x32, 0x9a, 0x01, 0x0a, 0x12, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x42, 0x6f,
	0x61, 0x72, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x48, 0x0a, 0x07, 0x45, 0x6e,
	0x64, 0x67, 0x61, 0x6d, 0x65, 0x12, 0x0f, 0x2e, 0x45, 0x6e, 0x64, 0x67, 0x61, 0x6d, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x45, 0x6e, 0x64, 0x67, 0x61, 0x6d, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14,
	0x3a, 0x01, 0x2a, 0x22, 0x0f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x64,
	0x67, 0x61, 0x6d, 0x65, 0x12, 0x3a, 0x0a, 0x0b, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x42, 0x6f,
	0x61, 0x72, 0x64, 0x12, 0x13, 0x2e, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x42, 0x6f, 0x61, 0x72,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x4c, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x15, 0x5a, 0x13, 0x63, 0x6d, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6c, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_protos_rank_proto_rawDescOnce sync.Once
	file_pkg_protos_rank_proto_rawDescData = file_pkg_protos_rank_proto_rawDesc
)

func file_pkg_protos_rank_proto_rawDescGZIP() []byte {
	file_pkg_protos_rank_proto_rawDescOnce.Do(func() {
		file_pkg_protos_rank_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_protos_rank_proto_rawDescData)
	})
	return file_pkg_protos_rank_proto_rawDescData
}

var file_pkg_protos_rank_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_pkg_protos_rank_proto_goTypes = []interface{}{
	(*EndgameRequest)(nil),        // 0: EndgameRequest
	(*EndgameResponse)(nil),       // 1: EndgameResponse
	(*LeaderBoardRequest)(nil),    // 2: LeaderBoardRequest
	(*LeaderBoardResponse)(nil),   // 3: LeaderBoardResponse
	(*Result)(nil),                // 4: Result
	(*wrapperspb.Int64Value)(nil), // 5: google.protobuf.Int64Value
}
var file_pkg_protos_rank_proto_depIdxs = []int32{
	4, // 0: LeaderBoardResponse.result:type_name -> Result
	5, // 1: Result.id:type_name -> google.protobuf.Int64Value
	0, // 2: LeaderBoardService.Endgame:input_type -> EndgameRequest
	2, // 3: LeaderBoardService.LeaderBoard:input_type -> LeaderBoardRequest
	1, // 4: LeaderBoardService.Endgame:output_type -> EndgameResponse
	3, // 5: LeaderBoardService.LeaderBoard:output_type -> LeaderBoardResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pkg_protos_rank_proto_init() }
func file_pkg_protos_rank_proto_init() {
	if File_pkg_protos_rank_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_protos_rank_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EndgameRequest); i {
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
		file_pkg_protos_rank_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EndgameResponse); i {
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
		file_pkg_protos_rank_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeaderBoardRequest); i {
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
		file_pkg_protos_rank_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeaderBoardResponse); i {
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
		file_pkg_protos_rank_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Result); i {
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
			RawDescriptor: file_pkg_protos_rank_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_protos_rank_proto_goTypes,
		DependencyIndexes: file_pkg_protos_rank_proto_depIdxs,
		MessageInfos:      file_pkg_protos_rank_proto_msgTypes,
	}.Build()
	File_pkg_protos_rank_proto = out.File
	file_pkg_protos_rank_proto_rawDesc = nil
	file_pkg_protos_rank_proto_goTypes = nil
	file_pkg_protos_rank_proto_depIdxs = nil
}
