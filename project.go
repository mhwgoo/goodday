package achieve

type project struct {
	name   string
	desc   string
	status statusID
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

var Feature = project{
	name:   "Features",
	desc:   "Panel to manage features",
	status: ON_TRACK,
	tasks: []task{
		{name: "terminal args", status: todo, priority: HIGH},
		{name: "wrote task", status: todo, priority: HIGH},
	},
}

var Bug = project{
	name:   "Bugs",
	desc:   "Panel to manage bugs",
	status: ON_TRACK,
}

var Learn = project{
	name:   "Learn",
	desc:   "Track repos to learn",
	status: ON_TRACK,
	tasks: []task{
		{name: "Youtube", status: todo, priority: MIDDLE},
		{name: "What Does It Mean To Have Character", status: doing, priority: HIGH},
	},
}

var Read = project{
	name:   "Read",
	desc:   "Track books and articles to read",
	status: ON_TRACK,
	tasks: []task{
		{name: "The Apology", status: doing, priority: MIDDLE},
		{name: "What Does It Mean To Have Character", status: doing, priority: HIGH},
	},
}

var Solve = project{
	name:   "Solve",
	desc:   "Track things to solve",
	status: ON_TRACK,
	tasks: []task{
		{name: "Virtualbox Guest Additions", status: todo, priority: MIDDLE},
		{name: "Dual Monitors", status: todo, priority: MIDDLE},
	},
}
