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
		ArgsFile ArgsFile
	}

	AnsibleResult struct {
		Msg     string `json:"msg"`
		Changed bool   `json:"changed"`
		Failed  bool   `json:"failed"`
	}

	ArgsFile struct {
		Name string
		Data []byte
	}
)

func NewAnsibleModule() (*AnsibleModule, error) {

	if len(os.Args) != 2 {
		return nil, fmt.Errorf("no arguments file passed to module")
	}

	argsFile := os.Args[1]

	data, err := ioutil.ReadFile(argsFile)

	if err != nil {
		return nil, fmt.Errorf("could not read configuration file: %s", argsFile)
	}

	return &AnsibleModule{
		Result: AnsibleResult{},
		ArgsFile: ArgsFile{
			Name: argsFile,
			Data: data,
		},
	}, nil
}

func (a *AnsibleModule) Exit() {
	returnResponse(a.Result)
}
func (a *AnsibleModule) Fail(msg string) {
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

func Fail(msg string) {
	result := AnsibleResult{
		Msg:     msg,
		Changed: false,
		Failed:  true,
	}

	returnResponse(result)
}
