package server

import (
	"database/sql"

	"github.com/gorilla/mux"
	"github.com/jacobtie/learning-resource-tracker/pkg/repositories"
	"github.com/jacobtie/learning-resource-tracker/pkg/server/learningresources"
	"github.com/jacobtie/learning-resource-tracker/pkg/services"
)

func mapRoutes(r *mux.Router, db *sql.DB) {
	r.Use(loggingMiddleware)
	r.Use(addJSONHeaderMiddleware)

	mapLearningResourcesRoutes(r, db)
}

func mapLearningResourcesRoutes(r *mux.Router, db *sql.DB) {
	categoryRepository := &repositories.CategoryRepository{DB: db}
	topicRepository := &repositories.TopicRepository{DB: db}
	resourceRepository := &repositories.ResourceRepository{DB: db}
	learningResourcesService := &services.LearningResourcesService{
		CategoryRepository: categoryRepository,
		TopicRepository:    topicRepository,
		ResourceRepository: resourceRepository,
	}
	learningResourcesRouter := &learningresources.Router{Service: learningResourcesService}

	r.HandleFunc("/learningresources", learningResourcesRouter.GetAll)
}
