package service

import (
	"github.com/gin-gonic/gin"
	"github.com/go_example/internal/meta"
)

type JenkinsService interface {
	GetQueuePendingTaskList(c *gin.Context) (taskIdList []int64, err error)
	CancelQueueTaskById(c *gin.Context, taskId int64) (bool, error)
	GetJobStatus(c *gin.Context) (int, error)
}

type jenkinsService struct {
	jenkinsClient meta.JenkinsClient
}

func NewJenkinsService(jenkinsClient meta.JenkinsClient) JenkinsService {
	return &jenkinsService{
		jenkinsClient: jenkinsClient,
	}
}

// GetQueuePendingTaskList 获取队列中等待的任务列表
func (svc *jenkinsService) GetQueuePendingTaskList(c *gin.Context) (taskIdList []int64, err error) {
	taskIdList = make([]int64, 0, 100)
	queue, err := svc.jenkinsClient.GetQueue(c)
	if err != nil {
		return taskIdList, err
	}
	for _, v := range queue.Raw.Items {
		taskIdList = append(taskIdList, v.ID)
	}
	return taskIdList, nil
}

// CancelQueueTaskById 根据任务ID停止队列中的等待任务
func (svc *jenkinsService) CancelQueueTaskById(c *gin.Context, taskId int64) (bool, error) {
	queue, err := svc.jenkinsClient.GetQueue(c)
	if err != nil {
		return false, err
	}
	return queue.CancelTask(c, taskId)
}

func (svc *jenkinsService) GetJobStatus(c *gin.Context) (int, error) {
	return 0, nil
}
