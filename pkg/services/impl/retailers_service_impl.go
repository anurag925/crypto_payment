package impl

import (
	"context"
	"errors"
	"github.com/anurag925/crypto_payment/pkg/models"
	"github.com/anurag925/crypto_payment/pkg/repositories"
	"github.com/anurag925/crypto_payment/pkg/repositories/postgresql"
	"github.com/anurag925/crypto_payment/pkg/services"
	"github.com/anurag925/crypto_payment/utils/logger"

	"gorm.io/gorm"
)

type retailerServiceImpl struct {
	retailerRepo repositories.RetailerRepository
}

var _ services.RetailerService = (*retailerServiceImpl)(nil)

func NewRetailerServiceImpl(retailerRepo repositories.RetailerRepository) *retailerServiceImpl {
	return &retailerServiceImpl{retailerRepo: retailerRepo}
}

func DefaultRetailerServiceImpl() *retailerServiceImpl {
	return NewRetailerServiceImpl(postgresql.DefaultRetailerRepositoryImpl())
}

func (s *retailerServiceImpl) FindById(ctx context.Context, id int64) (models.Retailer, error) {
	return s.retailerRepo.FindById(ctx, id)
}

func (s *retailerServiceImpl) AllRetailers(ctx context.Context) (services.AdminInfoRetailers, error) {
	retailers, err := s.retailerRepo.FindAll(ctx)
	if err != nil {
		logger.Error(ctx, "retailer fetch error", "error", err)
		return services.AdminInfoRetailers{}, err
	}
	adminRetailerInfo := make([]services.AdminInfoRetailer, 0)
	for _, retailer := range retailers {
		wallet, err := s.retailerRepo.Wallet(ctx, &retailer)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			logger.Error(ctx, "wallet fetch error", "error", err)
			return services.AdminInfoRetailers{}, err
		}
		retailerInfo := services.AdminInfoRetailer{
			ID:               retailer.ID,
			RetailerName:     retailer.RetailName.String,
			Status:           retailer.Account.Status.String(),
			OTV:              retailer.OverallTxnValue.String,
			MTV:              retailer.MonthlyLimit.String,
			PendingPayout:    wallet.Pending.String,
			AvailableBalance: wallet.Balance.String,
		}
		adminRetailerInfo = append(adminRetailerInfo, retailerInfo)
	}
	return services.AdminInfoRetailers{Retailers: adminRetailerInfo}, nil
}

func (s *retailerServiceImpl) RetailersDetail(ctx context.Context, retailerId int64) (services.RetailersDetail, error) {
	retailer, err := s.retailerRepo.FindById(ctx, retailerId)
	if err != nil {
		logger.Error(ctx, "retailer fetch error", "error", err)
		return services.RetailersDetail{}, err
	}
	account, err := s.retailerRepo.Account(ctx, &retailer)
	if err != nil {
		logger.Error(ctx, "retailer account fetch error", "error", err)
		return services.RetailersDetail{}, err
	}
	kyc, err := s.retailerRepo.Kyc(ctx, &retailer)
	if err != nil {
		logger.Error(ctx, "retailer kyc fetch error", "error", err)
		return services.RetailersDetail{}, err
	}
	cryptoWallets, err := s.retailerRepo.CryptoWallets(ctx, &retailer)
	if err != nil {
		logger.Error(ctx, "retailer cryptoWallets error", "error", err)
		return services.RetailersDetail{}, err
	}
	contact, err := s.retailerRepo.Contact(ctx, &retailer)
	if err != nil {
		logger.Error(ctx, "retailer contact error", "error", err)
		return services.RetailersDetail{}, err
	}
	shareholders, err := s.retailerRepo.Shareholders(ctx, &retailer)
	if err != nil {
		logger.Error(ctx, "retailer shareholders error", "error", err)
		return services.RetailersDetail{}, err
	}
	documents, err := s.retailerRepo.Documents(ctx, &retailer)
	if err != nil {
		logger.Error(ctx, "retailer documents error", "error", err)
		return services.RetailersDetail{}, err
	}
	address, err := s.retailerRepo.Address(ctx, &retailer)
	if err != nil {
		logger.Error(ctx, "retailer address error", "error", err)
		return services.RetailersDetail{}, err
	}
	txnFeeConfigs, err := s.retailerRepo.TxnFeeConfigs(ctx, &retailer)
	if err != nil {
		logger.Error(ctx, "retailer txnFeeConfigs error", "error", err)
		return services.RetailersDetail{}, err
	}
	return services.RetailersDetail{
		Account:       account,
		Retailer:      retailer,
		Kyc:           kyc,
		Address:       address,
		CryptoWallets: cryptoWallets,
		Contact:       contact,
		Shareholders:  shareholders,
		Documents:     documents,
		TxnFeeConfigs: txnFeeConfigs,
	}, nil
}

func (s *retailerServiceImpl) Create(ctx context.Context, r *models.Retailer) error {
	return s.retailerRepo.Create(ctx, r)
}

func (s *retailerServiceImpl) Update(ctx context.Context, r *models.Retailer) error {
	return s.retailerRepo.Save(ctx, r)
}

func (s *retailerServiceImpl) CreateRetailersAdmin(ctx context.Context, request services.RetailersCreateData) (services.RetailersDetail, error) {
	logger.Info(ctx, "create retailer", "request", request)
	response, err := DefaultRetailersAdminHelperServiceImpl().CreateRetailer(ctx, request)
	if err != nil {
		logger.Error(ctx, "create retailer error", "error", err)
		return response, err
	}
	return response, nil
}

func (s *retailerServiceImpl) CreateRetailersConfigurationsAdmin(ctx context.Context, request services.RetailersCreateData) (services.RetailersDetail, error) {
	logger.Info(ctx, "create retailer", "request", request)
	response, err := DefaultRetailersAdminHelperServiceImpl().CreateRetailersConfigurations(ctx, request)
	if err != nil {
		logger.Error(ctx, "create retailer error", "error", err)
		return response, err
	}
	return response, nil
}
