package dao

import (
	"fmt"
	"github.com/spf13/cast"
	"gorm.io/gorm"
	"moony-task-go/common/model"
	"moony-task-go/core/global"
	"moony-task-go/utils" // 确保这里的路径与你的项目结构相匹配
	"strconv"
	"strings"
)

type IJobsDAO interface {
	Get(jobId int64) (*model.Jobs, error)
	GetAll(ComReq model.ComReq) ([]*model.Jobs, error)
	GetRecentJobs(ComReq model.ComReq) ([]*model.Jobs, error)
	GetJobsNearby(lat, lng float64, radius int, ComReq model.ComReq) ([]*model.Jobs, error)
	Create(job *model.Jobs) (*model.Jobs, error)
	Update(job *model.Jobs) (*model.Jobs, error)
	Delete(jobId int64) error
	FilterJobs(req model.JobFilterRequest, comReq model.ComReq) ([]*model.JobRsp, error)
}

type JobsDAO struct {
	DB *gorm.DB
}

func NewJobsDAO(db *gorm.DB) IJobsDAO {
	return &JobsDAO{DB: db}
}

func (dao *JobsDAO) TableName() string {
	return "jobs"
}

func (dao *JobsDAO) Get(jobId int64) (*model.Jobs, error) {
	var job model.Jobs
	result := dao.DB.Table(dao.TableName()).Where("job_id = ?", jobId).First(&job)
	return &job, result.Error
}

func (dao *JobsDAO) GetAll(ComReq model.ComReq) ([]*model.Jobs, error) {
	var jobs []*model.Jobs
	result := dao.DB.Table(dao.TableName()).Offset((ComReq.Page - 1) * ComReq.Size).Limit(ComReq.Size).Find(&jobs)
	return jobs, result.Error
}

func (dao *JobsDAO) GetRecentJobs(ComReq model.ComReq) ([]*model.Jobs, error) {
	var jobs []*model.Jobs
	result := dao.DB.Table(dao.TableName()).Order("create_time DESC").Offset((ComReq.Page - 1) * ComReq.Size).Limit(ComReq.Size).Find(&jobs)
	return jobs, result.Error
}

func (dao *JobsDAO) GetJobsNearby(lat, lng float64, radius int, ComReq model.ComReq) ([]*model.Jobs, error) {
	// 这里仅作示例，实际应用中应考虑更优化的地理空间查询策略
	var jobs []*model.Jobs
	result := dao.DB.Table(dao.TableName()).Find(&jobs)
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
	result := dao.DB.Table(dao.TableName()).Create(job)
	return job, result.Error
}

func (dao *JobsDAO) Update(job *model.Jobs) (*model.Jobs, error) {
	result := dao.DB.Table(dao.TableName()).Save(job)
	return job, result.Error
}

func (dao *JobsDAO) Delete(jobId int64) error {
	result := dao.DB.Table(dao.TableName()).Delete(&model.Jobs{}, jobId)
	return result.Error
}

func (dao *JobsDAO) FilterJobs(req model.JobFilterRequest, comReq model.ComReq) ([]*model.JobRsp, error) {
	// Start with a basic query for JobRsp model
	var jobs []model.Jobs
	query := dao.DB.Table(dao.TableName()).Model(&jobs)

	// Apply filters from JobFilterRequest
	if req.CategoryIds != "" {
		// 将逗号分隔的字符串转换为ID数组
		typeIds := strings.Split(req.CategoryIds, ",")
		var categoryIds []int

		for _, idStr := range typeIds {
			id, err := strconv.Atoi(idStr)
			if err != nil {
				continue // 如果转换失败，跳过这个ID
			}
			categoryIds = append(categoryIds, id)
		}

		if len(categoryIds) > 0 {
			// 使用IN子句来查找匹配的类别ID
			query = query.Joins("JOIN job_category_relations ON job_category_relations.job_id = jobs.id").
				Where("job_category_relations.category_id IN ?", categoryIds)
		}
	}

	// 筛选地区
	if req.Area.CityId != "" && req.Area.Level != 0 {
		// 将逗号分隔的字符串转换为地区名称数组
		var areasId []int
		if req.Area.Level == 2 {
			areasId, _ = NewAreaDAO(global.Db).GetAreaIDs(cast.ToInt(req.Area.CityId))
		} else {
			areasIdStr := strings.Split(req.Area.CityId, ",")
			for _, idStr := range areasIdStr {
				id, err := strconv.Atoi(idStr)
				if err != nil {
					fmt.Println("Error converting string to int:", err)
					continue
				}
				areasId = append(areasId, id)
			}
		}
		query = query.Where("city_id IN ?", areasId)
	}

	// 应用筛选条件
	for _, filter := range req.Filters {
		if isValidFilter(filter.Key) {
			query = query.Where(fmt.Sprintf("%s IN (?)", filter.Key), filter.Values)
		}
	}
	// 使用直接的 SQL 地理位置距离计算进行筛选
	if req.Sort.Lat != 0 && req.Sort.Lng != 0 {
		const maxDistance = 5000 // 5公里，单位为米
		// 直接在 SQL 查询中加入地理距离计算
		query = query.Where("ST_Distance_Sphere(point(lng, lat), point(?, ?)) <= ?", req.Sort.Lng, req.Sort.Lat, maxDistance)
	}

	// Sorting based on the provided sort criteria
	if req.Sort.Desc != "" {
		var sortField string
		switch req.Sort.Desc {
		case "date_desc":
			sortField = "create_time DESC"
		case "date_asc":
			sortField = "create_time ASC"
		default:
			sortField = "create_time DESC" // default sorting
		}
		query = query.Order(sortField)
	}

	// Pagination handling
	if comReq.Size > 0 {
		offset := (comReq.Page - 1) * comReq.Size
		query = query.Offset(offset).Limit(comReq.Size)
	}

	// 执行查询
	result := query.Find(&jobs)
	if result.Error != nil {
		return nil, result.Error
	}

	// 将Jobs数据转换为JobRsp数据
	var jobsRsp []*model.JobRsp
	for _, job := range jobs {
		shortname, _ := NewAreaDAO(global.Db).GetAreaNameById(job.CityId)
		jobRsp := &model.JobRsp{
			Id:          job.JobId,
			Title:       job.Title,
			Description: job.Description,
			Salary:      job.Salary,
			//Requirements: []string{}, // 这里可能需要处理Requirements字段的分解或转换
			//Company:      job.UserId, // 如果CompanyId与Company名称需要转换，需要额外逻辑处理
			Location: shortname,
		}
		jobsRsp = append(jobsRsp, jobRsp)
	}
	return jobsRsp, nil
}

// isValidFilter 验证是否是有效的筛选字段
func isValidFilter(key string) bool {
	switch key {
	case "gender_requirement", "work_period":
		return true
	default:
		return false
	}
}
