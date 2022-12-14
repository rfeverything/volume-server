// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: proto/volume.proto

package volume

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

type WriteReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VolumeId uint32 `protobuf:"varint,1,opt,name=volume_id,json=volumeId,proto3" json:"volume_id,omitempty"`
	Fid      uint64 `protobuf:"varint,2,opt,name=fid,proto3" json:"fid,omitempty"`
	Size     int32  `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	Blob     []byte `protobuf:"bytes,4,opt,name=blob,proto3" json:"blob,omitempty"`
}

func (x *WriteReq) Reset() {
	*x = WriteReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_volume_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteReq) ProtoMessage() {}

func (x *WriteReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_volume_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteReq.ProtoReflect.Descriptor instead.
func (*WriteReq) Descriptor() ([]byte, []int) {
	return file_proto_volume_proto_rawDescGZIP(), []int{0}
}

func (x *WriteReq) GetVolumeId() uint32 {
	if x != nil {
		return x.VolumeId
	}
	return 0
}

func (x *WriteReq) GetFid() uint64 {
	if x != nil {
		return x.Fid
	}
	return 0
}

func (x *WriteReq) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *WriteReq) GetBlob() []byte {
	if x != nil {
		return x.Blob
	}
	return nil
}

type WriteResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *WriteResp) Reset() {
	*x = WriteResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_volume_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteResp) ProtoMessage() {}

func (x *WriteResp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_volume_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteResp.ProtoReflect.Descriptor instead.
func (*WriteResp) Descriptor() ([]byte, []int) {
	return file_proto_volume_proto_rawDescGZIP(), []int{1}
}

type StatusReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StatusReq) Reset() {
	*x = StatusReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_volume_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusReq) ProtoMessage() {}

func (x *StatusReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_volume_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusReq.ProtoReflect.Descriptor instead.
func (*StatusReq) Descriptor() ([]byte, []int) {
	return file_proto_volume_proto_rawDescGZIP(), []int{2}
}

type StatusResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total      uint64 `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Free       uint64 `protobuf:"varint,2,opt,name=free,proto3" json:"free,omitempty"`
	Used       uint64 `protobuf:"varint,3,opt,name=used,proto3" json:"used,omitempty"`
	BlockSize  uint64 `protobuf:"varint,4,opt,name=block_size,json=blockSize,proto3" json:"block_size,omitempty"`
	BlockCount uint64 `protobuf:"varint,5,opt,name=block_count,json=blockCount,proto3" json:"block_count,omitempty"`
}

func (x *StatusResp) Reset() {
	*x = StatusResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_volume_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusResp) ProtoMessage() {}

func (x *StatusResp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_volume_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusResp.ProtoReflect.Descriptor instead.
func (*StatusResp) Descriptor() ([]byte, []int) {
	return file_proto_volume_proto_rawDescGZIP(), []int{3}
}

func (x *StatusResp) GetTotal() uint64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *StatusResp) GetFree() uint64 {
	if x != nil {
		return x.Free
	}
	return 0
}

func (x *StatusResp) GetUsed() uint64 {
	if x != nil {
		return x.Used
	}
	return 0
}

func (x *StatusResp) GetBlockSize() uint64 {
	if x != nil {
		return x.BlockSize
	}
	return 0
}

func (x *StatusResp) GetBlockCount() uint64 {
	if x != nil {
		return x.BlockCount
	}
	return 0
}

type ReadReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VolumeId uint32 `protobuf:"varint,1,opt,name=volume_id,json=volumeId,proto3" json:"volume_id,omitempty"`
	Fid      uint64 `protobuf:"varint,2,opt,name=fid,proto3" json:"fid,omitempty"`
	BlockId  uint64 `protobuf:"varint,3,opt,name=block_id,json=blockId,proto3" json:"block_id,omitempty"`
	Offset   uint64 `protobuf:"varint,4,opt,name=offset,proto3" json:"offset,omitempty"`
	Size     int32  `protobuf:"varint,5,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *ReadReq) Reset() {
	*x = ReadReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_volume_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadReq) ProtoMessage() {}

func (x *ReadReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_volume_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadReq.ProtoReflect.Descriptor instead.
func (*ReadReq) Descriptor() ([]byte, []int) {
	return file_proto_volume_proto_rawDescGZIP(), []int{4}
}

func (x *ReadReq) GetVolumeId() uint32 {
	if x != nil {
		return x.VolumeId
	}
	return 0
}

func (x *ReadReq) GetFid() uint64 {
	if x != nil {
		return x.Fid
	}
	return 0
}

func (x *ReadReq) GetBlockId() uint64 {
	if x != nil {
		return x.BlockId
	}
	return 0
}

func (x *ReadReq) GetOffset() uint64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ReadReq) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

type ReadResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReadResp) Reset() {
	*x = ReadResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_volume_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadResp) ProtoMessage() {}

func (x *ReadResp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_volume_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadResp.ProtoReflect.Descriptor instead.
func (*ReadResp) Descriptor() ([]byte, []int) {
	return file_proto_volume_proto_rawDescGZIP(), []int{5}
}

var File_proto_volume_proto protoreflect.FileDescriptor

var file_proto_volume_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x22, 0x61, 0x0a, 0x08,
	0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x1b, 0x0a, 0x09, 0x76, 0x6f, 0x6c, 0x75,
	0x6d, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x76, 0x6f, 0x6c,
	0x75, 0x6d, 0x65, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x03, 0x66, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x62,
	0x6c, 0x6f, 0x62, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x62, 0x6c, 0x6f, 0x62, 0x22,
	0x0b, 0x0a, 0x09, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x0b, 0x0a, 0x09,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x22, 0x8a, 0x01, 0x0a, 0x0a, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x12,
	0x0a, 0x04, 0x66, 0x72, 0x65, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x66, 0x72,
	0x65, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x04, 0x75, 0x73, 0x65, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x7f, 0x0a, 0x07, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65,
	0x71, 0x12, 0x1b, 0x0a, 0x09, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x49, 0x64, 0x12, 0x10,
	0x0a, 0x03, 0x66, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x66, 0x69, 0x64,
	0x12, 0x19, 0x0a, 0x08, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x07, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x0a, 0x0a, 0x08, 0x52, 0x65, 0x61, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x32, 0x9e, 0x01, 0x0a, 0x06, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x2e,
	0x0a, 0x05, 0x57, 0x72, 0x69, 0x74, 0x65, 0x12, 0x10, 0x2e, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65,
	0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x76, 0x6f, 0x6c, 0x75,
	0x6d, 0x65, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x2b,
	0x0a, 0x04, 0x52, 0x65, 0x61, 0x64, 0x12, 0x0f, 0x2e, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x2e,
	0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65,
	0x2e, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x0c, 0x56,
	0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x11, 0x2e, 0x76, 0x6f,
	0x6c, 0x75, 0x6d, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x12,
	0x2e, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_volume_proto_rawDescOnce sync.Once
	file_proto_volume_proto_rawDescData = file_proto_volume_proto_rawDesc
)

func file_proto_volume_proto_rawDescGZIP() []byte {
	file_proto_volume_proto_rawDescOnce.Do(func() {
		file_proto_volume_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_volume_proto_rawDescData)
	})
	return file_proto_volume_proto_rawDescData
}

var file_proto_volume_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_volume_proto_goTypes = []interface{}{
	(*WriteReq)(nil),   // 0: volume.WriteReq
	(*WriteResp)(nil),  // 1: volume.WriteResp
	(*StatusReq)(nil),  // 2: volume.StatusReq
	(*StatusResp)(nil), // 3: volume.StatusResp
	(*ReadReq)(nil),    // 4: volume.ReadReq
	(*ReadResp)(nil),   // 5: volume.ReadResp
}
var file_proto_volume_proto_depIdxs = []int32{
	0, // 0: volume.Volume.Write:input_type -> volume.WriteReq
	4, // 1: volume.Volume.Read:input_type -> volume.ReadReq
	2, // 2: volume.Volume.VolumeStatus:input_type -> volume.StatusReq
	1, // 3: volume.Volume.Write:output_type -> volume.WriteResp
	5, // 4: volume.Volume.Read:output_type -> volume.ReadResp
	3, // 5: volume.Volume.VolumeStatus:output_type -> volume.StatusResp
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_volume_proto_init() }
func file_proto_volume_proto_init() {
	if File_proto_volume_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_volume_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteReq); i {
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
		file_proto_volume_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteResp); i {
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
		file_proto_volume_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusReq); i {
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
		file_proto_volume_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusResp); i {
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
		file_proto_volume_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadReq); i {
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
		file_proto_volume_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadResp); i {
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
			RawDescriptor: file_proto_volume_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_volume_proto_goTypes,
		DependencyIndexes: file_proto_volume_proto_depIdxs,
		MessageInfos:      file_proto_volume_proto_msgTypes,
	}.Build()
	File_proto_volume_proto = out.File
	file_proto_volume_proto_rawDesc = nil
	file_proto_volume_proto_goTypes = nil
	file_proto_volume_proto_depIdxs = nil
}
