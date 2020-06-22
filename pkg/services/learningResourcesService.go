package services

import (
	"sync"

	"github.com/jacobtie/learning-resource-tracker/pkg/models"
)

type categoryRepository interface {
	GetTopCategories() models.Categories
	GetSubCategories(parentID int, depth int) models.Categories
}

type topicRepository interface {
	GetTopTopics() models.Topics
	GetCategoryTopics(categoryID int) models.Topics
}

type resourceRepository interface {
	GetTopicResources(topicID int) models.Resources
}

type tagRepository interface {
	GetCategoryTags(categoryID int) models.Tags
	GetTopicTags(topicID int) models.Tags
	GetResourceTags(resourceID int) models.Tags
}

// LearningResourcesService ...
type LearningResourcesService struct {
	CategoryRepository categoryRepository
	TopicRepository    topicRepository
	ResourceRepository resourceRepository
	TagRepository      tagRepository
}

// GetLearningResources ...
func (cr *LearningResourcesService) GetLearningResources() LearningResources {
	var categories models.Categories
	var topics models.Topics

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		categories = cr.getTopCategories()
		wg.Done()
	}()
	go func() {
		topics = cr.getTopTopics()
		wg.Done()
	}()
	wg.Wait()

	learningResources := LearningResources{Categories: categories, Topics: topics}

	return learningResources
}

func (cr *LearningResourcesService) getTopCategories() models.Categories {
	topCategories := cr.CategoryRepository.GetTopCategories()
	cr.hydrateCategories(topCategories, 1)

	return topCategories
}

func (cr *LearningResourcesService) getTopTopics() models.Topics {
	topTopics := cr.TopicRepository.GetTopTopics()
	cr.hydrateTopics(topTopics)

	return topTopics
}
