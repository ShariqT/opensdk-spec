package utils

import(
	"github.com/forestgiant/sliceutil"
)

func processRawClasses(data interface{}) []map[string]interface{} {
	results := make([]map[string]interface{},0)
	properties := []string{"description", "namespace", "implements", "inherits", "properties", "methods"}
	for _, value := range data.([]interface{}) {
		structuredClassMap := make(map[string]interface{})
		for key, value := range value.(map[string]interface{}) {
			if (sliceutil.Contains(properties, key) == false) {
				structuredClassMap["name"] = key
			} else if (key =="methods") {
				structuredClassMap[key] = processRawFunctions(value)
			} else {
				structuredClassMap[key] = value
			}
		}
		results =append(results, structuredClassMap)
	}
return results
}