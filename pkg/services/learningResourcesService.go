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

// LearningResourcesService ...
type LearningResourcesService struct {
	CategoryRepository categoryRepository
	TopicRepository    topicRepository
	ResourceRepository resourceRepository
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
	var wg sync.WaitGroup
	wg.Add(len(topCategories))
	depth := 1
	for i := range topCategories {
		index := i
		go func() {
			cr.hydrateSubItems(&topCategories[index], depth)
			wg.Done()
		}()
	}
	wg.Wait()

	return topCategories
}

func (cr *LearningResourcesService) getTopTopics() models.Topics {
	topTopics := cr.TopicRepository.GetTopTopics()
	var wg sync.WaitGroup
	wg.Add(len(topTopics))
	for i := range topTopics {
		index := i
		go func() {
			cr.hydrateSubResources(&topTopics[index])
			wg.Done()
		}()
	}
	wg.Wait()

	return topTopics
}
