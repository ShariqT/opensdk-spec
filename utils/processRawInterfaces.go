package utils

import(
	"github.com/forestgiant/sliceutil"
)

func processRawInterfaces(data interface{}) []map[string]interface{} {
	results := make([]map[string]interface{},0)
	properties := []string{"description", "namespace", "methods"}
	for _, value := range data.([]interface{}) {
		structuredInterfaceMap := make(map[string]interface{})
		for key, value := range value.(map[string]interface{}) {
			if (sliceutil.Contains(properties, key) == false) {
				structuredInterfaceMap["name"] = key
			} else if (key == "methods") {
				structuredInterfaceMap[key] = processRawFunctions(value)
			} else {
				structuredInterfaceMap[key] = value
			}
		}
		results =append(results, structuredInterfaceMap)
	}
return results
}