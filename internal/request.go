/*
 * Copyright (c) 2023, Patrick Wilmes <patrick.wilmes@bit-lake.com>
 *
 * SPDX-License-Identifier: BSD-2-Clause
 */
package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetJson(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(r.Body)

	return json.NewDecoder(r.Body).Decode(target)
}

const (
	PostRequest = "POST"
	PutRequest  = "PUT"
)

func ExecuteWriteRequest(requestKind string, url string, body interface{}) error {
	payloadBuf := new(bytes.Buffer)
	err := json.NewEncoder(payloadBuf).Encode(body)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(requestKind, url, payloadBuf)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)

	responseBody, _ := io.ReadAll(resp.Body)
	fmt.Println("response Body:", string(responseBody))
	return nil
}
