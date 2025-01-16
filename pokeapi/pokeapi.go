package pokeapi

const (
	BaseUrl     = "https://pokeapi.co/api/v2/"
	LocationUrl = "https://pokeapi.co/api/v2/location-area/"
)

// Structs to unmarshal info from the location area api
type LocationAreaResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
