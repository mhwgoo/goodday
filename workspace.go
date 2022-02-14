package achieve

type workspace struct {
	name     string
	desc     string
	projects []project
	// create time
	// lastMod time
	// owner user
	// members []user
}

var workspaces = []workspace{achieve, personal}

var achieve = workspace{
	name:     "Achieve",
	desc:     "A task management tool written in go",
	projects: []project{Feature, Bug},
}

var personal = workspace{
	name:     "Personal",
	desc:     "Keep track of day-to-day personal tasks",
	projects: []project{Read, Solve},
}
