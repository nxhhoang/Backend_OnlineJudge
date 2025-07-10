package models

type Problem struct {
	Name      string   `json:"name"`
	ShortName string   `json:"short-name"`
	Tags      []string `json:"tags"`
}
