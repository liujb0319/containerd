// Code generated by protoc-gen-gogo.
// source: user.proto
// DO NOT EDIT!

/*
	Package user is a generated protocol buffer package.

	It is generated from these files:
		user.proto

	It has these top-level messages:
		User
*/
package user

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import strings "strings"
import github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"
import sort "sort"
import strconv "strconv"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type User struct {
	Uid            uint32   `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Gid            uint32   `protobuf:"varint,2,opt,name=gid,proto3" json:"gid,omitempty"`
	AdditionalGids []uint32 `protobuf:"varint,3,rep,packed,name=additional_gids,json=additionalGids" json:"additional_gids,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptorUser, []int{0} }

func init() {
	proto.RegisterType((*User)(nil), "containerd.v1.types.User")
}
func (this *User) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&user.User{")
	s = append(s, "Uid: "+fmt.Sprintf("%#v", this.Uid)+",\n")
	s = append(s, "Gid: "+fmt.Sprintf("%#v", this.Gid)+",\n")
	s = append(s, "AdditionalGids: "+fmt.Sprintf("%#v", this.AdditionalGids)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringUser(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func extensionToGoStringUser(m github_com_gogo_protobuf_proto.Message) string {
	e := github_com_gogo_protobuf_proto.GetUnsafeExtensionsMap(m)
	if e == nil {
		return "nil"
	}
	s := "proto.NewUnsafeXXX_InternalExtensions(map[int32]proto.Extension{"
	keys := make([]int, 0, len(e))
	for k := range e {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)
	ss := []string{}
	for _, k := range keys {
		ss = append(ss, strconv.Itoa(k)+": "+e[int32(k)].GoString())
	}
	s += strings.Join(ss, ",") + "})"
	return s
}
func (m *User) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *User) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Uid != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintUser(dAtA, i, uint64(m.Uid))
	}
	if m.Gid != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintUser(dAtA, i, uint64(m.Gid))
	}
	if len(m.AdditionalGids) > 0 {
		dAtA2 := make([]byte, len(m.AdditionalGids)*10)
		var j1 int
		for _, num := range m.AdditionalGids {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		dAtA[i] = 0x1a
		i++
		i = encodeVarintUser(dAtA, i, uint64(j1))
		i += copy(dAtA[i:], dAtA2[:j1])
	}
	return i, nil
}

func encodeFixed64User(dAtA []byte, offset int, v uint64) int {
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
func encodeFixed32User(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintUser(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *User) Size() (n int) {
	var l int
	_ = l
	if m.Uid != 0 {
		n += 1 + sovUser(uint64(m.Uid))
	}
	if m.Gid != 0 {
		n += 1 + sovUser(uint64(m.Gid))
	}
	if len(m.AdditionalGids) > 0 {
		l = 0
		for _, e := range m.AdditionalGids {
			l += sovUser(uint64(e))
		}
		n += 1 + sovUser(uint64(l)) + l
	}
	return n
}

func sovUser(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozUser(x uint64) (n int) {
	return sovUser(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *User) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&User{`,
		`Uid:` + fmt.Sprintf("%v", this.Uid) + `,`,
		`Gid:` + fmt.Sprintf("%v", this.Gid) + `,`,
		`AdditionalGids:` + fmt.Sprintf("%v", this.AdditionalGids) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringUser(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *User) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUser
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
			return fmt.Errorf("proto: User: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: User: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uid", wireType)
			}
			m.Uid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Uid |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Gid", wireType)
			}
			m.Gid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Gid |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowUser
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthUser
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v uint32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowUser
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (uint32(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.AdditionalGids = append(m.AdditionalGids, v)
				}
			} else if wireType == 0 {
				var v uint32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowUser
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (uint32(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.AdditionalGids = append(m.AdditionalGids, v)
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field AdditionalGids", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipUser(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthUser
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
func skipUser(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowUser
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
					return 0, ErrIntOverflowUser
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
					return 0, ErrIntOverflowUser
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
				return 0, ErrInvalidLengthUser
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowUser
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
				next, err := skipUser(dAtA[start:])
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
	ErrInvalidLengthUser = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowUser   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("user.proto", fileDescriptorUser) }

var fileDescriptorUser = []byte{
	// 162 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x2d, 0x4e, 0x2d,
	0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x4e, 0xce, 0xcf, 0x2b, 0x49, 0xcc, 0xcc, 0x4b,
	0x2d, 0x4a, 0xd1, 0x2b, 0x33, 0xd4, 0x2b, 0xa9, 0x2c, 0x48, 0x2d, 0x56, 0x0a, 0xe6, 0x62, 0x09,
	0x2d, 0x4e, 0x2d, 0x12, 0x12, 0xe0, 0x62, 0x2e, 0xcd, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0,
	0x0d, 0x02, 0x31, 0x41, 0x22, 0xe9, 0x99, 0x29, 0x12, 0x4c, 0x10, 0x91, 0xf4, 0xcc, 0x14, 0x21,
	0x75, 0x2e, 0xfe, 0xc4, 0x94, 0x94, 0xcc, 0x92, 0xcc, 0xfc, 0xbc, 0xc4, 0x9c, 0xf8, 0xf4, 0xcc,
	0x94, 0x62, 0x09, 0x66, 0x05, 0x66, 0x0d, 0xde, 0x20, 0x3e, 0x84, 0xb0, 0x7b, 0x66, 0x4a, 0xb1,
	0x93, 0xcc, 0x89, 0x87, 0x72, 0x0c, 0x37, 0x1e, 0xca, 0x31, 0x7c, 0x78, 0x28, 0xc7, 0xd8, 0xf0,
	0x48, 0x8e, 0xf1, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x4c,
	0x62, 0x03, 0x3b, 0xc7, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x28, 0xf0, 0x6e, 0xd2, 0x9c, 0x00,
	0x00, 0x00,
}