package impl

import (
	"context"
	"github.com/anurag925/crypto_payment/pkg/models"
	"github.com/anurag925/crypto_payment/pkg/services"
	"github.com/anurag925/crypto_payment/utils/encryption/seed"
	"github.com/anurag925/crypto_payment/utils/logger"
)

type retailersAdminHelperServiceImpl struct {
	accountService      services.AccountService
	retailerService     services.RetailerService
	addressService      services.AddressService
	contactService      services.ContactService
	kycService          services.KycService
	documentService     services.DocumentService
	cryptoWalletService services.CryptoWalletService
	txnFeeConfigService services.TxnFeeConfigService
}

func NewRetailersAdminHelperServiceImpl(
	accountService services.AccountService,
	retailerService services.RetailerService,
	addressService services.AddressService,
	contactService services.ContactService,
	kycService services.KycService,
	documentService services.DocumentService,
	cryptoWalletService services.CryptoWalletService,
	txnFeeConfigService services.TxnFeeConfigService,
) *retailersAdminHelperServiceImpl {
	return &retailersAdminHelperServiceImpl{
		accountService:      accountService,
		retailerService:     retailerService,
		addressService:      addressService,
		contactService:      contactService,
		kycService:          kycService,
		documentService:     documentService,
		cryptoWalletService: cryptoWalletService,
		txnFeeConfigService: txnFeeConfigService,
	}
}

func DefaultRetailersAdminHelperServiceImpl() *retailersAdminHelperServiceImpl {
	return NewRetailersAdminHelperServiceImpl(
		DefaultAccountServiceImpl(),
		DefaultRetailerServiceImpl(),
		DefaultAddressServiceImpl(),
		DefaultContactServiceImpl(),
		DefaultKycServiceImpl(),
		DefaultDocumentServiceImpl(),
		DefaultCryptoWalletServiceImpl(),
		DefaultTxnFeeConfigServiceImpl(),
	)
}

func (r *retailersAdminHelperServiceImpl) CreateRetailer(ctx context.Context, req services.RetailersCreateData) (services.RetailersDetail, error) {
	res := services.RetailersDetail{
		Shareholders: []models.Contact{},
		Documents:    []models.Document{},
	}
	logger.Info(ctx, "creating retailer")
	account := &req.Account
	account.Password.SetValid(seed.Generate(10))
	account.Type = models.AccountTypeRetailer
	if err := r.accountService.Create(ctx, account); err != nil {
		logger.Error(ctx, "error create account", "error", err)
		return res, err
	}
	res.Account = *account
	logger.Info(ctx, "account creation successful")
	retailer := &req.Retailer
	retailer.AccountID = account.ID
	if err := r.retailerService.Create(ctx, retailer); err != nil {
		logger.Error(ctx, "error create retailer", "error", err)
		return res, err
	}
	res.Retailer = *retailer
	logger.Info(ctx, "retailer creation successful")
	kyc := &models.Kyc{}
	kyc.Status = models.KycStatusPending
	kyc.AccountID = account.ID
	if err := r.kycService.Create(ctx, kyc); err != nil {
		logger.Error(ctx, "error create kyc", "error", err)
		return res, err
	}
	res.Kyc = *kyc
	logger.Info(ctx, "kyc creation successful")
	address := &req.Address
	address.AccountID = account.ID
	address.Type = models.AddressTypeBusiness
	if err := r.addressService.Create(ctx, address); err != nil {
		logger.Error(ctx, "error create address", "error", err)
		return res, err
	}
	res.Address = *address
	logger.Info(ctx, "address creation successful")
	contacts := &req.Contacts
	for _, contact := range *contacts {
		contact.RetailerID = retailer.ID
		if err := r.contactService.Create(ctx, &contact); err != nil {
			logger.Error(ctx, "error create contact", "error", err)
			return res, err
		}
		if contact.PointOfContact {
			res.Contact = contact
		} else {
			res.Shareholders = append(res.Shareholders, contact)
		}
	}
	logger.Info(ctx, "contacts creation successful")
	documents := req.Documents
	for _, document := range documents {
		document.KycID = kyc.ID
		if err := r.documentService.Create(ctx, &document); err != nil {
			logger.Error(ctx, "error create contact", "error", err)
			return res, err
		}
		res.Documents = append(res.Documents, document)
	}
	logger.Info(ctx, "documents creation successful")
	logger.Info(ctx, "creating retailer done")
	return res, nil
}

func (r *retailersAdminHelperServiceImpl) CreateRetailersConfigurations(ctx context.Context, req services.RetailersCreateData) (services.RetailersDetail, error) {
	logger.Info(ctx, "creating retailer configurations")
	res := services.RetailersDetail{
		CryptoWallets: []models.CryptoWallet{},
		TxnFeeConfigs: []models.TxnFeeConfig{},
	}
	retailerConfigs := req.Retailer
	retailer, err := r.retailerService.FindById(ctx, retailerConfigs.ID)
	if err != nil {
		logger.Error(ctx, "error find retailer", "error", err)
		return res, err
	}
	retailer.DailyLimit = retailerConfigs.DailyLimit
	retailer.WeeklyLimit = retailerConfigs.WeeklyLimit
	retailer.MonthlyLimit = retailerConfigs.MonthlyLimit
	retailer.SingleTxnLimit = retailerConfigs.SingleTxnLimit
	if err := r.retailerService.Update(ctx, &retailer); err != nil {
		logger.Error(ctx, "error update retailer", "error", err)
		return res, err
	}
	res.Retailer = retailer
	logger.Info(ctx, "retailer update successful")
	for _, cryptoWallet := range req.CryptoWallets {
		if err := r.cryptoWalletService.CreateCryptoWallet(ctx, retailer, &cryptoWallet); err != nil {
			logger.Error(ctx, "error create contact", "error", err)
			return res, err
		}
		res.CryptoWallets = append(res.CryptoWallets, cryptoWallet)
	}
	logger.Info(ctx, "cryptoWallet creation successful")
	for _, txnFeeConfig := range req.TxnFeeConfigs {
		txnFeeConfig.RetailerID = retailer.ID
		if err := r.txnFeeConfigService.Create(ctx, &txnFeeConfig); err != nil {
			logger.Error(ctx, "error create contact", "error", err)
			return res, err
		}
		res.TxnFeeConfigs = append(res.TxnFeeConfigs, txnFeeConfig)
	}
	logger.Info(ctx, "txnFeeConfig creation successful")
	logger.Info(ctx, "creating retailer configurations done")
	return res, nil
}
