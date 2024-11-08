package types

type ConstantVar struct {
	Name string `mapstructure:"name"`
	Description string `mapstructure:"description"`
	Type string `mapstructure:"datatype"`
	Value string `mapstructure:"value"`
	Namespace string `mapstructure:"namespace",omitempty`
}

