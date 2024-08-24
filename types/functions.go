package types


type FuncParam struct {
	Type string
	Name string
	Description string
}

type FuncReturn struct {
	Type string
	Description string
}


type Function struct {
	Name string
	Description string
	Namespace string
	Params []FuncParam
	Returns []FuncReturn
	OnFailure []string
}