// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sxpayment.proto

package sxpayment

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type RequestHeader struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestHeader) Reset()         { *m = RequestHeader{} }
func (m *RequestHeader) String() string { return proto.CompactTextString(m) }
func (*RequestHeader) ProtoMessage()    {}
func (*RequestHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_47b245f9b6c7f8c8, []int{0}
}

func (m *RequestHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestHeader.Unmarshal(m, b)
}
func (m *RequestHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestHeader.Marshal(b, m, deterministic)
}
func (m *RequestHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestHeader.Merge(m, src)
}
func (m *RequestHeader) XXX_Size() int {
	return xxx_messageInfo_RequestHeader.Size(m)
}
func (m *RequestHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestHeader.DiscardUnknown(m)
}

var xxx_messageInfo_RequestHeader proto.InternalMessageInfo

func (m *RequestHeader) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type RequestField struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	CustomerPayment      bool     `protobuf:"varint,3,opt,name=customerPayment,proto3" json:"customerPayment,omitempty"`
	CustomerID           int32    `protobuf:"varint,4,opt,name=customerID,proto3" json:"customerID,omitempty"`
	VendorID             int32    `protobuf:"varint,5,opt,name=vendorID,proto3" json:"vendorID,omitempty"`
	PaidAmount           float64  `protobuf:"fixed64,6,opt,name=paidAmount,proto3" json:"paidAmount,omitempty"`
	TransactionIDs       []int32  `protobuf:"varint,7,rep,packed,name=transactionIDs,proto3" json:"transactionIDs,omitempty"`
	PaymentType          string   `protobuf:"bytes,8,opt,name=paymentType,proto3" json:"paymentType,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestField) Reset()         { *m = RequestField{} }
func (m *RequestField) String() string { return proto.CompactTextString(m) }
func (*RequestField) ProtoMessage()    {}
func (*RequestField) Descriptor() ([]byte, []int) {
	return fileDescriptor_47b245f9b6c7f8c8, []int{1}
}

func (m *RequestField) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestField.Unmarshal(m, b)
}
func (m *RequestField) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestField.Marshal(b, m, deterministic)
}
func (m *RequestField) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestField.Merge(m, src)
}
func (m *RequestField) XXX_Size() int {
	return xxx_messageInfo_RequestField.Size(m)
}
func (m *RequestField) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestField.DiscardUnknown(m)
}

var xxx_messageInfo_RequestField proto.InternalMessageInfo

func (m *RequestField) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RequestField) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *RequestField) GetCustomerPayment() bool {
	if m != nil {
		return m.CustomerPayment
	}
	return false
}

func (m *RequestField) GetCustomerID() int32 {
	if m != nil {
		return m.CustomerID
	}
	return 0
}

func (m *RequestField) GetVendorID() int32 {
	if m != nil {
		return m.VendorID
	}
	return 0
}

func (m *RequestField) GetPaidAmount() float64 {
	if m != nil {
		return m.PaidAmount
	}
	return 0
}

func (m *RequestField) GetTransactionIDs() []int32 {
	if m != nil {
		return m.TransactionIDs
	}
	return nil
}

func (m *RequestField) GetPaymentType() string {
	if m != nil {
		return m.PaymentType
	}
	return ""
}

type RawRequest struct {
	Header               *RequestHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *RawRequest) Reset()         { *m = RawRequest{} }
func (m *RawRequest) String() string { return proto.CompactTextString(m) }
func (*RawRequest) ProtoMessage()    {}
func (*RawRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_47b245f9b6c7f8c8, []int{2}
}

func (m *RawRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RawRequest.Unmarshal(m, b)
}
func (m *RawRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RawRequest.Marshal(b, m, deterministic)
}
func (m *RawRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RawRequest.Merge(m, src)
}
func (m *RawRequest) XXX_Size() int {
	return xxx_messageInfo_RawRequest.Size(m)
}
func (m *RawRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RawRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RawRequest proto.InternalMessageInfo

func (m *RawRequest) GetHeader() *RequestHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type GetRequest struct {
	Header               *RequestHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Id                   int32          `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_47b245f9b6c7f8c8, []int{3}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetHeader() *RequestHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *GetRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type CreateRequest struct {
	Header               *RequestHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Field                *RequestField  `protobuf:"bytes,2,opt,name=field,proto3" json:"field,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_47b245f9b6c7f8c8, []int{4}
}

func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (m *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(m, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetHeader() *RequestHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *CreateRequest) GetField() *RequestField {
	if m != nil {
		return m.Field
	}
	return nil
}

type UpdateRequest struct {
	Header               *RequestHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Id                   int32          `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Action               string         `protobuf:"bytes,3,opt,name=action,proto3" json:"action,omitempty"`
	Field                *RequestField  `protobuf:"bytes,4,opt,name=field,proto3" json:"field,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *UpdateRequest) Reset()         { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()    {}
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_47b245f9b6c7f8c8, []int{5}
}

func (m *UpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRequest.Unmarshal(m, b)
}
func (m *UpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRequest.Marshal(b, m, deterministic)
}
func (m *UpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRequest.Merge(m, src)
}
func (m *UpdateRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateRequest.Size(m)
}
func (m *UpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRequest proto.InternalMessageInfo

func (m *UpdateRequest) GetHeader() *RequestHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *UpdateRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UpdateRequest) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *UpdateRequest) GetField() *RequestField {
	if m != nil {
		return m.Field
	}
	return nil
}

type RawResponse struct {
	StatusCode           int32             `protobuf:"varint,1,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
	Message              string            `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data                 *RawResponse_Data `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *RawResponse) Reset()         { *m = RawResponse{} }
func (m *RawResponse) String() string { return proto.CompactTextString(m) }
func (*RawResponse) ProtoMessage()    {}
func (*RawResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_47b245f9b6c7f8c8, []int{6}
}

func (m *RawResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RawResponse.Unmarshal(m, b)
}
func (m *RawResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RawResponse.Marshal(b, m, deterministic)
}
func (m *RawResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RawResponse.Merge(m, src)
}
func (m *RawResponse) XXX_Size() int {
	return xxx_messageInfo_RawResponse.Size(m)
}
func (m *RawResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RawResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RawResponse proto.InternalMessageInfo

func (m *RawResponse) GetStatusCode() int32 {
	if m != nil {
		return m.StatusCode
	}
	return 0
}

func (m *RawResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *RawResponse) GetData() *RawResponse_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

type RawResponse_Data struct {
	PaymentType          []string `protobuf:"bytes,1,rep,name=paymentType,proto3" json:"paymentType,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RawResponse_Data) Reset()         { *m = RawResponse_Data{} }
func (m *RawResponse_Data) String() string { return proto.CompactTextString(m) }
func (*RawResponse_Data) ProtoMessage()    {}
func (*RawResponse_Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_47b245f9b6c7f8c8, []int{6, 0}
}

func (m *RawResponse_Data) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RawResponse_Data.Unmarshal(m, b)
}
func (m *RawResponse_Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RawResponse_Data.Marshal(b, m, deterministic)
}
func (m *RawResponse_Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RawResponse_Data.Merge(m, src)
}
func (m *RawResponse_Data) XXX_Size() int {
	return xxx_messageInfo_RawResponse_Data.Size(m)
}
func (m *RawResponse_Data) XXX_DiscardUnknown() {
	xxx_messageInfo_RawResponse_Data.DiscardUnknown(m)
}

var xxx_messageInfo_RawResponse_Data proto.InternalMessageInfo

func (m *RawResponse_Data) GetPaymentType() []string {
	if m != nil {
		return m.PaymentType
	}
	return nil
}

type DataResponse struct {
	StatusCode           int32              `protobuf:"varint,1,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
	Message              string             `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data                 *DataResponse_Data `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *DataResponse) Reset()         { *m = DataResponse{} }
func (m *DataResponse) String() string { return proto.CompactTextString(m) }
func (*DataResponse) ProtoMessage()    {}
func (*DataResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_47b245f9b6c7f8c8, []int{7}
}

func (m *DataResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataResponse.Unmarshal(m, b)
}
func (m *DataResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataResponse.Marshal(b, m, deterministic)
}
func (m *DataResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataResponse.Merge(m, src)
}
func (m *DataResponse) XXX_Size() int {
	return xxx_messageInfo_DataResponse.Size(m)
}
func (m *DataResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DataResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DataResponse proto.InternalMessageInfo

func (m *DataResponse) GetStatusCode() int32 {
	if m != nil {
		return m.StatusCode
	}
	return 0
}

func (m *DataResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *DataResponse) GetData() *DataResponse_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

type DataResponse_Data struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	CustomerPayment      bool     `protobuf:"varint,4,opt,name=customerPayment,proto3" json:"customerPayment,omitempty"`
	CustomerID           int32    `protobuf:"varint,5,opt,name=customerID,proto3" json:"customerID,omitempty"`
	Customer             string   `protobuf:"bytes,6,opt,name=customer,proto3" json:"customer,omitempty"`
	VendorID             int32    `protobuf:"varint,7,opt,name=vendorID,proto3" json:"vendorID,omitempty"`
	Vendor               string   `protobuf:"bytes,8,opt,name=vendor,proto3" json:"vendor,omitempty"`
	PaidAmount           float64  `protobuf:"fixed64,9,opt,name=paidAmount,proto3" json:"paidAmount,omitempty"`
	PaymentType          string   `protobuf:"bytes,10,opt,name=paymentType,proto3" json:"paymentType,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataResponse_Data) Reset()         { *m = DataResponse_Data{} }
func (m *DataResponse_Data) String() string { return proto.CompactTextString(m) }
func (*DataResponse_Data) ProtoMessage()    {}
func (*DataResponse_Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_47b245f9b6c7f8c8, []int{7, 0}
}

func (m *DataResponse_Data) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataResponse_Data.Unmarshal(m, b)
}
func (m *DataResponse_Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataResponse_Data.Marshal(b, m, deterministic)
}
func (m *DataResponse_Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataResponse_Data.Merge(m, src)
}
func (m *DataResponse_Data) XXX_Size() int {
	return xxx_messageInfo_DataResponse_Data.Size(m)
}
func (m *DataResponse_Data) XXX_DiscardUnknown() {
	xxx_messageInfo_DataResponse_Data.DiscardUnknown(m)
}

var xxx_messageInfo_DataResponse_Data proto.InternalMessageInfo

func (m *DataResponse_Data) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DataResponse_Data) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DataResponse_Data) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *DataResponse_Data) GetCustomerPayment() bool {
	if m != nil {
		return m.CustomerPayment
	}
	return false
}

func (m *DataResponse_Data) GetCustomerID() int32 {
	if m != nil {
		return m.CustomerID
	}
	return 0
}

func (m *DataResponse_Data) GetCustomer() string {
	if m != nil {
		return m.Customer
	}
	return ""
}

func (m *DataResponse_Data) GetVendorID() int32 {
	if m != nil {
		return m.VendorID
	}
	return 0
}

func (m *DataResponse_Data) GetVendor() string {
	if m != nil {
		return m.Vendor
	}
	return ""
}

func (m *DataResponse_Data) GetPaidAmount() float64 {
	if m != nil {
		return m.PaidAmount
	}
	return 0
}

func (m *DataResponse_Data) GetPaymentType() string {
	if m != nil {
		return m.PaymentType
	}
	return ""
}

func init() {
	proto.RegisterType((*RequestHeader)(nil), "sxpayment.RequestHeader")
	proto.RegisterType((*RequestField)(nil), "sxpayment.RequestField")
	proto.RegisterType((*RawRequest)(nil), "sxpayment.RawRequest")
	proto.RegisterType((*GetRequest)(nil), "sxpayment.GetRequest")
	proto.RegisterType((*CreateRequest)(nil), "sxpayment.CreateRequest")
	proto.RegisterType((*UpdateRequest)(nil), "sxpayment.UpdateRequest")
	proto.RegisterType((*RawResponse)(nil), "sxpayment.RawResponse")
	proto.RegisterType((*RawResponse_Data)(nil), "sxpayment.RawResponse.Data")
	proto.RegisterType((*DataResponse)(nil), "sxpayment.DataResponse")
	proto.RegisterType((*DataResponse_Data)(nil), "sxpayment.DataResponse.Data")
}

func init() { proto.RegisterFile("sxpayment.proto", fileDescriptor_47b245f9b6c7f8c8) }

var fileDescriptor_47b245f9b6c7f8c8 = []byte{
	// 559 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0xdd, 0x8a, 0x13, 0x31,
	0x14, 0x26, 0xf3, 0xb7, 0xed, 0x99, 0xed, 0x2e, 0x04, 0xed, 0x0e, 0x55, 0x64, 0x18, 0x50, 0xe6,
	0xc6, 0xba, 0x54, 0xbd, 0x5a, 0x10, 0xa4, 0xc5, 0x75, 0x6f, 0x44, 0x82, 0x3e, 0x40, 0x6c, 0xa2,
	0x0e, 0x3a, 0x3f, 0x4e, 0x52, 0xd7, 0x7d, 0x0f, 0x2f, 0xbc, 0x16, 0x04, 0x5f, 0xcf, 0x37, 0x90,
	0x49, 0x32, 0x6d, 0x66, 0xba, 0x85, 0xb2, 0x7b, 0x37, 0xe7, 0x4b, 0x4e, 0xce, 0x97, 0x73, 0xbe,
	0x2f, 0x03, 0xc7, 0xe2, 0x47, 0x45, 0xaf, 0x72, 0x5e, 0xc8, 0x69, 0x55, 0x97, 0xb2, 0xc4, 0xc3,
	0x35, 0x90, 0x3c, 0x84, 0x11, 0xe1, 0xdf, 0x56, 0x5c, 0xc8, 0xd7, 0x9c, 0x32, 0x5e, 0xe3, 0x3b,
	0xe0, 0xcb, 0xf2, 0x0b, 0x2f, 0x22, 0x14, 0xa3, 0x74, 0x48, 0x74, 0x90, 0xfc, 0x74, 0xe0, 0xd0,
	0xec, 0x7b, 0x95, 0xf1, 0xaf, 0x0c, 0x63, 0xf0, 0x0a, 0x9a, 0x73, 0xb3, 0x4b, 0x7d, 0xe3, 0x18,
	0x42, 0xc6, 0xc5, 0xb2, 0xce, 0x2a, 0x99, 0x95, 0x45, 0xe4, 0xa8, 0x25, 0x1b, 0xc2, 0x29, 0x1c,
	0x2f, 0x57, 0x42, 0x96, 0x39, 0xaf, 0xdf, 0x6a, 0x02, 0x91, 0x1b, 0xa3, 0x74, 0x40, 0xfa, 0x30,
	0x7e, 0x00, 0xd0, 0x42, 0x17, 0x8b, 0xc8, 0x8b, 0x51, 0xea, 0x13, 0x0b, 0xc1, 0x13, 0x18, 0x7c,
	0xe7, 0x05, 0x2b, 0x9b, 0x55, 0x5f, 0xad, 0xae, 0xe3, 0x26, 0xb7, 0xa2, 0x19, 0x7b, 0x99, 0x97,
	0xab, 0x42, 0x46, 0x41, 0x8c, 0x52, 0x44, 0x2c, 0x04, 0x3f, 0x82, 0x23, 0x59, 0xd3, 0x42, 0xd0,
	0x65, 0x43, 0xea, 0x62, 0x21, 0xa2, 0x83, 0xd8, 0x4d, 0x7d, 0xd2, 0x43, 0x9b, 0xfb, 0x98, 0x36,
	0xbd, 0xbb, 0xaa, 0x78, 0x34, 0xd0, 0xf7, 0xb1, 0xa0, 0xe4, 0x05, 0x00, 0xa1, 0x97, 0xa6, 0x31,
	0xf8, 0x14, 0x82, 0xcf, 0xaa, 0x89, 0xaa, 0x2b, 0xe1, 0x2c, 0x9a, 0x6e, 0x1a, 0xdf, 0x69, 0x32,
	0x31, 0xfb, 0x92, 0x37, 0x00, 0xe7, 0x5c, 0xde, 0x38, 0x1f, 0x1f, 0x81, 0x93, 0x31, 0xd5, 0x68,
	0x9f, 0x38, 0x19, 0x4b, 0x2a, 0x18, 0xcd, 0x6b, 0x4e, 0x25, 0xbf, 0xf9, 0x91, 0x8f, 0xc1, 0xff,
	0xd8, 0x4c, 0x58, 0x9d, 0x1a, 0xce, 0x4e, 0xb6, 0x13, 0x94, 0x00, 0x88, 0xde, 0x95, 0xfc, 0x42,
	0x30, 0x7a, 0x5f, 0xb1, 0x5b, 0x95, 0xec, 0xdd, 0x02, 0x8f, 0x21, 0xd0, 0x43, 0x50, 0xe2, 0x18,
	0x12, 0x13, 0x6d, 0xa8, 0x79, 0x7b, 0x51, 0xfb, 0x8d, 0x20, 0x54, 0xd3, 0x11, 0x55, 0x59, 0x08,
	0xde, 0xc8, 0x42, 0x48, 0x2a, 0x57, 0x62, 0x5e, 0x32, 0x2d, 0x5c, 0x9f, 0x58, 0x08, 0x8e, 0xe0,
	0x20, 0xe7, 0x42, 0xd0, 0x4f, 0xdc, 0x48, 0xb7, 0x0d, 0xf1, 0x13, 0xf0, 0x18, 0x95, 0x54, 0xd1,
	0x09, 0x67, 0xf7, 0xec, 0xba, 0x9b, 0xf3, 0xa7, 0x0b, 0x2a, 0x29, 0x51, 0x1b, 0x27, 0x29, 0x78,
	0x4d, 0xd4, 0x57, 0x10, 0x8a, 0xdd, 0xbe, 0x82, 0xfe, 0xb8, 0x70, 0xa8, 0x12, 0x6f, 0xcf, 0xf2,
	0xb4, 0xc3, 0xf2, 0xbe, 0xc5, 0xd2, 0x2e, 0x60, 0xd3, 0xfc, 0xeb, 0x18, 0x9e, 0x7a, 0x02, 0xda,
	0xcb, 0xcd, 0x04, 0x5a, 0x77, 0x3b, 0xbb, 0xdd, 0xed, 0xee, 0xe5, 0x6e, 0x6f, 0x1f, 0x77, 0xfb,
	0xd7, 0xb9, 0xbb, 0x8d, 0x94, 0x7f, 0x87, 0x64, 0x1d, 0x77, 0x9c, 0x7f, 0xd0, 0x73, 0xfe, 0x18,
	0x02, 0xfd, 0x6d, 0xcc, 0x6a, 0xa2, 0xde, 0x8b, 0x30, 0xdc, 0x7a, 0x11, 0x7a, 0x73, 0x82, 0x2d,
	0xa7, 0xcf, 0xfe, 0x21, 0x18, 0xcc, 0x5b, 0x0a, 0xcf, 0xc0, 0x25, 0xf4, 0x12, 0xdf, 0xed, 0x0b,
	0x41, 0x69, 0x70, 0x32, 0xbe, 0x5e, 0x1f, 0xf8, 0x0c, 0x02, 0x6d, 0x4e, 0x6c, 0x5b, 0xa2, 0xe3,
	0xd7, 0xc9, 0xc9, 0x8e, 0xa9, 0xe1, 0xe7, 0xe0, 0x9e, 0x73, 0xd9, 0x29, 0xb9, 0x79, 0x39, 0x76,
	0xa7, 0x9d, 0x41, 0xa0, 0xdd, 0xd9, 0xa9, 0xd9, 0x31, 0xec, 0xce, 0xe4, 0x0f, 0x81, 0xfa, 0x5b,
	0x3c, 0xfd, 0x1f, 0x00, 0x00, 0xff, 0xff, 0xe0, 0x75, 0xbe, 0x40, 0x40, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CustomerClient is the client API for Customer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CustomerClient interface {
	Raw(ctx context.Context, in *RawRequest, opts ...grpc.CallOption) (*RawResponse, error)
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*DataResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*DataResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*DataResponse, error)
}

type customerClient struct {
	cc *grpc.ClientConn
}

func NewCustomerClient(cc *grpc.ClientConn) CustomerClient {
	return &customerClient{cc}
}

func (c *customerClient) Raw(ctx context.Context, in *RawRequest, opts ...grpc.CallOption) (*RawResponse, error) {
	out := new(RawResponse)
	err := c.cc.Invoke(ctx, "/sxpayment.Customer/Raw", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*DataResponse, error) {
	out := new(DataResponse)
	err := c.cc.Invoke(ctx, "/sxpayment.Customer/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*DataResponse, error) {
	out := new(DataResponse)
	err := c.cc.Invoke(ctx, "/sxpayment.Customer/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*DataResponse, error) {
	out := new(DataResponse)
	err := c.cc.Invoke(ctx, "/sxpayment.Customer/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomerServer is the server API for Customer service.
type CustomerServer interface {
	Raw(context.Context, *RawRequest) (*RawResponse, error)
	Create(context.Context, *CreateRequest) (*DataResponse, error)
	Get(context.Context, *GetRequest) (*DataResponse, error)
	Update(context.Context, *UpdateRequest) (*DataResponse, error)
}

// UnimplementedCustomerServer can be embedded to have forward compatible implementations.
type UnimplementedCustomerServer struct {
}

func (*UnimplementedCustomerServer) Raw(ctx context.Context, req *RawRequest) (*RawResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Raw not implemented")
}
func (*UnimplementedCustomerServer) Create(ctx context.Context, req *CreateRequest) (*DataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedCustomerServer) Get(ctx context.Context, req *GetRequest) (*DataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedCustomerServer) Update(ctx context.Context, req *UpdateRequest) (*DataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}

func RegisterCustomerServer(s *grpc.Server, srv CustomerServer) {
	s.RegisterService(&_Customer_serviceDesc, srv)
}

func _Customer_Raw_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RawRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServer).Raw(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sxpayment.Customer/Raw",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServer).Raw(ctx, req.(*RawRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Customer_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sxpayment.Customer/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Customer_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sxpayment.Customer/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Customer_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sxpayment.Customer/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Customer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sxpayment.Customer",
	HandlerType: (*CustomerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Raw",
			Handler:    _Customer_Raw_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Customer_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Customer_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Customer_Update_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sxpayment.proto",
}
