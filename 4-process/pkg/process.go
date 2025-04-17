package pkg

import (
	"fmt"
	"os/exec"

	"github.com/fatih/color"
)

func ProcessCommands(commands *Commands) {
	if commands == nil {
		fmt.Println("No commands to process")
		return
	}
	for _, command := range commands.Commands {
		color.Cyan("About to execute: %s %v -> %s\n", command.Command, command.Args, command.Explanation)
		color.Yellow("Do you want to proceed? (y/n): ")

		var response string
		fmt.Scanln(&response)

		if response != "y" && response != "Y" {
			color.Red("Command skipped")
			continue
		}

		fmt.Printf("Executing: %s %v\n", command.Command, command.Args)
		cmd := exec.Command(command.Command, command.Args...)
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Printf("Error: %s\n", err)
		}
		color.Green("Output: %s\n", output)
	}
}
