package grapi

// ReadList represents the ReadList model
type ReadList struct {
	ID     string
	Name   string
	Status Status
}

// Status type to define the read lists status (enum from proto file)
type Status int

const (
	// Inactive define inactive read list status
	Inactive Status = iota
	// Active define active read list status
	Active
)

// Repository provide operations to operate read lists
type Repository interface {
	Store(ReadList) error
	AddArticle(string, Article) error
	FetchArticles(string) ([]Article, error)
}
