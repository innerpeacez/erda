// Code generated by protoc-gen-go-http. DO NOT EDIT.
// Source: channel.proto

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

// NotifyChannelServiceHandler is the server API for NotifyChannelService service.
type NotifyChannelServiceHandler interface {
	// POST /api/notify-channel
	CreateNotifyChannel(context.Context, *CreateNotifyChannelRequest) (*CreateNotifyChannelResponse, error)
	// GET /api/notify-channels
	GetNotifyChannels(context.Context, *GetNotifyChannelsRequest) (*GetNotifyChannelsResponse, error)
	// PUT /api/notify-channel
	UpdateNotifyChannel(context.Context, *UpdateNotifyChannelRequest) (*UpdateNotifyChannelResponse, error)
	// GET /api/notify-channel
	GetNotifyChannel(context.Context, *GetNotifyChannelRequest) (*GetNotifyChannelResponse, error)
	// DELETE /api/notify-channel
	DeleteNotifyChannel(context.Context, *DeleteNotifyChannelRequest) (*DeleteNotifyChannelResponse, error)
}

// RegisterNotifyChannelServiceHandler register NotifyChannelServiceHandler to http.Router.
func RegisterNotifyChannelServiceHandler(r http.Router, srv NotifyChannelServiceHandler, opts ...http.HandleOption) {
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

	add_CreateNotifyChannel := func(method, path string, fn func(context.Context, *CreateNotifyChannelRequest) (*CreateNotifyChannelResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*CreateNotifyChannelRequest))
		}
		var CreateNotifyChannel_info transport.ServiceInfo
		if h.Interceptor != nil {
			CreateNotifyChannel_info = transport.NewServiceInfo("erda.core.services.notify.channel.NotifyChannelService", "CreateNotifyChannel", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, CreateNotifyChannel_info)
				}
				r = r.WithContext(ctx)
				var in CreateNotifyChannelRequest
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

	add_GetNotifyChannels := func(method, path string, fn func(context.Context, *GetNotifyChannelsRequest) (*GetNotifyChannelsResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*GetNotifyChannelsRequest))
		}
		var GetNotifyChannels_info transport.ServiceInfo
		if h.Interceptor != nil {
			GetNotifyChannels_info = transport.NewServiceInfo("erda.core.services.notify.channel.NotifyChannelService", "GetNotifyChannels", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, GetNotifyChannels_info)
				}
				r = r.WithContext(ctx)
				var in GetNotifyChannelsRequest
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

	add_UpdateNotifyChannel := func(method, path string, fn func(context.Context, *UpdateNotifyChannelRequest) (*UpdateNotifyChannelResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*UpdateNotifyChannelRequest))
		}
		var UpdateNotifyChannel_info transport.ServiceInfo
		if h.Interceptor != nil {
			UpdateNotifyChannel_info = transport.NewServiceInfo("erda.core.services.notify.channel.NotifyChannelService", "UpdateNotifyChannel", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, UpdateNotifyChannel_info)
				}
				r = r.WithContext(ctx)
				var in UpdateNotifyChannelRequest
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

	add_GetNotifyChannel := func(method, path string, fn func(context.Context, *GetNotifyChannelRequest) (*GetNotifyChannelResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*GetNotifyChannelRequest))
		}
		var GetNotifyChannel_info transport.ServiceInfo
		if h.Interceptor != nil {
			GetNotifyChannel_info = transport.NewServiceInfo("erda.core.services.notify.channel.NotifyChannelService", "GetNotifyChannel", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, GetNotifyChannel_info)
				}
				r = r.WithContext(ctx)
				var in GetNotifyChannelRequest
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

	add_DeleteNotifyChannel := func(method, path string, fn func(context.Context, *DeleteNotifyChannelRequest) (*DeleteNotifyChannelResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*DeleteNotifyChannelRequest))
		}
		var DeleteNotifyChannel_info transport.ServiceInfo
		if h.Interceptor != nil {
			DeleteNotifyChannel_info = transport.NewServiceInfo("erda.core.services.notify.channel.NotifyChannelService", "DeleteNotifyChannel", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, DeleteNotifyChannel_info)
				}
				r = r.WithContext(ctx)
				var in DeleteNotifyChannelRequest
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

	add_CreateNotifyChannel("POST", "/api/notify-channel", srv.CreateNotifyChannel)
	add_GetNotifyChannels("GET", "/api/notify-channels", srv.GetNotifyChannels)
	add_UpdateNotifyChannel("PUT", "/api/notify-channel", srv.UpdateNotifyChannel)
	add_GetNotifyChannel("GET", "/api/notify-channel", srv.GetNotifyChannel)
	add_DeleteNotifyChannel("DELETE", "/api/notify-channel", srv.DeleteNotifyChannel)
}
