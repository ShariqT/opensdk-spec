package utils
import(
	"reflect"
)
func UnwrapRawYAML(data interface{}) map[string]interface{} {
	results := make(map[string]interface{})
	switch reflect.TypeOf(data).Kind() {
		case reflect.Map:
			for key, value := range data.(map[string]interface{}) {
				switch key {
					case "constants":
						results[key] = processRawConstants(value)
					case "enums":
						results[key] = processRawEnums(value)
					default:
						results[key] = value
				}
			}
	}
	return results
	
}