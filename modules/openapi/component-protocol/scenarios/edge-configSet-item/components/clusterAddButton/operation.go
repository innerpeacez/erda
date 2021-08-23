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

package clusteraddbutton

import (
	"fmt"

	protocol "github.com/erda-project/erda/modules/openapi/component-protocol"
)

type ComponentClusterAddButton struct {
	ctxBundle protocol.ContextBundle
}

func RenderCreator() protocol.CompRender {
	return &ComponentClusterAddButton{}
}
func (c *ComponentClusterAddButton) SetBundle(ctxBundle protocol.ContextBundle) error {
	if ctxBundle.Bdl == nil {
		return fmt.Errorf("invalie bundle")
	}
	c.ctxBundle = ctxBundle
	return nil
}
