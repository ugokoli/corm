package corm

import (
	"errors"
	"fmt"
	"github.com/ugokoli/corm/utility"
	"reflect"
	"strings"
	"time"
)

type BaseModel struct {
	ID        string    `corm:"name:id;partition_key" json:"id"`
	UUID      string    `corm:"partition_key;type:uuid" json:"uuid"`
	CreatedBy string    `corm:"name:created_by" json:"created_by"`
	CreatedAt time.Time `corm:"name:created_at" json:"created_at"`
}

// Parse all fields including fields from injected embedded struct models
func parseModel(model interface{}) ([]columnData, error) {
	rT := reflect.TypeOf(model)
	rV := reflect.ValueOf(model)
	var data []columnData

	// Iterate over all available fields and read the tag value
	for i := 0; i < rT.NumField(); i++ {
		fieldT := rT.Field(i)
		fieldV := rV.Field(i)

		fData := columnData{
			Name: utility.ToSnakeCase(fieldT.Name),
			Type: fieldT.Type.String(),
		}

		// If it is an embedded struct field, otherwise known as an Anonymous field
		if fieldT.Anonymous && fieldV.Kind() == reflect.Struct {
			if result, err := parseModel(fieldV.Interface()); err != nil {
				return nil, err
			} else {
				data = append(data, result...)
			}
			continue
		}

		cormTagStr, ok := fieldT.Tag.Lookup("corm") // follows "name:value;name:value;name" scheme
		if ok {
			if cormTagStr == "-" {
				continue
			}

			// -, column, default, primary_key, type, size, unique, index
			cormTag := strings.Split(cormTagStr, ";")
			for _, config := range cormTag {
				configName := config
				var configValue string
				if strings.Contains(config, ":") {
					tagConfig := strings.Split(config, ":")
					configName = tagConfig[0]
					configValue = tagConfig[1]
				}

				switch strings.ToLower(configName) {
				case "name":
					fData.Name = configValue
					break
				case "type":
					fData.Type = configValue
					break
				case "default":
					fData.Default = configValue
					break
				case "size":
					fData.Size = configValue
					break
				case "partition_key":
					fData.Partition = true
					break
				case "cluster_column":
					fData.Cluster = true
					break
				case "static_column":
					fData.Static = true
					break
				case "unique":
					fData.Unique = true
					break
				case "index":
					fData.Index = true
					break
				}
			}
		}

		// Update fData.Type to cassandra valid type names if supplied type is supported
		var dType string
		if dType, ok = cassandraDataTypes[fData.Type]; !ok {
			return nil, errors.New(fmt.Sprintf("%s.%s data type of %s is not a valid type supported for cassandra data type", rT.Name(), fieldT.Name, fData.Type))
		}
		fData.Type = dType

		var numDefSet int
		columnDef := [3]bool{fData.Partition, fData.Cluster, fData.Static}
		for _, value := range columnDef {
			if value {
				numDefSet++
			}
		}
		if numDefSet > 1 {
			return nil, errors.New(fmt.Sprintf("%s.%s can only be either a partition key, cluster coulmn or static column", rT.Name(), fieldT.Name))
		}

		data = append(data, fData)
	}

	return data, nil
}
