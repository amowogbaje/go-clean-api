package domain

// Media represents our database entity
type Media struct {
	ID   int    `json:"id"`
	URL  string `json:"urls"` // Plural as requested
	Type string `json:"type"` // "image" or "video"
}

// MediaRepository handles database interactions
type MediaRepository interface {
	Create(media *Media) error
	FetchAll() ([]Media, error)
}

// MediaUsecase handles business logic
type MediaUsecase interface {
	ProcessUpload(filename string) (*Media, error)
	GetAllMedia() ([]Media, error)
}