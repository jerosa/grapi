package listing

import grapi "github.com/jerosa/grapi/internal"

// Service AddArticle which allows to list articles
type Service interface {
	ListArticles(readListID string) ([]grapi.Article, error)
}

type service struct {
	repository grapi.Repository
}

func (s service) ListArticles(readListID string) ([]grapi.Article, error) {
	return s.repository.FetchArticles(readListID)
}

// NewService creates a new service for listing articles
func NewService(repository grapi.Repository) Service {
	return service{repository: repository}
}
