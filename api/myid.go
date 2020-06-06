package api

import (
	"bytes"
	"encoding/json"
	"github.com/myidapp/backend/models"
	"io/ioutil"
	"net/http"
	"strings"
)

type ResponseInterface struct {
	Error     string      `json:"error"`
	Result    interface{} `json:"result"`
}


type MYIDAPI struct {
	url    string
	client *http.Client
}

func (a *MYIDAPI) Init(url string) {
	a.url = url
	a.client = &http.Client{}
}

func (a *MYIDAPI) Request(method string, endpoint string, data interface{}) ([]byte, error) {
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(data)
	req, e := http.NewRequest(method, a.url+endpoint, buf)
	if e != nil {
		return []byte{}, e
	}

	resp, e := a.client.Do(req)
	if e != nil {
		return []byte{}, e
	}
	defer resp.Body.Close()
	body, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		return []byte{}, e
	}
	return body, nil
}

func (a *MYIDAPI) Clear() (bool, error) {
	resp, err := a.Request("POST", "/clear", nil)
	if err != nil {
		return false, err
	}

	var response ResponseInterface
	json.NewDecoder(strings.NewReader(string(resp))).Decode(&response)

	return response.Result.(bool), nil
}

func (a *MYIDAPI) GetSignature(uuid string) (models.Signature, error) {
	resp, err := a.Request("GET", "/sign/"+uuid, nil)

	if err != nil {
		return models.Signature{}, err
	}
	var response struct {
		Error   string `json:"error"`
		Result  models.Signature `json:"result"`
	}
	json.NewDecoder(strings.NewReader(string(resp))).Decode(&response)

	return response.Result, nil
}

func (a *MYIDAPI) PostSignature(uuid string, signature string, public_key string) (bool, error) {
	resp, err := a.Request("POST", "/sign", models.Signature{UUID: uuid, Signature: signature, PublicKey: public_key})
	if err != nil {
		return false, err
	}

	var response ResponseInterface
	json.NewDecoder(strings.NewReader(string(resp))).Decode(&response)

	return response.Result.(bool), nil
}

