package types

type Enum struct {
	Name string `mapstructure:"name"`
	Description string `mapstructure:"description"`
	Type string `mapstructure:"datatype"`
	Values []EnumValue `mapstructure:"enumValues"`
}

type EnumValue struct {
	Name string `mapstructure:"name"`
	Value string `mapstructure:"value"`
}