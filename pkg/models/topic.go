package models

// Topic is a single topic
type Topic struct {
	TopicID   int       `json:"topicID"`
	Title     string    `json:"title"`
	Summary   *string   `json:"summary"`
	State     string    `json:"state"`
	Priority  Priority  `json:"priority"`
	Resources Resources `json:"resources"`
}

// Topics are a collection of topics
type Topics []Topic
