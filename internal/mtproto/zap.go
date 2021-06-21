package mtproto

import (
	"fmt"

	"code.cloudfoundry.org/bytefmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/gotd/td/bin"
)

type logType struct {
	ID   uint32
	Name string
}

func (l logType) MarshalLogObject(e zapcore.ObjectEncoder) error {
	typeIDStr := fmt.Sprintf("0x%x", l.ID)
	e.AddString("type_id", typeIDStr)
	if l.Name != "" {
		e.AddString("type_name", l.Name)
	}
	return nil
}

type logSize struct {
	Size uint64
}

func (l logSize) MarshalLogObject(e zapcore.ObjectEncoder) error {
	e.AddUint64("size_bytes", l.Size)
	// Don’t print small sizes (like 88B). That’s confusing for humans.
	if l.Size >= bytefmt.KILOBYTE {
		e.AddString("size_human", bytefmt.ByteSize(l.Size)+"iB") // IEC
	}
	return nil
}

func (c *Conn) logWithBuffer(b *bin.Buffer) *zap.Logger {
	return c.logWithType(b).With(zap.Inline(logSize{
		Size: uint64(b.Len()),
	}))
}

func (c *Conn) logWithType(b *bin.Buffer) *zap.Logger {
	id, err := b.PeekID()
	if err != nil {
		// Type info not available.
		return c.log
	}

	return c.logWithTypeID(id)
}

func (c *Conn) logWithTypeID(id uint32) *zap.Logger {
	return c.log.With(zap.Inline(logType{
		ID:   id,
		Name: c.types.Get(id),
	}))
}
