// Code generated by gotdgen, DO NOT EDIT.

package tdapi

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdjson"
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
	_ = tdjson.Encoder{}
)

// EmojiCategory represents TL type `emojiCategory#d8681d22`.
type EmojiCategory struct {
	// Name of the category
	Name string
	// Unique identifier of the custom emoji, which represents icon of the category
	IconCustomEmojiID int64
	// List of emojis in the category
	Emojis []string
}

// EmojiCategoryTypeID is TL type id of EmojiCategory.
const EmojiCategoryTypeID = 0xd8681d22

// Ensuring interfaces in compile-time for EmojiCategory.
var (
	_ bin.Encoder     = &EmojiCategory{}
	_ bin.Decoder     = &EmojiCategory{}
	_ bin.BareEncoder = &EmojiCategory{}
	_ bin.BareDecoder = &EmojiCategory{}
)

func (e *EmojiCategory) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.Name == "") {
		return false
	}
	if !(e.IconCustomEmojiID == 0) {
		return false
	}
	if !(e.Emojis == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *EmojiCategory) String() string {
	if e == nil {
		return "EmojiCategory(nil)"
	}
	type Alias EmojiCategory
	return fmt.Sprintf("EmojiCategory%+v", Alias(*e))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*EmojiCategory) TypeID() uint32 {
	return EmojiCategoryTypeID
}

// TypeName returns name of type in TL schema.
func (*EmojiCategory) TypeName() string {
	return "emojiCategory"
}

// TypeInfo returns info about TL type.
func (e *EmojiCategory) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "emojiCategory",
		ID:   EmojiCategoryTypeID,
	}
	if e == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Name",
			SchemaName: "name",
		},
		{
			Name:       "IconCustomEmojiID",
			SchemaName: "icon_custom_emoji_id",
		},
		{
			Name:       "Emojis",
			SchemaName: "emojis",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (e *EmojiCategory) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode emojiCategory#d8681d22 as nil")
	}
	b.PutID(EmojiCategoryTypeID)
	return e.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (e *EmojiCategory) EncodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode emojiCategory#d8681d22 as nil")
	}
	b.PutString(e.Name)
	b.PutLong(e.IconCustomEmojiID)
	b.PutInt(len(e.Emojis))
	for _, v := range e.Emojis {
		b.PutString(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (e *EmojiCategory) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode emojiCategory#d8681d22 to nil")
	}
	if err := b.ConsumeID(EmojiCategoryTypeID); err != nil {
		return fmt.Errorf("unable to decode emojiCategory#d8681d22: %w", err)
	}
	return e.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (e *EmojiCategory) DecodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode emojiCategory#d8681d22 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode emojiCategory#d8681d22: field name: %w", err)
		}
		e.Name = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode emojiCategory#d8681d22: field icon_custom_emoji_id: %w", err)
		}
		e.IconCustomEmojiID = value
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode emojiCategory#d8681d22: field emojis: %w", err)
		}

		if headerLen > 0 {
			e.Emojis = make([]string, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode emojiCategory#d8681d22: field emojis: %w", err)
			}
			e.Emojis = append(e.Emojis, value)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (e *EmojiCategory) EncodeTDLibJSON(b tdjson.Encoder) error {
	if e == nil {
		return fmt.Errorf("can't encode emojiCategory#d8681d22 as nil")
	}
	b.ObjStart()
	b.PutID("emojiCategory")
	b.Comma()
	b.FieldStart("name")
	b.PutString(e.Name)
	b.Comma()
	b.FieldStart("icon_custom_emoji_id")
	b.PutLong(e.IconCustomEmojiID)
	b.Comma()
	b.FieldStart("emojis")
	b.ArrStart()
	for _, v := range e.Emojis {
		b.PutString(v)
		b.Comma()
	}
	b.StripComma()
	b.ArrEnd()
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (e *EmojiCategory) DecodeTDLibJSON(b tdjson.Decoder) error {
	if e == nil {
		return fmt.Errorf("can't decode emojiCategory#d8681d22 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("emojiCategory"); err != nil {
				return fmt.Errorf("unable to decode emojiCategory#d8681d22: %w", err)
			}
		case "name":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode emojiCategory#d8681d22: field name: %w", err)
			}
			e.Name = value
		case "icon_custom_emoji_id":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode emojiCategory#d8681d22: field icon_custom_emoji_id: %w", err)
			}
			e.IconCustomEmojiID = value
		case "emojis":
			if err := b.Arr(func(b tdjson.Decoder) error {
				value, err := b.String()
				if err != nil {
					return fmt.Errorf("unable to decode emojiCategory#d8681d22: field emojis: %w", err)
				}
				e.Emojis = append(e.Emojis, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode emojiCategory#d8681d22: field emojis: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetName returns value of Name field.
func (e *EmojiCategory) GetName() (value string) {
	if e == nil {
		return
	}
	return e.Name
}

// GetIconCustomEmojiID returns value of IconCustomEmojiID field.
func (e *EmojiCategory) GetIconCustomEmojiID() (value int64) {
	if e == nil {
		return
	}
	return e.IconCustomEmojiID
}

// GetEmojis returns value of Emojis field.
func (e *EmojiCategory) GetEmojis() (value []string) {
	if e == nil {
		return
	}
	return e.Emojis
}
