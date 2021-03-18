package dialogs

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/gotd/td/telegram/internal/rpcmock"
	"github.com/gotd/td/tg"
)

func generateDialogs(count int) []tg.DialogClass {
	r := make([]tg.DialogClass, 0, count)

	for i := 0; i < count; i++ {
		r = append(r, &tg.Dialog{
			Peer: &tg.PeerChannel{ChannelID: i},
		})
	}

	return r
}

func result(r []tg.DialogClass, count int) tg.MessagesDialogsClass {
	msgs := make([]tg.MessageClass, 0, len(r))
	for i, dlg := range r {
		msgs = append(msgs, &tg.Message{
			ID:     i,
			PeerID: dlg.GetPeer(),
		})
	}

	chats := make([]tg.ChatClass, 0, len(r))
	for i, dlg := range r {
		id := dlg.GetPeer().(*tg.PeerChannel).ChannelID
		chats = append(chats, &tg.Channel{
			ID:         id,
			AccessHash: 10,
			Photo: &tg.ChatPhoto{
				PhotoSmall: tg.FileLocationToBeDeprecated{
					VolumeID: int64(i),
					LocalID:  i,
				},
				PhotoBig: tg.FileLocationToBeDeprecated{
					VolumeID: int64(i),
					LocalID:  i,
				},
			},
		})
	}

	return &tg.MessagesDialogsSlice{
		Dialogs:  r,
		Messages: msgs,
		Chats:    chats,
		Count:    count,
	}
}

func TestIterator(t *testing.T) {
	ctx := context.Background()
	mock := rpcmock.NewMock(t, require.New(t))
	limit := 10
	totalRows := 3 * limit
	expected := generateDialogs(totalRows)

	mock.Expect().ThenResult(result(expected[0:limit], totalRows))
	mock.Expect().ThenResult(result(expected[limit:2*limit], totalRows))
	mock.Expect().ThenResult(result(expected[2*limit:3*limit], totalRows))
	mock.Expect().ThenResult(result(expected[3*limit:], totalRows))

	iter := NewQueryBuilder(mock).GetDialogs().BatchSize(10).Iter()
	i := 0
	for iter.Next(ctx) {
		mock.Equal(expected[i].GetPeer(), iter.Value().Dialog.GetPeer())
		i++
	}
	mock.NoError(iter.Err())
	mock.Equal(totalRows, i)

	total, err := iter.Total(ctx)
	mock.NoError(err)
	mock.Equal(totalRows, total)

	mock.ExpectCall(&tg.MessagesGetDialogsRequest{
		OffsetPeer: &tg.InputPeerEmpty{},
		Limit:      1,
	}).ThenResult(result(expected[:0], totalRows))
	total, err = iter.FetchTotal(ctx)
	mock.NoError(err)
	mock.Equal(totalRows, total)
}
