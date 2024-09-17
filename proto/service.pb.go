// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: proto/service.proto

package snapwall

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

type ServiceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time        string `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	Type        string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Source      string `protobuf:"bytes,3,opt,name=source,proto3" json:"source,omitempty"`
	Destination string `protobuf:"bytes,4,opt,name=destination,proto3" json:"destination,omitempty"`
	Port        string `protobuf:"bytes,5,opt,name=port,proto3" json:"port,omitempty"`
	Protocol    string `protobuf:"bytes,6,opt,name=protocol,proto3" json:"protocol,omitempty"`
	Severity    string `protobuf:"bytes,7,opt,name=severity,proto3" json:"severity,omitempty"`
	Pid         int32  `protobuf:"varint,8,opt,name=pid,proto3" json:"pid,omitempty"`
	Path        string `protobuf:"bytes,9,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *ServiceRequest) Reset() {
	*x = ServiceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceRequest) ProtoMessage() {}

func (x *ServiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceRequest.ProtoReflect.Descriptor instead.
func (*ServiceRequest) Descriptor() ([]byte, []int) {
	return file_proto_service_proto_rawDescGZIP(), []int{0}
}

func (x *ServiceRequest) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

func (x *ServiceRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ServiceRequest) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *ServiceRequest) GetDestination() string {
	if x != nil {
		return x.Destination
	}
	return ""
}

func (x *ServiceRequest) GetPort() string {
	if x != nil {
		return x.Port
	}
	return ""
}

func (x *ServiceRequest) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

func (x *ServiceRequest) GetSeverity() string {
	if x != nil {
		return x.Severity
	}
	return ""
}

func (x *ServiceRequest) GetPid() int32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *ServiceRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type ServiceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time        string `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	Type        string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Source      string `protobuf:"bytes,3,opt,name=source,proto3" json:"source,omitempty"`
	Destination string `protobuf:"bytes,4,opt,name=destination,proto3" json:"destination,omitempty"`
	Port        string `protobuf:"bytes,5,opt,name=port,proto3" json:"port,omitempty"`
	Protocol    string `protobuf:"bytes,6,opt,name=protocol,proto3" json:"protocol,omitempty"`
	Severity    string `protobuf:"bytes,7,opt,name=severity,proto3" json:"severity,omitempty"`
	Pid         int32  `protobuf:"varint,8,opt,name=pid,proto3" json:"pid,omitempty"`
	Path        string `protobuf:"bytes,9,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *ServiceResponse) Reset() {
	*x = ServiceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceResponse) ProtoMessage() {}

func (x *ServiceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceResponse.ProtoReflect.Descriptor instead.
func (*ServiceResponse) Descriptor() ([]byte, []int) {
	return file_proto_service_proto_rawDescGZIP(), []int{1}
}

func (x *ServiceResponse) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

func (x *ServiceResponse) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ServiceResponse) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *ServiceResponse) GetDestination() string {
	if x != nil {
		return x.Destination
	}
	return ""
}

func (x *ServiceResponse) GetPort() string {
	if x != nil {
		return x.Port
	}
	return ""
}

func (x *ServiceResponse) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

func (x *ServiceResponse) GetSeverity() string {
	if x != nil {
		return x.Severity
	}
	return ""
}

func (x *ServiceResponse) GetPid() int32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *ServiceResponse) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type NodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cpu     string `protobuf:"bytes,1,opt,name=cpu,proto3" json:"cpu,omitempty"`
	Memory  string `protobuf:"bytes,2,opt,name=memory,proto3" json:"memory,omitempty"`
	Disk    string `protobuf:"bytes,3,opt,name=disk,proto3" json:"disk,omitempty"`
	Network string `protobuf:"bytes,4,opt,name=network,proto3" json:"network,omitempty"`
}

func (x *NodeRequest) Reset() {
	*x = NodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeRequest) ProtoMessage() {}

func (x *NodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeRequest.ProtoReflect.Descriptor instead.
func (*NodeRequest) Descriptor() ([]byte, []int) {
	return file_proto_service_proto_rawDescGZIP(), []int{2}
}

func (x *NodeRequest) GetCpu() string {
	if x != nil {
		return x.Cpu
	}
	return ""
}

func (x *NodeRequest) GetMemory() string {
	if x != nil {
		return x.Memory
	}
	return ""
}

func (x *NodeRequest) GetDisk() string {
	if x != nil {
		return x.Disk
	}
	return ""
}

func (x *NodeRequest) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

type NodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Error  string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *NodeResponse) Reset() {
	*x = NodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeResponse) ProtoMessage() {}

func (x *NodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeResponse.ProtoReflect.Descriptor instead.
func (*NodeResponse) Descriptor() ([]byte, []int) {
	return file_proto_service_proto_rawDescGZIP(), []int{3}
}

func (x *NodeResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *NodeResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_proto_service_proto protoreflect.FileDescriptor

var file_proto_service_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0xe4,
	0x01, 0x0a, 0x0e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x70, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0xe5, 0x01, 0x0a, 0x0f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x6f, 0x72, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x73,
	0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73,
	0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x70, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0x65, 0x0a,
	0x0b, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x63, 0x70, 0x75, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x70, 0x75, 0x12, 0x16,
	0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x69, 0x73, 0x6b, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x69, 0x73, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x22, 0x3c, 0x0a, 0x0c, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x32, 0x88, 0x01, 0x0a, 0x06, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x3d, 0x0a,
	0x04, 0x53, 0x65, 0x6e, 0x64, 0x12, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x12, 0x3f, 0x0a, 0x0c,
	0x53, 0x65, 0x6e, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x14, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4e, 0x6f, 0x64,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x42, 0x20, 0x5a,
	0x1e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x6e, 0x73,
	0x68, 0x61, 0x6c, 0x31, 0x30, 0x31, 0x2f, 0x73, 0x6e, 0x61, 0x70, 0x77, 0x61, 0x6c, 0x6c, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_service_proto_rawDescOnce sync.Once
	file_proto_service_proto_rawDescData = file_proto_service_proto_rawDesc
)

func file_proto_service_proto_rawDescGZIP() []byte {
	file_proto_service_proto_rawDescOnce.Do(func() {
		file_proto_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_service_proto_rawDescData)
	})
	return file_proto_service_proto_rawDescData
}

var file_proto_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_service_proto_goTypes = []any{
	(*ServiceRequest)(nil),  // 0: service.ServiceRequest
	(*ServiceResponse)(nil), // 1: service.ServiceResponse
	(*NodeRequest)(nil),     // 2: service.NodeRequest
	(*NodeResponse)(nil),    // 3: service.NodeResponse
}
var file_proto_service_proto_depIdxs = []int32{
	0, // 0: service.Sender.Send:input_type -> service.ServiceRequest
	2, // 1: service.Sender.SendNodeData:input_type -> service.NodeRequest
	1, // 2: service.Sender.Send:output_type -> service.ServiceResponse
	3, // 3: service.Sender.SendNodeData:output_type -> service.NodeResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_service_proto_init() }
func file_proto_service_proto_init() {
	if File_proto_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ServiceRequest); i {
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
		file_proto_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ServiceResponse); i {
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
		file_proto_service_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*NodeRequest); i {
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
		file_proto_service_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*NodeResponse); i {
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
			RawDescriptor: file_proto_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_service_proto_goTypes,
		DependencyIndexes: file_proto_service_proto_depIdxs,
		MessageInfos:      file_proto_service_proto_msgTypes,
	}.Build()
	File_proto_service_proto = out.File
	file_proto_service_proto_rawDesc = nil
	file_proto_service_proto_goTypes = nil
	file_proto_service_proto_depIdxs = nil
}
