package model

import "example/user/industrialExercise/internal/service"

//Request is the struct that holds the target user profile url and the operator to deal with it
type Request struct{
	Url      string
	Operator service.Operator
}