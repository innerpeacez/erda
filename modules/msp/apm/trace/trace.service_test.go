// Copyright (c) 2021 Terminus, Inc.
//
// This program is free software: you can use, redistribute, and/or modify
// it under the terms of the GNU Affero General Public License, version 3
// or later ("AGPL"), as published by the Free Software Foundation.
//
// This program is distributed in the hope that it will be useful, but WITHOUT
// ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or
// FITNESS FOR A PARTICULAR PURPOSE.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package trace

import (
	"context"
	"encoding/json"
	"reflect"
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"

	"github.com/erda-project/erda-infra/base/servicehub"
	"github.com/erda-project/erda-proto-go/msp/apm/trace/pb"
	"github.com/erda-project/erda/modules/msp/apm/trace/db"
)

func Test_traceService_GetSpans(t *testing.T) {
	type args struct {
		ctx context.Context
		req *pb.GetSpansRequest
	}
	tests := []struct {
		name     string
		service  string
		config   string
		args     args
		wantResp *pb.GetSpansResponse
		wantErr  bool
	}{
		// TODO: Add test cases.
		//		{
		//			"case 1",
		//			"erda.msp.apm.trace.TraceService",
		//			`
		//erda.msp.apm.trace:
		//`,
		//			args{
		//				context.TODO(),
		//				&pb.GetSpansRequest{
		//					// TODO: setup fields
		//				},
		//			},
		//			&pb.GetSpansResponse{
		//				// TODO: setup fields.
		//			},
		//			false,
		//		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hub := servicehub.New()
			events := hub.Events()
			go func() {
				hub.RunWithOptions(&servicehub.RunOptions{Content: tt.config})
			}()
			err := <-events.Started()
			if err != nil {
				t.Error(err)
				return
			}
			srv := hub.Service(tt.service).(pb.TraceServiceServer)
			got, err := srv.GetSpans(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("traceService.GetSpans() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.wantResp) {
				t.Errorf("traceService.GetSpans() = %v, want %v", got, tt.wantResp)
			}
		})
	}
}

func Test_traceService_GetTraces(t *testing.T) {
	type args struct {
		ctx context.Context
		req *pb.GetTracesRequest
	}
	tests := []struct {
		name     string
		service  string
		config   string
		args     args
		wantResp *pb.GetTracesResponse
		wantErr  bool
	}{
		// TODO: Add test cases.
		//		{
		//			"case 1",
		//			"erda.msp.apm.trace.TraceService",
		//			`
		//erda.msp.apm.trace:
		//`,
		//			args{
		//				context.TODO(),
		//				&pb.GetTracesRequest{
		//					// TODO: setup fields
		//				},
		//			},
		//			&pb.GetTracesResponse{
		//				// TODO: setup fields.
		//			},
		//			false,
		//		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hub := servicehub.New()
			events := hub.Events()
			go func() {
				hub.RunWithOptions(&servicehub.RunOptions{Content: tt.config})
			}()
			err := <-events.Started()
			if err != nil {
				t.Error(err)
				return
			}
			srv := hub.Service(tt.service).(pb.TraceServiceServer)
			got, err := srv.GetTraces(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("traceService.GetTraces() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.wantResp) {
				t.Errorf("traceService.GetTraces() = %v, want %v", got, tt.wantResp)
			}
		})
	}
}

func Test_traceService_GetTraceDebugHistories(t *testing.T) {
	type args struct {
		ctx context.Context
		req *pb.GetTraceDebugHistoriesRequest
	}
	tests := []struct {
		name     string
		service  string
		config   string
		args     args
		wantResp *pb.GetTraceDebugHistoriesResponse
		wantErr  bool
	}{
		// TODO: Add test cases.
		//		{
		//			"case 1",
		//			"erda.msp.apm.trace.TraceService",
		//			`
		//erda.msp.apm.trace:
		//`,
		//			args{
		//				context.TODO(),
		//				&pb.GetTraceDebugHistoriesRequest{
		//					// TODO: setup fields
		//				},
		//			},
		//			&pb.GetTraceDebugHistoriesResponse{
		//				// TODO: setup fields.
		//			},
		//			false,
		//		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hub := servicehub.New()
			events := hub.Events()
			go func() {
				hub.RunWithOptions(&servicehub.RunOptions{Content: tt.config})
			}()
			err := <-events.Started()
			if err != nil {
				t.Error(err)
				return
			}
			srv := hub.Service(tt.service).(pb.TraceServiceServer)
			got, err := srv.GetTraceDebugHistories(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("traceService.GetTraceDebugHistories() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.wantResp) {
				t.Errorf("traceService.GetTraceDebugHistories() = %v, want %v", got, tt.wantResp)
			}
		})
	}
}

func Test_traceService_GetTraceDebugByRequestID(t *testing.T) {
	type args struct {
		ctx context.Context
		req *pb.GetTraceDebugRequest
	}
	tests := []struct {
		name     string
		service  string
		config   string
		args     args
		wantResp *pb.GetTraceDebugResponse
		wantErr  bool
	}{
		// TODO: Add test cases.
		//		{
		//			"case 1",
		//			"erda.msp.apm.trace.TraceService",
		//			`
		//erda.msp.apm.trace:
		//`,
		//			args{
		//				context.TODO(),
		//				&pb.GetTraceDebugRequest{
		//					// TODO: setup fields
		//				},
		//			},
		//			&pb.GetTraceDebugResponse{
		//				// TODO: setup fields.
		//			},
		//			false,
		//		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hub := servicehub.New()
			events := hub.Events()
			go func() {
				hub.RunWithOptions(&servicehub.RunOptions{Content: tt.config})
			}()
			err := <-events.Started()
			if err != nil {
				t.Error(err)
				return
			}
			srv := hub.Service(tt.service).(pb.TraceServiceServer)
			got, err := srv.GetTraceDebugByRequestID(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("traceService.GetTraceDebugByRequestID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.wantResp) {
				t.Errorf("traceService.GetTraceDebugByRequestID() = %v, want %v", got, tt.wantResp)
			}
		})
	}
}

func Test_traceService_CreateTraceDebug(t *testing.T) {
	type args struct {
		ctx context.Context
		req *pb.CreateTraceDebugRequest
	}
	tests := []struct {
		name     string
		service  string
		config   string
		args     args
		wantResp *pb.CreateTraceDebugResponse
		wantErr  bool
	}{
		// TODO: Add test cases.
		//		{
		//			"case 1",
		//			"erda.msp.apm.trace.TraceService",
		//			`
		//erda.msp.apm.trace:
		//`,
		//			args{
		//				context.TODO(),
		//				&pb.CreateTraceDebugRequest{
		//					// TODO: setup fields
		//				},
		//			},
		//			&pb.CreateTraceDebugResponse{
		//				// TODO: setup fields.
		//			},
		//			false,
		//		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hub := servicehub.New()
			events := hub.Events()
			go func() {
				hub.RunWithOptions(&servicehub.RunOptions{Content: tt.config})
			}()
			err := <-events.Started()
			if err != nil {
				t.Error(err)
				return
			}
			srv := hub.Service(tt.service).(pb.TraceServiceServer)
			got, err := srv.CreateTraceDebug(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("traceService.CreateTraceDebug() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.wantResp) {
				t.Errorf("traceService.CreateTraceDebug() = %v, want %v", got, tt.wantResp)
			}
		})
	}
}

func Test_traceService_StopTraceDebug(t *testing.T) {
	type args struct {
		ctx context.Context
		req *pb.StopTraceDebugRequest
	}
	tests := []struct {
		name     string
		service  string
		config   string
		args     args
		wantResp *pb.StopTraceDebugResponse
		wantErr  bool
	}{
		// TODO: Add test cases.
		//		{
		//			"case 1",
		//			"erda.msp.apm.trace.TraceService",
		//			`
		//erda.msp.apm.trace:
		//`,
		//			args{
		//				context.TODO(),
		//				&pb.StopTraceDebugRequest{
		//					// TODO: setup fields
		//				},
		//			},
		//			&pb.StopTraceDebugResponse{
		//				// TODO: setup fields.
		//			},
		//			false,
		//		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hub := servicehub.New()
			events := hub.Events()
			go func() {
				hub.RunWithOptions(&servicehub.RunOptions{Content: tt.config})
			}()
			err := <-events.Started()
			if err != nil {
				t.Error(err)
				return
			}
			srv := hub.Service(tt.service).(pb.TraceServiceServer)
			got, err := srv.StopTraceDebug(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("traceService.StopTraceDebug() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.wantResp) {
				t.Errorf("traceService.StopTraceDebug() = %v, want %v", got, tt.wantResp)
			}
		})
	}
}

func Test_traceService_GetTraceDebugHistoryStatusByRequestID(t *testing.T) {
	type args struct {
		ctx context.Context
		req *pb.GetTraceDebugStatusByRequestIDRequest
	}
	tests := []struct {
		name     string
		service  string
		config   string
		args     args
		wantResp *pb.GetTraceDebugStatusByRequestIDResponse
		wantErr  bool
	}{
		// TODO: Add test cases.
		//		{
		//			"case 1",
		//			"erda.msp.apm.trace.TraceService",
		//			`
		//erda.msp.apm.trace:
		//`,
		//			args{
		//				context.TODO(),
		//				&pb.GetTraceDebugStatusByRequestIDRequest{
		//					// TODO: setup fields
		//				},
		//			},
		//			&pb.GetTraceDebugStatusByRequestIDResponse{
		//				// TODO: setup fields.
		//			},
		//			false,
		//		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hub := servicehub.New()
			events := hub.Events()
			go func() {
				hub.RunWithOptions(&servicehub.RunOptions{Content: tt.config})
			}()
			err := <-events.Started()
			if err != nil {
				t.Error(err)
				return
			}
			srv := hub.Service(tt.service).(pb.TraceServiceServer)
			got, err := srv.GetTraceDebugHistoryStatusByRequestID(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("traceService.GetTraceDebugHistoryStatusByRequestID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.wantResp) {
				t.Errorf("traceService.GetTraceDebugHistoryStatusByRequestID() = %v, want %v", got, tt.wantResp)
			}
		})
	}
}

func Test_composeTraceRequestHistory(t *testing.T) {
	key := uuid.NewV4().String()
	req := pb.CreateTraceDebugRequest{
		Method:    "GET",
		Url:       "http://erda.cloud",
		Body:      "",
		Query:     map[string]string{},
		Header:    map[string]string{},
		ScopeID:   key,
		ProjectID: "1",
	}

	queryString, err := json.Marshal(req.Query)
	if err != nil {
		return
	}
	headerString, err := json.Marshal(req.Header)
	if err != nil {
		return
	}
	bodyValid := json.Valid([]byte(req.Body))
	if req.Body != "" && !bodyValid {
		return
	}
	if req.CreateTime == "" || req.UpdateTime == "" {
		req.CreateTime = time.Now().Format(layout)
		req.UpdateTime = time.Now().Format(layout)
	}
	createTime, err := time.ParseInLocation(layout, req.CreateTime, time.Local)
	if err != nil {
		return
	}
	updateTime, err := time.ParseInLocation(layout, req.UpdateTime, time.Local)
	if err != nil {
		return
	}
	h := &db.TraceRequestHistory{
		TerminusKey:    req.ScopeID,
		Url:            req.Url,
		QueryString:    string(queryString),
		Header:         string(headerString),
		Body:           req.Body,
		Method:         req.Method,
		Status:         int(req.Status),
		ResponseBody:   req.ResponseBody,
		ResponseStatus: int(req.ResponseCode),
		CreateTime:     createTime,
		UpdateTime:     updateTime,
	}
	type args struct {
		req *pb.CreateTraceDebugRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *db.TraceRequestHistory
		wantErr bool
	}{
		{"case-1", args{&req}, h, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := composeTraceRequestHistory(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("composeTraceRequestHistory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == nil {
				t.Errorf("composeTraceRequestHistory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			tt.want.RequestId = got.RequestId
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("composeTraceRequestHistory() got = %v, want %v", got, tt.want)
			}
		})
	}
}
