//
//Define messages to create new contract to TKS-Contract service.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: contract.proto

package pbgo

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// CreateContractRequest is a request to create new contract to TKS-Contract service.
type CreateContractRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContractorName    string                 `protobuf:"bytes,1,opt,name=contractor_name,json=contractorName,proto3" json:"contractor_name,omitempty"`          // Name of contractor
	ContractId        string                 `protobuf:"bytes,2,opt,name=contract_id,json=contractId,proto3" json:"contract_id,omitempty"`                      // Global Unique ID created by CBP
	Quota             *ContractQuota         `protobuf:"bytes,3,opt,name=quota,proto3" json:"quota,omitempty"`                                                  // Resource Quota for this contract
	AvailableServices []string               `protobuf:"bytes,4,rep,name=available_services,json=availableServices,proto3" json:"available_services,omitempty"` // Available service list
	LastUpdated       *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
}

func (x *CreateContractRequest) Reset() {
	*x = CreateContractRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateContractRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateContractRequest) ProtoMessage() {}

func (x *CreateContractRequest) ProtoReflect() protoreflect.Message {
	mi := &file_contract_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateContractRequest.ProtoReflect.Descriptor instead.
func (*CreateContractRequest) Descriptor() ([]byte, []int) {
	return file_contract_proto_rawDescGZIP(), []int{0}
}

func (x *CreateContractRequest) GetContractorName() string {
	if x != nil {
		return x.ContractorName
	}
	return ""
}

func (x *CreateContractRequest) GetContractId() string {
	if x != nil {
		return x.ContractId
	}
	return ""
}

func (x *CreateContractRequest) GetQuota() *ContractQuota {
	if x != nil {
		return x.Quota
	}
	return nil
}

func (x *CreateContractRequest) GetAvailableServices() []string {
	if x != nil {
		return x.AvailableServices
	}
	return nil
}

func (x *CreateContractRequest) GetLastUpdated() *timestamppb.Timestamp {
	if x != nil {
		return x.LastUpdated
	}
	return nil
}

// CreateContractResponse is a response to the CreateContractRequest request
// from TKS-Contract service.
type CreateContractResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code   `protobuf:"varint,1,opt,name=code,proto3,enum=pbgo.Code" json:"code,omitempty"`
	Error   *Error `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	McOpsId string `protobuf:"bytes,3,opt,name=mc_ops_id,json=mcOpsId,proto3" json:"mc_ops_id,omitempty"` // Global Unique ID created by contract service
}

func (x *CreateContractResponse) Reset() {
	*x = CreateContractResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateContractResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateContractResponse) ProtoMessage() {}

func (x *CreateContractResponse) ProtoReflect() protoreflect.Message {
	mi := &file_contract_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateContractResponse.ProtoReflect.Descriptor instead.
func (*CreateContractResponse) Descriptor() ([]byte, []int) {
	return file_contract_proto_rawDescGZIP(), []int{1}
}

func (x *CreateContractResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_OK
}

func (x *CreateContractResponse) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *CreateContractResponse) GetMcOpsId() string {
	if x != nil {
		return x.McOpsId
	}
	return ""
}

// ContractQuota is a resource quota for total usage in clusters.
type ContractQuota struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cpu      int64 `protobuf:"zigzag64,1,opt,name=cpu,proto3" json:"cpu,omitempty"`                           // cpu unit(GB) from 10 to 10000
	Memory   int64 `protobuf:"zigzag64,2,opt,name=memory,proto3" json:"memory,omitempty"`                     // memory unit(MB) from 40 to 40000
	Block    int64 `protobuf:"zigzag64,3,opt,name=block,proto3" json:"block,omitempty"`                       // block storage(MB) from 0 to 256000000
	BlockSsd int64 `protobuf:"zigzag64,4,opt,name=block_ssd,json=blockSsd,proto3" json:"block_ssd,omitempty"` // block ssd storage(MB) from 0 to 256000000
	Fs       int64 `protobuf:"zigzag64,5,opt,name=fs,proto3" json:"fs,omitempty"`                             // filesystem (MB) from 0 to 256000000
	FsSsd    int64 `protobuf:"zigzag64,6,opt,name=fs_ssd,json=fsSsd,proto3" json:"fs_ssd,omitempty"`          // SSD filesystem (MB) from 0 to 256000000
}

func (x *ContractQuota) Reset() {
	*x = ContractQuota{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContractQuota) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContractQuota) ProtoMessage() {}

func (x *ContractQuota) ProtoReflect() protoreflect.Message {
	mi := &file_contract_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContractQuota.ProtoReflect.Descriptor instead.
func (*ContractQuota) Descriptor() ([]byte, []int) {
	return file_contract_proto_rawDescGZIP(), []int{2}
}

func (x *ContractQuota) GetCpu() int64 {
	if x != nil {
		return x.Cpu
	}
	return 0
}

func (x *ContractQuota) GetMemory() int64 {
	if x != nil {
		return x.Memory
	}
	return 0
}

func (x *ContractQuota) GetBlock() int64 {
	if x != nil {
		return x.Block
	}
	return 0
}

func (x *ContractQuota) GetBlockSsd() int64 {
	if x != nil {
		return x.BlockSsd
	}
	return 0
}

func (x *ContractQuota) GetFs() int64 {
	if x != nil {
		return x.Fs
	}
	return 0
}

func (x *ContractQuota) GetFsSsd() int64 {
	if x != nil {
		return x.FsSsd
	}
	return 0
}

type UpdateQuotaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContractId string         `protobuf:"bytes,1,opt,name=contract_id,json=contractId,proto3" json:"contract_id,omitempty"` // Global Unique ID created by CBP
	Quota      *ContractQuota `protobuf:"bytes,2,opt,name=quota,proto3" json:"quota,omitempty"`                             // Resource Quota
}

func (x *UpdateQuotaRequest) Reset() {
	*x = UpdateQuotaRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateQuotaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateQuotaRequest) ProtoMessage() {}

func (x *UpdateQuotaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_contract_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateQuotaRequest.ProtoReflect.Descriptor instead.
func (*UpdateQuotaRequest) Descriptor() ([]byte, []int) {
	return file_contract_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateQuotaRequest) GetContractId() string {
	if x != nil {
		return x.ContractId
	}
	return ""
}

func (x *UpdateQuotaRequest) GetQuota() *ContractQuota {
	if x != nil {
		return x.Quota
	}
	return nil
}

type UpdateQuotaResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code         Code           `protobuf:"varint,1,opt,name=code,proto3,enum=pbgo.Code" json:"code,omitempty"`
	Error        *Error         `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	ContractId   string         `protobuf:"bytes,3,opt,name=contract_id,json=contractId,proto3" json:"contract_id,omitempty"`       // Global Unique ID created by CBP
	PrevQuota    *ContractQuota `protobuf:"bytes,4,opt,name=prev_quota,json=prevQuota,proto3" json:"prev_quota,omitempty"`          // Previous Quota
	CurrentQuota *ContractQuota `protobuf:"bytes,5,opt,name=current_quota,json=currentQuota,proto3" json:"current_quota,omitempty"` // Current Quota
}

func (x *UpdateQuotaResponse) Reset() {
	*x = UpdateQuotaResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateQuotaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateQuotaResponse) ProtoMessage() {}

func (x *UpdateQuotaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_contract_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateQuotaResponse.ProtoReflect.Descriptor instead.
func (*UpdateQuotaResponse) Descriptor() ([]byte, []int) {
	return file_contract_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateQuotaResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_OK
}

func (x *UpdateQuotaResponse) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *UpdateQuotaResponse) GetContractId() string {
	if x != nil {
		return x.ContractId
	}
	return ""
}

func (x *UpdateQuotaResponse) GetPrevQuota() *ContractQuota {
	if x != nil {
		return x.PrevQuota
	}
	return nil
}

func (x *UpdateQuotaResponse) GetCurrentQuota() *ContractQuota {
	if x != nil {
		return x.CurrentQuota
	}
	return nil
}

var File_contract_proto protoreflect.FileDescriptor

var file_contract_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x04, 0x70, 0x62, 0x67, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfa, 0x01, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x27, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x6f, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x05, 0x71, 0x75, 0x6f,
	0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x62, 0x67, 0x6f, 0x2e,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x52, 0x05, 0x71,
	0x75, 0x6f, 0x74, 0x61, 0x12, 0x2d, 0x0a, 0x12, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c,
	0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x11, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x12, 0x3d, 0x0a, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x22, 0x77, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x70, 0x62, 0x67,
	0x6f, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x21, 0x0a, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x62,
	0x67, 0x6f, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12,
	0x1a, 0x0a, 0x09, 0x6d, 0x63, 0x5f, 0x6f, 0x70, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x63, 0x4f, 0x70, 0x73, 0x49, 0x64, 0x22, 0x93, 0x01, 0x0a, 0x0d,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x12, 0x10, 0x0a,
	0x03, 0x63, 0x70, 0x75, 0x18, 0x01, 0x20, 0x01, 0x28, 0x12, 0x52, 0x03, 0x63, 0x70, 0x75, 0x12,
	0x16, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x12, 0x52,
	0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x12, 0x52, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x1b, 0x0a,
	0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x73, 0x73, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x12,
	0x52, 0x08, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x73, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x66, 0x73,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x12, 0x52, 0x02, 0x66, 0x73, 0x12, 0x15, 0x0a, 0x06, 0x66, 0x73,
	0x5f, 0x73, 0x73, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x12, 0x52, 0x05, 0x66, 0x73, 0x53, 0x73,
	0x64, 0x22, 0x60, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x51, 0x75, 0x6f, 0x74, 0x61,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x05, 0x71, 0x75, 0x6f, 0x74,
	0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x62, 0x67, 0x6f, 0x2e, 0x43,
	0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x52, 0x05, 0x71, 0x75,
	0x6f, 0x74, 0x61, 0x22, 0xe7, 0x01, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x51, 0x75,
	0x6f, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x70, 0x62, 0x67, 0x6f,
	0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x21, 0x0a, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x62, 0x67,
	0x6f, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x1f,
	0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x49, 0x64, 0x12,
	0x32, 0x0a, 0x0a, 0x70, 0x72, 0x65, 0x76, 0x5f, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x62, 0x67, 0x6f, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x52, 0x09, 0x70, 0x72, 0x65, 0x76, 0x51, 0x75,
	0x6f, 0x74, 0x61, 0x12, 0x38, 0x0a, 0x0d, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x71,
	0x75, 0x6f, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x62, 0x67,
	0x6f, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x52,
	0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x32, 0xa6, 0x01,
	0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x4d, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x12, 0x1b, 0x2e, 0x70, 0x62, 0x67, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1c, 0x2e, 0x70, 0x62, 0x67, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x44, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x12,
	0x18, 0x2e, 0x70, 0x62, 0x67, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x51, 0x75, 0x6f,
	0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x62, 0x67, 0x6f,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x64, 0x65,
	0x76, 0x2f, 0x74, 0x6b, 0x73, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x62, 0x67, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_contract_proto_rawDescOnce sync.Once
	file_contract_proto_rawDescData = file_contract_proto_rawDesc
)

func file_contract_proto_rawDescGZIP() []byte {
	file_contract_proto_rawDescOnce.Do(func() {
		file_contract_proto_rawDescData = protoimpl.X.CompressGZIP(file_contract_proto_rawDescData)
	})
	return file_contract_proto_rawDescData
}

var file_contract_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_contract_proto_goTypes = []interface{}{
	(*CreateContractRequest)(nil),  // 0: pbgo.CreateContractRequest
	(*CreateContractResponse)(nil), // 1: pbgo.CreateContractResponse
	(*ContractQuota)(nil),          // 2: pbgo.ContractQuota
	(*UpdateQuotaRequest)(nil),     // 3: pbgo.UpdateQuotaRequest
	(*UpdateQuotaResponse)(nil),    // 4: pbgo.UpdateQuotaResponse
	(*timestamppb.Timestamp)(nil),  // 5: google.protobuf.Timestamp
	(Code)(0),                      // 6: pbgo.Code
	(*Error)(nil),                  // 7: pbgo.Error
}
var file_contract_proto_depIdxs = []int32{
	2,  // 0: pbgo.CreateContractRequest.quota:type_name -> pbgo.ContractQuota
	5,  // 1: pbgo.CreateContractRequest.last_updated:type_name -> google.protobuf.Timestamp
	6,  // 2: pbgo.CreateContractResponse.code:type_name -> pbgo.Code
	7,  // 3: pbgo.CreateContractResponse.error:type_name -> pbgo.Error
	2,  // 4: pbgo.UpdateQuotaRequest.quota:type_name -> pbgo.ContractQuota
	6,  // 5: pbgo.UpdateQuotaResponse.code:type_name -> pbgo.Code
	7,  // 6: pbgo.UpdateQuotaResponse.error:type_name -> pbgo.Error
	2,  // 7: pbgo.UpdateQuotaResponse.prev_quota:type_name -> pbgo.ContractQuota
	2,  // 8: pbgo.UpdateQuotaResponse.current_quota:type_name -> pbgo.ContractQuota
	0,  // 9: pbgo.ContractService.CreateContract:input_type -> pbgo.CreateContractRequest
	3,  // 10: pbgo.ContractService.UpdateQuota:input_type -> pbgo.UpdateQuotaRequest
	1,  // 11: pbgo.ContractService.CreateContract:output_type -> pbgo.CreateContractResponse
	4,  // 12: pbgo.ContractService.UpdateQuota:output_type -> pbgo.UpdateQuotaResponse
	11, // [11:13] is the sub-list for method output_type
	9,  // [9:11] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_contract_proto_init() }
func file_contract_proto_init() {
	if File_contract_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_contract_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateContractRequest); i {
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
		file_contract_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateContractResponse); i {
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
		file_contract_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContractQuota); i {
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
		file_contract_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateQuotaRequest); i {
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
		file_contract_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateQuotaResponse); i {
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
			RawDescriptor: file_contract_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_contract_proto_goTypes,
		DependencyIndexes: file_contract_proto_depIdxs,
		MessageInfos:      file_contract_proto_msgTypes,
	}.Build()
	File_contract_proto = out.File
	file_contract_proto_rawDesc = nil
	file_contract_proto_goTypes = nil
	file_contract_proto_depIdxs = nil
}