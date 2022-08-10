package temphttp

import (
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
	req.Header.Add("Authorization", "Bearer 368d4a81335fda597f947d3d45971abc")
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

	return string(body)
}
