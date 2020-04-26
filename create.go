package corm

import (
	"fmt"
	"github.com/ugokoli/corm/logger"
	"reflect"
	"strings"
)

func (d *DB) Create(model interface{}) *DB {
	var stmt string
	var values []interface{}
	var err error

	if stmt, values, err = generateInsertRecordCQL(model); err != nil {
		d.Error = logger.Error("%v", err)
		return d
	}

	if err = d.session.Query(stmt, values...).Exec(); err != nil {
		d.Error = logger.Error("%v", err)
		return d
	}

	return d
}

func generateInsertRecordCQL(model interface{}) (string, []interface{}, error) {
	tableName := getModelName(model)

	var fields []columnData
	var columns []string
	var valueTags []string
	var values []interface{}
	var err error

	if fields, err = parseModel(model); err != nil {
		return "", values, err
	}

	for _, field := range fields {
		value := field.Value
		rT := reflect.ValueOf(value)
		if rT.IsZero() {
			if field.Default == nil {
				continue
			}
			value = field.Default
		}

		columns = append(columns, field.Name)
		valueTags = append(valueTags, "?")
		values = append(values, value)
	}

	return fmt.Sprintf(`INSERT INTO %s (%s) VALUES (%s)`, tableName, strings.Join(columns, ", "), strings.Join(valueTags, ", ")), values, nil
}
