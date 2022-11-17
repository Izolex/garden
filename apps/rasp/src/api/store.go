package api

import (
	"main/work"
	"shared/model/entity/periphery"
)

type Store interface {
	Measurement(work.Id, periphery.Name, float32)
}

func NewStore(model RequestModel) Store {
	return &store{model}
}

type store struct {
	model RequestModel
}

func (s *store) Measurement(workId work.Id, peripheryName periphery.Name, value float32) {
	s.model.Insert("/api/v1/measurement", map[string]interface{}{
		"workId":      workId,
		"peripheryId": peripheryName,
		"value":       value,
	})
}
