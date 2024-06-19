// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: order_products.proto

package order_product_service

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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_products_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_order_products_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_order_products_proto_rawDescGZIP(), []int{0}
}

type OrderProductPrimaryKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *OrderProductPrimaryKey) Reset() {
	*x = OrderProductPrimaryKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_products_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderProductPrimaryKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderProductPrimaryKey) ProtoMessage() {}

func (x *OrderProductPrimaryKey) ProtoReflect() protoreflect.Message {
	mi := &file_order_products_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderProductPrimaryKey.ProtoReflect.Descriptor instead.
func (*OrderProductPrimaryKey) Descriptor() ([]byte, []int) {
	return file_order_products_proto_rawDescGZIP(), []int{1}
}

func (x *OrderProductPrimaryKey) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CreateOrderProduct struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId     string  `protobuf:"bytes,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Count         int64   `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	DiscountPrice float64 `protobuf:"fixed64,3,opt,name=discount_price,json=discountPrice,proto3" json:"discount_price,omitempty"`
	Price         float64 `protobuf:"fixed64,4,opt,name=price,proto3" json:"price,omitempty"`
	OrderId       string  `protobuf:"bytes,5,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
}

func (x *CreateOrderProduct) Reset() {
	*x = CreateOrderProduct{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_products_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOrderProduct) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderProduct) ProtoMessage() {}

func (x *CreateOrderProduct) ProtoReflect() protoreflect.Message {
	mi := &file_order_products_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderProduct.ProtoReflect.Descriptor instead.
func (*CreateOrderProduct) Descriptor() ([]byte, []int) {
	return file_order_products_proto_rawDescGZIP(), []int{2}
}

func (x *CreateOrderProduct) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *CreateOrderProduct) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *CreateOrderProduct) GetDiscountPrice() float64 {
	if x != nil {
		return x.DiscountPrice
	}
	return 0
}

func (x *CreateOrderProduct) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *CreateOrderProduct) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

type OrderProduct struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ProductId     string  `protobuf:"bytes,2,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Count         int64   `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
	DiscountPrice float64 `protobuf:"fixed64,4,opt,name=discount_price,json=discountPrice,proto3" json:"discount_price,omitempty"`
	Price         float64 `protobuf:"fixed64,5,opt,name=price,proto3" json:"price,omitempty"`
	OrderId       string  `protobuf:"bytes,6,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	CreatedAt     string  `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     string  `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *OrderProduct) Reset() {
	*x = OrderProduct{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_products_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderProduct) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderProduct) ProtoMessage() {}

func (x *OrderProduct) ProtoReflect() protoreflect.Message {
	mi := &file_order_products_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderProduct.ProtoReflect.Descriptor instead.
func (*OrderProduct) Descriptor() ([]byte, []int) {
	return file_order_products_proto_rawDescGZIP(), []int{3}
}

func (x *OrderProduct) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *OrderProduct) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *OrderProduct) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *OrderProduct) GetDiscountPrice() float64 {
	if x != nil {
		return x.DiscountPrice
	}
	return 0
}

func (x *OrderProduct) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *OrderProduct) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *OrderProduct) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *OrderProduct) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type UpdateOrderProduct struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ProductId     string  `protobuf:"bytes,2,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Count         int64   `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
	DiscountPrice float64 `protobuf:"fixed64,4,opt,name=discount_price,json=discountPrice,proto3" json:"discount_price,omitempty"`
	Price         float64 `protobuf:"fixed64,5,opt,name=price,proto3" json:"price,omitempty"`
	OrderId       string  `protobuf:"bytes,6,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	DeletedAt     int64   `protobuf:"varint,7,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
}

func (x *UpdateOrderProduct) Reset() {
	*x = UpdateOrderProduct{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_products_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateOrderProduct) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateOrderProduct) ProtoMessage() {}

func (x *UpdateOrderProduct) ProtoReflect() protoreflect.Message {
	mi := &file_order_products_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateOrderProduct.ProtoReflect.Descriptor instead.
func (*UpdateOrderProduct) Descriptor() ([]byte, []int) {
	return file_order_products_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateOrderProduct) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateOrderProduct) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *UpdateOrderProduct) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *UpdateOrderProduct) GetDiscountPrice() float64 {
	if x != nil {
		return x.DiscountPrice
	}
	return 0
}

func (x *UpdateOrderProduct) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *UpdateOrderProduct) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *UpdateOrderProduct) GetDeletedAt() int64 {
	if x != nil {
		return x.DeletedAt
	}
	return 0
}

type GetListOrderProductRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int64 `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit  int64 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *GetListOrderProductRequest) Reset() {
	*x = GetListOrderProductRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_products_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListOrderProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListOrderProductRequest) ProtoMessage() {}

func (x *GetListOrderProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_order_products_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListOrderProductRequest.ProtoReflect.Descriptor instead.
func (*GetListOrderProductRequest) Descriptor() ([]byte, []int) {
	return file_order_products_proto_rawDescGZIP(), []int{5}
}

func (x *GetListOrderProductRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetListOrderProductRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetListOrderProductResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count         int64           `protobuf:"varint,1,opt,name=Count,proto3" json:"Count,omitempty"`
	OrderProducts []*OrderProduct `protobuf:"bytes,2,rep,name=OrderProducts,proto3" json:"OrderProducts,omitempty"`
}

func (x *GetListOrderProductResponse) Reset() {
	*x = GetListOrderProductResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_products_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListOrderProductResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListOrderProductResponse) ProtoMessage() {}

func (x *GetListOrderProductResponse) ProtoReflect() protoreflect.Message {
	mi := &file_order_products_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListOrderProductResponse.ProtoReflect.Descriptor instead.
func (*GetListOrderProductResponse) Descriptor() ([]byte, []int) {
	return file_order_products_proto_rawDescGZIP(), []int{6}
}

func (x *GetListOrderProductResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *GetListOrderProductResponse) GetOrderProducts() []*OrderProduct {
	if x != nil {
		return x.OrderProducts
	}
	return nil
}

var File_order_products_proto protoreflect.FileDescriptor

var file_order_products_proto_rawDesc = []byte{
	0x0a, 0x14, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x28, 0x0a, 0x16, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50,
	0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0xa1, 0x01, 0x0a, 0x12, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0d, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x22, 0xe9, 0x01,
	0x0a, 0x0c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0d, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xd0, 0x01, 0x0a, 0x12, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0d, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x4a, 0x0a, 0x1a,
	0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x77, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x4c,
	0x69, 0x73, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x42, 0x0a,
	0x0d, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x52, 0x0d, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x73, 0x32, 0xae, 0x03, 0x0a, 0x0d, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x73, 0x12, 0x4c, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x22, 0x2e,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x1a, 0x1c, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x22,
	0x00, 0x12, 0x51, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x26, 0x2e, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x2e, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72,
	0x79, 0x4b, 0x65, 0x79, 0x1a, 0x1c, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x22, 0x00, 0x12, 0x63, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x2a,
	0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x2e,
	0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x4c,
	0x69, 0x73, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4c, 0x0a, 0x06, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x22, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x1a, 0x1c, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x12, 0x26, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50,
	0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x1a, 0x15, 0x2e, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x00, 0x42, 0x20, 0x5a, 0x1e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_order_products_proto_rawDescOnce sync.Once
	file_order_products_proto_rawDescData = file_order_products_proto_rawDesc
)

func file_order_products_proto_rawDescGZIP() []byte {
	file_order_products_proto_rawDescOnce.Do(func() {
		file_order_products_proto_rawDescData = protoimpl.X.CompressGZIP(file_order_products_proto_rawDescData)
	})
	return file_order_products_proto_rawDescData
}

var file_order_products_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_order_products_proto_goTypes = []interface{}{
	(*Empty)(nil),                       // 0: order_products.Empty
	(*OrderProductPrimaryKey)(nil),      // 1: order_products.OrderProductPrimaryKey
	(*CreateOrderProduct)(nil),          // 2: order_products.CreateOrderProduct
	(*OrderProduct)(nil),                // 3: order_products.OrderProduct
	(*UpdateOrderProduct)(nil),          // 4: order_products.UpdateOrderProduct
	(*GetListOrderProductRequest)(nil),  // 5: order_products.GetListOrderProductRequest
	(*GetListOrderProductResponse)(nil), // 6: order_products.GetListOrderProductResponse
}
var file_order_products_proto_depIdxs = []int32{
	3, // 0: order_products.GetListOrderProductResponse.OrderProducts:type_name -> order_products.OrderProduct
	2, // 1: order_products.OrderProducts.Create:input_type -> order_products.CreateOrderProduct
	1, // 2: order_products.OrderProducts.GetById:input_type -> order_products.OrderProductPrimaryKey
	5, // 3: order_products.OrderProducts.GetAll:input_type -> order_products.GetListOrderProductRequest
	4, // 4: order_products.OrderProducts.Update:input_type -> order_products.UpdateOrderProduct
	1, // 5: order_products.OrderProducts.Delete:input_type -> order_products.OrderProductPrimaryKey
	3, // 6: order_products.OrderProducts.Create:output_type -> order_products.OrderProduct
	3, // 7: order_products.OrderProducts.GetById:output_type -> order_products.OrderProduct
	6, // 8: order_products.OrderProducts.GetAll:output_type -> order_products.GetListOrderProductResponse
	3, // 9: order_products.OrderProducts.Update:output_type -> order_products.OrderProduct
	0, // 10: order_products.OrderProducts.Delete:output_type -> order_products.Empty
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_order_products_proto_init() }
func file_order_products_proto_init() {
	if File_order_products_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_order_products_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_order_products_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderProductPrimaryKey); i {
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
		file_order_products_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOrderProduct); i {
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
		file_order_products_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderProduct); i {
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
		file_order_products_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateOrderProduct); i {
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
		file_order_products_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListOrderProductRequest); i {
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
		file_order_products_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListOrderProductResponse); i {
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
			RawDescriptor: file_order_products_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_order_products_proto_goTypes,
		DependencyIndexes: file_order_products_proto_depIdxs,
		MessageInfos:      file_order_products_proto_msgTypes,
	}.Build()
	File_order_products_proto = out.File
	file_order_products_proto_rawDesc = nil
	file_order_products_proto_goTypes = nil
	file_order_products_proto_depIdxs = nil
}
