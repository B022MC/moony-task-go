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
	areaService := service.NewAreaService()                        // 创建区域服务实例
	areaController := ApiController.NewAreaController(areaService) // 创建控制器实例

	// 注册区域路由
	apiGroup.GET("/areas/:id", areaController.GetArea)            // 获取单个区域信息
	apiGroup.GET("/subAreas", areaController.GetSubAreas)         // 根据父ID获取子区域列表
	apiGroup.GET("/areasByLevel", areaController.GetAreasByLevel) // 根据等级获取区域列表

	// Assuming WorkExperienceService and WorkExperienceController have been initialized here
	workExperienceService := service.NewWorkExperienceService()                                  // Initialize your service
	workExperienceController := ApiController.NewWorkExperienceController(workExperienceService) // Initialize your controller
	apiGroup.POST("/workExperiences", workExperienceController.CreateWorkExperience)
	apiGroup.GET("/workExperiences/:id", workExperienceController.GetWorkExperience)
	apiGroup.PUT("/workExperiences/:id", workExperienceController.UpdateWorkExperience)
	apiGroup.DELETE("/workExperiences/:id", workExperienceController.DeleteWorkExperience)
	apiGroup.GET("/users/:userId/workExperiences", workExperienceController.GetAllWorkExperiencesByUserId)

	// 注册自我评估(SelfEvaluation)路由
	selfEvaluationService := service.NewSelfEvaluationService()                                  // 已创建服务实例
	selfEvaluationController := ApiController.NewSelfEvaluationController(selfEvaluationService) // 已创建控制器实例
	apiGroup.POST("/selfEvaluations", selfEvaluationController.CreateSelfEvaluation)
	apiGroup.GET("/selfEvaluations/:id", selfEvaluationController.GetSelfEvaluation)
	apiGroup.PUT("/selfEvaluations/:id", selfEvaluationController.UpdateSelfEvaluation)
	apiGroup.DELETE("/selfEvaluations/:id", selfEvaluationController.DeleteSelfEvaluation)
	apiGroup.GET("/users/:userId/selfEvaluations", selfEvaluationController.GetAllSelfEvaluationsByUserId)

	// 注册技能证书(SkillCertificate)路由
	skillCertificateService := service.NewSkillCertificateService()                                    // 已创建服务实例
	skillCertificateController := ApiController.NewSkillCertificateController(skillCertificateService) // 已创建控制器实例
	apiGroup.POST("/skillCertificates", skillCertificateController.CreateSkillCertificate)
	apiGroup.GET("/skillCertificates/:id", skillCertificateController.GetSkillCertificate)
	apiGroup.PUT("/skillCertificates/:id", skillCertificateController.UpdateSkillCertificate)
	apiGroup.DELETE("/skillCertificates/:id", skillCertificateController.DeleteSkillCertificate)
	apiGroup.GET("/users/:userId/skillCertificates", skillCertificateController.GetAllSkillCertificatesByUserId)

	// 注册 JobExpectation 相关的路由
	jobExpectationService := service.NewJobExpectationService()                                  // 假设已经有了这个函数
	jobExpectationController := ApiController.NewJobExpectationController(jobExpectationService) // 假设已经有了这个函数
	apiGroup.GET("/jobExpectations/:id", jobExpectationController.GetJobExpectation)
	apiGroup.GET("/users/:userId/jobExpectations", jobExpectationController.GetAllJobExpectationsByUserId)
	apiGroup.POST("/jobExpectations", jobExpectationController.CreateJobExpectation)
	apiGroup.PUT("/jobExpectations/:id", jobExpectationController.UpdateJobExpectation)
	apiGroup.DELETE("/jobExpectations/:id", jobExpectationController.DeleteJobExpectation)

	// 注册 Education 相关的路由
	educationService := service.NewEducationService()                             // 假设已经有了这个函数
	educationController := ApiController.NewEducationController(educationService) // 假设已经有了这个函数
	apiGroup.GET("/educations/:id", educationController.GetEducation)
	apiGroup.GET("/users/:userId/educations", educationController.GetAllEducationsByUserId)
	apiGroup.POST("/educations", educationController.CreateEducation)
	apiGroup.PUT("/educations/:id", educationController.UpdateEducation)
	apiGroup.DELETE("/educations/:id", educationController.DeleteEducation)
}
