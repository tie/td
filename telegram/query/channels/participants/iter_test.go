package participants

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/gotd/td/telegram/internal/rpcmock"
	"github.com/gotd/td/tg"
)

func generateParticipants(count int) []tg.ChannelParticipantClass {
	r := make([]tg.ChannelParticipantClass, 0, count)

	for i := 0; i < count; i++ {
		r = append(r, &tg.ChannelParticipant{
			UserID: i,
			Date:   i,
		})
	}

	return r
}

func result(r []tg.ChannelParticipantClass, count int) tg.ChannelsChannelParticipantsClass {
	return &tg.ChannelsChannelParticipants{
		Participants: r,
		Count:        count,
	}
}

func TestIterator(t *testing.T) {
	ctx := context.Background()
	mock := rpcmock.NewMock(t, require.New(t))
	limit := 10
	totalRecords := 3 * limit
	expected := generateParticipants(totalRecords)
	ch := &tg.InputChannel{
		ChannelID:  10,
		AccessHash: 10,
	}

	mock.ExpectCall(&tg.ChannelsGetParticipantsRequest{
		Channel: ch,
		Filter:  &tg.ChannelParticipantsRecent{},
		Offset:  0,
		Limit:   limit,
	}).ThenResult(result(expected[0:limit], totalRecords))
	mock.ExpectCall(&tg.ChannelsGetParticipantsRequest{
		Channel: ch,
		Filter:  &tg.ChannelParticipantsRecent{},
		Offset:  limit,
		Limit:   limit,
	}).ThenResult(result(expected[limit:2*limit], totalRecords))
	mock.ExpectCall(&tg.ChannelsGetParticipantsRequest{
		Channel: ch,
		Filter:  &tg.ChannelParticipantsRecent{},
		Offset:  2 * limit,
		Limit:   limit,
	}).ThenResult(result(expected[2*limit:3*limit], totalRecords))
	mock.ExpectCall(&tg.ChannelsGetParticipantsRequest{
		Channel: ch,
		Filter:  &tg.ChannelParticipantsRecent{},
		Offset:  3 * limit,
		Limit:   limit,
	}).ThenResult(result(expected[3*limit:], totalRecords))

	iter := NewQueryBuilder(mock).GetParticipants(ch).BatchSize(10).Iter()
	i := 0
	for iter.Next(ctx) {
		mock.Equal(expected[i], iter.Value().Participant)
		i++
	}
	mock.NoError(iter.Err())
	mock.Equal(totalRecords, i)

	total, err := iter.Total(ctx)
	mock.NoError(err)
	mock.Equal(totalRecords, total)

	mock.ExpectCall(&tg.ChannelsGetParticipantsRequest{
		Channel: ch,
		Filter:  &tg.ChannelParticipantsRecent{},
		Offset:  0,
		Limit:   1,
	}).ThenResult(result(expected[:0], totalRecords))
	total, err = iter.FetchTotal(ctx)
	mock.NoError(err)
	mock.Equal(totalRecords, total)
}
