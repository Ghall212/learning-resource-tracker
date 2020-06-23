package repositories

import (
	"database/sql"

	"github.com/jacobtie/learning-resource-tracker/pkg/models"
)

// ResourceRepository ...
type ResourceRepository struct {
	DB *sql.DB
}

// GetTopicResources ...
func (rr *ResourceRepository) GetTopicResources(topicID int) (models.Resources, error) {
	rows, err := rr.DB.Query(`
							SELECT
								r.resource_item_id,
								r.title,
								r.URL,
								r.summary,
								s.label
							FROM
								ResourceItem r
								INNER JOIN State s USING (state_id)
							WHERE s.label <> 'deleted' AND r.topic_id = ?
							;
							`, topicID)
	if err != nil {
		return nil, err
	}

	resources := make(models.Resources, 0)
	for rows.Next() {
		var resource models.Resource
		err = rows.Scan(
			&resource.ResourceID,
			&resource.Title,
			&resource.URL,
			&resource.Summary,
			&resource.State,
		)
		if err != nil {
			return nil, err
		}
		resources = append(resources, resource)
	}

	return resources, nil
}
