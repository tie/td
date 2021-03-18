package downloader

import (
	"context"

	"golang.org/x/xerrors"

	"github.com/gotd/td/tg"
)

var errHashesNotSupported = xerrors.New("this schema does not support hashes fetch")

// web is a web file download schema.
// See https://core.telegram.org/api/files#downloading-webfiles.
type web struct {
	rpc tg.Invoker

	location tg.InputWebFileLocationClass
}

var _ schema = web{}

func (w web) Chunk(ctx context.Context, offset, limit int) (chunk, error) {
	file, err := tg.UploadGetWebFile(ctx, w.rpc, &tg.UploadGetWebFileRequest{
		Location: w.location,
		Offset:   offset,
		Limit:    limit,
	})
	if err != nil {
		return chunk{}, err
	}

	return chunk{data: file.Bytes, tag: file.FileType}, nil
}

func (w web) Hashes(ctx context.Context, offset int) ([]tg.FileHash, error) {
	return nil, errHashesNotSupported
}
