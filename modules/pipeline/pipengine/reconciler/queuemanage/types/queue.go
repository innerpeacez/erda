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

package types

import (
	"github.com/erda-project/erda-proto-go/pipeline/pb"
	"github.com/erda-project/erda/apistructs"
	"github.com/erda-project/erda/modules/pipeline/spec"
)

type Queue interface {
	ID() string
	IsStrictMode() bool
	OccupiedResource() apistructs.PipelineAppliedResource
	Usage(pipelineCaches map[uint64]*spec.Pipeline) pb.QueueUsage
	AddPipelineIntoQueue(p *spec.Pipeline, doneCh chan struct{})
	PopOutPipeline(p *spec.Pipeline, markAsFailed ...bool)
	Update(pq *apistructs.PipelineQueue)
	RangePendingQueue(mgr QueueManager)
	QueueValidator
}
