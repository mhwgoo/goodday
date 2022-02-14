package achieve

type task struct {
	name string
	desc string
	//dueDate  time.time
	status   status
	priority levelID
}

type status string

var (
	todo  status = "TODO"
	doing status = "DOING"
	done  status = "DONE"
)

type levelID int

const (
	HIGH levelID = iota
	MIDDLE
	LOW
)
