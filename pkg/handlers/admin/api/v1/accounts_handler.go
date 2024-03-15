package v1

import (
	"github.com/anurag925/crypto_payment/pkg/handlers"

	"github.com/labstack/echo/v4"
)

type Report struct {
	Title      string `json:"title"`
	Value      string `json:"value"`
	BadgeValue string `json:"badgeValue"`
	Color      string `json:"color"`
	Result     string `json:"result"`
}

type Metric struct {
	Title    string   `json:"title"`
	Currency string   `json:"currency,omitempty"`
	Reports  []Report `json:"reports"`
}

var fakeDashboardData = []Metric{
	{
		Title:    "Processing Volume",
		Currency: "EUR",
		Reports: []Report{
			{
				Title:      "Daily",
				Value:      "1020000",
				BadgeValue: "2",
				Color:      "success",
				Result:     "up",
			},
			{
				Title:      "Weekly",
				Value:      "12000000",
				BadgeValue: "2",
				Color:      "danger",
				Result:     "down",
			},
			{
				Title:      "Monthly",
				Value:      "350000",
				BadgeValue: "2",
				Color:      "success",
				Result:     "up",
			},
			{
				Title:      "Overall",
				Value:      "4280000000",
				BadgeValue: "2",
				Color:      "success",
				Result:     "up",
			},
		},
	},
	{
		Title: "Number of Transactions",
		Reports: []Report{
			{
				Title:      "Daily",
				Value:      "1200",
				BadgeValue: "2",
				Color:      "success",
				Result:     "up",
			},
			{
				Title:      "Weekly",
				Value:      "120000",
				BadgeValue: "2",
				Color:      "danger",
				Result:     "down",
			},
			{
				Title:      "Monthly",
				Value:      "1207767676",
				BadgeValue: "2",
				Color:      "success",
				Result:     "up",
			},
			{
				Title:      "Overall",
				Value:      "123727642",
				BadgeValue: "2",
				Color:      "success",
				Result:     "up",
			},
		},
	},
	{
		Title:    "Fees earned crypto_payment",
		Currency: "EUR",
		Reports: []Report{
			{
				Title:      "Daily",
				Value:      "1020",
				BadgeValue: "2",
				Color:      "success",
				Result:     "up",
			},
			{
				Title:      "Weekly",
				Value:      "400",
				BadgeValue: "2",
				Color:      "danger",
				Result:     "down",
			},
			{
				Title:      "Monthly",
				Value:      "40000",
				BadgeValue: "2",
				Color:      "success",
				Result:     "up",
			},
			{
				Title:      "Overall",
				Value:      "1200000",
				BadgeValue: "2",
				Color:      "success",
				Result:     "up",
			},
		},
	},
	{
		Title:    "Fees earned Payment Gateway",
		Currency: "EUR",
		Reports: []Report{
			{
				Title:      "Daily",
				Value:      "1020",
				BadgeValue: "2",
				Color:      "success",
				Result:     "up",
			},
			{
				Title:      "Weekly",
				Value:      "4000",
				BadgeValue: "2",
				Color:      "danger",
				Result:     "down",
			},
			{
				Title:      "Monthly",
				Value:      "4000",
				BadgeValue: "2",
				Color:      "success",
				Result:     "up",
			},
			{
				Title:      "Overall",
				Value:      "120000",
				BadgeValue: "2",
				Color:      "success",
				Result:     "up",
			},
		},
	},
}

type accountHandler struct {
}

func NewAccountHandler() *accountHandler {
	return &accountHandler{}
}

func (h *accountHandler) Dashboard(c echo.Context) error {
	return handlers.SuccessResponse(c, fakeDashboardData)
}
