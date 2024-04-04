// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.0
// source: api/mesh/v1alpha1/zoneegressoverview.proto

package v1alpha1

import (
	reflect "reflect"
	sync "sync"
)

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"

	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

import (
	_ "github.com/apache/dubbo-kubernetes/api/mesh"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// ZoneEgressOverview defines the projected state of a ZoneEgress.
type ZoneEgressOverview struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ZoneEgress        *ZoneEgress        `protobuf:"bytes,1,opt,name=zoneEgress,proto3" json:"zoneEgress,omitempty"`
	ZoneEgressInsight *ZoneEgressInsight `protobuf:"bytes,2,opt,name=zoneEgressInsight,proto3" json:"zoneEgressInsight,omitempty"`
}

func (x *ZoneEgressOverview) Reset() {
	*x = ZoneEgressOverview{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_mesh_v1alpha1_zoneegressoverview_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ZoneEgressOverview) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ZoneEgressOverview) ProtoMessage() {}

func (x *ZoneEgressOverview) ProtoReflect() protoreflect.Message {
	mi := &file_api_mesh_v1alpha1_zoneegressoverview_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ZoneEgressOverview.ProtoReflect.Descriptor instead.
func (*ZoneEgressOverview) Descriptor() ([]byte, []int) {
	return file_api_mesh_v1alpha1_zoneegressoverview_proto_rawDescGZIP(), []int{0}
}

func (x *ZoneEgressOverview) GetZoneEgress() *ZoneEgress {
	if x != nil {
		return x.ZoneEgress
	}
	return nil
}

func (x *ZoneEgressOverview) GetZoneEgressInsight() *ZoneEgressInsight {
	if x != nil {
		return x.ZoneEgressInsight
	}
	return nil
}

var File_api_mesh_v1alpha1_zoneegressoverview_proto protoreflect.FileDescriptor

var file_api_mesh_v1alpha1_zoneegressoverview_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2f, 0x7a, 0x6f, 0x6e, 0x65, 0x65, 0x67, 0x72, 0x65, 0x73, 0x73, 0x6f, 0x76,
	0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x64, 0x75,
	0x62, 0x62, 0x6f, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x1a, 0x16, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x61, 0x70, 0x69, 0x2f, 0x6d,
	0x65, 0x73, 0x68, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x7a, 0x6f, 0x6e,
	0x65, 0x65, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x61,
	0x70, 0x69, 0x2f, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2f, 0x7a, 0x6f, 0x6e, 0x65, 0x65, 0x67, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x73, 0x69, 0x67,
	0x68, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8d, 0x02, 0x0a, 0x12, 0x5a, 0x6f, 0x6e,
	0x65, 0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x4f, 0x76, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x12,
	0x3f, 0x0a, 0x0a, 0x7a, 0x6f, 0x6e, 0x65, 0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x64, 0x75, 0x62, 0x62, 0x6f, 0x2e, 0x6d, 0x65, 0x73, 0x68,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x5a, 0x6f, 0x6e, 0x65, 0x45, 0x67,
	0x72, 0x65, 0x73, 0x73, 0x52, 0x0a, 0x7a, 0x6f, 0x6e, 0x65, 0x45, 0x67, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x54, 0x0a, 0x11, 0x7a, 0x6f, 0x6e, 0x65, 0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x49, 0x6e,
	0x73, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x64, 0x75,
	0x62, 0x62, 0x6f, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x5a, 0x6f, 0x6e, 0x65, 0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x73, 0x69,
	0x67, 0x68, 0x74, 0x52, 0x11, 0x7a, 0x6f, 0x6e, 0x65, 0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x49,
	0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x3a, 0x60, 0xaa, 0x8c, 0x89, 0xa6, 0x01, 0x1c, 0x0a, 0x1a,
	0x5a, 0x6f, 0x6e, 0x65, 0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x4f, 0x76, 0x65, 0x72, 0x76, 0x69,
	0x65, 0x77, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0xaa, 0x8c, 0x89, 0xa6, 0x01, 0x14,
	0x12, 0x12, 0x5a, 0x6f, 0x6e, 0x65, 0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x4f, 0x76, 0x65, 0x72,
	0x76, 0x69, 0x65, 0x77, 0xaa, 0x8c, 0x89, 0xa6, 0x01, 0x02, 0x18, 0x01, 0xaa, 0x8c, 0x89, 0xa6,
	0x01, 0x06, 0x22, 0x04, 0x6d, 0x65, 0x73, 0x68, 0xaa, 0x8c, 0x89, 0xa6, 0x01, 0x02, 0x30, 0x01,
	0xaa, 0x8c, 0x89, 0xa6, 0x01, 0x02, 0x60, 0x01, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2f, 0x64, 0x75,
	0x62, 0x62, 0x6f, 0x2d, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_mesh_v1alpha1_zoneegressoverview_proto_rawDescOnce sync.Once
	file_api_mesh_v1alpha1_zoneegressoverview_proto_rawDescData = file_api_mesh_v1alpha1_zoneegressoverview_proto_rawDesc
)

func file_api_mesh_v1alpha1_zoneegressoverview_proto_rawDescGZIP() []byte {
	file_api_mesh_v1alpha1_zoneegressoverview_proto_rawDescOnce.Do(func() {
		file_api_mesh_v1alpha1_zoneegressoverview_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_mesh_v1alpha1_zoneegressoverview_proto_rawDescData)
	})
	return file_api_mesh_v1alpha1_zoneegressoverview_proto_rawDescData
}

var file_api_mesh_v1alpha1_zoneegressoverview_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_api_mesh_v1alpha1_zoneegressoverview_proto_goTypes = []interface{}{
	(*ZoneEgressOverview)(nil), // 0: dubbo.mesh.v1alpha1.ZoneEgressOverview
	(*ZoneEgress)(nil),         // 1: dubbo.mesh.v1alpha1.ZoneEgress
	(*ZoneEgressInsight)(nil),  // 2: dubbo.mesh.v1alpha1.ZoneEgressInsight
}
var file_api_mesh_v1alpha1_zoneegressoverview_proto_depIdxs = []int32{
	1, // 0: dubbo.mesh.v1alpha1.ZoneEgressOverview.zoneEgress:type_name -> dubbo.mesh.v1alpha1.ZoneEgress
	2, // 1: dubbo.mesh.v1alpha1.ZoneEgressOverview.zoneEgressInsight:type_name -> dubbo.mesh.v1alpha1.ZoneEgressInsight
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_mesh_v1alpha1_zoneegressoverview_proto_init() }
func file_api_mesh_v1alpha1_zoneegressoverview_proto_init() {
	if File_api_mesh_v1alpha1_zoneegressoverview_proto != nil {
		return
	}
	file_api_mesh_v1alpha1_zoneegress_proto_init()
	file_api_mesh_v1alpha1_zoneegressinsight_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_api_mesh_v1alpha1_zoneegressoverview_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ZoneEgressOverview); i {
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
			RawDescriptor: file_api_mesh_v1alpha1_zoneegressoverview_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_mesh_v1alpha1_zoneegressoverview_proto_goTypes,
		DependencyIndexes: file_api_mesh_v1alpha1_zoneegressoverview_proto_depIdxs,
		MessageInfos:      file_api_mesh_v1alpha1_zoneegressoverview_proto_msgTypes,
	}.Build()
	File_api_mesh_v1alpha1_zoneegressoverview_proto = out.File
	file_api_mesh_v1alpha1_zoneegressoverview_proto_rawDesc = nil
	file_api_mesh_v1alpha1_zoneegressoverview_proto_goTypes = nil
	file_api_mesh_v1alpha1_zoneegressoverview_proto_depIdxs = nil
}
