package controller

import (
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
	//taskList, err := ctr.jenkinsSvc.GetQueuePendingTaskList(c)
	//if err != nil {
	//	zap.L().Error("jenkinsSvc.GetQueuePendingTaskList", zap.Error(err))
	//}
	//for _, taskId := range taskList {
	//	res, err := ctr.jenkinsSvc.CancelQueueTaskById(c, taskId)
	//	if err != nil {
	//		zap.L().Error("jenkinsSvc.CancelQueueTaskById", zap.Error(err))
	//	}
	//	fmt.Printf("停止队列中的任务%d,res:%v", taskId, res)
	//}

	//nextBuildNum, status, err := ctr.jenkinsSvc.GetJobLatestStatusAndNextBuildNumByName(c, "go-test")
	//if err != nil {
	//	zap.L().Error("jenkinsSvc.GetJobLatestStatusAndNextBuildNumByName", zap.Error(err))
	//}
	//fmt.Println(nextBuildNum, status)
	log, err := ctr.jenkinsSvc.GetJobLastBuildLog(c, "go-test")
	if err != nil {
		zap.L().Error("jenkinsSvc.GetJobLastBuildLog", zap.Error(err))
	}

	//resMap, err := ctr.jenkinsSvc.GetAllJobLatestStatus(c)
	//if err != nil {
	//	zap.L().Error("jenkinsSvc.ValidateAllJobStatus", zap.Error(err))
	//}
	//err := ctr.jenkinsSvc.CreateJobFolder(c, "xudttest")
	//fmt.Println(err)
	//err := ctr.jenkinsSvc.CreateJobFolder(c, "folder1_test", "xudttest")
	//fmt.Println(err)
	//job, err := ctr.jenkinsSvc.CreateJobInFolder(c, "./assets/jenkins/job.xml", "job_test", "xudttest", "folder1_test")
	//fmt.Println(job, err)
	//job := ctr.jenkinsSvc.RenameJob(c, "xudttest/job/folder1_test/job/job_test", "job_test1")
	//fmt.Printf("%+v", job)
	//taskId, err := ctr.jenkinsSvc.BuildJob(c, "go-test", nil)
	//if err != nil {
	//	zap.L().Error("jenkinsSvc.BuildJob", zap.Error(err))
	//}

	return negotiate.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			//"taskNum": taskList,
			//"status":  status,
			"log": log,
			//"resMap": resMap,
			//"taskId": taskId,
		},
	})
}
