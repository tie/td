package downloader

import (
	"context"

	"github.com/gotd/td/tg"
)

// schema is simple interface for different download schemas.
type schema interface {
	Chunk(ctx context.Context, offset, limit int) (chunk, error)
	Hashes(ctx context.Context, offset int) ([]tg.FileHash, error)
}

type chunk struct {
	data []byte
	tag  tg.StorageFileTypeClass
}
