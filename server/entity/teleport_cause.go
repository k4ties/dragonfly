package entity

import "github.com/df-mc/dragonfly/server/world"

type TeleportCause interface {
	TeleportCause()
}

type TeleportCauseExternal struct {
	teleportCause
}

type TeleportCauseRespawn struct {
	teleportCause
}

type TeleportCauseProjectile struct {
	teleportCause

	Owner,
	Projectile world.Entity
}

type teleportCause struct{}

func (teleportCause) TeleportCause() {}
