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

package main

import (
	"github.com/erda-project/erda-infra/base/servicehub"
	"github.com/erda-project/erda/pkg/common"

	// providers
	_ "github.com/erda-project/erda-infra/providers/component-protocol"
	_ "github.com/erda-project/erda-infra/providers/grpcclient"
	_ "github.com/erda-project/erda-infra/providers/i18n"
	_ "github.com/erda-project/erda-infra/providers/mysql"
	_ "github.com/erda-project/erda-infra/providers/serviceregister"
	_ "github.com/erda-project/erda-proto-go/cmp/dashboard/client"
	_ "github.com/erda-project/erda-proto-go/core/pipeline/cms/client"
	_ "github.com/erda-project/erda-proto-go/core/pipeline/definition/client"
	_ "github.com/erda-project/erda-proto-go/orchestrator/addon/mysql/client"
	_ "github.com/erda-project/erda/modules/dop"
	_ "github.com/erda-project/erda/modules/dop/providers/autotest/testplan"
	_ "github.com/erda-project/erda/modules/dop/providers/contribution"
	_ "github.com/erda-project/erda/modules/dop/providers/project/home"
	_ "github.com/erda-project/erda/modules/pipeline/providers/definition_client"

	// components
	_ "github.com/erda-project/erda/modules/dop/component-protocol/components"
)

func main() {
	common.Run(&servicehub.RunOptions{
		ConfigFile: "conf/dop/dop.yaml",
	})
}
