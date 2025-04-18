package session

import (
	"github.com/df-mc/dragonfly/server/event"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
	"math"
)

// connWrapper ...
type connWrapper struct {
	Conn
	session *Session
	get     func() Handler
}

// ReadPacket ...
func (c *connWrapper) ReadPacket() (packet.Packet, error) {
	pkt, err := c.Conn.ReadPacket()
	if err != nil {
		return pkt, err
	}

	ctx := event.C(c.session)
	if c.get().HandleClientPacket(ctx, pkt); ctx.Cancelled() {
		return zeroPacket{}, nil
	}

	return pkt, nil
}

// WritePacket ...
func (c *connWrapper) WritePacket(pk packet.Packet) error {
	ctx := event.C(c.session)
	if c.get().HandleServerPacket(ctx, pk); ctx.Cancelled() {
		return nil
	}
	return c.Conn.WritePacket(pk)
}

// zeroPacket is zero implementation of gophertunnel Packet interface, to avoid using nil.
type zeroPacket struct{}

func (zeroPacket) ID() uint32          { return math.MaxUint32 }
func (zeroPacket) Marshal(protocol.IO) {}
