package corm

import (
	"testing"
	"time"
)

type SampleModel struct {
	BaseModel
	Password    string    `corm:"name:password" json:"-"`
	IgnoreMe    float32   `corm:"-" json:"email"`
	Email       string    `corm:"name:email;UNIQUE;INDEX" json:"email"`
	DateOfBirth time.Time `corm:"name:date_of_birth;default:CURRENT_TIMESTAMP" json:"date_of_birth"`
	IsActive    int32     `corm:"default:0" json:"is_active"`
	UpdatedBy   int32     `corm:"name:updated_by" json:"updated_by"`
	UpdatedAt   time.Time `corm:"name:updated_at" json:"updated_at"`
}

// TableName sets the insert table name for this struct type
func (a SampleModel) TableName() string {
	return "custom_sample_model_name"
}

func TestParseModel(t *testing.T) {
	length := 10

	sampleModel := SampleModel{}
	var fields []columnData
	if fields, _ = parseModel(sampleModel); len(fields) != 10 {
		t.Errorf("Unexpected model field length. Expected %d, but got %d fields", length, len(fields))
	}

	//TODO: perform other test cases
}
