package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const PostCode = 201

func PostHeader(url string, msg []byte, headers map[string]string) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, strings.NewReader(string(msg)))
	if err != nil {
		return "", err
	}
	for key, header := range headers {
		req.Header.Set(key, header)
	}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != PostCode {
		return "", errors.New(fmt.Sprintf("%d%s", resp.StatusCode, string(body)))
	}
	return string(body), nil
}

func Post[T any](url string, msg []byte, headers map[string]string) (t T, err error) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, strings.NewReader(string(msg)))
	if err != nil {
		return t, err
	}
	for key, header := range headers {
		req.Header.Set(key, header)
	}
	resp, err := client.Do(req)
	defer resp.Body.Close()

	if resp.StatusCode != PostCode {
		if body, err := ioutil.ReadAll(resp.Body); err != nil {
			log.Println(err)
			return t, errors.New(fmt.Sprintf("%d", resp.StatusCode))
		} else {
			return t, errors.New(fmt.Sprintf("%d%s", resp.StatusCode, string(body)))
		}
	}

	err = json.NewDecoder(resp.Body).Decode(&t)
	if err != nil {
		return t, err
	}
	return
}
