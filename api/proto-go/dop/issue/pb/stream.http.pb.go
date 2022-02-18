// Code generated by protoc-gen-go-http. DO NOT EDIT.
// Source: stream.proto

package pb

import (
	context "context"
	http1 "net/http"

	transport "github.com/erda-project/erda-infra/pkg/transport"
	http "github.com/erda-project/erda-infra/pkg/transport/http"
	urlenc "github.com/erda-project/erda-infra/pkg/urlenc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "github.com/erda-project/erda-infra/pkg/transport/http" package it is being compiled against.
const _ = http.SupportPackageIsVersion1

// CommentIssueStreamServiceHandler is the server API for CommentIssueStreamService service.
type CommentIssueStreamServiceHandler interface {
	// POST /api/issues/actions/batch-create-comment-stream
	BatchCreateIssueStream(context.Context, *CommentIssueStreamBatchCreateRequest) (*CommentIssueStreamBatchCreateResponse, error)
}

// RegisterCommentIssueStreamServiceHandler register CommentIssueStreamServiceHandler to http.Router.
func RegisterCommentIssueStreamServiceHandler(r http.Router, srv CommentIssueStreamServiceHandler, opts ...http.HandleOption) {
	h := http.DefaultHandleOptions()
	for _, op := range opts {
		op(h)
	}
	encodeFunc := func(fn func(http1.ResponseWriter, *http1.Request) (interface{}, error)) http.HandlerFunc {
		handler := func(w http1.ResponseWriter, r *http1.Request) {
			out, err := fn(w, r)
			if err != nil {
				h.Error(w, r, err)
				return
			}
			if err := h.Encode(w, r, out); err != nil {
				h.Error(w, r, err)
			}
		}
		if h.HTTPInterceptor != nil {
			handler = h.HTTPInterceptor(handler)
		}
		return handler
	}

	add_BatchCreateIssueStream := func(method, path string, fn func(context.Context, *CommentIssueStreamBatchCreateRequest) (*CommentIssueStreamBatchCreateResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*CommentIssueStreamBatchCreateRequest))
		}
		var BatchCreateIssueStream_info transport.ServiceInfo
		if h.Interceptor != nil {
			BatchCreateIssueStream_info = transport.NewServiceInfo("erda.dop.issue.CommentIssueStreamService", "BatchCreateIssueStream", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, BatchCreateIssueStream_info)
				}
				r = r.WithContext(ctx)
				var in CommentIssueStreamBatchCreateRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_BatchCreateIssueStream("POST", "/api/issues/actions/batch-create-comment-stream", srv.BatchCreateIssueStream)
}
