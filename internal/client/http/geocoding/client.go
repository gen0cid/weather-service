package geocoding

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type City struct {
	Name      string  `json:"name"`
	Country   string  `json:"country"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
type Client struct {
	client *http.Client
}

func NewClient(httpClient *http.Client) *Client {
	return &Client{
		client: httpClient,
	}
}

func (c *Client) GetCoords(city string) (City, error) {
	res, err := c.client.Get(
		fmt.Sprintf("https://geocoding-api.open-meteo.com/v1/search?name=%s&count=1&language=ru&format=json",
			city,
		),
	)

	if err != nil {
		return City{}, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return City{}, fmt.Errorf("status code %d", res.StatusCode)
	}

	var geoResp struct {
		Results []City `json:"results"`
	}

	if err := json.NewDecoder(res.Body).Decode(&geoResp); err != nil {
		return City{}, err
	}

	return geoResp.Results[0], nil
}
