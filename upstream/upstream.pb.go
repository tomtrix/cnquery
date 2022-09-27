// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: upstream.proto

package upstream

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

type ServiceAccountCredentials struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mrn         string `protobuf:"bytes,1,opt,name=mrn,proto3" json:"mrn,omitempty"`
	ParentMrn   string `protobuf:"bytes,2,opt,name=parent_mrn,json=parentMrn,proto3" json:"parent_mrn,omitempty"`
	PrivateKey  string `protobuf:"bytes,3,opt,name=private_key,json=privateKey,proto3" json:"private_key,omitempty"`
	Certificate string `protobuf:"bytes,4,opt,name=certificate,proto3" json:"certificate,omitempty"`
	ApiEndpoint string `protobuf:"bytes,5,opt,name=api_endpoint,json=apiEndpoint,proto3" json:"api_endpoint,omitempty"`
}

func (x *ServiceAccountCredentials) Reset() {
	*x = ServiceAccountCredentials{}
	if protoimpl.UnsafeEnabled {
		mi := &file_upstream_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceAccountCredentials) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceAccountCredentials) ProtoMessage() {}

func (x *ServiceAccountCredentials) ProtoReflect() protoreflect.Message {
	mi := &file_upstream_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceAccountCredentials.ProtoReflect.Descriptor instead.
func (*ServiceAccountCredentials) Descriptor() ([]byte, []int) {
	return file_upstream_proto_rawDescGZIP(), []int{0}
}

func (x *ServiceAccountCredentials) GetMrn() string {
	if x != nil {
		return x.Mrn
	}
	return ""
}

func (x *ServiceAccountCredentials) GetParentMrn() string {
	if x != nil {
		return x.ParentMrn
	}
	return ""
}

func (x *ServiceAccountCredentials) GetPrivateKey() string {
	if x != nil {
		return x.PrivateKey
	}
	return ""
}

func (x *ServiceAccountCredentials) GetCertificate() string {
	if x != nil {
		return x.Certificate
	}
	return ""
}

func (x *ServiceAccountCredentials) GetApiEndpoint() string {
	if x != nil {
		return x.ApiEndpoint
	}
	return ""
}

type Ping struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Ping) Reset() {
	*x = Ping{}
	if protoimpl.UnsafeEnabled {
		mi := &file_upstream_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ping) ProtoMessage() {}

func (x *Ping) ProtoReflect() protoreflect.Message {
	mi := &file_upstream_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ping.ProtoReflect.Descriptor instead.
func (*Ping) Descriptor() ([]byte, []int) {
	return file_upstream_proto_rawDescGZIP(), []int{1}
}

type Pong struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Pong) Reset() {
	*x = Pong{}
	if protoimpl.UnsafeEnabled {
		mi := &file_upstream_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pong) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pong) ProtoMessage() {}

func (x *Pong) ProtoReflect() protoreflect.Message {
	mi := &file_upstream_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pong.ProtoReflect.Descriptor instead.
func (*Pong) Descriptor() ([]byte, []int) {
	return file_upstream_proto_rawDescGZIP(), []int{2}
}

type AgentRegistrationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token     string     `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Name      string     `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	AgentInfo *AgentInfo `protobuf:"bytes,3,opt,name=agent_info,json=agentInfo,proto3" json:"agent_info,omitempty"`
}

func (x *AgentRegistrationRequest) Reset() {
	*x = AgentRegistrationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_upstream_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgentRegistrationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentRegistrationRequest) ProtoMessage() {}

func (x *AgentRegistrationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_upstream_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentRegistrationRequest.ProtoReflect.Descriptor instead.
func (*AgentRegistrationRequest) Descriptor() ([]byte, []int) {
	return file_upstream_proto_rawDescGZIP(), []int{3}
}

func (x *AgentRegistrationRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *AgentRegistrationRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AgentRegistrationRequest) GetAgentInfo() *AgentInfo {
	if x != nil {
		return x.AgentInfo
	}
	return nil
}

type AgentInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mrn              string            `protobuf:"bytes,1,opt,name=mrn,proto3" json:"mrn,omitempty"`
	Version          string            `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Build            string            `protobuf:"bytes,3,opt,name=build,proto3" json:"build,omitempty"`
	PlatformName     string            `protobuf:"bytes,4,opt,name=platform_name,json=platformName,proto3" json:"platform_name,omitempty"`
	PlatformRelease  string            `protobuf:"bytes,5,opt,name=platform_release,json=platformRelease,proto3" json:"platform_release,omitempty"`
	PlatformArch     string            `protobuf:"bytes,6,opt,name=platform_arch,json=platformArch,proto3" json:"platform_arch,omitempty"`
	PlatformIp       string            `protobuf:"bytes,7,opt,name=platform_ip,json=platformIp,proto3" json:"platform_ip,omitempty"`
	PlatformHostname string            `protobuf:"bytes,8,opt,name=platform_hostname,json=platformHostname,proto3" json:"platform_hostname,omitempty"`
	Labels           map[string]string `protobuf:"bytes,18,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	PlatformId       string            `protobuf:"bytes,20,opt,name=platform_id,json=platformId,proto3" json:"platform_id,omitempty"`
}

func (x *AgentInfo) Reset() {
	*x = AgentInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_upstream_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgentInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentInfo) ProtoMessage() {}

func (x *AgentInfo) ProtoReflect() protoreflect.Message {
	mi := &file_upstream_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentInfo.ProtoReflect.Descriptor instead.
func (*AgentInfo) Descriptor() ([]byte, []int) {
	return file_upstream_proto_rawDescGZIP(), []int{4}
}

func (x *AgentInfo) GetMrn() string {
	if x != nil {
		return x.Mrn
	}
	return ""
}

func (x *AgentInfo) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *AgentInfo) GetBuild() string {
	if x != nil {
		return x.Build
	}
	return ""
}

func (x *AgentInfo) GetPlatformName() string {
	if x != nil {
		return x.PlatformName
	}
	return ""
}

func (x *AgentInfo) GetPlatformRelease() string {
	if x != nil {
		return x.PlatformRelease
	}
	return ""
}

func (x *AgentInfo) GetPlatformArch() string {
	if x != nil {
		return x.PlatformArch
	}
	return ""
}

func (x *AgentInfo) GetPlatformIp() string {
	if x != nil {
		return x.PlatformIp
	}
	return ""
}

func (x *AgentInfo) GetPlatformHostname() string {
	if x != nil {
		return x.PlatformHostname
	}
	return ""
}

func (x *AgentInfo) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *AgentInfo) GetPlatformId() string {
	if x != nil {
		return x.PlatformId
	}
	return ""
}

type AgentRegistrationConfirmation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AgentMrn   string                     `protobuf:"bytes,1,opt,name=agent_mrn,json=agentMrn,proto3" json:"agent_mrn,omitempty"`
	Credential *ServiceAccountCredentials `protobuf:"bytes,2,opt,name=credential,proto3" json:"credential,omitempty"`
}

func (x *AgentRegistrationConfirmation) Reset() {
	*x = AgentRegistrationConfirmation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_upstream_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgentRegistrationConfirmation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentRegistrationConfirmation) ProtoMessage() {}

func (x *AgentRegistrationConfirmation) ProtoReflect() protoreflect.Message {
	mi := &file_upstream_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentRegistrationConfirmation.ProtoReflect.Descriptor instead.
func (*AgentRegistrationConfirmation) Descriptor() ([]byte, []int) {
	return file_upstream_proto_rawDescGZIP(), []int{5}
}

func (x *AgentRegistrationConfirmation) GetAgentMrn() string {
	if x != nil {
		return x.AgentMrn
	}
	return ""
}

func (x *AgentRegistrationConfirmation) GetCredential() *ServiceAccountCredentials {
	if x != nil {
		return x.Credential
	}
	return nil
}

type Mrn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mrn string `protobuf:"bytes,1,opt,name=mrn,proto3" json:"mrn,omitempty"`
}

func (x *Mrn) Reset() {
	*x = Mrn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_upstream_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Mrn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Mrn) ProtoMessage() {}

func (x *Mrn) ProtoReflect() protoreflect.Message {
	mi := &file_upstream_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Mrn.ProtoReflect.Descriptor instead.
func (*Mrn) Descriptor() ([]byte, []int) {
	return file_upstream_proto_rawDescGZIP(), []int{6}
}

func (x *Mrn) GetMrn() string {
	if x != nil {
		return x.Mrn
	}
	return ""
}

type Confirmation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mrn string `protobuf:"bytes,1,opt,name=mrn,proto3" json:"mrn,omitempty"`
}

func (x *Confirmation) Reset() {
	*x = Confirmation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_upstream_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Confirmation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Confirmation) ProtoMessage() {}

func (x *Confirmation) ProtoReflect() protoreflect.Message {
	mi := &file_upstream_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Confirmation.ProtoReflect.Descriptor instead.
func (*Confirmation) Descriptor() ([]byte, []int) {
	return file_upstream_proto_rawDescGZIP(), []int{7}
}

func (x *Confirmation) GetMrn() string {
	if x != nil {
		return x.Mrn
	}
	return ""
}

var File_upstream_proto protoreflect.FileDescriptor

var file_upstream_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x75, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x12, 0x6d, 0x6f, 0x6e, 0x64, 0x6f, 0x6f, 0x2e, 0x75, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x2e, 0x76, 0x31, 0x22, 0xb2, 0x01, 0x0a, 0x19, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61,
	0x6c, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x72, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6d, 0x72, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x6d,
	0x72, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x4d, 0x72, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x6b,
	0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74,
	0x65, 0x4b, 0x65, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x65, 0x72, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x70, 0x69, 0x5f, 0x65, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x70,
	0x69, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x22, 0x06, 0x0a, 0x04, 0x50, 0x69, 0x6e,
	0x67, 0x22, 0x06, 0x0a, 0x04, 0x50, 0x6f, 0x6e, 0x67, 0x22, 0x82, 0x01, 0x0a, 0x18, 0x41, 0x67,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x3c, 0x0a, 0x0a, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6d, 0x6f, 0x6e, 0x64, 0x6f, 0x6f, 0x2e, 0x75, 0x70,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x09, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0xaf,
	0x03, 0x0a, 0x09, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x10, 0x0a, 0x03,
	0x6d, 0x72, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x72, 0x6e, 0x12, 0x18,
	0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x75, 0x69, 0x6c,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x12, 0x23,
	0x0a, 0x0d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5f,
	0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x12, 0x23,
	0x0a, 0x0d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x61, 0x72, 0x63, 0x68, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x41,
	0x72, 0x63, 0x68, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5f,
	0x69, 0x70, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x49, 0x70, 0x12, 0x2b, 0x0a, 0x11, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x5f, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x10, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x48, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x41, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x12, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x29, 0x2e, 0x6d, 0x6f, 0x6e, 0x64, 0x6f, 0x6f, 0x2e, 0x75, 0x70, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x5f, 0x69, 0x64, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x49, 0x64, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0x8b, 0x01, 0x0a, 0x1d, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x72, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x4d, 0x72, 0x6e, 0x12,
	0x4d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x6d, 0x6f, 0x6e, 0x64, 0x6f, 0x6f, 0x2e, 0x75, 0x70, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61,
	0x6c, 0x73, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x22, 0x17,
	0x0a, 0x03, 0x4d, 0x72, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x72, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6d, 0x72, 0x6e, 0x22, 0x20, 0x0a, 0x0c, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x72, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x72, 0x6e, 0x32, 0x8e, 0x02, 0x0a, 0x0c, 0x41, 0x67,
	0x65, 0x6e, 0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x70, 0x0a, 0x0d, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x2c, 0x2e, 0x6d, 0x6f,
	0x6e, 0x64, 0x6f, 0x6f, 0x2e, 0x75, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x76, 0x31,
	0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x6d, 0x6f, 0x6e, 0x64,
	0x6f, 0x6f, 0x2e, 0x75, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4c, 0x0a, 0x0f,
	0x55, 0x6e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x12,
	0x17, 0x2e, 0x6d, 0x6f, 0x6e, 0x64, 0x6f, 0x6f, 0x2e, 0x75, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x72, 0x6e, 0x1a, 0x20, 0x2e, 0x6d, 0x6f, 0x6e, 0x64, 0x6f,
	0x6f, 0x2e, 0x75, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3e, 0x0a, 0x08, 0x50, 0x69,
	0x6e, 0x67, 0x50, 0x6f, 0x6e, 0x67, 0x12, 0x18, 0x2e, 0x6d, 0x6f, 0x6e, 0x64, 0x6f, 0x6f, 0x2e,
	0x75, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x69, 0x6e, 0x67,
	0x1a, 0x18, 0x2e, 0x6d, 0x6f, 0x6e, 0x64, 0x6f, 0x6f, 0x2e, 0x75, 0x70, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x6e, 0x67, 0x42, 0x20, 0x5a, 0x1e, 0x67, 0x6f,
	0x2e, 0x6d, 0x6f, 0x6e, 0x64, 0x6f, 0x6f, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6e, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x2f, 0x75, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_upstream_proto_rawDescOnce sync.Once
	file_upstream_proto_rawDescData = file_upstream_proto_rawDesc
)

func file_upstream_proto_rawDescGZIP() []byte {
	file_upstream_proto_rawDescOnce.Do(func() {
		file_upstream_proto_rawDescData = protoimpl.X.CompressGZIP(file_upstream_proto_rawDescData)
	})
	return file_upstream_proto_rawDescData
}

var file_upstream_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_upstream_proto_goTypes = []interface{}{
	(*ServiceAccountCredentials)(nil),     // 0: mondoo.upstream.v1.ServiceAccountCredentials
	(*Ping)(nil),                          // 1: mondoo.upstream.v1.Ping
	(*Pong)(nil),                          // 2: mondoo.upstream.v1.Pong
	(*AgentRegistrationRequest)(nil),      // 3: mondoo.upstream.v1.AgentRegistrationRequest
	(*AgentInfo)(nil),                     // 4: mondoo.upstream.v1.AgentInfo
	(*AgentRegistrationConfirmation)(nil), // 5: mondoo.upstream.v1.AgentRegistrationConfirmation
	(*Mrn)(nil),                           // 6: mondoo.upstream.v1.Mrn
	(*Confirmation)(nil),                  // 7: mondoo.upstream.v1.Confirmation
	nil,                                   // 8: mondoo.upstream.v1.AgentInfo.LabelsEntry
}
var file_upstream_proto_depIdxs = []int32{
	4, // 0: mondoo.upstream.v1.AgentRegistrationRequest.agent_info:type_name -> mondoo.upstream.v1.AgentInfo
	8, // 1: mondoo.upstream.v1.AgentInfo.labels:type_name -> mondoo.upstream.v1.AgentInfo.LabelsEntry
	0, // 2: mondoo.upstream.v1.AgentRegistrationConfirmation.credential:type_name -> mondoo.upstream.v1.ServiceAccountCredentials
	3, // 3: mondoo.upstream.v1.AgentManager.RegisterAgent:input_type -> mondoo.upstream.v1.AgentRegistrationRequest
	6, // 4: mondoo.upstream.v1.AgentManager.UnRegisterAgent:input_type -> mondoo.upstream.v1.Mrn
	1, // 5: mondoo.upstream.v1.AgentManager.PingPong:input_type -> mondoo.upstream.v1.Ping
	5, // 6: mondoo.upstream.v1.AgentManager.RegisterAgent:output_type -> mondoo.upstream.v1.AgentRegistrationConfirmation
	7, // 7: mondoo.upstream.v1.AgentManager.UnRegisterAgent:output_type -> mondoo.upstream.v1.Confirmation
	2, // 8: mondoo.upstream.v1.AgentManager.PingPong:output_type -> mondoo.upstream.v1.Pong
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_upstream_proto_init() }
func file_upstream_proto_init() {
	if File_upstream_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_upstream_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceAccountCredentials); i {
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
		file_upstream_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ping); i {
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
		file_upstream_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pong); i {
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
		file_upstream_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AgentRegistrationRequest); i {
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
		file_upstream_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AgentInfo); i {
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
		file_upstream_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AgentRegistrationConfirmation); i {
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
		file_upstream_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Mrn); i {
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
		file_upstream_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Confirmation); i {
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
			RawDescriptor: file_upstream_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_upstream_proto_goTypes,
		DependencyIndexes: file_upstream_proto_depIdxs,
		MessageInfos:      file_upstream_proto_msgTypes,
	}.Build()
	File_upstream_proto = out.File
	file_upstream_proto_rawDesc = nil
	file_upstream_proto_goTypes = nil
	file_upstream_proto_depIdxs = nil
}
