// Code generated by protoc-gen-gogo.
// source: metrics.proto
// DO NOT EDIT!

/*
	Package metrics is a generated protocol buffer package.

	It is generated from these files:
		metrics.proto

	It has these top-level messages:
		LabelPair
		Gauge
		Counter
		Metric
		MetricFamily
*/
package metrics

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"
import github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type MetricType int32

const (
	MetricType_COUNTER MetricType = 0
	MetricType_GAUGE   MetricType = 1
)

var MetricType_name = map[int32]string{
	0: "COUNTER",
	1: "GAUGE",
}
var MetricType_value = map[string]int32{
	"COUNTER": 0,
	"GAUGE":   1,
}

func (x MetricType) Enum() *MetricType {
	p := new(MetricType)
	*p = x
	return p
}
func (x MetricType) String() string {
	return proto.EnumName(MetricType_name, int32(x))
}
func (x *MetricType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MetricType_value, data, "MetricType")
	if err != nil {
		return err
	}
	*x = MetricType(value)
	return nil
}
func (MetricType) EnumDescriptor() ([]byte, []int) { return fileDescriptorMetrics, []int{0} }

type LabelPair struct {
	Name  string `protobuf:"bytes,1,req,name=name" json:"name"`
	Value string `protobuf:"bytes,2,req,name=value" json:"value"`
}

func (m *LabelPair) Reset()                    { *m = LabelPair{} }
func (m *LabelPair) String() string            { return proto.CompactTextString(m) }
func (*LabelPair) ProtoMessage()               {}
func (*LabelPair) Descriptor() ([]byte, []int) { return fileDescriptorMetrics, []int{0} }

func (m *LabelPair) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LabelPair) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type Gauge struct {
	Value float64 `protobuf:"fixed64,1,req,name=value" json:"value"`
}

func (m *Gauge) Reset()                    { *m = Gauge{} }
func (m *Gauge) String() string            { return proto.CompactTextString(m) }
func (*Gauge) ProtoMessage()               {}
func (*Gauge) Descriptor() ([]byte, []int) { return fileDescriptorMetrics, []int{1} }

func (m *Gauge) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type Counter struct {
	Value float64 `protobuf:"fixed64,1,req,name=value" json:"value"`
}

func (m *Counter) Reset()                    { *m = Counter{} }
func (m *Counter) String() string            { return proto.CompactTextString(m) }
func (*Counter) ProtoMessage()               {}
func (*Counter) Descriptor() ([]byte, []int) { return fileDescriptorMetrics, []int{2} }

func (m *Counter) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type Metric struct {
	Label   []*LabelPair `protobuf:"bytes,1,rep,name=label" json:"label,omitempty"`
	Gauge   *Gauge       `protobuf:"bytes,2,opt,name=gauge" json:"gauge,omitempty"`
	Counter *Counter     `protobuf:"bytes,3,opt,name=counter" json:"counter,omitempty"`
}

func (m *Metric) Reset()                    { *m = Metric{} }
func (m *Metric) String() string            { return proto.CompactTextString(m) }
func (*Metric) ProtoMessage()               {}
func (*Metric) Descriptor() ([]byte, []int) { return fileDescriptorMetrics, []int{3} }

func (m *Metric) GetLabel() []*LabelPair {
	if m != nil {
		return m.Label
	}
	return nil
}

func (m *Metric) GetGauge() *Gauge {
	if m != nil {
		return m.Gauge
	}
	return nil
}

func (m *Metric) GetCounter() *Counter {
	if m != nil {
		return m.Counter
	}
	return nil
}

type MetricFamily struct {
	Name string `protobuf:"bytes,1,req,name=name" json:"name"`
	// reserved 2;
	Type   MetricType `protobuf:"varint,3,req,name=type,enum=metrics.MetricType" json:"type"`
	Metric []*Metric  `protobuf:"bytes,4,rep,name=metric" json:"metric,omitempty"`
}

func (m *MetricFamily) Reset()                    { *m = MetricFamily{} }
func (m *MetricFamily) String() string            { return proto.CompactTextString(m) }
func (*MetricFamily) ProtoMessage()               {}
func (*MetricFamily) Descriptor() ([]byte, []int) { return fileDescriptorMetrics, []int{4} }

func (m *MetricFamily) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MetricFamily) GetType() MetricType {
	if m != nil {
		return m.Type
	}
	return MetricType_COUNTER
}

func (m *MetricFamily) GetMetric() []*Metric {
	if m != nil {
		return m.Metric
	}
	return nil
}

func init() {
	proto.RegisterType((*LabelPair)(nil), "metrics.LabelPair")
	proto.RegisterType((*Gauge)(nil), "metrics.Gauge")
	proto.RegisterType((*Counter)(nil), "metrics.Counter")
	proto.RegisterType((*Metric)(nil), "metrics.Metric")
	proto.RegisterType((*MetricFamily)(nil), "metrics.MetricFamily")
	proto.RegisterEnum("metrics.MetricType", MetricType_name, MetricType_value)
}
func (m *LabelPair) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *LabelPair) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintMetrics(data, i, uint64(len(m.Name)))
	i += copy(data[i:], m.Name)
	data[i] = 0x12
	i++
	i = encodeVarintMetrics(data, i, uint64(len(m.Value)))
	i += copy(data[i:], m.Value)
	return i, nil
}

func (m *Gauge) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Gauge) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0x9
	i++
	i = encodeFixed64Metrics(data, i, uint64(math.Float64bits(float64(m.Value))))
	return i, nil
}

func (m *Counter) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Counter) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0x9
	i++
	i = encodeFixed64Metrics(data, i, uint64(math.Float64bits(float64(m.Value))))
	return i, nil
}

func (m *Metric) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Metric) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Label) > 0 {
		for _, msg := range m.Label {
			data[i] = 0xa
			i++
			i = encodeVarintMetrics(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.Gauge != nil {
		data[i] = 0x12
		i++
		i = encodeVarintMetrics(data, i, uint64(m.Gauge.Size()))
		n1, err := m.Gauge.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.Counter != nil {
		data[i] = 0x1a
		i++
		i = encodeVarintMetrics(data, i, uint64(m.Counter.Size()))
		n2, err := m.Counter.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}

func (m *MetricFamily) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *MetricFamily) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintMetrics(data, i, uint64(len(m.Name)))
	i += copy(data[i:], m.Name)
	data[i] = 0x18
	i++
	i = encodeVarintMetrics(data, i, uint64(m.Type))
	if len(m.Metric) > 0 {
		for _, msg := range m.Metric {
			data[i] = 0x22
			i++
			i = encodeVarintMetrics(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeFixed64Metrics(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Metrics(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintMetrics(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *LabelPair) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	n += 1 + l + sovMetrics(uint64(l))
	l = len(m.Value)
	n += 1 + l + sovMetrics(uint64(l))
	return n
}

func (m *Gauge) Size() (n int) {
	var l int
	_ = l
	n += 9
	return n
}

func (m *Counter) Size() (n int) {
	var l int
	_ = l
	n += 9
	return n
}

func (m *Metric) Size() (n int) {
	var l int
	_ = l
	if len(m.Label) > 0 {
		for _, e := range m.Label {
			l = e.Size()
			n += 1 + l + sovMetrics(uint64(l))
		}
	}
	if m.Gauge != nil {
		l = m.Gauge.Size()
		n += 1 + l + sovMetrics(uint64(l))
	}
	if m.Counter != nil {
		l = m.Counter.Size()
		n += 1 + l + sovMetrics(uint64(l))
	}
	return n
}

func (m *MetricFamily) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	n += 1 + l + sovMetrics(uint64(l))
	n += 1 + sovMetrics(uint64(m.Type))
	if len(m.Metric) > 0 {
		for _, e := range m.Metric {
			l = e.Size()
			n += 1 + l + sovMetrics(uint64(l))
		}
	}
	return n
}

func sovMetrics(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozMetrics(x uint64) (n int) {
	return sovMetrics(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *LabelPair) Unmarshal(data []byte) error {
	var hasFields [1]uint64
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMetrics
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: LabelPair: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LabelPair: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetrics
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMetrics
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(data[iNdEx:postIndex])
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetrics
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMetrics
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(data[iNdEx:postIndex])
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000002)
		default:
			iNdEx = preIndex
			skippy, err := skipMetrics(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMetrics
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("name")
	}
	if hasFields[0]&uint64(0x00000002) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("value")
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Gauge) Unmarshal(data []byte) error {
	var hasFields [1]uint64
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMetrics
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Gauge: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Gauge: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 8
			v = uint64(data[iNdEx-8])
			v |= uint64(data[iNdEx-7]) << 8
			v |= uint64(data[iNdEx-6]) << 16
			v |= uint64(data[iNdEx-5]) << 24
			v |= uint64(data[iNdEx-4]) << 32
			v |= uint64(data[iNdEx-3]) << 40
			v |= uint64(data[iNdEx-2]) << 48
			v |= uint64(data[iNdEx-1]) << 56
			m.Value = float64(math.Float64frombits(v))
			hasFields[0] |= uint64(0x00000001)
		default:
			iNdEx = preIndex
			skippy, err := skipMetrics(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMetrics
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("value")
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Counter) Unmarshal(data []byte) error {
	var hasFields [1]uint64
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMetrics
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Counter: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Counter: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 8
			v = uint64(data[iNdEx-8])
			v |= uint64(data[iNdEx-7]) << 8
			v |= uint64(data[iNdEx-6]) << 16
			v |= uint64(data[iNdEx-5]) << 24
			v |= uint64(data[iNdEx-4]) << 32
			v |= uint64(data[iNdEx-3]) << 40
			v |= uint64(data[iNdEx-2]) << 48
			v |= uint64(data[iNdEx-1]) << 56
			m.Value = float64(math.Float64frombits(v))
			hasFields[0] |= uint64(0x00000001)
		default:
			iNdEx = preIndex
			skippy, err := skipMetrics(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMetrics
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("value")
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Metric) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMetrics
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Metric: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Metric: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Label", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetrics
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMetrics
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Label = append(m.Label, &LabelPair{})
			if err := m.Label[len(m.Label)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Gauge", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetrics
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMetrics
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Gauge == nil {
				m.Gauge = &Gauge{}
			}
			if err := m.Gauge.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Counter", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetrics
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMetrics
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Counter == nil {
				m.Counter = &Counter{}
			}
			if err := m.Counter.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMetrics(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMetrics
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
func (m *MetricFamily) Unmarshal(data []byte) error {
	var hasFields [1]uint64
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMetrics
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MetricFamily: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MetricFamily: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetrics
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMetrics
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(data[iNdEx:postIndex])
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000001)
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetrics
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Type |= (MetricType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			hasFields[0] |= uint64(0x00000002)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metric", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetrics
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMetrics
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Metric = append(m.Metric, &Metric{})
			if err := m.Metric[len(m.Metric)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMetrics(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMetrics
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("name")
	}
	if hasFields[0]&uint64(0x00000002) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("type")
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipMetrics(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMetrics
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
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
					return 0, ErrIntOverflowMetrics
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
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
					return 0, ErrIntOverflowMetrics
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthMetrics
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowMetrics
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
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
				next, err := skipMetrics(data[start:])
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
	ErrInvalidLengthMetrics = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMetrics   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("metrics.proto", fileDescriptorMetrics) }

var fileDescriptorMetrics = []byte{
	// 278 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0xcd, 0x4d, 0x2d, 0x29,
	0xca, 0x4c, 0x2e, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x72, 0x95, 0x4c, 0xb8,
	0x38, 0x7d, 0x12, 0x93, 0x52, 0x73, 0x02, 0x12, 0x33, 0x8b, 0x84, 0x84, 0xb8, 0x58, 0xf2, 0x12,
	0x73, 0x53, 0x25, 0x18, 0x15, 0x98, 0x34, 0x38, 0x9d, 0x58, 0x4e, 0xdc, 0x93, 0x67, 0x10, 0x12,
	0xe6, 0x62, 0x2d, 0x4b, 0xcc, 0x29, 0x4d, 0x95, 0x60, 0x42, 0x08, 0x2a, 0xc9, 0x70, 0xb1, 0xba,
	0x27, 0x96, 0xa6, 0xa7, 0x22, 0x64, 0x41, 0x5a, 0x18, 0xa1, 0xb2, 0x72, 0x5c, 0xec, 0xce, 0xf9,
	0xa5, 0x79, 0x25, 0xa9, 0x45, 0xd8, 0xe5, 0x73, 0xb9, 0xd8, 0x7c, 0xc1, 0xd6, 0x0b, 0x29, 0x72,
	0xb1, 0xe6, 0x80, 0x6c, 0x97, 0x60, 0x54, 0x60, 0xd6, 0xe0, 0x36, 0x12, 0xd2, 0x83, 0xb9, 0x12,
	0xe1, 0x26, 0x59, 0x2e, 0xd6, 0x74, 0x90, 0x55, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0xdc, 0x46, 0x7c,
	0x70, 0x25, 0x10, 0x07, 0x28, 0x72, 0xb1, 0x27, 0x43, 0xec, 0x92, 0x60, 0x06, 0x2b, 0x10, 0x80,
	0x2b, 0x80, 0xba, 0x41, 0x29, 0x87, 0x8b, 0x07, 0x62, 0x9d, 0x5b, 0x62, 0x6e, 0x66, 0x4e, 0x25,
	0x56, 0x5f, 0xaa, 0x73, 0xb1, 0x94, 0x54, 0x16, 0xa4, 0x4a, 0x30, 0x2b, 0x30, 0x69, 0xf0, 0x19,
	0x09, 0xc3, 0xcd, 0x80, 0x68, 0x0c, 0xa9, 0x2c, 0x48, 0x85, 0x2a, 0x94, 0xe7, 0x62, 0x83, 0xc8,
	0x49, 0xb0, 0x80, 0x9d, 0xcc, 0x8f, 0xa6, 0x54, 0x4b, 0x85, 0x8b, 0x0b, 0xa1, 0x49, 0x88, 0x9b,
	0x8b, 0xdd, 0xd9, 0x3f, 0xd4, 0x2f, 0xc4, 0x35, 0x48, 0x80, 0x41, 0x88, 0x93, 0x8b, 0xd5, 0xdd,
	0x31, 0xd4, 0xdd, 0x55, 0x80, 0xd1, 0x49, 0xe0, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18,
	0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0x01, 0x10, 0x00, 0x00, 0xff, 0xff, 0xcf, 0x95,
	0x2c, 0x6f, 0xa1, 0x01, 0x00, 0x00,
}