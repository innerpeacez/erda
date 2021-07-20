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

package expression

import (
	context "context"
	pb "github.com/erda-project/erda-proto-go/msp/apm/expression/pb"
	"github.com/erda-project/erda/pkg/common/errors"
)

type expressionService struct {
	p *provider
}

func (s *expressionService) GetExpression(ctx context.Context, req *pb.GetExpressionRequest) (*pb.GetExpressionResponse, error) {
	if req.Type == "" {
		return nil, errors.NewMissingParameterError(req.Type)
	}
	var expressions []*pb.Expression
	readExpression(GetFS(req.Type), req.Type, &expressions)
	return &pb.GetExpressionResponse{Data: expressions}, nil
}
