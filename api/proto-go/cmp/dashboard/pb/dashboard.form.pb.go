// Code generated by protoc-gen-go-form. DO NOT EDIT.
// Source: dashboard.proto

package pb

import (
	url "net/url"
	strconv "strconv"

	urlenc "github.com/erda-project/erda-infra/pkg/urlenc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "github.com/erda-project/erda-infra/pkg/urlenc" package it is being compiled against.
var _ urlenc.URLValuesUnmarshaler = (*GetClustersResourcesRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetClusterResourcesResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*ClusterResourceDetail)(nil)
var _ urlenc.URLValuesUnmarshaler = (*HostResourceDetail)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetNamespacesResourcesRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*ClusterNamespacePair)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetNamespacesResourcesResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*ClusterResourceItem)(nil)
var _ urlenc.URLValuesUnmarshaler = (*NamespaceResourceDetail)(nil)

// GetClustersResourcesRequest implement urlenc.URLValuesUnmarshaler.
func (m *GetClustersResourcesRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "clusterNames":
				m.ClusterNames = vals
			}
		}
	}
	return nil
}

// GetClusterResourcesResponse implement urlenc.URLValuesUnmarshaler.
func (m *GetClusterResourcesResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "total":
				val, err := strconv.ParseUint(vals[0], 10, 32)
				if err != nil {
					return err
				}
				m.Total = uint32(val)
			}
		}
	}
	return nil
}

// ClusterResourceDetail implement urlenc.URLValuesUnmarshaler.
func (m *ClusterResourceDetail) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "success":
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.Success = val
			case "err":
				m.Err = vals[0]
			case "clusterName":
				m.ClusterName = vals[0]
			}
		}
	}
	return nil
}

// HostResourceDetail implement urlenc.URLValuesUnmarshaler.
func (m *HostResourceDetail) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "host":
				m.Host = vals[0]
			case "cpuAllocatable":
				val, err := strconv.ParseUint(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.CpuAllocatable = val
			case "cpuTotal":
				val, err := strconv.ParseUint(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.CpuTotal = val
			case "cpuRequest":
				val, err := strconv.ParseUint(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.CpuRequest = val
			case "memAllocatable":
				val, err := strconv.ParseUint(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.MemAllocatable = val
			case "memTotal":
				val, err := strconv.ParseUint(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.MemTotal = val
			case "memRequest":
				val, err := strconv.ParseUint(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.MemRequest = val
			case "labels":
				m.Labels = vals
			}
		}
	}
	return nil
}

// GetNamespacesResourcesRequest implement urlenc.URLValuesUnmarshaler.
func (m *GetNamespacesResourcesRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}

// ClusterNamespacePair implement urlenc.URLValuesUnmarshaler.
func (m *ClusterNamespacePair) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "clusterName":
				m.ClusterName = vals[0]
			case "namespace":
				m.Namespace = vals[0]
			}
		}
	}
	return nil
}

// GetNamespacesResourcesResponse implement urlenc.URLValuesUnmarshaler.
func (m *GetNamespacesResourcesResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "total":
				val, err := strconv.ParseUint(vals[0], 10, 32)
				if err != nil {
					return err
				}
				m.Total = uint32(val)
			}
		}
	}
	return nil
}

// ClusterResourceItem implement urlenc.URLValuesUnmarshaler.
func (m *ClusterResourceItem) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "success":
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.Success = val
			case "err":
				m.Err = vals[0]
			case "clusterName":
				m.ClusterName = vals[0]
			}
		}
	}
	return nil
}

// NamespaceResourceDetail implement urlenc.URLValuesUnmarshaler.
func (m *NamespaceResourceDetail) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "namespace":
				m.Namespace = vals[0]
			case "cpuRequest":
				val, err := strconv.ParseUint(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.CpuRequest = val
			case "memRequest":
				val, err := strconv.ParseUint(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.MemRequest = val
			}
		}
	}
	return nil
}
