package grapi

import (
	"crypto/rand"
	"time"

	"github.com/oklog/ulid/v2"
)

// Read: https://blog.kowalczyk.info/article/JyRZ/generating-good-unique-ids-in-go.html

// NewULID generate a new ULID
func NewULID() string {
	t := time.Now().UTC()
	id := ulid.MustNew(ulid.Timestamp(t), rand.Reader)

	return id.String()
}
