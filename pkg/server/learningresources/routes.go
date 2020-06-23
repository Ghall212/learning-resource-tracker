package learningresources

import (
	"encoding/json"
	"net/http"

	"github.com/jacobtie/learning-resource-tracker/pkg/services"
)

type service interface {
	GetLearningResources() (*services.LearningResources, error)
}

// Router ...
type Router struct {
	Service service
}

// GetAll ...
func (cr *Router) GetAll(w http.ResponseWriter, r *http.Request) {
	categories, err := cr.Service.GetLearningResources()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if err := json.NewEncoder(w).Encode(categories); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
