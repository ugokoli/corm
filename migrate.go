package corm

import (
	"errors"
	"fmt"
	"github.com/ugokoli/corm/logger"
	"github.com/ugokoli/corm/utility"
	"reflect"
	"strings"
)

func (d *DB) AutoMigrate(models ...interface{}) {
	for _, model := range models {
		var query string
		var err error

		if query, err = generateCreateTableCQL(model); err != nil {
			d.Error = logger.Error("%v", err)
			continue
		}

		if err = d.runMigrateQuery(query); err != nil {
			d.Error = logger.Error("%v", err)
		}
	}
}

func generateCreateTableCQL(model interface{}) (string, error) {
	r := reflect.TypeOf(model)

	tableName := utility.ToSnakeCase(r.Name())
	m, ok := model.(ModelInterface)
	if ok {
		tableName = m.TableName()
	}

	var fields []columnData
	var err error

	if fields, err = parseModel(model); err != nil {
		return "", err
	}

	var columns []string
	var partitionKeys []string
	var clusteringColumns []string
	isStatic := map[bool]string{true: " static", false: ""}
	for _, field := range fields {
		columns = append(columns, fmt.Sprintf("%s %s%s", field.Name, field.Type, isStatic[field.Static]))

		if field.Partition {
			partitionKeys = append(partitionKeys, field.Name)
		}
		if field.Cluster {
			clusteringColumns = append(clusteringColumns, field.Name)
		}
	}

	partitionSpread := strings.Join(partitionKeys, ",")
	if len(partitionKeys) == 0 {
		return "", errors.New(fmt.Sprintf("table %s must have at least one(1) partition key", tableName))
	} else if len(partitionKeys) > 1 {
		partitionSpread = fmt.Sprintf("(%s)", partitionSpread)
	}

	primaryKeys := []string{partitionSpread}
	primaryKeys = append(primaryKeys, clusteringColumns...)

	return fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s(%s, PRIMARY KEY(%s));", tableName, strings.Join(columns, ", "), strings.Join(primaryKeys, ",")), nil
}