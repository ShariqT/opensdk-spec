package utils

import(
	"github.com/forestgiant/sliceutil"
)

func processRawConstants(data interface{}) []map[string]interface{} {
	results := make([]map[string]interface{},0)
	properties := []string{"type", "value", "namespace", "description"}
	for _, value := range data.([]interface{}) {
		structuredConstantMap := make(map[string]interface{})
		for key, value := range value.(map[string]interface{}) {
			if (sliceutil.Contains(properties, key) == false) {
				structuredConstantMap["name"] = key
			} else {
				structuredConstantMap[key] = value
			}
		}
		results =append(results, structuredConstantMap)
	}
	return results
}
