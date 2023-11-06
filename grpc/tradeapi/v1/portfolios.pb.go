// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: grpc/tradeapi/v1/portfolios.proto

package tradeapi

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_grpc_tradeapi_v1_portfolios_proto protoreflect.FileDescriptor

var file_grpc_tradeapi_v1_portfolios_proto_rawDesc = []byte{
	0x0a, 0x21, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x74, 0x72, 0x61, 0x64, 0x65, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x10, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x74, 0x72, 0x61, 0x64, 0x65, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a, 0x22, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x72, 0x61,
	0x64, 0x65, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c,
	0x69, 0x6f, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x6b, 0x0a, 0x0a, 0x50, 0x6f, 0x72,
	0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x73, 0x12, 0x5d, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x50, 0x6f,
	0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x12, 0x26, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x74, 0x72, 0x61, 0x64, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50,
	0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x25, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x72, 0x61, 0x64, 0x65, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x42, 0x19, 0xaa, 0x02, 0x16, 0x46, 0x69, 0x6e, 0x61, 0x6d,
	0x2e, 0x54, 0x72, 0x61, 0x64, 0x65, 0x41, 0x70, 0x69, 0x2e, 0x47, 0x72, 0x70, 0x63, 0x2e, 0x56,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_grpc_tradeapi_v1_portfolios_proto_goTypes = []interface{}{
	(*GetPortfolioRequest)(nil), // 0: proto.tradeapi.v1.GetPortfolioRequest
	(*GetPortfolioResult)(nil),  // 1: proto.tradeapi.v1.GetPortfolioResult
}
var file_grpc_tradeapi_v1_portfolios_proto_depIdxs = []int32{
	0, // 0: grpc.tradeapi.v1.Portfolios.GetPortfolio:input_type -> proto.tradeapi.v1.GetPortfolioRequest
	1, // 1: grpc.tradeapi.v1.Portfolios.GetPortfolio:output_type -> proto.tradeapi.v1.GetPortfolioResult
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_grpc_tradeapi_v1_portfolios_proto_init() }
func file_grpc_tradeapi_v1_portfolios_proto_init() {
	if File_grpc_tradeapi_v1_portfolios_proto != nil {
		return
	}
	file_proto_tradeapi_v1_portfolios_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_grpc_tradeapi_v1_portfolios_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_tradeapi_v1_portfolios_proto_goTypes,
		DependencyIndexes: file_grpc_tradeapi_v1_portfolios_proto_depIdxs,
	}.Build()
	File_grpc_tradeapi_v1_portfolios_proto = out.File
	file_grpc_tradeapi_v1_portfolios_proto_rawDesc = nil
	file_grpc_tradeapi_v1_portfolios_proto_goTypes = nil
	file_grpc_tradeapi_v1_portfolios_proto_depIdxs = nil
}
