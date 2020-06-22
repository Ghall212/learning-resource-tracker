package learningresources

import (
	"encoding/json"
	"net/http"

	"github.com/jacobtie/learning-resource-tracker/pkg/services"
)

type service interface {
	GetLearningResources() services.LearningResources
}

// Router ...
type Router struct {
	CategoryService service
}

// GetAll ...
func (cr *Router) GetAll(w http.ResponseWriter, r *http.Request) {
	categories := cr.CategoryService.GetLearningResources()

	if err := json.NewEncoder(w).Encode(categories); err != nil {
		panic(err)
	}
}