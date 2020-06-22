package services

import "github.com/jacobtie/learning-resource-tracker/pkg/models"

// LearningResources ...
type LearningResources struct {
	Categories models.Categories `json:"categories"`
	Topics     models.Topics     `json:"topics"`
}
