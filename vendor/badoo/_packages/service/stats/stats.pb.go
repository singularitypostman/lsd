// Code generated by protoc-gen-gogo.
// source: stats/stats.proto
// DO NOT EDIT!

/*
Package badoo_service is a generated protocol buffer package.

It is generated from these files:
	stats/stats.proto

It has these top-level messages:
	ResponseMemoryStats
	ResponseProcStats
	ResponseGeneric
	RequestStats
	RequestMemoryStats
	RequestProcStats
	ResponseStats
	RequestVersion
	ResponseVersion
	RequestZlogNotice
	RequestConfigJson
	ResponseConfigJson
	RequestReturnMemoryToOs
*/
package badoo_service

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type RequestMsgid int32

const (
	RequestMsgid_REQUEST_STATS               RequestMsgid = 1
	RequestMsgid_REQUEST_VERSION             RequestMsgid = 2
	RequestMsgid_REQUEST_MEMORY_STATS        RequestMsgid = 3
	RequestMsgid_REQUEST_PROC_STATS          RequestMsgid = 4
	RequestMsgid_REQUEST_ZLOG_NOTICE         RequestMsgid = 5
	RequestMsgid_REQUEST_CONFIG_JSON         RequestMsgid = 6
	RequestMsgid_REQUEST_RETURN_MEMORY_TO_OS RequestMsgid = 7
)

var RequestMsgid_name = map[int32]string{
	1: "REQUEST_STATS",
	2: "REQUEST_VERSION",
	3: "REQUEST_MEMORY_STATS",
	4: "REQUEST_PROC_STATS",
	5: "REQUEST_ZLOG_NOTICE",
	6: "REQUEST_CONFIG_JSON",
	7: "REQUEST_RETURN_MEMORY_TO_OS",
}
var RequestMsgid_value = map[string]int32{
	"REQUEST_STATS":               1,
	"REQUEST_VERSION":             2,
	"REQUEST_MEMORY_STATS":        3,
	"REQUEST_PROC_STATS":          4,
	"REQUEST_ZLOG_NOTICE":         5,
	"REQUEST_CONFIG_JSON":         6,
	"REQUEST_RETURN_MEMORY_TO_OS": 7,
}

func (x RequestMsgid) Enum() *RequestMsgid {
	p := new(RequestMsgid)
	*p = x
	return p
}
func (x RequestMsgid) String() string {
	return proto.EnumName(RequestMsgid_name, int32(x))
}
func (x *RequestMsgid) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(RequestMsgid_value, data, "RequestMsgid")
	if err != nil {
		return err
	}
	*x = RequestMsgid(value)
	return nil
}

type ResponseMsgid int32

const (
	ResponseMsgid_RESPONSE_GENERIC      ResponseMsgid = 1
	ResponseMsgid_RESPONSE_STATS        ResponseMsgid = 2
	ResponseMsgid_RESPONSE_VERSION      ResponseMsgid = 3
	ResponseMsgid_RESPONSE_MEMORY_STATS ResponseMsgid = 4
	ResponseMsgid_RESPONSE_PROC_STATS   ResponseMsgid = 5
	ResponseMsgid_RESPONSE_CONFIG_JSON  ResponseMsgid = 6
)

var ResponseMsgid_name = map[int32]string{
	1: "RESPONSE_GENERIC",
	2: "RESPONSE_STATS",
	3: "RESPONSE_VERSION",
	4: "RESPONSE_MEMORY_STATS",
	5: "RESPONSE_PROC_STATS",
	6: "RESPONSE_CONFIG_JSON",
}
var ResponseMsgid_value = map[string]int32{
	"RESPONSE_GENERIC":      1,
	"RESPONSE_STATS":        2,
	"RESPONSE_VERSION":      3,
	"RESPONSE_MEMORY_STATS": 4,
	"RESPONSE_PROC_STATS":   5,
	"RESPONSE_CONFIG_JSON":  6,
}

func (x ResponseMsgid) Enum() *ResponseMsgid {
	p := new(ResponseMsgid)
	*p = x
	return p
}
func (x ResponseMsgid) String() string {
	return proto.EnumName(ResponseMsgid_name, int32(x))
}
func (x *ResponseMsgid) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ResponseMsgid_value, data, "ResponseMsgid")
	if err != nil {
		return err
	}
	*x = ResponseMsgid(value)
	return nil
}

type Errno int32

const (
	Errno_ERRNO_GENERIC Errno = 1
)

var Errno_name = map[int32]string{
	1: "ERRNO_GENERIC",
}
var Errno_value = map[string]int32{
	"ERRNO_GENERIC": 1,
}

func (x Errno) Enum() *Errno {
	p := new(Errno)
	*p = x
	return p
}
func (x Errno) String() string {
	return proto.EnumName(Errno_name, int32(x))
}
func (x *Errno) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Errno_value, data, "Errno")
	if err != nil {
		return err
	}
	*x = Errno(value)
	return nil
}

type AllocatorType int32

const (
	AllocatorType_LIBC     AllocatorType = 1
	AllocatorType_JEMALLOC AllocatorType = 2
)

var AllocatorType_name = map[int32]string{
	1: "LIBC",
	2: "JEMALLOC",
}
var AllocatorType_value = map[string]int32{
	"LIBC":     1,
	"JEMALLOC": 2,
}

func (x AllocatorType) Enum() *AllocatorType {
	p := new(AllocatorType)
	*p = x
	return p
}
func (x AllocatorType) String() string {
	return proto.EnumName(AllocatorType_name, int32(x))
}
func (x *AllocatorType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(AllocatorType_value, data, "AllocatorType")
	if err != nil {
		return err
	}
	*x = AllocatorType(value)
	return nil
}

type ResponseMemoryStats struct {
	Type             *AllocatorType                `protobuf:"varint,1,req,name=type,enum=badoo.service.AllocatorType" json:"type,omitempty"`
	MemoryUsed       *uint64                       `protobuf:"varint,2,req,name=memory_used" json:"memory_used,omitempty"`
	Jemalloc         *ResponseMemoryStatsJemallocT `protobuf:"bytes,3,opt,name=jemalloc" json:"jemalloc,omitempty"`
	XXX_unrecognized []byte                        `json:"-"`
}

func (m *ResponseMemoryStats) Reset()         { *m = ResponseMemoryStats{} }
func (m *ResponseMemoryStats) String() string { return proto.CompactTextString(m) }
func (*ResponseMemoryStats) ProtoMessage()    {}

func (m *ResponseMemoryStats) GetType() AllocatorType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return AllocatorType_LIBC
}

func (m *ResponseMemoryStats) GetMemoryUsed() uint64 {
	if m != nil && m.MemoryUsed != nil {
		return *m.MemoryUsed
	}
	return 0
}

func (m *ResponseMemoryStats) GetJemalloc() *ResponseMemoryStatsJemallocT {
	if m != nil {
		return m.Jemalloc
	}
	return nil
}

type ResponseMemoryStatsJemallocT struct {
	Allocated        *uint64 `protobuf:"varint,1,opt,name=allocated" json:"allocated,omitempty"`
	Active           *uint64 `protobuf:"varint,2,opt,name=active" json:"active,omitempty"`
	Metadata         *uint64 `protobuf:"varint,3,opt,name=metadata" json:"metadata,omitempty"`
	Resident         *uint64 `protobuf:"varint,4,opt,name=resident" json:"resident,omitempty"`
	Mapped           *uint64 `protobuf:"varint,5,opt,name=mapped" json:"mapped,omitempty"`
	Retained         *uint64 `protobuf:"varint,6,opt,name=retained" json:"retained,omitempty"`
	FullStats        *string `protobuf:"bytes,7,opt,name=full_stats" json:"full_stats,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ResponseMemoryStatsJemallocT) Reset()         { *m = ResponseMemoryStatsJemallocT{} }
func (m *ResponseMemoryStatsJemallocT) String() string { return proto.CompactTextString(m) }
func (*ResponseMemoryStatsJemallocT) ProtoMessage()    {}

func (m *ResponseMemoryStatsJemallocT) GetAllocated() uint64 {
	if m != nil && m.Allocated != nil {
		return *m.Allocated
	}
	return 0
}

func (m *ResponseMemoryStatsJemallocT) GetActive() uint64 {
	if m != nil && m.Active != nil {
		return *m.Active
	}
	return 0
}

func (m *ResponseMemoryStatsJemallocT) GetMetadata() uint64 {
	if m != nil && m.Metadata != nil {
		return *m.Metadata
	}
	return 0
}

func (m *ResponseMemoryStatsJemallocT) GetResident() uint64 {
	if m != nil && m.Resident != nil {
		return *m.Resident
	}
	return 0
}

func (m *ResponseMemoryStatsJemallocT) GetMapped() uint64 {
	if m != nil && m.Mapped != nil {
		return *m.Mapped
	}
	return 0
}

func (m *ResponseMemoryStatsJemallocT) GetRetained() uint64 {
	if m != nil && m.Retained != nil {
		return *m.Retained
	}
	return 0
}

func (m *ResponseMemoryStatsJemallocT) GetFullStats() string {
	if m != nil && m.FullStats != nil {
		return *m.FullStats
	}
	return ""
}

type ResponseProcStats struct {
	Size_            *uint64 `protobuf:"varint,1,req,name=size" json:"size,omitempty"`
	Resident         *uint64 `protobuf:"varint,2,req,name=resident" json:"resident,omitempty"`
	Shared           *uint64 `protobuf:"varint,3,req,name=shared" json:"shared,omitempty"`
	Text             *uint64 `protobuf:"varint,4,req,name=text" json:"text,omitempty"`
	Data             *uint64 `protobuf:"varint,5,req,name=data" json:"data,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ResponseProcStats) Reset()         { *m = ResponseProcStats{} }
func (m *ResponseProcStats) String() string { return proto.CompactTextString(m) }
func (*ResponseProcStats) ProtoMessage()    {}

func (m *ResponseProcStats) GetSize_() uint64 {
	if m != nil && m.Size_ != nil {
		return *m.Size_
	}
	return 0
}

func (m *ResponseProcStats) GetResident() uint64 {
	if m != nil && m.Resident != nil {
		return *m.Resident
	}
	return 0
}

func (m *ResponseProcStats) GetShared() uint64 {
	if m != nil && m.Shared != nil {
		return *m.Shared
	}
	return 0
}

func (m *ResponseProcStats) GetText() uint64 {
	if m != nil && m.Text != nil {
		return *m.Text
	}
	return 0
}

func (m *ResponseProcStats) GetData() uint64 {
	if m != nil && m.Data != nil {
		return *m.Data
	}
	return 0
}

type ResponseGeneric struct {
	ErrorCode        *int32  `protobuf:"zigzag32,1,req,name=error_code" json:"error_code,omitempty"`
	ErrorText        *string `protobuf:"bytes,2,opt,name=error_text" json:"error_text,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ResponseGeneric) Reset()         { *m = ResponseGeneric{} }
func (m *ResponseGeneric) String() string { return proto.CompactTextString(m) }
func (*ResponseGeneric) ProtoMessage()    {}

func (m *ResponseGeneric) GetErrorCode() int32 {
	if m != nil && m.ErrorCode != nil {
		return *m.ErrorCode
	}
	return 0
}

func (m *ResponseGeneric) GetErrorText() string {
	if m != nil && m.ErrorText != nil {
		return *m.ErrorText
	}
	return ""
}

type RequestStats struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *RequestStats) Reset()         { *m = RequestStats{} }
func (m *RequestStats) String() string { return proto.CompactTextString(m) }
func (*RequestStats) ProtoMessage()    {}

type RequestMemoryStats struct {
	Full             *bool  `protobuf:"varint,1,opt,name=full,def=0" json:"full,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *RequestMemoryStats) Reset()         { *m = RequestMemoryStats{} }
func (m *RequestMemoryStats) String() string { return proto.CompactTextString(m) }
func (*RequestMemoryStats) ProtoMessage()    {}

const Default_RequestMemoryStats_Full bool = false

func (m *RequestMemoryStats) GetFull() bool {
	if m != nil && m.Full != nil {
		return *m.Full
	}
	return Default_RequestMemoryStats_Full
}

type RequestProcStats struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *RequestProcStats) Reset()         { *m = RequestProcStats{} }
func (m *RequestProcStats) String() string { return proto.CompactTextString(m) }
func (*RequestProcStats) ProtoMessage()    {}

type ResponseStats struct {
	Uptime            *uint32                   `protobuf:"varint,1,req,name=uptime" json:"uptime,omitempty"`
	RusageSelf        *ResponseStatsRusage      `protobuf:"bytes,2,req,name=rusage_self" json:"rusage_self,omitempty"`
	Ports             []*ResponseStatsPortStats `protobuf:"bytes,3,rep,name=ports" json:"ports,omitempty"`
	Connections       *uint32                   `protobuf:"varint,4,opt,name=connections" json:"connections,omitempty"`
	InitPhaseDuration *uint32                   `protobuf:"varint,5,opt,name=init_phase_duration" json:"init_phase_duration,omitempty"`
	XXX_unrecognized  []byte                    `json:"-"`
}

func (m *ResponseStats) Reset()         { *m = ResponseStats{} }
func (m *ResponseStats) String() string { return proto.CompactTextString(m) }
func (*ResponseStats) ProtoMessage()    {}

func (m *ResponseStats) GetUptime() uint32 {
	if m != nil && m.Uptime != nil {
		return *m.Uptime
	}
	return 0
}

func (m *ResponseStats) GetRusageSelf() *ResponseStatsRusage {
	if m != nil {
		return m.RusageSelf
	}
	return nil
}

func (m *ResponseStats) GetPorts() []*ResponseStatsPortStats {
	if m != nil {
		return m.Ports
	}
	return nil
}

func (m *ResponseStats) GetConnections() uint32 {
	if m != nil && m.Connections != nil {
		return *m.Connections
	}
	return 0
}

func (m *ResponseStats) GetInitPhaseDuration() uint32 {
	if m != nil && m.InitPhaseDuration != nil {
		return *m.InitPhaseDuration
	}
	return 0
}

type ResponseStatsRusage struct {
	RuUtime          *float32 `protobuf:"fixed32,1,req,name=ru_utime" json:"ru_utime,omitempty"`
	RuStime          *float32 `protobuf:"fixed32,2,req,name=ru_stime" json:"ru_stime,omitempty"`
	RuMaxrss         *uint64  `protobuf:"varint,3,opt,name=ru_maxrss" json:"ru_maxrss,omitempty"`
	RuMinflt         *uint64  `protobuf:"varint,4,opt,name=ru_minflt" json:"ru_minflt,omitempty"`
	RuMajflt         *uint64  `protobuf:"varint,5,opt,name=ru_majflt" json:"ru_majflt,omitempty"`
	RuInblock        *uint64  `protobuf:"varint,6,opt,name=ru_inblock" json:"ru_inblock,omitempty"`
	RuOublock        *uint64  `protobuf:"varint,7,opt,name=ru_oublock" json:"ru_oublock,omitempty"`
	RuNvcsw          *uint64  `protobuf:"varint,8,opt,name=ru_nvcsw" json:"ru_nvcsw,omitempty"`
	RuNivcsw         *uint64  `protobuf:"varint,9,opt,name=ru_nivcsw" json:"ru_nivcsw,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *ResponseStatsRusage) Reset()         { *m = ResponseStatsRusage{} }
func (m *ResponseStatsRusage) String() string { return proto.CompactTextString(m) }
func (*ResponseStatsRusage) ProtoMessage()    {}

func (m *ResponseStatsRusage) GetRuUtime() float32 {
	if m != nil && m.RuUtime != nil {
		return *m.RuUtime
	}
	return 0
}

func (m *ResponseStatsRusage) GetRuStime() float32 {
	if m != nil && m.RuStime != nil {
		return *m.RuStime
	}
	return 0
}

func (m *ResponseStatsRusage) GetRuMaxrss() uint64 {
	if m != nil && m.RuMaxrss != nil {
		return *m.RuMaxrss
	}
	return 0
}

func (m *ResponseStatsRusage) GetRuMinflt() uint64 {
	if m != nil && m.RuMinflt != nil {
		return *m.RuMinflt
	}
	return 0
}

func (m *ResponseStatsRusage) GetRuMajflt() uint64 {
	if m != nil && m.RuMajflt != nil {
		return *m.RuMajflt
	}
	return 0
}

func (m *ResponseStatsRusage) GetRuInblock() uint64 {
	if m != nil && m.RuInblock != nil {
		return *m.RuInblock
	}
	return 0
}

func (m *ResponseStatsRusage) GetRuOublock() uint64 {
	if m != nil && m.RuOublock != nil {
		return *m.RuOublock
	}
	return 0
}

func (m *ResponseStatsRusage) GetRuNvcsw() uint64 {
	if m != nil && m.RuNvcsw != nil {
		return *m.RuNvcsw
	}
	return 0
}

func (m *ResponseStatsRusage) GetRuNivcsw() uint64 {
	if m != nil && m.RuNivcsw != nil {
		return *m.RuNivcsw
	}
	return 0
}

type ResponseStatsPortStats struct {
	Proto            *string                                `protobuf:"bytes,1,req,name=proto" json:"proto,omitempty"`
	Address          *string                                `protobuf:"bytes,2,opt,name=address" json:"address,omitempty"`
	LqCur            *uint32                                `protobuf:"varint,3,opt,name=lq_cur" json:"lq_cur,omitempty"`
	LqMax            *uint32                                `protobuf:"varint,4,opt,name=lq_max" json:"lq_max,omitempty"`
	ConnCur          *uint64                                `protobuf:"varint,5,opt,name=conn_cur" json:"conn_cur,omitempty"`
	ConnTotal        *uint64                                `protobuf:"varint,6,opt,name=conn_total" json:"conn_total,omitempty"`
	ConnAborted      *uint64                                `protobuf:"varint,11,opt,name=conn_aborted" json:"conn_aborted,omitempty"`
	Requests         *uint64                                `protobuf:"varint,7,opt,name=requests" json:"requests,omitempty"`
	BytesRead        *uint64                                `protobuf:"varint,8,opt,name=bytes_read" json:"bytes_read,omitempty"`
	BytesWritten     *uint64                                `protobuf:"varint,9,opt,name=bytes_written" json:"bytes_written,omitempty"`
	RequestStats     []*ResponseStatsPortStatsRequestStatsT `protobuf:"bytes,10,rep,name=request_stats" json:"request_stats,omitempty"`
	XXX_unrecognized []byte                                 `json:"-"`
}

func (m *ResponseStatsPortStats) Reset()         { *m = ResponseStatsPortStats{} }
func (m *ResponseStatsPortStats) String() string { return proto.CompactTextString(m) }
func (*ResponseStatsPortStats) ProtoMessage()    {}

func (m *ResponseStatsPortStats) GetProto() string {
	if m != nil && m.Proto != nil {
		return *m.Proto
	}
	return ""
}

func (m *ResponseStatsPortStats) GetAddress() string {
	if m != nil && m.Address != nil {
		return *m.Address
	}
	return ""
}

func (m *ResponseStatsPortStats) GetLqCur() uint32 {
	if m != nil && m.LqCur != nil {
		return *m.LqCur
	}
	return 0
}

func (m *ResponseStatsPortStats) GetLqMax() uint32 {
	if m != nil && m.LqMax != nil {
		return *m.LqMax
	}
	return 0
}

func (m *ResponseStatsPortStats) GetConnCur() uint64 {
	if m != nil && m.ConnCur != nil {
		return *m.ConnCur
	}
	return 0
}

func (m *ResponseStatsPortStats) GetConnTotal() uint64 {
	if m != nil && m.ConnTotal != nil {
		return *m.ConnTotal
	}
	return 0
}

func (m *ResponseStatsPortStats) GetConnAborted() uint64 {
	if m != nil && m.ConnAborted != nil {
		return *m.ConnAborted
	}
	return 0
}

func (m *ResponseStatsPortStats) GetRequests() uint64 {
	if m != nil && m.Requests != nil {
		return *m.Requests
	}
	return 0
}

func (m *ResponseStatsPortStats) GetBytesRead() uint64 {
	if m != nil && m.BytesRead != nil {
		return *m.BytesRead
	}
	return 0
}

func (m *ResponseStatsPortStats) GetBytesWritten() uint64 {
	if m != nil && m.BytesWritten != nil {
		return *m.BytesWritten
	}
	return 0
}

func (m *ResponseStatsPortStats) GetRequestStats() []*ResponseStatsPortStatsRequestStatsT {
	if m != nil {
		return m.RequestStats
	}
	return nil
}

type ResponseStatsPortStatsRequestStatsT struct {
	Name             *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Count            *uint64 `protobuf:"varint,2,req,name=count" json:"count,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ResponseStatsPortStatsRequestStatsT) Reset()         { *m = ResponseStatsPortStatsRequestStatsT{} }
func (m *ResponseStatsPortStatsRequestStatsT) String() string { return proto.CompactTextString(m) }
func (*ResponseStatsPortStatsRequestStatsT) ProtoMessage()    {}

func (m *ResponseStatsPortStatsRequestStatsT) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *ResponseStatsPortStatsRequestStatsT) GetCount() uint64 {
	if m != nil && m.Count != nil {
		return *m.Count
	}
	return 0
}

type RequestVersion struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *RequestVersion) Reset()         { *m = RequestVersion{} }
func (m *RequestVersion) String() string { return proto.CompactTextString(m) }
func (*RequestVersion) ProtoMessage()    {}

// This message exists in misc-proto.git too. If you change something,
// keep it in sync.
type ResponseVersion struct {
	Version          *string `protobuf:"bytes,1,req,name=version" json:"version,omitempty"`
	BuildId          *string `protobuf:"bytes,2,opt,name=build_id" json:"build_id,omitempty"`
	AutoBuildTag     *string `protobuf:"bytes,3,opt,name=auto_build_tag" json:"auto_build_tag,omitempty"`
	BuildDate        *string `protobuf:"bytes,4,opt,name=build_date" json:"build_date,omitempty"`
	BuildHost        *string `protobuf:"bytes,5,opt,name=build_host" json:"build_host,omitempty"`
	BuildCc          *string `protobuf:"bytes,6,opt,name=build_cc" json:"build_cc,omitempty"`
	BuildConfigure   *string `protobuf:"bytes,7,opt,name=build_configure" json:"build_configure,omitempty"`
	VcsType          *string `protobuf:"bytes,8,opt,name=vcs_type" json:"vcs_type,omitempty"`
	VcsBasename      *string `protobuf:"bytes,9,opt,name=vcs_basename" json:"vcs_basename,omitempty"`
	VcsNum           *string `protobuf:"bytes,10,opt,name=vcs_num" json:"vcs_num,omitempty"`
	VcsDate          *string `protobuf:"bytes,11,opt,name=vcs_date" json:"vcs_date,omitempty"`
	VcsBranch        *string `protobuf:"bytes,12,opt,name=vcs_branch" json:"vcs_branch,omitempty"`
	VcsTag           *string `protobuf:"bytes,13,opt,name=vcs_tag" json:"vcs_tag,omitempty"`
	VcsTick          *string `protobuf:"bytes,14,opt,name=vcs_tick" json:"vcs_tick,omitempty"`
	VcsFullHash      *string `protobuf:"bytes,15,opt,name=vcs_full_hash" json:"vcs_full_hash,omitempty"`
	VcsShortHash     *string `protobuf:"bytes,16,opt,name=vcs_short_hash" json:"vcs_short_hash,omitempty"`
	VcsWcModified    *string `protobuf:"bytes,17,opt,name=vcs_wc_modified" json:"vcs_wc_modified,omitempty"`
	Maintainer       *string `protobuf:"bytes,18,opt,name=maintainer" json:"maintainer,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ResponseVersion) Reset()         { *m = ResponseVersion{} }
func (m *ResponseVersion) String() string { return proto.CompactTextString(m) }
func (*ResponseVersion) ProtoMessage()    {}

func (m *ResponseVersion) GetVersion() string {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return ""
}

func (m *ResponseVersion) GetBuildId() string {
	if m != nil && m.BuildId != nil {
		return *m.BuildId
	}
	return ""
}

func (m *ResponseVersion) GetAutoBuildTag() string {
	if m != nil && m.AutoBuildTag != nil {
		return *m.AutoBuildTag
	}
	return ""
}

func (m *ResponseVersion) GetBuildDate() string {
	if m != nil && m.BuildDate != nil {
		return *m.BuildDate
	}
	return ""
}

func (m *ResponseVersion) GetBuildHost() string {
	if m != nil && m.BuildHost != nil {
		return *m.BuildHost
	}
	return ""
}

func (m *ResponseVersion) GetBuildCc() string {
	if m != nil && m.BuildCc != nil {
		return *m.BuildCc
	}
	return ""
}

func (m *ResponseVersion) GetBuildConfigure() string {
	if m != nil && m.BuildConfigure != nil {
		return *m.BuildConfigure
	}
	return ""
}

func (m *ResponseVersion) GetVcsType() string {
	if m != nil && m.VcsType != nil {
		return *m.VcsType
	}
	return ""
}

func (m *ResponseVersion) GetVcsBasename() string {
	if m != nil && m.VcsBasename != nil {
		return *m.VcsBasename
	}
	return ""
}

func (m *ResponseVersion) GetVcsNum() string {
	if m != nil && m.VcsNum != nil {
		return *m.VcsNum
	}
	return ""
}

func (m *ResponseVersion) GetVcsDate() string {
	if m != nil && m.VcsDate != nil {
		return *m.VcsDate
	}
	return ""
}

func (m *ResponseVersion) GetVcsBranch() string {
	if m != nil && m.VcsBranch != nil {
		return *m.VcsBranch
	}
	return ""
}

func (m *ResponseVersion) GetVcsTag() string {
	if m != nil && m.VcsTag != nil {
		return *m.VcsTag
	}
	return ""
}

func (m *ResponseVersion) GetVcsTick() string {
	if m != nil && m.VcsTick != nil {
		return *m.VcsTick
	}
	return ""
}

func (m *ResponseVersion) GetVcsFullHash() string {
	if m != nil && m.VcsFullHash != nil {
		return *m.VcsFullHash
	}
	return ""
}

func (m *ResponseVersion) GetVcsShortHash() string {
	if m != nil && m.VcsShortHash != nil {
		return *m.VcsShortHash
	}
	return ""
}

func (m *ResponseVersion) GetVcsWcModified() string {
	if m != nil && m.VcsWcModified != nil {
		return *m.VcsWcModified
	}
	return ""
}

func (m *ResponseVersion) GetMaintainer() string {
	if m != nil && m.Maintainer != nil {
		return *m.Maintainer
	}
	return ""
}

type RequestZlogNotice struct {
	Text             *string `protobuf:"bytes,1,req,name=text" json:"text,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *RequestZlogNotice) Reset()         { *m = RequestZlogNotice{} }
func (m *RequestZlogNotice) String() string { return proto.CompactTextString(m) }
func (*RequestZlogNotice) ProtoMessage()    {}

func (m *RequestZlogNotice) GetText() string {
	if m != nil && m.Text != nil {
		return *m.Text
	}
	return ""
}

type RequestConfigJson struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *RequestConfigJson) Reset()         { *m = RequestConfigJson{} }
func (m *RequestConfigJson) String() string { return proto.CompactTextString(m) }
func (*RequestConfigJson) ProtoMessage()    {}

type ResponseConfigJson struct {
	Json             *string `protobuf:"bytes,1,req,name=json" json:"json,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ResponseConfigJson) Reset()         { *m = ResponseConfigJson{} }
func (m *ResponseConfigJson) String() string { return proto.CompactTextString(m) }
func (*ResponseConfigJson) ProtoMessage()    {}

func (m *ResponseConfigJson) GetJson() string {
	if m != nil && m.Json != nil {
		return *m.Json
	}
	return ""
}

type RequestReturnMemoryToOs struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *RequestReturnMemoryToOs) Reset()         { *m = RequestReturnMemoryToOs{} }
func (m *RequestReturnMemoryToOs) String() string { return proto.CompactTextString(m) }
func (*RequestReturnMemoryToOs) ProtoMessage()    {}

func init() {
	proto.RegisterEnum("badoo.service.RequestMsgid", RequestMsgid_name, RequestMsgid_value)
	proto.RegisterEnum("badoo.service.ResponseMsgid", ResponseMsgid_name, ResponseMsgid_value)
	proto.RegisterEnum("badoo.service.Errno", Errno_name, Errno_value)
	proto.RegisterEnum("badoo.service.AllocatorType", AllocatorType_name, AllocatorType_value)
}
