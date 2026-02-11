package entity

import "github.com/df-mc/dragonfly/server/world"

// ProjectileIntersect used to allow ProjectileBehaviour to intersect living
// entity.
type ProjectileIntersect func(ent Living, owner *world.EntityHandle) bool
