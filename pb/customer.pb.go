// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.8
// source: pb/customer.proto

package pb

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

type CustomerLocation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address    string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Village    string `protobuf:"bytes,2,opt,name=village,proto3" json:"village,omitempty"`
	District   string `protobuf:"bytes,3,opt,name=district,proto3" json:"district,omitempty"`
	City       string `protobuf:"bytes,4,opt,name=city,proto3" json:"city,omitempty"`
	Province   string `protobuf:"bytes,5,opt,name=province,proto3" json:"province,omitempty"`
	PostalCode string `protobuf:"bytes,6,opt,name=postal_code,json=postalCode,proto3" json:"postal_code,omitempty"`
}

func (x *CustomerLocation) Reset() {
	*x = CustomerLocation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_customer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomerLocation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerLocation) ProtoMessage() {}

func (x *CustomerLocation) ProtoReflect() protoreflect.Message {
	mi := &file_pb_customer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerLocation.ProtoReflect.Descriptor instead.
func (*CustomerLocation) Descriptor() ([]byte, []int) {
	return file_pb_customer_proto_rawDescGZIP(), []int{0}
}

func (x *CustomerLocation) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *CustomerLocation) GetVillage() string {
	if x != nil {
		return x.Village
	}
	return ""
}

func (x *CustomerLocation) GetDistrict() string {
	if x != nil {
		return x.District
	}
	return ""
}

func (x *CustomerLocation) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *CustomerLocation) GetProvince() string {
	if x != nil {
		return x.Province
	}
	return ""
}

func (x *CustomerLocation) GetPostalCode() string {
	if x != nil {
		return x.PostalCode
	}
	return ""
}

type CustomerDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Npsn     string            `protobuf:"bytes,1,opt,name=npsn,proto3" json:"npsn,omitempty"`
	Name     string            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email    string            `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Wa       string            `protobuf:"bytes,4,opt,name=wa,proto3" json:"wa,omitempty"`
	Type     string            `protobuf:"bytes,5,opt,name=type,proto3" json:"type,omitempty"`
	Level    string            `protobuf:"bytes,6,opt,name=level,proto3" json:"level,omitempty"`
	About    string            `protobuf:"bytes,7,opt,name=about,proto3" json:"about,omitempty"`
	Location *CustomerLocation `protobuf:"bytes,8,opt,name=location,proto3" json:"location,omitempty"`
}

func (x *CustomerDetail) Reset() {
	*x = CustomerDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_customer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomerDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerDetail) ProtoMessage() {}

func (x *CustomerDetail) ProtoReflect() protoreflect.Message {
	mi := &file_pb_customer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerDetail.ProtoReflect.Descriptor instead.
func (*CustomerDetail) Descriptor() ([]byte, []int) {
	return file_pb_customer_proto_rawDescGZIP(), []int{1}
}

func (x *CustomerDetail) GetNpsn() string {
	if x != nil {
		return x.Npsn
	}
	return ""
}

func (x *CustomerDetail) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CustomerDetail) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CustomerDetail) GetWa() string {
	if x != nil {
		return x.Wa
	}
	return ""
}

func (x *CustomerDetail) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *CustomerDetail) GetLevel() string {
	if x != nil {
		return x.Level
	}
	return ""
}

func (x *CustomerDetail) GetAbout() string {
	if x != nil {
		return x.About
	}
	return ""
}

func (x *CustomerDetail) GetLocation() *CustomerLocation {
	if x != nil {
		return x.Location
	}
	return nil
}

type CustomerCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User   string          `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Pass   string          `protobuf:"bytes,2,opt,name=pass,proto3" json:"pass,omitempty"`
	Detail *CustomerDetail `protobuf:"bytes,3,opt,name=detail,proto3" json:"detail,omitempty"`
}

func (x *CustomerCreateRequest) Reset() {
	*x = CustomerCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_customer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomerCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerCreateRequest) ProtoMessage() {}

func (x *CustomerCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_customer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerCreateRequest.ProtoReflect.Descriptor instead.
func (*CustomerCreateRequest) Descriptor() ([]byte, []int) {
	return file_pb_customer_proto_rawDescGZIP(), []int{2}
}

func (x *CustomerCreateRequest) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *CustomerCreateRequest) GetPass() string {
	if x != nil {
		return x.Pass
	}
	return ""
}

func (x *CustomerCreateRequest) GetDetail() *CustomerDetail {
	if x != nil {
		return x.Detail
	}
	return nil
}

type Customer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId string          `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	User       string          `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	Pass       string          `protobuf:"bytes,3,opt,name=pass,proto3" json:"pass,omitempty"`
	Detail     *CustomerDetail `protobuf:"bytes,4,opt,name=detail,proto3" json:"detail,omitempty"`
	Status     string          `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	ExpUntil   int64           `protobuf:"varint,6,opt,name=exp_until,json=expUntil,proto3" json:"exp_until,omitempty"`
	CreatedAt  int64           `protobuf:"varint,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt  int64           `protobuf:"varint,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt  int64           `protobuf:"varint,9,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
}

func (x *Customer) Reset() {
	*x = Customer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_customer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Customer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Customer) ProtoMessage() {}

func (x *Customer) ProtoReflect() protoreflect.Message {
	mi := &file_pb_customer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Customer.ProtoReflect.Descriptor instead.
func (*Customer) Descriptor() ([]byte, []int) {
	return file_pb_customer_proto_rawDescGZIP(), []int{3}
}

func (x *Customer) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *Customer) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *Customer) GetPass() string {
	if x != nil {
		return x.Pass
	}
	return ""
}

func (x *Customer) GetDetail() *CustomerDetail {
	if x != nil {
		return x.Detail
	}
	return nil
}

func (x *Customer) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Customer) GetExpUntil() int64 {
	if x != nil {
		return x.ExpUntil
	}
	return 0
}

func (x *Customer) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Customer) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *Customer) GetDeletedAt() int64 {
	if x != nil {
		return x.DeletedAt
	}
	return 0
}

type CustomerFindOneRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	User       string `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	Pass       string `protobuf:"bytes,3,opt,name=pass,proto3" json:"pass,omitempty"`
}

func (x *CustomerFindOneRequest) Reset() {
	*x = CustomerFindOneRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_customer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomerFindOneRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerFindOneRequest) ProtoMessage() {}

func (x *CustomerFindOneRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_customer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerFindOneRequest.ProtoReflect.Descriptor instead.
func (*CustomerFindOneRequest) Descriptor() ([]byte, []int) {
	return file_pb_customer_proto_rawDescGZIP(), []int{4}
}

func (x *CustomerFindOneRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *CustomerFindOneRequest) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *CustomerFindOneRequest) GetPass() string {
	if x != nil {
		return x.Pass
	}
	return ""
}

type CustomerFindAllRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sort    string `protobuf:"bytes,1,opt,name=sort,proto3" json:"sort,omitempty"`
	Page    int64  `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	PerPage int64  `protobuf:"varint,3,opt,name=perPage,proto3" json:"perPage,omitempty"`
	Search  string `protobuf:"bytes,4,opt,name=search,proto3" json:"search,omitempty"`
	Status  string `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *CustomerFindAllRequest) Reset() {
	*x = CustomerFindAllRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_customer_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomerFindAllRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerFindAllRequest) ProtoMessage() {}

func (x *CustomerFindAllRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_customer_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerFindAllRequest.ProtoReflect.Descriptor instead.
func (*CustomerFindAllRequest) Descriptor() ([]byte, []int) {
	return file_pb_customer_proto_rawDescGZIP(), []int{5}
}

func (x *CustomerFindAllRequest) GetSort() string {
	if x != nil {
		return x.Sort
	}
	return ""
}

func (x *CustomerFindAllRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *CustomerFindAllRequest) GetPerPage() int64 {
	if x != nil {
		return x.PerPage
	}
	return 0
}

func (x *CustomerFindAllRequest) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

func (x *CustomerFindAllRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type CustomerFindAllResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Customers  []*Customer `protobuf:"bytes,1,rep,name=customers,proto3" json:"customers,omitempty"`
	Rows       int64       `protobuf:"varint,2,opt,name=rows,proto3" json:"rows,omitempty"`
	Pages      int64       `protobuf:"varint,3,opt,name=pages,proto3" json:"pages,omitempty"`
	PerPage    int64       `protobuf:"varint,4,opt,name=per_page,json=perPage,proto3" json:"per_page,omitempty"`
	ActivePage int64       `protobuf:"varint,5,opt,name=active_page,json=activePage,proto3" json:"active_page,omitempty"`
	Total      int64       `protobuf:"varint,6,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *CustomerFindAllResponse) Reset() {
	*x = CustomerFindAllResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_customer_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomerFindAllResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerFindAllResponse) ProtoMessage() {}

func (x *CustomerFindAllResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_customer_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerFindAllResponse.ProtoReflect.Descriptor instead.
func (*CustomerFindAllResponse) Descriptor() ([]byte, []int) {
	return file_pb_customer_proto_rawDescGZIP(), []int{6}
}

func (x *CustomerFindAllResponse) GetCustomers() []*Customer {
	if x != nil {
		return x.Customers
	}
	return nil
}

func (x *CustomerFindAllResponse) GetRows() int64 {
	if x != nil {
		return x.Rows
	}
	return 0
}

func (x *CustomerFindAllResponse) GetPages() int64 {
	if x != nil {
		return x.Pages
	}
	return 0
}

func (x *CustomerFindAllResponse) GetPerPage() int64 {
	if x != nil {
		return x.PerPage
	}
	return 0
}

func (x *CustomerFindAllResponse) GetActivePage() int64 {
	if x != nil {
		return x.ActivePage
	}
	return 0
}

func (x *CustomerFindAllResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type CustomerChangeStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	Status     string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *CustomerChangeStatusRequest) Reset() {
	*x = CustomerChangeStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_customer_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomerChangeStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerChangeStatusRequest) ProtoMessage() {}

func (x *CustomerChangeStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_customer_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerChangeStatusRequest.ProtoReflect.Descriptor instead.
func (*CustomerChangeStatusRequest) Descriptor() ([]byte, []int) {
	return file_pb_customer_proto_rawDescGZIP(), []int{7}
}

func (x *CustomerChangeStatusRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *CustomerChangeStatusRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type CustomerUpdateDetailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId string          `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	Detail     *CustomerDetail `protobuf:"bytes,2,opt,name=detail,proto3" json:"detail,omitempty"`
}

func (x *CustomerUpdateDetailRequest) Reset() {
	*x = CustomerUpdateDetailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_customer_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomerUpdateDetailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerUpdateDetailRequest) ProtoMessage() {}

func (x *CustomerUpdateDetailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_customer_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerUpdateDetailRequest.ProtoReflect.Descriptor instead.
func (*CustomerUpdateDetailRequest) Descriptor() ([]byte, []int) {
	return file_pb_customer_proto_rawDescGZIP(), []int{8}
}

func (x *CustomerUpdateDetailRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *CustomerUpdateDetailRequest) GetDetail() *CustomerDetail {
	if x != nil {
		return x.Detail
	}
	return nil
}

type CustomerDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	Hard       bool   `protobuf:"varint,2,opt,name=hard,proto3" json:"hard,omitempty"`
}

func (x *CustomerDeleteRequest) Reset() {
	*x = CustomerDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_customer_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomerDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerDeleteRequest) ProtoMessage() {}

func (x *CustomerDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_customer_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerDeleteRequest.ProtoReflect.Descriptor instead.
func (*CustomerDeleteRequest) Descriptor() ([]byte, []int) {
	return file_pb_customer_proto_rawDescGZIP(), []int{9}
}

func (x *CustomerDeleteRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *CustomerDeleteRequest) GetHard() bool {
	if x != nil {
		return x.Hard
	}
	return false
}

var File_pb_customer_proto protoreflect.FileDescriptor

var file_pb_customer_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x62, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x70, 0x62, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb3, 0x01, 0x0a, 0x10, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x69, 0x6c, 0x6c, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x69, 0x6c, 0x6c, 0x61, 0x67, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x70,
	0x6f, 0x73, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x64, 0x65, 0x22, 0xcd, 0x01, 0x0a,
	0x0e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x70, 0x73, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x70, 0x73, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x0e, 0x0a,
	0x02, 0x77, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x77, 0x61, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x62, 0x6f, 0x75, 0x74,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x62, 0x6f, 0x75, 0x74, 0x12, 0x2d, 0x0a,
	0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x68, 0x0a, 0x15,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x73,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x73, 0x73, 0x12, 0x27, 0x0a,
	0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x06,
	0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x22, 0x8e, 0x02, 0x0a, 0x08, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x73, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x73, 0x73, 0x12, 0x27, 0x0a, 0x06,
	0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x06, 0x64,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1b, 0x0a,
	0x09, 0x65, 0x78, 0x70, 0x5f, 0x75, 0x6e, 0x74, 0x69, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x65, 0x78, 0x70, 0x55, 0x6e, 0x74, 0x69, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x61, 0x0a, 0x16, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x73, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x73, 0x73, 0x22, 0x8a, 0x01, 0x0a, 0x16, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x70, 0x65, 0x72, 0x50, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07,
	0x70, 0x65, 0x72, 0x50, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xbe, 0x01, 0x0a, 0x17, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x09, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x52, 0x09, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x72, 0x6f, 0x77, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x72, 0x6f, 0x77, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x65, 0x72, 0x5f, 0x70, 0x61,
	0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x70, 0x65, 0x72, 0x50, 0x61, 0x67,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x70, 0x61, 0x67, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x50, 0x61,
	0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x56, 0x0a, 0x1b, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x22, 0x67, 0x0a, 0x1b, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x27, 0x0a, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x52, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x22, 0x4c, 0x0a, 0x15, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x04, 0x68, 0x61, 0x72, 0x64, 0x32, 0xee, 0x02, 0x0a, 0x0f, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2a, 0x0a, 0x06, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x16, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x06, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x07, 0x46, 0x69, 0x6e, 0x64, 0x4f,
	0x6e, 0x65, 0x12, 0x17, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x46, 0x69, 0x6e,
	0x64, 0x4f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x09, 0x2e, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x07, 0x46, 0x69, 0x6e, 0x64,
	0x41, 0x6c, 0x6c, 0x12, 0x17, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x46, 0x69,
	0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x0c, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x0c,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x1c, 0x2e, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x36, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x16, 0x2e, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x12, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_customer_proto_rawDescOnce sync.Once
	file_pb_customer_proto_rawDescData = file_pb_customer_proto_rawDesc
)

func file_pb_customer_proto_rawDescGZIP() []byte {
	file_pb_customer_proto_rawDescOnce.Do(func() {
		file_pb_customer_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_customer_proto_rawDescData)
	})
	return file_pb_customer_proto_rawDescData
}

var file_pb_customer_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_pb_customer_proto_goTypes = []interface{}{
	(*CustomerLocation)(nil),            // 0: CustomerLocation
	(*CustomerDetail)(nil),              // 1: CustomerDetail
	(*CustomerCreateRequest)(nil),       // 2: CustomerCreateRequest
	(*Customer)(nil),                    // 3: Customer
	(*CustomerFindOneRequest)(nil),      // 4: CustomerFindOneRequest
	(*CustomerFindAllRequest)(nil),      // 5: CustomerFindAllRequest
	(*CustomerFindAllResponse)(nil),     // 6: CustomerFindAllResponse
	(*CustomerChangeStatusRequest)(nil), // 7: CustomerChangeStatusRequest
	(*CustomerUpdateDetailRequest)(nil), // 8: CustomerUpdateDetailRequest
	(*CustomerDeleteRequest)(nil),       // 9: CustomerDeleteRequest
	(*Empty)(nil),                       // 10: Empty
	(*OperationResponse)(nil),           // 11: OperationResponse
}
var file_pb_customer_proto_depIdxs = []int32{
	0,  // 0: CustomerDetail.location:type_name -> CustomerLocation
	1,  // 1: CustomerCreateRequest.detail:type_name -> CustomerDetail
	1,  // 2: Customer.detail:type_name -> CustomerDetail
	3,  // 3: CustomerFindAllResponse.customers:type_name -> Customer
	1,  // 4: CustomerUpdateDetailRequest.detail:type_name -> CustomerDetail
	2,  // 5: CustomerService.Create:input_type -> CustomerCreateRequest
	4,  // 6: CustomerService.FindOne:input_type -> CustomerFindOneRequest
	5,  // 7: CustomerService.FindAll:input_type -> CustomerFindAllRequest
	7,  // 8: CustomerService.ChangeStatus:input_type -> CustomerChangeStatusRequest
	8,  // 9: CustomerService.UpdateDetail:input_type -> CustomerUpdateDetailRequest
	9,  // 10: CustomerService.Delete:input_type -> CustomerDeleteRequest
	10, // 11: CustomerService.Create:output_type -> Empty
	3,  // 12: CustomerService.FindOne:output_type -> Customer
	6,  // 13: CustomerService.FindAll:output_type -> CustomerFindAllResponse
	11, // 14: CustomerService.ChangeStatus:output_type -> OperationResponse
	11, // 15: CustomerService.UpdateDetail:output_type -> OperationResponse
	11, // 16: CustomerService.Delete:output_type -> OperationResponse
	11, // [11:17] is the sub-list for method output_type
	5,  // [5:11] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_pb_customer_proto_init() }
func file_pb_customer_proto_init() {
	if File_pb_customer_proto != nil {
		return
	}
	file_pb_response_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_pb_customer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomerLocation); i {
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
		file_pb_customer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomerDetail); i {
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
		file_pb_customer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomerCreateRequest); i {
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
		file_pb_customer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Customer); i {
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
		file_pb_customer_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomerFindOneRequest); i {
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
		file_pb_customer_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomerFindAllRequest); i {
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
		file_pb_customer_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomerFindAllResponse); i {
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
		file_pb_customer_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomerChangeStatusRequest); i {
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
		file_pb_customer_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomerUpdateDetailRequest); i {
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
		file_pb_customer_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomerDeleteRequest); i {
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
			RawDescriptor: file_pb_customer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_customer_proto_goTypes,
		DependencyIndexes: file_pb_customer_proto_depIdxs,
		MessageInfos:      file_pb_customer_proto_msgTypes,
	}.Build()
	File_pb_customer_proto = out.File
	file_pb_customer_proto_rawDesc = nil
	file_pb_customer_proto_goTypes = nil
	file_pb_customer_proto_depIdxs = nil
}