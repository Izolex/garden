package work

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"shared/model/entity/job"
)

type WorkPOSTEntity struct {
	Id     Id                     `json:"id" binding:"required" validate:"required"`
	Name   job.Name               `json:"name" binding:"required" validate:"required" example:"1"` // Job name
	Params map[string]interface{} `json:"params" binding:"required" validate:"required" example:"{\"pin\":14}"`
}

func NewWorkPOST(manager Manager) gin.HandlerFunc {
	return func(c *gin.Context) {
		entity := new(WorkPOSTEntity)
		if err := c.ShouldBindJSON(entity); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		params, err := json.Marshal(entity.Params)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		err = manager.Do(entity.Id, entity.Name, params)
		if err != nil {
			c.Status(http.StatusNotFound)
			return
		}

		c.Status(http.StatusCreated)
	}
}
