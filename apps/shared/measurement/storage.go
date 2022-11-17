//go:generate mockgen -source=storage.go -destination=mock/storage.go -package=mock
package measurement

import (
	"gorm.io/gorm"
	"shared/model/entity/measurement"
)

type Inserter interface {
	Insert(measurement *measurement.Entity) error
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

func (s *storage) Insert(measurement *measurement.Entity) error {
	result := s.db.Create(measurement)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
