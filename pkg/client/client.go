package client

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	URL string
}

func NewClient(url string) *Client {
	return &Client{
		URL: url,
	}
}

func (s *Client) doRequest(req *http.Request) ([]byte, error) {
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
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("%s", body)
	}
	return body, nil
}

func (s *Client) GetRequest(id string) ([]byte, error) {
	url := fmt.Sprintf("%s/%s", s.URL, id)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	bytes, err := s.doRequest(req)
	if err != nil {
		return nil, err
	}

	// var data model.ZipCodeInfo
	// err = json.Unmarshal(bytes, &data)
	// if err != nil {
	// 	return nil, err
	// }

	return bytes, nil
}
