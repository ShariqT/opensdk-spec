package utils


func GenerateClassURL(name string) string {
	return "./class_" + name + ".html"
}

func GenerateInterfaceURL(name string) string {
    return "./interface_" + name + ".html"
}

func GenerateFunctionURL(name string) string {
    return "./function_" + name + ".html"
}

func GenerateExceptionURL(name string) string {
    return "./index.html#" + name
}