package constants

type TaskStatus int32

type TaskType string

const (
	Active    TaskStatus = 1
	Completed TaskStatus = 2
	Closed    TaskStatus = 3

	Video      TaskType = "video"
	OwnerCheck TaskType = "ownerCheck"
	Text       TaskType = "text"
	Picture    TaskType = "picture"
)
