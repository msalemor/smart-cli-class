package cmd

import (
	"fmt"
	"mytool/pkg"

	"github.com/spf13/cobra"
)

var azCmd = &cobra.Command{
	Use:   "az",
	Short: "Run Kubernetes commands",
	Long:  "Run Kubernetes commands using the Azure CLI",
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Printf("Running Kubernetes commands with prompt: %s\n", prompt)
		commands, err := pkg.ChatCompletion(prompt)
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		pkg.ProcessCommands(commands)
	},
}

func init() {
	rootCmd.AddCommand(azCmd)
}
