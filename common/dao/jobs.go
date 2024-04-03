package dao

import (
	"gorm.io/gorm"
	"moony-task-go/common/model"
	"moony-task-go/utils" // 确保这里的路径与你的项目结构相匹配
)

type IJobsDAO interface {
	Get(jobId int64) (*model.Jobs, error)
	GetAll(ComReq model.ComReq) ([]*model.Jobs, error)
	GetRecentJobs(ComReq model.ComReq) ([]*model.Jobs, error)
	GetJobsNearby(lat, lng float64, radius int, ComReq model.ComReq) ([]*model.Jobs, error)
	Create(job *model.Jobs) (*model.Jobs, error)
	Update(job *model.Jobs) (*model.Jobs, error)
	Delete(jobId int64) error
}

type JobsDAO struct {
	DB *gorm.DB
}

func NewJobsDAO(db *gorm.DB) IJobsDAO {
	return &JobsDAO{DB: db}
}

func (dao *JobsDAO) Get(jobId int64) (*model.Jobs, error) {
	var job model.Jobs
	result := dao.DB.Where("job_id = ?", jobId).First(&job)
	return &job, result.Error
}

func (dao *JobsDAO) GetAll(ComReq model.ComReq) ([]*model.Jobs, error) {
	var jobs []*model.Jobs
	result := dao.DB.Offset((ComReq.Page - 1) * ComReq.Size).Limit(ComReq.Size).Find(&jobs)
	return jobs, result.Error
}

func (dao *JobsDAO) GetRecentJobs(ComReq model.ComReq) ([]*model.Jobs, error) {
	var jobs []*model.Jobs
	result := dao.DB.Order("create_time DESC").Offset((ComReq.Page - 1) * ComReq.Size).Limit(ComReq.Size).Find(&jobs)
	return jobs, result.Error
}

func (dao *JobsDAO) GetJobsNearby(lat, lng float64, radius int, ComReq model.ComReq) ([]*model.Jobs, error) {
	// 这里仅作示例，实际应用中应考虑更优化的地理空间查询策略
	var jobs []*model.Jobs
	result := dao.DB.Find(&jobs)
	if result.Error != nil {
		return nil, result.Error
	}

	var nearbyJobs []*model.Jobs
	for _, job := range jobs {
		if distance := utils.CalculateDistance(lat, lng, job.Lat, job.Lng); distance <= float64(radius) {
			nearbyJobs = append(nearbyJobs, job)
		}
	}

	// 假设ComReq.Size足够大，这里不进一步实现分页逻辑
	return nearbyJobs, nil
}

func (dao *JobsDAO) Create(job *model.Jobs) (*model.Jobs, error) {
	result := dao.DB.Create(job)
	return job, result.Error
}

func (dao *JobsDAO) Update(job *model.Jobs) (*model.Jobs, error) {
	result := dao.DB.Save(job)
	return job, result.Error
}

func (dao *JobsDAO) Delete(jobId int64) error {
	result := dao.DB.Delete(&model.Jobs{}, jobId)
	return result.Error
}
