//go:generate mockgen -source=manager.go -destination=mock/manager.go -package=mock
package work

import (
	"fmt"
	"main/io"
	"shared/model/entity/job"
	"sync"
)

type Manager interface {
	Do(Id, job.Name, Params) error
	Run()
	GetRunnersCount() int
	StopAll()
}

func NewManager(pins io.Pins, jobs JobMap) Manager {
	return &manager{
		pins:        pins,
		jobs:        jobs,
		runners:     make(map[Id]runner, 0),
		requestChan: make(chan request),
		stopChan:    make(chan struct{}),
	}
}

type manager struct {
	sync.RWMutex
	pins        io.Pins
	jobs        JobMap
	runners     map[Id]runner
	requestChan chan request
	stopChan    chan struct{}
}

type runner struct {
	stop func()
}

type request struct {
	id     Id
	name   job.Name
	params Params
}

func (m *manager) Do(id Id, name job.Name, params Params) error {
	m.Lock()
	defer m.Unlock()

	if _, exists := m.jobs[name]; !exists {
		return fmt.Errorf("job does not exists: %d", name)
	}

	m.requestChan <- request{id, name, params}

	return nil
}

func (m *manager) Run() {
	for {
		select {
		case <-m.stopChan:
			return
		case req := <-m.requestChan:
			stop := m.jobs[req.name](req.id, m.pins, req.params, func() {
				m.Lock()
				delete(m.runners, req.id)
				m.Unlock()
			})
			m.Lock()
			m.runners[req.id] = runner{stop}
			m.Unlock()
		}
	}
}

func (m *manager) StopAll() {
	m.Lock()
	defer m.Unlock()

	m.stopChan <- struct{}{}

	for id, runner := range m.runners {
		runner.stop()
		delete(m.runners, id)
	}
}

func (m *manager) GetRunnersCount() int {
	m.RLock()
	defer m.RUnlock()

	return len(m.runners)
}
