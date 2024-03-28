package controller

type AdData struct {
}

func AdDataInstance() *AdData {
	return &AdData{}
}

// ReportUserData 指定日期用户数据上报
//func (a *AdData) ReportUserData(ctx *gin.Context) {
//	endDate := ctx.Query("endDate")
//	startDate := ctx.Query("startDate")
//
//	service.AdDataInstance().UserListRepost(startDate, endDate)
//
//	ctx.JSON(http.StatusOK, common.NewRspOk())
//}

// ReportOrderDate 指定日期订单数据上报
//func (a *AdData) ReportOrderDate(ctx *gin.Context) {
//	endDate := ctx.Query("endDate")
//	startDate := ctx.Query("startDate")
//
//	service.AdDataInstance().OrderListRepost(startDate, endDate)
//
//	ctx.JSON(http.StatusOK, common.NewRspOk())
//}
