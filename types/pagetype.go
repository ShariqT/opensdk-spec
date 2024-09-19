package types

type PageType string

const(
	PageTypeIndex PageType = "index"
	PageTypeClass PageType = "class"
	PageTypeFunction PageType = "function"
	PageTypeInterface PageType = "interface"
	PageTypeEnum PageType = "enum"
	PageTypeError PageType = "error"
)