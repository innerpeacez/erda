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

package schedulepolicy

import (
	"github.com/erda-project/erda/apistructs"
)

// Pass2ScheduleInfo request -> Pass1ScheduleInfo(LabelInfo) ------------> Pass2ScheduleInfo(apistructs.ScheduleInfo)
//                                                            filters
type Pass2ScheduleInfo apistructs.ScheduleInfo

func (p *Pass2ScheduleInfo) validate() error {
	return nil
}

type Pass2ScheduleInfo2 apistructs.ScheduleInfo2

func (p *Pass2ScheduleInfo2) validate() error {
	return nil
}
