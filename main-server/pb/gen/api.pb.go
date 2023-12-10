// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.6
// source: api.proto

package pb

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type GlobalRatingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Skip  uint32 `protobuf:"varint,1,opt,name=skip,proto3" json:"skip,omitempty"`
	Limit uint32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *GlobalRatingRequest) Reset() {
	*x = GlobalRatingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GlobalRatingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GlobalRatingRequest) ProtoMessage() {}

func (x *GlobalRatingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GlobalRatingRequest.ProtoReflect.Descriptor instead.
func (*GlobalRatingRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{0}
}

func (x *GlobalRatingRequest) GetSkip() uint32 {
	if x != nil {
		return x.Skip
	}
	return 0
}

func (x *GlobalRatingRequest) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GlobalRating struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Datetime string   `protobuf:"bytes,1,opt,name=datetime,proto3" json:"datetime,omitempty"`
	Stocks   []*Stock `protobuf:"bytes,2,rep,name=stocks,proto3" json:"stocks,omitempty"`
}

func (x *GlobalRating) Reset() {
	*x = GlobalRating{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GlobalRating) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GlobalRating) ProtoMessage() {}

func (x *GlobalRating) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GlobalRating.ProtoReflect.Descriptor instead.
func (*GlobalRating) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{1}
}

func (x *GlobalRating) GetDatetime() string {
	if x != nil {
		return x.Datetime
	}
	return ""
}

func (x *GlobalRating) GetStocks() []*Stock {
	if x != nil {
		return x.Stocks
	}
	return nil
}

type StockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SecId string `protobuf:"bytes,1,opt,name=sec_id,json=secId,proto3" json:"sec_id,omitempty"`
}

func (x *StockRequest) Reset() {
	*x = StockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StockRequest) ProtoMessage() {}

func (x *StockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StockRequest.ProtoReflect.Descriptor instead.
func (*StockRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{2}
}

func (x *StockRequest) GetSecId() string {
	if x != nil {
		return x.SecId
	}
	return ""
}

type Stock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Industry   uint32  `protobuf:"varint,1,opt,name=industry,proto3" json:"industry,omitempty"`
	SecId      string  `protobuf:"bytes,2,opt,name=sec_id,json=secId,proto3" json:"sec_id,omitempty"`
	StockPrice float32 `protobuf:"fixed32,3,opt,name=stock_price,json=stockPrice,proto3" json:"stock_price,omitempty"`
	Delta      float32 `protobuf:"fixed32,4,opt,name=delta,proto3" json:"delta,omitempty"`
}

func (x *Stock) Reset() {
	*x = Stock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Stock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stock) ProtoMessage() {}

func (x *Stock) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stock.ProtoReflect.Descriptor instead.
func (*Stock) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{3}
}

func (x *Stock) GetIndustry() uint32 {
	if x != nil {
		return x.Industry
	}
	return 0
}

func (x *Stock) GetSecId() string {
	if x != nil {
		return x.SecId
	}
	return ""
}

func (x *Stock) GetStockPrice() float32 {
	if x != nil {
		return x.StockPrice
	}
	return 0
}

func (x *Stock) GetDelta() float32 {
	if x != nil {
		return x.Delta
	}
	return 0
}

type IndustriesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *IndustriesRequest) Reset() {
	*x = IndustriesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IndustriesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IndustriesRequest) ProtoMessage() {}

func (x *IndustriesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IndustriesRequest.ProtoReflect.Descriptor instead.
func (*IndustriesRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{4}
}

type Industries struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Industries []*Industries_Industry `protobuf:"bytes,1,rep,name=industries,proto3" json:"industries,omitempty"`
}

func (x *Industries) Reset() {
	*x = Industries{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Industries) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Industries) ProtoMessage() {}

func (x *Industries) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Industries.ProtoReflect.Descriptor instead.
func (*Industries) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{5}
}

func (x *Industries) GetIndustries() []*Industries_Industry {
	if x != nil {
		return x.Industries
	}
	return nil
}

type Industries_Industry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Industries_Industry) Reset() {
	*x = Industries_Industry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Industries_Industry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Industries_Industry) ProtoMessage() {}

func (x *Industries_Industry) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Industries_Industry.ProtoReflect.Descriptor instead.
func (*Industries_Industry) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{5, 0}
}

func (x *Industries_Industry) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Industries_Industry) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_api_proto protoreflect.FileDescriptor

var file_api_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x74, 0x72, 0x61,
	0x64, 0x65, 0x72, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x3f, 0x0a, 0x13, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x52, 0x61, 0x74, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6b, 0x69, 0x70,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x73, 0x6b, 0x69, 0x70, 0x12, 0x14, 0x0a, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x22, 0x51, 0x0a, 0x0c, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x52, 0x61, 0x74, 0x69,
	0x6e, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x25,
	0x0a, 0x06, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x74, 0x72, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x06, 0x73,
	0x74, 0x6f, 0x63, 0x6b, 0x73, 0x22, 0x25, 0x0a, 0x0c, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x65, 0x63, 0x49, 0x64, 0x22, 0x71, 0x0a, 0x05,
	0x53, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x69, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72,
	0x79, 0x12, 0x15, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x73, 0x65, 0x63, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x6f, 0x63,
	0x6b, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x73,
	0x74, 0x6f, 0x63, 0x6b, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x65, 0x6c,
	0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x22,
	0x13, 0x0a, 0x11, 0x49, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x79, 0x0a, 0x0a, 0x49, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x69,
	0x65, 0x73, 0x12, 0x3b, 0x0a, 0x0a, 0x69, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x69, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x74, 0x72, 0x61, 0x64, 0x65, 0x72, 0x2e,
	0x49, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x64, 0x75, 0x73,
	0x74, 0x72, 0x79, 0x52, 0x0a, 0x69, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x69, 0x65, 0x73, 0x1a,
	0x2e, 0x0a, 0x08, 0x49, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x32,
	0xfe, 0x01, 0x0a, 0x06, 0x54, 0x72, 0x61, 0x64, 0x65, 0x72, 0x12, 0x55, 0x0a, 0x0f, 0x47, 0x65,
	0x74, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x1b, 0x2e,
	0x74, 0x72, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x52, 0x61, 0x74,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x74, 0x72, 0x61,
	0x64, 0x65, 0x72, 0x2e, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67,
	0x22, 0x0f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x09, 0x12, 0x07, 0x2f, 0x72, 0x61, 0x74, 0x69, 0x6e,
	0x67, 0x12, 0x48, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x14, 0x2e,
	0x74, 0x72, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x74, 0x72, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x6f,
	0x63, 0x6b, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x12, 0x0f, 0x2f, 0x73, 0x74, 0x6f,
	0x63, 0x6b, 0x2f, 0x7b, 0x73, 0x65, 0x63, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x53, 0x0a, 0x0d, 0x47,
	0x65, 0x74, 0x49, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x69, 0x65, 0x73, 0x12, 0x19, 0x2e, 0x74,
	0x72, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x49, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x69, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x74, 0x72, 0x61, 0x64, 0x65, 0x72,
	0x2e, 0x49, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x69, 0x65, 0x73, 0x22, 0x13, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x0d, 0x12, 0x0b, 0x2f, 0x69, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x69, 0x65, 0x73,
	0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72,
	0x75, 0x73, 0x6c, 0x61, 0x6e, 0x31, 0x36, 0x62, 0x6d, 0x73, 0x74, 0x75, 0x2f, 0x67, 0x6f, 0x2d,
	0x61, 0x6c, 0x67, 0x6f, 0x2d, 0x68, 0x61, 0x63, 0x6b, 0x61, 0x74, 0x68, 0x6f, 0x6e, 0x2f, 0x6d,
	0x61, 0x69, 0x6e, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_rawDescOnce sync.Once
	file_api_proto_rawDescData = file_api_proto_rawDesc
)

func file_api_proto_rawDescGZIP() []byte {
	file_api_proto_rawDescOnce.Do(func() {
		file_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_rawDescData)
	})
	return file_api_proto_rawDescData
}

var file_api_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_proto_goTypes = []interface{}{
	(*GlobalRatingRequest)(nil), // 0: trader.GlobalRatingRequest
	(*GlobalRating)(nil),        // 1: trader.GlobalRating
	(*StockRequest)(nil),        // 2: trader.StockRequest
	(*Stock)(nil),               // 3: trader.Stock
	(*IndustriesRequest)(nil),   // 4: trader.IndustriesRequest
	(*Industries)(nil),          // 5: trader.Industries
	(*Industries_Industry)(nil), // 6: trader.Industries.Industry
}
var file_api_proto_depIdxs = []int32{
	3, // 0: trader.GlobalRating.stocks:type_name -> trader.Stock
	6, // 1: trader.Industries.industries:type_name -> trader.Industries.Industry
	0, // 2: trader.Trader.GetGlobalRating:input_type -> trader.GlobalRatingRequest
	2, // 3: trader.Trader.GetStock:input_type -> trader.StockRequest
	4, // 4: trader.Trader.GetIndustries:input_type -> trader.IndustriesRequest
	1, // 5: trader.Trader.GetGlobalRating:output_type -> trader.GlobalRating
	3, // 6: trader.Trader.GetStock:output_type -> trader.Stock
	5, // 7: trader.Trader.GetIndustries:output_type -> trader.Industries
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_proto_init() }
func file_api_proto_init() {
	if File_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GlobalRatingRequest); i {
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
		file_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GlobalRating); i {
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
		file_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StockRequest); i {
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
		file_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Stock); i {
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
		file_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IndustriesRequest); i {
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
		file_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Industries); i {
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
		file_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Industries_Industry); i {
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
			RawDescriptor: file_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_goTypes,
		DependencyIndexes: file_api_proto_depIdxs,
		MessageInfos:      file_api_proto_msgTypes,
	}.Build()
	File_api_proto = out.File
	file_api_proto_rawDesc = nil
	file_api_proto_goTypes = nil
	file_api_proto_depIdxs = nil
}