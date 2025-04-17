package main

import (
	"mytool/cmd"
)

func main() {
	// panic if settings are not set
	// pkg.GetSettings()
	// commands := pkg.Commands{
	// 	Commands: []pkg.Command{
	// 		{
	// 			Command:     "kubectl",
	// 			Args:        []string{"get", "all", "-A"},
	// 			Explanation: "List all resources in all namespaces",
	// 		},
	// 	}}
	// pkg.ProcessCommands(&commands)

	// cmds, _ := pkg.ChatCompletion("List all resources in all namespaces")
	// pkg.ProcessCommands(cmds)

	// // Test chat completion
	// cmds, err := pkg.ChatCompletion("List the namespaces")
	// if err != nil {
	// 	fmt.Println("Error: ", err)
	// 	return
	// }
	// fmt.Println("Commands: ", cmds)
	cmd.Execute()
}
