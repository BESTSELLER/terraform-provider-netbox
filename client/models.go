package client

import "time"

type ReponseAvailablePrefixes struct {
	ID     int `json:"id"`
	Family struct {
		Value int    `json:"value"`
		Label string `json:"label"`
	} `json:"family"`
	Prefix string      `json:"prefix"`
	Site   interface{} `json:"site"`
	Vrf    interface{} `json:"vrf"`
	Tenant interface{} `json:"tenant"`
	Vlan   interface{} `json:"vlan"`
	Status struct {
		Value string `json:"value"`
		Label string `json:"label"`
		ID    int    `json:"id"`
	} `json:"status"`
	Role        interface{}   `json:"role"`
	IsPool      bool          `json:"is_pool"`
	Description string        `json:"description"`
	Tags        []interface{} `json:"tags"`
	Created     string        `json:"created"`
	LastUpdated time.Time     `json:"last_updated"`
}
type AvailablePrefixes struct {
	PrefixLength int    `json:"prefix_length"`
	Site         int    `json:"site,omitempty"`
	Tenant       int    `json:"tenant,omitempty"`
	Status       string `json:"status,omitempty"`
	Role         int    `json:"role,omitempty"`
	Description  string `json:"description,omitempty"`
}

type ResponeListOfPrefixes struct {
	Count    int                        `json:"count"`
	Next     interface{}                `json:"next"`
	Previous interface{}                `json:"previous"`
	Results  []ReponseAvailablePrefixes `json:"results"`
}

type GetAvailablePrefixResponse struct {
	Prefix         string `json:"prefix"`
	Description    string `json:"description"`
	PrefixLength   int    `json:"prefix_length"`
	PrefixID       int    `json:"prefix_id"`
	ParentPrefixID int    `json:"parent_prefix_id"`
	ID             string `json:"id"`
}
