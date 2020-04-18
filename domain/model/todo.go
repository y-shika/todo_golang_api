package model

type (
	// Todo is todo's mdoel.
	Todo struct {
		ID     string `json:"id"`
		Title  string `json:"title"`
		Active bool   `json:"active"`
		Detail string `json:"detail"`
		// createdAt, updatedAtはあるがクライアントにそれを返さなくても良いはずなので要素には加えない
	}
)
