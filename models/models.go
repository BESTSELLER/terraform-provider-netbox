package models

import "time"

// PathAvailablePrefixes is the path to available prefixes
var PathAvailablePrefixes = "/ipam/prefixes/"

// PathDeviceTypes is the path to device types
var PathDeviceTypes = "/dcim/device-types/"

// ReponseAvailablePrefixes is the response for available prefixes
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

// AvailablePrefixes is the body for available prefixes
type AvailablePrefixes struct {
	PrefixLength int    `json:"prefix_length"`
	Site         int    `json:"site,omitempty"`
	Tenant       int    `json:"tenant,omitempty"`
	Status       string `json:"status,omitempty"`
	Role         int    `json:"role,omitempty"`
	Description  string `json:"description,omitempty"`
}

// ResponeListOfPrefixes is the response for list of prefixes
type ResponeListOfPrefixes struct {
	Count    int                        `json:"count"`
	Next     interface{}                `json:"next"`
	Previous interface{}                `json:"previous"`
	Results  []ReponseAvailablePrefixes `json:"results"`
}

// GetAvailablePrefixResponse is the response for get available prefix
type GetAvailablePrefixResponse struct {
	Prefix         string `json:"prefix"`
	Description    string `json:"description"`
	PrefixLength   int    `json:"prefix_length"`
	PrefixID       int    `json:"prefix_id"`
	ParentPrefixID int    `json:"parent_prefix_id"`
	ID             string `json:"id"`
}

// ResponseSites is the response for sites
type ResponseSites struct {
	Count    int         `json:"count"`
	Next     interface{} `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []struct {
		ID     int    `json:"id"`
		Name   string `json:"name"`
		Slug   string `json:"slug"`
		Status struct {
			Value string `json:"value"`
			Label string `json:"label"`
			ID    int    `json:"id"`
		} `json:"status"`
		Region interface{} `json:"region"`
		Tenant struct {
			ID   int    `json:"id"`
			URL  string `json:"url"`
			Name string `json:"name"`
			Slug string `json:"slug"`
		} `json:"tenant"`
		Facility        string        `json:"facility"`
		Asn             interface{}   `json:"asn"`
		TimeZone        interface{}   `json:"time_zone"`
		Description     string        `json:"description"`
		PhysicalAddress string        `json:"physical_address"`
		ShippingAddress string        `json:"shipping_address"`
		Latitude        interface{}   `json:"latitude"`
		Longitude       interface{}   `json:"longitude"`
		ContactName     string        `json:"contact_name"`
		ContactPhone    string        `json:"contact_phone"`
		ContactEmail    string        `json:"contact_email"`
		Comments        string        `json:"comments"`
		Tags            []interface{} `json:"tags"`
		CustomFields    struct {
		} `json:"custom_fields"`
		Created             string      `json:"created"`
		LastUpdated         time.Time   `json:"last_updated"`
		CircuitCount        interface{} `json:"circuit_count"`
		DeviceCount         interface{} `json:"device_count"`
		PrefixCount         int         `json:"prefix_count"`
		RackCount           interface{} `json:"rack_count"`
		VirtualmachineCount interface{} `json:"virtualmachine_count"`
		VlanCount           interface{} `json:"vlan_count"`
	} `json:"results"`
}

// ResponsePrefix is the response for prefix
type ResponsePrefix struct {
	Count    int         `json:"count"`
	Next     interface{} `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []struct {
		ID     int `json:"id"`
		Family struct {
			Value int    `json:"value"`
			Label string `json:"label"`
		} `json:"family"`
		Prefix string `json:"prefix"`
		Scope  struct {
			ID   int    `json:"id"`
			URL  string `json:"url"`
			Name string `json:"name"`
			Slug string `json:"slug"`
		} `json:"scope"`
		Vrf    interface{} `json:"vrf"`
		Tenant struct {
			ID   int    `json:"id"`
			URL  string `json:"url"`
			Name string `json:"name"`
			Slug string `json:"slug"`
		} `json:"tenant"`
		Vlan   interface{} `json:"vlan"`
		Status struct {
			Value string `json:"value"`
			Label string `json:"label"`
			ID    int    `json:"id"`
		} `json:"status"`
		Role struct {
			ID   int    `json:"id"`
			URL  string `json:"url"`
			Name string `json:"name"`
			Slug string `json:"slug"`
		} `json:"role"`
		IsPool       bool          `json:"is_pool"`
		Description  string        `json:"description"`
		Tags         []interface{} `json:"tags"`
		CustomFields struct {
		} `json:"custom_fields"`
		Created     string    `json:"created"`
		LastUpdated time.Time `json:"last_updated"`
	} `json:"results"`
}

// ResponseDeviceTypes is the response for device types
type ResponseDeviceTypes struct {
	ID           int    `json:"id"`
	DisplayName  string `json:"display"`
	Manufacturer struct {
		ID          int    `json:"id"`
		DisplayName string `json:"display"`
		Name        string `json:"name"`
		Slug        string `json:"slug"`
	} `json:"manufacturer"`
	Model        string      `json:"model"`
	Slug         string      `json:"slug"`
	PartNumber   string      `json:"part_number"`
	Descrption   string      `json:"description"`
	CustomFields interface{} `json:"custom_fields"`
}
