package session

import (
	"github.com/df-mc/dragonfly/server/event"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

// packetHandler represents a type that is able to handle a specific type of incoming packets from the client.
type packetHandler interface {
	// Handle handles an incoming packet from the client. The session of the client is passed to the packetHandler.
	// Handle returns an error if the packet was in any way invalid.
	Handle(p packet.Packet, s *Session, tx *world.Tx, c Controllable) error
}

// Context ...
type Context = event.Context[*Session]

// Handler is packet handler to the Session.
type Handler interface {
	// HandleClientPacket handles all client packets.
	HandleClientPacket(ctx *Context, pk packet.Packet)
	// HandleServerPacket handles all server packets.
	HandleServerPacket(ctx *Context, pk packet.Packet)
}

// nopHandler is no-operation implementation of Handler.
type nopHandler struct{}

func (nopHandler) HandleClientPacket(*Context, packet.Packet) {}
func (nopHandler) HandleServerPacket(*Context, packet.Packet) {}
