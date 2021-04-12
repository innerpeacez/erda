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

package cmdb

import (
	"github.com/erda-project/erda/apistructs"
	"github.com/erda-project/erda/modules/openapi/api/apis"
	"github.com/erda-project/erda/modules/openapi/api/spec"
)

var CMDB_ORG_CREATE = apis.ApiSpec{
	Path:         "/api/orgs",
	BackendPath:  "/api/orgs",
	Host:         "cmdb.marathon.l4lb.thisdcos.directory:9093",
	Scheme:       "http",
	Method:       "POST",
	CheckLogin:   true,
	CheckToken:   true,
	IsOpenAPI:    true,
	RequestType:  apistructs.OrgCreateRequest{},
	ResponseType: apistructs.OrgCreateResponse{},
	Doc:          "summary: 创建组织",
	Audit: func(ctx *spec.AuditContext) error {
		var requestBody apistructs.OrgCreateRequest
		if err := ctx.BindRequestData(&requestBody); err != nil {
			return err
		}
		var responseBody apistructs.OrgCreateResponse
		if err := ctx.BindResponseData(&responseBody); err != nil {
			return err
		}
		return ctx.CreateAudit(&apistructs.Audit{
			ScopeType:    apistructs.SysScope,
			ScopeID:      1,
			OrgID:        responseBody.Data.ID,
			TemplateName: apistructs.CreateOrgTemplate,
			Context:      map[string]interface{}{"orgName": requestBody.Name},
		})
	},
}
