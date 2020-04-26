package corm

import (
	"testing"
	"time"
)

type SampleModel struct {
	Model
	Password    string    `corm:"name:password" json:"-"`
	IgnoreMe    float32   `corm:"-" json:"email"`
	Email       string    `corm:"name:email;UNIQUE;INDEX" json:"email"`
	DateOfBirth time.Time `corm:"name:date_of_birth;default:CURRENT_TIMESTAMP" json:"date_of_birth"`
	IsActive    int32     `corm:"default:1" json:"is_active"`
	UpdatedBy   string    `corm:"name:updated_by" json:"updated_by"`
	UpdatedAt   time.Time `corm:"name:updated_at" json:"updated_at"`
}

func (a SampleModel) TableName() string {
	return "custom_sample_model_name"
}

func (a SampleModel) WithOptions() string {
	return "WITH comment='A sample model'"
}

func TestGetModelName(t *testing.T) {
	expectedName := "custom_sample_model_name"
	sampleModel := SampleModel{}
	var resultName string

	if resultName = getModelName(sampleModel); resultName != expectedName {
		t.Errorf("Unexpected model name. Expected %s, but got %s", expectedName, resultName)
	}

	//TODO: perform other test cases
}

func TestGetModelWithOptions(t *testing.T) {
	expectedOptions := "WITH comment='A sample model'"
	sampleModel := SampleModel{}
	var resultOptions string

	if resultOptions = getModelWithOptions(sampleModel); resultOptions != expectedOptions {
		t.Errorf("Unexpected model with options. Expected %s, but got %s", expectedOptions, resultOptions)
	}

	//TODO: perform other test cases
}

func TestParseModel(t *testing.T) {
	expectedLength := 10
	sampleModel := SampleModel{}
	var resultFields []columnData

	if resultFields, _ = parseModel(sampleModel); len(resultFields) != expectedLength {
		t.Errorf("Unexpected model field expectedLength. Expected %d, but got %d resultFields", expectedLength, len(resultFields))
	}

	//TODO: perform other test cases
}
