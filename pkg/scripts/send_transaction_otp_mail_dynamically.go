package scripts

import (
	"fmt"
	"reflect"

	"github.com/anurag925/crypto_payment/pkg/mailers"
	"github.com/anurag925/crypto_payment/pkg/models"
)

func SendTransactionOtpMailDynamically() {
	method := reflect.ValueOf(mailers.NewOtpMailer()).MethodByName("SendTransactionOtpMail")
	// Check if the method exists
	if method.IsValid() {
		// Call the method with an empty slice of arguments
		returnValues := method.Call([]reflect.Value{reflect.ValueOf(models.Account{Email: "dev@crypto_payment.io"}), reflect.ValueOf(123456)})
		if len(returnValues) > 0 {
			err := returnValues[0].Interface()
			if err != nil {
				// Handle the error
				fmt.Printf("Error: %v\n", err)
			}
		}
	} else {
		fmt.Printf("Method %s does not exist\n", "SendTransactionOtpMail")
	}
	// fmt.Printf("error %+v", err)
}
