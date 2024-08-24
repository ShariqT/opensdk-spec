package utils


import(
	"github.com/forestgiant/sliceutil"
)

func processRawFunctions(data interface{}) []map[string]interface{} {
	results := make([]map[string]interface{},0)
	properties := []string{"description", "namespace", "params", "returns", "onfailure"}
	for _, value := range data.([]interface{}) {
		structuredFunctionMap := make(map[string]interface{})
		for key, value := range value.(map[string]interface{}) {
			if (sliceutil.Contains(properties, key) == false) {
				structuredFunctionMap["name"] = key
			} else {
				structuredFunctionMap[key] = value
			}
		}
		results =append(results, structuredFunctionMap)
	}
		return results
}