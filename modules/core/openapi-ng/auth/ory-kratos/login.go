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

package orykratos

import (
	"net/http"

	"github.com/erda-project/erda/modules/core/openapi-ng/common"
)

func (p *provider) LoginURL(rw http.ResponseWriter, r *http.Request) {
	common.ResponseJSON(rw, &struct {
		URL string `json:"url"`
	}{
		URL: "/uc/auth/login",
	})
}

func (p *provider) Logout(rw http.ResponseWriter, r *http.Request) {
	common.ResponseJSON(rw, &struct {
		URL string `json:"url"`
	}{
		URL: "/.ory/kratos/public/self-service/browser/flows/logout",
	})
}
