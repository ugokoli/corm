package corm

import "testing"

func TestGenerateCreateTableCQL(t *testing.T) {
	expectedResult := "CREATE TABLE IF NOT EXISTS custom_sample_model_name(id varchar, uuid uuid, created_by varchar, created_at timestamp, password varchar, email varchar, date_of_birth timestamp, is_active int, updated_by varchar, updated_at timestamp, PRIMARY KEY((id,uuid))) WITH comment='A sample model';"

	sampleModel := SampleModel{}
	var query string
	if query, _ = generateCreateTableCQL(sampleModel); query != expectedResult {
		t.Errorf("Unexpected CREATE TABLE CQL statement. Expected %s, but got %s", expectedResult, query)
	}

	//TODO: perform other test cases
}
