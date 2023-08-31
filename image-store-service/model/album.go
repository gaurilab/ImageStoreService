package models

type Album struct {
	ID    int      `json:"id"`
	Name  string   `json:"name"`
	Images []Image `json:"images"`
}

type Image struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

