package corm

import (
	"testing"
	"time"
)

func TestGenerateInsertRecordCQL(t *testing.T) {
	nowTime := time.Now()
	sampleModel := SampleModel{
		BaseModel: BaseModel{
			ID:        "123e4567-e89b-12d3-a456-426655440000",
			CreatedAt: nowTime,
		},
		Password:    "revolve",
		IgnoreMe:    27,
		Email:       "hi@ugokoli.com",
		DateOfBirth: nowTime,
	}

	expectedStmt := "INSERT INTO custom_sample_model_name (id, created_at, password, email, date_of_birth, is_active) VALUES (?, ?, ?, ?, ?, ?)"
	var expectedValues []interface{}
	expectedValues = append(expectedValues, "123e4567-e89b-12d3-a456-426655440000", nowTime, "revolve", "hi@ugokoli.com", nowTime, 1)

	var stmt string
	var values []interface{}

	if stmt, values, _ = generateInsertRecordCQL(sampleModel); stmt != expectedStmt {
		t.Errorf("Unexpected INSERT CQL statement. Expected %s, but got %s", expectedStmt, stmt)
	} else if len(values) != len(expectedValues) {
		t.Errorf("Unexpected INSERT CQL values count. Expected %s count, but got %s", expectedValues, values)
	}

	for i, value := range values {
		if expectedValues[i] != value {
			t.Errorf("Unexpected INSERT CQL value. Expected %v, but got %v", expectedValues[i], value)
		}
	}

	//TODO: perform other test cases
}
