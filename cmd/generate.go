package cmd

import(
	"fmt"
	// "encoding/json"
	"os"
	"github.com/spf13/cobra"
	"context"
	"github.com/mitchellh/mapstructure"
	"gopkg.in/yaml.v3"
	"github.com/shariqtorres/opensdk-server/types"
	"github.com/shariqtorres/opensdk-server/utils"
	"github.com/shariqtorres/opensdk-server/templates"

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
		// s, _ := json.MarshalIndent(sdkDocument, "", "  ")
		// fmt.Print(string(s))
		// fmt.Printf("%+v\n", sdkDocument)
		os.Mkdir("docs", 0755)
		fp, err := os.Create("docs/index.html")
		if err != nil {
			fmt.Println(err)
			return
		}
		component := templates.IndexPage(sdkDocument)
		
		err = templates.Page(component).Render(context.Background(), fp) 
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Generrated docs/index.html")
	},
}