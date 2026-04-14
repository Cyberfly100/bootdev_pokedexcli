package pokeapi

import "fmt"

func (c *Client) GetAreas(url *string) (AreaResponse, error) {
	var response AreaResponse
	if url == nil {
		url = new(string)
		*url = "location-area?offset=0&limit=20"
	}
	err := c.fetchData(url, &response)
	if err != nil {
		return AreaResponse{}, err
	}

	return response, nil
}

func (c *Client) ExploreArea(area string) (ExploreResponse, error) {
	var response ExploreResponse
	url := fmt.Sprintf("location-area/%s", area)
	err := c.fetchData(&url, &response)
	if err != nil {
		return ExploreResponse{}, err
	}

	return response, nil
}
