//Package models contains all struct types used to fetch the data.
package models

// Result is a type used to unmarshal after fetching data from the Here Maps.
type Result struct {
	Results struct {
		Items []Item `json:"items"`
		Next  string `json:"next"`
	} `json:"results"`
	Search Search `json:"search"`
}

// Item is type to fetch one item
type Item struct {
	Position      [2]float32 `json:"position"`
	Distance      float32    `json:"distance"`
	Title         string     `json:"title"`
	Category      Category   `json:"category"`
	AverageRating float32    `json:"averageRating"`
	Icon          string     `json:"icon"`
	Vicinity      string     `json:"vicinity"`
	Having        []string   `json:"having"`
	ChainIds      []string   `json:"chainIds"`
	Type          string     `json:"type"`
	Href          string     `json:"href"`
	ID            string     `json:"id"`
}

type Category struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Type   string `json:"type"`
	System string `json:"system"`
	Href   string `json:"href"`
}

type Search struct {
	Context Context `json:"context"`
}

type Context struct {
	Location Location `json:"location"`
	Href     string   `json:"href"`
	Type     string   `json:"type"`
}

type Location struct {
	Position []float32 `json:"position"`
	Address  Address   `json:"address"`
}

type Address struct {
	Text        string `json:"text"`
	House       string `json:"house"`
	Street      string `jaon:"street"`
	PostalCode  string `json:"postalCode"`
	District    string `json:"district"`
	City        string `json:"city"`
	Country     string `json:"country"`
	County      string `json:"county"`
	StateCode   string `json:"stateCode"`
	CountryCode string `json:"countryCode"`
}
