package achiever

type project struct {
	name   string
	desc   string
	status statusID
	kind   ptype
	tasks  []task
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

var projects = []project{achiever, cambridge, fakeua, learn, read, solve}

var achiever = project{
	name:   "achiever",
	desc:   "Panel to manage achiever features",
	status: ON_TRACK,
	kind:   work,
	tasks: []task{
		{name: "terminal commands to run different modules", kind: feature, status: todo, priority: p1},
		{name: "establish local db using sqlite and gorm", kind: feature, status: todo, priority: p1},
		{name: "create terminal ui using ncurses or some lib like urwid", kind: feature, status: todo, priority: p1},
		{name: "mix all tasks, sort and view them from different perspectives", kind: feature, status: todo, priority: p2},
		{name: "display data statistics", kind: feature, status: todo, priority: p2},
		{name: "display random quotes", kind: feature, status: done, priority: p1},
	},
}

var cambridge = project{
	name:   "cambridge",
	desc:   "Panel to manage cambridge features",
	status: ON_TRACK,
	kind:   work,
	tasks: []task{
		{name: "terminal args", kind: feature, status: done, priority: p3},
		{name: "remove log file", kind: feature, status: done, priority: p3},
		{name: "cache parsing result", desc: "wait till achiever finishes sqlite db", kind: feature, status: todo, priority: p3},
		{name: "words db", desc: "wait till achiever finishes sqlite db", kind: feature, status: todo, priority: p2},
	},
}

var fakeua = project{
	name:   "fakeua",
	desc:   "Panel to manage fakeua features",
	status: ON_TRACK,
	kind:   work,
	tasks: []task{
		{name: "terminal args", kind: feature, status: done, priority: p1},
	},
}

var learn = project{
	name:   "learn",
	desc:   "Track repos to learn",
	status: ON_TRACK,
	kind:   personal,
	tasks: []task{
		{name: "sqlite as db", kind: action, status: done, priority: p1},
		{name: "gorm as orm", kind: action, status: doing, priority: p1},
		{name: "cache", kind: action, status: done, priority: p1},
		{name: "chi as backend framework", kind: action, status: todo, priority: p1},
		{name: "curses for terminal ui", kind: action, status: todo, priority: p1},
		{name: "clash for learning go", kind: action, status: todo, priority: p1},
		{name: "dmenu for learning how to do a project", kind: action, status: todo, priority: p2},
		{name: "write medium article", kind: action, status: todo, priority: p2},
	},
}

var read = project{
	name:   "read",
	desc:   "Track books and articles to read",
	status: ON_TRACK,
	kind:   personal,
	tasks: []task{
		{name: "The Apology", kind: action, status: doing, priority: p3},
		{name: "What Does It Mean To Have Character", kind: action, status: doing, priority: p3},
		{name: "Suckless Articles", kind: action, status: todo, priority: p3},
	},
}

var solve = project{
	name:   "solve",
	desc:   "Track things to solve",
	status: ON_TRACK,
	kind:   personal,
	tasks: []task{
		{name: "Virtualbox Guest Additions", kind: action, status: todo, priority: p3},
		{name: "Dual Monitors", kind: action, status: todo, priority: p3},
	},
}
