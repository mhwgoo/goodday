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

var workspaces = []workspace{achieve, cambridge, fakeua, personal}

var achieve = workspace{
	name:     "Achieve",
	desc:     "A task management tool written in go",
	projects: []project{achieve_feature, achieve_bug},
}

var cambridge = workspace{
	name:     "Cambridge",
	desc:     "A terminal dictionary written in go",
	projects: []project{cambridge_feature, cambridge_bug},
}

var fakeua = workspace{
	name:     "Fakeua",
	desc:     "Fast random fake useragent written in python",
	projects: []project{fakeua_feature, fakeua_bug},
}

var personal = workspace{
	name:     "Personal",
	desc:     "Keep track of day-to-day personal tasks",
	projects: []project{Read, Solve},
}
