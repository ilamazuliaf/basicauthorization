package models

type (
	Person struct {
		Name    string `json:"name"`
		Address string `json:"address"`
		Age     int    `json:"age"`
	}
)
