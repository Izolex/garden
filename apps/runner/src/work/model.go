package work

import (
	"gorm.io/gorm"
	"shared/model/entity/job"
	"shared/model/entity/work"
)

type Model interface {
	GetLast(raspId uint, jobId job.Name) (*work.Entity, error)
}

func NewModel(db *gorm.DB) Model {
	return &model{db}
}

type model struct {
	db *gorm.DB
}

func (m *model) GetLast(raspId uint, jobId job.Name) (*work.Entity, error) {
	var workEntity work.Entity

	result := m.db.Limit(1).Find(&workEntity, &work.Entity{RaspberryId: raspId, JobId: uint(jobId)})
	if result.Error != nil {
		return nil, result.Error
	}

	return &workEntity, nil
}
