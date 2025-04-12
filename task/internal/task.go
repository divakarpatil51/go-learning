package internal

import "strconv"

type Task struct {
	Id          int
	Description string
	CreatedAt   string
	Status      string
	Deleted     bool
}

type TaskStatus int

const (
	TODO TaskStatus = iota
	IN_PROGRESS
	BLOCKED
	DONE
)

var StatusName = map[TaskStatus]string{
	TODO:        "TODO",
	IN_PROGRESS: "IN_PROGRESS",
	BLOCKED:     "BLOCKED",
	DONE:        "DONE",
}

func (t *Task) ToCSVFormat() []string {
	return []string{
		strconv.Itoa(t.Id),
		t.Description,
		t.CreatedAt,
		t.Status,
		strconv.FormatBool(t.Deleted),
	}
}
