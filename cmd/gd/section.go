package main

import "time"

type Section struct {
	Name   string
	Desc   string
	Status StatusID
	Kind   Ptype
	Tasks  []Task
	Start  time.Time
	End    time.Time
}

type StatusID int

const (
	ON_TRACK StatusID = iota
	AT_RISK
	OFF_TRACK
	ON_HOLD
	COMPLETE
)

type Ptype string

var (
	Work     Ptype = "Work"
	Personal Ptype = "Personal"
)

var Sections = []Section{Goodday, Cambridge, Learn, Read, Solve}

var Goodday = Section{
	Name:   "Goodday",
	Desc:   "Panel to manage Goodday project",
	Status: ON_TRACK,
	Kind:   Work,
	Tasks: []Task{
		{Name: "add more terminal commands", Kind: Feature, Status: Todo, Priority: P1},
		{Name: "add more tables to db", Kind: Feature, Status: Todo, Priority: P1},
		{Name: "create terminal ui using tcell", Kind: Feature, Status: Todo, Priority: P1},
	},
}

var Cambridge = Section{
	Name:   "Cambridge",
	Desc:   "Panel to manage Cambridge project",
	Status: ON_TRACK,
	Kind:   Work,
	Tasks: []Task{
		{Name: "steal ui from telescope", Kind: Feature, Status: Todo, Priority: P1},
		{Name: "use hashset to manipulate data", Kind: Feature, Status: Todo, Priority: P1},
		{Name: "integrate more dicts, like urban and webster", Kind: Epic, Status: Todo, Priority: P1},
	},
}

var fakeua = Section{
	Name:   "fakeua",
	Desc:   "Panel to manage fakeua project",
	Status: ON_TRACK,
	Kind:   Work,
	Tasks: []Task{
		{Name: "add terminal args", Kind: Feature, Status: Done, Priority: P1},
	},
}

var Learn = Section{
	Name:   "Learn",
	Desc:   "Track repos to Learn",
	Status: ON_TRACK,
	Kind:   Personal,
	Tasks: []Task{
		{Name: "chi for web dev", Kind: Action, Status: Todo, Priority: P1},
		{Name: "lf for go", Kind: Action, Status: Todo, Priority: P1},
	},
}

var Read = Section{
	Name:   "Read",
	Desc:   "Track books and articles to Read",
	Status: ON_TRACK,
	Kind:   Personal,
	Tasks: []Task{
		{Name: "Suckless Articles", Kind: Action, Status: Todo, Priority: P3},
		{Name: "The Apology", Kind: Action, Status: Doing, Priority: P3},
		{Name: "传习录", Kind: Action, Status: Doing, Priority: P3},
		{Name: "算法之美", Kind: Action, Status: Todo, Priority: P3},
	},
}

var Solve = Section{
	Name:   "Solve",
	Desc:   "Track things to Solve",
	Status: ON_TRACK,
	Kind:   Personal,
	Tasks: []Task{
		{Name: "Virtualbox Guest Additions", Kind: Action, Status: Todo, Priority: P3},
		{Name: "Dual Monitors", Kind: Action, Status: Todo, Priority: P3},
	},
}
