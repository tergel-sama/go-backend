// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import ()

// Contains list of Question
type Question struct {
	Id              int32
	QuestionGroupId int32
	Question        string
}

// Contains list of QuestionAnswer
type QuestionAnswer struct {
	Id         int32
	QuestionId int32
	AnswerText string
	Score      int32
}

// Contains list of QuestionGroup
type QuestionGroup struct {
	Id     int32
	TestId int32
	Name   string
	Order  int32
}

// Contains list of QuestionGroupAnswer
type QuestionGroupAnswer struct {
	Id              int32
	QuestionGroupId int32
	MaxScr          int32
	MinScr          int32
	AnswerText      string
}

// Contains list of test
type Test struct {
	Id            int32
	Name          string
	Desc          string
	Img           string
	Minute        int32
	AgeCls        string
	BeforeDesc    string
	AfterDesc     string
	ExampleReport string
	IsActive      bool
}
