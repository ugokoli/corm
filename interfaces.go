package corm

type ModelInterface interface {
	TableName() string
}

type BeforeCreateInterface interface {
	BeforeCreate() error
}

type CreateWithInterface interface {
	CreateWith() string
}
