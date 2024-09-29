package utils

import(
    "strings"
)

func splitName(name string) []string {
	var partsOfName []string
	partsOfName = strings.Split(name, "#")
	return partsOfName
}
func GenerateClassURL(name string) string {
    if strings.Contains(  name, "#") {
        partsOfName := splitName(name)
        return "./class_" + partsOfName[1] + ".html"
    } else {
        return "./class_" + name + ".html"
    }
}

func GenerateInterfaceURL(name string) string {
    if strings.Contains(  name, "#") {
        partsOfName := splitName(name)
        return "./interface_" + partsOfName[1] + ".html"
    } else {
        return "./interface_" + name + ".html"
    }
}

func GenerateFunctionURL(name string) string {
    if strings.Contains(  name, "#") {
        partsOfName := splitName(name)
        return "./function_" + partsOfName[1] + ".html"
    } else {
        return "./function_" + name + ".html"
    }
}

func GenerateExceptionURL(name string) string {
        return "./index.html?open=errors#" + name
}