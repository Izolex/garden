package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"shared/api/request"
	"shared/app/logger"
	"time"
)

func NewSender(
	apiUrl string,
	tickerDuration time.Duration,
	client *http.Client,
	requestFactory request.Factory,
	model RequestModel,
	logger logger.Logger,
) *sender {
	return &sender{
		apiUrl,
		tickerDuration,
		client,
		requestFactory,
		model,
		logger,
	}
}

type sender struct {
	apiUrl         string
	tickerDuration time.Duration
	client         *http.Client
	requestFactory request.Factory
	model          RequestModel
	logger         logger.Logger
}

func (s *sender) Run() {
	ticker := time.NewTicker(s.tickerDuration)
	defer ticker.Stop()

	for {
		<-ticker.C

		row, err := s.model.Get()
		if err != nil {
			panic(err)
		}
		if row == nil {
			continue
		}

		var body map[string]interface{}
		err = json.Unmarshal(row.Body, &body)
		if err != nil {
			s.logger.Error(fmt.Errorf("api worker: %v", err))
			continue
		}

		req, err := s.requestFactory.New(http.MethodPost, s.apiUrl+row.Endpoint, body)
		if err != nil {
			panic(err)
		}

		res, err := s.client.Do(req)
		if err != nil {
			s.logger.Error(fmt.Errorf("api worker: call failed %v", err))
			continue
		}
		if res.StatusCode != http.StatusCreated {
			s.logger.Error(fmt.Errorf("api worker: bad status %d", res.StatusCode))
			continue
		}

		err = s.model.Delete(row.ID)
		if err != nil {
			s.logger.Error(fmt.Errorf("api worker: delete failed for id \"%d\" %v", row.ID, err))
			continue
		}
	}
}
