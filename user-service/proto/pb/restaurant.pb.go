// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        v5.29.3
// source: proto/restaurant.proto

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

type CreateRestaurantRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"` // User ID from the User Service
	Email         string                 `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`                 // Email of the user
	Name          string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`                   // Restaurant name
	Address       string                 `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`             // Restaurant address
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateRestaurantRequest) Reset() {
	*x = CreateRestaurantRequest{}
	mi := &file_proto_restaurant_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateRestaurantRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRestaurantRequest) ProtoMessage() {}

func (x *CreateRestaurantRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_restaurant_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRestaurantRequest.ProtoReflect.Descriptor instead.
func (*CreateRestaurantRequest) Descriptor() ([]byte, []int) {
	return file_proto_restaurant_proto_rawDescGZIP(), []int{0}
}

func (x *CreateRestaurantRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreateRestaurantRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateRestaurantRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateRestaurantRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type CreateRestaurantResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"` // Indicates if the operation was successful
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`  // Additional information or error details
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateRestaurantResponse) Reset() {
	*x = CreateRestaurantResponse{}
	mi := &file_proto_restaurant_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateRestaurantResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRestaurantResponse) ProtoMessage() {}

func (x *CreateRestaurantResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_restaurant_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRestaurantResponse.ProtoReflect.Descriptor instead.
func (*CreateRestaurantResponse) Descriptor() ([]byte, []int) {
	return file_proto_restaurant_proto_rawDescGZIP(), []int{1}
}

func (x *CreateRestaurantResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *CreateRestaurantResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_proto_restaurant_proto protoreflect.FileDescriptor

var file_proto_restaurant_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61,
	0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75,
	0x72, 0x61, 0x6e, 0x74, 0x22, 0x76, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x4e, 0x0a, 0x18,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x72, 0x0a, 0x11,
	0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x5d, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x74, 0x61,
	0x75, 0x72, 0x61, 0x6e, 0x74, 0x12, 0x23, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61,
	0x6e, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72,
	0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x72, 0x65, 0x73,
	0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x08, 0x5a, 0x06, 0x70, 0x62, 0x2f, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_restaurant_proto_rawDescOnce sync.Once
	file_proto_restaurant_proto_rawDescData = file_proto_restaurant_proto_rawDesc
)

func file_proto_restaurant_proto_rawDescGZIP() []byte {
	file_proto_restaurant_proto_rawDescOnce.Do(func() {
		file_proto_restaurant_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_restaurant_proto_rawDescData)
	})
	return file_proto_restaurant_proto_rawDescData
}

var file_proto_restaurant_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_restaurant_proto_goTypes = []any{
	(*CreateRestaurantRequest)(nil),  // 0: restaurant.CreateRestaurantRequest
	(*CreateRestaurantResponse)(nil), // 1: restaurant.CreateRestaurantResponse
}
var file_proto_restaurant_proto_depIdxs = []int32{
	0, // 0: restaurant.RestaurantService.CreateRestaurant:input_type -> restaurant.CreateRestaurantRequest
	1, // 1: restaurant.RestaurantService.CreateRestaurant:output_type -> restaurant.CreateRestaurantResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_restaurant_proto_init() }
func file_proto_restaurant_proto_init() {
	if File_proto_restaurant_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_restaurant_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_restaurant_proto_goTypes,
		DependencyIndexes: file_proto_restaurant_proto_depIdxs,
		MessageInfos:      file_proto_restaurant_proto_msgTypes,
	}.Build()
	File_proto_restaurant_proto = out.File
	file_proto_restaurant_proto_rawDesc = nil
	file_proto_restaurant_proto_goTypes = nil
	file_proto_restaurant_proto_depIdxs = nil
}
