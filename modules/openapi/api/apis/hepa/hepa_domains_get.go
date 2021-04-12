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

package hepa

import (
	"github.com/erda-project/erda/modules/openapi/api/apis"
)

var HEPA_DOMAINS_GET = apis.ApiSpec{
	Path:        "/api/domains",
	BackendPath: "/api/gateway/domains",
	Host:        "hepa.marathon.l4lb.thisdcos.directory:8080",
	K8SHost:     "hepa.default.svc.cluster.local:8080",
	Scheme:      "http",
	Method:      "GET",
	CheckLogin:  true,
	Doc:         `获取域名列表`,
}
