package services

import (
	"sync"

	"github.com/jacobtie/learning-resource-tracker/pkg/models"
)

type categoryRepository interface {
	GetTopCategories() (models.Categories, error)
	GetSubCategories(parentID int, depth int) (models.Categories, error)
}

type topicRepository interface {
	GetTopTopics() (models.Topics, error)
	GetCategoryTopics(categoryID int) (models.Topics, error)
}

type resourceRepository interface {
	GetTopicResources(topicID int) (models.Resources, error)
}

type tagRepository interface {
	GetCategoryTags(categoryID int) (models.Tags, error)
	GetTopicTags(topicID int) (models.Tags, error)
	GetResourceTags(resourceID int) (models.Tags, error)
}

// LearningResourcesService ...
type LearningResourcesService struct {
	CategoryRepository categoryRepository
	TopicRepository    topicRepository
	ResourceRepository resourceRepository
	TagRepository      tagRepository
}

// GetLearningResources ...
func (cr *LearningResourcesService) GetLearningResources() (*LearningResources, error) {
	var categories models.Categories
	var topics models.Topics
	var innerErr error

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		var err error
		categories, err = cr.getTopCategories()
		if err != nil {
			innerErr = err
		}
		wg.Done()
	}()
	go func() {
		var err error
		topics, err = cr.getTopTopics()
		if err != nil {
			innerErr = err
		}
		wg.Done()
	}()
	wg.Wait()

	if innerErr != nil {
		return nil, innerErr
	}

	learningResources := &LearningResources{Categories: categories, Topics: topics}

	return learningResources, nil
}

func (cr *LearningResourcesService) getTopCategories() (models.Categories, error) {
	topCategories, err := cr.CategoryRepository.GetTopCategories()
	if err != nil {
		return nil, err
	}
	cr.hydrateCategories(topCategories, 1)

	return topCategories, nil
}

func (cr *LearningResourcesService) getTopTopics() (models.Topics, error) {
	topTopics, err := cr.TopicRepository.GetTopTopics()
	if err != nil {
		return nil, err
	}
	cr.hydrateTopics(topTopics)

	return topTopics, nil
}
