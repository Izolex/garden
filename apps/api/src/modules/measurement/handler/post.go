package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shared/measurement"
	measurementModel "shared/model/entity/measurement"
)

type POSTEntity struct {
	WorkId      uint    `json:"workId" binding:"required" validate:"required" example:"77"`
	PeripheryId uint    `json:"peripheryId" binding:"required" validate:"required" example:"4"`
	Value       float32 `json:"value" binding:"required" validate:"required"`
}

// @Summary  Measurement result
// @Accept   json
// @Param    string       body    POSTEntity  true  "Measurement informations"
// @Param    X-Signature  header  string      true  "Request signature"
// @Success  201          "Created"
// @Failure  400          "Bad Request"
// @Failure  401          "Unauthorized"
// @Failure  500          "Internal Server&nbsp;Error"
// @Router   /api/v1/measurement [post]
func NewPOST(inserter measurement.Inserter) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestEntity := new(POSTEntity)
		if err := c.ShouldBindJSON(requestEntity); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		err := inserter.Insert(&measurementModel.Entity{
			WorkId:      requestEntity.WorkId,
			PeripheryId: requestEntity.PeripheryId,
			Value:       requestEntity.Value,
		})
		if err != nil {
			panic(err)
		}

		c.Status(http.StatusCreated)
	}
}
