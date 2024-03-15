package v1

import (
	"fmt"
	"github.com/anurag925/crypto_payment/pkg/handlers"
	"github.com/anurag925/crypto_payment/pkg/models"
	"github.com/anurag925/crypto_payment/pkg/services"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type documentHandler struct {
	service services.DocumentService
}

func NewDocumentHandler(s services.DocumentService) *documentHandler {
	return &documentHandler{s}
}

func (h *documentHandler) CreateDocument(c echo.Context) error {
	d := models.Document{}
	if err := c.Bind(&d); err != nil {
		return handlers.BadRequestResponse(c, "bad document body given", err)
	}
	if err := h.service.Create(handlers.Context(c), &d); err != nil {
		return err
	}
	return handlers.CreatedResponse(c, d)
}

type fileUploadData struct {
	KycID        string `json:"kyc_id" query:"kyc_id" validate:"required"`
	DocumentType string `json:"document_type" query:"document_type" validate:"required"`
	DocumentName string `json:"document_name" query:"document_name" validate:"required"`
	Access       string `json:"access" query:"access" validate:"required"`
	ContentType  string `json:"content_type" query:"content_type" validate:"required"`
}

func (h *documentHandler) UploadDocument(c echo.Context) error {
	data := fileUploadData{}
	if err := c.Bind(&data); err != nil {
		return handlers.BadRequestResponse(c, "bad document body given", err)
	}
	if err := validator.New().Struct(data); err != nil {
		return handlers.BadRequestResponse(c, fmt.Sprintf("bad document body given error: %v", err), err)
	}
	path := fmt.Sprintf("%s/%s/%s", data.KycID, data.DocumentType, data.DocumentName)
	url, err := h.service.Upload(handlers.Context(c), path, data.Access, data.ContentType)
	if err != nil {
		return err
	}
	return handlers.SuccessResponse(c, url)
}
