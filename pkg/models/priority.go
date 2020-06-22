package models

// Priority ...
type Priority struct {
	PriorityID int    `json:"priorityID"`
	Label      string `json:"label"`
	Worth      int    `json:"worth"`
}
