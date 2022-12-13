// TcaplusService option Ver.1
// Author: Tcaplus
// 2016-8-1
// Modified: 2018-02-12 Add Crypto options

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: tcaplusservice.optionv1.proto

package tcaplusservice

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var file_tcaplusservice_optionv1_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         60000,
		Name:          "tcaplusservice.tcaplus_primary_key",
		Tag:           "bytes,60000,opt,name=tcaplus_primary_key",
		Filename:      "tcaplusservice.optionv1.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: ([]string)(nil),
		Field:         60001,
		Name:          "tcaplusservice.tcaplus_index",
		Tag:           "bytes,60001,rep,name=tcaplus_index",
		Filename:      "tcaplusservice.optionv1.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         60002,
		Name:          "tcaplusservice.tcaplus_field_cipher_suite",
		Tag:           "bytes,60002,opt,name=tcaplus_field_cipher_suite",
		Filename:      "tcaplusservice.optionv1.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         60003,
		Name:          "tcaplusservice.tcaplus_record_cipher_suite",
		Tag:           "bytes,60003,opt,name=tcaplus_record_cipher_suite",
		Filename:      "tcaplusservice.optionv1.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         60004,
		Name:          "tcaplusservice.tcaplus_cipher_md5",
		Tag:           "bytes,60004,opt,name=tcaplus_cipher_md5",
		Filename:      "tcaplusservice.optionv1.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         60005,
		Name:          "tcaplusservice.tcaplus_sharding_key",
		Tag:           "bytes,60005,opt,name=tcaplus_sharding_key",
		Filename:      "tcaplusservice.optionv1.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         60006,
		Name:          "tcaplusservice.tcaplus_customattr",
		Tag:           "bytes,60006,opt,name=tcaplus_customattr",
		Filename:      "tcaplusservice.optionv1.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*uint32)(nil),
		Field:         60000,
		Name:          "tcaplusservice.tcaplus_size",
		Tag:           "varint,60000,opt,name=tcaplus_size",
		Filename:      "tcaplusservice.optionv1.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         60001,
		Name:          "tcaplusservice.tcaplus_desc",
		Tag:           "bytes,60001,opt,name=tcaplus_desc",
		Filename:      "tcaplusservice.optionv1.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         60002,
		Name:          "tcaplusservice.tcaplus_crypto",
		Tag:           "varint,60002,opt,name=tcaplus_crypto",
		Filename:      "tcaplusservice.optionv1.proto",
	},
}

// Extension fields to descriptorpb.MessageOptions.
var (
	// optional string tcaplus_primary_key = 60000;
	E_TcaplusPrimaryKey = &file_tcaplusservice_optionv1_proto_extTypes[0] //Tcaplus Primary Key
	// repeated string tcaplus_index = 60001;
	E_TcaplusIndex = &file_tcaplusservice_optionv1_proto_extTypes[1] //Tcaplus Index
	// optional string tcaplus_field_cipher_suite = 60002;
	E_TcaplusFieldCipherSuite = &file_tcaplusservice_optionv1_proto_extTypes[2] //Tcaplus field crypto suite name
	// optional string tcaplus_record_cipher_suite = 60003;
	E_TcaplusRecordCipherSuite = &file_tcaplusservice_optionv1_proto_extTypes[3] //Tcaplus record crypto suite name
	// optional string tcaplus_cipher_md5 = 60004;
	E_TcaplusCipherMd5 = &file_tcaplusservice_optionv1_proto_extTypes[4] //Tcaplus crypto key md5
	// optional string tcaplus_sharding_key = 60005;
	E_TcaplusShardingKey = &file_tcaplusservice_optionv1_proto_extTypes[5] //Tcaplus sharding key
	// optional string tcaplus_customattr = 60006;
	E_TcaplusCustomattr = &file_tcaplusservice_optionv1_proto_extTypes[6] //Tcaplus custom define attribute
)

// Extension fields to descriptorpb.FieldOptions.
var (
	// optional uint32 tcaplus_size = 60000;
	E_TcaplusSize = &file_tcaplusservice_optionv1_proto_extTypes[7] // Tcaplus field size
	// optional string tcaplus_desc = 60001;
	E_TcaplusDesc = &file_tcaplusservice_optionv1_proto_extTypes[8] // Tcaplus description
	// optional bool tcaplus_crypto = 60002;
	E_TcaplusCrypto = &file_tcaplusservice_optionv1_proto_extTypes[9] // Tcaplus crypto
)

var File_tcaplusservice_optionv1_proto protoreflect.FileDescriptor

var file_tcaplusservice_optionv1_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x74, 0x63, 0x61, 0x70, 0x6c, 0x75, 0x73, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0e, 0x74, 0x63, 0x61, 0x70, 0x6c, 0x75, 0x73, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a,
	0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x3a, 0x51, 0x0a, 0x13, 0x74, 0x63, 0x61, 0x70, 0x6c, 0x75, 0x73, 0x5f, 0x70, 0x72, 0x69,
	0x6d, 0x61, 0x72, 0x79, 0x5f, 0x6b, 0x65, 0x79, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xe0, 0xd4, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x11, 0x74, 0x63, 0x61, 0x70, 0x6c, 0x75, 0x73, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72,
	0x79, 0x4b, 0x65, 0x79, 0x3a, 0x46, 0x0a, 0x0d, 0x74, 0x63, 0x61, 0x70, 0x6c, 0x75, 0x73, 0x5f,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xe1, 0xd4, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c,
	0x74, 0x63, 0x61, 0x70, 0x6c, 0x75, 0x73, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x3a, 0x5e, 0x0a, 0x1a,
	0x74, 0x63, 0x61, 0x70, 0x6c, 0x75, 0x73, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x63, 0x69,
	0x70, 0x68, 0x65, 0x72, 0x5f, 0x73, 0x75, 0x69, 0x74, 0x65, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xe2, 0xd4, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x17, 0x74, 0x63, 0x61, 0x70, 0x6c, 0x75, 0x73, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x43, 0x69, 0x70, 0x68, 0x65, 0x72, 0x53, 0x75, 0x69, 0x74, 0x65, 0x3a, 0x60, 0x0a, 0x1b,
	0x74, 0x63, 0x61, 0x70, 0x6c, 0x75, 0x73, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x5f, 0x63,
	0x69, 0x70, 0x68, 0x65, 0x72, 0x5f, 0x73, 0x75, 0x69, 0x74, 0x65, 0x12, 0x1f, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xe3, 0xd4, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x18, 0x74, 0x63, 0x61, 0x70, 0x6c, 0x75, 0x73, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x43, 0x69, 0x70, 0x68, 0x65, 0x72, 0x53, 0x75, 0x69, 0x74, 0x65, 0x3a, 0x4f,
	0x0a, 0x12, 0x74, 0x63, 0x61, 0x70, 0x6c, 0x75, 0x73, 0x5f, 0x63, 0x69, 0x70, 0x68, 0x65, 0x72,
	0x5f, 0x6d, 0x64, 0x35, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xe4, 0xd4, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x74,
	0x63, 0x61, 0x70, 0x6c, 0x75, 0x73, 0x43, 0x69, 0x70, 0x68, 0x65, 0x72, 0x4d, 0x64, 0x35, 0x3a,
	0x53, 0x0a, 0x14, 0x74, 0x63, 0x61, 0x70, 0x6c, 0x75, 0x73, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x64,
	0x69, 0x6e, 0x67, 0x5f, 0x6b, 0x65, 0x79, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xe5, 0xd4, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x12, 0x74, 0x63, 0x61, 0x70, 0x6c, 0x75, 0x73, 0x53, 0x68, 0x61, 0x72, 0x64, 0x69, 0x6e,
	0x67, 0x4b, 0x65, 0x79, 0x3a, 0x50, 0x0a, 0x12, 0x74, 0x63, 0x61, 0x70, 0x6c, 0x75, 0x73, 0x5f,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x74, 0x72, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xe6, 0xd4, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x11, 0x74, 0x63, 0x61, 0x70, 0x6c, 0x75, 0x73, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x61, 0x74, 0x74, 0x72, 0x3a, 0x42, 0x0a, 0x0c, 0x74, 0x63, 0x61, 0x70, 0x6c, 0x75,
	0x73, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xe0, 0xd4, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x74,
	0x63, 0x61, 0x70, 0x6c, 0x75, 0x73, 0x53, 0x69, 0x7a, 0x65, 0x3a, 0x42, 0x0a, 0x0c, 0x74, 0x63,
	0x61, 0x70, 0x6c, 0x75, 0x73, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xe1, 0xd4, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x74, 0x63, 0x61, 0x70, 0x6c, 0x75, 0x73, 0x44, 0x65, 0x73, 0x63, 0x3a, 0x46,
	0x0a, 0x0e, 0x74, 0x63, 0x61, 0x70, 0x6c, 0x75, 0x73, 0x5f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f,
	0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0xe2, 0xd4, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x74, 0x63, 0x61, 0x70, 0x6c, 0x75, 0x73,
	0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x42, 0x21, 0x5a, 0x1f, 0x2e, 0x2f, 0x74, 0x63, 0x61, 0x70,
	0x6c, 0x75, 0x73, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x3b, 0x74, 0x63, 0x61, 0x70, 0x6c,
	0x75, 0x73, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
}

var file_tcaplusservice_optionv1_proto_goTypes = []interface{}{
	(*descriptorpb.MessageOptions)(nil), // 0: google.protobuf.MessageOptions
	(*descriptorpb.FieldOptions)(nil),   // 1: google.protobuf.FieldOptions
}
var file_tcaplusservice_optionv1_proto_depIdxs = []int32{
	0,  // 0: tcaplusservice.tcaplus_primary_key:extendee -> google.protobuf.MessageOptions
	0,  // 1: tcaplusservice.tcaplus_index:extendee -> google.protobuf.MessageOptions
	0,  // 2: tcaplusservice.tcaplus_field_cipher_suite:extendee -> google.protobuf.MessageOptions
	0,  // 3: tcaplusservice.tcaplus_record_cipher_suite:extendee -> google.protobuf.MessageOptions
	0,  // 4: tcaplusservice.tcaplus_cipher_md5:extendee -> google.protobuf.MessageOptions
	0,  // 5: tcaplusservice.tcaplus_sharding_key:extendee -> google.protobuf.MessageOptions
	0,  // 6: tcaplusservice.tcaplus_customattr:extendee -> google.protobuf.MessageOptions
	1,  // 7: tcaplusservice.tcaplus_size:extendee -> google.protobuf.FieldOptions
	1,  // 8: tcaplusservice.tcaplus_desc:extendee -> google.protobuf.FieldOptions
	1,  // 9: tcaplusservice.tcaplus_crypto:extendee -> google.protobuf.FieldOptions
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	0,  // [0:10] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_tcaplusservice_optionv1_proto_init() }
func file_tcaplusservice_optionv1_proto_init() {
	if File_tcaplusservice_optionv1_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tcaplusservice_optionv1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 10,
			NumServices:   0,
		},
		GoTypes:           file_tcaplusservice_optionv1_proto_goTypes,
		DependencyIndexes: file_tcaplusservice_optionv1_proto_depIdxs,
		ExtensionInfos:    file_tcaplusservice_optionv1_proto_extTypes,
	}.Build()
	File_tcaplusservice_optionv1_proto = out.File
	file_tcaplusservice_optionv1_proto_rawDesc = nil
	file_tcaplusservice_optionv1_proto_goTypes = nil
	file_tcaplusservice_optionv1_proto_depIdxs = nil
}
