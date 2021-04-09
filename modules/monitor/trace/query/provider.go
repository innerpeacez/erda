package query

import (
	"fmt"

	"github.com/erda-project/erda-infra/base/logs"
	"github.com/erda-project/erda-infra/base/servicehub"
	"github.com/erda-project/erda-infra/providers/cassandra"
	"github.com/gocql/gocql"
)

type define struct{}

func (d *define) Service() []string      { return []string{"trace-query"} }
func (d *define) Dependencies() []string { return []string{"cassandra"} }
func (d *define) Summary() string        { return "trace query" }
func (d *define) Description() string    { return d.Summary() }
func (d *define) Config() interface{}    { return &config{} }
func (d *define) Creator() servicehub.Creator {
	return func() servicehub.Provider {
		return &provider{}
	}
}

type config struct {
	Cassandra cassandra.SessionConfig `file:"cassandra"`
}

type provider struct {
	C                *config
	L                logs.Logger
	cassandraSession *gocql.Session
}

func (p *provider) Init(ctx servicehub.Context) error {
	c := ctx.Service("cassandra").(cassandra.Interface)
	session, err := c.Session(&p.C.Cassandra)
	p.cassandraSession = session
	if err != nil {
		return fmt.Errorf("fail to create cassandra session: %s", err)
	}
	return nil
}

// Start .
func (p *provider) Start() error {
	return nil
}

func (p *provider) Close() error {
	return nil
}

func (p *provider) selectSpans(traceId string, limit int64) []map[string]interface{} {
	query := p.cassandraSession.Query("SELECT * FROM spans WHERE trace_id = ? limit ?", traceId, limit)

	iter := query.Consistency(gocql.All).RetryPolicy(nil).Iter()

	list := make([]map[string]interface{}, 0, 10)
	for row := make(map[string]interface{}, 0); iter.MapScan(row); {
		list = append(list, row)
	}

	return list
}

func init() {
	servicehub.RegisterProvider("trace-query", &define{})
}
