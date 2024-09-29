package cmd

import(
	"fmt"
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
	Use: "generate [path to OpenSDK yaml file]",
	Short: "Generate OpenSDK Documentation",
	Long: "Generate OpenSDK Documentation in a directory called 'docs' in the current folder. Requires the path to an OpenSDK yaml file.",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
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
		fmt.Println("Generated docs/index.html")
		if len(sdkDocument.Classes) > 0 {
			for _, classInfo := range sdkDocument.Classes {
				contentComponent := templates.ClassPageContent(classInfo)
				fp, err = os.Create("docs/class_" + classInfo.Name + ".html")
				if err != nil {
					fmt.Println(err)
					fmt.Println("Failed to create file for class: " + classInfo.Name)
				}
				err = templates.Page(contentComponent).Render(context.Background(), fp)
				if err != nil {
					fmt.Println(err)
					fmt.Println("Failed to render class: " + classInfo.Name)
					return
				}
			}
		}
		if len(sdkDocument.Interfaces) > 0 {
			for _, interfaceInfo := range sdkDocument.Interfaces {
				contentComponent := templates.InterfacePageContent(interfaceInfo)
				fp, err = os.Create("docs/interface_" + interfaceInfo.Name + ".html")
				if err != nil {
					fmt.Println(err)
					fmt.Println("Failed to create file for interface: " + interfaceInfo.Name)
				}
				err = templates.Page(contentComponent).Render(context.Background(), fp)
				if err != nil {
					fmt.Println(err)
					fmt.Println("Failed to render interface: " + interfaceInfo.Name)	
					return
				}
			}
		}
		if len(sdkDocument.Functions) > 0 {
			for _, functionInfo := range sdkDocument.Functions {
				contentComponent := templates.FunctionPageContent(functionInfo)
				fp, err = os.Create("docs/function_" + functionInfo.Name + ".html")
				if err != nil {
					fmt.Println(err)
					fmt.Println("Failed to create file for function: " + functionInfo.Name)
				}
				err = templates.Page(contentComponent).Render(context.Background(), fp)
				if err != nil {
					fmt.Println(err)
					fmt.Println("Failed to render function: " + functionInfo.Name)
					return
				}
			}
		}
		
	},
}