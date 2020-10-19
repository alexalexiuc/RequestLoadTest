package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type RequestResponse struct {
	StatusCode int
	Status     string
	RespBody   []byte
}

func DoPostRequest(url string, body interface{}, headers map[string]string) (RequestResponse, LocalError) {
	bodyJSON, err := json.Marshal(body)
	if err != nil {
		return RequestResponse{}, RequestMarshalErr.WithError(err)
	}

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(bodyJSON))
	for key, val := range headers {
		req.Header.Set(key, val)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return RequestResponse{}, RequestExecErr.WithError(err)
	}
	defer resp.Body.Close()
	responseBody, _ := ioutil.ReadAll(resp.Body)

	reqResp := RequestResponse{
		StatusCode: resp.StatusCode,
		Status:     resp.Status,
		RespBody:   responseBody,
	}
	if resp.StatusCode != 200 {
		respErr := RequestRespWarn.WithError(nil)
		respErr.SetAdditionalDetails(string(responseBody))
		return reqResp, respErr
	}
	return reqResp, nil
}
