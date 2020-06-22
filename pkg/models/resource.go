package models

// Resource ...
type Resource struct {
	ResourceID int     `json:"resourceID"`
	Title      string  `json:"title"`
	URL        string  `json:"url"`
	Summary    *string `json:"summary"`
	State      string  `json:"state"`
}

// Resources ...
type Resources []Resource
