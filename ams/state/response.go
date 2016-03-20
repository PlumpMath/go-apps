// Copyright 2016
package state

type HyperLink struct {
	Relation string `json:"rel"`
	URL      string `json:"url"`
	Method   string `json:"method"`
}

type ResponseState struct {
	State
	Links []HyperLink `json:"links"`
}
