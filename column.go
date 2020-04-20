package corm

type (
	// http://cassandra.apache.org/doc/latest/cql/ddl.html#create-table
	// Partition and Cluster are types of PRIMARY KEY
	columnData struct {
		Index     bool
		Name      string
		Type      string
		Size      string
		Partition bool // partition key is a PRIMARY KEY
		Cluster   bool // clustering column is a PRIMARY KEY
		Static    bool // static column is not a PRIMARY KEY
		Unique    bool
		Default   interface{}
	}
)

var cassandraDataTypes = map[string]string{
	"uuid":      "uuid",
	"ascii":     "ascii",
	"string":    "varchar",
	"text":      "text",
	"blob":      "blob",
	"bool":      "boolean",
	"counter":   "counter",
	"inet":      "inet",
	"int":       "varint",
	"int8":      "tinyint",
	"int16":     "smallint",
	"int32":     "int",
	"int64":     "bigint",
	"float32":   "float",
	"float64":   "double",
	"decimal":   "decimal",
	"time":      "time",
	"date":      "date",
	"time.Time": "timestamp",
	"timeuuid":  "timeuuid",
	"duration":  "duration",
}
