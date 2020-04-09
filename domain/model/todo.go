package model

type (
	Todo struct {
		// TODO: 要素は仮決め
		ID     string `json:"id"`
		Title  string `json:"title"`
		Active bool   `json:"active"`
		Detail string `json:"detail"`
	}
)
