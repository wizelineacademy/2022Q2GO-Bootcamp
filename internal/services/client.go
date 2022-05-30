package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/krmirandas/2022Q2GO-Bootcamp/internal/model"
)

const baseURL string = "https://api.zippopotam.us"

type Client struct {
	Username string
	Password string
}

func NewBasicAuthClient(username, password string) *Client {
	return &Client{
		Username: username,
		Password: password,
	}
}

func (s *Client) doRequest(req *http.Request) ([]byte, error) {
	req.SetBasicAuth(s.Username, s.Password)
	req.Header.Add("accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if 200 != resp.StatusCode {
		return nil, fmt.Errorf("%s", body)
	}
	return body, nil
}

func (s *Client) GetTodo(id string) (*model.ZipCodeInfo, error) {
	url := fmt.Sprintf(baseURL+"/mx/%s", id)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	bytes, err := s.doRequest(req)
	if err != nil {
		return nil, err
	}

	var data model.ZipCodeInfo
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
