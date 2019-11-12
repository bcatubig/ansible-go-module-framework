package main

import (
	"encoding/json"
	"fmt"
	module "github.com/bcatubig/ansible-go-module-framework"
)

type ModuleArgs struct {
	Name string
}

func main() {
	m, err := module.NewAnsibleModule()

	if err != nil {
		module.Fail(err.Error())
	}

	var moduleArgs ModuleArgs

	err = json.Unmarshal(m.ArgsFile.Data, &moduleArgs)

	if err != nil {
		module.Fail("Configuration file not valid JSON: ")
	}

	if moduleArgs.Name == "" {
		moduleArgs.Name = "World"
	}

	m.Result.Msg = fmt.Sprintf("Hello, %s!", moduleArgs.Name)
	m.Exit()
}
