package types


type FuncParam struct {
	Type string `mapstructure:"datatype"`
	Name string `mapstructure:"name"`
	Description string `mapstructure:"description"`
}

type FuncReturn struct {
	Type string `mapstructure:"datatype"`
	Description string `mapstructure:"description"`
}


type Function struct {
	Name string
	Description string
	Namespace string
	Params []FuncParam
	Returns []FuncReturn
	OnFailure []string
}