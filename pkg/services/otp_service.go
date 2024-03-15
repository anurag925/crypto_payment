package services

import (
	"context"
	"github.com/anurag925/crypto_payment/pkg/models"

	"gopkg.in/guregu/null.v4"
)

type GenerateOtp struct {
	Receiver      string           `json:"receiver" validate:"required"`
	Type          models.OtpType   `json:"type" validate:"required"`
	Action        models.OtpAction `json:"action" validate:"required"`
	VerifyingID   null.Int         `json:"verifying_id"`
	VerifyingType null.String      `json:"verifying_type"`
}

type VerifyOtp struct {
	Receiver string           `json:"receiver"`
	Type     models.OtpType   `json:"type"`
	Action   models.OtpAction `json:"action"`
	Value    int              `json:"value"`
}

type OtpService interface {
	SendOtpMail(ctx context.Context, o GenerateOtp) error
	Verify(ctx context.Context, o VerifyOtp) error
}
