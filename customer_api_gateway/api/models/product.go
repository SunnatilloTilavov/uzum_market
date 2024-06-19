package models

type Empty1 struct{}

type ProductPrimaryKey struct {
    ID string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

type CreateProduct struct {
    Slug            string   `protobuf:"bytes,1,opt,name=slug,proto3" json:"slug,omitempty"`
    NameUz          string   `protobuf:"bytes,2,opt,name=name_uz,json=nameUz,proto3" json:"name_uz,omitempty"`
    NameRu          string   `protobuf:"bytes,3,opt,name=name_ru,json=nameRu,proto3" json:"name_ru,omitempty"`
    NameEn          string   `protobuf:"bytes,4,opt,name=name_en,json=nameEn,proto3" json:"name_en,omitempty"`
    DescriptionUz   string   `protobuf:"bytes,5,opt,name=description_uz,json=descriptionUz,proto3" json:"description_uz,omitempty"`
    DescriptionRu   string   `protobuf:"bytes,6,opt,name=description_ru,json=descriptionRu,proto3" json:"description_ru,omitempty"`
    DescriptionEn   string   `protobuf:"bytes,7,opt,name=description_en,json=descriptionEn,proto3" json:"description_en,omitempty"`
    Active          bool     `protobuf:"varint,8,opt,name=active,proto3" json:"active,omitempty"`
    OrderNo         int32    `protobuf:"varint,9,opt,name=order_no,json=orderNo,proto3" json:"order_no,omitempty"`
    InPrice         float32  `protobuf:"fixed32,10,opt,name=in_price,json=inPrice,proto3" json:"in_price,omitempty"`
    OutPrice        float32  `protobuf:"fixed32,11,opt,name=out_price,json=outPrice,proto3" json:"out_price,omitempty"`
    LeftCount       int32    `protobuf:"varint,12,opt,name=left_count,json=leftCount,proto3" json:"left_count,omitempty"`
    DiscountPercent float32  `protobuf:"fixed32,13,opt,name=discount_percent,json=discountPercent,proto3" json:"discount_percent,omitempty"`
    Image           []string `protobuf:"bytes,14,rep,name=image,proto3" json:"image,omitempty"`
}

type Product struct {
    ID              string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
    Slug            string   `protobuf:"bytes,2,opt,name=slug,proto3" json:"slug,omitempty"`
    NameUz          string   `protobuf:"bytes,3,opt,name=name_uz,json=nameUz,proto3" json:"name_uz,omitempty"`
    NameRu          string   `protobuf:"bytes,4,opt,name=name_ru,json=nameRu,proto3" json:"name_ru,omitempty"`
    NameEn          string   `protobuf:"bytes,5,opt,name=name_en,json=nameEn,proto3" json:"name_en,omitempty"`
    DescriptionUz   string   `protobuf:"bytes,6,opt,name=description_uz,json=descriptionUz,proto3" json:"description_uz,omitempty"`
    DescriptionRu   string   `protobuf:"bytes,7,opt,name=description_ru,json=descriptionRu,proto3" json:"description_ru,omitempty"`
    DescriptionEn   string   `protobuf:"bytes,8,opt,name=description_en,json=descriptionEn,proto3" json:"description_en,omitempty"`
    Active          bool     `protobuf:"varint,9,opt,name=active,proto3" json:"active,omitempty"`
    OrderNo         int32    `protobuf:"varint,10,opt,name=order_no,json=orderNo,proto3" json:"order_no,omitempty"`
    InPrice         float32  `protobuf:"fixed32,11,opt,name=in_price,json=inPrice,proto3" json:"in_price,omitempty"`
    OutPrice        float32  `protobuf:"fixed32,12,opt,name=out_price,json=outPrice,proto3" json:"out_price,omitempty"`
    LeftCount       int32    `protobuf:"varint,13,opt,name=left_count,json=leftCount,proto3" json:"left_count,omitempty"`
    DiscountPercent float32  `protobuf:"fixed32,14,opt,name=discount_percent,json=discountPercent,proto3" json:"discount_percent,omitempty"`
    Image           []string `protobuf:"bytes,15,rep,name=image,proto3" json:"image,omitempty"`
    CreatedAt       string   `protobuf:"bytes,16,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
    UpdatedAt       string   `protobuf:"bytes,17,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
    DeletedAt       string   `protobuf:"bytes,18,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
}

type UpdateProduct struct {
    ID              string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
    Slug            string   `protobuf:"bytes,2,opt,name=slug,proto3" json:"slug,omitempty"`
    NameUz          string   `protobuf:"bytes,3,opt,name=name_uz,json=nameUz,proto3" json:"name_uz,omitempty"`
    NameRu          string   `protobuf:"bytes,4,opt,name=name_ru,json=nameRu,proto3" json:"name_ru,omitempty"`
    NameEn          string   `protobuf:"bytes,5,opt,name=name_en,json=nameEn,proto3" json:"name_en,omitempty"`
    DescriptionUz   string   `protobuf:"bytes,6,opt,name=description_uz,json=descriptionUz,proto3" json:"description_uz,omitempty"`
    DescriptionRu   string   `protobuf:"bytes,7,opt,name=description_ru,json=descriptionRu,proto3" json:"description_ru,omitempty"`
    DescriptionEn   string   `protobuf:"bytes,8,opt,name=description_en,json=descriptionEn,proto3" json:"description_en,omitempty"`
    Active          bool     `protobuf:"varint,9,opt,name=active,proto3" json:"active,omitempty"`
    OrderNo         int32    `protobuf:"varint,10,opt,name=order_no,json=orderNo,proto3" json:"order_no,omitempty"`
    InPrice         float32  `protobuf:"fixed32,11,opt,name=in_price,json=inPrice,proto3" json:"in_price,omitempty"`
    OutPrice        float32  `protobuf:"fixed32,12,opt,name=out_price,json=outPrice,proto3" json:"out_price,omitempty"`
    LeftCount       int32    `protobuf:"varint,13,opt,name=left_count,json=leftCount,proto3" json:"left_count,omitempty"`
    DiscountPercent float32  `protobuf:"fixed32,14,opt,name=discount_percent,json=discountPercent,proto3" json:"discount_percent,omitempty"`
    Image           []string `protobuf:"bytes,15,rep,name=image,proto3" json:"image,omitempty"`
}

type GetListProductRequest struct {
    Offset int64  `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
    Limit  int64  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
    Search string `protobuf:"bytes,3,opt,name=search,proto3" json:"search,omitempty"`
}

type GetListProductResponse struct {
    Count    int64      `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
    Products []*Product `protobuf:"bytes,2,rep,name=products,proto3" json:"products,omitempty"`
}
