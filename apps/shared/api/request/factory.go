package request

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"shared/api/sign"
)

type Factory interface {
	New(method string, url string, body map[string]interface{}) (*http.Request, error)
}

type factory struct {
	signService sign.Service
}

func NewFactory(signService sign.Service) Factory {
	return &factory{signService}
}

func (f *factory) New(method string, url string, body map[string]interface{}) (*http.Request, error) {
	var err error

	data, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}
	request, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	if body != nil {
		defer request.Body.Close()
	}

	if body != nil {
		request.Header.Set("Content-type", "application/json")
	}
	if method == http.MethodPost {
		request.Header.Set("X-Signature", hex.EncodeToString(f.signService.Sign(request)))
	}

	return request, nil
}
