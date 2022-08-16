package service

import (
	"github.com/bndr/gojenkins"
	"github.com/gin-gonic/gin"
	"github.com/go_example/internal/meta"
	"io/ioutil"
)

type JenkinsService interface {
	GetQueuePendingTaskList(c *gin.Context) (taskIdList []int64, err error)
	CancelQueueTaskById(c *gin.Context, taskId int64) (bool, error)
	CancelQueueAllTask(c *gin.Context) error
	GetJobLatestStatusAndNextBuildNumByName(c *gin.Context, jobName string) (int, int, error)
	GetAllJobLatestStatus(c *gin.Context) (map[string]int, error)
	GetJobLastBuildLog(c *gin.Context, jobName string) (string, error)
	CreateJobFolder(c *gin.Context, folderName string, parents ...string) error
	CreateJobInFolder(c *gin.Context, configPath string, jobName string, parentIDs ...string) (*gojenkins.Job, error)
	RenameJob(c *gin.Context, oldName string, newName string) *gojenkins.Job
	BuildJob(c *gin.Context, jobName string, params map[string]string) (int64, error)
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
	for _, task := range queue.Raw.Items {
		taskIdList = append(taskIdList, task.ID)
	}
	return taskIdList, nil
}

// CancelQueueAllTask 停止队列中所有等待任务
func (svc *jenkinsService) CancelQueueAllTask(c *gin.Context) error {
	queue, err := svc.jenkinsClient.GetQueue(c)
	if err != nil {
		return err
	}
	for _, task := range queue.Raw.Items {
		if _, err = queue.CancelTask(c, task.ID); err != nil {
			return err
		}
	}
	return nil
}

// CancelQueueTaskById 根据任务ID停止队列中的等待任务
func (svc *jenkinsService) CancelQueueTaskById(c *gin.Context, taskId int64) (bool, error) {
	queue, err := svc.jenkinsClient.GetQueue(c)
	if err != nil {
		return false, err
	}
	// todo 这里返回结果有问题 - 已反馈给作者
	return queue.CancelTask(c, taskId)
}

// GetAllJobLatestStatus 获取所有job最近编译状态
func (svc *jenkinsService) GetAllJobLatestStatus(c *gin.Context) (map[string]int, error) {
	resMap := make(map[string]int, 0)
	jobList, err := svc.jenkinsClient.GetAllJobs(c)
	if err != nil {
		return nil, err
	}
	for _, job := range jobList {
		resMap[job.GetName()] = getJobStatus(job)
	}
	return resMap, nil
}

// GetJobLatestStatusByName 根据job的名称获取某个job最近一次的编译状态
func (svc *jenkinsService) GetJobLatestStatusAndNextBuildNumByName(c *gin.Context, jobName string) (int, int, error) {
	job, err := svc.jenkinsClient.GetJob(c, jobName)
	if err != nil {
		return 0, 0, err
	}
	return int(job.Raw.NextBuildNumber), getJobStatus(job), nil
}

// getJobStatus 根据job编译详情判断状态
// 1 成功 2 失败 0 未编译或其他
func getJobStatus(job *gojenkins.Job) int {
	if job.Raw.LastBuild.Number == 0 { // 未编译
		return 0
	}
	if job.Raw.LastBuild.Number == job.Raw.LastSuccessfulBuild.Number { // 编译成功
		return 1
	}
	if job.Raw.LastBuild.Number == job.Raw.LastFailedBuild.Number { // 编译失败
		return 2
	}
	return 0
}

// GetJobLastBuildLog 获取指定job最近一次编译日志
func (svc *jenkinsService) GetJobLastBuildLog(c *gin.Context, jobName string) (string, error) {
	job, err := svc.jenkinsClient.GetJob(c, jobName)
	if err != nil {
		return "", err
	}
	build, err := job.GetLastBuild(c)
	if err != nil {
		return "", err
	}
	// 实时日志
	resp, err := build.GetConsoleOutputFromIndex(c, 0)
	if err != nil {
		return "", err
	}
	return resp.Content, nil
}

// CreateJobFolder 创建job的文件夹
// 前提条件：安装Folders Plugin
func (svc *jenkinsService) CreateJobFolder(c *gin.Context, folderName string, parents ...string) error {
	_, err := svc.jenkinsClient.CreateFolder(c, folderName, parents...)
	if err != nil {
		return err
	}
	return nil
}

// CreateJobInFolder 创建一个job工程项目
func (svc *jenkinsService) CreateJobInFolder(c *gin.Context, configPath string, jobName string, parentIDs ...string) (*gojenkins.Job, error) {
	buf, err := ioutil.ReadFile(configPath)
	if err != nil {
		panic(err)
	}
	config := string(buf)
	return svc.jenkinsClient.CreateJobInFolder(c, config, jobName, parentIDs...)
}

// RenameJob 重命名job工程项目
// oldName为job的完整路径（/job后面的完整路径），例如：xudttest/job/folder1_test/job/job_test
// newName为新的job名称：例如job_test1
func (svc *jenkinsService) RenameJob(c *gin.Context, oldName string, newName string) *gojenkins.Job {
	return svc.jenkinsClient.RenameJob(c, oldName, newName)
}

// BuildJob 编译job，返回队列id
func (svc *jenkinsService) BuildJob(c *gin.Context, jobName string, params map[string]string) (int64, error) {
	return svc.jenkinsClient.BuildJob(c, jobName, params)
}
