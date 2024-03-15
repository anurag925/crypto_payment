package v1

import (
	"github.com/anurag925/crypto_payment/pkg/handlers"
	"github.com/anurag925/crypto_payment/pkg/services"
	"github.com/anurag925/crypto_payment/pkg/services/impl"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type orderHandler struct {
	service services.OrderService
}

func NewOrderHandler(service services.OrderService) *orderHandler {
	return &orderHandler{service: service}
}

func (h *orderHandler) PostOrder(c echo.Context) error {
	request := services.OrderCreateRequest{}
	if err := c.Bind(&request); err != nil {
		return handlers.BadRequestResponse(c, "invalid request data", err)
	}
	if err := impl.Validator().Struct(&request); err != nil {
		return handlers.BadRequestResponse(c, "data validation failed", err)
	}
	order, err := h.service.CreateBuyOrder(handlers.Context(c), request)
	if err != nil {
		// error sent at internal server error status 500
		return err
	}
	return handlers.CreatedResponse(c, order)
}

func (h *orderHandler) GetOrdersForRetailer(c echo.Context) error {
	startDate, err := time.Parse("02/01/2006", c.QueryParam("start_date"))
	if err != nil {
		return handlers.BadRequestResponse(c, "invalid start date", err)
	}
	endDate, err := time.Parse("02/01/2006", c.QueryParam("end_date"))
	if err != nil {
		return handlers.BadRequestResponse(c, "invalid end date", err)
	}
	page, err := strconv.ParseInt(c.QueryParam("page"), 10, 64)
	if err != nil {
		return handlers.BadRequestResponse(c, "invalid page number", err)
	}
	orders, err := h.service.OrdersForRetailer(handlers.Context(c), handlers.Retailer(c), startDate, endDate, int(page))
	if err != nil {
		// error sent at internal server error status 500
		return err
	}
	return handlers.SuccessResponse(c, orders)
}

func (h *orderHandler) GetOrderDetailForRetailer(c echo.Context) error {
	order_id, err := strconv.ParseInt(c.Param("order_id"), 10, 64)
	if err != nil {
		return handlers.BadRequestResponse(c, "invalid order id", err)
	}
	orderDetail, err := h.service.OrderDetailForRetailer(handlers.Context(c), handlers.Retailer(c), order_id)
	if err != nil {
		return err
	}
	return handlers.SuccessResponse(c, orderDetail)
}

func (h *orderHandler) GetSettlementAmount(c echo.Context) error {
	request := services.SettlementRequest{}
	if err := c.Bind(&request); err != nil {
		return handlers.BadRequestResponse(c, "invalid request data", err)
	}
	settlementAmount, err := h.service.SettlementAmount(handlers.Context(c), request)
	if err != nil {
		return err
	}
	return handlers.SuccessResponse(c, settlementAmount)
}
