package service

import "example/user/industrialExercise/internal/model"

type Operator interface{
	Operate(contents []byte, url string) Result
}

//UrlOperator is an implementation of Operator
type UrlOperator struct{
}

func (UrlOperator) Operate(contents []byte, url string) Result {
	// TODO: add business logic
	return Result{}
}


type Item struct{
	Username string
	Age string
	Id string
}

type Result struct{
	Items []Item
}


type ConcurrentEngine struct{
	Scheduler        Scheduler
	WorkerCount      int
	ItemChan         chan Item
	RequestProcessor Processor //adding processor so it can be referenced directly from this struct
}

func(e *ConcurrentEngine) Run(seeds ...model.Request){
	e.Scheduler.Run()
}

type Processor func(r model.Request) (Result, error)

type Scheduler interface{
	Submit(model.Request)
	WorkerChan() chan model.Request
	Run()
}

func Worker(r model.Request) (Result, error){
	var body []byte
	return r.Operator.Operate(body, r.Url), nil
}