package request

import (
	"net/http"
	"os"
	"strings"

	"github.com/fatih/color"
)

type Request struct {
	Method string
	Url    string
	Data   string
	Token  string
}

func MakeRequest(request *Request) *http.Response {
	httpRequest, err := http.NewRequest(request.Method, request.Url, strings.NewReader(request.Data))
	if err != nil {
		color.Red("Error creating request: %v", err.Error())
		os.Exit(1)
	}

	httpRequest.Header.Add("Authorization", request.Token)

	response, err := http.DefaultClient.Do(httpRequest)
	if err != nil {
		color.Red("Error sending request: %v", err.Error())
		os.Exit(1)
	}

	return response
}
