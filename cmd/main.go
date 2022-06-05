package cmd

import (
	"example/user/industrialExercise/internal/dao"
	"example/user/industrialExercise/internal/model"
	"example/user/industrialExercise/internal/scheduler"
	"example/user/industrialExercise/internal/service"
)

func main(){
	itemChan, err := dao.ItemSaver("user_profile")
	if err != nil{
		panic(err)
	}
	e := service.ConcurrentEngine{
		//Scheduler: &scheduler.SimpleScheduler{},
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      50,
		ItemChan:         itemChan,
		RequestProcessor: service.Worker,
	}

	e.Run(model.Request{
		Url:      "http://localhost:8080/api/",
		Operator: service.UrlOperator{},
	})
}



