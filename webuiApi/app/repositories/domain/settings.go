// Package domain contains different business domain DTO's
package domain

// Settings is a collection of settings
type Settings struct {
	Settings []Setting `json:"settings"`
}
type Setting struct {
	Version              string      `json:"version"`
	Region               string      `json:"region"`
	Endpoint             string      `json:"endpoint"`
	UsePathStyleEndpoint bool        `json:"use_path_style_endpoint"`
	Credentials          Credentials `json:"credentials"`
}

// Credentials is a a key and secret pair
type Credentials struct {
	Key    string `json:"key"`
	Secret string `json:"secret"`
}
