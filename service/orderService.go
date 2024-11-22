package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	iUtils "github.com/iWyh2/go-myUtils/utils"
	"github.com/ulule/deepcopier"
	"go-SkyTakeaway/common/constant"
	errs "go-SkyTakeaway/common/errors"
	"go-SkyTakeaway/global"
	"go-SkyTakeaway/model"
	"go-SkyTakeaway/model/dto"
	"go-SkyTakeaway/model/entity"
	pageResult "go-SkyTakeaway/model/result/page"
	"go-SkyTakeaway/model/vo"
	"go-SkyTakeaway/router/websocket"
	"gorm.io/gorm"
	"log"
	"strconv"
	"time"
)

// OrderSubmit 用户下单
func OrderSubmit(data *dto.OrderSubmitDTO, ctx *gin.Context) *vo.OrderSubmitVO {
	// 异常情况的处理（收货地址为空、超出配送范围、购物车为空）
	addressBook := QueryAddressBookById(strconv.Itoa(data.AddressBookId))
	if addressBook == nil {
		panic(errs.AddressBookIsNilError)
	}
	// 查询当前用户的购物车数据
	userId, ok := ctx.Get(constant.UserID)
	if !ok {
		panic(errs.MissUserIdError)
	}
	id, _ := strconv.Atoi(userId.(string))
	shoppingCartList := queryShoppingCart(entity.ShoppingCart{UserId: id})
	if shoppingCartList == nil || len(shoppingCartList) == 0 {
		panic(errs.ShoppingCartIsNilError)
	}
	// 构造订单数据
	var order entity.Order
	err := deepcopier.Copy(data).To(&order)
	if err != nil {
		panic(errs.ServerInternalError)
	}
	order.UserId = id
	order.Number = strconv.FormatInt(time.Now().UnixMilli(), 10)
	order.Phone = addressBook.Phone
	order.Address = addressBook.Detail
	order.Consignee = addressBook.Consignee
	order.Status = constant.PendingPayment
	order.PayStatus = constant.UnPaid
	// 向订单表插入1条数据
	if err := global.Db.Create(&order).Error; err != nil {
		panic(errs.DBError)
	}
	// 订单明细数据
	orderDetailList := make([]entity.OrderDetail, 0)
	for _, shoppingCart := range shoppingCartList {
		var orderDetail entity.OrderDetail
		if err := deepcopier.Copy(&shoppingCart).To(&orderDetail); err != nil {
			panic(errs.ServerInternalError)
		}
		orderDetail.OrderId = order.Id
		orderDetailList = append(orderDetailList, orderDetail)
	}
	// 向明细表插入n条数据
	insertBatchOrderDetail(orderDetailList)
	// 清理购物车中的数据
	CleanShoppingCart(ctx)
	// 封装返回结果
	return &vo.OrderSubmitVO{
		OrderId:     order.Id,
		OrderNumber: order.Number,
		OrderAmount: order.Amount,
		OrderTime:   order.OrderTime,
	}
}

// OrderConfirm 接单
func OrderConfirm(data *dto.OrderConfirmDTO) {
	var id int
	switch data.OrderId.(type) {
	case int:
		id = data.OrderId.(int)
	case string:
		id, _ = strconv.Atoi(data.OrderId.(string))
	case float64:
		id = int(data.OrderId.(float64))
	}
	// 更新订单为已接单
	UpdateOrder(entity.Order{
		Id:     id,
		Status: constant.Confirmed,
	})
}

// OrderRejection 拒单
func OrderRejection(data *dto.OrderRejectionDTO) {
	// 根据id查询订单
	order := getOrderById(strconv.Itoa(data.OrderId))
	// 订单只有存在且状态为2（待接单）才可以拒单
	if order == nil || order.Status != constant.ToBeConfirmed {
		panic(errs.OrderStatusError)
	}
	// 检查支付状态
	if order.PayStatus == constant.Paid {
		// 用户已支付，需要退款
		// 模拟微信退款
		log.Printf("已支付订单被拒单, 退款: [%v￥]", order.Amount)
	}
	// 根据订单id更新订单状态、拒单原因、取消时间
	order.Status = constant.Cancelled
	order.RejectionReason = data.RejectionReason
	order.CancelTime = model.LocalTime(time.Now())
	UpdateOrder(*order)
}

// OrderPayment 订单支付
func OrderPayment(orderData *dto.OrderPaymentDTO, ctx *gin.Context) *vo.OrderPaymentVO {
	// 调用微信支付接口，生成预支付交易单，此处进行模拟
	log.Printf("调用微信支付接口: [%v]", orderData)
	// 模拟支付成功，修改订单状态
	go paySuccess(orderData.OrderNumber, ctx)
	return &vo.OrderPaymentVO{
		NonceStr:   iUtils.UUID(),
		PaySign:    "iWyh2",
		SignType:   "iWyh2",
		PackageStr: iUtils.UUID(),
		TimeStamp:  strconv.FormatInt(time.Now().UnixMilli(), 10),
	}
}

// OrderDetail 根据订单id查询订单详情
func OrderDetail(orderId string) *vo.OrderVO {
	// 根据id查询订单
	order := getOrderById(orderId)
	// 查询该订单对应的菜品/套餐明细
	orderDetailList := GetOrderDetailByOrderId(orderId)
	// 将该订单及其详情封装到OrderVO并返回
	var orderVO vo.OrderVO
	err := deepcopier.Copy(order).To(&orderVO)
	if err != nil {
		panic(errs.ServerInternalError)
	}
	orderVO.OrderDetailList = orderDetailList
	return &orderVO
}

// OrderConditionSearch 订单搜索
func OrderConditionSearch(data *dto.OrderPageQueryDTO) *pageResult.Page[vo.OrderVO] {
	// 获取分页查询参数
	pageIndex := data.Page
	pageSize := data.PageSize
	// 准备数据容器
	orders := make([]entity.Order, 0)
	orderVOs := make([]vo.OrderVO, 0)
	// 准备返回数据
	page := &pageResult.Page[vo.OrderVO]{}
	// 指定查询模型，方便后续操作
	query := global.Db.Table("orders")
	// 设置模糊查询条件
	if data.Status != 0 {
		query = query.Where("status = ?", data.Status)
	}
	if data.Number != "" {
		query = query.Where("number like ?", "%"+data.Number+"%")
	}
	if data.Phone != "" {
		query = query.Where("phone like ?", "%"+data.Phone+"%")
	}
	if data.UserId != 0 {
		query = query.Where("user_id = ?", data.UserId)
	}
	if data.BeginTime != "" {
		beginTime, err := time.Parse("2006-01-02 15:04:05", data.BeginTime)
		if err != nil {
			panic(errs.ServerInternalError)
		}
		query = query.Where("order_time >= ?", model.LocalTime(beginTime))
	}
	if data.EndTime != "" {
		endTime, err := time.Parse("2006-01-02 15:04:05", data.EndTime)
		if err != nil {
			panic(errs.ServerInternalError)
		}
		query = query.Where("order_time <= ?", model.LocalTime(endTime))
	}
	// 设置按 order_time desc 排序
	if err := query.Order("order_time desc").Error; err != nil {
		panic(errs.DBError)
	}
	// 统计数据数量，执行分页查询
	if err := query.Count(&page.Total).
		Limit(pageSize).Offset((pageIndex - 1) * pageSize).
		Find(&orders).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			orders = nil
		}
		panic(errs.DBError)
	}
	// 处理数据
	if orders != nil && len(orders) > 0 {
		for _, order := range orders {
			orderId := order.Id
			// 查询订单明细
			orderDetails := GetOrderDetailByOrderId(strconv.Itoa(orderId))
			var orderVO vo.OrderVO
			if err := deepcopier.Copy(&order).To(&orderVO); err != nil {
				panic(errs.ServerInternalError)
			}
			orderVO.OrderDishes = getOrderDishes(orderDetails)
			orderVOs = append(orderVOs, orderVO)
		}
	}
	page.Records = orderVOs
	return page
}

// HistoryOrders 查询历史订单
func HistoryOrders(ctx *gin.Context) *pageResult.Page[vo.OrderVO] {
	// 获取分页查询参数
	pageIndex, _ := strconv.Atoi(ctx.Query("page"))
	pageSize, _ := strconv.Atoi(ctx.Query("pageSize"))
	// 准备数据容器
	orders := make([]entity.Order, 0)
	orderVOs := make([]vo.OrderVO, 0)
	// 准备返回数据
	page := &pageResult.Page[vo.OrderVO]{}
	// 指定查询模型，方便后续操作
	query := global.Db.Table("orders")
	// 设置模糊查询条件
	if ctx.Query("status") != "" {
		query = query.Where("status = ?", ctx.Query("status"))
	}
	// 设置按 order_time desc 排序
	if err := query.Order("order_time desc").Error; err != nil {
		panic(errs.DBError)
	}
	// 统计数据数量，执行分页查询
	if err := query.Count(&page.Total).
		Limit(pageSize).Offset((pageIndex - 1) * pageSize).
		Find(&orders).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			orders = nil
		}
		panic(errs.DBError)
	}
	// 处理数据
	if orders != nil && len(orders) > 0 {
		for _, order := range orders {
			orderId := order.Id
			// 查询订单明细
			orderDetails := GetOrderDetailByOrderId(strconv.Itoa(orderId))
			var orderVO vo.OrderVO
			if err := deepcopier.Copy(&order).To(&orderVO); err != nil {
				panic(errs.ServerInternalError)
			}
			orderVO.OrderDetailList = orderDetails
			orderVOs = append(orderVOs, orderVO)
		}
	}
	page.Records = orderVOs
	return page
}

// OrderStatistics 各个状态的订单数量统计
func OrderStatistics() *vo.OrderStatisticsVO {
	// 根据状态，分别查询出待接单、待派送、派送中的订单数量
	toBeConfirmed, confirmed, deliveryInProgress := int64(0), int64(0), int64(0)
	if err := global.Db.Table("orders").
		Where("status = ?", constant.ToBeConfirmed).
		Count(&toBeConfirmed).Error; err != nil {
		panic(errs.DBError)
	}
	if err := global.Db.Table("orders").
		Where("status = ?", constant.Confirmed).
		Count(&confirmed).Error; err != nil {
		panic(errs.DBError)
	}
	if err := global.Db.Table("orders").
		Where("status = ?", constant.DeliveryInProgress).
		Count(&deliveryInProgress).Error; err != nil {
		panic(errs.DBError)
	}
	// 将查询出的数据封装到orderStatisticsVO中响应
	return &vo.OrderStatisticsVO{
		ToBeConfirmed:      int(toBeConfirmed),
		Confirmed:          int(confirmed),
		DeliveryInProgress: int(deliveryInProgress),
	}
}

// CancelOrderByBusiness 商家取消订单
func CancelOrderByBusiness(data *dto.OrderCancelDTO) {
	// 根据id查询订单
	order := getOrderById(strconv.Itoa(data.OrderId))
	// 检查支付状态
	if order == nil {
		panic(errs.OrderNotExistError)
	}
	if order.PayStatus == constant.Paid {
		// 用户已支付，需要退款
		// 模拟微信退款
		log.Printf("已支付订单被取消订单, 退款: [%v￥]", order.Amount)
	}
	// 根据订单id更新订单状态、取消原因、取消时间
	order.Status = constant.Cancelled
	order.CancelReason = data.CancelReason
	order.CancelTime = model.LocalTime(time.Now())
	UpdateOrder(*order)
}

// CancelOrder 用户取消订单
func CancelOrder(orderId string) {
	// 根据id查询订单
	order := getOrderById(orderId)
	// 校验订单是否存在
	if order == nil {
		panic(errs.OrderNotExistError)
	}
	// 校验订单状态
	if order.Status > constant.ToBeConfirmed {
		panic(errs.OrderStatusError)
	}
	// 订单处于待接单状态下取消，需要进行退款
	if order.Status == constant.ToBeConfirmed {
		// 模拟微信退款
		log.Printf("待接单订单取消, 退款: [%v￥]", order.Amount)
		// 支付状态修改为 退款
		order.PayStatus = constant.Refund
	}
	// 更新订单状态、取消原因、取消时间
	order.Status = constant.Cancelled
	order.CancelReason = "用户取消"
	order.CancelTime = model.LocalTime(time.Now())
	UpdateOrder(*order)
}

// RepetitionOrder 再来一单
func RepetitionOrder(orderId string, ctx *gin.Context) {
	// 查询当前用户id
	userId, _ := strconv.Atoi(ctx.MustGet("userId").(string))
	// 根据订单id查询当前订单详情
	orderDetailList := GetOrderDetailByOrderId(orderId)
	// 将订单详情对象转换为购物车对象
	var shoppingCartList []entity.ShoppingCart
	for _, orderDetail := range orderDetailList {
		// 将原订单详情里面的菜品信息重新复制到购物车对象中
		var shoppingCart entity.ShoppingCart
		err := deepcopier.Copy(&orderDetail).To(&shoppingCart)
		if err != nil {
			panic(errs.ServerInternalError)
		}
		shoppingCart.Id = 0
		shoppingCart.UserId = userId
		shoppingCartList = append(shoppingCartList, shoppingCart)
	}
	// 将购物车对象批量添加到数据库
	InsertBatchShoppingCart(shoppingCartList)
}

// OrderReminder 用户催单
func OrderReminder(orderId string) {
	// 查询订单是否存在
	order := getOrderById(orderId)
	if order == nil {
		panic(errs.OrderNotExistError)
	}
	// 基于WebSocket实现催单
	jsonMap := map[string]any{
		"type":    2,
		"orderId": orderId,
		"content": "订单号: " + order.Number,
	}
	websocket.WSServer.SendToAllClients(jsonMap)
}

// 支付成功，修改订单状态
func paySuccess(orderNumber string, ctx *gin.Context) {
	userId, ok := ctx.Get(constant.UserID)
	if !ok {
		panic(errs.MissUserIdError)
	}
	// 根据订单号查询当前用户的订单
	order := getOrderByNumberAndUserId(orderNumber, userId.(string))
	// 根据订单id更新订单的状态、支付方式、支付状态、结账时间
	UpdateOrder(entity.Order{
		Id:           order.Id,
		Status:       constant.ToBeConfirmed,
		PayStatus:    constant.Paid,
		CheckoutTime: model.LocalTime(time.Now()),
	})
	// 基于WebSocket提醒商家来单了
	jsonMap := map[string]any{
		"type":    1,
		"orderId": order.Id,
		"content": "订单号: " + orderNumber,
	}
	websocket.WSServer.SendToAllClients(jsonMap)
}

// 根据订单号和用户id查询订单
func getOrderByNumberAndUserId(orderNumber, userId string) *entity.Order {
	var order entity.Order
	if err := global.Db.
		Where("number = ?", orderNumber).
		Where("user_id = ?", userId).
		First(&order).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		panic(errs.DBError)
	}
	return &order
}

// 根据订单详情获取菜品信息字符串
func getOrderDishes(orderDetails []entity.OrderDetail) string {
	orderDishes := ""
	// 将每一条订单菜品信息拼接为字符串（格式：宫保鸡丁*3;）
	for _, orderDetail := range orderDetails {
		orderDishes += orderDetail.Name + "*" + strconv.Itoa(orderDetail.Number) + ";"
	}
	return orderDishes
}

// 根据订单id查询订单
func getOrderById(orderId string) *entity.Order {
	var order entity.Order
	if err := global.Db.
		Where("id = ?", orderId).
		Find(&order).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		panic(errs.DBError)
	}
	return &order
}

// UpdateOrder 修改订单信息
func UpdateOrder(order entity.Order) {
	if err := global.Db.Table("orders").
		Updates(order).Error; err != nil {
		panic(errs.DBError)
	}
}

// GetOrderByStatusAndOrderTime 根据状态和下单时间查询订单
func GetOrderByStatusAndOrderTime(status int, orderTime model.LocalTime) []entity.Order {
	var orders []entity.Order
	if err := global.Db.
		Where("status = ?", status).
		Where("order_time < ?", orderTime).
		Find(&orders).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		panic(errs.DBError)
	}
	return orders
}

// OrderDelivery 订单派送
func OrderDelivery(orderId string) {
	// 根据id查询订单
	order := getOrderById(orderId)
	// 校验订单是否存在，并且状态为Confirmed
	if order == nil || order.Status != constant.Confirmed {
		panic(errs.OrderStatusError)
	}
	// 更新订单状态,状态转为派送中
	order.Status = constant.DeliveryInProgress
	UpdateOrder(*order)
}

// OrderComplete 完成订单
func OrderComplete(orderId string) {
	// 根据id查询订单
	order := getOrderById(orderId)
	// 校验订单是否存在，并且状态为DeliveryInProgress
	if order == nil || order.Status != constant.DeliveryInProgress {
		panic(errs.OrderStatusError)
	}
	// 更新订单状态,状态转为完成
	order.Status = constant.Completed
	order.DeliveryTime = model.LocalTime(time.Now())
	UpdateOrder(*order)
}

// GetDailyTurnover 获取每日营业额
func GetDailyTurnover(begin, end time.Time) float64 {
	var turnover float64
	if err := global.Db.Table("orders").
		Where("status = ?", constant.Completed).
		Where("order_time >= ?", model.LocalTime(begin)).
		Where("order_time <= ?", model.LocalTime(end)).
		Select("ifnull(sum(amount),0) as amount").
		Scan(&turnover).Error; err != nil {
		panic(errs.DBError)
	}
	return turnover
}

// GetDailyOrderCount 获取每日订单数
func GetDailyOrderCount(begin, end time.Time, status int) int {
	var orderCount int64
	if status != 0 {
		// 查询有效订单数量
		if err := global.Db.Table("orders").
			Where("order_time >= ?", model.LocalTime(begin)).
			Where("order_time <= ?", model.LocalTime(end)).
			Where("status = ?", status).
			Count(&orderCount).Error; err != nil {
			panic(errs.DBError)
		}
	} else {
		// 查询总订单数量
		if err := global.Db.Table("orders").
			Where("order_time >= ?", model.LocalTime(begin)).
			Where("order_time <= ?", model.LocalTime(end)).
			Count(&orderCount).Error; err != nil {
			panic(errs.DBError)
		}
	}
	return int(orderCount)
}

// GetSalesTop10 获取销量前十的商品
func GetSalesTop10(begin, end time.Time) []dto.GoodsSalesDTO {
	goodsSales := make([]dto.GoodsSalesDTO, 0)
	if err := global.Db.Table("order_detail").
		Select("order_detail.name, sum(order_detail.number) as number").
		Joins("left join orders on order_detail.order_id = orders.id").
		Where("orders.status = ?", constant.Completed).
		Where("orders.order_time >= ?", begin).
		Where("orders.order_time <= ?", end).
		Group("order_detail.name").
		Order("number desc").
		Limit(10).Offset(0).
		Scan(&goodsSales).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		panic(errs.DBError)
	}
	return goodsSales
}
