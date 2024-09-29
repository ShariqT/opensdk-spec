package cmd

import(
	"fmt"
	"github.com/spf13/cobra"
)
var rootCmd = &cobra.Command{
	Use: "opensdk",
	Short: "OpenSDK Generator",
	Long: "OpenSDK Generator builds documentation form an OpenSDK specification file",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("OpenSDK Server v0.1.0")
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}	
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}