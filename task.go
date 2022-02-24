package achieve

type task struct {
	name string
	desc string
	//dueDate  time.time
	kind     ttype
	status   status
	priority level
}

type ttype string

var (
	epic    ttype = "epic"
	feature ttype = "feature"
	bug     ttype = "bug"
	action  ttype = "action"
)

type status string

var (
	todo  status = "todo"
	doing status = "doing"
	done  status = "done"
)

type level string

var (
	p1 level = "p1"
	p2 level = "p2"
	p3 level = "p3"
)
