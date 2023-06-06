// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.2
// source: covid.proto

package pb

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

type CovidData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Year  int32 `protobuf:"varint,1,opt,name=year,proto3" json:"year,omitempty"`
	Cases int32 `protobuf:"varint,2,opt,name=cases,proto3" json:"cases,omitempty"`
}

func (x *CovidData) Reset() {
	*x = CovidData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_covid_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CovidData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CovidData) ProtoMessage() {}

func (x *CovidData) ProtoReflect() protoreflect.Message {
	mi := &file_covid_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CovidData.ProtoReflect.Descriptor instead.
func (*CovidData) Descriptor() ([]byte, []int) {
	return file_covid_proto_rawDescGZIP(), []int{0}
}

func (x *CovidData) GetYear() int32 {
	if x != nil {
		return x.Year
	}
	return 0
}

func (x *CovidData) GetCases() int32 {
	if x != nil {
		return x.Cases
	}
	return 0
}

type CovidDataList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*CovidData `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *CovidDataList) Reset() {
	*x = CovidDataList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_covid_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CovidDataList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CovidDataList) ProtoMessage() {}

func (x *CovidDataList) ProtoReflect() protoreflect.Message {
	mi := &file_covid_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CovidDataList.ProtoReflect.Descriptor instead.
func (*CovidDataList) Descriptor() ([]byte, []int) {
	return file_covid_proto_rawDescGZIP(), []int{1}
}

func (x *CovidDataList) GetData() []*CovidData {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_covid_proto protoreflect.FileDescriptor

var file_covid_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x63, 0x6f, 0x76, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x63,
	0x6f, 0x76, 0x69, 0x64, 0x22, 0x35, 0x0a, 0x09, 0x43, 0x6f, 0x76, 0x69, 0x64, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x79, 0x65, 0x61, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x61, 0x73, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x61, 0x73, 0x65, 0x73, 0x22, 0x35, 0x0a, 0x0d, 0x43,
	0x6f, 0x76, 0x69, 0x64, 0x44, 0x61, 0x74, 0x61, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x6f, 0x76,
	0x69, 0x64, 0x2e, 0x43, 0x6f, 0x76, 0x69, 0x64, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_covid_proto_rawDescOnce sync.Once
	file_covid_proto_rawDescData = file_covid_proto_rawDesc
)

func file_covid_proto_rawDescGZIP() []byte {
	file_covid_proto_rawDescOnce.Do(func() {
		file_covid_proto_rawDescData = protoimpl.X.CompressGZIP(file_covid_proto_rawDescData)
	})
	return file_covid_proto_rawDescData
}

var file_covid_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_covid_proto_goTypes = []interface{}{
	(*CovidData)(nil),     // 0: covid.CovidData
	(*CovidDataList)(nil), // 1: covid.CovidDataList
}
var file_covid_proto_depIdxs = []int32{
	0, // 0: covid.CovidDataList.data:type_name -> covid.CovidData
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_covid_proto_init() }
func file_covid_proto_init() {
	if File_covid_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_covid_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CovidData); i {
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
		file_covid_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CovidDataList); i {
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
			RawDescriptor: file_covid_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_covid_proto_goTypes,
		DependencyIndexes: file_covid_proto_depIdxs,
		MessageInfos:      file_covid_proto_msgTypes,
	}.Build()
	File_covid_proto = out.File
	file_covid_proto_rawDesc = nil
	file_covid_proto_goTypes = nil
	file_covid_proto_depIdxs = nil
}
