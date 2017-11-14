// Code generated by protoc-gen-gpbrpc-go.
// source: stats/stats.proto
// DO NOT EDIT!

package badoo_service

import "github.com/gogo/protobuf/proto"
import "badoo/_packages/gpbrpc"
import "fmt"

type GpbrpcType struct {
}

var Gpbrpc GpbrpcType

var RequestMsgid_gpbrpc_name = map[uint32]string{
	1: "request_stats",
	2: "request_version",
	3: "request_memory_stats",
	4: "request_proc_stats",
	5: "request_zlog_notice",
	6: "request_config_json",
	7: "request_return_memory_to_os",
}

var RequestMsgid_gpbrpc_value = map[string]uint32{
	"request_stats":               1,
	"request_version":             2,
	"request_memory_stats":        3,
	"request_proc_stats":          4,
	"request_zlog_notice":         5,
	"request_config_json":         6,
	"request_return_memory_to_os": 7,
}

var ResponseMsgid_gpbrpc_name = map[uint32]string{
	1: "response_generic",
	2: "response_stats",
	3: "response_version",
	4: "response_memory_stats",
	5: "response_proc_stats",
	6: "response_config_json",
}

var ResponseMsgid_gpbrpc_value = map[string]uint32{
	"response_generic":      1,
	"response_stats":        2,
	"response_version":      3,
	"response_memory_stats": 4,
	"response_proc_stats":   5,
	"response_config_json":  6,
}

func (GpbrpcType) GetRequestMsgid(msg proto.Message) uint32 {
	switch msg.(type) {
	case *RequestStats:
		return uint32(RequestMsgid_REQUEST_STATS)
	case *RequestVersion:
		return uint32(RequestMsgid_REQUEST_VERSION)
	case *RequestMemoryStats:
		return uint32(RequestMsgid_REQUEST_MEMORY_STATS)
	case *RequestProcStats:
		return uint32(RequestMsgid_REQUEST_PROC_STATS)
	case *RequestZlogNotice:
		return uint32(RequestMsgid_REQUEST_ZLOG_NOTICE)
	case *RequestConfigJson:
		return uint32(RequestMsgid_REQUEST_CONFIG_JSON)
	case *RequestReturnMemoryToOs:
		return uint32(RequestMsgid_REQUEST_RETURN_MEMORY_TO_OS)
	default:
		panic("you gave me the wrong message")
	}
}

func (GpbrpcType) GetRequestNameToIdMap() map[string]uint32 {
	return RequestMsgid_gpbrpc_value
}

func (GpbrpcType) GetRequestIdToNameMap() map[uint32]string {
	return RequestMsgid_gpbrpc_name
}

func (GpbrpcType) GetResponseNameToIdMap() map[string]uint32 {
	return ResponseMsgid_gpbrpc_value
}

func (GpbrpcType) GetResponseIdToNameMap() map[uint32]string {
	return ResponseMsgid_gpbrpc_name
}

func (GpbrpcType) GetResponseMsgid(msg proto.Message) uint32 {
	switch msg.(type) {
	case *ResponseGeneric:
		return uint32(ResponseMsgid_RESPONSE_GENERIC)
	case *ResponseStats:
		return uint32(ResponseMsgid_RESPONSE_STATS)
	case *ResponseVersion:
		return uint32(ResponseMsgid_RESPONSE_VERSION)
	case *ResponseMemoryStats:
		return uint32(ResponseMsgid_RESPONSE_MEMORY_STATS)
	case *ResponseProcStats:
		return uint32(ResponseMsgid_RESPONSE_PROC_STATS)
	case *ResponseConfigJson:
		return uint32(ResponseMsgid_RESPONSE_CONFIG_JSON)
	default:
		panic("you gave me the wrong message")
	}
}

func (GpbrpcType) GetPackageName() string {
	return "badoo.service"
}

func (GpbrpcType) GetRequestMsg(request_msgid uint32) proto.Message {
	switch RequestMsgid(request_msgid) {
	case RequestMsgid_REQUEST_STATS:
		return &RequestStats{}
	case RequestMsgid_REQUEST_VERSION:
		return &RequestVersion{}
	case RequestMsgid_REQUEST_MEMORY_STATS:
		return &RequestMemoryStats{}
	case RequestMsgid_REQUEST_PROC_STATS:
		return &RequestProcStats{}
	case RequestMsgid_REQUEST_ZLOG_NOTICE:
		return &RequestZlogNotice{}
	case RequestMsgid_REQUEST_CONFIG_JSON:
		return &RequestConfigJson{}
	case RequestMsgid_REQUEST_RETURN_MEMORY_TO_OS:
		return &RequestReturnMemoryToOs{}
	default:
		return nil
	}
}

func (GpbrpcType) GetResponseMsg(response_msgid uint32) proto.Message {
	switch ResponseMsgid(response_msgid) {
	case ResponseMsgid_RESPONSE_GENERIC:
		return &ResponseGeneric{}
	case ResponseMsgid_RESPONSE_STATS:
		return &ResponseStats{}
	case ResponseMsgid_RESPONSE_VERSION:
		return &ResponseVersion{}
	case ResponseMsgid_RESPONSE_MEMORY_STATS:
		return &ResponseMemoryStats{}
	case ResponseMsgid_RESPONSE_PROC_STATS:
		return &ResponseProcStats{}
	case ResponseMsgid_RESPONSE_CONFIG_JSON:
		return &ResponseConfigJson{}
	default:
		return nil
	}
}

type GpbrpcInterface interface {
	RequestStats(rctx gpbrpc.RequestT, request *RequestStats) gpbrpc.ResultT
	RequestVersion(rctx gpbrpc.RequestT, request *RequestVersion) gpbrpc.ResultT
	RequestMemoryStats(rctx gpbrpc.RequestT, request *RequestMemoryStats) gpbrpc.ResultT
	RequestProcStats(rctx gpbrpc.RequestT, request *RequestProcStats) gpbrpc.ResultT
	RequestZlogNotice(rctx gpbrpc.RequestT, request *RequestZlogNotice) gpbrpc.ResultT
	RequestConfigJson(rctx gpbrpc.RequestT, request *RequestConfigJson) gpbrpc.ResultT
	RequestReturnMemoryToOs(rctx gpbrpc.RequestT, request *RequestReturnMemoryToOs) gpbrpc.ResultT
}

func (GpbrpcType) Dispatch(rctx gpbrpc.RequestT, abstract_service interface{}) gpbrpc.ResultT {

	service := abstract_service.(GpbrpcInterface)

	switch RequestMsgid(rctx.MessageId) {
	case RequestMsgid_REQUEST_STATS:
		r := rctx.Message.(*RequestStats)
		return service.RequestStats(rctx, r)
	case RequestMsgid_REQUEST_VERSION:
		r := rctx.Message.(*RequestVersion)
		return service.RequestVersion(rctx, r)
	case RequestMsgid_REQUEST_MEMORY_STATS:
		r := rctx.Message.(*RequestMemoryStats)
		return service.RequestMemoryStats(rctx, r)
	case RequestMsgid_REQUEST_PROC_STATS:
		r := rctx.Message.(*RequestProcStats)
		return service.RequestProcStats(rctx, r)
	case RequestMsgid_REQUEST_ZLOG_NOTICE:
		r := rctx.Message.(*RequestZlogNotice)
		return service.RequestZlogNotice(rctx, r)
	case RequestMsgid_REQUEST_CONFIG_JSON:
		r := rctx.Message.(*RequestConfigJson)
		return service.RequestConfigJson(rctx, r)
	case RequestMsgid_REQUEST_RETURN_MEMORY_TO_OS:
		r := rctx.Message.(*RequestReturnMemoryToOs)
		return service.RequestReturnMemoryToOs(rctx, r)
	default:
		panic("screw you")
	}
}

var okResult = gpbrpc.Result(&ResponseGeneric{ErrorCode: proto.Int32(0)})

func (GpbrpcType) OK(args ...interface{}) gpbrpc.ResultT {
	if len(args) == 0 {
		return okResult
	}
	return gpbrpc.Result(&ResponseGeneric{ErrorCode: proto.Int32(0),
		ErrorText: proto.String(fmt.Sprint(args...))})
}

func (GpbrpcType) ErrorGeneric(args ...interface{}) gpbrpc.ResultT {
	return gpbrpc.Result(&ResponseGeneric{ErrorCode: proto.Int32(-int32(Errno_ERRNO_GENERIC)),
		ErrorText: proto.String(fmt.Sprint(args...))})
}

/*
	func ($receiver$) RequestStats(rctx gpbrpc.RequestT, request *$proto$.RequestStats) gpbrpc.ResultT {
		return $proto$.Gpbrpc.ErrorGeneric("not implemented")
	}

	func ($receiver$) RequestVersion(rctx gpbrpc.RequestT, request *$proto$.RequestVersion) gpbrpc.ResultT {
		return $proto$.Gpbrpc.ErrorGeneric("not implemented")
	}

	func ($receiver$) RequestMemoryStats(rctx gpbrpc.RequestT, request *$proto$.RequestMemoryStats) gpbrpc.ResultT {
		return $proto$.Gpbrpc.ErrorGeneric("not implemented")
	}

	func ($receiver$) RequestProcStats(rctx gpbrpc.RequestT, request *$proto$.RequestProcStats) gpbrpc.ResultT {
		return $proto$.Gpbrpc.ErrorGeneric("not implemented")
	}

	func ($receiver$) RequestZlogNotice(rctx gpbrpc.RequestT, request *$proto$.RequestZlogNotice) gpbrpc.ResultT {
		return $proto$.Gpbrpc.ErrorGeneric("not implemented")
	}

	func ($receiver$) RequestConfigJson(rctx gpbrpc.RequestT, request *$proto$.RequestConfigJson) gpbrpc.ResultT {
		return $proto$.Gpbrpc.ErrorGeneric("not implemented")
	}

	func ($receiver$) RequestReturnMemoryToOs(rctx gpbrpc.RequestT, request *$proto$.RequestReturnMemoryToOs) gpbrpc.ResultT {
		return $proto$.Gpbrpc.ErrorGeneric("not implemented")
	}

*/
