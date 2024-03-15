package impl

import (
	"context"
	"errors"
	"fmt"
	"github.com/anurag925/crypto_payment/app"
	"github.com/anurag925/crypto_payment/app/configs"
	"github.com/anurag925/crypto_payment/pkg/mailers"
	"github.com/anurag925/crypto_payment/pkg/models"
	"github.com/anurag925/crypto_payment/pkg/repositories"
	"github.com/anurag925/crypto_payment/pkg/repositories/postgresql"
	"github.com/anurag925/crypto_payment/pkg/services"
	"github.com/anurag925/crypto_payment/utils/logger"
	"math/rand"
	"reflect"
	"strings"
	"time"

	"gorm.io/gorm"
)

const validOtpDuration = 15 * time.Minute

type otpServiceImpl struct {
	otpRepo repositories.OtpRepository
}

func NewOtpServiceImpl(otpRepo repositories.OtpRepository) *otpServiceImpl {
	return &otpServiceImpl{otpRepo: otpRepo}
}

func DefaultOtpServiceImpl() *otpServiceImpl {
	return NewOtpServiceImpl(postgresql.DefaultOtpRepositoryImpl())
}

func (s *otpServiceImpl) SendOtpMail(ctx context.Context, o services.GenerateOtp) error {
	action := snakeToPascal(o.Action.String())
	logger.Info(ctx, "calling mailer action", "action", action)
	stringMethodName := fmt.Sprintf("Send%sMail", action)
	logger.Info(ctx, "calling mailer method name", "stringMethodName", stringMethodName)
	method := reflect.ValueOf(mailers.NewOtpMailer()).MethodByName(stringMethodName)
	logger.Info(ctx, "calling mailer method", "method", method)
	if !method.IsValid() {
		logger.Error(ctx, "the mailing method does not exists", "error", method)
		return errors.New("the mailing method does not exists")
	}
	account := models.Account{Email: o.Receiver}
	otp, err := s.Generate(ctx, o)
	if err != nil {
		return err
	}
	logger.Debug(ctx, "sending otp for verification", "otp", otp)
	// Call the method with an empty slice of arguments
	returnValues := method.Call([]reflect.Value{reflect.ValueOf(account), reflect.ValueOf(otp.Value)})
	if len(returnValues) > 0 {
		err := returnValues[0].Interface()
		if err != nil {
			logger.Error(ctx, "error in sending mail ", "error", err)
			return err.(error)
		}
	}

	return nil
}

func (s *otpServiceImpl) Generate(ctx context.Context, o services.GenerateOtp) (models.Otp, error) {
	newOtp, err := s.otpRepo.LastActive(ctx, o.Receiver, o.Type, o.Action)
	if (err != nil && errors.Is(err, gorm.ErrRecordNotFound)) || s.isExpired(ctx, newOtp.CreatedAt) {
		logger.Info(ctx, "generate new otp", "error", err)
		newOtp, err = s.generateNewOtp(ctx, o)
		if err != nil {
			return models.Otp{}, err
		}
	} else if err != nil {
		return models.Otp{}, err
	} else if s.isExpired(ctx, newOtp.CreatedAt) {
		logger.Info(ctx, "generate new otp cuz otp expired", "error", err)
		newOtp, err = s.generateNewOtp(ctx, o)
		if err != nil {
			return models.Otp{}, err
		}
	}
	return newOtp, nil
}

func (s *otpServiceImpl) Verify(ctx context.Context, o services.VerifyOtp) error {
	if app.Config().Env != configs.Production {
		if o.Value != 123456 {
			return errors.New("otp is mismatched")
		}
		return nil
	}
	otp, err := s.otpRepo.LastActive(ctx, o.Receiver, o.Type, o.Action)
	if err != nil {
		return err
	}
	if s.isExpired(ctx, otp.CreatedAt) {
		return errors.New("otp is expired")
	}
	if otp.Value != o.Value {
		return errors.New("otp is mismatched")
	}
	return nil
}

func (s *otpServiceImpl) generateNewOtp(ctx context.Context, o services.GenerateOtp) (models.Otp, error) {
	otp := models.Otp{
		Type:          o.Type,
		Value:         generateRandomNumber(),
		Action:        o.Action,
		Receiver:      o.Receiver,
		VerifyingID:   o.VerifyingID,
		VerifyingType: o.VerifyingType,
		Verified:      false,
		RetryCount:    0,
	}
	// if err := s.otpRepo.Create(ctx, &otp); err != nil {
	// 	return models.Otp{}, err
	// }
	return otp, nil
}

func (s *otpServiceImpl) isExpired(ctx context.Context, createdAt time.Time) bool {
	return time.Since(createdAt) > validOtpDuration
}

func generateRandomNumber() int {
	min := 100000 // Minimum 6-digit number
	max := 999999 // Maximum 6-digit number
	return rand.Intn(max-min+1) + min
}

func snakeToPascal(snakeCase string) string {
	words := strings.Split(snakeCase, "_")
	for i := 0; i < len(words); i++ {
		words[i] = strings.ToUpper(words[i][:1]) + words[i][1:]
	}
	return strings.Join(words, "")
}
