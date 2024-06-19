package models

type Empty struct{}

type CategoryPrimaryKey struct {
    ID string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

type CreateCategory struct {
    Slug     string `protobuf:"bytes,1,opt,name=slug,proto3" json:"slug,omitempty"`
    NameUz   string `protobuf:"bytes,2,opt,name=name_uz,json=nameUz,proto3" json:"name_uz,omitempty"`
    NameRu   string `protobuf:"bytes,3,opt,name=name_ru,json=nameRu,proto3" json:"name_ru,omitempty"`
    NameEn   string `protobuf:"bytes,4,opt,name=name_en,json=nameEn,proto3" json:"name_en,omitempty"`
    Active   bool   `protobuf:"varint,5,opt,name=active,proto3" json:"active,omitempty"`
    OrderNo  int32  `protobuf:"varint,6,opt,name=order_no,json=orderNo,proto3" json:"order_no,omitempty"`
    ParentID string `protobuf:"bytes,7,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
}

type Category struct {
    ID        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
    Slug      string `protobuf:"bytes,2,opt,name=slug,proto3" json:"slug,omitempty"`
    NameUz    string `protobuf:"bytes,3,opt,name=name_uz,json=nameUz,proto3" json:"name_uz,omitempty"`
    NameRu    string `protobuf:"bytes,4,opt,name=name_ru,json=nameRu,proto3" json:"name_ru,omitempty"`
    NameEn    string `protobuf:"bytes,5,opt,name=name_en,json=nameEn,proto3" json:"name_en,omitempty"`
    Active    bool   `protobuf:"varint,6,opt,name=active,proto3" json:"active,omitempty"`
    OrderNo   int32  `protobuf:"varint,7,opt,name=order_no,json=orderNo,proto3" json:"order_no,omitempty"`
    ParentID  string `protobuf:"bytes,8,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
    CreatedAt string `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
    UpdatedAt string `protobuf:"bytes,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
    DeletedAt int64  `protobuf:"varint,11,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
}

type UpdateCategory struct {
    ID       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
    Slug     string `protobuf:"bytes,2,opt,name=slug,proto3" json:"slug,omitempty"`
    NameUz   string `protobuf:"bytes,3,opt,name=name_uz,json=nameUz,proto3" json:"name_uz,omitempty"`
    NameRu   string `protobuf:"bytes,4,opt,name=name_ru,json=nameRu,proto3" json:"name_ru,omitempty"`
    NameEn   string `protobuf:"bytes,5,opt,name=name_en,json=nameEn,proto3" json:"name_en,omitempty"`
    Active   bool   `protobuf:"varint,6,opt,name=active,proto3" json:"active,omitempty"`
    OrderNo  int32  `protobuf:"varint,7,opt,name=order_no,json=orderNo,proto3" json:"order_no,omitempty"`
    ParentID string `protobuf:"bytes,8,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
}

type GetListCategoryRequest struct {
    Offset int64  `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
    Limit  int64  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
    Search string `protobuf:"bytes,3,opt,name=search,proto3" json:"search,omitempty"`
}

type GetListCategoryResponse struct {
    Count     int64      `protobuf:"varint,1,opt,name=Count,proto3" json:"Count,omitempty"`
    Categorys []*Category `protobuf:"bytes,2,rep,name=Categorys,proto3" json:"Categorys,omitempty"`
}
