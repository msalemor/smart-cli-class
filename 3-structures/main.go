package main

import (
	"mytool/cmd"
	"mytool/pkg"
)

func main() {
	// panic if settings are not set
	pkg.GetSettings()
	// // Test chat completion
	// cmds, err := pkg.ChatCompletion("List the namespaces")
	// if err != nil {
	// 	fmt.Println("Error: ", err)
	// 	return
	// }
	// fmt.Println("Commands: ", cmds)
	cmd.Execute()
}
