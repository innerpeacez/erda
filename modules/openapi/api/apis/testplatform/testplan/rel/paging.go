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

package rel

import (
	"net/http"

	"github.com/erda-project/erda/apistructs"
	"github.com/erda-project/erda/modules/openapi/api/apis"
)

var PAGING = apis.ApiSpec{
	Path:         "/api/testplans/<testPlanID>/testcase-relations",
	BackendPath:  "/api/testplans/<testPlanID>/testcase-relations",
	Host:         "dop.marathon.l4lb.thisdcos.directory:9527",
	Scheme:       "http",
	Method:       http.MethodGet,
	CheckLogin:   true,
	RequestType:  apistructs.TestPlanCaseRelPagingRequest{},
	ResponseType: apistructs.TestPlanCaseRelPagingResponse{},
	Doc:          "summary: 获取测试计划用例关系列表",
}
