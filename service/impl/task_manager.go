package impl

import (
	"github.com/Tsuirongu/tap/utils"
	"net/http"

	"github.com/Tsuirongu/tap/service/dao"
	"github.com/Tsuirongu/tap/service/entity"
	"github.com/gin-gonic/gin"
)

// TaskList get Task list
func TaskList(c *gin.Context) {
	user := &entity.User{}
	if err := c.BindJSON(user); err != nil {
		utils.ErrorResponse(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"list": dao.GetTaskList(),
	})
}

func TaskAdd(c *gin.Context) {
	task := &entity.TaskAddReq{}
	if err := c.BindJSON(task); err != nil {
		utils.ErrorResponse(c, err)
		return
	}

}
