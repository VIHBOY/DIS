// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: chat.proto

package chat

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Orden struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Producto string `protobuf:"bytes,2,opt,name=producto,proto3" json:"producto,omitempty"`
	Valor    string `protobuf:"bytes,3,opt,name=valor,proto3" json:"valor,omitempty"`
	Inicio   string `protobuf:"bytes,4,opt,name=inicio,proto3" json:"inicio,omitempty"`
	Destino  string `protobuf:"bytes,5,opt,name=destino,proto3" json:"destino,omitempty"`
	Tipo     string `protobuf:"bytes,6,opt,name=tipo,proto3" json:"tipo,omitempty"`
}

func (x *Orden) Reset() {
	*x = Orden{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Orden) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Orden) ProtoMessage() {}

func (x *Orden) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Orden.ProtoReflect.Descriptor instead.
func (*Orden) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{0}
}

func (x *Orden) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Orden) GetProducto() string {
	if x != nil {
		return x.Producto
	}
	return ""
}

func (x *Orden) GetValor() string {
	if x != nil {
		return x.Valor
	}
	return ""
}

func (x *Orden) GetInicio() string {
	if x != nil {
		return x.Inicio
	}
	return ""
}

func (x *Orden) GetDestino() string {
	if x != nil {
		return x.Destino
	}
	return ""
}

func (x *Orden) GetTipo() string {
	if x != nil {
		return x.Tipo
	}
	return ""
}

type Paquete struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Track    string `protobuf:"bytes,2,opt,name=track,proto3" json:"track,omitempty"`
	Tipo     string `protobuf:"bytes,3,opt,name=tipo,proto3" json:"tipo,omitempty"`
	Valor    int32  `protobuf:"varint,4,opt,name=valor,proto3" json:"valor,omitempty"`
	Intentos int32  `protobuf:"varint,5,opt,name=intentos,proto3" json:"intentos,omitempty"`
	Estado   string `protobuf:"bytes,6,opt,name=estado,proto3" json:"estado,omitempty"`
}

func (x *Paquete) Reset() {
	*x = Paquete{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Paquete) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Paquete) ProtoMessage() {}

func (x *Paquete) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Paquete.ProtoReflect.Descriptor instead.
func (*Paquete) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{1}
}

func (x *Paquete) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Paquete) GetTrack() string {
	if x != nil {
		return x.Track
	}
	return ""
}

func (x *Paquete) GetTipo() string {
	if x != nil {
		return x.Tipo
	}
	return ""
}

func (x *Paquete) GetValor() int32 {
	if x != nil {
		return x.Valor
	}
	return 0
}

func (x *Paquete) GetIntentos() int32 {
	if x != nil {
		return x.Intentos
	}
	return 0
}

func (x *Paquete) GetEstado() string {
	if x != nil {
		return x.Estado
	}
	return ""
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Body string `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{2}
}

func (x *Message) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

var File_chat_proto protoreflect.FileDescriptor

var file_chat_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63, 0x68,
	0x61, 0x74, 0x22, 0x8f, 0x01, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x6e, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x6f,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x6f, 0x72, 0x12, 0x16,
	0x0a, 0x06, 0x69, 0x6e, 0x69, 0x63, 0x69, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x69, 0x6e, 0x69, 0x63, 0x69, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e,
	0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x6f,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x70, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x69, 0x70, 0x6f, 0x22, 0x8d, 0x01, 0x0a, 0x07, 0x50, 0x61, 0x71, 0x75, 0x65, 0x74, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x70, 0x6f, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x69, 0x70, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x6f, 0x72,
	0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x6f, 0x73, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x6f, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x65, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x73,
	0x74, 0x61, 0x64, 0x6f, 0x22, 0x1d, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62,
	0x6f, 0x64, 0x79, 0x32, 0x81, 0x03, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x2d, 0x0a, 0x0b, 0x4d, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x4f, 0x72, 0x64,
	0x65, 0x6e, 0x12, 0x0d, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x1a, 0x0d, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x00, 0x12, 0x2b, 0x0a, 0x09, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x74, 0x61, 0x72, 0x12,
	0x0d, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x0d,
	0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12,
	0x2c, 0x0a, 0x0c, 0x4d, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x6e, 0x32, 0x12,
	0x0b, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x6e, 0x1a, 0x0d, 0x2e, 0x63,
	0x68, 0x61, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x29, 0x0a,
	0x07, 0x52, 0x65, 0x63, 0x69, 0x62, 0x69, 0x72, 0x12, 0x0d, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x0d, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x2a, 0x0a, 0x08, 0x52, 0x65, 0x63, 0x69,
	0x62, 0x69, 0x72, 0x32, 0x12, 0x0d, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x1a, 0x0d, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x50, 0x61, 0x71, 0x75, 0x65,
	0x74, 0x65, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x0d, 0x43, 0x61, 0x6d, 0x62, 0x69, 0x61, 0x72, 0x45,
	0x73, 0x74, 0x61, 0x64, 0x6f, 0x12, 0x0d, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x1a, 0x0d, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x0f, 0x43, 0x61, 0x6d, 0x62, 0x69, 0x61, 0x72,
	0x49, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x6f, 0x73, 0x12, 0x0d, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x0d, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x2d, 0x0a, 0x0b, 0x42, 0x75, 0x73, 0x63,
	0x61, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x6e, 0x12, 0x0d, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x0d, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chat_proto_rawDescOnce sync.Once
	file_chat_proto_rawDescData = file_chat_proto_rawDesc
)

func file_chat_proto_rawDescGZIP() []byte {
	file_chat_proto_rawDescOnce.Do(func() {
		file_chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_chat_proto_rawDescData)
	})
	return file_chat_proto_rawDescData
}

var file_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_chat_proto_goTypes = []interface{}{
	(*Orden)(nil),   // 0: chat.Orden
	(*Paquete)(nil), // 1: chat.Paquete
	(*Message)(nil), // 2: chat.Message
}
var file_chat_proto_depIdxs = []int32{
	2, // 0: chat.ChatService.MandarOrden:input_type -> chat.Message
	2, // 1: chat.ChatService.Consultar:input_type -> chat.Message
	0, // 2: chat.ChatService.MandarOrden2:input_type -> chat.Orden
	2, // 3: chat.ChatService.Recibir:input_type -> chat.Message
	2, // 4: chat.ChatService.Recibir2:input_type -> chat.Message
	2, // 5: chat.ChatService.CambiarEstado:input_type -> chat.Message
	2, // 6: chat.ChatService.CambiarIntentos:input_type -> chat.Message
	2, // 7: chat.ChatService.BuscarOrden:input_type -> chat.Message
	2, // 8: chat.ChatService.MandarOrden:output_type -> chat.Message
	2, // 9: chat.ChatService.Consultar:output_type -> chat.Message
	2, // 10: chat.ChatService.MandarOrden2:output_type -> chat.Message
	2, // 11: chat.ChatService.Recibir:output_type -> chat.Message
	1, // 12: chat.ChatService.Recibir2:output_type -> chat.Paquete
	2, // 13: chat.ChatService.CambiarEstado:output_type -> chat.Message
	2, // 14: chat.ChatService.CambiarIntentos:output_type -> chat.Message
	2, // 15: chat.ChatService.BuscarOrden:output_type -> chat.Message
	8, // [8:16] is the sub-list for method output_type
	0, // [0:8] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_chat_proto_init() }
func file_chat_proto_init() {
	if File_chat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Orden); i {
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
		file_chat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Paquete); i {
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
		file_chat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
			RawDescriptor: file_chat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chat_proto_goTypes,
		DependencyIndexes: file_chat_proto_depIdxs,
		MessageInfos:      file_chat_proto_msgTypes,
	}.Build()
	File_chat_proto = out.File
	file_chat_proto_rawDesc = nil
	file_chat_proto_goTypes = nil
	file_chat_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ChatServiceClient is the client API for ChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChatServiceClient interface {
	MandarOrden(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	Consultar(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	MandarOrden2(ctx context.Context, in *Orden, opts ...grpc.CallOption) (*Message, error)
	Recibir(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	Recibir2(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Paquete, error)
	CambiarEstado(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	CambiarIntentos(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	BuscarOrden(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
}

type chatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChatServiceClient(cc grpc.ClientConnInterface) ChatServiceClient {
	return &chatServiceClient{cc}
}

func (c *chatServiceClient) MandarOrden(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/chat.ChatService/MandarOrden", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) Consultar(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/chat.ChatService/Consultar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) MandarOrden2(ctx context.Context, in *Orden, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/chat.ChatService/MandarOrden2", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) Recibir(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/chat.ChatService/Recibir", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) Recibir2(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Paquete, error) {
	out := new(Paquete)
	err := c.cc.Invoke(ctx, "/chat.ChatService/Recibir2", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) CambiarEstado(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/chat.ChatService/CambiarEstado", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) CambiarIntentos(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/chat.ChatService/CambiarIntentos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) BuscarOrden(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/chat.ChatService/BuscarOrden", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatServiceServer is the server API for ChatService service.
type ChatServiceServer interface {
	MandarOrden(context.Context, *Message) (*Message, error)
	Consultar(context.Context, *Message) (*Message, error)
	MandarOrden2(context.Context, *Orden) (*Message, error)
	Recibir(context.Context, *Message) (*Message, error)
	Recibir2(context.Context, *Message) (*Paquete, error)
	CambiarEstado(context.Context, *Message) (*Message, error)
	CambiarIntentos(context.Context, *Message) (*Message, error)
	BuscarOrden(context.Context, *Message) (*Message, error)
}

// UnimplementedChatServiceServer can be embedded to have forward compatible implementations.
type UnimplementedChatServiceServer struct {
}

func (*UnimplementedChatServiceServer) MandarOrden(context.Context, *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MandarOrden not implemented")
}
func (*UnimplementedChatServiceServer) Consultar(context.Context, *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Consultar not implemented")
}
func (*UnimplementedChatServiceServer) MandarOrden2(context.Context, *Orden) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MandarOrden2 not implemented")
}
func (*UnimplementedChatServiceServer) Recibir(context.Context, *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Recibir not implemented")
}
func (*UnimplementedChatServiceServer) Recibir2(context.Context, *Message) (*Paquete, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Recibir2 not implemented")
}
func (*UnimplementedChatServiceServer) CambiarEstado(context.Context, *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CambiarEstado not implemented")
}
func (*UnimplementedChatServiceServer) CambiarIntentos(context.Context, *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CambiarIntentos not implemented")
}
func (*UnimplementedChatServiceServer) BuscarOrden(context.Context, *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BuscarOrden not implemented")
}

func RegisterChatServiceServer(s *grpc.Server, srv ChatServiceServer) {
	s.RegisterService(&_ChatService_serviceDesc, srv)
}

func _ChatService_MandarOrden_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).MandarOrden(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatService/MandarOrden",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).MandarOrden(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_Consultar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).Consultar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatService/Consultar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).Consultar(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_MandarOrden2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Orden)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).MandarOrden2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatService/MandarOrden2",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).MandarOrden2(ctx, req.(*Orden))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_Recibir_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).Recibir(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatService/Recibir",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).Recibir(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_Recibir2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).Recibir2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatService/Recibir2",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).Recibir2(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_CambiarEstado_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).CambiarEstado(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatService/CambiarEstado",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).CambiarEstado(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_CambiarIntentos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).CambiarIntentos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatService/CambiarIntentos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).CambiarIntentos(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_BuscarOrden_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).BuscarOrden(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatService/BuscarOrden",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).BuscarOrden(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _ChatService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chat.ChatService",
	HandlerType: (*ChatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MandarOrden",
			Handler:    _ChatService_MandarOrden_Handler,
		},
		{
			MethodName: "Consultar",
			Handler:    _ChatService_Consultar_Handler,
		},
		{
			MethodName: "MandarOrden2",
			Handler:    _ChatService_MandarOrden2_Handler,
		},
		{
			MethodName: "Recibir",
			Handler:    _ChatService_Recibir_Handler,
		},
		{
			MethodName: "Recibir2",
			Handler:    _ChatService_Recibir2_Handler,
		},
		{
			MethodName: "CambiarEstado",
			Handler:    _ChatService_CambiarEstado_Handler,
		},
		{
			MethodName: "CambiarIntentos",
			Handler:    _ChatService_CambiarIntentos_Handler,
		},
		{
			MethodName: "BuscarOrden",
			Handler:    _ChatService_BuscarOrden_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chat.proto",
}
