package pokeapi

import "fmt"

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

func (c *Client) ExploreArea(area string) (exploreResponse, error) {
	var response exploreResponse
	url := fmt.Sprintf("location-area/%s", area)
	err := c.fetchData(&url, &response)
	if err != nil {
		return exploreResponse{}, err
	}

	return response, nil
}
