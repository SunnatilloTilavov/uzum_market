package models

type Empty2 struct{}

type ProductCategory struct {
    ID         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
    ProductID  string `protobuf:"bytes,2,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
    CategoryID string `protobuf:"bytes,3,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
}

type CreateProductCategoryRequest struct {
    ProductID  string `protobuf:"bytes,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
    CategoryID string `protobuf:"bytes,2,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
}

type GetProductCategoriesByProductIDRequest struct {
    ProductID string `protobuf:"bytes,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
}

type GetProductCategoriesByProductIDResponse struct {
    Categories []*ProductCategory `protobuf:"bytes,1,rep,name=categories,proto3" json:"categories,omitempty"`
}

type GetProductCategoriesByCategoryIDRequest struct {
    CategoryID string `protobuf:"bytes,1,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
}

type GetProductCategoriesByCategoryIDResponse struct {
    Categories []*ProductCategory `protobuf:"bytes,1,rep,name=categories,proto3" json:"categories,omitempty"`
}

type DeleteProductCategoryRequest struct {
    ProductID  string `protobuf:"bytes,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
    CategoryID string `protobuf:"bytes,2,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
}
