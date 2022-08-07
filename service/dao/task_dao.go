package dao

import (
	"github.com/Tsuirongu/tap/service/constants"
	"github.com/Tsuirongu/tap/service/entity"
)

// GetTaskList fake dao
func GetTaskList() []entity.Task {
	return []entity.Task{
		{
			ID:          "1",
			Description: "test task",
			Owner:       "savagelu",
			Status:      constants.Active,
			Type:        constants.Text,
		},
		{
			ID:          "2",
			Description: "test task ",
			Owner:       "savagelu",
			Status:      constants.Completed,
			Type:        constants.OwnerCheck,
		},
	}
}

func AddTask(task *entity.Task) error {
	return nil
}
