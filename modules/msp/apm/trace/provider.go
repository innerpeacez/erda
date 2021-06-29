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

package trace

import (
	"github.com/erda-project/erda-infra/base/logs"
	"github.com/erda-project/erda-infra/base/servicehub"
	"github.com/erda-project/erda-infra/pkg/transport"
	metricpb "github.com/erda-project/erda-proto-go/core/monitor/metric/pb"
	"github.com/erda-project/erda-proto-go/msp/apm/trace/pb"
	"github.com/erda-project/erda/pkg/common/apis"
)

type config struct {
}

// +provider
type provider struct {
	Cfg          *config
	Log          logs.Logger
	Register     transport.Register `autowired:"service-register" optional:"true"`
	traceService *traceService
	//Metricq      metricq.Queryer              `autowired:"metrics-query"`
	Metric metricpb.MetricServiceServer `autowired:"erda.core.monitor.metric.MetricService" optional:"true"`
	//Spanq  query.SpanQueryAPI           `autowired:"trace-query"`
}

func (p *provider) Init(ctx servicehub.Context) error {
	p.traceService = &traceService{
		p: p,
		//metricq: p.Metricq,
		metricq: p.Metric,
		//spanq:   p.Spanq,
	}
	if p.Register != nil {
		pb.RegisterTraceServiceImp(p.Register, p.traceService, apis.Options())
	}
	return nil
}

func (p *provider) Provide(ctx servicehub.DependencyContext, args ...interface{}) interface{} {
	switch {
	case ctx.Service() == "erda.msp.apm.trace.TraceService" || ctx.Type() == pb.TraceServiceServerType() || ctx.Type() == pb.TraceServiceHandlerType():
		return p.traceService
	}
	return p
}

func init() {
	servicehub.Register("erda.msp.apm.trace", &servicehub.Spec{
		Services:             pb.ServiceNames(),
		Types:                pb.Types(),
		OptionalDependencies: []string{"service-register"},
		Description:          "",
		ConfigFunc: func() interface{} {
			return &config{}
		},
		Creator: func() servicehub.Provider {
			return &provider{}
		},
	})
}
