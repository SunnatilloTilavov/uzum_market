package models

type Empty4 struct{}

type ProductReviewPrimaryKey struct {
	ID string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

type CreateProductReviewRequest struct {
	CustomerId string  `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	ProductId  string  `protobuf:"bytes,2,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Text       string  `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
	Rating     float32 `protobuf:"fixed32,4,opt,name=rating,proto3" json:"rating,omitempty"`
	OrderId    string  `protobuf:"bytes,5,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
}


type ProductReview struct {
	Id         string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CustomerId string  `protobuf:"bytes,2,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	ProductId  string  `protobuf:"bytes,3,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Text       string  `protobuf:"bytes,4,opt,name=text,proto3" json:"text,omitempty"`
	Rating     float32 `protobuf:"fixed32,5,opt,name=rating,proto3" json:"rating,omitempty"`
	OrderId    string  `protobuf:"bytes,6,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	CreatedAt  string  `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

type UpdateProductReviewRequest struct {
	Id     string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Text   string  `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	Rating float32 `protobuf:"fixed32,3,opt,name=rating,proto3" json:"rating,omitempty"`
}

type GetProductReviewsByProductIDRequest struct {
	ProductId string `protobuf:"bytes,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
}

type GetProductReviewsByProductIDResponse struct {
	Reviews []*ProductReview `protobuf:"bytes,1,rep,name=reviews,proto3" json:"reviews,omitempty"`
}

type GetProductReviewsByCustomerIDRequest struct {
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
}

type GetProductReviewsByCustomerIDResponse struct {
	Reviews []*ProductReview `protobuf:"bytes,1,rep,name=reviews,proto3" json:"reviews,omitempty"`
}
