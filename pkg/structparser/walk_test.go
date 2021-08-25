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

package structparser

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testWalkType struct {
	a int `tagtagtag`
	b map[string]*bool
	c struct {
		d int
		f struct {
			g int
			h string
		}
	}
}

func TestBottomUpWalk(t *testing.T) {
	tp := reflect.TypeOf(testWalkType{})
	n := newNode(constructCtx{name: tp.Name()}, tp)
	BottomUpWalk(n, func(curr Node, children []Node) {
		fmt.Printf("%+v, %s\n", curr, curr.Name()) // debug print
		extra := curr.Extra()
		*extra = curr.Name()
		for _, c := range children {
			(*extra) = (*extra).(string) + (*c.Extra()).(string)
		}
	})
	assert.Equal(t, "testWalkTypeabcdfgh", (*n.Extra()).(string))
}
