// Code generated by protoc-gen-go-http. DO NOT EDIT.
// Source: release.proto

package pb

import (
	context "context"
	http1 "net/http"
	strconv "strconv"
	strings "strings"

	transport "github.com/erda-project/erda-infra/pkg/transport"
	http "github.com/erda-project/erda-infra/pkg/transport/http"
	httprule "github.com/erda-project/erda-infra/pkg/transport/http/httprule"
	runtime "github.com/erda-project/erda-infra/pkg/transport/http/runtime"
	urlenc "github.com/erda-project/erda-infra/pkg/urlenc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "github.com/erda-project/erda-infra/pkg/transport/http" package it is being compiled against.
const _ = http.SupportPackageIsVersion1

// ReleaseServiceHandler is the server API for ReleaseService service.
type ReleaseServiceHandler interface {
	// POST /api/releases
	CreateRelease(context.Context, *ReleaseCreateRequest) (*ReleaseCreateResponseData, error)
	// PUT /api/releases/{releaseID}
	UpdateRelease(context.Context, *ReleaseUpdateRequest) (*ReleaseUpdateResponse, error)
	// PUT /api/releases/{releaseID}/reference/actions/change
	UpdateReleaseReference(context.Context, *ReleaseReferenceUpdateRequest) (*ReleaseDataResponse, error)
	// GET /api/releases/{releaseID}/actions/get-plist
	GetIosPlist(context.Context, *GetIosPlistRequest) (*GetIosPlistResponse, error)
	// GET /api/releases/{releaseID}
	GetRelease(context.Context, *ReleaseGetRequest) (*ReleaseGetResponse, error)
	// DELETE /api/releases/{releaseID}
	DeleteRelease(context.Context, *ReleaseDeleteRequest) (*ReleaseDeleteResponse, error)
	// GET /api/releases
	ListRelease(context.Context, *ReleaseListRequest) (*ReleaseListResponse, error)
	// GET /api/releases/actions/get-name
	ListReleaseName(context.Context, *ListReleaseNameRequest) (*ListReleaseNameResponse, error)
	// GET /api/releases/actions/get-latest
	GetLatestReleases(context.Context, *GetLatestReleasesRequest) (*GetLatestReleasesResponse, error)
	// POST /gc
	ReleaseGC(context.Context, *ReleaseGCRequest) (*ReleaseDataResponse, error)
	// POST /api/releases/actions/upload
	UploadRelease(context.Context, *ReleaseUploadRequest) (*ReleaseUploadResponse, error)
	// GET /api/releases/actions/parse-version
	ParseReleaseFile(context.Context, *ParseReleaseFileRequest) (*ParseReleaseFileResponse, error)
	// PUT /api/releases/{releaseId}/actions/formal
	ToFormalRelease(context.Context, *FormalReleaseRequest) (*FormalReleaseResponse, error)
	// PUT /api/releases
	ToFormalReleases(context.Context, *FormalReleasesRequest) (*FormalReleasesResponse, error)
	// DELETE /api/releases
	DeleteReleases(context.Context, *ReleasesDeleteRequest) (*ReleasesDeleteResponse, error)
	// GET /api/releases/actions/check-version
	CheckVersion(context.Context, *CheckVersionRequest) (*CheckVersionResponse, error)
}

// RegisterReleaseServiceHandler register ReleaseServiceHandler to http.Router.
func RegisterReleaseServiceHandler(r http.Router, srv ReleaseServiceHandler, opts ...http.HandleOption) {
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

	add_CreateRelease := func(method, path string, fn func(context.Context, *ReleaseCreateRequest) (*ReleaseCreateResponseData, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*ReleaseCreateRequest))
		}
		var CreateRelease_info transport.ServiceInfo
		if h.Interceptor != nil {
			CreateRelease_info = transport.NewServiceInfo("erda.core.dicehub.release.ReleaseService", "CreateRelease", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, CreateRelease_info)
				}
				r = r.WithContext(ctx)
				var in ReleaseCreateRequest
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

	add_UpdateRelease := func(method, path string, fn func(context.Context, *ReleaseUpdateRequest) (*ReleaseUpdateResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*ReleaseUpdateRequest))
		}
		var UpdateRelease_info transport.ServiceInfo
		if h.Interceptor != nil {
			UpdateRelease_info = transport.NewServiceInfo("erda.core.dicehub.release.ReleaseService", "UpdateRelease", srv)
			handler = h.Interceptor(handler)
		}
		compiler, _ := httprule.Parse(path)
		temp := compiler.Compile()
		pattern, _ := runtime.NewPattern(httprule.SupportPackageIsVersion1, temp.OpCodes, temp.Pool, temp.Verb)
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, UpdateRelease_info)
				}
				r = r.WithContext(ctx)
				var in ReleaseUpdateRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				path := r.URL.Path
				if len(path) > 0 {
					components := strings.Split(path[1:], "/")
					last := len(components) - 1
					var verb string
					if idx := strings.LastIndex(components[last], ":"); idx >= 0 {
						c := components[last]
						components[last], verb = c[:idx], c[idx+1:]
					}
					vars, err := pattern.Match(components, verb)
					if err != nil {
						return nil, err
					}
					for k, val := range vars {
						switch k {
						case "releaseID":
							in.ReleaseID = val
						}
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

	add_UpdateReleaseReference := func(method, path string, fn func(context.Context, *ReleaseReferenceUpdateRequest) (*ReleaseDataResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*ReleaseReferenceUpdateRequest))
		}
		var UpdateReleaseReference_info transport.ServiceInfo
		if h.Interceptor != nil {
			UpdateReleaseReference_info = transport.NewServiceInfo("erda.core.dicehub.release.ReleaseService", "UpdateReleaseReference", srv)
			handler = h.Interceptor(handler)
		}
		compiler, _ := httprule.Parse(path)
		temp := compiler.Compile()
		pattern, _ := runtime.NewPattern(httprule.SupportPackageIsVersion1, temp.OpCodes, temp.Pool, temp.Verb)
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, UpdateReleaseReference_info)
				}
				r = r.WithContext(ctx)
				var in ReleaseReferenceUpdateRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				path := r.URL.Path
				if len(path) > 0 {
					components := strings.Split(path[1:], "/")
					last := len(components) - 1
					var verb string
					if idx := strings.LastIndex(components[last], ":"); idx >= 0 {
						c := components[last]
						components[last], verb = c[:idx], c[idx+1:]
					}
					vars, err := pattern.Match(components, verb)
					if err != nil {
						return nil, err
					}
					for k, val := range vars {
						switch k {
						case "releaseID":
							in.ReleaseID = val
						}
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

	add_GetIosPlist := func(method, path string, fn func(context.Context, *GetIosPlistRequest) (*GetIosPlistResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*GetIosPlistRequest))
		}
		var GetIosPlist_info transport.ServiceInfo
		if h.Interceptor != nil {
			GetIosPlist_info = transport.NewServiceInfo("erda.core.dicehub.release.ReleaseService", "GetIosPlist", srv)
			handler = h.Interceptor(handler)
		}
		compiler, _ := httprule.Parse(path)
		temp := compiler.Compile()
		pattern, _ := runtime.NewPattern(httprule.SupportPackageIsVersion1, temp.OpCodes, temp.Pool, temp.Verb)
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, GetIosPlist_info)
				}
				r = r.WithContext(ctx)
				var in GetIosPlistRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				path := r.URL.Path
				if len(path) > 0 {
					components := strings.Split(path[1:], "/")
					last := len(components) - 1
					var verb string
					if idx := strings.LastIndex(components[last], ":"); idx >= 0 {
						c := components[last]
						components[last], verb = c[:idx], c[idx+1:]
					}
					vars, err := pattern.Match(components, verb)
					if err != nil {
						return nil, err
					}
					for k, val := range vars {
						switch k {
						case "releaseID":
							in.ReleaseID = val
						}
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

	add_GetRelease := func(method, path string, fn func(context.Context, *ReleaseGetRequest) (*ReleaseGetResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*ReleaseGetRequest))
		}
		var GetRelease_info transport.ServiceInfo
		if h.Interceptor != nil {
			GetRelease_info = transport.NewServiceInfo("erda.core.dicehub.release.ReleaseService", "GetRelease", srv)
			handler = h.Interceptor(handler)
		}
		compiler, _ := httprule.Parse(path)
		temp := compiler.Compile()
		pattern, _ := runtime.NewPattern(httprule.SupportPackageIsVersion1, temp.OpCodes, temp.Pool, temp.Verb)
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, GetRelease_info)
				}
				r = r.WithContext(ctx)
				var in ReleaseGetRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				path := r.URL.Path
				if len(path) > 0 {
					components := strings.Split(path[1:], "/")
					last := len(components) - 1
					var verb string
					if idx := strings.LastIndex(components[last], ":"); idx >= 0 {
						c := components[last]
						components[last], verb = c[:idx], c[idx+1:]
					}
					vars, err := pattern.Match(components, verb)
					if err != nil {
						return nil, err
					}
					for k, val := range vars {
						switch k {
						case "releaseID":
							in.ReleaseID = val
						}
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

	add_DeleteRelease := func(method, path string, fn func(context.Context, *ReleaseDeleteRequest) (*ReleaseDeleteResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*ReleaseDeleteRequest))
		}
		var DeleteRelease_info transport.ServiceInfo
		if h.Interceptor != nil {
			DeleteRelease_info = transport.NewServiceInfo("erda.core.dicehub.release.ReleaseService", "DeleteRelease", srv)
			handler = h.Interceptor(handler)
		}
		compiler, _ := httprule.Parse(path)
		temp := compiler.Compile()
		pattern, _ := runtime.NewPattern(httprule.SupportPackageIsVersion1, temp.OpCodes, temp.Pool, temp.Verb)
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, DeleteRelease_info)
				}
				r = r.WithContext(ctx)
				var in ReleaseDeleteRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				path := r.URL.Path
				if len(path) > 0 {
					components := strings.Split(path[1:], "/")
					last := len(components) - 1
					var verb string
					if idx := strings.LastIndex(components[last], ":"); idx >= 0 {
						c := components[last]
						components[last], verb = c[:idx], c[idx+1:]
					}
					vars, err := pattern.Match(components, verb)
					if err != nil {
						return nil, err
					}
					for k, val := range vars {
						switch k {
						case "releaseID":
							in.ReleaseID = val
						}
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

	add_ListRelease := func(method, path string, fn func(context.Context, *ReleaseListRequest) (*ReleaseListResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*ReleaseListRequest))
		}
		var ListRelease_info transport.ServiceInfo
		if h.Interceptor != nil {
			ListRelease_info = transport.NewServiceInfo("erda.core.dicehub.release.ReleaseService", "ListRelease", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, ListRelease_info)
				}
				r = r.WithContext(ctx)
				var in ReleaseListRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				params := r.URL.Query()
				if vals := params["applicationId"]; len(vals) > 0 {
					in.ApplicationID = vals
				}
				if vals := params["branch"]; len(vals) > 0 {
					in.Branch = vals[0]
				}
				if vals := params["cluster"]; len(vals) > 0 {
					in.Cluster = vals[0]
				}
				if vals := params["commitId"]; len(vals) > 0 {
					in.CommitID = vals[0]
				}
				if vals := params["crossCluster"]; len(vals) > 0 {
					in.CrossCluster = vals[0]
				}
				if vals := params["crossClusterOrSpecifyCluster"]; len(vals) > 0 {
					in.CrossClusterOrSpecifyCluster = vals[0]
				}
				if vals := params["endTime"]; len(vals) > 0 {
					val, err := strconv.ParseInt(vals[0], 10, 64)
					if err != nil {
						return nil, err
					}
					in.EndTime = val
				}
				if vals := params["isFormal"]; len(vals) > 0 {
					in.IsFormal = vals[0]
				}
				if vals := params["isProjectRelease"]; len(vals) > 0 {
					in.IsProjectRelease = vals[0]
				}
				if vals := params["isStable"]; len(vals) > 0 {
					in.IsStable = vals[0]
				}
				if vals := params["isVersion"]; len(vals) > 0 {
					val, err := strconv.ParseBool(vals[0])
					if err != nil {
						return nil, err
					}
					in.IsVersion = val
				}
				if vals := params["latest"]; len(vals) > 0 {
					val, err := strconv.ParseBool(vals[0])
					if err != nil {
						return nil, err
					}
					in.IsLatest = val
				}
				if vals := params["order"]; len(vals) > 0 {
					in.Order = vals[0]
				}
				if vals := params["orderBy"]; len(vals) > 0 {
					in.OrderBy = vals[0]
				}
				if vals := params["pageNum"]; len(vals) > 0 {
					val, err := strconv.ParseInt(vals[0], 10, 64)
					if err != nil {
						return nil, err
					}
					in.PageNum = val
				}
				if vals := params["pageSize"]; len(vals) > 0 {
					val, err := strconv.ParseInt(vals[0], 10, 64)
					if err != nil {
						return nil, err
					}
					in.PageSize = val
				}
				if vals := params["projectId"]; len(vals) > 0 {
					val, err := strconv.ParseInt(vals[0], 10, 64)
					if err != nil {
						return nil, err
					}
					in.ProjectID = val
				}
				if vals := params["q"]; len(vals) > 0 {
					in.Query = vals[0]
				}
				if vals := params["releaseId"]; len(vals) > 0 {
					in.ReleaseID = vals[0]
				}
				if vals := params["releaseName"]; len(vals) > 0 {
					in.ReleaseName = vals[0]
				}
				if vals := params["tags"]; len(vals) > 0 {
					in.Tags = vals[0]
				}
				if vals := params["userId"]; len(vals) > 0 {
					in.UserID = vals
				}
				if vals := params["version"]; len(vals) > 0 {
					in.Version = vals[0]
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_ListReleaseName := func(method, path string, fn func(context.Context, *ListReleaseNameRequest) (*ListReleaseNameResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*ListReleaseNameRequest))
		}
		var ListReleaseName_info transport.ServiceInfo
		if h.Interceptor != nil {
			ListReleaseName_info = transport.NewServiceInfo("erda.core.dicehub.release.ReleaseService", "ListReleaseName", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, ListReleaseName_info)
				}
				r = r.WithContext(ctx)
				var in ListReleaseNameRequest
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

	add_GetLatestReleases := func(method, path string, fn func(context.Context, *GetLatestReleasesRequest) (*GetLatestReleasesResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*GetLatestReleasesRequest))
		}
		var GetLatestReleases_info transport.ServiceInfo
		if h.Interceptor != nil {
			GetLatestReleases_info = transport.NewServiceInfo("erda.core.dicehub.release.ReleaseService", "GetLatestReleases", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, GetLatestReleases_info)
				}
				r = r.WithContext(ctx)
				var in GetLatestReleasesRequest
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

	add_ReleaseGC := func(method, path string, fn func(context.Context, *ReleaseGCRequest) (*ReleaseDataResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*ReleaseGCRequest))
		}
		var ReleaseGC_info transport.ServiceInfo
		if h.Interceptor != nil {
			ReleaseGC_info = transport.NewServiceInfo("erda.core.dicehub.release.ReleaseService", "ReleaseGC", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, ReleaseGC_info)
				}
				r = r.WithContext(ctx)
				var in ReleaseGCRequest
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

	add_UploadRelease := func(method, path string, fn func(context.Context, *ReleaseUploadRequest) (*ReleaseUploadResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*ReleaseUploadRequest))
		}
		var UploadRelease_info transport.ServiceInfo
		if h.Interceptor != nil {
			UploadRelease_info = transport.NewServiceInfo("erda.core.dicehub.release.ReleaseService", "UploadRelease", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, UploadRelease_info)
				}
				r = r.WithContext(ctx)
				var in ReleaseUploadRequest
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

	add_ParseReleaseFile := func(method, path string, fn func(context.Context, *ParseReleaseFileRequest) (*ParseReleaseFileResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*ParseReleaseFileRequest))
		}
		var ParseReleaseFile_info transport.ServiceInfo
		if h.Interceptor != nil {
			ParseReleaseFile_info = transport.NewServiceInfo("erda.core.dicehub.release.ReleaseService", "ParseReleaseFile", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, ParseReleaseFile_info)
				}
				r = r.WithContext(ctx)
				var in ParseReleaseFileRequest
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

	add_ToFormalRelease := func(method, path string, fn func(context.Context, *FormalReleaseRequest) (*FormalReleaseResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*FormalReleaseRequest))
		}
		var ToFormalRelease_info transport.ServiceInfo
		if h.Interceptor != nil {
			ToFormalRelease_info = transport.NewServiceInfo("erda.core.dicehub.release.ReleaseService", "ToFormalRelease", srv)
			handler = h.Interceptor(handler)
		}
		compiler, _ := httprule.Parse(path)
		temp := compiler.Compile()
		pattern, _ := runtime.NewPattern(httprule.SupportPackageIsVersion1, temp.OpCodes, temp.Pool, temp.Verb)
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, ToFormalRelease_info)
				}
				r = r.WithContext(ctx)
				var in FormalReleaseRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				path := r.URL.Path
				if len(path) > 0 {
					components := strings.Split(path[1:], "/")
					last := len(components) - 1
					var verb string
					if idx := strings.LastIndex(components[last], ":"); idx >= 0 {
						c := components[last]
						components[last], verb = c[:idx], c[idx+1:]
					}
					vars, err := pattern.Match(components, verb)
					if err != nil {
						return nil, err
					}
					for k, val := range vars {
						switch k {
						case "releaseId":
							in.ReleaseId = val
						}
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

	add_ToFormalReleases := func(method, path string, fn func(context.Context, *FormalReleasesRequest) (*FormalReleasesResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*FormalReleasesRequest))
		}
		var ToFormalReleases_info transport.ServiceInfo
		if h.Interceptor != nil {
			ToFormalReleases_info = transport.NewServiceInfo("erda.core.dicehub.release.ReleaseService", "ToFormalReleases", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, ToFormalReleases_info)
				}
				r = r.WithContext(ctx)
				var in FormalReleasesRequest
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

	add_DeleteReleases := func(method, path string, fn func(context.Context, *ReleasesDeleteRequest) (*ReleasesDeleteResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*ReleasesDeleteRequest))
		}
		var DeleteReleases_info transport.ServiceInfo
		if h.Interceptor != nil {
			DeleteReleases_info = transport.NewServiceInfo("erda.core.dicehub.release.ReleaseService", "DeleteReleases", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, DeleteReleases_info)
				}
				r = r.WithContext(ctx)
				var in ReleasesDeleteRequest
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

	add_CheckVersion := func(method, path string, fn func(context.Context, *CheckVersionRequest) (*CheckVersionResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*CheckVersionRequest))
		}
		var CheckVersion_info transport.ServiceInfo
		if h.Interceptor != nil {
			CheckVersion_info = transport.NewServiceInfo("erda.core.dicehub.release.ReleaseService", "CheckVersion", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, CheckVersion_info)
				}
				r = r.WithContext(ctx)
				var in CheckVersionRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				params := r.URL.Query()
				if vals := params["appID"]; len(vals) > 0 {
					val, err := strconv.ParseInt(vals[0], 10, 64)
					if err != nil {
						return nil, err
					}
					in.AppID = val
				}
				if vals := params["isProjectRelease"]; len(vals) > 0 {
					val, err := strconv.ParseBool(vals[0])
					if err != nil {
						return nil, err
					}
					in.IsProjectRelease = val
				}
				if vals := params["orgID"]; len(vals) > 0 {
					val, err := strconv.ParseInt(vals[0], 10, 64)
					if err != nil {
						return nil, err
					}
					in.OrgID = val
				}
				if vals := params["projectID"]; len(vals) > 0 {
					val, err := strconv.ParseInt(vals[0], 10, 64)
					if err != nil {
						return nil, err
					}
					in.ProjectID = val
				}
				if vals := params["version"]; len(vals) > 0 {
					in.Version = vals[0]
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_CreateRelease("POST", "/api/releases", srv.CreateRelease)
	add_UpdateRelease("PUT", "/api/releases/{releaseID}", srv.UpdateRelease)
	add_UpdateReleaseReference("PUT", "/api/releases/{releaseID}/reference/actions/change", srv.UpdateReleaseReference)
	add_GetIosPlist("GET", "/api/releases/{releaseID}/actions/get-plist", srv.GetIosPlist)
	add_GetRelease("GET", "/api/releases/{releaseID}", srv.GetRelease)
	add_DeleteRelease("DELETE", "/api/releases/{releaseID}", srv.DeleteRelease)
	add_ListRelease("GET", "/api/releases", srv.ListRelease)
	add_ListReleaseName("GET", "/api/releases/actions/get-name", srv.ListReleaseName)
	add_GetLatestReleases("GET", "/api/releases/actions/get-latest", srv.GetLatestReleases)
	add_ReleaseGC("POST", "/gc", srv.ReleaseGC)
	add_UploadRelease("POST", "/api/releases/actions/upload", srv.UploadRelease)
	add_ParseReleaseFile("GET", "/api/releases/actions/parse-version", srv.ParseReleaseFile)
	add_ToFormalRelease("PUT", "/api/releases/{releaseId}/actions/formal", srv.ToFormalRelease)
	add_ToFormalReleases("PUT", "/api/releases", srv.ToFormalReleases)
	add_DeleteReleases("DELETE", "/api/releases", srv.DeleteReleases)
	add_CheckVersion("GET", "/api/releases/actions/check-version", srv.CheckVersion)
}
