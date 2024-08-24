package utils

import(
	"github.com/forestgiant/sliceutil"
)

func processRawMethods(data interface{}) []map[string]interface{} {
	results := make([]map[string]interface{},0)
	properties := []string{"description", "parameters", "return"}
	for _, value := range data.([]interface{}) {
		structuredMethodMap := make(map[string]interface{})
		for key, value := range value.(map[string]interface{}) {
			if (sliceutil.Contains(properties, key) == false) {
				structuredMethodMap["name"] = key
			} else {
				structuredMethodMap[key] = value
			}
		}
		results =append(results, structuredMethodMap)
	}
return results
}