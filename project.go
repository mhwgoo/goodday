package achieve

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

var projects = []project{achieve, cambridge, fakeua, learn, read, solve}

var achieve = project{
	name:   "achieve",
	desc:   "Panel to manage features",
	status: ON_TRACK,
	kind:   work,
	tasks: []task{
		{name: "establish local db using sqlite", kind: feature, status: todo, priority: HIGH},
		{name: "draw terminal dashboard", kind: feature, status: todo, priority: HIGH},
		{name: "enable mouse action on the dashboard", kind: feature, status: todo, priority: HIGH},
		{name: "unicode and emoji support", kind: feature, status: todo, priority: MIDDLE},
		{name: "color and font size support", kind: feature, status: todo, priority: MIDDLE},
		{name: "project and task sort algorithems", kind: feature, status: todo, priority: MIDDLE},
		{name: "data stattistics", kind: feature, status: todo, priority: MIDDLE},
	},
}

var cambridge = project{
	name:   "cambridge",
	desc:   "Panel to manage features",
	status: ON_TRACK,
	kind:   work,
	tasks: []task{
		{name: "terminal args", kind: feature, status: todo, priority: LOW},
		{name: "words db", kind: feature, status: todo, priority: MIDDLE},
	},
}

var fakeua = project{
	name:   "fakeua",
	desc:   "Panel to manage features",
	status: ON_TRACK,
	kind:   work,
	tasks: []task{
		{name: "terminal args", kind: feature, status: done, priority: HIGH},
	},
}

var learn = project{
	name:   "learn",
	desc:   "Track repos to learn",
	status: ON_TRACK,
	kind:   personal,
	tasks: []task{
		{name: "sqlite in go", kind: action, status: todo, priority: HIGH},
		{name: "terminal ui in go", kind: action, status: todo, priority: HIGH},
		{name: "dmenu", kind: action, status: todo, priority: MIDDLE},
		{name: "youtube", kind: action, status: todo, priority: MIDDLE},
	},
}

var read = project{
	name:   "read",
	desc:   "Track books and articles to read",
	status: ON_TRACK,
	kind:   personal,
	tasks: []task{
		{name: "The Apology", kind: action, status: doing, priority: LOW},
		{name: "What Does It Mean To Have Character", kind: action, status: doing, priority: LOW},
		{name: "Suckless Articles", kind: action, status: todo, priority: LOW},
	},
}

var solve = project{
	name:   "solve",
	desc:   "Track things to solve",
	status: ON_TRACK,
	kind:   personal,
	tasks: []task{
		{name: "Virtualbox Guest Additions", kind: action, status: todo, priority: LOW},
		{name: "Dual Monitors", kind: action, status: todo, priority: LOW},
		{name: "Chinese Input", kind: action, status: todo, priority: LOW},
	},
}
