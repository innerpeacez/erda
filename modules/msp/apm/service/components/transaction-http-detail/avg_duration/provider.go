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

package avg_duration

import (
	"reflect"

	"github.com/erda-project/erda-infra/base/logs"
	"github.com/erda-project/erda-infra/base/servicehub"
	"github.com/erda-project/erda-infra/providers/component-protocol/components/linegraph"
	"github.com/erda-project/erda-infra/providers/component-protocol/components/linegraph/impl"
	"github.com/erda-project/erda-infra/providers/component-protocol/cptype"
	"github.com/erda-project/erda-infra/providers/component-protocol/protocol"
	"github.com/erda-project/erda-infra/providers/i18n"
	metricpb "github.com/erda-project/erda-proto-go/core/monitor/metric/pb"
)

type provider struct {
	impl.DefaultLineGraph
	Log    logs.Logger
	I18n   i18n.Translator              `autowired:"i18n" translator:"msp-i18n"`
	Metric metricpb.MetricServiceServer `autowired:"erda.core.monitor.metric.MetricService"`
}

// RegisterInitializeOp .
func (p *provider) RegisterInitializeOp() (opFunc cptype.OperationFunc) {
	return func(sdk *cptype.SDK) {
		data := linegraph.Data{}
		p.StdDataPtr = &data
	}
}

// RegisterRenderingOp .
func (p *provider) RegisterRenderingOp() (opFunc cptype.OperationFunc) {
	return p.RegisterInitializeOp()
}

// Init .
func (p *provider) Init(ctx servicehub.Context) error {
	p.DefaultLineGraph = impl.DefaultLineGraph{}
	v := reflect.ValueOf(p)
	v.Elem().FieldByName("Impl").Set(v)
	compName := "avgDuration"
	if ctx.Label() != "" {
		compName = ctx.Label()
	}
	protocol.MustRegisterComponent(&protocol.CompRenderSpec{
		Scenario: "transaction-http-detail",
		CompName: compName,
		Creator:  func() cptype.IComponent { return p },
	})
	return nil
}

// Provide .
func (p *provider) Provide(ctx servicehub.DependencyContext, args ...interface{}) interface{} {
	return p
}

func init() {
	servicehub.Register("component-protocol.components.transaction-http-detail.avgDuration", &servicehub.Spec{
		Creator: func() servicehub.Provider { return &provider{} },
	})
}
