package scripts

import (
	"fmt"

	"github.com/anurag925/crypto_payment/pkg/mailers"
	"github.com/anurag925/crypto_payment/pkg/models"
)

func SendTransactionOtpMail() {
	err := mailers.NewOtpMailer().SendTransactionOtpMail(models.Account{
		Email: "dev@crypto_payment.io",
	}, 123456)
	fmt.Printf("error %+v", err)
}
