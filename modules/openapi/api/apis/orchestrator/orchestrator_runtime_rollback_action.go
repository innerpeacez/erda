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

package orchestrator

import "github.com/erda-project/erda/modules/openapi/api/apis"

var ORCHESTRATOR_RUNTIME_ROLLBACK_ACTION = apis.ApiSpec{
	Path:        "/api/runtimes/<runtimeId>/actions/rollback-action",
	BackendPath: "/api/runtimes/<runtimeId>/actions/rollback-action",
	Host:        "orchestrator.marathon.l4lb.thisdcos.directory:8081",
	Scheme:      "http",
	Method:      "POST",
	CheckLogin:  true,
	CheckToken:  true,
	Doc:         `rollback 的 action 调用 api`,
}
