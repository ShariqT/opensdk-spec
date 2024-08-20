package utils


import(
	"github.com/forestgiant/sliceutil"
)
func processRawEnums(data interface{}) []map[string]interface{} {
	results := make([]map[string]interface{},0)
	properties := []string{"type", "values"}
	for _, value := range data.([]interface{}) {
		structuredEnumMap := make(map[string]interface{})
		for key, value := range value.(map[string]interface{}) {
			if (sliceutil.Contains(properties, key) == false) {
				structuredEnumMap["name"] = key
			} else {
				structuredEnumMap[key] = value
			}
		}
		results =append(results, structuredEnumMap)
	}
	return results
}
