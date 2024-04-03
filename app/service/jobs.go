package service

import (
	"moony-task-go/common/dao"
	"moony-task-go/common/model"
)

// IJobsService 定义了兼职信息服务接口
type IJobsService interface {
	GetJob(jobId int64) (*model.Jobs, error)
	GetAllJobs(comReq model.ComReq) ([]*model.Jobs, error)
	GetRecentJobs(comReq model.ComReq) ([]*model.Jobs, error)
	GetJobsNearby(lat, lng float64, radius int, comReq model.ComReq) ([]*model.Jobs, error)
	CreateJob(job *model.Jobs) (*model.Jobs, error)
	UpdateJob(job *model.Jobs) (*model.Jobs, error)
	DeleteJob(jobId int64) error
}

type JobsService struct {
	jobsDAO dao.IJobsDAO
}

// NewJobsService 创建新的兼职信息服务实例
func NewJobsService(jobsDAO dao.IJobsDAO) IJobsService {
	return &JobsService{
		jobsDAO: jobsDAO,
	}
}

func (s *JobsService) GetJob(jobId int64) (*model.Jobs, error) {
	return s.jobsDAO.Get(jobId)
}

func (s *JobsService) GetAllJobs(comReq model.ComReq) ([]*model.Jobs, error) {
	return s.jobsDAO.GetAll(comReq)
}

func (s *JobsService) GetRecentJobs(comReq model.ComReq) ([]*model.Jobs, error) {
	return s.jobsDAO.GetRecentJobs(comReq)
}

func (s *JobsService) GetJobsNearby(lat, lng float64, radius int, comReq model.ComReq) ([]*model.Jobs, error) {
	return s.jobsDAO.GetJobsNearby(lat, lng, radius, comReq)
}

func (s *JobsService) CreateJob(job *model.Jobs) (*model.Jobs, error) {
	return s.jobsDAO.Create(job)
}

func (s *JobsService) UpdateJob(job *model.Jobs) (*model.Jobs, error) {
	return s.jobsDAO.Update(job)
}

func (s *JobsService) DeleteJob(jobId int64) error {
	return s.jobsDAO.Delete(jobId)
}
