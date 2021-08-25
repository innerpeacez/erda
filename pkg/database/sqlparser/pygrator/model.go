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

package pygrator

import (
	"io"
	"strings"

	"github.com/pingcap/parser/ast"
	"github.com/pingcap/parser/mysql"
	"github.com/pkg/errors"

	"github.com/erda-project/erda/pkg/strutil"
)

const ModelPattern = `class {{.ModelName}}(django.db.models.Model):
    """
    generated by erda-cli
    """

    {{range .Fields}}{{.Name}} = django.db.models.{{.Type}}({{.Option}})
    {{end}}
    class Meta:
        db_table = "{{.TableName}}"`

type Model struct {
	ModelName string
	TableName string
	Fields    []*Fields
}

type Fields struct {
	Name   string
	Type   string
	Option string
}

func GenModel(rw io.ReadWriter, model Model) error {
	return generate(rw, "ModelPattern", ModelPattern, model)
}

// ColToDjangoField
// https://docs.djangoproject.com/zh-hans/3.2/ref/models/fields/#model-field-types
func ColToDjangoField(col *ast.ColumnDef) (*Fields, error) {
	if col == nil {
		return nil, errors.New("invalid column, it is nil")
	}
	if col.Tp == nil {
		return nil, errors.New("invalid column, its Tp is nil")
	}
	if col.Name == nil {
		return nil, errors.New("invalid column, its Name is nil")
	}

	var options []string
	var field = Fields{Name: col.Name.String()}
	switch col.Tp.Tp {
	case mysql.TypeDecimal, mysql.TypeNewDecimal:
		field.Type = "DecimalField"
	case mysql.TypeTiny, mysql.TypeBit:
		field.Type = "BooleanField"
	case mysql.TypeShort, mysql.TypeYear:
		field.Type = "IntegerField"
	case mysql.TypeLong, mysql.TypeLonglong:
		field.Type = "BigIntegerField"
	case mysql.TypeFloat, mysql.TypeDouble, mysql.TypeInt24:
		field.Type = "FloatField"
	case mysql.TypeTimestamp, mysql.TypeDatetime:
		field.Type = "DateTimeField"
		for _, opt := range col.Options {
			if opt.Tp == ast.ColumnOptionDefaultValue {
				if expr, ok := opt.Expr.(*ast.FuncCallExpr); ok && strings.EqualFold(expr.FnName.String(), "CURRENT_TIMESTAMP") {
					options = append(options, "auto_now=True")
				}
			}
			if opt.Tp == ast.ColumnOptionOnUpdate {
				if expr, ok := opt.Expr.(*ast.FuncCallExpr); ok && strings.EqualFold(expr.FnName.String(), "CURRENT_TIMESTAMP") {
					options = append(options, "auto_now_add=True")
				}
			}
		}
	case mysql.TypeDate, mysql.TypeNewDate:
		field.Type = "DateField"
	case mysql.TypeDuration:
		field.Type = "TimeField"
	case mysql.TypeVarchar, mysql.TypeEnum, mysql.TypeSet:
		field.Type = "CharField"
	case mysql.TypeJSON:
		field.Type = "JSONField"
	case mysql.TypeBlob, mysql.TypeTinyBlob, mysql.TypeMediumBlob, mysql.TypeLongBlob:
		field.Type = "TextField"
	case mysql.TypeVarString, mysql.TypeString:
		field.Type = "TextField"
	default:
		field.Type = "TextField"
	}
	if len(options) > 0 {
		field.Option = strings.Join(options, ", ")
	}

	return &field, nil
}

func CreateTableStmtToModel(stmt *ast.CreateTableStmt) (*Model, error) {
	if stmt == nil {
		return nil, errors.New("invalid CreateTableStmt, it is nil")
	}

	var model Model
	if stmt.Table == nil {
		return nil, errors.New("invalid CreateTableStmt, Table is nil")
	}
	model.TableName = stmt.Table.Name.String()
	model.ModelName = strutil.SnakeToUpCamel(stmt.Table.Name.String())

	for _, col := range stmt.Cols {
		field, err := ColToDjangoField(col)
		if err != nil {
			return nil, errors.Wrapf(err, "invalid CreateTableStmt, TableName: %s", model.TableName)
		}
		model.Fields = append(model.Fields, field)
	}

	return &model, nil
}
