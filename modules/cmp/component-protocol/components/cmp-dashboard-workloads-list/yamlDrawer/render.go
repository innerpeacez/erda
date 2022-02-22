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

package yamlDrawer

import (
	"context"
	"encoding/json"

	"github.com/erda-project/erda-infra/base/servicehub"
	"github.com/erda-project/erda-infra/providers/component-protocol/cpregister/base"
	"github.com/erda-project/erda-infra/providers/component-protocol/cptype"
	"github.com/erda-project/erda-infra/providers/component-protocol/utils/cputil"
)

func init() {
	base.InitProviderWithCreator("cmp-dashboard-workloads-list", "yamlDrawer", func() servicehub.Provider {
		return &ComponentYamlDrawer{}
	})
}

func (d *ComponentYamlDrawer) Render(ctx context.Context, component *cptype.Component, _ cptype.Scenario,
	event cptype.ComponentEvent, gs *cptype.GlobalStateData) error {
	d.InitComponent(ctx)
	d.Props.Title = d.sdk.I18n("viewOrEditYaml")
	d.Props.Size = "l"

	workloadID, _ := (*gs)["workloadID"].(string)
	d.State.Visible = workloadID != ""
	d.Transfer(component)
	return nil
}

func (d *ComponentYamlDrawer) InitComponent(ctx context.Context) {
	sdk := cputil.SDK(ctx)
	d.sdk = sdk
}

func (d *ComponentYamlDrawer) GenComponentState(c *cptype.Component) error {
	if c == nil || c.State == nil {
		return nil
	}
	var state State
	jsonData, err := json.Marshal(c.State)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(jsonData, &state); err != nil {
		return err
	}
	d.State = state
	return nil
}

func (d *ComponentYamlDrawer) Transfer(c *cptype.Component) {
	c.Props = cputil.MustConvertProps(d.Props)
	c.State = map[string]interface{}{
		"visible": d.State.Visible,
	}
}
