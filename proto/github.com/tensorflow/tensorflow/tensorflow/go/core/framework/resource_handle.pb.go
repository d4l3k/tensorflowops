// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tensorflow/core/framework/resource_handle.proto

package framework

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Protocol buffer representing a handle to a tensorflow resource. Handles are
// not valid across executions, but can be serialized back and forth from within
// a single run.
type ResourceHandleProto struct {
	// Unique name for the device containing the resource.
	Device string `protobuf:"bytes,1,opt,name=device,proto3" json:"device,omitempty"`
	// Container in which this resource is placed.
	Container string `protobuf:"bytes,2,opt,name=container,proto3" json:"container,omitempty"`
	// Unique name of this resource.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Hash code for the type of the resource. Is only valid in the same device
	// and in the same execution.
	HashCode uint64 `protobuf:"varint,4,opt,name=hash_code,json=hashCode,proto3" json:"hash_code,omitempty"`
	// For debug-only, the name of the type pointed to by this handle, if
	// available.
	MaybeTypeName string `protobuf:"bytes,5,opt,name=maybe_type_name,json=maybeTypeName,proto3" json:"maybe_type_name,omitempty"`
}

func (m *ResourceHandleProto) Reset()      { *m = ResourceHandleProto{} }
func (*ResourceHandleProto) ProtoMessage() {}
func (*ResourceHandleProto) Descriptor() ([]byte, []int) {
	return fileDescriptorResourceHandle, []int{0}
}

func (m *ResourceHandleProto) GetDevice() string {
	if m != nil {
		return m.Device
	}
	return ""
}

func (m *ResourceHandleProto) GetContainer() string {
	if m != nil {
		return m.Container
	}
	return ""
}

func (m *ResourceHandleProto) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ResourceHandleProto) GetHashCode() uint64 {
	if m != nil {
		return m.HashCode
	}
	return 0
}

func (m *ResourceHandleProto) GetMaybeTypeName() string {
	if m != nil {
		return m.MaybeTypeName
	}
	return ""
}

func init() {
	proto.RegisterType((*ResourceHandleProto)(nil), "tensorflow.ResourceHandleProto")
}
func (this *ResourceHandleProto) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ResourceHandleProto)
	if !ok {
		that2, ok := that.(ResourceHandleProto)
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
	if this.Device != that1.Device {
		return false
	}
	if this.Container != that1.Container {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.HashCode != that1.HashCode {
		return false
	}
	if this.MaybeTypeName != that1.MaybeTypeName {
		return false
	}
	return true
}
func (this *ResourceHandleProto) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 9)
	s = append(s, "&framework.ResourceHandleProto{")
	s = append(s, "Device: "+fmt.Sprintf("%#v", this.Device)+",\n")
	s = append(s, "Container: "+fmt.Sprintf("%#v", this.Container)+",\n")
	s = append(s, "Name: "+fmt.Sprintf("%#v", this.Name)+",\n")
	s = append(s, "HashCode: "+fmt.Sprintf("%#v", this.HashCode)+",\n")
	s = append(s, "MaybeTypeName: "+fmt.Sprintf("%#v", this.MaybeTypeName)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringResourceHandle(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *ResourceHandleProto) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResourceHandleProto) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Device) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintResourceHandle(dAtA, i, uint64(len(m.Device)))
		i += copy(dAtA[i:], m.Device)
	}
	if len(m.Container) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintResourceHandle(dAtA, i, uint64(len(m.Container)))
		i += copy(dAtA[i:], m.Container)
	}
	if len(m.Name) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintResourceHandle(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if m.HashCode != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintResourceHandle(dAtA, i, uint64(m.HashCode))
	}
	if len(m.MaybeTypeName) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintResourceHandle(dAtA, i, uint64(len(m.MaybeTypeName)))
		i += copy(dAtA[i:], m.MaybeTypeName)
	}
	return i, nil
}

func encodeVarintResourceHandle(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ResourceHandleProto) Size() (n int) {
	var l int
	_ = l
	l = len(m.Device)
	if l > 0 {
		n += 1 + l + sovResourceHandle(uint64(l))
	}
	l = len(m.Container)
	if l > 0 {
		n += 1 + l + sovResourceHandle(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovResourceHandle(uint64(l))
	}
	if m.HashCode != 0 {
		n += 1 + sovResourceHandle(uint64(m.HashCode))
	}
	l = len(m.MaybeTypeName)
	if l > 0 {
		n += 1 + l + sovResourceHandle(uint64(l))
	}
	return n
}

func sovResourceHandle(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozResourceHandle(x uint64) (n int) {
	return sovResourceHandle(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *ResourceHandleProto) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ResourceHandleProto{`,
		`Device:` + fmt.Sprintf("%v", this.Device) + `,`,
		`Container:` + fmt.Sprintf("%v", this.Container) + `,`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`HashCode:` + fmt.Sprintf("%v", this.HashCode) + `,`,
		`MaybeTypeName:` + fmt.Sprintf("%v", this.MaybeTypeName) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringResourceHandle(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *ResourceHandleProto) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowResourceHandle
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
			return fmt.Errorf("proto: ResourceHandleProto: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResourceHandleProto: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Device", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResourceHandle
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
				return ErrInvalidLengthResourceHandle
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Device = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Container", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResourceHandle
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
				return ErrInvalidLengthResourceHandle
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Container = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResourceHandle
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
				return ErrInvalidLengthResourceHandle
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HashCode", wireType)
			}
			m.HashCode = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResourceHandle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.HashCode |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaybeTypeName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResourceHandle
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
				return ErrInvalidLengthResourceHandle
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MaybeTypeName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipResourceHandle(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthResourceHandle
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
func skipResourceHandle(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowResourceHandle
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
					return 0, ErrIntOverflowResourceHandle
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
					return 0, ErrIntOverflowResourceHandle
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
				return 0, ErrInvalidLengthResourceHandle
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowResourceHandle
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
				next, err := skipResourceHandle(dAtA[start:])
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
	ErrInvalidLengthResourceHandle = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowResourceHandle   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("tensorflow/core/framework/resource_handle.proto", fileDescriptorResourceHandle)
}

var fileDescriptorResourceHandle = []byte{
	// 291 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x3f, 0x4b, 0xf3, 0x50,
	0x14, 0xc6, 0x73, 0xde, 0xb7, 0x16, 0x73, 0x41, 0x85, 0x2b, 0x48, 0x40, 0xb9, 0x14, 0x07, 0xe9,
	0x94, 0x0c, 0xce, 0x2e, 0x75, 0x71, 0x12, 0x09, 0x4e, 0x2e, 0xe1, 0xe6, 0xe6, 0x34, 0x09, 0x36,
	0x39, 0xe1, 0x26, 0xb5, 0x64, 0xf3, 0x23, 0x38, 0x3b, 0x3b, 0xf8, 0x51, 0x1c, 0x3b, 0x3a, 0x9a,
	0xeb, 0xe2, 0xd8, 0xd1, 0x51, 0xbc, 0x14, 0x23, 0xc5, 0xed, 0x9c, 0xe7, 0xcf, 0xf0, 0xfc, 0x58,
	0xd0, 0x60, 0x59, 0x93, 0x9e, 0xce, 0x68, 0x11, 0x28, 0xd2, 0x18, 0x4c, 0xb5, 0x2c, 0x70, 0x41,
	0xfa, 0x36, 0xd0, 0x58, 0xd3, 0x5c, 0x2b, 0x8c, 0x32, 0x59, 0x26, 0x33, 0xf4, 0x2b, 0x4d, 0x0d,
	0x71, 0xd6, 0x17, 0x8e, 0x9f, 0x80, 0xed, 0x87, 0xeb, 0xd4, 0x85, 0x0d, 0x5d, 0xd9, 0xcc, 0x01,
	0x1b, 0x26, 0x78, 0x97, 0x2b, 0xf4, 0x60, 0x04, 0x63, 0x37, 0x5c, 0x7f, 0xfc, 0x88, 0xb9, 0x8a,
	0xca, 0x46, 0xe6, 0x25, 0x6a, 0xef, 0x9f, 0xb5, 0x7a, 0x81, 0x73, 0x36, 0x28, 0x65, 0x81, 0xde,
	0x7f, 0x6b, 0xd8, 0x9b, 0x1f, 0x32, 0x37, 0x93, 0x75, 0x16, 0x29, 0x4a, 0xd0, 0x1b, 0x8c, 0x60,
	0x3c, 0x08, 0xb7, 0xbf, 0x85, 0x73, 0x4a, 0x90, 0x9f, 0xb0, 0xbd, 0x42, 0xb6, 0x31, 0x46, 0x4d,
	0x5b, 0x61, 0x64, 0xbb, 0x5b, 0xb6, 0xbb, 0x63, 0xe5, 0xeb, 0xb6, 0xc2, 0x4b, 0x59, 0xe0, 0xe4,
	0x11, 0x96, 0x9d, 0x70, 0x5e, 0x3b, 0xe1, 0xac, 0x3a, 0x01, 0xf7, 0x46, 0xc0, 0xb3, 0x11, 0xf0,
	0x62, 0x04, 0x2c, 0x8d, 0x80, 0x37, 0x23, 0xe0, 0xc3, 0x08, 0x67, 0x65, 0x04, 0x3c, 0xbc, 0x0b,
	0x87, 0x79, 0xa4, 0x53, 0xbf, 0x5f, 0xe9, 0xff, 0x10, 0x99, 0xec, 0x6e, 0x8c, 0x85, 0x9b, 0xb3,
	0x34, 0x6f, 0xb2, 0x79, 0xec, 0x2b, 0x2a, 0x7e, 0x93, 0xfc, 0xfb, 0x4c, 0x69, 0x03, 0xf1, 0x27,
	0x40, 0x3c, 0xb4, 0x58, 0x4f, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x6b, 0x88, 0x17, 0x0a, 0x89,
	0x01, 0x00, 0x00,
}