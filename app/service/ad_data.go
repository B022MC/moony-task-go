package service

var (
	ReqUserUrl  = "/api/client/user"
	ReqOrderUrl = "/api/client/order"
)

type AdData struct {
}

func AdDataInstance() *AdData {
	return &AdData{}
}

type AdDataRsp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// OrderData 订单数据上报
//func (a *AdData) OrderData(orderId int64, status int) error {
//
//	order, err := dao.OrderInstance().Get(orderId)
//	if err != nil {
//		log.Errorf("order.data dao.OrderInstance().Get err=[%s]", err.Error())
//		return err
//	}
//	if order == nil {
//		return errors.New("order.data order is nil")
//	}
//	var itemsInfo []model.OrderItemInfo // 存储订单中每个商品的信息
//	//goosInfo := make(map[string]interface{})
//	json.Unmarshal([]byte(order.ItemsInfo), &itemsInfo)
//	// 提取 name 并使用逗号连接
//	var names []string
//	var goodsIds []string
//	for _, item := range itemsInfo {
//		names = append(names, item.Name)
//		goodsIds = append(goodsIds, item.GoodsId)
//	}
//	user, err := dao.UserInstance().Get(order.UserId)
//	if err != nil {
//		log.Errorf("order.data dao.UserInstance().Get err=[%s]", err.Error())
//		return err
//	}
//
//	postData := make(map[string]interface{})
//	postData["orderId"] = cast.ToString(orderId)
//	postData["orderNo"] = cast.ToString(order.OutTradeNo)
//	postData["goodsId"] = cast.ToString(goodsIds)
//	postData["goodsName"] = cast.ToString(names)
//	postData["userId"] = cast.ToString(order.UserId)
//	postData["totalFee"] = cast.ToString(order.TotalFee)
//	postData["status"] = cast.ToString(status)
//	postData["payType"] = cast.ToString(order.PayType)
//	postData["isReport"] = cast.ToString(order.IsReport)
//	postData["payTime"] = cast.ToString(order.PayTime)
//
//	if user == nil {
//		postData["adId"] = ""
//		postData["appId"] = cast.ToString(order.AppId)
//		postData["hashId"] = cast.ToString(utils.Hash64(fmt.Sprintf("%d", time.Now().UnixNano())))
//		postData["materialId"] = ""
//		postData["adSource"] = "nature"
//	} else {
//		postData["adId"] = cast.ToString(user.AdId)
//		postData["appId"] = cast.ToString(user.AppId)
//		postData["hashId"] = cast.ToString(user.HashId)
//		postData["materialId"] = cast.ToString(user.MaterialId)
//		if user.AdSource == "" {
//			postData["adSource"] = "nature"
//		} else {
//			postData["adSource"] = cast.ToString(user.AdSource)
//		}
//	}
//
//	var rsp AdDataRsp
//	cfg := config.GetConfig().Ad
//	postJson, err := utils.HttpPostJson(cfg.DataUrl+ReqOrderUrl, utils.EncodeJSON(postData), "")
//	if err != nil {
//		log.Errorf("order.data utils.HttpPostJson err=[%s]", err.Error())
//		return err
//	}
//	json.Unmarshal(postJson, &rsp)
//	if rsp.Code != 0 {
//		log.Errorf("order.data response code=[%d] err=[%s]", rsp.Code, err.Error())
//		return errors.New(fmt.Sprintf("order.data response code=[%d] err=[%s]", rsp.Code, err.Error()))
//	}
//
//	return nil
//}

// OrderListRepost 指定日期订单数据上报
//func (a *AdData) OrderListRepost(startDate, endDate string) {
//	startTime := utils.DateToTime(startDate)
//	endTime := utils.DateToTime(endDate) + 86400
//
//	list, err := dao.OrderInstance().GetOrderList(startTime, endTime)
//	if err != nil {
//		log.Errorf("order.data.OrderListRepost dao.OrderInstance().GetOrderList err=[%s]", err.Error())
//		return
//	}
//	for _, order := range list {
//		var status int
//		if order.Status == model.OrderStatusPayed {
//			status = model.ReportOrderStatusOne
//		} else if order.Status == model.OrderStatusRefund {
//			status = model.ReportOrderStatusTwo
//		}
//		orderId := order.Id
//		if err := a.OrderData(orderId, status); err != nil {
//			log.Errorf("order.data.OrderListRepost a.OrderData err=[%s]", err.Error())
//		}
//	}
//
//	return
//}
