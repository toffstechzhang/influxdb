// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: meta.proto

package backend

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

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

// StoreTaskMeta is the internal state of a task.
type StoreTaskMeta struct {
	MaxConcurrency int32 `protobuf:"varint,1,opt,name=max_concurrency,json=maxConcurrency,proto3" json:"max_concurrency,omitempty"`
	// latest_completed is the unix timestamp of the latest "naturally" completed run.
	// If a run for time t finishes before a run for time t - u, latest_completed will reflect time t.
	LatestCompleted int64 `protobuf:"varint,2,opt,name=latest_completed,json=latestCompleted,proto3" json:"latest_completed,omitempty"`
	// status indicates if the task is enabled or disabled.
	Status string `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	// currently_running is the collection of runs in-progress.
	// If a runner crashes or otherwise disappears, this indicates to the new runner what needs to be picked up.
	CurrentlyRunning []*StoreTaskMetaRun `protobuf:"bytes,4,rep,name=currently_running,json=currentlyRunning,proto3" json:"currently_running,omitempty"`
	// effective_cron is the effective cron string as reported by the task's options.
	EffectiveCron string `protobuf:"bytes,5,opt,name=effective_cron,json=effectiveCron,proto3" json:"effective_cron,omitempty"`
	// Task's configured delay, in seconds.
	Offset int32 `protobuf:"varint,6,opt,name=offset,proto3" json:"offset,omitempty"`
	// The Authorization ID associated with the task.
	AuthorizationID uint64                    `protobuf:"varint,7,opt,name=authorization_id,json=authorizationId,proto3" json:"authorization_id,omitempty"`
	ManualRuns      []*StoreTaskMetaManualRun `protobuf:"bytes,16,rep,name=manual_runs,json=manualRuns,proto3" json:"manual_runs,omitempty"`
}

func (m *StoreTaskMeta) Reset()         { *m = StoreTaskMeta{} }
func (m *StoreTaskMeta) String() string { return proto.CompactTextString(m) }
func (*StoreTaskMeta) ProtoMessage()    {}
func (*StoreTaskMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_meta_0932e91782f0009e, []int{0}
}
func (m *StoreTaskMeta) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StoreTaskMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StoreTaskMeta.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *StoreTaskMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoreTaskMeta.Merge(dst, src)
}
func (m *StoreTaskMeta) XXX_Size() int {
	return m.Size()
}
func (m *StoreTaskMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_StoreTaskMeta.DiscardUnknown(m)
}

var xxx_messageInfo_StoreTaskMeta proto.InternalMessageInfo

func (m *StoreTaskMeta) GetMaxConcurrency() int32 {
	if m != nil {
		return m.MaxConcurrency
	}
	return 0
}

func (m *StoreTaskMeta) GetLatestCompleted() int64 {
	if m != nil {
		return m.LatestCompleted
	}
	return 0
}

func (m *StoreTaskMeta) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *StoreTaskMeta) GetCurrentlyRunning() []*StoreTaskMetaRun {
	if m != nil {
		return m.CurrentlyRunning
	}
	return nil
}

func (m *StoreTaskMeta) GetEffectiveCron() string {
	if m != nil {
		return m.EffectiveCron
	}
	return ""
}

func (m *StoreTaskMeta) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *StoreTaskMeta) GetAuthorizationID() uint64 {
	if m != nil {
		return m.AuthorizationID
	}
	return 0
}

func (m *StoreTaskMeta) GetManualRuns() []*StoreTaskMetaManualRun {
	if m != nil {
		return m.ManualRuns
	}
	return nil
}

type StoreTaskMetaRun struct {
	// now is the unix timestamp of the "now" value for the run.
	Now   int64  `protobuf:"varint,1,opt,name=now,proto3" json:"now,omitempty"`
	Try   uint32 `protobuf:"varint,2,opt,name=try,proto3" json:"try,omitempty"`
	RunID uint64 `protobuf:"varint,3,opt,name=run_id,json=runId,proto3" json:"run_id,omitempty"`
	// range_start is the start of the manual run's time range.
	RangeStart int64 `protobuf:"varint,4,opt,name=range_start,json=rangeStart,proto3" json:"range_start,omitempty"`
	// range_end is the end of the manual run's time range.
	RangeEnd int64 `protobuf:"varint,5,opt,name=range_end,json=rangeEnd,proto3" json:"range_end,omitempty"`
	// requested_at is the unix timestamp indicating when this run was requested.
	// It is the same value as the "parent" StoreTaskMetaManualRun, if this run was the result of a manual request.
	RequestedAt int64 `protobuf:"varint,6,opt,name=requested_at,json=requestedAt,proto3" json:"requested_at,omitempty"`
}

func (m *StoreTaskMetaRun) Reset()         { *m = StoreTaskMetaRun{} }
func (m *StoreTaskMetaRun) String() string { return proto.CompactTextString(m) }
func (*StoreTaskMetaRun) ProtoMessage()    {}
func (*StoreTaskMetaRun) Descriptor() ([]byte, []int) {
	return fileDescriptor_meta_0932e91782f0009e, []int{1}
}
func (m *StoreTaskMetaRun) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StoreTaskMetaRun) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StoreTaskMetaRun.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *StoreTaskMetaRun) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoreTaskMetaRun.Merge(dst, src)
}
func (m *StoreTaskMetaRun) XXX_Size() int {
	return m.Size()
}
func (m *StoreTaskMetaRun) XXX_DiscardUnknown() {
	xxx_messageInfo_StoreTaskMetaRun.DiscardUnknown(m)
}

var xxx_messageInfo_StoreTaskMetaRun proto.InternalMessageInfo

func (m *StoreTaskMetaRun) GetNow() int64 {
	if m != nil {
		return m.Now
	}
	return 0
}

func (m *StoreTaskMetaRun) GetTry() uint32 {
	if m != nil {
		return m.Try
	}
	return 0
}

func (m *StoreTaskMetaRun) GetRunID() uint64 {
	if m != nil {
		return m.RunID
	}
	return 0
}

func (m *StoreTaskMetaRun) GetRangeStart() int64 {
	if m != nil {
		return m.RangeStart
	}
	return 0
}

func (m *StoreTaskMetaRun) GetRangeEnd() int64 {
	if m != nil {
		return m.RangeEnd
	}
	return 0
}

func (m *StoreTaskMetaRun) GetRequestedAt() int64 {
	if m != nil {
		return m.RequestedAt
	}
	return 0
}

// StoreTaskMetaManualRun indicates a manually requested run for a time range.
// It has a start and end pair of unix timestamps indicating the time range covered by the request.
type StoreTaskMetaManualRun struct {
	// start is the earliest allowable unix time stamp for this queue of runs.
	Start int64 `protobuf:"varint,1,opt,name=start,proto3" json:"start,omitempty"`
	// end is the latest allowable unix time stamp for this queue of runs.
	End int64 `protobuf:"varint,2,opt,name=end,proto3" json:"end,omitempty"`
	// latest_completed is the timestamp of the latest completed run from this queue.
	LatestCompleted int64 `protobuf:"varint,3,opt,name=latest_completed,json=latestCompleted,proto3" json:"latest_completed,omitempty"`
	// requested_at is the unix timestamp indicating when this run was requested.
	RequestedAt int64 `protobuf:"varint,4,opt,name=requested_at,json=requestedAt,proto3" json:"requested_at,omitempty"`
	// run_id is set ahead of time for retries of individual runs. Manually run time ranges do not receive an ID.
	RunID uint64 `protobuf:"varint,5,opt,name=run_id,json=runId,proto3" json:"run_id,omitempty"`
}

func (m *StoreTaskMetaManualRun) Reset()         { *m = StoreTaskMetaManualRun{} }
func (m *StoreTaskMetaManualRun) String() string { return proto.CompactTextString(m) }
func (*StoreTaskMetaManualRun) ProtoMessage()    {}
func (*StoreTaskMetaManualRun) Descriptor() ([]byte, []int) {
	return fileDescriptor_meta_0932e91782f0009e, []int{2}
}
func (m *StoreTaskMetaManualRun) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StoreTaskMetaManualRun) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StoreTaskMetaManualRun.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *StoreTaskMetaManualRun) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoreTaskMetaManualRun.Merge(dst, src)
}
func (m *StoreTaskMetaManualRun) XXX_Size() int {
	return m.Size()
}
func (m *StoreTaskMetaManualRun) XXX_DiscardUnknown() {
	xxx_messageInfo_StoreTaskMetaManualRun.DiscardUnknown(m)
}

var xxx_messageInfo_StoreTaskMetaManualRun proto.InternalMessageInfo

func (m *StoreTaskMetaManualRun) GetStart() int64 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *StoreTaskMetaManualRun) GetEnd() int64 {
	if m != nil {
		return m.End
	}
	return 0
}

func (m *StoreTaskMetaManualRun) GetLatestCompleted() int64 {
	if m != nil {
		return m.LatestCompleted
	}
	return 0
}

func (m *StoreTaskMetaManualRun) GetRequestedAt() int64 {
	if m != nil {
		return m.RequestedAt
	}
	return 0
}

func (m *StoreTaskMetaManualRun) GetRunID() uint64 {
	if m != nil {
		return m.RunID
	}
	return 0
}

func init() {
	proto.RegisterType((*StoreTaskMeta)(nil), "com.influxdata.platform.task.backend.StoreTaskMeta")
	proto.RegisterType((*StoreTaskMetaRun)(nil), "com.influxdata.platform.task.backend.StoreTaskMetaRun")
	proto.RegisterType((*StoreTaskMetaManualRun)(nil), "com.influxdata.platform.task.backend.StoreTaskMetaManualRun")
}
func (m *StoreTaskMeta) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StoreTaskMeta) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.MaxConcurrency != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintMeta(dAtA, i, uint64(m.MaxConcurrency))
	}
	if m.LatestCompleted != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintMeta(dAtA, i, uint64(m.LatestCompleted))
	}
	if len(m.Status) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintMeta(dAtA, i, uint64(len(m.Status)))
		i += copy(dAtA[i:], m.Status)
	}
	if len(m.CurrentlyRunning) > 0 {
		for _, msg := range m.CurrentlyRunning {
			dAtA[i] = 0x22
			i++
			i = encodeVarintMeta(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.EffectiveCron) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintMeta(dAtA, i, uint64(len(m.EffectiveCron)))
		i += copy(dAtA[i:], m.EffectiveCron)
	}
	if m.Offset != 0 {
		dAtA[i] = 0x30
		i++
		i = encodeVarintMeta(dAtA, i, uint64(m.Offset))
	}
	if m.AuthorizationID != 0 {
		dAtA[i] = 0x38
		i++
		i = encodeVarintMeta(dAtA, i, uint64(m.AuthorizationID))
	}
	if len(m.ManualRuns) > 0 {
		for _, msg := range m.ManualRuns {
			dAtA[i] = 0x82
			i++
			dAtA[i] = 0x1
			i++
			i = encodeVarintMeta(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *StoreTaskMetaRun) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StoreTaskMetaRun) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Now != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintMeta(dAtA, i, uint64(m.Now))
	}
	if m.Try != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintMeta(dAtA, i, uint64(m.Try))
	}
	if m.RunID != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintMeta(dAtA, i, uint64(m.RunID))
	}
	if m.RangeStart != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintMeta(dAtA, i, uint64(m.RangeStart))
	}
	if m.RangeEnd != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintMeta(dAtA, i, uint64(m.RangeEnd))
	}
	if m.RequestedAt != 0 {
		dAtA[i] = 0x30
		i++
		i = encodeVarintMeta(dAtA, i, uint64(m.RequestedAt))
	}
	return i, nil
}

func (m *StoreTaskMetaManualRun) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StoreTaskMetaManualRun) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Start != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintMeta(dAtA, i, uint64(m.Start))
	}
	if m.End != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintMeta(dAtA, i, uint64(m.End))
	}
	if m.LatestCompleted != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintMeta(dAtA, i, uint64(m.LatestCompleted))
	}
	if m.RequestedAt != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintMeta(dAtA, i, uint64(m.RequestedAt))
	}
	if m.RunID != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintMeta(dAtA, i, uint64(m.RunID))
	}
	return i, nil
}

func encodeVarintMeta(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *StoreTaskMeta) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MaxConcurrency != 0 {
		n += 1 + sovMeta(uint64(m.MaxConcurrency))
	}
	if m.LatestCompleted != 0 {
		n += 1 + sovMeta(uint64(m.LatestCompleted))
	}
	l = len(m.Status)
	if l > 0 {
		n += 1 + l + sovMeta(uint64(l))
	}
	if len(m.CurrentlyRunning) > 0 {
		for _, e := range m.CurrentlyRunning {
			l = e.Size()
			n += 1 + l + sovMeta(uint64(l))
		}
	}
	l = len(m.EffectiveCron)
	if l > 0 {
		n += 1 + l + sovMeta(uint64(l))
	}
	if m.Offset != 0 {
		n += 1 + sovMeta(uint64(m.Offset))
	}
	if m.AuthorizationID != 0 {
		n += 1 + sovMeta(uint64(m.AuthorizationID))
	}
	if len(m.ManualRuns) > 0 {
		for _, e := range m.ManualRuns {
			l = e.Size()
			n += 2 + l + sovMeta(uint64(l))
		}
	}
	return n
}

func (m *StoreTaskMetaRun) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Now != 0 {
		n += 1 + sovMeta(uint64(m.Now))
	}
	if m.Try != 0 {
		n += 1 + sovMeta(uint64(m.Try))
	}
	if m.RunID != 0 {
		n += 1 + sovMeta(uint64(m.RunID))
	}
	if m.RangeStart != 0 {
		n += 1 + sovMeta(uint64(m.RangeStart))
	}
	if m.RangeEnd != 0 {
		n += 1 + sovMeta(uint64(m.RangeEnd))
	}
	if m.RequestedAt != 0 {
		n += 1 + sovMeta(uint64(m.RequestedAt))
	}
	return n
}

func (m *StoreTaskMetaManualRun) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Start != 0 {
		n += 1 + sovMeta(uint64(m.Start))
	}
	if m.End != 0 {
		n += 1 + sovMeta(uint64(m.End))
	}
	if m.LatestCompleted != 0 {
		n += 1 + sovMeta(uint64(m.LatestCompleted))
	}
	if m.RequestedAt != 0 {
		n += 1 + sovMeta(uint64(m.RequestedAt))
	}
	if m.RunID != 0 {
		n += 1 + sovMeta(uint64(m.RunID))
	}
	return n
}

func sovMeta(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozMeta(x uint64) (n int) {
	return sovMeta(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *StoreTaskMeta) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMeta
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
			return fmt.Errorf("proto: StoreTaskMeta: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StoreTaskMeta: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxConcurrency", wireType)
			}
			m.MaxConcurrency = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMeta
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxConcurrency |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LatestCompleted", wireType)
			}
			m.LatestCompleted = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMeta
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LatestCompleted |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMeta
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
				return ErrInvalidLengthMeta
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Status = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentlyRunning", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMeta
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
				return ErrInvalidLengthMeta
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CurrentlyRunning = append(m.CurrentlyRunning, &StoreTaskMetaRun{})
			if err := m.CurrentlyRunning[len(m.CurrentlyRunning)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EffectiveCron", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMeta
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
				return ErrInvalidLengthMeta
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EffectiveCron = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Offset", wireType)
			}
			m.Offset = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMeta
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Offset |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthorizationID", wireType)
			}
			m.AuthorizationID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMeta
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AuthorizationID |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 16:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ManualRuns", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMeta
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
				return ErrInvalidLengthMeta
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ManualRuns = append(m.ManualRuns, &StoreTaskMetaManualRun{})
			if err := m.ManualRuns[len(m.ManualRuns)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMeta(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMeta
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
func (m *StoreTaskMetaRun) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMeta
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
			return fmt.Errorf("proto: StoreTaskMetaRun: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StoreTaskMetaRun: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Now", wireType)
			}
			m.Now = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMeta
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Now |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Try", wireType)
			}
			m.Try = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMeta
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Try |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RunID", wireType)
			}
			m.RunID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMeta
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RunID |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RangeStart", wireType)
			}
			m.RangeStart = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMeta
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RangeStart |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RangeEnd", wireType)
			}
			m.RangeEnd = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMeta
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RangeEnd |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestedAt", wireType)
			}
			m.RequestedAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMeta
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RequestedAt |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMeta(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMeta
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
func (m *StoreTaskMetaManualRun) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMeta
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
			return fmt.Errorf("proto: StoreTaskMetaManualRun: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StoreTaskMetaManualRun: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Start", wireType)
			}
			m.Start = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMeta
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Start |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field End", wireType)
			}
			m.End = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMeta
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.End |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LatestCompleted", wireType)
			}
			m.LatestCompleted = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMeta
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LatestCompleted |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestedAt", wireType)
			}
			m.RequestedAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMeta
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RequestedAt |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RunID", wireType)
			}
			m.RunID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMeta
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RunID |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMeta(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMeta
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
func skipMeta(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMeta
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
					return 0, ErrIntOverflowMeta
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
					return 0, ErrIntOverflowMeta
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
				return 0, ErrInvalidLengthMeta
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowMeta
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
				next, err := skipMeta(dAtA[start:])
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
	ErrInvalidLengthMeta = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMeta   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("meta.proto", fileDescriptor_meta_0932e91782f0009e) }

var fileDescriptor_meta_0932e91782f0009e = []byte{
	// 516 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0x1b, 0xd2, 0x74, 0xd4, 0xa5, 0x6b, 0x30, 0xd3, 0x14, 0x81, 0x94, 0x66, 0x15, 0x88,
	0x72, 0x09, 0x12, 0x48, 0x9c, 0x10, 0xd2, 0x56, 0x38, 0xec, 0xb0, 0x8b, 0xc7, 0x09, 0x09, 0x45,
	0x5e, 0xe2, 0x94, 0xa8, 0x89, 0x3d, 0x9c, 0x67, 0x68, 0xf9, 0x14, 0x7c, 0x14, 0xae, 0x5c, 0x38,
	0x73, 0xdc, 0x91, 0xd3, 0x84, 0xd2, 0x2f, 0x82, 0x6c, 0x77, 0x65, 0x1b, 0x3d, 0xa0, 0xdd, 0xde,
	0xfb, 0xc7, 0x7e, 0xfe, 0xff, 0xde, 0x7b, 0x41, 0xa8, 0x62, 0x40, 0xe3, 0x53, 0x29, 0x40, 0xe0,
	0x87, 0xa9, 0xa8, 0xe2, 0x82, 0xe7, 0xa5, 0x9a, 0x67, 0x54, 0xab, 0x25, 0x85, 0x5c, 0xc8, 0x2a,
	0x06, 0x5a, 0xcf, 0xe2, 0x13, 0x9a, 0xce, 0x18, 0xcf, 0xee, 0xef, 0x4c, 0xc5, 0x54, 0x98, 0x0b,
	0x4f, 0x75, 0x64, 0xef, 0x8e, 0x7e, 0xb8, 0xa8, 0x7f, 0x0c, 0x42, 0xb2, 0xb7, 0xb4, 0x9e, 0x1d,
	0x31, 0xa0, 0xf8, 0x31, 0x1a, 0x54, 0x74, 0x9e, 0xa4, 0x82, 0xa7, 0x4a, 0x4a, 0xc6, 0xd3, 0x45,
	0xe0, 0x44, 0xce, 0xd8, 0x23, 0xdb, 0x15, 0x9d, 0x4f, 0xfe, 0xaa, 0xf8, 0x09, 0xf2, 0x4b, 0x0a,
	0xac, 0x86, 0x24, 0x15, 0xd5, 0x69, 0xc9, 0x80, 0x65, 0xc1, 0xad, 0xc8, 0x19, 0xbb, 0x64, 0x60,
	0xf5, 0xc9, 0x85, 0x8c, 0x77, 0x51, 0xa7, 0x06, 0x0a, 0xaa, 0x0e, 0xdc, 0xc8, 0x19, 0x77, 0xc9,
	0x2a, 0xc3, 0x29, 0xba, 0x6b, 0xcb, 0x41, 0xb9, 0x48, 0xa4, 0xe2, 0xbc, 0xe0, 0xd3, 0xa0, 0x1d,
	0xb9, 0xe3, 0xde, 0xb3, 0x17, 0xf1, 0xff, 0x50, 0xc5, 0x57, 0xbc, 0x13, 0xc5, 0x89, 0xbf, 0x2e,
	0x48, 0x6c, 0x3d, 0xfc, 0x08, 0x6d, 0xb3, 0x3c, 0x67, 0x29, 0x14, 0x9f, 0x58, 0x92, 0x4a, 0xc1,
	0x03, 0xcf, 0x98, 0xe8, 0xaf, 0xd5, 0x89, 0x14, 0x5c, 0x7b, 0x14, 0x79, 0x5e, 0x33, 0x08, 0x3a,
	0x06, 0x77, 0x95, 0xe1, 0x57, 0xc8, 0xa7, 0x0a, 0x3e, 0x08, 0x59, 0x7c, 0xa1, 0x50, 0x08, 0x9e,
	0x14, 0x59, 0xb0, 0x15, 0x39, 0xe3, 0xf6, 0xc1, 0xbd, 0xe6, 0x7c, 0x38, 0xd8, 0xbf, 0xfc, 0xed,
	0xf0, 0x35, 0x19, 0x5c, 0x39, 0x7c, 0x98, 0xe1, 0xf7, 0xa8, 0x57, 0x51, 0xae, 0x68, 0xa9, 0x01,
	0xeb, 0xc0, 0x37, 0x74, 0x2f, 0x6f, 0x40, 0x77, 0x64, 0xaa, 0x68, 0x46, 0x54, 0x5d, 0x84, 0xf5,
	0xe8, 0xbb, 0x83, 0xfc, 0xeb, 0x4d, 0xc0, 0x3e, 0x72, 0xb9, 0xf8, 0x6c, 0xe6, 0xe6, 0x12, 0x1d,
	0x6a, 0x05, 0xe4, 0xc2, 0xcc, 0xa7, 0x4f, 0x74, 0x88, 0x23, 0xd4, 0x91, 0xca, 0xd0, 0xb8, 0x86,
	0xa6, 0xdb, 0x9c, 0x0f, 0x3d, 0xa2, 0x34, 0x83, 0x27, 0x95, 0x76, 0x3e, 0x44, 0x3d, 0x49, 0xf9,
	0x94, 0x25, 0x35, 0x50, 0x09, 0x41, 0xdb, 0x54, 0x43, 0x46, 0x3a, 0xd6, 0x0a, 0x7e, 0x80, 0xba,
	0xf6, 0x00, 0xe3, 0x99, 0x69, 0xaa, 0x4b, 0x6e, 0x1b, 0xe1, 0x0d, 0xcf, 0xf0, 0x1e, 0xba, 0x23,
	0xd9, 0x47, 0xc5, 0x6a, 0x60, 0x59, 0x42, 0x6d, 0x57, 0x5d, 0xd2, 0x5b, 0x6b, 0xfb, 0x30, 0xfa,
	0xe6, 0xa0, 0xdd, 0xcd, 0x88, 0x78, 0x07, 0x79, 0xf6, 0x55, 0xcb, 0x60, 0x13, 0x4d, 0xa1, 0x9f,
	0xb2, 0x5b, 0xa6, 0xc3, 0x8d, 0x4b, 0xe8, 0x6e, 0x5e, 0xc2, 0xeb, 0x86, 0xda, 0xff, 0x18, 0xba,
	0xd4, 0x13, 0x6f, 0x73, 0x4f, 0x0e, 0xf6, 0x7e, 0x36, 0xa1, 0x73, 0xd6, 0x84, 0xce, 0xef, 0x26,
	0x74, 0xbe, 0x2e, 0xc3, 0xd6, 0xd9, 0x32, 0x6c, 0xfd, 0x5a, 0x86, 0xad, 0x77, 0x5b, 0xab, 0xa1,
	0x9d, 0x74, 0xcc, 0x9f, 0xf5, 0xfc, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xca, 0xa2, 0x56, 0x8b,
	0xa3, 0x03, 0x00, 0x00,
}
