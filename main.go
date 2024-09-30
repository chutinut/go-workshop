package main

import (
	greeting "github.com/chutinut/go-workshop/internal-package/0-greeting"
	variable "github.com/chutinut/go-workshop/internal-package/1-variables"
	controlStructure "github.com/chutinut/go-workshop/internal-package/2-control-structure"
)

func main() {
	greeting.SayHello()
	variable.ShowVariables()
	controlStructure.ShowControlStructure()
}
