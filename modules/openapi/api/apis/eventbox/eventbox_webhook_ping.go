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

package eventbox

import (
	"github.com/erda-project/erda/apistructs"
	"github.com/erda-project/erda/modules/openapi/api/apis"
)

var EVENTBOX_WEBHOOK_PING = apis.ApiSpec{
	Path:         "/api/webhooks/<id>/actions/ping",
	BackendPath:  "/api/dice/eventbox/webhooks/<id>/actions/ping",
	Host:         "eventbox.marathon.l4lb.thisdcos.directory:9528",
	Scheme:       "http",
	Method:       "POST",
	CheckLogin:   true,
	RequestType:  apistructs.WebhookPingRequest{},
	ResponseType: apistructs.WebhookPingResponse{},
	Doc:          `ping webhook, 发送 ping 事件`,
	IsOpenAPI:    true,
}
