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

package client

import (
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
	"helm.sh/helm/v3/pkg/repo"

	"github.com/erda-project/erda/bundle"
	"github.com/erda-project/erda/modules/cluster-init/config"
	erdahelm "github.com/erda-project/erda/pkg/helm"
	kc "github.com/erda-project/erda/pkg/k8sclient/config"
)

const (
	defaultRepoName   = "stable"
	InstallModeRemote = "REMOTE"
	RepoModeRemote    = "REMOTE"
	RepoModeLocal     = "LOCAL"
	LocalRepoPath     = "/app/charts"
	ErdaBaseCharts    = "erda-base"
	ErdaAddonsCharts  = "erda-addons"
	ErdaCharts        = "erda"
)

type Option func(client *Client)

type Client struct {
	config *config.Config
}

func New(opts ...Option) *Client {
	c := Client{}
	for _, op := range opts {
		op(&c)
	}

	return &c
}

func WithConfig(cfg *config.Config) Option {
	return func(c *Client) {
		c.config = cfg
	}
}

func (c *Client) Execute() error {
	logrus.Debugf("load config: %+v", c.config)

	opts, err := c.newHelmClientOptions()
	if err != nil {
		return fmt.Errorf("get helm client error: %v", err)
	}

	hc, err := erdahelm.New(opts...)
	if err != nil {
		return err
	}

	switch strings.ToUpper(c.config.RepoMode) {
	case RepoModeRemote:
		// TODO: support repo auth info.
		e := &repo.Entry{Name: defaultRepoName, URL: c.config.RepoURL}

		if err = hc.AddOrUpdateRepo(e); err != nil {
			return err
		}
	}

	if c.config.Reinstall {
		charts := c.getInitCharts()
		for _, chart := range charts {
			chart.Action = erdahelm.ActionUninstall
		}
		m := erdahelm.Manager{
			HelmClient:    hc,
			Charts:        charts,
			LocalRepoName: defaultRepoName,
		}

		if err := m.Execute(); err != nil {
			return err
		}
	}

	m := erdahelm.Manager{
		HelmClient:    hc,
		Charts:        c.getInitCharts(),
		LocalRepoName: defaultRepoName,
	}

	if err := m.Execute(); err != nil {
		return err
	}

	return nil
}

// newHelmClientOptions create helm client options
func (c *Client) newHelmClientOptions() ([]erdahelm.Option, error) {
	opts := make([]erdahelm.Option, 0)

	switch strings.ToUpper(c.config.InstallMode) {
	case InstallModeRemote:
		b := bundle.New(bundle.WithClusterManager())
		cluster, err := b.GetCluster(c.config.TargetCluster)
		if err != nil {
			return nil, err
		}

		rc, err := kc.ParseManageConfig(c.config.TargetCluster, cluster.ManageConfig)
		if err != nil {
			return nil, err
		}

		opts = append(opts, erdahelm.WithRESTClientGetter(erdahelm.NewRESTClientGetterImpl(rc)))
	}

	switch strings.ToUpper(c.config.RepoMode) {
	case RepoModeLocal:
		opts = append(opts, erdahelm.WithLocalChartDiscoverDir(LocalRepoPath))
	}

	return opts, nil
}

func (c *Client) getInitCharts() []*erdahelm.ChartSpec {
	return []*erdahelm.ChartSpec{
		{
			ReleaseName: ErdaBaseCharts,
			ChartName:   ErdaBaseCharts,
			Version:     c.config.Version,
			Action:      erdahelm.ActionInstall,
			Values:      c.config.ErdaBaseValues,
		},
		{
			ReleaseName: ErdaAddonsCharts,
			ChartName:   ErdaAddonsCharts,
			Version:     c.config.Version,
			Action:      erdahelm.ActionInstall,
			Values:      c.config.ErdaAddonsValues,
		},
		{
			ReleaseName: ErdaCharts,
			ChartName:   ErdaCharts,
			Version:     c.config.Version,
			Action:      erdahelm.ActionInstall,
			Values:      c.config.ErdaValues,
		},
	}
}
