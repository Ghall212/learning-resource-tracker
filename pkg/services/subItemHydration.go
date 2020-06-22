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
	wg.Add(3)
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
	go func() {
		cr.hydrateCategoryTags(category)
		wg.Done()
	}()
	wg.Wait()
}

func (cr *LearningResourcesService) hydrateSubCategories(category *models.Category, depth int) {
	category.SubCategories = cr.CategoryRepository.GetSubCategories(category.CategoryID, depth)
	cr.hydrateCategories(category.SubCategories, depth)
}

func (cr *LearningResourcesService) hydrateSubTopics(category *models.Category) {
	category.Topics = cr.TopicRepository.GetCategoryTopics(category.CategoryID)
	cr.hydrateTopics(category.Topics)
}

func (cr *LearningResourcesService) hydrateCategories(categories models.Categories, depth int) {
	var wg sync.WaitGroup
	wg.Add(len(categories))
	for i := range categories {
		index := i
		go func() {
			cr.hydrateSubItems(&categories[index], depth)
			wg.Done()
		}()
	}
	wg.Wait()
}

func (cr *LearningResourcesService) hydrateTopics(topics models.Topics) {
	var wg sync.WaitGroup
	wg.Add(len(topics) * 2)
	for i := range topics {
		index := i
		go func() {
			cr.hydrateSubResources(&topics[index])
			wg.Done()
		}()
		go func() {
			cr.hydrateTopicTags(&topics[index])
			wg.Done()
		}()
	}
	wg.Wait()
}

func (cr *LearningResourcesService) hydrateSubResources(topic *models.Topic) {
	topic.Resources = cr.ResourceRepository.GetTopicResources(topic.TopicID)
	var wg sync.WaitGroup
	wg.Add(len(topic.Resources))
	for i := range topic.Resources {
		index := i
		go func() {
			cr.hydrateResourceTags(&topic.Resources[index])
			wg.Done()
		}()
	}
	wg.Wait()
}

func (cr *LearningResourcesService) hydrateCategoryTags(category *models.Category) {
	category.Tags = cr.TagRepository.GetCategoryTags(category.CategoryID)
}

func (cr *LearningResourcesService) hydrateTopicTags(topic *models.Topic) {
	topic.Tags = cr.TagRepository.GetTopicTags(topic.TopicID)
}

func (cr *LearningResourcesService) hydrateResourceTags(resource *models.Resource) {
	resource.Tags = cr.TagRepository.GetResourceTags(resource.ResourceID)
}
