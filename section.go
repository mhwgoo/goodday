package goodday

import "time"

type section struct {
	name   string
	desc   string
	status statusID
	kind   ptype
	tasks  []task
	start  time.Time
	end    time.Time
}

type statusID int

const (
	ON_TRACK statusID = iota
	AT_RISK
	OFF_TRACK
	ON_HOLD
	COMPLETE
)

type ptype string

var (
	work     ptype = "work"
	personal ptype = "personal"
)

var sections = []section{goodday, cambridge, learn, read, solve}

var goodday = section{
	name:   "goodday",
	desc:   "Panel to manage goodday project",
	status: ON_TRACK,
	kind:   work,
	tasks: []task{
		{name: "add more terminal commands", kind: feature, status: todo, priority: p1},
		{name: "establish local db using gorm", kind: feature, status: todo, priority: p1},
		{name: "create terminal ui using tcell", kind: feature, status: todo, priority: p1},
	},
}

var cambridge = section{
	name:   "cambridge",
	desc:   "Panel to manage cambridge project",
	status: ON_TRACK,
	kind:   work,
	tasks: []task{
		{name: "steal ui from telescope", kind: feature, status: todo, priority: p1},
		{name: "use hashset to manipulate data", kind: feature, status: todo, priority: p1},
		{name: "integrate more dicts, like urban and webster", kind: epic, status: todo, priority: p1},
	},
}

var fakeua = section{
	name:   "fakeua",
	desc:   "Panel to manage fakeua project",
	status: ON_TRACK,
	kind:   work,
	tasks: []task{
		{name: "add terminal args", kind: feature, status: done, priority: p1},
	},
}

var learn = section{
	name:   "learn",
	desc:   "Track repos to learn",
	status: ON_TRACK,
	kind:   personal,
	tasks: []task{
		{name: "chi", kind: action, status: todo, priority: p1},
		{name: "lf for go", kind: action, status: todo, priority: p1},
	},
}

var read = section{
	name:   "read",
	desc:   "Track books and articles to read",
	status: ON_TRACK,
	kind:   personal,
	tasks: []task{
		{name: "Suckless Articles", kind: action, status: todo, priority: p3},
		{name: "The Apology", kind: action, status: doing, priority: p3},
		{name: "传习录", kind: action, status: doing, priority: p3},
		{name: "算法之美", kind: action, status: todo, priority: p3},
	},
}

var solve = section{
	name:   "solve",
	desc:   "Track things to solve",
	status: ON_TRACK,
	kind:   personal,
	tasks: []task{
		{name: "Virtualbox Guest Additions", kind: action, status: todo, priority: p3},
		{name: "Dual Monitors", kind: action, status: todo, priority: p3},
	},
}
