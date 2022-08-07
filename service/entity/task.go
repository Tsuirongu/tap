package entity

import "github.com/Tsuirongu/tap/service/constants"

// Task task...
type Task struct {
	ID          string               `json:"id"` // task uuid
	Avatar      string               `json:"avatar"`
	Description string               `json:"description"` // describe what the task is
	Owner       string               `json:"owner"`
	Status      constants.TaskStatus `json:"taskStatus"`
	Type        constants.TaskType   `json:"type"`
}

// Comment ...
type Comment struct {
	CommentID  string `json:"CommentID"`
	User       string `json:"user"`
	Avatar     string `json:"avatar"`
	Likes      int32  `json:"likes"`
	MarkDown   string `json:"markDown"` // user comment content
	CreateTime string `json:"createTime"`
	DeleteTime string `json:"deleteTime"`
}

type TaskAddReq struct {
	Owner       string               `json:"owner"`
	Avatar      string               `json:"avatar"`
	Description string               `json:"description"` // describe what the task is
	Status      constants.TaskStatus `json:"taskStatus"`
	Type        constants.TaskType   `json:"type"`
}
