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

package types

import (
	"strings"

	"github.com/pkg/errors"

	"github.com/erda-project/erda/apistructs"
)

func TestTypeValues() []apistructs.TestType {
	return []apistructs.TestType{apistructs.UT, apistructs.IT}
}

func TestTypeValueOf(tptype string) (apistructs.TestType, error) {
	switch strings.TrimSpace(tptype) {
	case string(apistructs.UT):
		return apistructs.UT, nil
	case string(apistructs.IT):
		return apistructs.IT, nil
	default:
		return "", errors.Errorf("not supported yet %s", tptype)
	}
}
