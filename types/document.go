package types

type SDKDocument struct {
	Version string `mapstructure:"version"`
	Name string `mapstructure:"name"`
	Description string `mapstructure:"description"`
	Constants []ConstantVar	`mapstructure:"constants"`
	Enums []Enum `mapstructure:"enums"`
	Functions []Function `mapstructure:"functions"`
	Classes []Class `mapstructure:"classes"`
	Interfaces []Interface `mapstructure:"interfaces"`
	Errors []Error `mapstructure:"errors"`
}