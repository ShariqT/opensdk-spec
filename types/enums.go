package types

type Enum struct {
	Name string `mapstructure:"name"`
	Type string `mapstructure:"type"`
	Values []EnumValue `mapstructure:"values"`
}

type EnumValue struct {
	Name string `mapstructure:"name"`
	Value string `mapstructure:"value"`
}