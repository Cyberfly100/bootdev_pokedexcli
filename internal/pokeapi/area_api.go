package pokeapi

type areaResponse struct {
	Count    int     `json:"count"`
	Next     string  `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetAreas(url string) (areaResponse, error) {
	var response areaResponse
	err := fetchData(url, &response)
	if err != nil {
		return areaResponse{}, err
	}

	return response, nil
}
