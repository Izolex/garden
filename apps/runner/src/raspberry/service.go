package raspberry

import (
	"encoding/json"
	"fmt"
	"net/http"
	"shared/api"
	"shared/api/request"
	"shared/app/logger"
	jobModel "shared/model/entity/job"
	"shared/model/entity/raspberry"
	"shared/model/entity/work"
	sharedWork "shared/work"
	"time"
)

type Params interface{}
type Service interface {
	Work(raspberry.Entity, jobModel.Name, Params)
	newEntity(uint, jobModel.Name, Params) *work.Entity
	newRequest(raspberry.Entity, *work.Entity) *http.Request
}

type service struct {
	httpClient api.Client
	logger     logger.Logger
	storage    sharedWork.Storage

	requestFactory request.Factory
}

func NewService(httpClient api.Client, logger logger.Logger, storage sharedWork.Storage, requestFactory request.Factory) Service {
	return &service{httpClient, logger, storage, requestFactory}
}

func (s *service) Work(rasp raspberry.Entity, jobName jobModel.Name, params Params) {
	entity := s.newEntity(rasp.ID, jobName, params)
	request := s.newRequest(rasp, entity)
	s.sendRequest(request)
}

func (s *service) newEntity(raspberryId uint, jobName jobModel.Name, params Params) *work.Entity {
	jsonParams, err := json.Marshal(params)
	if err != nil {
		panic(err)
	}
	workEntity := &work.Entity{
		RaspberryId: raspberryId,
		JobId:       uint(jobName),
		Params:      jsonParams,
	}
	err = s.storage.Insert(workEntity)
	if err != nil {
		panic(err)
	}

	return workEntity
}

func (s *service) newRequest(rasp raspberry.Entity, entity *work.Entity) *http.Request {
	var params map[string]interface{}
	err := json.Unmarshal(entity.Params, &params)
	if err != nil {
		panic(err)
	}
	req, err := s.requestFactory.New(http.MethodPost, fmt.Sprintf("http://%s/work", rasp.Address), map[string]interface{}{
		"id":     entity.ID,
		"name":   entity.JobId,
		"params": params,
	})
	if err != nil {
		panic(err)
	}
	return req
}

// todo what if we cannot send request to raspberry???
func (s *service) sendRequest(request *http.Request) {
	var resp *http.Response
	var err error

	retries := 0
	for {
		resp, err = s.httpClient.Do(request)
		if err != nil {
			s.logger.Error(fmt.Errorf("raspberry service: can not send request on %d retry", retries))
			retries++
			if retries > 10 {
				s.logger.Error(fmt.Errorf("raspberry service: can not send request %dx times, what now", retries))
				return
			}
			time.Sleep(1 * time.Second)
			continue
		}
		break
	}

	if resp.StatusCode != http.StatusCreated {
		s.logger.Error(fmt.Errorf("work status %d", resp.StatusCode))
	}
}
