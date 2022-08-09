package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go_example/internal/service"
	"github.com/go_example/internal/utils/negotiate"
	"go.uber.org/zap"
	"net/http"
)

type JenkinsController interface {
	Index(c *gin.Context) (int, gin.Negotiate)
}

type jenkinsController struct {
	jenkinsSvc service.JenkinsService
}

func NewJenkinsController(
	jenkinsSvc service.JenkinsService,
) JenkinsController {
	return &jenkinsController{
		jenkinsSvc: jenkinsSvc,
	}
}

func (ctr *jenkinsController) Index(c *gin.Context) (int, gin.Negotiate) {
	taskList, err := ctr.jenkinsSvc.GetQueuePendingTaskList(c)
	if err != nil {
		zap.L().Error("jenkinsSvc.GetQueuePendingTaskList", zap.Error(err))
	}
	for _, taskId := range taskList {
		res, err := ctr.jenkinsSvc.CancelQueueTaskById(c, taskId)
		if err != nil {
			zap.L().Error("jenkinsSvc.CancelQueueTaskById", zap.Error(err))
		}
		fmt.Printf("停止队列中的任务%d,res:%v", taskId, res)
	}
	return negotiate.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"taskNum": taskList,
		},
	})
}
