package main

import (
	"mytool/cmd"
	"mytool/pkg"
)

func main() {
	pkg.GetSettings()
	cmd.Execute()
}
