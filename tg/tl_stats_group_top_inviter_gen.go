// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdp"
	"github.com/gotd/td/tgerr"
)

// No-op definition for keeping imports.
var (
	_ = bin.Buffer{}
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = multierr.AppendInto
	_ = sort.Ints
	_ = tdp.Format
	_ = tgerr.Error{}
)

// StatsGroupTopInviter represents TL type `statsGroupTopInviter#535f779d`.
// Information about an active supergroup inviter
//
// See https://core.telegram.org/constructor/statsGroupTopInviter for reference.
type StatsGroupTopInviter struct {
	// User ID
	UserID int64
	// Number of invitations for statistics¹ period in consideration
	//
	// Links:
	//  1) https://core.telegram.org/api/stats
	Invitations int
}

// StatsGroupTopInviterTypeID is TL type id of StatsGroupTopInviter.
const StatsGroupTopInviterTypeID = 0x535f779d

func (s *StatsGroupTopInviter) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.UserID == 0) {
		return false
	}
	if !(s.Invitations == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *StatsGroupTopInviter) String() string {
	if s == nil {
		return "StatsGroupTopInviter(nil)"
	}
	type Alias StatsGroupTopInviter
	return fmt.Sprintf("StatsGroupTopInviter%+v", Alias(*s))
}

// FillFrom fills StatsGroupTopInviter from given interface.
func (s *StatsGroupTopInviter) FillFrom(from interface {
	GetUserID() (value int64)
	GetInvitations() (value int)
}) {
	s.UserID = from.GetUserID()
	s.Invitations = from.GetInvitations()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StatsGroupTopInviter) TypeID() uint32 {
	return StatsGroupTopInviterTypeID
}

// TypeName returns name of type in TL schema.
func (*StatsGroupTopInviter) TypeName() string {
	return "statsGroupTopInviter"
}

// TypeInfo returns info about TL type.
func (s *StatsGroupTopInviter) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "statsGroupTopInviter",
		ID:   StatsGroupTopInviterTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "UserID",
			SchemaName: "user_id",
		},
		{
			Name:       "Invitations",
			SchemaName: "invitations",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *StatsGroupTopInviter) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode statsGroupTopInviter#535f779d as nil")
	}
	b.PutID(StatsGroupTopInviterTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *StatsGroupTopInviter) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode statsGroupTopInviter#535f779d as nil")
	}
	b.PutLong(s.UserID)
	b.PutInt(s.Invitations)
	return nil
}

// GetUserID returns value of UserID field.
func (s *StatsGroupTopInviter) GetUserID() (value int64) {
	return s.UserID
}

// GetInvitations returns value of Invitations field.
func (s *StatsGroupTopInviter) GetInvitations() (value int) {
	return s.Invitations
}

// Decode implements bin.Decoder.
func (s *StatsGroupTopInviter) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode statsGroupTopInviter#535f779d to nil")
	}
	if err := b.ConsumeID(StatsGroupTopInviterTypeID); err != nil {
		return fmt.Errorf("unable to decode statsGroupTopInviter#535f779d: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *StatsGroupTopInviter) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode statsGroupTopInviter#535f779d to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode statsGroupTopInviter#535f779d: field user_id: %w", err)
		}
		s.UserID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode statsGroupTopInviter#535f779d: field invitations: %w", err)
		}
		s.Invitations = value
	}
	return nil
}

// Ensuring interfaces in compile-time for StatsGroupTopInviter.
var (
	_ bin.Encoder     = &StatsGroupTopInviter{}
	_ bin.Decoder     = &StatsGroupTopInviter{}
	_ bin.BareEncoder = &StatsGroupTopInviter{}
	_ bin.BareDecoder = &StatsGroupTopInviter{}
)
