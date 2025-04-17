package cmd

import (
	"github.com/spf13/cobra"
)

var prompt string

var rootCmd = &cobra.Command{
	Use:   "mytool",
	Short: "MyTool CLI",
	Long:  "MyTool CLI to interact with Kubernetes and OpenAI",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&prompt, "prompt", "p", "", "Prompt for the CLI")
	rootCmd.MarkPersistentFlagRequired("prompt")
}
