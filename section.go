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
	desc:   "Panel to manage goodday project",
	status: ON_TRACK,
	kind:   work,
	tasks: []task{
		{name: "add terminal commands to run different modules", kind: feature, status: todo, priority: p1},
		{name: "establish local db using gorm", kind: feature, status: todo, priority: p1},
		{name: "create terminal ui using tcell", kind: feature, status: todo, priority: p1},
		{name: "mix all tasks, sort and view them from different perspectives", kind: feature, status: todo, priority: p2},
		{name: "display data statistics", kind: feature, status: todo, priority: p2},
	},
}

var cambridge = section{
	name:   "cambridge",
	desc:   "Panel to manage cambridge project",
	status: ON_TRACK,
	kind:   work,
	tasks: []task{
		{name: "use tcell to redesign ui, may also steal some layout of telescope", kind: feature, status: todo, priority: p1},
		{name: "use hashset to manipulate data", kind: feature, status: todo, priority: p1},
		{name: "integrate more dicts, like other cambridge dics & urban", kind: epic, status: todo, priority: p1},
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
		{name: "gorm", kind: action, status: doing, priority: p1},
		{name: "chi", kind: action, status: todo, priority: p1},
		{name: "lf for go", kind: action, status: todo, priority: p1},
		{name: "ytfzf for bash", kind: action, status: todo, priority: p1},
		{name: "dmenu for c", kind: action, status: todo, priority: p2},
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
		{name: "王阳明心学", kind: action, status: doing, priority: p3},
		{name: "一本关于算法的书", kind: action, status: todo, priority: p3},
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
