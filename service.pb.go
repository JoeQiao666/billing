// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.0
// source: service.proto

package billing

import (
	_ "github.com/cloudstateio/go-support/cloudstate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Price struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId   string `protobuf:"bytes,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	ProductName string `protobuf:"bytes,2,opt,name=product_name,json=productName,proto3" json:"product_name,omitempty"`
	Price       int32  `protobuf:"varint,3,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *Price) Reset() {
	*x = Price{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Price) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Price) ProtoMessage() {}

func (x *Price) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Price.ProtoReflect.Descriptor instead.
func (*Price) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{0}
}

func (x *Price) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *Price) GetProductName() string {
	if x != nil {
		return x.ProductName
	}
	return ""
}

func (x *Price) GetPrice() int32 {
	if x != nil {
		return x.Price
	}
	return 0
}

type Prices struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Prices []*Price `protobuf:"bytes,1,rep,name=prices,proto3" json:"prices,omitempty"`
}

func (x *Prices) Reset() {
	*x = Prices{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Prices) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Prices) ProtoMessage() {}

func (x *Prices) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Prices.ProtoReflect.Descriptor instead.
func (*Prices) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{1}
}

func (x *Prices) GetPrices() []*Price {
	if x != nil {
		return x.Prices
	}
	return nil
}

type AddOrUpdatePrices struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BillingCenterId string   `protobuf:"bytes,1,opt,name=billing_center_id,json=billingCenterId,proto3" json:"billing_center_id,omitempty"`
	Prices          []*Price `protobuf:"bytes,2,rep,name=prices,proto3" json:"prices,omitempty"`
}

func (x *AddOrUpdatePrices) Reset() {
	*x = AddOrUpdatePrices{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddOrUpdatePrices) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddOrUpdatePrices) ProtoMessage() {}

func (x *AddOrUpdatePrices) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddOrUpdatePrices.ProtoReflect.Descriptor instead.
func (*AddOrUpdatePrices) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{2}
}

func (x *AddOrUpdatePrices) GetBillingCenterId() string {
	if x != nil {
		return x.BillingCenterId
	}
	return ""
}

func (x *AddOrUpdatePrices) GetPrices() []*Price {
	if x != nil {
		return x.Prices
	}
	return nil
}

type DeletePrices struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BillingCenterId string   `protobuf:"bytes,1,opt,name=billing_center_id,json=billingCenterId,proto3" json:"billing_center_id,omitempty"`
	Prices          []*Price `protobuf:"bytes,2,rep,name=prices,proto3" json:"prices,omitempty"`
}

func (x *DeletePrices) Reset() {
	*x = DeletePrices{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePrices) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePrices) ProtoMessage() {}

func (x *DeletePrices) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePrices.ProtoReflect.Descriptor instead.
func (*DeletePrices) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{3}
}

func (x *DeletePrices) GetBillingCenterId() string {
	if x != nil {
		return x.BillingCenterId
	}
	return ""
}

func (x *DeletePrices) GetPrices() []*Price {
	if x != nil {
		return x.Prices
	}
	return nil
}

type QueryPrices struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BillingCenterId string `protobuf:"bytes,1,opt,name=billing_center_id,json=billingCenterId,proto3" json:"billing_center_id,omitempty"`
}

func (x *QueryPrices) Reset() {
	*x = QueryPrices{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryPrices) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryPrices) ProtoMessage() {}

func (x *QueryPrices) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryPrices.ProtoReflect.Descriptor instead.
func (*QueryPrices) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{4}
}

func (x *QueryPrices) GetBillingCenterId() string {
	if x != nil {
		return x.BillingCenterId
	}
	return ""
}

type Product struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId string `protobuf:"bytes,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Quantity  int32  `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *Product) Reset() {
	*x = Product{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product) ProtoMessage() {}

func (x *Product) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Product.ProtoReflect.Descriptor instead.
func (*Product) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{5}
}

func (x *Product) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *Product) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Product) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type CalculatePrice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BillingCenterId string     `protobuf:"bytes,1,opt,name=billing_center_id,json=billingCenterId,proto3" json:"billing_center_id,omitempty"`
	Products        []*Product `protobuf:"bytes,2,rep,name=products,proto3" json:"products,omitempty"`
}

func (x *CalculatePrice) Reset() {
	*x = CalculatePrice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalculatePrice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalculatePrice) ProtoMessage() {}

func (x *CalculatePrice) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalculatePrice.ProtoReflect.Descriptor instead.
func (*CalculatePrice) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{6}
}

func (x *CalculatePrice) GetBillingCenterId() string {
	if x != nil {
		return x.BillingCenterId
	}
	return ""
}

func (x *CalculatePrice) GetProducts() []*Product {
	if x != nil {
		return x.Products
	}
	return nil
}

type ProductsWithTotalPrice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Products   []*Product `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"`
	TotalPrice int32      `protobuf:"varint,2,opt,name=totalPrice,proto3" json:"totalPrice,omitempty"`
}

func (x *ProductsWithTotalPrice) Reset() {
	*x = ProductsWithTotalPrice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductsWithTotalPrice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductsWithTotalPrice) ProtoMessage() {}

func (x *ProductsWithTotalPrice) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductsWithTotalPrice.ProtoReflect.Descriptor instead.
func (*ProductsWithTotalPrice) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{7}
}

func (x *ProductsWithTotalPrice) GetProducts() []*Product {
	if x != nil {
		return x.Products
	}
	return nil
}

func (x *ProductsWithTotalPrice) GetTotalPrice() int32 {
	if x != nil {
		return x.TotalPrice
	}
	return 0
}

type Resp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductsWithTotalPrice *ProductsWithTotalPrice `protobuf:"bytes,1,opt,name=productsWithTotalPrice,proto3" json:"productsWithTotalPrice,omitempty"`
}

func (x *Resp) Reset() {
	*x = Resp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Resp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Resp) ProtoMessage() {}

func (x *Resp) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Resp.ProtoReflect.Descriptor instead.
func (*Resp) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{8}
}

func (x *Resp) GetProductsWithTotalPrice() *ProductsWithTotalPrice {
	if x != nil {
		return x.ProductsWithTotalPrice
	}
	return nil
}

var File_service_proto protoreflect.FileDescriptor

var file_service_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x6b, 0x65, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x5f, 0x0a, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x22, 0x30, 0x0a, 0x06, 0x50, 0x72, 0x69, 0x63, 0x65, 0x73, 0x12, 0x26, 0x0a, 0x06, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x62, 0x69,
	0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x06, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x73, 0x22, 0x6c, 0x0a, 0x11, 0x41, 0x64, 0x64, 0x4f, 0x72, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x73, 0x12, 0x2f, 0x0a, 0x11, 0x62, 0x69, 0x6c, 0x6c,
	0x69, 0x6e, 0x67, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x03, 0xc0, 0x43, 0x01, 0x52, 0x0f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e,
	0x67, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x06, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x62, 0x69, 0x6c, 0x6c,
	0x69, 0x6e, 0x67, 0x2e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x06, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x73, 0x22, 0x67, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x73, 0x12, 0x2f, 0x0a, 0x11, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xc0, 0x43,
	0x01, 0x52, 0x0f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x26, 0x0a, 0x06, 0x70, 0x72, 0x69, 0x63, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x52, 0x06, 0x70, 0x72, 0x69, 0x63, 0x65, 0x73, 0x22, 0x3e, 0x0a, 0x0b, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x50, 0x72, 0x69, 0x63, 0x65, 0x73, 0x12, 0x2f, 0x0a, 0x11, 0x62, 0x69, 0x6c,
	0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xc0, 0x43, 0x01, 0x52, 0x0f, 0x62, 0x69, 0x6c, 0x6c, 0x69,
	0x6e, 0x67, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x49, 0x64, 0x22, 0x58, 0x0a, 0x07, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x22, 0x6f, 0x0a, 0x0e, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74,
	0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x2f, 0x0a, 0x11, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e,
	0x67, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x03, 0xc0, 0x43, 0x01, 0x52, 0x0f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x43,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x62, 0x69, 0x6c, 0x6c,
	0x69, 0x6e, 0x67, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x08, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x73, 0x22, 0x66, 0x0a, 0x16, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x73, 0x57, 0x69, 0x74, 0x68, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12,
	0x2c, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x1e, 0x0a,
	0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x22, 0x5f, 0x0a,
	0x04, 0x52, 0x65, 0x73, 0x70, 0x12, 0x57, 0x0a, 0x16, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x73, 0x57, 0x69, 0x74, 0x68, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x57, 0x69, 0x74, 0x68, 0x54, 0x6f, 0x74, 0x61,
	0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x16, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73,
	0x57, 0x69, 0x74, 0x68, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x32, 0xf7,
	0x03, 0x0a, 0x07, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x12, 0x8b, 0x01, 0x0a, 0x19, 0x41,
	0x64, 0x64, 0x4f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x73, 0x50, 0x72, 0x69, 0x63, 0x65, 0x73, 0x12, 0x1a, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69,
	0x6e, 0x67, 0x2e, 0x41, 0x64, 0x64, 0x4f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x73, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x3a, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x34, 0x22, 0x2f, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x7b,
	0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x7d, 0x2f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x61, 0x64, 0x64, 0x4f, 0x72, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x7c, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x50, 0x72, 0x69, 0x63, 0x65, 0x73,
	0x12, 0x15, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x35, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2f, 0x22, 0x2a, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e,
	0x67, 0x2f, 0x7b, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x67, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x73, 0x12, 0x14, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x50, 0x72, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x0f, 0x2e, 0x62, 0x69, 0x6c, 0x6c,
	0x69, 0x6e, 0x67, 0x2e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x73, 0x22, 0x33, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x2d, 0x12, 0x23, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x7b, 0x62, 0x69,
	0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d,
	0x2f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x69, 0x63, 0x65, 0x73, 0x12,
	0x77, 0x0a, 0x13, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x74, 0x61,
	0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x17, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67,
	0x2e, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x1a,
	0x0d, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x22, 0x38,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x32, 0x22, 0x2d, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67,
	0x2f, 0x7b, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x61, 0x6c, 0x63,
	0x75, 0x6c, 0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4a, 0x6f, 0x65, 0x51, 0x69, 0x61, 0x6f, 0x36, 0x36,
	0x36, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x3b, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e,
	0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_proto_rawDescOnce sync.Once
	file_service_proto_rawDescData = file_service_proto_rawDesc
)

func file_service_proto_rawDescGZIP() []byte {
	file_service_proto_rawDescOnce.Do(func() {
		file_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_proto_rawDescData)
	})
	return file_service_proto_rawDescData
}

var file_service_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_service_proto_goTypes = []interface{}{
	(*Price)(nil),                  // 0: billing.Price
	(*Prices)(nil),                 // 1: billing.Prices
	(*AddOrUpdatePrices)(nil),      // 2: billing.AddOrUpdatePrices
	(*DeletePrices)(nil),           // 3: billing.DeletePrices
	(*QueryPrices)(nil),            // 4: billing.QueryPrices
	(*Product)(nil),                // 5: billing.Product
	(*CalculatePrice)(nil),         // 6: billing.CalculatePrice
	(*ProductsWithTotalPrice)(nil), // 7: billing.ProductsWithTotalPrice
	(*Resp)(nil),                   // 8: billing.Resp
	(*emptypb.Empty)(nil),          // 9: google.protobuf.Empty
}
var file_service_proto_depIdxs = []int32{
	0,  // 0: billing.Prices.prices:type_name -> billing.Price
	0,  // 1: billing.AddOrUpdatePrices.prices:type_name -> billing.Price
	0,  // 2: billing.DeletePrices.prices:type_name -> billing.Price
	5,  // 3: billing.CalculatePrice.products:type_name -> billing.Product
	5,  // 4: billing.ProductsWithTotalPrice.products:type_name -> billing.Product
	7,  // 5: billing.Resp.productsWithTotalPrice:type_name -> billing.ProductsWithTotalPrice
	2,  // 6: billing.Billing.AddOrUpdateProductsPrices:input_type -> billing.AddOrUpdatePrices
	3,  // 7: billing.Billing.DeleteProductsPrices:input_type -> billing.DeletePrices
	4,  // 8: billing.Billing.GetPrices:input_type -> billing.QueryPrices
	6,  // 9: billing.Billing.CalculateTotalPrice:input_type -> billing.CalculatePrice
	9,  // 10: billing.Billing.AddOrUpdateProductsPrices:output_type -> google.protobuf.Empty
	9,  // 11: billing.Billing.DeleteProductsPrices:output_type -> google.protobuf.Empty
	1,  // 12: billing.Billing.GetPrices:output_type -> billing.Prices
	8,  // 13: billing.Billing.CalculateTotalPrice:output_type -> billing.Resp
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_service_proto_init() }
func file_service_proto_init() {
	if File_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Price); i {
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
		file_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Prices); i {
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
		file_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddOrUpdatePrices); i {
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
		file_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePrices); i {
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
		file_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryPrices); i {
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
		file_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Product); i {
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
		file_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalculatePrice); i {
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
		file_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductsWithTotalPrice); i {
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
		file_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Resp); i {
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
			RawDescriptor: file_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_proto_goTypes,
		DependencyIndexes: file_service_proto_depIdxs,
		MessageInfos:      file_service_proto_msgTypes,
	}.Build()
	File_service_proto = out.File
	file_service_proto_rawDesc = nil
	file_service_proto_goTypes = nil
	file_service_proto_depIdxs = nil
}
