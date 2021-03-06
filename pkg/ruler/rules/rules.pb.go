// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: rules.proto

package rules

import (
	fmt "fmt"
	_ "github.com/cortexproject/cortex/pkg/ingester/client"
	github_com_cortexproject_cortex_pkg_ingester_client "github.com/cortexproject/cortex/pkg/ingester/client"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "github.com/golang/protobuf/ptypes/duration"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// RuleGroupDesc is a proto representation of a cortex rule group
type RuleGroupDesc struct {
	Name      string        `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Namespace string        `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Interval  time.Duration `protobuf:"bytes,3,opt,name=interval,proto3,stdduration" json:"interval"`
	Rules     []*RuleDesc   `protobuf:"bytes,4,rep,name=rules,proto3" json:"rules,omitempty"`
	User      string        `protobuf:"bytes,6,opt,name=user,proto3" json:"user,omitempty"`
}

func (m *RuleGroupDesc) Reset()      { *m = RuleGroupDesc{} }
func (*RuleGroupDesc) ProtoMessage() {}
func (*RuleGroupDesc) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e722d3e922f0937, []int{0}
}
func (m *RuleGroupDesc) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RuleGroupDesc) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RuleGroupDesc.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RuleGroupDesc) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RuleGroupDesc.Merge(m, src)
}
func (m *RuleGroupDesc) XXX_Size() int {
	return m.Size()
}
func (m *RuleGroupDesc) XXX_DiscardUnknown() {
	xxx_messageInfo_RuleGroupDesc.DiscardUnknown(m)
}

var xxx_messageInfo_RuleGroupDesc proto.InternalMessageInfo

func (m *RuleGroupDesc) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RuleGroupDesc) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *RuleGroupDesc) GetInterval() time.Duration {
	if m != nil {
		return m.Interval
	}
	return 0
}

func (m *RuleGroupDesc) GetRules() []*RuleDesc {
	if m != nil {
		return m.Rules
	}
	return nil
}

func (m *RuleGroupDesc) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

// RuleDesc is a proto representation of a Prometheus Rule
type RuleDesc struct {
	Expr        string                                                             `protobuf:"bytes,1,opt,name=expr,proto3" json:"expr,omitempty"`
	Record      string                                                             `protobuf:"bytes,2,opt,name=record,proto3" json:"record,omitempty"`
	Alert       string                                                             `protobuf:"bytes,3,opt,name=alert,proto3" json:"alert,omitempty"`
	For         time.Duration                                                      `protobuf:"bytes,4,opt,name=for,proto3,stdduration" json:"for"`
	Labels      []github_com_cortexproject_cortex_pkg_ingester_client.LabelAdapter `protobuf:"bytes,5,rep,name=labels,proto3,customtype=github.com/cortexproject/cortex/pkg/ingester/client.LabelAdapter" json:"labels"`
	Annotations []github_com_cortexproject_cortex_pkg_ingester_client.LabelAdapter `protobuf:"bytes,6,rep,name=annotations,proto3,customtype=github.com/cortexproject/cortex/pkg/ingester/client.LabelAdapter" json:"annotations"`
}

func (m *RuleDesc) Reset()      { *m = RuleDesc{} }
func (*RuleDesc) ProtoMessage() {}
func (*RuleDesc) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e722d3e922f0937, []int{1}
}
func (m *RuleDesc) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RuleDesc) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RuleDesc.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RuleDesc) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RuleDesc.Merge(m, src)
}
func (m *RuleDesc) XXX_Size() int {
	return m.Size()
}
func (m *RuleDesc) XXX_DiscardUnknown() {
	xxx_messageInfo_RuleDesc.DiscardUnknown(m)
}

var xxx_messageInfo_RuleDesc proto.InternalMessageInfo

func (m *RuleDesc) GetExpr() string {
	if m != nil {
		return m.Expr
	}
	return ""
}

func (m *RuleDesc) GetRecord() string {
	if m != nil {
		return m.Record
	}
	return ""
}

func (m *RuleDesc) GetAlert() string {
	if m != nil {
		return m.Alert
	}
	return ""
}

func (m *RuleDesc) GetFor() time.Duration {
	if m != nil {
		return m.For
	}
	return 0
}

func init() {
	proto.RegisterType((*RuleGroupDesc)(nil), "rules.RuleGroupDesc")
	proto.RegisterType((*RuleDesc)(nil), "rules.RuleDesc")
}

func init() { proto.RegisterFile("rules.proto", fileDescriptor_8e722d3e922f0937) }

var fileDescriptor_8e722d3e922f0937 = []byte{
	// 456 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x51, 0x41, 0x8b, 0xd3, 0x40,
	0x14, 0xce, 0x6c, 0xd3, 0x98, 0x4e, 0x59, 0xac, 0x83, 0xc8, 0xb8, 0xc8, 0xb4, 0x2c, 0x08, 0xbd,
	0x98, 0xc0, 0x8a, 0x27, 0x0f, 0x6a, 0x59, 0x50, 0x8a, 0x07, 0xc9, 0xd1, 0xdb, 0x34, 0x7d, 0x1b,
	0xa3, 0xb3, 0x99, 0x30, 0x33, 0x11, 0x2f, 0x82, 0x3f, 0xc1, 0xa3, 0x3f, 0xc1, 0x9f, 0xb2, 0x37,
	0x7b, 0x5c, 0x3c, 0xac, 0x36, 0xbd, 0x78, 0xdc, 0x3f, 0x20, 0xc8, 0xcc, 0xa4, 0x6e, 0x8f, 0x22,
	0xec, 0x29, 0xef, 0xcb, 0x9b, 0xef, 0xbd, 0xef, 0xfb, 0x1e, 0x1e, 0xaa, 0x46, 0x80, 0x4e, 0x6a,
	0x25, 0x8d, 0x24, 0x7d, 0x07, 0x0e, 0x1e, 0x14, 0xa5, 0x79, 0xd3, 0x2c, 0x92, 0x5c, 0x9e, 0xa6,
	0x85, 0x2c, 0x64, 0xea, 0xba, 0x8b, 0xe6, 0xc4, 0x21, 0x07, 0x5c, 0xe5, 0x59, 0x07, 0xac, 0x90,
	0xb2, 0x10, 0x70, 0xf5, 0x6a, 0xd9, 0x28, 0x6e, 0x4a, 0x59, 0x75, 0xfd, 0xa7, 0x3b, 0xe3, 0x72,
	0xa9, 0x0c, 0x7c, 0xa8, 0x95, 0x7c, 0x0b, 0xb9, 0xe9, 0x50, 0x5a, 0xbf, 0x2b, 0xd2, 0xb2, 0x2a,
	0x40, 0x1b, 0x50, 0x69, 0x2e, 0x4a, 0xa8, 0xb6, 0x2d, 0x3f, 0xe1, 0xf0, 0x1b, 0xc2, 0xfb, 0x59,
	0x23, 0xe0, 0xb9, 0x92, 0x4d, 0x7d, 0x0c, 0x3a, 0x27, 0x04, 0x87, 0x15, 0x3f, 0x05, 0x8a, 0x26,
	0x68, 0x3a, 0xc8, 0x5c, 0x4d, 0xee, 0xe1, 0x81, 0xfd, 0xea, 0x9a, 0xe7, 0x40, 0xf7, 0x5c, 0xe3,
	0xea, 0x07, 0x79, 0x82, 0xe3, 0xb2, 0x32, 0xa0, 0xde, 0x73, 0x41, 0x7b, 0x13, 0x34, 0x1d, 0x1e,
	0xdd, 0x4d, 0xbc, 0xf0, 0x64, 0x2b, 0x3c, 0x39, 0xee, 0x84, 0xcf, 0xe2, 0xb3, 0x8b, 0x71, 0xf0,
	0xe5, 0xc7, 0x18, 0x65, 0x7f, 0x49, 0xe4, 0x3e, 0xf6, 0xf1, 0xd0, 0x70, 0xd2, 0x9b, 0x0e, 0x8f,
	0x6e, 0x26, 0x3e, 0x39, 0xab, 0xcb, 0x4a, 0xca, 0x7c, 0xd7, 0x2a, 0x6b, 0x34, 0x28, 0x1a, 0x79,
	0x65, 0xb6, 0x9e, 0x87, 0x71, 0x7f, 0x14, 0xcd, 0xc3, 0xf8, 0xc6, 0x28, 0x9e, 0x87, 0x71, 0x3c,
	0x1a, 0x1c, 0xfe, 0xde, 0xc3, 0xf1, 0x96, 0x69, 0x29, 0x36, 0x93, 0xad, 0x19, 0x5b, 0x93, 0x3b,
	0x38, 0x52, 0x90, 0x4b, 0xb5, 0xec, 0x9c, 0x74, 0x88, 0xdc, 0xc6, 0x7d, 0x2e, 0x40, 0x19, 0xe7,
	0x61, 0x90, 0x79, 0x40, 0x1e, 0xe1, 0xde, 0x89, 0x54, 0x34, 0xfc, 0x77, 0x5f, 0xf6, 0x3d, 0xd1,
	0x38, 0x12, 0x7c, 0x01, 0x42, 0xd3, 0xbe, 0xf3, 0x74, 0x2b, 0xe9, 0x62, 0x7f, 0x69, 0xff, 0xbe,
	0xe2, 0xa5, 0x9a, 0xbd, 0xb0, 0x8c, 0xef, 0x17, 0xe3, 0xff, 0x39, 0xa2, 0x1f, 0xf3, 0x6c, 0xc9,
	0x6b, 0x03, 0x2a, 0xeb, 0x56, 0x91, 0x8f, 0x78, 0xc8, 0xab, 0x4a, 0x1a, 0xa7, 0x48, 0xd3, 0xe8,
	0xfa, 0x37, 0xef, 0xee, 0x73, 0x57, 0xd8, 0x9f, 0x3d, 0x5e, 0xad, 0x59, 0x70, 0xbe, 0x66, 0xc1,
	0xe5, 0x9a, 0xa1, 0x4f, 0x2d, 0x43, 0x5f, 0x5b, 0x86, 0xce, 0x5a, 0x86, 0x56, 0x2d, 0x43, 0x3f,
	0x5b, 0x86, 0x7e, 0xb5, 0x2c, 0xb8, 0x6c, 0x19, 0xfa, 0xbc, 0x61, 0xc1, 0x6a, 0xc3, 0x82, 0xf3,
	0x0d, 0x0b, 0x5e, 0xfb, 0x13, 0x2f, 0x22, 0x17, 0xec, 0xc3, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x15, 0x3f, 0x81, 0xc2, 0x3c, 0x03, 0x00, 0x00,
}

func (this *RuleGroupDesc) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RuleGroupDesc)
	if !ok {
		that2, ok := that.(RuleGroupDesc)
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
	if this.Namespace != that1.Namespace {
		return false
	}
	if this.Interval != that1.Interval {
		return false
	}
	if len(this.Rules) != len(that1.Rules) {
		return false
	}
	for i := range this.Rules {
		if !this.Rules[i].Equal(that1.Rules[i]) {
			return false
		}
	}
	if this.User != that1.User {
		return false
	}
	return true
}
func (this *RuleDesc) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RuleDesc)
	if !ok {
		that2, ok := that.(RuleDesc)
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
	if this.Expr != that1.Expr {
		return false
	}
	if this.Record != that1.Record {
		return false
	}
	if this.Alert != that1.Alert {
		return false
	}
	if this.For != that1.For {
		return false
	}
	if len(this.Labels) != len(that1.Labels) {
		return false
	}
	for i := range this.Labels {
		if !this.Labels[i].Equal(that1.Labels[i]) {
			return false
		}
	}
	if len(this.Annotations) != len(that1.Annotations) {
		return false
	}
	for i := range this.Annotations {
		if !this.Annotations[i].Equal(that1.Annotations[i]) {
			return false
		}
	}
	return true
}
func (this *RuleGroupDesc) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 9)
	s = append(s, "&rules.RuleGroupDesc{")
	s = append(s, "Name: "+fmt.Sprintf("%#v", this.Name)+",\n")
	s = append(s, "Namespace: "+fmt.Sprintf("%#v", this.Namespace)+",\n")
	s = append(s, "Interval: "+fmt.Sprintf("%#v", this.Interval)+",\n")
	if this.Rules != nil {
		s = append(s, "Rules: "+fmt.Sprintf("%#v", this.Rules)+",\n")
	}
	s = append(s, "User: "+fmt.Sprintf("%#v", this.User)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *RuleDesc) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 10)
	s = append(s, "&rules.RuleDesc{")
	s = append(s, "Expr: "+fmt.Sprintf("%#v", this.Expr)+",\n")
	s = append(s, "Record: "+fmt.Sprintf("%#v", this.Record)+",\n")
	s = append(s, "Alert: "+fmt.Sprintf("%#v", this.Alert)+",\n")
	s = append(s, "For: "+fmt.Sprintf("%#v", this.For)+",\n")
	s = append(s, "Labels: "+fmt.Sprintf("%#v", this.Labels)+",\n")
	s = append(s, "Annotations: "+fmt.Sprintf("%#v", this.Annotations)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringRules(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *RuleGroupDesc) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RuleGroupDesc) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RuleGroupDesc) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.User) > 0 {
		i -= len(m.User)
		copy(dAtA[i:], m.User)
		i = encodeVarintRules(dAtA, i, uint64(len(m.User)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Rules) > 0 {
		for iNdEx := len(m.Rules) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Rules[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintRules(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	n1, err1 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.Interval, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.Interval):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintRules(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x1a
	if len(m.Namespace) > 0 {
		i -= len(m.Namespace)
		copy(dAtA[i:], m.Namespace)
		i = encodeVarintRules(dAtA, i, uint64(len(m.Namespace)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintRules(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RuleDesc) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RuleDesc) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RuleDesc) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Annotations) > 0 {
		for iNdEx := len(m.Annotations) - 1; iNdEx >= 0; iNdEx-- {
			{
				size := m.Annotations[iNdEx].Size()
				i -= size
				if _, err := m.Annotations[iNdEx].MarshalTo(dAtA[i:]); err != nil {
					return 0, err
				}
				i = encodeVarintRules(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.Labels) > 0 {
		for iNdEx := len(m.Labels) - 1; iNdEx >= 0; iNdEx-- {
			{
				size := m.Labels[iNdEx].Size()
				i -= size
				if _, err := m.Labels[iNdEx].MarshalTo(dAtA[i:]); err != nil {
					return 0, err
				}
				i = encodeVarintRules(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	n2, err2 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.For, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.For):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintRules(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x22
	if len(m.Alert) > 0 {
		i -= len(m.Alert)
		copy(dAtA[i:], m.Alert)
		i = encodeVarintRules(dAtA, i, uint64(len(m.Alert)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Record) > 0 {
		i -= len(m.Record)
		copy(dAtA[i:], m.Record)
		i = encodeVarintRules(dAtA, i, uint64(len(m.Record)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Expr) > 0 {
		i -= len(m.Expr)
		copy(dAtA[i:], m.Expr)
		i = encodeVarintRules(dAtA, i, uint64(len(m.Expr)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintRules(dAtA []byte, offset int, v uint64) int {
	offset -= sovRules(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RuleGroupDesc) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovRules(uint64(l))
	}
	l = len(m.Namespace)
	if l > 0 {
		n += 1 + l + sovRules(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.Interval)
	n += 1 + l + sovRules(uint64(l))
	if len(m.Rules) > 0 {
		for _, e := range m.Rules {
			l = e.Size()
			n += 1 + l + sovRules(uint64(l))
		}
	}
	l = len(m.User)
	if l > 0 {
		n += 1 + l + sovRules(uint64(l))
	}
	return n
}

func (m *RuleDesc) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Expr)
	if l > 0 {
		n += 1 + l + sovRules(uint64(l))
	}
	l = len(m.Record)
	if l > 0 {
		n += 1 + l + sovRules(uint64(l))
	}
	l = len(m.Alert)
	if l > 0 {
		n += 1 + l + sovRules(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.For)
	n += 1 + l + sovRules(uint64(l))
	if len(m.Labels) > 0 {
		for _, e := range m.Labels {
			l = e.Size()
			n += 1 + l + sovRules(uint64(l))
		}
	}
	if len(m.Annotations) > 0 {
		for _, e := range m.Annotations {
			l = e.Size()
			n += 1 + l + sovRules(uint64(l))
		}
	}
	return n
}

func sovRules(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRules(x uint64) (n int) {
	return sovRules(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *RuleGroupDesc) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForRules := "[]*RuleDesc{"
	for _, f := range this.Rules {
		repeatedStringForRules += strings.Replace(f.String(), "RuleDesc", "RuleDesc", 1) + ","
	}
	repeatedStringForRules += "}"
	s := strings.Join([]string{`&RuleGroupDesc{`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`Namespace:` + fmt.Sprintf("%v", this.Namespace) + `,`,
		`Interval:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.Interval), "Duration", "duration.Duration", 1), `&`, ``, 1) + `,`,
		`Rules:` + repeatedStringForRules + `,`,
		`User:` + fmt.Sprintf("%v", this.User) + `,`,
		`}`,
	}, "")
	return s
}
func (this *RuleDesc) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&RuleDesc{`,
		`Expr:` + fmt.Sprintf("%v", this.Expr) + `,`,
		`Record:` + fmt.Sprintf("%v", this.Record) + `,`,
		`Alert:` + fmt.Sprintf("%v", this.Alert) + `,`,
		`For:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.For), "Duration", "duration.Duration", 1), `&`, ``, 1) + `,`,
		`Labels:` + fmt.Sprintf("%v", this.Labels) + `,`,
		`Annotations:` + fmt.Sprintf("%v", this.Annotations) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringRules(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *RuleGroupDesc) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRules
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
			return fmt.Errorf("proto: RuleGroupDesc: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RuleGroupDesc: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRules
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
				return ErrInvalidLengthRules
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRules
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Namespace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRules
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
				return ErrInvalidLengthRules
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRules
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Namespace = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Interval", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRules
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRules
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRules
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.Interval, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rules", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRules
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRules
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRules
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Rules = append(m.Rules, &RuleDesc{})
			if err := m.Rules[len(m.Rules)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field User", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRules
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
				return ErrInvalidLengthRules
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRules
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.User = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRules(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRules
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthRules
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
func (m *RuleDesc) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRules
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
			return fmt.Errorf("proto: RuleDesc: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RuleDesc: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Expr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRules
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
				return ErrInvalidLengthRules
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRules
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Expr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Record", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRules
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
				return ErrInvalidLengthRules
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRules
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Record = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Alert", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRules
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
				return ErrInvalidLengthRules
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRules
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Alert = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field For", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRules
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRules
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRules
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.For, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Labels", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRules
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRules
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRules
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Labels = append(m.Labels, github_com_cortexproject_cortex_pkg_ingester_client.LabelAdapter{})
			if err := m.Labels[len(m.Labels)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Annotations", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRules
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRules
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRules
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Annotations = append(m.Annotations, github_com_cortexproject_cortex_pkg_ingester_client.LabelAdapter{})
			if err := m.Annotations[len(m.Annotations)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRules(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRules
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthRules
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
func skipRules(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRules
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
					return 0, ErrIntOverflowRules
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
					return 0, ErrIntOverflowRules
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
				return 0, ErrInvalidLengthRules
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthRules
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowRules
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
				next, err := skipRules(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthRules
				}
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
	ErrInvalidLengthRules = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRules   = fmt.Errorf("proto: integer overflow")
)
