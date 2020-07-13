package inmemory

import (
	"errors"

	grapi "github.com/jerosa/grapi/internal"
)

var (
	ReadListNotFound      = errors.New("read list not found")
	ReadListAlreadyExists = errors.New("read list already exists")
)

type inmemoryReadListRepo struct {
	readLists map[string]grapi.ReadList
	articles  map[string][]grapi.Article
}

func NewInMemoryReadListRepository() grapi.Repository {
	return inmemoryReadListRepo{
		readLists: make(map[string]grapi.ReadList),
		articles:  make(map[string][]grapi.Article),
	}
}

func (r inmemoryReadListRepo) Store(rl grapi.ReadList) error {
	if _, ok := r.readLists[rl.ID]; !ok {
		return ReadListAlreadyExists
	}
	r.readLists[rl.ID] = rl
	return nil
}

func (r inmemoryReadListRepo) AddArticle(ID string, article grapi.Article) error {
	if _, ok := r.readLists[ID]; !ok {
		return ReadListNotFound
	}
	r.articles[ID] = append(r.articles[ID], article)
	return nil
}

func (r inmemoryReadListRepo) FetchArticles(ID string) ([]grapi.Article, error) {
	if _, ok := r.readLists[ID]; !ok {
		return []grapi.Article{}, ReadListNotFound
	}
	return r.articles[ID], nil
}
