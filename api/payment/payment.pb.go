// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: api/payment/payment.proto

package payment

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

type ChargeAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StudentID string `protobuf:"bytes,1,opt,name=studentID,proto3" json:"studentID,omitempty"`
	Code      string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *ChargeAccountRequest) Reset() {
	*x = ChargeAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_payment_payment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChargeAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChargeAccountRequest) ProtoMessage() {}

func (x *ChargeAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_payment_payment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChargeAccountRequest.ProtoReflect.Descriptor instead.
func (*ChargeAccountRequest) Descriptor() ([]byte, []int) {
	return file_api_payment_payment_proto_rawDescGZIP(), []int{0}
}

func (x *ChargeAccountRequest) GetStudentID() string {
	if x != nil {
		return x.StudentID
	}
	return ""
}

func (x *ChargeAccountRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type AppResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *AppResponse) Reset() {
	*x = AppResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_payment_payment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppResponse) ProtoMessage() {}

func (x *AppResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_payment_payment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppResponse.ProtoReflect.Descriptor instead.
func (*AppResponse) Descriptor() ([]byte, []int) {
	return file_api_payment_payment_proto_rawDescGZIP(), []int{1}
}

func (x *AppResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GeneratePaymentCodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount float32 `protobuf:"fixed32,1,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *GeneratePaymentCodeRequest) Reset() {
	*x = GeneratePaymentCodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_payment_payment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeneratePaymentCodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeneratePaymentCodeRequest) ProtoMessage() {}

func (x *GeneratePaymentCodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_payment_payment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeneratePaymentCodeRequest.ProtoReflect.Descriptor instead.
func (*GeneratePaymentCodeRequest) Descriptor() ([]byte, []int) {
	return file_api_payment_payment_proto_rawDescGZIP(), []int{2}
}

func (x *GeneratePaymentCodeRequest) GetAmount() float32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type GetAmountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StudentID string `protobuf:"bytes,1,opt,name=studentID,proto3" json:"studentID,omitempty"`
}

func (x *GetAmountRequest) Reset() {
	*x = GetAmountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_payment_payment_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAmountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAmountRequest) ProtoMessage() {}

func (x *GetAmountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_payment_payment_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAmountRequest.ProtoReflect.Descriptor instead.
func (*GetAmountRequest) Descriptor() ([]byte, []int) {
	return file_api_payment_payment_proto_rawDescGZIP(), []int{3}
}

func (x *GetAmountRequest) GetStudentID() string {
	if x != nil {
		return x.StudentID
	}
	return ""
}

type GetAmountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount float32 `protobuf:"fixed32,1,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *GetAmountResponse) Reset() {
	*x = GetAmountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_payment_payment_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAmountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAmountResponse) ProtoMessage() {}

func (x *GetAmountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_payment_payment_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAmountResponse.ProtoReflect.Descriptor instead.
func (*GetAmountResponse) Descriptor() ([]byte, []int) {
	return file_api_payment_payment_proto_rawDescGZIP(), []int{4}
}

func (x *GetAmountResponse) GetAmount() float32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type PayClassRoomRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClassRoomID string `protobuf:"bytes,1,opt,name=classRoomID,proto3" json:"classRoomID,omitempty"`
	StudentID   string `protobuf:"bytes,2,opt,name=studentID,proto3" json:"studentID,omitempty"`
	Month       int32  `protobuf:"varint,3,opt,name=month,proto3" json:"month,omitempty"`
}

func (x *PayClassRoomRequest) Reset() {
	*x = PayClassRoomRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_payment_payment_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayClassRoomRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayClassRoomRequest) ProtoMessage() {}

func (x *PayClassRoomRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_payment_payment_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayClassRoomRequest.ProtoReflect.Descriptor instead.
func (*PayClassRoomRequest) Descriptor() ([]byte, []int) {
	return file_api_payment_payment_proto_rawDescGZIP(), []int{5}
}

func (x *PayClassRoomRequest) GetClassRoomID() string {
	if x != nil {
		return x.ClassRoomID
	}
	return ""
}

func (x *PayClassRoomRequest) GetStudentID() string {
	if x != nil {
		return x.StudentID
	}
	return ""
}

func (x *PayClassRoomRequest) GetMonth() int32 {
	if x != nil {
		return x.Month
	}
	return 0
}

var File_api_payment_payment_proto protoreflect.FileDescriptor

var file_api_payment_payment_proto_rawDesc = []byte{
	0x0a, 0x19, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x70, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x48, 0x0a, 0x14, 0x43,
	0x68, 0x61, 0x72, 0x67, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x49,
	0x44, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x27, 0x0a, 0x0b, 0x41, 0x70, 0x70, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x34,
	0x0a, 0x1a, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x22, 0x30, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x75, 0x64,
	0x65, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x75,
	0x64, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x22, 0x2b, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x22, 0x6b, 0x0a, 0x13, 0x50, 0x61, 0x79, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x52,
	0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6c,
	0x61, 0x73, 0x73, 0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09,
	0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f,
	0x6e, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68,
	0x32, 0xb4, 0x02, 0x0a, 0x0e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x0d, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x15, 0x2e, 0x43, 0x68, 0x61, 0x72, 0x67, 0x65, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x41, 0x70,
	0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x13, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x1b, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x50, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0c, 0x2e, 0x41, 0x70, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x34, 0x0a, 0x0c, 0x70, 0x61, 0x79, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x6f, 0x6f, 0x6d, 0x12,
	0x14, 0x2e, 0x50, 0x61, 0x79, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x41, 0x70, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x09, 0x67, 0x65, 0x74, 0x41, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x11, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x0f, 0x67,
	0x65, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x11,
	0x2e, 0x47, 0x65, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x12, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x48, 0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x61,
	0x72, 0x69, 0x6e, 0x69, 0x2e, 0x70, 0x61, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x61, 0x70, 0x69,
	0x50, 0x01, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c,
	0x65, 0x73, 0x2d, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_payment_payment_proto_rawDescOnce sync.Once
	file_api_payment_payment_proto_rawDescData = file_api_payment_payment_proto_rawDesc
)

func file_api_payment_payment_proto_rawDescGZIP() []byte {
	file_api_payment_payment_proto_rawDescOnce.Do(func() {
		file_api_payment_payment_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_payment_payment_proto_rawDescData)
	})
	return file_api_payment_payment_proto_rawDescData
}

var file_api_payment_payment_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_payment_payment_proto_goTypes = []interface{}{
	(*ChargeAccountRequest)(nil),       // 0: ChargeAccountRequest
	(*AppResponse)(nil),                // 1: AppResponse
	(*GeneratePaymentCodeRequest)(nil), // 2: GeneratePaymentCodeRequest
	(*GetAmountRequest)(nil),           // 3: GetAmountRequest
	(*GetAmountResponse)(nil),          // 4: GetAmountResponse
	(*PayClassRoomRequest)(nil),        // 5: PayClassRoomRequest
}
var file_api_payment_payment_proto_depIdxs = []int32{
	0, // 0: PaymentService.chargeAccount:input_type -> ChargeAccountRequest
	2, // 1: PaymentService.generatePaymentCode:input_type -> GeneratePaymentCodeRequest
	5, // 2: PaymentService.payClassRoom:input_type -> PayClassRoomRequest
	3, // 3: PaymentService.getAmount:input_type -> GetAmountRequest
	3, // 4: PaymentService.getAmountByCode:input_type -> GetAmountRequest
	1, // 5: PaymentService.chargeAccount:output_type -> AppResponse
	1, // 6: PaymentService.generatePaymentCode:output_type -> AppResponse
	1, // 7: PaymentService.payClassRoom:output_type -> AppResponse
	4, // 8: PaymentService.getAmount:output_type -> GetAmountResponse
	4, // 9: PaymentService.getAmountByCode:output_type -> GetAmountResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_payment_payment_proto_init() }
func file_api_payment_payment_proto_init() {
	if File_api_payment_payment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_payment_payment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChargeAccountRequest); i {
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
		file_api_payment_payment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppResponse); i {
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
		file_api_payment_payment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GeneratePaymentCodeRequest); i {
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
		file_api_payment_payment_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAmountRequest); i {
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
		file_api_payment_payment_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAmountResponse); i {
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
		file_api_payment_payment_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PayClassRoomRequest); i {
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
			RawDescriptor: file_api_payment_payment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_payment_payment_proto_goTypes,
		DependencyIndexes: file_api_payment_payment_proto_depIdxs,
		MessageInfos:      file_api_payment_payment_proto_msgTypes,
	}.Build()
	File_api_payment_payment_proto = out.File
	file_api_payment_payment_proto_rawDesc = nil
	file_api_payment_payment_proto_goTypes = nil
	file_api_payment_payment_proto_depIdxs = nil
}
