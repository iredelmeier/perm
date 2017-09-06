// Code generated by protoc-gen-gogo.
// source: role.proto
// DO NOT EDIT!

package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type AssignRoleRequest struct {
	Actor   string   `protobuf:"bytes,1,opt,name=actor,proto3" json:"actor,omitempty"`
	Role    string   `protobuf:"bytes,2,opt,name=role,proto3" json:"role,omitempty"`
	Context *Context `protobuf:"bytes,3,opt,name=context" json:"context,omitempty"`
}

func (m *AssignRoleRequest) Reset()                    { *m = AssignRoleRequest{} }
func (m *AssignRoleRequest) String() string            { return proto.CompactTextString(m) }
func (*AssignRoleRequest) ProtoMessage()               {}
func (*AssignRoleRequest) Descriptor() ([]byte, []int) { return fileDescriptorRole, []int{0} }

func (m *AssignRoleRequest) GetActor() string {
	if m != nil {
		return m.Actor
	}
	return ""
}

func (m *AssignRoleRequest) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func (m *AssignRoleRequest) GetContext() *Context {
	if m != nil {
		return m.Context
	}
	return nil
}

type AssignRoleResponse struct {
}

func (m *AssignRoleResponse) Reset()                    { *m = AssignRoleResponse{} }
func (m *AssignRoleResponse) String() string            { return proto.CompactTextString(m) }
func (*AssignRoleResponse) ProtoMessage()               {}
func (*AssignRoleResponse) Descriptor() ([]byte, []int) { return fileDescriptorRole, []int{1} }

type HasRoleRequest struct {
	Actor   string   `protobuf:"bytes,1,opt,name=actor,proto3" json:"actor,omitempty"`
	Role    string   `protobuf:"bytes,2,opt,name=role,proto3" json:"role,omitempty"`
	Context *Context `protobuf:"bytes,3,opt,name=context" json:"context,omitempty"`
}

func (m *HasRoleRequest) Reset()                    { *m = HasRoleRequest{} }
func (m *HasRoleRequest) String() string            { return proto.CompactTextString(m) }
func (*HasRoleRequest) ProtoMessage()               {}
func (*HasRoleRequest) Descriptor() ([]byte, []int) { return fileDescriptorRole, []int{2} }

func (m *HasRoleRequest) GetActor() string {
	if m != nil {
		return m.Actor
	}
	return ""
}

func (m *HasRoleRequest) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func (m *HasRoleRequest) GetContext() *Context {
	if m != nil {
		return m.Context
	}
	return nil
}

type HasRoleResponse struct {
	HasRole bool `protobuf:"varint,1,opt,name=hasRole,proto3" json:"hasRole,omitempty"`
}

func (m *HasRoleResponse) Reset()                    { *m = HasRoleResponse{} }
func (m *HasRoleResponse) String() string            { return proto.CompactTextString(m) }
func (*HasRoleResponse) ProtoMessage()               {}
func (*HasRoleResponse) Descriptor() ([]byte, []int) { return fileDescriptorRole, []int{3} }

func (m *HasRoleResponse) GetHasRole() bool {
	if m != nil {
		return m.HasRole
	}
	return false
}

func init() {
	proto.RegisterType((*AssignRoleRequest)(nil), "cloud_foundry.perm.protos.AssignRoleRequest")
	proto.RegisterType((*AssignRoleResponse)(nil), "cloud_foundry.perm.protos.AssignRoleResponse")
	proto.RegisterType((*HasRoleRequest)(nil), "cloud_foundry.perm.protos.HasRoleRequest")
	proto.RegisterType((*HasRoleResponse)(nil), "cloud_foundry.perm.protos.HasRoleResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for RoleService service

type RoleServiceClient interface {
	AssignRole(ctx context.Context, in *AssignRoleRequest, opts ...grpc.CallOption) (*AssignRoleResponse, error)
	HasRole(ctx context.Context, in *HasRoleRequest, opts ...grpc.CallOption) (*HasRoleResponse, error)
}

type roleServiceClient struct {
	cc *grpc.ClientConn
}

func NewRoleServiceClient(cc *grpc.ClientConn) RoleServiceClient {
	return &roleServiceClient{cc}
}

func (c *roleServiceClient) AssignRole(ctx context.Context, in *AssignRoleRequest, opts ...grpc.CallOption) (*AssignRoleResponse, error) {
	out := new(AssignRoleResponse)
	err := grpc.Invoke(ctx, "/cloud_foundry.perm.protos.RoleService/AssignRole", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleServiceClient) HasRole(ctx context.Context, in *HasRoleRequest, opts ...grpc.CallOption) (*HasRoleResponse, error) {
	out := new(HasRoleResponse)
	err := grpc.Invoke(ctx, "/cloud_foundry.perm.protos.RoleService/HasRole", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RoleService service

type RoleServiceServer interface {
	AssignRole(context.Context, *AssignRoleRequest) (*AssignRoleResponse, error)
	HasRole(context.Context, *HasRoleRequest) (*HasRoleResponse, error)
}

func RegisterRoleServiceServer(s *grpc.Server, srv RoleServiceServer) {
	s.RegisterService(&_RoleService_serviceDesc, srv)
}

func _RoleService_AssignRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssignRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).AssignRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud_foundry.perm.protos.RoleService/AssignRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).AssignRole(ctx, req.(*AssignRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleService_HasRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HasRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).HasRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud_foundry.perm.protos.RoleService/HasRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).HasRole(ctx, req.(*HasRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RoleService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cloud_foundry.perm.protos.RoleService",
	HandlerType: (*RoleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AssignRole",
			Handler:    _RoleService_AssignRole_Handler,
		},
		{
			MethodName: "HasRole",
			Handler:    _RoleService_HasRole_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "role.proto",
}

func (m *AssignRoleRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AssignRoleRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Actor) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRole(dAtA, i, uint64(len(m.Actor)))
		i += copy(dAtA[i:], m.Actor)
	}
	if len(m.Role) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintRole(dAtA, i, uint64(len(m.Role)))
		i += copy(dAtA[i:], m.Role)
	}
	if m.Context != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintRole(dAtA, i, uint64(m.Context.Size()))
		n1, err := m.Context.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func (m *AssignRoleResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AssignRoleResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *HasRoleRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HasRoleRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Actor) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRole(dAtA, i, uint64(len(m.Actor)))
		i += copy(dAtA[i:], m.Actor)
	}
	if len(m.Role) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintRole(dAtA, i, uint64(len(m.Role)))
		i += copy(dAtA[i:], m.Role)
	}
	if m.Context != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintRole(dAtA, i, uint64(m.Context.Size()))
		n2, err := m.Context.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}

func (m *HasRoleResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HasRoleResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.HasRole {
		dAtA[i] = 0x8
		i++
		if m.HasRole {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}

func encodeFixed64Role(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Role(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintRole(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *AssignRoleRequest) Size() (n int) {
	var l int
	_ = l
	l = len(m.Actor)
	if l > 0 {
		n += 1 + l + sovRole(uint64(l))
	}
	l = len(m.Role)
	if l > 0 {
		n += 1 + l + sovRole(uint64(l))
	}
	if m.Context != nil {
		l = m.Context.Size()
		n += 1 + l + sovRole(uint64(l))
	}
	return n
}

func (m *AssignRoleResponse) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *HasRoleRequest) Size() (n int) {
	var l int
	_ = l
	l = len(m.Actor)
	if l > 0 {
		n += 1 + l + sovRole(uint64(l))
	}
	l = len(m.Role)
	if l > 0 {
		n += 1 + l + sovRole(uint64(l))
	}
	if m.Context != nil {
		l = m.Context.Size()
		n += 1 + l + sovRole(uint64(l))
	}
	return n
}

func (m *HasRoleResponse) Size() (n int) {
	var l int
	_ = l
	if m.HasRole {
		n += 2
	}
	return n
}

func sovRole(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozRole(x uint64) (n int) {
	return sovRole(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AssignRoleRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRole
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AssignRoleRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AssignRoleRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Actor", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRole
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRole
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Actor = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Role", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRole
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRole
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Role = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Context", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRole
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRole
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Context == nil {
				m.Context = &Context{}
			}
			if err := m.Context.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRole(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRole
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AssignRoleResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRole
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AssignRoleResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AssignRoleResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipRole(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRole
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *HasRoleRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRole
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: HasRoleRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HasRoleRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Actor", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRole
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRole
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Actor = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Role", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRole
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRole
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Role = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Context", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRole
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRole
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Context == nil {
				m.Context = &Context{}
			}
			if err := m.Context.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRole(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRole
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *HasRoleResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRole
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: HasRoleResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HasRoleResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HasRole", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRole
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.HasRole = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipRole(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRole
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipRole(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRole
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRole
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRole
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthRole
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowRole
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipRole(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthRole = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRole   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("role.proto", fileDescriptorRole) }

var fileDescriptorRole = []byte{
	// 296 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0xca, 0xcf, 0x49,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x4c, 0xce, 0xc9, 0x2f, 0x4d, 0x89, 0x4f, 0xcb,
	0x2f, 0xcd, 0x4b, 0x29, 0xaa, 0xd4, 0x2b, 0x48, 0x2d, 0xca, 0x85, 0xc8, 0x14, 0x4b, 0xe9, 0xa6,
	0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0xa7, 0xe7, 0xa7, 0xe7, 0xeb, 0x83,
	0xc5, 0x93, 0x4a, 0xd3, 0xc0, 0x3c, 0x30, 0x07, 0xcc, 0x82, 0xa8, 0x97, 0xe2, 0x4d, 0xce, 0xcf,
	0x2b, 0x49, 0xad, 0x28, 0x81, 0x70, 0x95, 0xaa, 0xb9, 0x04, 0x1d, 0x8b, 0x8b, 0x33, 0xd3, 0xf3,
	0x82, 0xf2, 0x73, 0x52, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x44, 0xb8, 0x58, 0x13,
	0x93, 0x4b, 0xf2, 0x8b, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x20, 0x1c, 0x21, 0x21, 0x2e,
	0x16, 0x90, 0x8b, 0x24, 0x98, 0xc0, 0x82, 0x60, 0xb6, 0x90, 0x0d, 0x17, 0x3b, 0xd4, 0x3c, 0x09,
	0x66, 0x05, 0x46, 0x0d, 0x6e, 0x23, 0x25, 0x3d, 0x9c, 0x2e, 0xd5, 0x73, 0x86, 0xa8, 0x0c, 0x82,
	0x69, 0x51, 0x12, 0xe1, 0x12, 0x42, 0xb6, 0xbc, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x55, 0xa9, 0x82,
	0x8b, 0xcf, 0x23, 0xb1, 0x78, 0x20, 0xdc, 0xa3, 0xcd, 0xc5, 0x0f, 0xb7, 0x19, 0xe2, 0x18, 0x21,
	0x09, 0x2e, 0xf6, 0x0c, 0x88, 0x10, 0xd8, 0x72, 0x8e, 0x20, 0x18, 0xd7, 0xe8, 0x16, 0x23, 0x17,
	0x37, 0x88, 0x11, 0x9c, 0x5a, 0x54, 0x96, 0x99, 0x9c, 0x2a, 0x94, 0xc9, 0xc5, 0x85, 0xf0, 0x8c,
	0x90, 0x0e, 0x1e, 0x7b, 0x31, 0x02, 0x5c, 0x4a, 0x97, 0x48, 0xd5, 0x50, 0x47, 0x25, 0x70, 0xb1,
	0x43, 0xdd, 0x29, 0xa4, 0x89, 0x47, 0x27, 0x6a, 0x28, 0x4a, 0x69, 0x11, 0xa3, 0x14, 0x62, 0x83,
	0x93, 0xc4, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe3,
	0xb1, 0x1c, 0x43, 0x14, 0x1b, 0x44, 0x65, 0x12, 0x84, 0x36, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff,
	0x45, 0x0f, 0xb1, 0x90, 0x9e, 0x02, 0x00, 0x00,
}
