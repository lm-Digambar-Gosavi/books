package models

type Books struct {
	ID            int     `json:"id"`
	Name          string  `json:"name"`
	AuthorName    string  `json:"author_name"`
	Price         float64 `json:"price"`
	Available     int     `json:"available"`
	Issued        int     `json:"issued"`
	Publisher     string  `json:"publisher"`
	PublishedYear int     `json:"published_year"`
	Description   string  `json:"description"`
}
