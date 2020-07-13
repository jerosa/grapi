package adding

import grapi "github.com/jerosa/grapi/internal"

// Service AddArticle which allows to add articles
type Service interface {
	AddArticle(readListID string, name string, link string) (string, error)
}

type service struct {
	repository grapi.Repository
}

func (s service) AddArticle(readListID string, name string, link string) (string, error) {
	id := grapi.NewULID()
	article := grapi.Article{ID: id, ReadListID: readListID, Name: name, Link: link}
	err := s.repository.AddArticle(readListID, article)
	return id, err
}

// NewService creates a new service for adding articles
func NewService(repository grapi.Repository) Service {
	return service{repository: repository}
}
