package scheduler

import (
	"example/user/industrialExercise/internal/model"
)

type QueuedScheduler struct{

}

func (s *QueuedScheduler) Run() { //The run func of a single worker
	// TODO: do something
}

func (s *QueuedScheduler) Submit(model.Request) { //The run func of a single worker
	// TODO: do something
}

func (s *QueuedScheduler) WorkerChan() chan model.Request { //The run func of a single worker
	// TODO: do something
	return nil
}