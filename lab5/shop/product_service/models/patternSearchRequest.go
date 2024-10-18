package models

type PatternSearchRequest struct {
	NamePattern     string `json:"namePattern" example:"%_%"`
	LastNamePattern string `json:"lastNamePattern" example:"%_%"`
}
