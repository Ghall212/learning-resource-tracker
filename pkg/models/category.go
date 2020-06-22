package models

// Category is a collection of topics
type Category struct {
	CategoryID    int        `json:"categoryID"`
	Label         string     `json:"label"`
	ParentID      *int       `json:"parentID"`
	Depth         int        `json:"depth"`
	SubCategories Categories `json:"subCategories"`
	Topics        Topics     `json:"topics"`
}

// Categories ...
type Categories []Category
