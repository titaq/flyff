// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: flyff.proto

package domain

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type CharacterProps struct {
	ConnId string `protobuf:"bytes,1,opt,name=conn_id,json=connId,proto3" json:"conn_id,omitempty"`
	Id     string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *CharacterProps) Reset()      { *m = CharacterProps{} }
func (*CharacterProps) ProtoMessage() {}
func (*CharacterProps) Descriptor() ([]byte, []int) {
	return fileDescriptor_e39a2e871e89f4fa, []int{0}
}
func (m *CharacterProps) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CharacterProps) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CharacterProps.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CharacterProps) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CharacterProps.Merge(m, src)
}
func (m *CharacterProps) XXX_Size() int {
	return m.Size()
}
func (m *CharacterProps) XXX_DiscardUnknown() {
	xxx_messageInfo_CharacterProps.DiscardUnknown(m)
}

var xxx_messageInfo_CharacterProps proto.InternalMessageInfo

func (m *CharacterProps) GetConnId() string {
	if m != nil {
		return m.ConnId
	}
	return ""
}

func (m *CharacterProps) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type CharacterState struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *CharacterState) Reset()      { *m = CharacterState{} }
func (*CharacterState) ProtoMessage() {}
func (*CharacterState) Descriptor() ([]byte, []int) {
	return fileDescriptor_e39a2e871e89f4fa, []int{1}
}
func (m *CharacterState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CharacterState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CharacterState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CharacterState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CharacterState.Merge(m, src)
}
func (m *CharacterState) XXX_Size() int {
	return m.Size()
}
func (m *CharacterState) XXX_DiscardUnknown() {
	xxx_messageInfo_CharacterState.DiscardUnknown(m)
}

var xxx_messageInfo_CharacterState proto.InternalMessageInfo

func (m *CharacterState) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*CharacterProps)(nil), "domain.CharacterProps")
	proto.RegisterType((*CharacterState)(nil), "domain.CharacterState")
}

func init() { proto.RegisterFile("flyff.proto", fileDescriptor_e39a2e871e89f4fa) }

var fileDescriptor_e39a2e871e89f4fa = []byte{
	// 182 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0xcb, 0xa9, 0x4c,
	0x4b, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4b, 0xc9, 0xcf, 0x4d, 0xcc, 0xcc, 0x53,
	0xb2, 0xe4, 0xe2, 0x73, 0xce, 0x48, 0x2c, 0x4a, 0x4c, 0x2e, 0x49, 0x2d, 0x0a, 0x28, 0xca, 0x2f,
	0x28, 0x16, 0x12, 0xe7, 0x62, 0x4f, 0xce, 0xcf, 0xcb, 0x8b, 0xcf, 0x4c, 0x91, 0x60, 0x54, 0x60,
	0xd4, 0xe0, 0x0c, 0x62, 0x03, 0x71, 0x3d, 0x53, 0x84, 0xf8, 0xb8, 0x98, 0x32, 0x53, 0x24, 0x98,
	0xc0, 0x62, 0x4c, 0x99, 0x29, 0x4a, 0x2a, 0x48, 0x5a, 0x83, 0x4b, 0x12, 0x4b, 0x52, 0x85, 0x84,
	0xb8, 0x58, 0xf2, 0x12, 0x73, 0x53, 0xa1, 0xfa, 0xc0, 0x6c, 0x27, 0x93, 0x0b, 0x0f, 0xe5, 0x18,
	0x6e, 0x3c, 0x94, 0x63, 0xf8, 0xf0, 0x50, 0x8e, 0xb1, 0xe1, 0x91, 0x1c, 0xe3, 0x8a, 0x47, 0x72,
	0x8c, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x8b, 0x47,
	0x72, 0x0c, 0x1f, 0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d,
	0xc7, 0x72, 0x0c, 0x49, 0x6c, 0x60, 0x57, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xdf, 0xe7,
	0x96, 0x00, 0xb4, 0x00, 0x00, 0x00,
}

func (this *CharacterProps) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CharacterProps)
	if !ok {
		that2, ok := that.(CharacterProps)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.ConnId != that1.ConnId {
		return false
	}
	if this.Id != that1.Id {
		return false
	}
	return true
}
func (this *CharacterState) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CharacterState)
	if !ok {
		that2, ok := that.(CharacterState)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	return true
}
func (this *CharacterProps) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&domain.CharacterProps{")
	s = append(s, "ConnId: "+fmt.Sprintf("%#v", this.ConnId)+",\n")
	s = append(s, "Id: "+fmt.Sprintf("%#v", this.Id)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *CharacterState) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&domain.CharacterState{")
	s = append(s, "Name: "+fmt.Sprintf("%#v", this.Name)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringFlyff(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *CharacterProps) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CharacterProps) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CharacterProps) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintFlyff(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ConnId) > 0 {
		i -= len(m.ConnId)
		copy(dAtA[i:], m.ConnId)
		i = encodeVarintFlyff(dAtA, i, uint64(len(m.ConnId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CharacterState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CharacterState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CharacterState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintFlyff(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintFlyff(dAtA []byte, offset int, v uint64) int {
	offset -= sovFlyff(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CharacterProps) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ConnId)
	if l > 0 {
		n += 1 + l + sovFlyff(uint64(l))
	}
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovFlyff(uint64(l))
	}
	return n
}

func (m *CharacterState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovFlyff(uint64(l))
	}
	return n
}

func sovFlyff(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozFlyff(x uint64) (n int) {
	return sovFlyff(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *CharacterProps) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&CharacterProps{`,
		`ConnId:` + fmt.Sprintf("%v", this.ConnId) + `,`,
		`Id:` + fmt.Sprintf("%v", this.Id) + `,`,
		`}`,
	}, "")
	return s
}
func (this *CharacterState) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&CharacterState{`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringFlyff(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *CharacterProps) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFlyff
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CharacterProps: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CharacterProps: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConnId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFlyff
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthFlyff
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFlyff
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ConnId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFlyff
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthFlyff
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFlyff
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFlyff(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFlyff
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
func (m *CharacterState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFlyff
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CharacterState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CharacterState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFlyff
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthFlyff
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFlyff
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFlyff(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFlyff
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
func skipFlyff(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFlyff
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
					return 0, ErrIntOverflowFlyff
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowFlyff
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
			if length < 0 {
				return 0, ErrInvalidLengthFlyff
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupFlyff
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthFlyff
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthFlyff        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFlyff          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupFlyff = fmt.Errorf("proto: unexpected end of group")
)
