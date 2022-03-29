package main

import "time"

type Task struct {
	Name     string
	Desc     string
	Kind     Ttype
	Status   Status
	Priority Level
	Start    time.Time
	End      time.Time
}

type Ttype string

var (
	Epic    Ttype = "epic"
	Feature Ttype = "feature"
	Bug     Ttype = "bug"
	Action  Ttype = "action"
)

type Status string

var (
	Todo    Status = "todo"
	Doing   Status = "doing"
	Done    Status = "done"
	Pending Status = "pending"
)

type Level string

var (
	P1 Level = "p1"
	P2 Level = "p2"
	P3 Level = "p3"
)
