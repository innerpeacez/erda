// Copyright (c) 2021 Terminus, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package dicehub

import (
	"github.com/erda-project/erda/apistructs"
	"github.com/erda-project/erda/modules/openapi/api/apis"
)

var EXTENSION_VERSION_QUERY = apis.ApiSpec{
	Path:         "/api/extensions/<name>",
	BackendPath:  "/api/extensions/<name>",
	Host:         "dicehub.marathon.l4lb.thisdcos.directory:10000",
	Scheme:       "http",
	Method:       "GET",
	CheckLogin:   false,
	CheckToken:   false,
	RequestType:  apistructs.ExtensionVersionQueryRequest{},
	ResponseType: apistructs.ExtensionVersionQueryResponse{},
	Doc:          "summary: 查询扩展版本",
}
