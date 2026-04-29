package meteocoding

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
	client *http.Client
}

func NewClient(httpClient *http.Client) *Client {
	return &Client{
		client: httpClient,
	}
}

type Response struct {
	Current struct {
		Time        string  `json:"time"`
		Temperature float64 `json:"temperature_2m"`
	}
}

func (c *Client) GetTemperature(lat, long float64) (Response, error) {
	resp, err := c.client.Get(
		fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%v&longitude=%v&current=temperature_2m",
			lat, long),
	)
	if err != nil {
		return Response{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Response{}, fmt.Errorf("status code %d", resp.StatusCode)
	}

	var respStruct Response
	if err := json.NewDecoder(resp.Body).Decode(&respStruct); err != nil {
		return Response{}, err
	}

	return respStruct, nil
}
