package utils

import(
	"github.com/forestgiant/sliceutil"
)

func processRawErrors(data interface{}) []map[string]interface{} {
	results := make([]map[string]interface{},0)
	properties := []string{"description", "namespace", "message"}
	for _, value := range data.([]interface{}) {
		structuredInterfaceMap := make(map[string]interface{})
		for key, value := range value.(map[string]interface{}) {
			if (sliceutil.Contains(properties, key) == false) {
				structuredInterfaceMap["name"] = key
			} else {
				structuredInterfaceMap[key] = value
			}
		}
		results =append(results, structuredInterfaceMap)
	}
return results
}