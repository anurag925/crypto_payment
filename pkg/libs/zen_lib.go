package libs

import (
	"github.com/anurag925/crypto_payment/pkg/models"
	"time"
)

type ZenPaymentCreateRequest struct {
	Authorization         *ZenSeparateCurrencyAuthorization `json:"authorization,omitempty"`
	MerchantTransactionID string                            `json:"merchantTransactionId"`
	PaymentChannel        string                            `json:"paymentChannel"`
	Amount                string                            `json:"amount"`
	Currency              string                            `json:"currency"`
	CustomIpnURL          string                            `json:"customIpnUrl"`
	Comment               string                            `json:"comment"`
	FraudFields           ZenFingerPrint                    `json:"fraudFields"`
	Items                 []ZenItem                         `json:"items"`
	Customer              ZenCustomer                       `json:"customer"`
	PaymentSpecificData   any                               `json:"paymentSpecificData"`
}

type ZenSeparateCurrencyAuthorization struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

type ZenFingerPrint struct {
	FingerPrintID string `json:"fingerPrintId"`
}

type ZenItem struct {
	Code            string `json:"code"`
	Category        string `json:"category"`
	Name            string `json:"name"`
	Price           string `json:"price"`
	Quantity        int    `json:"quantity"`
	LineAmountTotal string `json:"lineAmountTotal"`
}

type ZenCustomer struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	IP        string `json:"ip"`
}

type OneTimePaymentSpecificData struct {
	Type            string                `json:"type,omitempty"`
	Descriptor      string                `json:"descriptor,omitempty"`
	Card            OnetimeCardDetails    `json:"card,omitempty"`
	Skip3Ds         bool                  `json:"skip3ds,omitempty"`
	BrowserDetails  models.BrowserDetails `json:"browserDetails,omitempty"`
	ReturnVerifyURL string                `json:"returnVerifyUrl,omitempty"`
}

type OnetimeCardDetails struct {
	Number     string `json:"number"`
	ExpiryDate string `json:"expiryDate"`
	Cvv        string `json:"cvv"`
}
type GeneralPaymentSpecificData struct {
	Type      string `json:"type,omitempty"`
	ReturnUrl string `json:"returnUrl,omitempty"`
}

type TokenPaymentSpecificData struct {
	Type            string                `json:"type,omitempty"`
	Descriptor      string                `json:"descriptor,omitempty"`
	BrowserDetails  models.BrowserDetails `json:"browserDetails,omitempty"`
	Token           string                `json:"token,omitempty"`
	ReturnVerifyURL string                `json:"returnVerifyUrl,omitempty"`
}

type ZenPaymentCreateResponse struct {
	RedirectURL           string    `json:"redirectUrl"`
	ID                    string    `json:"id"`
	MerchantTransactionID string    `json:"merchantTransactionId"`
	Amount                string    `json:"amount"`
	Currency              string    `json:"currency"`
	CreatedAt             time.Time `json:"createdAt"`
	ModifiedAt            time.Time `json:"modifiedAt"`
	Type                  string    `json:"type"`
	Status                string    `json:"status"`
	PaymentChannel        string    `json:"paymentChannel"`
	Actions               struct {
		Capture       bool `json:"capture"`
		Cancel        bool `json:"cancel"`
		Refund        bool `json:"refund"`
		Redirect      bool `json:"redirect"`
		Authorization bool `json:"authorization"`
	} `json:"actions"`
	FraudFields struct {
		FingerPrintID  string `json:"fingerPrintId"`
		Channel        string `json:"channel"`
		APIIntegration any    `json:"apiIntegration"`
	} `json:"fraudFields"`
	Meta struct {
		CustomIpnURL string `json:"customIpnUrl"`
		ThreeDs      struct {
			Xid string `json:"xid"`
		} `json:"threeDs"`
	} `json:"meta"`
	Customer struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Email     string `json:"email"`
		Phone     string `json:"phone"`
		IP        string `json:"ip"`
		Country   string `json:"country"`
	} `json:"customer"`
	Items []struct {
		Code                string `json:"code"`
		Category            string `json:"category"`
		Name                string `json:"name"`
		Price               string `json:"price"`
		AuthPrice           string `json:"authPrice"`
		Quantity            int    `json:"quantity"`
		LineAmountTotal     string `json:"lineAmountTotal"`
		AuthLineAmountTotal string `json:"authLineAmountTotal"`
	} `json:"items"`
	Cashback struct {
		Active bool `json:"active"`
	} `json:"cashback"`
	Source struct {
		Channel string `json:"channel"`
	} `json:"source"`
	MerchantAction struct {
		Action string `json:"action"`
		Data   struct {
			RedirectURL string `json:"redirectUrl"`
		} `json:"data"`
	} `json:"merchantAction"`
	VerifyReturnmac string `json:"verifyReturnmac"`
	CardInfo        struct {
		Bank              string `json:"bank"`
		Country           string `json:"country"`
		Organization      string `json:"organization"`
		OrganizationBrand string `json:"organizationBrand"`
		Segment           string `json:"segment"`
		Type              string `json:"type"`
		LastFourDigits    string `json:"lastFourDigits"`
		ExpirationDate    string `json:"expirationDate"`
		Bin               string `json:"bin"`
		Eea               bool   `json:"eea"`
		Commercial        bool   `json:"commercial"`
	} `json:"cardInfo"`
	Authorization struct {
		Amount   string `json:"amount"`
		Currency string `json:"currency"`
	} `json:"authorization"`
}

// type ZenPaymentCreateResponse struct {
// 	ID             string `json:"id"`
// 	RedirectURL    string `json:"redirectUrl"`
// 	ImageURL       string `json:"imageUrl"`
// 	MerchantAction struct {
// 		Action string `json:"action"`
// 		Data   struct {
// 			RedirectURL string `json:"redirectUrl"`
// 		} `json:"data"`
// 	} `json:"merchantAction"`
// 	MerchantTransactionID       string `json:"merchantTransactionId"`
// 	OriginMerchantTransactionID string `json:"originMerchantTransactionId"`
// 	Amount                      string `json:"amount"`
// 	Currency                    string `json:"currency"`
// 	Fee                         struct {
// 		Amount   string `json:"amount"`
// 		Currency string `json:"currency"`
// 	} `json:"fee"`
// 	SubsidiaryData struct {
// 		FeeAmount   string `json:"feeAmount"`
// 		GrossAmount string `json:"grossAmount"`
// 	} `json:"subsidiaryData"`
// 	Authorization struct {
// 		Amount   string `json:"amount"`
// 		Currency string `json:"currency"`
// 		Fee      string `json:"fee"`
// 	} `json:"authorization"`
// 	CreatedAt           time.Time `json:"createdAt"`
// 	ModifiedAt          time.Time `json:"modifiedAt"`
// 	Type                string    `json:"type"`
// 	Status              string    `json:"status"`
// 	TopupTransferStatus string    `json:"topupTransferStatus"`
// 	PaymentChannel      string    `json:"paymentChannel"`
// 	Actions             struct {
// 		Refund        bool `json:"refund"`
// 		Cancel        bool `json:"cancel"`
// 		Capture       bool `json:"capture"`
// 		Redirect      bool `json:"redirect"`
// 		Authorization bool `json:"authorization"`
// 	} `json:"actions"`
// 	FraudFields struct {
// 		Property1 string `json:"property1"`
// 		Property2 string `json:"property2"`
// 	} `json:"fraudFields"`
// 	RejectCode   string `json:"rejectCode"`
// 	RejectReason string `json:"rejectReason"`
// 	Refunds      []struct {
// 		ID                          string    `json:"id"`
// 		MerchantTransactionID       string    `json:"merchantTransactionId"`
// 		OriginMerchantTransactionID string    `json:"originMerchantTransactionId"`
// 		Amount                      string    `json:"amount"`
// 		Currency                    string    `json:"currency"`
// 		CreatedAt                   time.Time `json:"createdAt"`
// 		Status                      string    `json:"status"`
// 	} `json:"refunds"`
// 	Meta struct {
// 		PayoutBtcAddress  string `json:"payoutBtcAddress"`
// 		BtcAmount         string `json:"btcAmount"`
// 		FeeOwner          string `json:"feeOwner"`
// 		ReturnURL         string `json:"returnUrl"`
// 		CustomIpnURL      string `json:"customIpnUrl"`
// 		ReturnVerifyURL   string `json:"returnVerifyUrl"`
// 		AvsResult         string `json:"avsResult"`
// 		CvvResult         string `json:"cvvResult"`
// 		AuthorisationCode string `json:"authorisationCode"`
// 		ResultCode        string `json:"resultCode"`
// 		ThreeDs           struct {
// 			Xid           string `json:"xid"`
// 			Version       string `json:"version"`
// 			Eci           string `json:"eci"`
// 			Cavv          string `json:"cavv"`
// 			CavvAlgorithm string `json:"cavvAlgorithm"`
// 		} `json:"threeDs"`
// 		DestinationCurrency     string `json:"destinationCurrency"`
// 		DestinationAddress      string `json:"destinationAddress"`
// 		AmountToWithdraw        string `json:"amountToWithdraw"`
// 		PaymentCryptoAddress    string `json:"paymentCryptoAddress"`
// 		CryptoAmount            string `json:"cryptoAmount"`
// 		CryptoCurrencyShortName string `json:"cryptoCurrencyShortName"`
// 		CryptoCurrencyFullName  string `json:"cryptoCurrencyFullName"`
// 		CryptoNetworkFee        string `json:"cryptoNetworkFee"`
// 		Tracking                string `json:"tracking"`
// 		QrCodeData              string `json:"qrCodeData"`
// 		CryptoCurrency          string `json:"cryptoCurrency"`
// 		PromoCode               string `json:"promoCode"`
// 		WalletCurrency          string `json:"walletCurrency"`
// 		CurrencyName            string `json:"currencyName"`
// 		CurrencyCode            string `json:"currencyCode"`
// 		CurrencyLogoURL         string `json:"currencyLogoUrl"`
// 		Network                 string `json:"network"`
// 		NetworkCode             string `json:"networkCode"`
// 		NetworkLogoURL          string `json:"networkLogoUrl"`
// 	} `json:"meta"`
// 	Customer struct {
// 		ID          string `json:"id"`
// 		UserID      string `json:"userId"`
// 		TenantID    int    `json:"tenantId"`
// 		Segment     string `json:"segment"`
// 		FirstName   string `json:"firstName"`
// 		LastName    string `json:"lastName"`
// 		Email       string `json:"email"`
// 		Phone       string `json:"phone"`
// 		Information string `json:"information"`
// 		AccountID   string `json:"accountId"`
// 		IP          string `json:"ip"`
// 	} `json:"customer"`
// 	CardInfo struct {
// 		MerchantCardToken string `json:"merchantCardToken"`
// 		Bank              string `json:"bank"`
// 		Country           string `json:"country"`
// 		Organization      string `json:"organization"`
// 		OrganizationBrand string `json:"organizationBrand"`
// 		Token             string `json:"token"`
// 		Segment           string `json:"segment"`
// 		Type              string `json:"type"`
// 		// LastFourDigits    int    `json:"lastFourDigits"`
// 		// ExpirationDate    time.Time `json:"expirationDate"`
// 		Bin int `json:"bin"`
// 	} `json:"cardInfo"`
// 	BillingAddress struct {
// 		ID             string `json:"id"`
// 		UserID         string `json:"userId"`
// 		TenantID       int    `json:"tenantId"`
// 		Segment        string `json:"segment"`
// 		FirstName      string `json:"firstName"`
// 		LastName       string `json:"lastName"`
// 		Country        string `json:"country"`
// 		Street         string `json:"street"`
// 		City           string `json:"city"`
// 		CountryState   string `json:"countryState"`
// 		Province       string `json:"province"`
// 		BuildingNumber string `json:"buildingNumber"`
// 		RoomNumber     string `json:"roomNumber"`
// 		Postcode       string `json:"postcode"`
// 		CompanyName    string `json:"companyName"`
// 		Phone          string `json:"phone"`
// 		TaxID          string `json:"taxId"`
// 	} `json:"billingAddress"`
// 	ShippingAddress struct {
// 		ID             string `json:"id"`
// 		UserID         string `json:"userId"`
// 		TenantID       int    `json:"tenantId"`
// 		Segment        string `json:"segment"`
// 		FirstName      string `json:"firstName"`
// 		LastName       string `json:"lastName"`
// 		Country        string `json:"country"`
// 		Street         string `json:"street"`
// 		City           string `json:"city"`
// 		CountryState   string `json:"countryState"`
// 		Province       string `json:"province"`
// 		BuildingNumber string `json:"buildingNumber"`
// 		RoomNumber     string `json:"roomNumber"`
// 		Postcode       string `json:"postcode"`
// 		CompanyName    string `json:"companyName"`
// 		Phone          string `json:"phone"`
// 	} `json:"shippingAddress"`
// 	Items []struct {
// 		Code            string `json:"code"`
// 		Category        string `json:"category"`
// 		Type            string `json:"type"`
// 		Name            string `json:"name"`
// 		Price           string `json:"price"`
// 		Quantity        int    `json:"quantity"`
// 		LineAmountTotal string `json:"lineAmountTotal"`
// 	} `json:"items"`
// 	VerifyReturnmac string `json:"verifyReturnmac"`
// 	Cashback        struct {
// 		Active bool `json:"active"`
// 		Values []struct {
// 			Type  string `json:"type"`
// 			Value string `json:"value"`
// 		} `json:"values"`
// 		Client struct {
// 			ID       string `json:"id"`
// 			Type     string `json:"type"`
// 			TenantID int    `json:"tenantId"`
// 			Email    string `json:"email"`
// 		} `json:"client"`
// 	} `json:"cashback"`
// 	Source struct {
// 		Channel         string `json:"channel"`
// 		PluginName      string `json:"pluginName"`
// 		PluginVersion   string `json:"pluginVersion"`
// 		PlatformName    string `json:"platformName"`
// 		PlatformVersion string `json:"platformVersion"`
// 	} `json:"source"`
// }

type CallbackRequest struct {
	Type                  string `json:"type"`
	TransactionID         string `json:"transactionId"`
	MerchantTransactionID string `json:"merchantTransactionId"`
	Amount                string `json:"amount"`
	Currency              string `json:"currency"`
	Status                string `json:"status"`
	Hash                  string `json:"hash"`
	Signature             string `json:"signature"`
	PaymentMethod         struct {
		Name       string `json:"name"`
		Channel    string `json:"channel"`
		Parameters struct {
		} `json:"parameters"`
	} `json:"paymentMethod"`
	Customer struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Email     string `json:"email"`
		IP        string `json:"ip"`
		Country   string `json:"country"`
	} `json:"customer"`
	SecurityStatus string `json:"securityStatus"`
	RiskData       struct {
	} `json:"riskData"`
	Email string `json:"email"`
}
