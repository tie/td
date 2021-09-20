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

// ThemeSettings represents TL type `themeSettings#fa58b6d4`.
// Theme settings
//
// See https://core.telegram.org/constructor/themeSettings for reference.
type ThemeSettings struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// MessageColorsAnimated field of ThemeSettings.
	MessageColorsAnimated bool
	// Base theme
	BaseTheme BaseThemeClass
	// Accent color, RGB24 format
	AccentColor int
	// OutboxAccentColor field of ThemeSettings.
	//
	// Use SetOutboxAccentColor and GetOutboxAccentColor helpers.
	OutboxAccentColor int
	// MessageColors field of ThemeSettings.
	//
	// Use SetMessageColors and GetMessageColors helpers.
	MessageColors []int
	// Wallpaper
	//
	// Use SetWallpaper and GetWallpaper helpers.
	Wallpaper WallPaperClass
}

// ThemeSettingsTypeID is TL type id of ThemeSettings.
const ThemeSettingsTypeID = 0xfa58b6d4

func (t *ThemeSettings) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.Flags.Zero()) {
		return false
	}
	if !(t.MessageColorsAnimated == false) {
		return false
	}
	if !(t.BaseTheme == nil) {
		return false
	}
	if !(t.AccentColor == 0) {
		return false
	}
	if !(t.OutboxAccentColor == 0) {
		return false
	}
	if !(t.MessageColors == nil) {
		return false
	}
	if !(t.Wallpaper == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *ThemeSettings) String() string {
	if t == nil {
		return "ThemeSettings(nil)"
	}
	type Alias ThemeSettings
	return fmt.Sprintf("ThemeSettings%+v", Alias(*t))
}

// FillFrom fills ThemeSettings from given interface.
func (t *ThemeSettings) FillFrom(from interface {
	GetMessageColorsAnimated() (value bool)
	GetBaseTheme() (value BaseThemeClass)
	GetAccentColor() (value int)
	GetOutboxAccentColor() (value int, ok bool)
	GetMessageColors() (value []int, ok bool)
	GetWallpaper() (value WallPaperClass, ok bool)
}) {
	t.MessageColorsAnimated = from.GetMessageColorsAnimated()
	t.BaseTheme = from.GetBaseTheme()
	t.AccentColor = from.GetAccentColor()
	if val, ok := from.GetOutboxAccentColor(); ok {
		t.OutboxAccentColor = val
	}

	if val, ok := from.GetMessageColors(); ok {
		t.MessageColors = val
	}

	if val, ok := from.GetWallpaper(); ok {
		t.Wallpaper = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ThemeSettings) TypeID() uint32 {
	return ThemeSettingsTypeID
}

// TypeName returns name of type in TL schema.
func (*ThemeSettings) TypeName() string {
	return "themeSettings"
}

// TypeInfo returns info about TL type.
func (t *ThemeSettings) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "themeSettings",
		ID:   ThemeSettingsTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "MessageColorsAnimated",
			SchemaName: "message_colors_animated",
			Null:       !t.Flags.Has(2),
		},
		{
			Name:       "BaseTheme",
			SchemaName: "base_theme",
		},
		{
			Name:       "AccentColor",
			SchemaName: "accent_color",
		},
		{
			Name:       "OutboxAccentColor",
			SchemaName: "outbox_accent_color",
			Null:       !t.Flags.Has(3),
		},
		{
			Name:       "MessageColors",
			SchemaName: "message_colors",
			Null:       !t.Flags.Has(0),
		},
		{
			Name:       "Wallpaper",
			SchemaName: "wallpaper",
			Null:       !t.Flags.Has(1),
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (t *ThemeSettings) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode themeSettings#fa58b6d4 as nil")
	}
	b.PutID(ThemeSettingsTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *ThemeSettings) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode themeSettings#fa58b6d4 as nil")
	}
	if !(t.MessageColorsAnimated == false) {
		t.Flags.Set(2)
	}
	if !(t.OutboxAccentColor == 0) {
		t.Flags.Set(3)
	}
	if !(t.MessageColors == nil) {
		t.Flags.Set(0)
	}
	if !(t.Wallpaper == nil) {
		t.Flags.Set(1)
	}
	if err := t.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode themeSettings#fa58b6d4: field flags: %w", err)
	}
	if t.BaseTheme == nil {
		return fmt.Errorf("unable to encode themeSettings#fa58b6d4: field base_theme is nil")
	}
	if err := t.BaseTheme.Encode(b); err != nil {
		return fmt.Errorf("unable to encode themeSettings#fa58b6d4: field base_theme: %w", err)
	}
	b.PutInt(t.AccentColor)
	if t.Flags.Has(3) {
		b.PutInt(t.OutboxAccentColor)
	}
	if t.Flags.Has(0) {
		b.PutVectorHeader(len(t.MessageColors))
		for _, v := range t.MessageColors {
			b.PutInt(v)
		}
	}
	if t.Flags.Has(1) {
		if t.Wallpaper == nil {
			return fmt.Errorf("unable to encode themeSettings#fa58b6d4: field wallpaper is nil")
		}
		if err := t.Wallpaper.Encode(b); err != nil {
			return fmt.Errorf("unable to encode themeSettings#fa58b6d4: field wallpaper: %w", err)
		}
	}
	return nil
}

// SetMessageColorsAnimated sets value of MessageColorsAnimated conditional field.
func (t *ThemeSettings) SetMessageColorsAnimated(value bool) {
	if value {
		t.Flags.Set(2)
		t.MessageColorsAnimated = true
	} else {
		t.Flags.Unset(2)
		t.MessageColorsAnimated = false
	}
}

// GetMessageColorsAnimated returns value of MessageColorsAnimated conditional field.
func (t *ThemeSettings) GetMessageColorsAnimated() (value bool) {
	return t.Flags.Has(2)
}

// GetBaseTheme returns value of BaseTheme field.
func (t *ThemeSettings) GetBaseTheme() (value BaseThemeClass) {
	return t.BaseTheme
}

// GetAccentColor returns value of AccentColor field.
func (t *ThemeSettings) GetAccentColor() (value int) {
	return t.AccentColor
}

// SetOutboxAccentColor sets value of OutboxAccentColor conditional field.
func (t *ThemeSettings) SetOutboxAccentColor(value int) {
	t.Flags.Set(3)
	t.OutboxAccentColor = value
}

// GetOutboxAccentColor returns value of OutboxAccentColor conditional field and
// boolean which is true if field was set.
func (t *ThemeSettings) GetOutboxAccentColor() (value int, ok bool) {
	if !t.Flags.Has(3) {
		return value, false
	}
	return t.OutboxAccentColor, true
}

// SetMessageColors sets value of MessageColors conditional field.
func (t *ThemeSettings) SetMessageColors(value []int) {
	t.Flags.Set(0)
	t.MessageColors = value
}

// GetMessageColors returns value of MessageColors conditional field and
// boolean which is true if field was set.
func (t *ThemeSettings) GetMessageColors() (value []int, ok bool) {
	if !t.Flags.Has(0) {
		return value, false
	}
	return t.MessageColors, true
}

// SetWallpaper sets value of Wallpaper conditional field.
func (t *ThemeSettings) SetWallpaper(value WallPaperClass) {
	t.Flags.Set(1)
	t.Wallpaper = value
}

// GetWallpaper returns value of Wallpaper conditional field and
// boolean which is true if field was set.
func (t *ThemeSettings) GetWallpaper() (value WallPaperClass, ok bool) {
	if !t.Flags.Has(1) {
		return value, false
	}
	return t.Wallpaper, true
}

// Decode implements bin.Decoder.
func (t *ThemeSettings) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode themeSettings#fa58b6d4 to nil")
	}
	if err := b.ConsumeID(ThemeSettingsTypeID); err != nil {
		return fmt.Errorf("unable to decode themeSettings#fa58b6d4: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *ThemeSettings) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode themeSettings#fa58b6d4 to nil")
	}
	{
		if err := t.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode themeSettings#fa58b6d4: field flags: %w", err)
		}
	}
	t.MessageColorsAnimated = t.Flags.Has(2)
	{
		value, err := DecodeBaseTheme(b)
		if err != nil {
			return fmt.Errorf("unable to decode themeSettings#fa58b6d4: field base_theme: %w", err)
		}
		t.BaseTheme = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode themeSettings#fa58b6d4: field accent_color: %w", err)
		}
		t.AccentColor = value
	}
	if t.Flags.Has(3) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode themeSettings#fa58b6d4: field outbox_accent_color: %w", err)
		}
		t.OutboxAccentColor = value
	}
	if t.Flags.Has(0) {
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode themeSettings#fa58b6d4: field message_colors: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int()
			if err != nil {
				return fmt.Errorf("unable to decode themeSettings#fa58b6d4: field message_colors: %w", err)
			}
			t.MessageColors = append(t.MessageColors, value)
		}
	}
	if t.Flags.Has(1) {
		value, err := DecodeWallPaper(b)
		if err != nil {
			return fmt.Errorf("unable to decode themeSettings#fa58b6d4: field wallpaper: %w", err)
		}
		t.Wallpaper = value
	}
	return nil
}

// Ensuring interfaces in compile-time for ThemeSettings.
var (
	_ bin.Encoder     = &ThemeSettings{}
	_ bin.Decoder     = &ThemeSettings{}
	_ bin.BareEncoder = &ThemeSettings{}
	_ bin.BareDecoder = &ThemeSettings{}
)
