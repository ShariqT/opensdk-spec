package cmd

import(
	"fmt"
	"github.com/spf13/cobra"
)
var rootCmd = &cobra.Command{
	Use: "version",
	Short: "OpenSDK Server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("OpenSDK Server v0.0.1")
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}	
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}