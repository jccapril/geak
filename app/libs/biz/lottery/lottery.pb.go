// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.15.1
// source: lottery.proto

package lottery

import (
	m "biz/m"
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

type GetLastestLotteryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ssq bool `protobuf:"varint,1,opt,name=ssq,proto3" json:"ssq,omitempty"`
	Dlt bool `protobuf:"varint,2,opt,name=dlt,proto3" json:"dlt,omitempty"`
}

func (x *GetLastestLotteryRequest) Reset() {
	*x = GetLastestLotteryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lottery_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLastestLotteryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLastestLotteryRequest) ProtoMessage() {}

func (x *GetLastestLotteryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lottery_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLastestLotteryRequest.ProtoReflect.Descriptor instead.
func (*GetLastestLotteryRequest) Descriptor() ([]byte, []int) {
	return file_lottery_proto_rawDescGZIP(), []int{0}
}

func (x *GetLastestLotteryRequest) GetSsq() bool {
	if x != nil {
		return x.Ssq
	}
	return false
}

func (x *GetLastestLotteryRequest) GetDlt() bool {
	if x != nil {
		return x.Dlt
	}
	return false
}

type GetLastestLotteryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrCode int64        `protobuf:"varint,1,opt,name=errCode,proto3" json:"errCode,omitempty"`
	ErrMsg  string       `protobuf:"bytes,2,opt,name=errMsg,proto3" json:"errMsg,omitempty"`
	Lottery []*m.Lottery `protobuf:"bytes,3,rep,name=lottery,proto3" json:"lottery,omitempty"`
}

func (x *GetLastestLotteryResponse) Reset() {
	*x = GetLastestLotteryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lottery_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLastestLotteryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLastestLotteryResponse) ProtoMessage() {}

func (x *GetLastestLotteryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_lottery_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLastestLotteryResponse.ProtoReflect.Descriptor instead.
func (*GetLastestLotteryResponse) Descriptor() ([]byte, []int) {
	return file_lottery_proto_rawDescGZIP(), []int{1}
}

func (x *GetLastestLotteryResponse) GetErrCode() int64 {
	if x != nil {
		return x.ErrCode
	}
	return 0
}

func (x *GetLastestLotteryResponse) GetErrMsg() string {
	if x != nil {
		return x.ErrMsg
	}
	return ""
}

func (x *GetLastestLotteryResponse) GetLottery() []*m.Lottery {
	if x != nil {
		return x.Lottery
	}
	return nil
}

var File_lottery_proto protoreflect.FileDescriptor

var file_lottery_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6c, 0x6f, 0x74, 0x74, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0c, 0x6a, 0x65, 0x61, 0x6b, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x1a, 0x0c, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3e, 0x0a, 0x18, 0x47,
	0x65, 0x74, 0x4c, 0x61, 0x73, 0x74, 0x65, 0x73, 0x74, 0x4c, 0x6f, 0x74, 0x74, 0x65, 0x72, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x73, 0x71, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x73, 0x73, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x6c, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x64, 0x6c, 0x74, 0x22, 0x76, 0x0a, 0x19, 0x47,
	0x65, 0x74, 0x4c, 0x61, 0x73, 0x74, 0x65, 0x73, 0x74, 0x4c, 0x6f, 0x74, 0x74, 0x65, 0x72, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x72, 0x72, 0x43,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x65, 0x72, 0x72, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x65, 0x72, 0x72, 0x4d, 0x73, 0x67, 0x12, 0x27, 0x0a, 0x07, 0x6c, 0x6f,
	0x74, 0x74, 0x65, 0x72, 0x79, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6a, 0x65,
	0x61, 0x6b, 0x2e, 0x4c, 0x6f, 0x74, 0x74, 0x65, 0x72, 0x79, 0x52, 0x07, 0x6c, 0x6f, 0x74, 0x74,
	0x65, 0x72, 0x79, 0x32, 0x6f, 0x0a, 0x07, 0x4c, 0x6f, 0x74, 0x74, 0x65, 0x72, 0x79, 0x12, 0x64,
	0x0a, 0x11, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x73, 0x74, 0x65, 0x73, 0x74, 0x4c, 0x6f, 0x74, 0x74,
	0x65, 0x72, 0x79, 0x12, 0x26, 0x2e, 0x6a, 0x65, 0x61, 0x6b, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x73, 0x74, 0x65, 0x73, 0x74, 0x4c, 0x6f, 0x74,
	0x74, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x6a, 0x65,
	0x61, 0x6b, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x61,
	0x73, 0x74, 0x65, 0x73, 0x74, 0x4c, 0x6f, 0x74, 0x74, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0d, 0x5a, 0x0b, 0x62, 0x69, 0x7a, 0x2f, 0x6c, 0x6f, 0x74, 0x74,
	0x65, 0x72, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_lottery_proto_rawDescOnce sync.Once
	file_lottery_proto_rawDescData = file_lottery_proto_rawDesc
)

func file_lottery_proto_rawDescGZIP() []byte {
	file_lottery_proto_rawDescOnce.Do(func() {
		file_lottery_proto_rawDescData = protoimpl.X.CompressGZIP(file_lottery_proto_rawDescData)
	})
	return file_lottery_proto_rawDescData
}

var file_lottery_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_lottery_proto_goTypes = []interface{}{
	(*GetLastestLotteryRequest)(nil),  // 0: jeak.gateway.GetLastestLotteryRequest
	(*GetLastestLotteryResponse)(nil), // 1: jeak.gateway.GetLastestLotteryResponse
	(*m.Lottery)(nil),                 // 2: jeak.Lottery
}
var file_lottery_proto_depIdxs = []int32{
	2, // 0: jeak.gateway.GetLastestLotteryResponse.lottery:type_name -> jeak.Lottery
	0, // 1: jeak.gateway.Lottery.GetLastestLottery:input_type -> jeak.gateway.GetLastestLotteryRequest
	1, // 2: jeak.gateway.Lottery.GetLastestLottery:output_type -> jeak.gateway.GetLastestLotteryResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_lottery_proto_init() }
func file_lottery_proto_init() {
	if File_lottery_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_lottery_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLastestLotteryRequest); i {
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
		file_lottery_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLastestLotteryResponse); i {
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
			RawDescriptor: file_lottery_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_lottery_proto_goTypes,
		DependencyIndexes: file_lottery_proto_depIdxs,
		MessageInfos:      file_lottery_proto_msgTypes,
	}.Build()
	File_lottery_proto = out.File
	file_lottery_proto_rawDesc = nil
	file_lottery_proto_goTypes = nil
	file_lottery_proto_depIdxs = nil
}
