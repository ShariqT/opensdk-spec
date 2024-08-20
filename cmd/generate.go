package cmd

import(
	"fmt"
	"os"
	"github.com/spf13/cobra"

	"github.com/mitchellh/mapstructure"
	"gopkg.in/yaml.v3"
	"github.com/shariqtorres/opensdk-server/types"
	"github.com/shariqtorres/opensdk-server/utils"

)

var sdkDocument types.SDKDocument
var raw interface{}

var generateCmd = &cobra.Command{
	Use: "generate",
	Short: "Generate OpenSDK Server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Generate OpenSDK Server")
		filename := args[0]
		f, err := os.ReadFile(filename)
		if err != nil {
			fmt.Println(err)
			return
		}
		if err := yaml.Unmarshal(f, &raw); err != nil {
			fmt.Println("yaml.Unmarshal failed:", err)
			return
		}
		structuredYamlMap := utils.UnwrapRawYAML(raw)
		decoder, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{WeaklyTypedInput: true, Result: &sdkDocument})
		if err := decoder.Decode(structuredYamlMap); err != nil {
			fmt.Println("mapstructure.Decode failed:", err)
			return
		}
		fmt.Printf("%+v\n", sdkDocument)
	},
}