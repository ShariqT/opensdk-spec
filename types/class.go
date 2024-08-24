package types

type Class struct {
	Name string `mapstructure:"name"`
	Description string `mapstructure:"description"`
	Namespace string `mapstructure:"namespace"`
	Implements []string `mapstructure:"implements"`
	Inherits []string `mapstructure:"inherits"`
	Properties []Property `mapstructure:"properties"`
	Methods []Function `mapstructure:"methods"`
}

type Property struct {
	Type string `mapstructure:"type"`
	Name string `mapstructure:"name"`
	Description string `mapstructure:"description"`
}

