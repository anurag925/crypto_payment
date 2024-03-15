package impl

import (
	"context"
	"testing"
)

func Test_coinbaseLibImpl_GetUSDTtoEUROExchangeRate(t *testing.T) {
	type args struct {
		ctx      context.Context
		currency string
	}
	tests := []struct {
		name    string
		l       coinbaseLibImpl
		args    args
		want    float64
		wantErr bool
	}{
		{
			name:    "Get valid exchange rate",
			l:       coinbaseLibImpl{},
			args:    args{context.Background(), "EUR"},
			want:    0.9,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := coinbaseLibImpl{}
			got, err := l.GetUSDTtoEUROExchangeRate(tt.args.ctx, "EUR")
			if (err != nil) != tt.wantErr {
				t.Errorf("coinbaseLibImpl.GetUSDTtoEUROExchangeRate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == 0 {
				t.Errorf("coinbaseLibImpl.GetUSDTtoEUROExchangeRate() = %v, want %v", got, tt.want)
			}
		})
	}
}
