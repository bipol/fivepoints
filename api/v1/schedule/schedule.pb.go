// Code generated by protoc-gen-go. DO NOT EDIT.
// source: schedule.proto

package schedule

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type GetArrivalEstimatesRequest struct {
	StartDate            *timestamp.Timestamp `protobuf:"bytes,1,opt,name=StartDate,proto3" json:"StartDate,omitempty"`
	EndDate              *timestamp.Timestamp `protobuf:"bytes,2,opt,name=EndDate,proto3" json:"EndDate,omitempty"`
	Station              string               `protobuf:"bytes,3,opt,name=Station,proto3" json:"Station,omitempty"`
	Destination          string               `protobuf:"bytes,4,opt,name=Destination,proto3" json:"Destination,omitempty"`
	LastEvaluatedKey     string               `protobuf:"bytes,5,opt,name=LastEvaluatedKey,proto3" json:"LastEvaluatedKey,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *GetArrivalEstimatesRequest) Reset()         { *m = GetArrivalEstimatesRequest{} }
func (m *GetArrivalEstimatesRequest) String() string { return proto.CompactTextString(m) }
func (*GetArrivalEstimatesRequest) ProtoMessage()    {}
func (*GetArrivalEstimatesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d00842e68e05382a, []int{0}
}

func (m *GetArrivalEstimatesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetArrivalEstimatesRequest.Unmarshal(m, b)
}
func (m *GetArrivalEstimatesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetArrivalEstimatesRequest.Marshal(b, m, deterministic)
}
func (m *GetArrivalEstimatesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetArrivalEstimatesRequest.Merge(m, src)
}
func (m *GetArrivalEstimatesRequest) XXX_Size() int {
	return xxx_messageInfo_GetArrivalEstimatesRequest.Size(m)
}
func (m *GetArrivalEstimatesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetArrivalEstimatesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetArrivalEstimatesRequest proto.InternalMessageInfo

func (m *GetArrivalEstimatesRequest) GetStartDate() *timestamp.Timestamp {
	if m != nil {
		return m.StartDate
	}
	return nil
}

func (m *GetArrivalEstimatesRequest) GetEndDate() *timestamp.Timestamp {
	if m != nil {
		return m.EndDate
	}
	return nil
}

func (m *GetArrivalEstimatesRequest) GetStation() string {
	if m != nil {
		return m.Station
	}
	return ""
}

func (m *GetArrivalEstimatesRequest) GetDestination() string {
	if m != nil {
		return m.Destination
	}
	return ""
}

func (m *GetArrivalEstimatesRequest) GetLastEvaluatedKey() string {
	if m != nil {
		return m.LastEvaluatedKey
	}
	return ""
}

type ArrivalEstimate struct {
	PrimaryKey           string   `protobuf:"bytes,1,opt,name=PrimaryKey,proto3" json:"PrimaryKey,omitempty"`
	SortKey              string   `protobuf:"bytes,2,opt,name=SortKey,proto3" json:"SortKey,omitempty"`
	Destination          string   `protobuf:"bytes,3,opt,name=Destination,proto3" json:"Destination,omitempty"`
	Direction            string   `protobuf:"bytes,4,opt,name=Direction,proto3" json:"Direction,omitempty"`
	EventTime            string   `protobuf:"bytes,5,opt,name=EventTime,proto3" json:"EventTime,omitempty"`
	Line                 string   `protobuf:"bytes,6,opt,name=Line,proto3" json:"Line,omitempty"`
	NextArrival          string   `protobuf:"bytes,7,opt,name=NextArrival,proto3" json:"NextArrival,omitempty"`
	Station              string   `protobuf:"bytes,8,opt,name=Station,proto3" json:"Station,omitempty"`
	TrainID              string   `protobuf:"bytes,9,opt,name=TrainID,proto3" json:"TrainID,omitempty"`
	WaitingSeconds       string   `protobuf:"bytes,10,opt,name=WaitingSeconds,proto3" json:"WaitingSeconds,omitempty"`
	WaitingTime          string   `protobuf:"bytes,11,opt,name=WaitingTime,proto3" json:"WaitingTime,omitempty"`
	TTL                  int64    `protobuf:"varint,12,opt,name=TTL,proto3" json:"TTL,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ArrivalEstimate) Reset()         { *m = ArrivalEstimate{} }
func (m *ArrivalEstimate) String() string { return proto.CompactTextString(m) }
func (*ArrivalEstimate) ProtoMessage()    {}
func (*ArrivalEstimate) Descriptor() ([]byte, []int) {
	return fileDescriptor_d00842e68e05382a, []int{1}
}

func (m *ArrivalEstimate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArrivalEstimate.Unmarshal(m, b)
}
func (m *ArrivalEstimate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArrivalEstimate.Marshal(b, m, deterministic)
}
func (m *ArrivalEstimate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArrivalEstimate.Merge(m, src)
}
func (m *ArrivalEstimate) XXX_Size() int {
	return xxx_messageInfo_ArrivalEstimate.Size(m)
}
func (m *ArrivalEstimate) XXX_DiscardUnknown() {
	xxx_messageInfo_ArrivalEstimate.DiscardUnknown(m)
}

var xxx_messageInfo_ArrivalEstimate proto.InternalMessageInfo

func (m *ArrivalEstimate) GetPrimaryKey() string {
	if m != nil {
		return m.PrimaryKey
	}
	return ""
}

func (m *ArrivalEstimate) GetSortKey() string {
	if m != nil {
		return m.SortKey
	}
	return ""
}

func (m *ArrivalEstimate) GetDestination() string {
	if m != nil {
		return m.Destination
	}
	return ""
}

func (m *ArrivalEstimate) GetDirection() string {
	if m != nil {
		return m.Direction
	}
	return ""
}

func (m *ArrivalEstimate) GetEventTime() string {
	if m != nil {
		return m.EventTime
	}
	return ""
}

func (m *ArrivalEstimate) GetLine() string {
	if m != nil {
		return m.Line
	}
	return ""
}

func (m *ArrivalEstimate) GetNextArrival() string {
	if m != nil {
		return m.NextArrival
	}
	return ""
}

func (m *ArrivalEstimate) GetStation() string {
	if m != nil {
		return m.Station
	}
	return ""
}

func (m *ArrivalEstimate) GetTrainID() string {
	if m != nil {
		return m.TrainID
	}
	return ""
}

func (m *ArrivalEstimate) GetWaitingSeconds() string {
	if m != nil {
		return m.WaitingSeconds
	}
	return ""
}

func (m *ArrivalEstimate) GetWaitingTime() string {
	if m != nil {
		return m.WaitingTime
	}
	return ""
}

func (m *ArrivalEstimate) GetTTL() int64 {
	if m != nil {
		return m.TTL
	}
	return 0
}

type GetArrivalEstimatesResponse struct {
	ArrivalEstimates     []*ArrivalEstimate `protobuf:"bytes,1,rep,name=ArrivalEstimates,proto3" json:"ArrivalEstimates,omitempty"`
	LastEvaluatedKey     string             `protobuf:"bytes,2,opt,name=LastEvaluatedKey,proto3" json:"LastEvaluatedKey,omitempty"`
	ResultLength         int32              `protobuf:"varint,3,opt,name=ResultLength,proto3" json:"ResultLength,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *GetArrivalEstimatesResponse) Reset()         { *m = GetArrivalEstimatesResponse{} }
func (m *GetArrivalEstimatesResponse) String() string { return proto.CompactTextString(m) }
func (*GetArrivalEstimatesResponse) ProtoMessage()    {}
func (*GetArrivalEstimatesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d00842e68e05382a, []int{2}
}

func (m *GetArrivalEstimatesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetArrivalEstimatesResponse.Unmarshal(m, b)
}
func (m *GetArrivalEstimatesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetArrivalEstimatesResponse.Marshal(b, m, deterministic)
}
func (m *GetArrivalEstimatesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetArrivalEstimatesResponse.Merge(m, src)
}
func (m *GetArrivalEstimatesResponse) XXX_Size() int {
	return xxx_messageInfo_GetArrivalEstimatesResponse.Size(m)
}
func (m *GetArrivalEstimatesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetArrivalEstimatesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetArrivalEstimatesResponse proto.InternalMessageInfo

func (m *GetArrivalEstimatesResponse) GetArrivalEstimates() []*ArrivalEstimate {
	if m != nil {
		return m.ArrivalEstimates
	}
	return nil
}

func (m *GetArrivalEstimatesResponse) GetLastEvaluatedKey() string {
	if m != nil {
		return m.LastEvaluatedKey
	}
	return ""
}

func (m *GetArrivalEstimatesResponse) GetResultLength() int32 {
	if m != nil {
		return m.ResultLength
	}
	return 0
}

func init() {
	proto.RegisterType((*GetArrivalEstimatesRequest)(nil), "schedule.GetArrivalEstimatesRequest")
	proto.RegisterType((*ArrivalEstimate)(nil), "schedule.ArrivalEstimate")
	proto.RegisterType((*GetArrivalEstimatesResponse)(nil), "schedule.GetArrivalEstimatesResponse")
}

func init() { proto.RegisterFile("schedule.proto", fileDescriptor_d00842e68e05382a) }

var fileDescriptor_d00842e68e05382a = []byte{
	// 499 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xd5, 0x26, 0x6d, 0xd3, 0x4c, 0xaa, 0x36, 0x5a, 0x2e, 0xdb, 0x50, 0x95, 0xc8, 0x2a, 0x28,
	0xaa, 0xd4, 0x44, 0x04, 0x0e, 0x88, 0x1b, 0x52, 0x22, 0x84, 0x88, 0x10, 0x72, 0x22, 0x71, 0xde,
	0x26, 0x43, 0xba, 0x92, 0xb3, 0x0e, 0xbb, 0x63, 0x8b, 0x5e, 0x39, 0x71, 0xe7, 0xca, 0x67, 0xf0,
	0x27, 0xfc, 0x02, 0x77, 0x2e, 0x7c, 0x00, 0xf2, 0xae, 0x4d, 0x5c, 0x27, 0xa8, 0x37, 0xcf, 0x7b,
	0x6f, 0xfc, 0xde, 0xce, 0xce, 0xc2, 0xb1, 0x9d, 0xdf, 0xe0, 0x22, 0x89, 0xb0, 0xbf, 0x36, 0x31,
	0xc5, 0xfc, 0xb0, 0xa8, 0x3b, 0x8f, 0x96, 0x71, 0xbc, 0x8c, 0x70, 0xe0, 0xf0, 0xeb, 0xe4, 0xe3,
	0x80, 0xd4, 0x0a, 0x2d, 0xc9, 0xd5, 0xda, 0x4b, 0x3b, 0x67, 0xb9, 0x40, 0xae, 0xd5, 0x40, 0x6a,
	0x1d, 0x93, 0x24, 0x15, 0x6b, 0xeb, 0xd9, 0xe0, 0x0f, 0x83, 0xce, 0x6b, 0xa4, 0x57, 0xc6, 0xa8,
	0x54, 0x46, 0x63, 0x4b, 0x6a, 0x25, 0x09, 0x6d, 0x88, 0x9f, 0x12, 0xb4, 0xc4, 0x5f, 0x40, 0x73,
	0x4a, 0xd2, 0xd0, 0x48, 0x12, 0x0a, 0xd6, 0x65, 0xbd, 0xd6, 0xb0, 0xd3, 0xf7, 0x3f, 0xec, 0x17,
	0x8e, 0xfd, 0x59, 0xe1, 0x18, 0x6e, 0xc4, 0xfc, 0x39, 0x34, 0xc6, 0x7a, 0xe1, 0xfa, 0x6a, 0xf7,
	0xf6, 0x15, 0x52, 0x2e, 0xa0, 0x31, 0xf5, 0x01, 0x45, 0xbd, 0xcb, 0x7a, 0xcd, 0xb0, 0x28, 0x79,
	0x17, 0x5a, 0x23, 0xb4, 0xa4, 0xb4, 0x67, 0xf7, 0x1c, 0x5b, 0x86, 0xf8, 0x25, 0xb4, 0x27, 0xd2,
	0xd2, 0x38, 0x95, 0x51, 0x22, 0x09, 0x17, 0x6f, 0xf1, 0x56, 0xec, 0x3b, 0xd9, 0x16, 0x1e, 0xfc,
	0xae, 0xc1, 0x49, 0xe5, 0xcc, 0xfc, 0x1c, 0xe0, 0xbd, 0x51, 0x2b, 0x69, 0x6e, 0xb3, 0x4e, 0xe6,
	0x3a, 0x4b, 0x88, 0xcb, 0x16, 0x1b, 0xca, 0xc8, 0x5a, 0x9e, 0xcd, 0x97, 0xd5, 0x6c, 0xf5, 0xed,
	0x6c, 0x67, 0xd0, 0x1c, 0x29, 0x83, 0xf3, 0x52, 0xf6, 0x0d, 0x90, 0xb1, 0xe3, 0x14, 0x35, 0x65,
	0x03, 0xc9, 0x23, 0x6f, 0x00, 0xce, 0x61, 0x6f, 0xa2, 0x34, 0x8a, 0x03, 0x47, 0xb8, 0xef, 0xcc,
	0xf1, 0x1d, 0x7e, 0x2e, 0xae, 0x4d, 0x34, 0xbc, 0x63, 0x09, 0x2a, 0x4f, 0xf2, 0xf0, 0xee, 0x24,
	0x05, 0x34, 0x66, 0x46, 0x2a, 0xfd, 0x66, 0x24, 0x9a, 0x9e, 0xc9, 0x4b, 0xfe, 0x04, 0x8e, 0x3f,
	0x48, 0x45, 0x4a, 0x2f, 0xa7, 0x38, 0x8f, 0xf5, 0xc2, 0x0a, 0x70, 0x82, 0x0a, 0x9a, 0xb9, 0xe7,
	0x88, 0x4b, 0xdc, 0xf2, 0xee, 0x25, 0x88, 0xb7, 0xa1, 0x3e, 0x9b, 0x4d, 0xc4, 0x51, 0x97, 0xf5,
	0xea, 0x61, 0xf6, 0x19, 0xfc, 0x60, 0xf0, 0x70, 0xe7, 0xa2, 0xd9, 0x75, 0xac, 0x2d, 0xf2, 0x31,
	0xb4, 0xab, 0x9c, 0x60, 0xdd, 0x7a, 0xaf, 0x35, 0x3c, 0xed, 0xff, 0x5b, 0xfe, 0x8a, 0x22, 0xdc,
	0x6a, 0xd9, 0xb9, 0x04, 0xb5, 0xdd, 0x4b, 0xc0, 0x03, 0x38, 0x0a, 0xd1, 0x26, 0x11, 0x4d, 0x50,
	0x2f, 0xe9, 0xc6, 0xdd, 0xdb, 0x7e, 0x78, 0x07, 0x1b, 0x7e, 0x67, 0x70, 0x32, 0xcd, 0xed, 0xa7,
	0x68, 0x52, 0x35, 0x47, 0xfe, 0x95, 0xc1, 0x83, 0x1d, 0x47, 0xe1, 0x17, 0x9b, 0xa0, 0xff, 0x7f,
	0x52, 0x9d, 0xc7, 0xf7, 0xa8, 0xfc, 0x3c, 0x82, 0x8b, 0x2f, 0x3f, 0x7f, 0x7d, 0xab, 0x9d, 0x07,
	0xa7, 0x83, 0xf4, 0xe9, 0x60, 0x89, 0x74, 0x25, 0xbd, 0xf2, 0x0a, 0x0b, 0xe9, 0x4b, 0x76, 0x79,
	0x7d, 0xe0, 0x1e, 0xd3, 0xb3, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xc3, 0xa1, 0xc7, 0x94, 0x20,
	0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ScheduleServiceClient is the client API for ScheduleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ScheduleServiceClient interface {
	GetArrivalEstimates(ctx context.Context, in *GetArrivalEstimatesRequest, opts ...grpc.CallOption) (*GetArrivalEstimatesResponse, error)
}

type scheduleServiceClient struct {
	cc *grpc.ClientConn
}

func NewScheduleServiceClient(cc *grpc.ClientConn) ScheduleServiceClient {
	return &scheduleServiceClient{cc}
}

func (c *scheduleServiceClient) GetArrivalEstimates(ctx context.Context, in *GetArrivalEstimatesRequest, opts ...grpc.CallOption) (*GetArrivalEstimatesResponse, error) {
	out := new(GetArrivalEstimatesResponse)
	err := c.cc.Invoke(ctx, "/schedule.ScheduleService/GetArrivalEstimates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ScheduleServiceServer is the server API for ScheduleService service.
type ScheduleServiceServer interface {
	GetArrivalEstimates(context.Context, *GetArrivalEstimatesRequest) (*GetArrivalEstimatesResponse, error)
}

// UnimplementedScheduleServiceServer can be embedded to have forward compatible implementations.
type UnimplementedScheduleServiceServer struct {
}

func (*UnimplementedScheduleServiceServer) GetArrivalEstimates(ctx context.Context, req *GetArrivalEstimatesRequest) (*GetArrivalEstimatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArrivalEstimates not implemented")
}

func RegisterScheduleServiceServer(s *grpc.Server, srv ScheduleServiceServer) {
	s.RegisterService(&_ScheduleService_serviceDesc, srv)
}

func _ScheduleService_GetArrivalEstimates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetArrivalEstimatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScheduleServiceServer).GetArrivalEstimates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/schedule.ScheduleService/GetArrivalEstimates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScheduleServiceServer).GetArrivalEstimates(ctx, req.(*GetArrivalEstimatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ScheduleService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "schedule.ScheduleService",
	HandlerType: (*ScheduleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetArrivalEstimates",
			Handler:    _ScheduleService_GetArrivalEstimates_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "schedule.proto",
}