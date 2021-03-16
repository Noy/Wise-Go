package wisego

import (
	"bytes"
	"github.com/google/uuid"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type Wise struct {
	SandBox bool
	APIKey, KeyFile string
}

func (w *Wise) sendRequest(method, endPoint string, body io.Reader, withUUID bool) (*http.Response, error, string){
	var apiURL = "https://api.transferwise.com"
	if w.SandBox {
		apiURL = "https://api.sandbox.transferwise.tech"
	}
	req, err := http.NewRequest(method, apiURL+endPoint, body)
	if err != nil {
		return nil, err, ""
	}
	req.Header.Set("Authorization", "Bearer " + w.APIKey)
	var newUUID string
	if withUUID {
		newUUID = uuid.New().String()
		req.Header.Set("X-idempotence-uuid", newUUID)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err, ""
	}
	return resp, nil, newUUID
}

func (w *Wise) BaseTransferWiseGetRequest(endPoint string) ([]byte, error) {
	resp, err, _ := w.sendRequest("GET", endPoint, nil, false)
	if err != nil {
		log.Println(err.Error())
	}
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == 403 {
		log.Println("403 detected, doing secure layer retry.")
		// generating from the current response
		signature, oneTimeToken := generateNewSignedToken(w.KeyFile, resp)
		return w.retryRequestWithToken("GET", endPoint, signature, oneTimeToken, nil, false), nil
	} else {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return body, nil
	}
}

func (w *Wise) PostTransferWiseRequest(endPoint string, withUUID bool, data string) ([]byte, string) {
	resp, err, newUUID := w.sendRequest("POST", endPoint, bytes.NewBufferString(data), withUUID)
	defer resp.Body.Close()
	if err != nil {
		log.Println(err.Error())
	}
	if resp.StatusCode == 403 {
		log.Println("403 detected, doing secure layer retry.")
		signature, oneTimeToken := generateNewSignedToken(w.KeyFile, resp)
		return w.retryRequestWithToken("POST", endPoint, signature, oneTimeToken, bytes.NewBufferString(data), withUUID), newUUID
	} else {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(err.Error())
		}
		return body, newUUID
	}
}