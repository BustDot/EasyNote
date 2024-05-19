// Code generated by Kitex v0.9.1. DO NOT EDIT.

package noteservice

import (
	"context"
	note "easy_note/kitex_gen/note/note"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"CreateNote": kitex.NewMethodInfo(
		createNoteHandler,
		newNoteServiceCreateNoteArgs,
		newNoteServiceCreateNoteResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"DeleteNote": kitex.NewMethodInfo(
		deleteNoteHandler,
		newNoteServiceDeleteNoteArgs,
		newNoteServiceDeleteNoteResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"UpdateNote": kitex.NewMethodInfo(
		updateNoteHandler,
		newNoteServiceUpdateNoteArgs,
		newNoteServiceUpdateNoteResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"MGetNote": kitex.NewMethodInfo(
		mGetNoteHandler,
		newNoteServiceMGetNoteArgs,
		newNoteServiceMGetNoteResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"QueryNote": kitex.NewMethodInfo(
		queryNoteHandler,
		newNoteServiceQueryNoteArgs,
		newNoteServiceQueryNoteResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	noteServiceServiceInfo                = NewServiceInfo()
	noteServiceServiceInfoForClient       = NewServiceInfoForClient()
	noteServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return noteServiceServiceInfo
}

// for client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return noteServiceServiceInfoForStreamClient
}

// for stream client
func serviceInfoForClient() *kitex.ServiceInfo {
	return noteServiceServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "NoteService"
	handlerType := (*note.NoteService)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "note",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.9.1",
		Extra:           extra,
	}
	return svcInfo
}

func createNoteHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*note.NoteServiceCreateNoteArgs)
	realResult := result.(*note.NoteServiceCreateNoteResult)
	success, err := handler.(note.NoteService).CreateNote(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newNoteServiceCreateNoteArgs() interface{} {
	return note.NewNoteServiceCreateNoteArgs()
}

func newNoteServiceCreateNoteResult() interface{} {
	return note.NewNoteServiceCreateNoteResult()
}

func deleteNoteHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*note.NoteServiceDeleteNoteArgs)
	realResult := result.(*note.NoteServiceDeleteNoteResult)
	success, err := handler.(note.NoteService).DeleteNote(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newNoteServiceDeleteNoteArgs() interface{} {
	return note.NewNoteServiceDeleteNoteArgs()
}

func newNoteServiceDeleteNoteResult() interface{} {
	return note.NewNoteServiceDeleteNoteResult()
}

func updateNoteHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*note.NoteServiceUpdateNoteArgs)
	realResult := result.(*note.NoteServiceUpdateNoteResult)
	success, err := handler.(note.NoteService).UpdateNote(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newNoteServiceUpdateNoteArgs() interface{} {
	return note.NewNoteServiceUpdateNoteArgs()
}

func newNoteServiceUpdateNoteResult() interface{} {
	return note.NewNoteServiceUpdateNoteResult()
}

func mGetNoteHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*note.NoteServiceMGetNoteArgs)
	realResult := result.(*note.NoteServiceMGetNoteResult)
	success, err := handler.(note.NoteService).MGetNote(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newNoteServiceMGetNoteArgs() interface{} {
	return note.NewNoteServiceMGetNoteArgs()
}

func newNoteServiceMGetNoteResult() interface{} {
	return note.NewNoteServiceMGetNoteResult()
}

func queryNoteHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*note.NoteServiceQueryNoteArgs)
	realResult := result.(*note.NoteServiceQueryNoteResult)
	success, err := handler.(note.NoteService).QueryNote(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newNoteServiceQueryNoteArgs() interface{} {
	return note.NewNoteServiceQueryNoteArgs()
}

func newNoteServiceQueryNoteResult() interface{} {
	return note.NewNoteServiceQueryNoteResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) CreateNote(ctx context.Context, req *note.CreateNoteReq) (r *note.CreateNoteResp, err error) {
	var _args note.NoteServiceCreateNoteArgs
	_args.Req = req
	var _result note.NoteServiceCreateNoteResult
	if err = p.c.Call(ctx, "CreateNote", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DeleteNote(ctx context.Context, req *note.DeleteNoteReq) (r *note.DeleteNoteResp, err error) {
	var _args note.NoteServiceDeleteNoteArgs
	_args.Req = req
	var _result note.NoteServiceDeleteNoteResult
	if err = p.c.Call(ctx, "DeleteNote", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateNote(ctx context.Context, req *note.UpdateNoteReq) (r *note.UpdateNoteResp, err error) {
	var _args note.NoteServiceUpdateNoteArgs
	_args.Req = req
	var _result note.NoteServiceUpdateNoteResult
	if err = p.c.Call(ctx, "UpdateNote", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MGetNote(ctx context.Context, req *note.MGetNoteReq) (r *note.MGetNoteResp, err error) {
	var _args note.NoteServiceMGetNoteArgs
	_args.Req = req
	var _result note.NoteServiceMGetNoteResult
	if err = p.c.Call(ctx, "MGetNote", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) QueryNote(ctx context.Context, req *note.QueryNoteReq) (r *note.QueryNoteResp, err error) {
	var _args note.NoteServiceQueryNoteArgs
	_args.Req = req
	var _result note.NoteServiceQueryNoteResult
	if err = p.c.Call(ctx, "QueryNote", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}