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

package resourcecalculator

import (
	"sort"
	"strings"

	"github.com/pkg/errors"
)

const (
	Dev Workspace = iota
	Test
	Staging
	Prod
)

var Workspaces = []Workspace{Prod, Staging, Test, Dev}

type Workspace int

type Calculator struct {
	ClusterName string
	CPU         *ResourceCalculator
	Mem         *ResourceCalculator
}

type ResourceCalculator struct {
	Type string
	M    map[string]uint64
}

func New(clusterName string) *Calculator {
	return &Calculator{
		ClusterName: clusterName,
		CPU: &ResourceCalculator{
			Type: "CPU",
			M:    make(map[string]uint64),
		},
		Mem: &ResourceCalculator{
			Type: "Memory",
			M:    make(map[string]uint64),
		},
	}
}

func (q *ResourceCalculator) AddValue(value uint64, workspace ...Workspace) {
	workspaces := WorkspacesString(workspace)
	if length := len(workspaces); length == 0 || length > 4 {
		return
	}
	w := strings.Join(workspaces, ":")
	if _, ok := q.M[w]; ok {
		q.M[w] += value
	} else {
		q.M[w] = value
	}
}

func (q ResourceCalculator) TotalQuotable() uint64 {
	var sum uint64
	for _, v := range q.M {
		sum += v
	}
	return sum
}

func (q *ResourceCalculator) TotalForWorkspace(workspace Workspace) uint64 {
	var (
		sum uint64
		w   = WorkspaceString(workspace)
	)
	if w == "" {
		return 0
	}
	for k, v := range q.M {
		if strings.Contains(k, w) {
			sum += v
		}
	}
	return sum
}

func (q *ResourceCalculator) Quota(workspace Workspace, quota uint64) error {
	if totalForWorkspace := q.TotalForWorkspace(workspace); quota > totalForWorkspace {
		return errors.Errorf("the resource %v is not enough, total: %v, your request：%v",
			q.Type, totalForWorkspace, quota)
	}

	// 按优先级减扣
	p := priority(workspace)
	for _, v := range p {
		if q.M[v] >= quota {
			q.M[v] -= quota
			return nil
		}
		quota -= q.M[v]
		q.M[v] = 0
	}
	return nil
}

func WorkspaceString(workspace Workspace) string {
	switch workspace {
	case Prod:
		return "prod"
	case Staging:
		return "staging"
	case Test:
		return "test"
	case Dev:
		return "dev"
	default:
		return ""
	}
}

func WorkspacesString(workspaces []Workspace) []string {
	var m = make(map[Workspace]bool)
	for _, w := range workspaces {
		m[w] = true
	}
	workspaces = []Workspace{}
	for v := range m {
		workspaces = append(workspaces, v)
	}
	sort.Slice(workspaces, func(i, j int) bool {
		return workspaces[i] < workspaces[j]
	})
	var result []string
	for _, v := range workspaces {
		result = append(result, WorkspaceString(v))
	}
	return result
}

func CoreToMillcore(v float64) uint64 {
	return uint64(v * 1000)
}

func MillcoreToCore(v uint64) float64 {
	return float64(v) / 1000
}

func GibibyteToByte(v float64) uint64 {
	return uint64(v * 1024 * 1024 * 1024)
}
func ByteToGibibyte(v uint64) float64 {
	return float64(v) / (1024 * 1024 * 1024)
}

func priority(workspace Workspace) []string {
	switch workspace {
	case Prod:
		return []string{
			"prod",
			"dev:prod", "test:prod", "staging:prod",
			"dev:test:prod", "dev:staging:prod", "test:staging:prod",
			"dev:test:staging:prod",
		}
	case Staging:
		return []string{
			"staging",
			"dev:staging", "test:staging", "staging:prod",
			"dev:test:staging", "dev:staging:prod", "test:staging:prod",
			"dev:test:staging:prod",
		}
	case Test:
		return []string{
			"test",
			"dev:test", "test:staging", "test:prod",
			"dev:test:staging", "dev:test:prod", "test:staging:prod",
			"dev:test:staging:prod",
		}
	case Dev:
		return []string{
			"dev",
			"dev:test", "dev:staging", "dev:prod",
			"dev:test:staging", "dev:test:prod", "dev:staging:prod",
			"dev:test:staging:prod",
		}
	}
	return nil
}
