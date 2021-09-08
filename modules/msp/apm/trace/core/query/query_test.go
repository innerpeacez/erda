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

package query

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/erda-project/erda-infra/providers/i18n"
	"github.com/erda-project/erda-proto-go/msp/apm/trace/pb"
)

func TestDepthCopyQueryConditions(t *testing.T) {
	tests := []struct {
		name string
		want *pb.TraceQueryConditions
	}{
		{"case1", &TraceQueryConditions},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DepthCopyQueryConditions()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DepthCopyQueryConditions() = %v, want %v", got, tt.want)
			}
			// got point
			gotPoint := getMemoryPoint(got)
			gotPointOthers := getMemoryPoint(got.Others)
			gotPointLimit := getMemoryPoint(got.Limit)
			gotPointSort := getMemoryPoint(got.Sort)
			gotPointTraceStatus := getMemoryPoint(got.TraceStatus)

			// TraceQueryConditions point
			wantPoint := getMemoryPoint(tt.want)
			wantPointOthers := getMemoryPoint(tt.want.Others)
			wantPointLimit := getMemoryPoint(tt.want.Limit)
			wantPointSort := getMemoryPoint(tt.want.Sort)
			wantPointTraceStatus := getMemoryPoint(tt.want.TraceStatus)

			if gotPoint == wantPoint {
				t.Errorf("gotPointServiceName = %v, wantPointServiceName %v", gotPoint, wantPoint)
			}
			if gotPointOthers == wantPointOthers {
				t.Errorf("gotPointOthers = %v, wantPointOthers %v", gotPointOthers, wantPointOthers)
			}
			if gotPointLimit == wantPointLimit {
				t.Errorf("gotPointServiceName = %v, wantPointServiceName %v", gotPointLimit, wantPointLimit)
			}
			if gotPointSort == wantPointSort {
				t.Errorf("gotPointServiceName = %v, wantPointServiceName %v", gotPointSort, wantPointSort)
			}
			if gotPointTraceStatus == wantPointTraceStatus {
				t.Errorf("gotPointServiceName = %v, wantPointServiceName %v", gotPointTraceStatus, wantPointTraceStatus)
			}
		})
	}
}

func getMemoryPoint(need interface{}) string {
	return fmt.Sprintf("%p", need)
}

func Test_clone(t *testing.T) {
	type args struct {
		src *pb.TraceQueryConditions
	}
	tests := []struct {
		name    string
		args    args
		want    *pb.TraceQueryConditions
		wantErr bool
	}{
		{"case1", args{src: &TraceQueryConditions}, &TraceQueryConditions, false},
		{"case2", args{src: nil}, &TraceQueryConditions, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := clone(tt.args.src)
			if (err != nil) != tt.wantErr {
				t.Errorf("clone() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("clone() got = %v, want %v", got, tt.want)
			}
			gotPoint := getMemoryPoint(got)
			wantPoint := getMemoryPoint(tt.want)
			if gotPoint == wantPoint {
				t.Errorf("gotPointServiceName = %v, wantPointServiceName %v", gotPoint, wantPoint)
			}
		})
	}
}

func TestTranslateCondition(t *testing.T) {
	type args struct {
		i18n i18n.Translator
		lang i18n.LanguageCodes
		key  string
	}
	i18n := new(i18n.Translator)
	tests := []struct {
		name string
		args args
		want string
	}{
		{"case1", args{i18n: *i18n, lang: nil, key: "test"}, "test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TranslateCondition(tt.args.i18n, tt.args.lang, tt.args.key); got != tt.want {
				t.Errorf("TranslateCondition() = %v, want %v", got, tt.want)
			}
		})
	}
}
