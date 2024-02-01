package model

type ElementType string

type WorkSuitability struct {
	ID   string      `json:"id"`
	Name string      `json:"name"`
	Type ElementType `json:"element_type"`
}
