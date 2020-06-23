package services

import (
	"errors"
	"sync"

	"github.com/jacobtie/learning-resource-tracker/pkg/models"
)

const maxCategoryDepth int = 4

func (cr *LearningResourcesService) hydrateSubItems(category *models.Category, depth int) error {
	var innerErr error
	nextDepth := depth + 1
	if nextDepth > maxCategoryDepth+1 {
		return errors.New("Exceeded max depth")
	}
	var wg sync.WaitGroup
	wg.Add(2)
	if nextDepth <= maxCategoryDepth {
		go func() {
			if err := cr.hydrateSubCategories(category, nextDepth); err != nil {
				innerErr = err
			}
			wg.Done()
		}()
	} else {
		wg.Done()
	}
	go func() {
		if err := cr.hydrateSubTopics(category); err != nil {
			innerErr = err
		}
		wg.Done()
	}()
	wg.Wait()

	if innerErr != nil {
		return innerErr
	}

	return nil
}

func (cr *LearningResourcesService) hydrateSubCategories(category *models.Category, depth int) error {
	var err error
	category.SubCategories, err = cr.CategoryRepository.GetSubCategories(category.CategoryID, depth)
	if err != nil {
		return err
	}
	if err = cr.hydrateCategories(category.SubCategories, depth); err != nil {
		return err
	}
	return nil
}

func (cr *LearningResourcesService) hydrateSubTopics(category *models.Category) error {
	var err error
	category.Topics, err = cr.TopicRepository.GetCategoryTopics(category.CategoryID)
	if err != nil {
		return err
	}
	if err = cr.hydrateTopics(category.Topics); err != nil {
		return err
	}
	return nil
}

func (cr *LearningResourcesService) hydrateCategories(categories models.Categories, depth int) error {
	var innerErr error

	var wg sync.WaitGroup
	wg.Add(len(categories) * 2)
	for i := range categories {
		index := i
		go func() {
			if err := cr.hydrateSubItems(&categories[index], depth); err != nil {
				innerErr = err
			}
			wg.Done()
		}()
		go func() {
			if err := cr.hydrateCategoryTags(&categories[index]); err != nil {
				innerErr = err
			}
			wg.Done()
		}()
	}
	wg.Wait()

	if innerErr != nil {
		return innerErr
	}

	return nil
}

func (cr *LearningResourcesService) hydrateTopics(topics models.Topics) error {
	var innerErr error
	var wg sync.WaitGroup
	wg.Add(len(topics) * 2)
	for i := range topics {
		index := i
		go func() {
			if err := cr.hydrateSubResources(&topics[index]); err != nil {
				innerErr = err
			}
			wg.Done()
		}()
		go func() {
			if err := cr.hydrateTopicTags(&topics[index]); err != nil {
				innerErr = err
			}
			wg.Done()
		}()
	}
	wg.Wait()

	if innerErr != nil {
		return innerErr
	}

	return nil
}

func (cr *LearningResourcesService) hydrateSubResources(topic *models.Topic) error {
	var err error
	topic.Resources, err = cr.ResourceRepository.GetTopicResources(topic.TopicID)
	if err != nil {
		return err
	}

	var innerErr error
	var wg sync.WaitGroup
	wg.Add(len(topic.Resources))
	for i := range topic.Resources {
		index := i
		go func() {
			if err = cr.hydrateResourceTags(&topic.Resources[index]); err != nil {
				innerErr = err
			}
			wg.Done()
		}()
	}
	wg.Wait()

	if innerErr != nil {
		return innerErr
	}

	return nil
}

func (cr *LearningResourcesService) hydrateCategoryTags(category *models.Category) error {
	var err error
	category.Tags, err = cr.TagRepository.GetCategoryTags(category.CategoryID)
	if err != nil {
		return err
	}
	return nil
}

func (cr *LearningResourcesService) hydrateTopicTags(topic *models.Topic) error {
	var err error
	topic.Tags, err = cr.TagRepository.GetTopicTags(topic.TopicID)
	if err != nil {
		return err
	}
	return nil
}

func (cr *LearningResourcesService) hydrateResourceTags(resource *models.Resource) error {
	var err error
	if resource.Tags, err = cr.TagRepository.GetResourceTags(resource.ResourceID); err != nil {
		return err
	}
	return nil
}
