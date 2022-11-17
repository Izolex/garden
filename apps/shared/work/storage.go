//go:generate mockgen -source=storage.go -destination=mock/storage.go -package=mock
package work

import (
	"gorm.io/gorm"
	"shared/model/entity/work"
)

type Inserter interface {
	Insert(work *work.Entity) error
}

type Storage interface {
	Inserter
}

func NewStorage(db *gorm.DB) Storage {
	return &storage{db}
}

type storage struct {
	db *gorm.DB
}

func (s *storage) Insert(work *work.Entity) error {
	result := s.db.Create(work)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
