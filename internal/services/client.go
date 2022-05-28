package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const baseURL string = "https://api.coingecko.com/api/v3"

type Client struct {
	Username string
	Password string
}

type Todo struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
	Done    bool   `json:"done"`
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

func (s *Client) GetTodo() (map[string]interface{}, error) {
	url := fmt.Sprintf(baseURL + "/exchange_rates")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	bytes, err := s.doRequest(req)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	err = json.Unmarshal([]byte(bytes), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
