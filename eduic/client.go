package eduic

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

const (
	url             = "http://api.frdic.com/api/open/v1/studylist/words"
	contentTypeJson = "application/json"
	language        = "en"
)

var categoryId = "0"

type AuthProvider func() string

type Client struct {
	auth string
}

// NewClient construct a new client by a function which provide the auth
func NewClient(authProvider AuthProvider) *Client {
	auth := authProvider()
	if len(auth) == 0 {
		log.Fatal("Please config your eduic token in .wbcfg.toml, you can found it in \"http://my.eudic.net/OpenAPI/Authorization\"")
	}
	return &Client{auth}
}

// ListWords returns a slice of explainations
func (client *Client) ListWords(page int, size int) ([]Explain, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add("language", language)
	q.Add("category_id", categoryId)
	q.Add("page", strconv.Itoa(page))
	q.Add("page_size", strconv.Itoa(size))
	req.URL.RawQuery = q.Encode()
	commanHeader(req.Header, client.auth)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// handle response
	if resp.StatusCode != 200 {
		msg := fmt.Sprintf("eduic api:ListWords response failed, code: %d", resp.StatusCode)
		return nil, errors.New(msg)
	}

	respByte, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil
	}
	var resBody Response[[]Explain]
	json.Unmarshal(respByte, &resBody)
	return resBody.Data, nil
}

// AddWords will call eduic api to delete words
func (client *Client) DelWords(words ...string) error {
	data := WordBody{categoryId, language, words}
	body, err := json.Marshal(data)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodDelete, url, bytes.NewBuffer(body))
	if err != nil {
		return nil
	}
	commanHeader(req.Header, client.auth)
	req.Header.Add("Content-Type", contentTypeJson)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	// handle response
	if resp.StatusCode != 204 {
		_, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		msg := fmt.Sprintf("eduic api response:DelWords failed, code: %d", resp.StatusCode)
		return errors.New(msg)
	}
	return nil
}

// AddWords will call eduic api to add words
func (client *Client) AddWords(words ...string) (string, error) {
	data := WordBody{categoryId, language, words}
	body, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}
	commanHeader(req.Header, client.auth)
	req.Header.Add("Content-Type", contentTypeJson)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// handle response
	if resp.StatusCode != 201 {
		respByte, err := io.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}
		msg := fmt.Sprintf("eduic api response:AddWords failed, code: %d, body: %s", resp.StatusCode, string(respByte))
		return "", errors.New(msg)
	}

	respByte, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", nil
	}
	var resBody Response[any]
	json.Unmarshal(respByte, &resBody)
	return resBody.Message, nil
}

func commanHeader(header http.Header, auth string) {
	header.Add("Authorization", auth)
	// Remove User-Agent, the default user-agent setting will cause 5oo internal error
	header.Set("User-Agent", "")
}
