package temphttp

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Client struct {
	baseURL string
	client  *http.Client
}

func NewClient(baseURL string, client *http.Client) *Client {
	return &Client{
		baseURL: baseURL,
		client:  client,
	}
}

func (c *Client) Convert(numbers []string) string {
	req, err := http.NewRequest(http.MethodPost, c.baseURL+"/add?", nil)
	q := req.URL.Query()

	for _, num := range numbers {
		q.Add("num", fmt.Sprint(num))
	}

	req.URL.RawQuery = q.Encode()
	if err != nil {
		log.Fatal(err)
	}
	res, err := c.client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	var result string
	err = json.Unmarshal(body, &result)
	return result
}
