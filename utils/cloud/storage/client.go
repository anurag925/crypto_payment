package storage

import (
	"context"
	"io"
)

type Client interface {
	Upload(ctx context.Context, body io.ReadSeeker, m Metadata) error
}

type Metadata struct {
	Bucket      string
	Key         string
	ContentType string
	ACL         string
}
