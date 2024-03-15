package impl

import (
	"context"
	"encoding/json"
	"github.com/anurag925/crypto_payment/app"
	"github.com/anurag925/crypto_payment/app/core"
	"github.com/anurag925/crypto_payment/pkg/libs"
	"testing"
)

func init() {
	app.New(core.GetBackendApp())
}

func Test_zenPaymentLibImpl_Callback(t *testing.T) {
	type args struct {
		ctx     context.Context
		request string
	}
	tests := []struct {
		name    string
		l       *zenPaymentLibImpl
		args    args
		wantErr bool
	}{
		{
			name: "success",
			l:    &zenPaymentLibImpl{},
			args: args{
				ctx: context.Background(),
				request: `{
					"type": "TRT_PURCHASE",
					"transactionId": "2d36ff20-017d-4c63-b626-407edb369cc2",
					"merchantTransactionId": "feb78e88-47bc-428a-8ea4-806535aaf2de",
					"amount": "100",
					"currency": "PLN",
					"status": "PENDING",
					"hash": "28EE6604A8A40ACC8B8CE0B8DE9AAC87A4E24BBF0388A48ED164E512C8073C7E",
					"signature": "D3739E5ADCC20E436DEE8F386C81B1C3ACCCE0558FAB65EC924564F998F983EE",
					"paymentMethod": {
					  "name": "PME_PBZ",
					  "channel": "PCL_PBZ",
					  "parameters": {}
					},
					"customer": {
					  "firstName": "John",
					  "lastName": "Doe",
					  "email": "john@doe.pl",
					  "ip": "172.89.0.1",
					  "country": "US"
					},
					"securityStatus": "pending",
					"riskData": {},
					"email": "john@doe.pl"
				  }`,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &zenPaymentLibImpl{}
			request := libs.CallbackRequest{}
			if err := json.Unmarshal([]byte(tt.args.request), &request); err != nil {
				t.Errorf("zenPaymentLibImpl.Callback() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err := l.Callback(tt.args.ctx, request); (err != nil) != tt.wantErr {
				t.Errorf("zenPaymentLibImpl.Callback() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_zenPaymentLibImpl_checkHash(t *testing.T) {
	type args struct {
		ctx        context.Context
		s          string
		hashString string
	}
	tests := []struct {
		name string
		l    *zenPaymentLibImpl
		args args
		want bool
	}{
		{
			name: "success",
			l:    &zenPaymentLibImpl{},
			args: args{
				ctx:        context.TODO(),
				s:          "feb78e88-47bc-428a-8ea4-806535aaf2dePLN100PENDINGaeb8e7bf-0009-4f30-b521-1136fd336ae6",
				hashString: "28EE6604A8A40ACC8B8CE0B8DE9AAC87A4E24BBF0388A48ED164E512C8073C7E",
			},
			want: true,
		},
		{
			name: "success",
			l:    &zenPaymentLibImpl{},
			args: args{
				ctx:        context.TODO(),
				s:          "feb78e88-47bc-428a-8ea4-806535aaf2dePLN100PENDINGaeb8e7bf-0009-4f30-b521-1136fd336ae6",
				hashString: "28EE6604A8A40ACC8B8CE0B8DE9AAC87A4E24BBF0388A48E64E512C8073C7E",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &zenPaymentLibImpl{}
			if got := l.checkHash(tt.args.ctx, tt.args.s, tt.args.hashString); got != tt.want {
				t.Errorf("zenPaymentLibImpl.checkHash() = %v, want %v", got, tt.want)
			}
		})
	}
}
