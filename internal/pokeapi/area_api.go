package pokeapi

type areaResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (c *Client) GetAreas(url *string) (areaResponse, error) {
	var response areaResponse
	if url == nil {
		url = new(string)
		*url = "location-area?offset=0&limit=20"
	}
	err := c.fetchData(url, &response)
	if err != nil {
		return areaResponse{}, err
	}

	return response, nil
}
