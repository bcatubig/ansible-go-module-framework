package module

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type (
	AnsibleModule struct {
		Result   AnsibleResult
		ArgsFile []byte
	}

	AnsibleResult struct {
		Msg     string `json:"msg"`
		Changed bool   `json:"changed"`
		Failed  bool   `json:"failed"`
	}
)

func NewAnsibleModule(argSpec interface{}) (*AnsibleModule, error) {
	if len(os.Args) != 2 {
		return nil, fmt.Errorf("no argument file provided")
	}

	argsFile := os.Args[1]

	args, err := ioutil.ReadAll(argsFile)

	if err != nil {
		return nil, fmt.Errorf("could not read configuration file: %s", argsFile)
	}

	return &AnsibleModule{
		Result:   AnsibleResult{},
		ArgsFile: args,
	}, nil
}

func (a *AnsibleModule) ExitJSON() {
	returnResponse(a.Result)
}
func (a *AnsibleModule) FailJSON(msg string) {
	a.Result.Msg = msg
	a.Result.Failed = true
	returnResponse(a.Result)
}

func returnResponse(responseBody AnsibleResult) {
	var response []byte
	var err error

	response, err = json.Marshal(responseBody)

	if err != nil {
		response, _ = json.Marshal(&AnsibleResult{
			Msg:     "Invalid response object",
			Changed: false,
			Failed:  true,
		})
		os.Exit(-1)
	}

	fmt.Println(string(response))

	if responseBody.Failed {
		os.Exit(1)
	}

	os.Exit(0)
}
