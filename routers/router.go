package routers

import (
	"github.com/gin-gonic/gin"
	ApiController "moony-task-go/app/controller"
	"moony-task-go/app/service"
	"moony-task-go/core/middleware"
	"net/http"
)

// InitRouter 初始化路由
func InitRouter(engine *gin.Engine) {
	engine.StaticFS("/h5", http.Dir("./config/html"))
	engine.GET("/health", ApiController.BaseInstance().Health)
	initApiRouter(engine)
}

func initApiRouter(engine *gin.Engine) {
	apiGroup := engine.Group("/api")
	apiGroup.Use(middleware.ApiInstance().Recovery)
	apiGroup.Use(middleware.ApiInstance().ApiBefore)
	apiGroup.GET("/launch", ApiController.BaseInstance().Launch) //配置获取

	apiGroup.POST("/user/login", ApiController.UserInstance().Login) //用户登录
	//apiGroup.POST("/user/logout", ApiController.UserInstance().Logout) //退出登录
	apiGroup.GET("/user", ApiController.UserInstance().GetUserInfo)  //用户详情
	apiGroup.PUT("/user", ApiController.UserInstance().SaveUserInfo) //编辑信息
	//apiGroup.Any("/user/cancel", ApiController.UserInstance().Cancel)  //微信注销用户

	// 为区域路由创建服务实例并注册路由
	areaService := service.NewAreaService()                        //
	areaController := ApiController.NewAreaController(areaService) //
	apiGroup.GET("/areas/:id", areaController.GetArea)             // 获取单个区域信息
	apiGroup.GET("/subAreas", areaController.GetSubAreas)          // 根据父ID获取子区域列表
	apiGroup.GET("/areasByLevel", areaController.GetAreasByLevel)  // 根据等级获取区域列表

	// 初始化 UserResume 相关服务和控制器
	userResumeService := service.NewUserResumeService()                              //
	userResumeController := ApiController.NewUserResumeController(userResumeService) //
	apiGroup.GET("/user/resume", userResumeController.GetUserResumeByUserId)

	// Assuming WorkExperienceService and WorkExperienceController have been initialized here
	workExperienceService := service.NewWorkExperienceService()                                  //
	workExperienceController := ApiController.NewWorkExperienceController(workExperienceService) //
	apiGroup.POST("/workExperiences", workExperienceController.CreateWorkExperience)
	apiGroup.GET("/workExperiences", workExperienceController.GetWorkExperience)
	apiGroup.PUT("/workExperiences", workExperienceController.UpdateWorkExperience)
	apiGroup.DELETE("/workExperiences", workExperienceController.DeleteWorkExperience)
	apiGroup.GET("/users/workExperiences", workExperienceController.GetAllWorkExperiencesByUserId)

	// 注册自我评估(SelfEvaluation)路由
	selfEvaluationService := service.NewSelfEvaluationService()                                  //
	selfEvaluationController := ApiController.NewSelfEvaluationController(selfEvaluationService) //
	apiGroup.POST("/selfEvaluations", selfEvaluationController.CreateSelfEvaluation)
	apiGroup.GET("/selfEvaluations", selfEvaluationController.GetSelfEvaluation)
	apiGroup.PUT("/selfEvaluations", selfEvaluationController.UpdateSelfEvaluation)
	apiGroup.DELETE("/selfEvaluations", selfEvaluationController.DeleteSelfEvaluation)
	apiGroup.GET("/users/selfEvaluations", selfEvaluationController.GetAllSelfEvaluationsByUserId)

	// 注册技能证书(SkillCertificate)路由
	skillCertificateService := service.NewSkillCertificateService()                                    // 已创建服务实例
	skillCertificateController := ApiController.NewSkillCertificateController(skillCertificateService) // 已创建控制器实例
	apiGroup.POST("/skillCertificates", skillCertificateController.CreateSkillCertificate)
	apiGroup.GET("/skillCertificates", skillCertificateController.GetSkillCertificate)
	apiGroup.PUT("/skillCertificates", skillCertificateController.UpdateSkillCertificate)
	apiGroup.DELETE("/skillCertificates", skillCertificateController.DeleteSkillCertificate)
	apiGroup.GET("/users/skillCertificates", skillCertificateController.GetAllSkillCertificatesByUserId)

	// 注册 JobExpectation 相关的路由
	jobExpectationService := service.NewJobExpectationService()                                  //
	jobExpectationController := ApiController.NewJobExpectationController(jobExpectationService) //
	apiGroup.GET("/jobExpectations", jobExpectationController.GetJobExpectation)
	apiGroup.GET("/users/jobExpectations", jobExpectationController.GetAllJobExpectationsByUserId)
	apiGroup.POST("/jobExpectations", jobExpectationController.CreateJobExpectation)
	apiGroup.PUT("/jobExpectations", jobExpectationController.UpdateJobExpectation)
	apiGroup.DELETE("/jobExpectations", jobExpectationController.DeleteJobExpectation)

	// 注册 Education 相关的路由
	educationService := service.NewEducationService()                             //
	educationController := ApiController.NewEducationController(educationService) //
	apiGroup.GET("/educations", educationController.GetEducation)
	apiGroup.GET("/users/educations", educationController.GetAllEducationsByUserId)
	apiGroup.POST("/educations", educationController.CreateEducation)
	apiGroup.PUT("/educations", educationController.UpdateEducation)
	apiGroup.DELETE("/educations", educationController.DeleteEducation)
}
