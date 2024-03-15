package aes

import (
	"testing"
)

func TestEncrypt(t *testing.T) {
	type args struct {
		keyString       string
		stringToEncrypt string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Success",
			args: args{
				keyString:       "34753778214125442A472D4B6150645367566B59703373357638792F423F4528",
				stringToEncrypt: "this is the fake string",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotEncryptedString, err := Encrypt(tt.args.keyString, tt.args.stringToEncrypt)
			if (err != nil) != tt.wantErr {
				t.Errorf("Encrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotEncryptedString == tt.args.stringToEncrypt {
				t.Errorf("Encrypt() = %v, want %v", gotEncryptedString, tt.args.stringToEncrypt)
			}
		})
	}
}

func TestDecrypt(t *testing.T) {
	type args struct {
		keyString       string
		stringToDecrypt string
	}
	tests := []struct {
		name                string
		args                args
		wantDecryptedString string
		wantErr             bool
	}{
		{
			name: "Success",
			args: args{
				keyString:       "34753778214125442A472D4B6150645367566B59703373357638792F423F4528",
				stringToDecrypt: "5UjZXapYwS6YZUCEhxJWiFEHl3T2naOGIDcOuQQmY5dw0ZGuZfUmJ1Ahb8cSW5OW",
			},
			wantDecryptedString: "this is the fake string",
			wantErr:             false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDecryptedString, err := Decrypt(tt.args.keyString, tt.args.stringToDecrypt)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotDecryptedString != tt.wantDecryptedString {
				t.Errorf("Decrypt() = %v, want %v", gotDecryptedString, tt.wantDecryptedString)
			}
		})
	}
}
