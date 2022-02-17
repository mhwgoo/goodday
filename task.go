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
	HIGH   level = "high"
	MIDDLE level = "middle"
	LOW    level = "low"
)
