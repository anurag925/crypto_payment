package impl

import (
	"context"
	"fmt"
	"github.com/anurag925/crypto_payment/pkg/models"
	"github.com/anurag925/crypto_payment/pkg/repositories"
	"github.com/anurag925/crypto_payment/pkg/services"
	"reflect"
	"testing"
)

func Test_accountServiceImpl_Login(t *testing.T) {
	type fields struct {
		accountRepo repositories.AccountRepository
	}
	type args struct {
		ctx context.Context
		req services.LoginRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    services.LoginResponse
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &accountServiceImpl{
				accountRepo: tt.fields.accountRepo,
			}
			got, err := s.Login(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("accountServiceImpl.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("accountServiceImpl.Login() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_accountServiceImpl_generateUserToken(t *testing.T) {
	type fields struct {
		accountRepo repositories.AccountRepository
	}
	type args struct {
		ctx context.Context
		a   models.Account
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name:   "success",
			fields: fields{},
			args: args{
				ctx: context.Background(),
				a:   models.Account{ID: 1, Email: "anuragle100@gmail.com"},
			},
			want:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFudXJhZ2xlMTAwQGdtYWlsLmNvbSIsImlkIjoxfQ._lqV_C-lGveqcUYjrgQYqjnOkFs2gNcxlpmFPhOs1OY",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &accountServiceImpl{
				accountRepo: tt.fields.accountRepo,
			}
			got, err := s.generateUserToken(tt.args.ctx, tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("accountServiceImpl.generateUserToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("accountServiceImpl.generateUserToken() = %v, want %v", got, tt.want)
			}
			fmt.Println(got)
		})
	}
}
