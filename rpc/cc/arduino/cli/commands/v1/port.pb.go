// This file is part of arduino-cli.
//
// Copyright 2021 ARDUINO SA (http://www.arduino.cc/)
//
// This software is released under the GNU General Public License version 3,
// which covers the main part of arduino-cli.
// The terms of this license can be found at:
// https://www.gnu.org/licenses/gpl-3.0.en.html
//
// You can be released from the requirements of the above licenses by purchasing
// a commercial license. Buying such a license is mandatory if you want to
// modify or otherwise use the software for commercial activities involving the
// Arduino software without disclosing the source code of your own applications.
// To purchase a commercial license, send an email to license@arduino.cc.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: cc/arduino/cli/commands/v1/port.proto

package commands

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

// Port represents a board port that may be used to upload or to monitor a board
type Port struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Address of the port (e.g., `/dev/ttyACM0`).
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// The port label to show on the GUI (e.g. "ttyACM0")
	Label string `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
	// Protocol of the port (e.g., `serial`, `network`, ...).
	Protocol string `protobuf:"bytes,3,opt,name=protocol,proto3" json:"protocol,omitempty"`
	// A human friendly description of the protocol (e.g., "Serial Port (USB)").
	ProtocolLabel string `protobuf:"bytes,4,opt,name=protocol_label,json=protocolLabel,proto3" json:"protocol_label,omitempty"`
	// A set of properties of the port
	Properties map[string]string `protobuf:"bytes,5,rep,name=properties,proto3" json:"properties,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Port) Reset() {
	*x = Port{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cc_arduino_cli_commands_v1_port_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Port) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Port) ProtoMessage() {}

func (x *Port) ProtoReflect() protoreflect.Message {
	mi := &file_cc_arduino_cli_commands_v1_port_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Port.ProtoReflect.Descriptor instead.
func (*Port) Descriptor() ([]byte, []int) {
	return file_cc_arduino_cli_commands_v1_port_proto_rawDescGZIP(), []int{0}
}

func (x *Port) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Port) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *Port) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

func (x *Port) GetProtocolLabel() string {
	if x != nil {
		return x.ProtocolLabel
	}
	return ""
}

func (x *Port) GetProperties() map[string]string {
	if x != nil {
		return x.Properties
	}
	return nil
}

var File_cc_arduino_cli_commands_v1_port_proto protoreflect.FileDescriptor

var file_cc_arduino_cli_commands_v1_port_proto_rawDesc = []byte{
	0x0a, 0x25, 0x63, 0x63, 0x2f, 0x61, 0x72, 0x64, 0x75, 0x69, 0x6e, 0x6f, 0x2f, 0x63, 0x6c, 0x69,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f, 0x72,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x63, 0x63, 0x2e, 0x61, 0x72, 0x64, 0x75,
	0x69, 0x6e, 0x6f, 0x2e, 0x63, 0x6c, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73,
	0x2e, 0x76, 0x31, 0x22, 0x8a, 0x02, 0x0a, 0x04, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12,
	0x50, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x63, 0x63, 0x2e, 0x61, 0x72, 0x64, 0x75, 0x69, 0x6e, 0x6f,
	0x2e, 0x63, 0x6c, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x6f, 0x72, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65,
	0x73, 0x1a, 0x3d, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x42, 0x48, 0x5a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61,
	0x72, 0x64, 0x75, 0x69, 0x6e, 0x6f, 0x2f, 0x61, 0x72, 0x64, 0x75, 0x69, 0x6e, 0x6f, 0x2d, 0x63,
	0x6c, 0x69, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x63, 0x2f, 0x61, 0x72, 0x64, 0x75, 0x69, 0x6e,
	0x6f, 0x2f, 0x63, 0x6c, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2f, 0x76,
	0x31, 0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_cc_arduino_cli_commands_v1_port_proto_rawDescOnce sync.Once
	file_cc_arduino_cli_commands_v1_port_proto_rawDescData = file_cc_arduino_cli_commands_v1_port_proto_rawDesc
)

func file_cc_arduino_cli_commands_v1_port_proto_rawDescGZIP() []byte {
	file_cc_arduino_cli_commands_v1_port_proto_rawDescOnce.Do(func() {
		file_cc_arduino_cli_commands_v1_port_proto_rawDescData = protoimpl.X.CompressGZIP(file_cc_arduino_cli_commands_v1_port_proto_rawDescData)
	})
	return file_cc_arduino_cli_commands_v1_port_proto_rawDescData
}

var file_cc_arduino_cli_commands_v1_port_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_cc_arduino_cli_commands_v1_port_proto_goTypes = []interface{}{
	(*Port)(nil), // 0: cc.arduino.cli.commands.v1.Port
	nil,          // 1: cc.arduino.cli.commands.v1.Port.PropertiesEntry
}
var file_cc_arduino_cli_commands_v1_port_proto_depIdxs = []int32{
	1, // 0: cc.arduino.cli.commands.v1.Port.properties:type_name -> cc.arduino.cli.commands.v1.Port.PropertiesEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_cc_arduino_cli_commands_v1_port_proto_init() }
func file_cc_arduino_cli_commands_v1_port_proto_init() {
	if File_cc_arduino_cli_commands_v1_port_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cc_arduino_cli_commands_v1_port_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Port); i {
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
			RawDescriptor: file_cc_arduino_cli_commands_v1_port_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cc_arduino_cli_commands_v1_port_proto_goTypes,
		DependencyIndexes: file_cc_arduino_cli_commands_v1_port_proto_depIdxs,
		MessageInfos:      file_cc_arduino_cli_commands_v1_port_proto_msgTypes,
	}.Build()
	File_cc_arduino_cli_commands_v1_port_proto = out.File
	file_cc_arduino_cli_commands_v1_port_proto_rawDesc = nil
	file_cc_arduino_cli_commands_v1_port_proto_goTypes = nil
	file_cc_arduino_cli_commands_v1_port_proto_depIdxs = nil
}