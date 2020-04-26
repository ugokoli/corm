package corm

type TableNameInterface interface {
	TableName() string
}

type BeforeCreateInterface interface {
	BeforeCreate() error
}

type WithOptionsInterface interface {
	WithOptions() string
}
