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

	//apiGroup.GET("/city", ApiController.CityInstance().GetList) //城市列表

	// 为区域路由创建服务实例并注册路由
	areaService := service.NewAreaService()                        // 创建区域服务实例
	areaController := ApiController.NewAreaController(areaService) // 创建控制器实例

	// 注册区域路由
	apiGroup.GET("/areas/:id", areaController.GetArea)            // 获取单个区域信息
	apiGroup.GET("/subAreas", areaController.GetSubAreas)         // 根据父ID获取子区域列表
	apiGroup.GET("/areasByLevel", areaController.GetAreasByLevel) // 根据等级获取区域列表

}
