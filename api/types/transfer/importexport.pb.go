//
//Copyright The containerd Authors.
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.1
// source: github.com/containerd/containerd/api/types/transfer/importexport.proto

package transfer

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

type ImageImportStream struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Stream is used to identify the binary input stream for the import operation.
	// The stream uses the transfer binary stream protocol with the client as the sender.
	// The binary data is expected to be a raw tar stream.
	Stream        string `protobuf:"bytes,1,opt,name=stream,proto3" json:"stream,omitempty"`
	MediaType     string `protobuf:"bytes,2,opt,name=media_type,json=mediaType,proto3" json:"media_type,omitempty"`
	ForceCompress bool   `protobuf:"varint,3,opt,name=force_compress,json=forceCompress,proto3" json:"force_compress,omitempty"`
}

func (x *ImageImportStream) Reset() {
	*x = ImageImportStream{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_containerd_containerd_api_types_transfer_importexport_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImageImportStream) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageImportStream) ProtoMessage() {}

func (x *ImageImportStream) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_containerd_containerd_api_types_transfer_importexport_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageImportStream.ProtoReflect.Descriptor instead.
func (*ImageImportStream) Descriptor() ([]byte, []int) {
	return file_github_com_containerd_containerd_api_types_transfer_importexport_proto_rawDescGZIP(), []int{0}
}

func (x *ImageImportStream) GetStream() string {
	if x != nil {
		return x.Stream
	}
	return ""
}

func (x *ImageImportStream) GetMediaType() string {
	if x != nil {
		return x.MediaType
	}
	return ""
}

func (x *ImageImportStream) GetForceCompress() bool {
	if x != nil {
		return x.ForceCompress
	}
	return false
}

type ImageExportStream struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Stream is used to identify the binary output stream for the export operation.
	// The stream uses the transfer binary stream protocol with the server as the sender.
	// The binary data is expected to be a raw tar stream.
	Stream string `protobuf:"bytes,1,opt,name=stream,proto3" json:"stream,omitempty"`
}

func (x *ImageExportStream) Reset() {
	*x = ImageExportStream{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_containerd_containerd_api_types_transfer_importexport_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImageExportStream) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageExportStream) ProtoMessage() {}

func (x *ImageExportStream) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_containerd_containerd_api_types_transfer_importexport_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageExportStream.ProtoReflect.Descriptor instead.
func (*ImageExportStream) Descriptor() ([]byte, []int) {
	return file_github_com_containerd_containerd_api_types_transfer_importexport_proto_rawDescGZIP(), []int{1}
}

func (x *ImageExportStream) GetStream() string {
	if x != nil {
		return x.Stream
	}
	return ""
}

var File_github_com_containerd_containerd_api_types_transfer_importexport_proto protoreflect.FileDescriptor

var file_github_com_containerd_containerd_api_types_transfer_importexport_proto_rawDesc = []byte{
	0x0a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x64, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x2f, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x78, 0x70, 0x6f,
	0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69,
	0x6e, 0x65, 0x72, 0x64, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x65, 0x72, 0x22, 0x71, 0x0a, 0x11, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x6d, 0x70, 0x6f,
	0x72, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x25, 0x0a, 0x0e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x43, 0x6f,
	0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x22, 0x2b, 0x0a, 0x11, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x45,
	0x78, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x64, 0x2f, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_github_com_containerd_containerd_api_types_transfer_importexport_proto_rawDescOnce sync.Once
	file_github_com_containerd_containerd_api_types_transfer_importexport_proto_rawDescData = file_github_com_containerd_containerd_api_types_transfer_importexport_proto_rawDesc
)

func file_github_com_containerd_containerd_api_types_transfer_importexport_proto_rawDescGZIP() []byte {
	file_github_com_containerd_containerd_api_types_transfer_importexport_proto_rawDescOnce.Do(func() {
		file_github_com_containerd_containerd_api_types_transfer_importexport_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_containerd_containerd_api_types_transfer_importexport_proto_rawDescData)
	})
	return file_github_com_containerd_containerd_api_types_transfer_importexport_proto_rawDescData
}

var file_github_com_containerd_containerd_api_types_transfer_importexport_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_github_com_containerd_containerd_api_types_transfer_importexport_proto_goTypes = []interface{}{
	(*ImageImportStream)(nil), // 0: containerd.types.transfer.ImageImportStream
	(*ImageExportStream)(nil), // 1: containerd.types.transfer.ImageExportStream
}
var file_github_com_containerd_containerd_api_types_transfer_importexport_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_github_com_containerd_containerd_api_types_transfer_importexport_proto_init() }
func file_github_com_containerd_containerd_api_types_transfer_importexport_proto_init() {
	if File_github_com_containerd_containerd_api_types_transfer_importexport_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_containerd_containerd_api_types_transfer_importexport_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImageImportStream); i {
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
		file_github_com_containerd_containerd_api_types_transfer_importexport_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImageExportStream); i {
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
			RawDescriptor: file_github_com_containerd_containerd_api_types_transfer_importexport_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_containerd_containerd_api_types_transfer_importexport_proto_goTypes,
		DependencyIndexes: file_github_com_containerd_containerd_api_types_transfer_importexport_proto_depIdxs,
		MessageInfos:      file_github_com_containerd_containerd_api_types_transfer_importexport_proto_msgTypes,
	}.Build()
	File_github_com_containerd_containerd_api_types_transfer_importexport_proto = out.File
	file_github_com_containerd_containerd_api_types_transfer_importexport_proto_rawDesc = nil
	file_github_com_containerd_containerd_api_types_transfer_importexport_proto_goTypes = nil
	file_github_com_containerd_containerd_api_types_transfer_importexport_proto_depIdxs = nil
}
