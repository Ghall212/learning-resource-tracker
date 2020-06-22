package services

import (
	"sync"

	"github.com/jacobtie/learning-resource-tracker/pkg/models"
)

const maxCategoryDepth int = 4

func (cr *LearningResourcesService) hydrateSubItems(category *models.Category, depth int) {
	nextDepth := depth + 1
	if nextDepth > maxCategoryDepth+1 {
		return
	}
	var wg sync.WaitGroup
	wg.Add(2)
	if nextDepth <= maxCategoryDepth {
		go func() {
			cr.hydrateSubCategories(category, nextDepth)
			wg.Done()
		}()
	} else {
		wg.Done()
	}
	go func() {
		cr.hydrateSubTopics(category)
		wg.Done()
	}()
	wg.Wait()
}

func (cr *LearningResourcesService) hydrateSubCategories(category *models.Category, depth int) {
	category.SubCategories = cr.CategoryRepository.GetSubCategories(category.CategoryID, depth)
	var wg sync.WaitGroup
	wg.Add(len(category.SubCategories))
	for i := range category.SubCategories {
		index := i
		go func() {
			cr.hydrateSubItems(&category.SubCategories[index], depth)
			wg.Done()
		}()
	}
	wg.Wait()
}

func (cr *LearningResourcesService) hydrateSubTopics(category *models.Category) {
	category.Topics = cr.TopicRepository.GetCategoryTopics(category.CategoryID)
	var wg sync.WaitGroup
	wg.Add(len(category.Topics))
	for i := range category.Topics {
		index := i
		go func() {
			cr.hydrateSubResources(&category.Topics[index])
			wg.Done()
		}()
	}
	wg.Wait()
}

func (cr *LearningResourcesService) hydrateSubResources(topic *models.Topic) {
	topic.Resources = cr.ResourceRepository.GetTopicResources(topic.TopicID)
}
