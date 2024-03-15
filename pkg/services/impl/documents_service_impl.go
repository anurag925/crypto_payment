package impl

import (
	"context"
	"fmt"
	"os"

	"github.com/anurag925/crypto_payment/app"
	"github.com/anurag925/crypto_payment/pkg/models"
	"github.com/anurag925/crypto_payment/pkg/repositories"
	"github.com/anurag925/crypto_payment/pkg/repositories/postgresql"
	"github.com/anurag925/crypto_payment/pkg/services"
	"github.com/anurag925/crypto_payment/utils/cloud/session"
	"github.com/anurag925/crypto_payment/utils/cloud/storage"
	"github.com/anurag925/crypto_payment/utils/logger"
)

type documentServiceImpl struct {
	documentRepo repositories.DocumentRepository
}

var _ services.DocumentService = (*documentServiceImpl)(nil)

func NewDocumentServiceImpl(documentRepo repositories.DocumentRepository) *documentServiceImpl {
	return &documentServiceImpl{documentRepo: documentRepo}
}

func DefaultDocumentServiceImpl() *documentServiceImpl {
	return NewDocumentServiceImpl(postgresql.DefaultDocumentRepositoryImpl())
}

func (s *documentServiceImpl) Create(ctx context.Context, d *models.Document) error {
	return s.documentRepo.Create(ctx, d)
}

func (s *documentServiceImpl) Upload(ctx context.Context, path, access, contentType string) (string, error) {
	awsSession, err := session.NewAwsSession(
		app.Config().AccessKeyID, app.Config().SecretAccessKey, app.Config().Region)
	if err != nil {
		return "", err
	}
	storageClient := storage.NewAwsS3(awsSession)
	key := fmt.Sprintf("crypto_payment/kyc/documents/%s", path)
	logger.Info(ctx, "upload document", "key", key)
	fileMetadata := storage.Metadata{
		Bucket:      os.Getenv("DOCUMENT_S3_BUCKET"),
		Key:         key,
		ContentType: contentType,
	}
	return storageClient.PresignedPutUrl(ctx, fileMetadata)
}
