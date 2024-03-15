package impl

import (
	"context"
	"fmt"
	"testing"

	"github.com/anurag925/crypto_payment/pkg/models"
	"github.com/anurag925/crypto_payment/pkg/repositories"
	"github.com/anurag925/crypto_payment/pkg/repositories/postgresql"
	"github.com/anurag925/crypto_payment/pkg/services"

	"gopkg.in/guregu/null.v4"
)

func Test_otpServiceImpl_Generate(t *testing.T) {
	type fields struct {
		otpRepo repositories.OtpRepository
	}
	type args struct {
		ctx context.Context
		o   services.GenerateOtp
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    models.Otp
		wantErr bool
	}{
		{
			name:    "success otp generation",
			fields:  fields{postgresql.DefaultOtpRepositoryImpl()},
			args:    args{context.Background(), services.GenerateOtp{Receiver: "test@gmail.com", Type: models.OtpTypeEmail, Action: models.OtpActionTransactionOtp, VerifyingID: null.IntFrom(1), VerifyingType: null.StringFrom("Order")}},
			want:    models.Otp{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &otpServiceImpl{
				otpRepo: tt.fields.otpRepo,
			}
			got, err := s.Generate(tt.args.ctx, tt.args.o)
			if (err != nil) != tt.wantErr {
				t.Errorf("otpServiceImpl.Generate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println(got)
			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("otpServiceImpl.Generate() = %v, want %v", got, tt.want)
			// }
		})
	}
}

func Test_otpServiceImpl_SendOtpMail(t *testing.T) {
	type fields struct {
		otpRepo repositories.OtpRepository
	}
	type args struct {
		ctx context.Context
		o   services.GenerateOtp
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "sending otp to dev",
			fields: fields{postgresql.DefaultOtpRepositoryImpl()},
			args: args{context.Background(), services.GenerateOtp{
				Receiver:      "dev@crypto_payment.io",
				Type:          models.OtpTypeEmail,
				Action:        models.OtpActionTransactionOtp,
				VerifyingID:   null.IntFrom(1),
				VerifyingType: null.StringFrom("Order"),
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &otpServiceImpl{
				otpRepo: tt.fields.otpRepo,
			}
			if err := s.SendOtpMail(tt.args.ctx, tt.args.o); (err != nil) != tt.wantErr {
				t.Errorf("otpServiceImpl.SendOtpMail() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
