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

var sections = []section{goodday, cambridge, fakeua, learn, read, solve}

var goodday = section{
	name:   "goodday",
	desc:   "Panel to manage goodday features",
	status: ON_TRACK,
	kind:   work,
	tasks: []task{
		{name: "terminal commands to run different modules", kind: feature, status: todo, priority: p1},
		{name: "establish local db using sqlite and gorm", kind: feature, status: todo, priority: p1},
		{name: "create terminal ui using ncurses", kind: feature, status: todo, priority: p1},
		{name: "mix all tasks, sort and view them from different perspectives", kind: feature, status: todo, priority: p2},
		{name: "display data statistics", kind: feature, status: todo, priority: p2},
		{name: "display random quotes", kind: feature, status: done, priority: p1},
		{name: "rename to goodday", kind: feature, status: done, priority: p1},
	},
}

var cambridge = section{
	name:   "cambridge",
	desc:   "Panel to manage cambridge features",
	status: ON_TRACK,
	kind:   work,
	tasks: []task{
		{name: "terminal args", kind: feature, status: done, priority: p3},
		{name: "remove log file", kind: feature, status: done, priority: p3},
		{name: "add caching mechanism", kind: feature, status: done, priority: p1},
		{name: "establish words db", kind: feature, status: done, priority: p1},
		{name: "use python orm", kind: feature, status: todo, priority: p1},
		{name: "integrate fzf", kind: feature, status: done, priority: p1},
		{name: "integrate curses", kind: feature, status: todo, priority: p1},
		{name: "release version 2.0", kind: epic, status: todo, priority: p1},
		{name: "integrate more dicts, like other cambridge dics & urban", kind: epic, status: todo, priority: p1},
	},
}

var fakeua = section{
	name:   "fakeua",
	desc:   "Panel to manage fakeua features",
	status: ON_TRACK,
	kind:   work,
	tasks: []task{
		{name: "terminal args", kind: feature, status: done, priority: p1},
	},
}

var learn = section{
	name:   "learn",
	desc:   "Track repos to learn",
	status: ON_TRACK,
	kind:   personal,
	tasks: []task{
		{name: "sqlite", kind: action, status: done, priority: p1},
		{name: "gorm", kind: action, status: doing, priority: p1},
		{name: "python orm", kind: action, status: todo, priority: p1},
		{name: "cache mechanism", kind: action, status: done, priority: p1},
		{name: "chi as backend framework", kind: action, status: todo, priority: p1},
		{name: "curses", kind: action, status: done, priority: p1},
		{name: "fzf for go", kind: action, status: todo, priority: p1},
		{name: "ytfzf for bash", kind: action, status: todo, priority: p1},
		{name: "dmenu for c", kind: action, status: todo, priority: p2},
		{name: "write medium article", kind: action, status: todo, priority: p2},
		{name: "make coding videos", kind: action, status: todo, priority: p2},
	},
}

var read = section{
	name:   "read",
	desc:   "Track books and articles to read",
	status: ON_TRACK,
	kind:   personal,
	tasks: []task{
		{name: "What Does It Mean To Have Character", kind: action, status: doing, priority: p3},
		{name: "Suckless Articles", kind: action, status: todo, priority: p3},
		{name: "The Apology", kind: action, status: doing, priority: p3},
		{name: "王阳明心学", kind: action, status: todo, priority: p3},
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
