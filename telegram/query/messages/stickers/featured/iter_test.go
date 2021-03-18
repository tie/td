package featured

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/gotd/td/telegram/internal/rpcmock"
	"github.com/gotd/td/tg"
)

func generateStickers(count int) []tg.StickerSetCoveredClass {
	r := make([]tg.StickerSetCoveredClass, 0, count)

	for i := 0; i < count; i++ {
		r = append(r, &tg.StickerSetCovered{
			Set: tg.StickerSet{
				ID:         int64(i + 1),
				AccessHash: int64(i + 1),
			},
			Cover: &tg.Document{
				ID:            int64(i + 1),
				AccessHash:    int64(i + 1),
				FileReference: []uint8{uint8(i)},
			},
		})
	}

	return r
}

func result(r []tg.StickerSetCoveredClass, count int) tg.MessagesFeaturedStickersClass {
	return &tg.MessagesFeaturedStickers{
		Sets:  r,
		Count: count,
	}
}

func TestIterator(t *testing.T) {
	ctx := context.Background()
	mock := rpcmock.NewMock(t, require.New(t))
	limit := 10
	totalRecords := 3 * limit
	expected := generateStickers(totalRecords)

	mock.ExpectCall(&tg.MessagesGetOldFeaturedStickersRequest{
		Offset: 0,
		Limit:  limit,
	}).ThenResult(result(expected[0:limit], totalRecords))
	mock.ExpectCall(&tg.MessagesGetOldFeaturedStickersRequest{
		Offset: limit,
		Limit:  limit,
	}).ThenResult(result(expected[limit:2*limit], totalRecords))
	mock.ExpectCall(&tg.MessagesGetOldFeaturedStickersRequest{
		Offset: 2 * limit,
		Limit:  limit,
	}).ThenResult(result(expected[2*limit:3*limit], totalRecords))
	mock.ExpectCall(&tg.MessagesGetOldFeaturedStickersRequest{
		Offset: 3 * limit,
		Limit:  limit,
	}).ThenResult(result(expected[3*limit:], totalRecords))

	iter := NewQueryBuilder(mock).GetOldFeaturedStickers().BatchSize(10).Iter()
	i := 0
	for iter.Next(ctx) {
		mock.Equal(expected[i], iter.Value().Sticker)
		i++
	}
	mock.NoError(iter.Err())
	mock.Equal(totalRecords, i)

	total, err := iter.Total(ctx)
	mock.NoError(err)
	mock.Equal(totalRecords, total)

	mock.ExpectCall(&tg.MessagesGetOldFeaturedStickersRequest{
		Offset: 0,
		Limit:  1,
	}).ThenResult(result(expected[:0], totalRecords))
	total, err = iter.FetchTotal(ctx)
	mock.NoError(err)
	mock.Equal(totalRecords, total)
}
