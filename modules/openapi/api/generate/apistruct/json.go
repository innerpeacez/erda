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

package apistruct

import (
	"encoding/json"
	"reflect"
	"strings"

	"github.com/erda-project/erda/pkg/strutil"
)

type JSON map[string]interface{}
type SliceJSON []interface{}

/*
path: URL path
method: http method
summary: 综合性描述
m: 生成的json结构
req: 请求类型结构体
resp: 应答类型结构体
*/
func ToJson(path, method, summary string, group string, m JSON, req, resp interface{}) {
	paths := m["paths"].(JSON)
	definitions := m["definitions"].(JSON)
	reqparams, _ := structToParam(context{request: true}, req)
	_, respparam := structToParam(context{}, resp)

	// convert param to JSON
	reqByte, err := json.Marshal(reqparams)
	if err != nil {
		panic(err)
	}
	respByte, err := json.Marshal(respparam)
	if err != nil {
		panic(err)
	}
	var (
		reqJSON  = make(SliceJSON, 0)
		respJSON = make(JSON)
	)
	if err := json.Unmarshal(reqByte, &reqJSON); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(respByte, &respJSON); err != nil {
		panic(err)
	}
	if _, ok := paths[path]; !ok {
		paths[path] = make(JSON)
	}
	if _, ok := paths[path].(JSON)[method]; !ok {
		paths[path].(JSON)[method] = make(JSON)
	}
	paths[path].(JSON)[method].(JSON)["parameters"] = reqJSON
	paths[path].(JSON)[method].(JSON)["produces"] = []string{"application/json"}
	paths[path].(JSON)[method].(JSON)["responses"] = make(JSON)
	paths[path].(JSON)[method].(JSON)["summary"] = strings.TrimLeft(summary, "summary:")
	if group != "" {
		paths[path].(JSON)[method].(JSON)["tags"] = []string{group}
	}
	paths[path].(JSON)[method].(JSON)["responses"].(JSON)["200"] = make(JSON)
	paths[path].(JSON)[method].(JSON)["responses"].(JSON)["200"].(JSON)["description"] = "OK"
	respTp := reflect.TypeOf(resp)
	paths[path].(JSON)[method].(JSON)["responses"].(JSON)["200"].(JSON)["schema"] = JSON{
		"$ref": strutil.Concat("#/definitions/", respTp.Name()),
	}

	// definitions
	definitions[respTp.Name()] = respJSON
}

func EventToJson(event interface{}, summary string, m JSON) {
	t := reflect.TypeOf(event)
	_, eventSwagger := structToParam(context{}, event)
	eventByte, err := json.Marshal(eventSwagger)
	if err != nil {
		panic(err)
	}
	var eventJSON = make(JSON)
	if err := json.Unmarshal(eventByte, &eventJSON); err != nil {
		panic(err)
	}
	m[t.Name()] = eventJSON
	m[t.Name()].(JSON)["description"] = summary
}
