package creating

import grapi "github.com/jerosa/grapi/internal"

// Service CreateReadList wich allows to create reading lists
type Service interface {
	Create(name string, status grapi.Status) (string, error)
}

type service struct {
	repository grapi.Repository
}

func (s service) Create(name string, status grapi.Status) (string, error) {
	id := grapi.NewULID()
	r := grapi.ReadList{ID: id, Name: name, Status: status}
	err := s.repository.Store(r)
	if err != nil {
		return "", err
	}
	return id, nil
}

// NewService creates a new service for creating reading lists
func NewService(repository grapi.Repository) Service {
	return service{repository: repository}
}
