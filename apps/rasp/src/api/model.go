//go:generate mockgen -source=model.go -destination=mock/model.go -package=mock
package api

import (
	"encoding/json"
	"gorm.io/gorm"
	"shared/model/database"
)

type Entity struct {
	ID       uint
	Endpoint string
	Body     database.JSON
}

func (Entity) TableName() string {
	return "request"
}

type RequestModel interface {
	Insert(endpoint string, body map[string]interface{}) error
	Get() (*Entity, error)
	Delete(id uint) error
}

func NewRequestModel(db *gorm.DB) RequestModel {
	return &model{db: db}
}

type model struct {
	db *gorm.DB
}

func (model *model) Get() (*Entity, error) {
	var requests []Entity
	result := model.db.Limit(1).Find(&requests)
	if result.Error != nil {
		return nil, result.Error
	}
	if len(requests) == 0 {
		return nil, nil
	}

	return &requests[0], nil
}

func (model *model) Insert(endpoint string, body map[string]interface{}) error {
	bodyJson, err := json.Marshal(body)
	if err != nil {
		return err
	}

	result := model.db.Create(&Entity{
		Endpoint: endpoint,
		Body:     bodyJson,
	})
	return result.Error
}

func (model *model) Delete(id uint) error {
	result := model.db.Delete(&Entity{
		ID: id,
	})
	return result.Error
}
