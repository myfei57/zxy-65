package disconn

import "portpower/internal/ship"

type LimitEnforcer struct {
	registry *ship.Registry
	cache    map[string]ship.Profile
}

func NewLimitEnforcer(reg *ship.Registry) *LimitEnforcer {
	return &LimitEnforcer{registry: reg, cache: map[string]ship.Profile{}}
}

func (e *LimitEnforcer) Allow(shipID string, demand float64) bool {
	if p, ok := e.cache[shipID]; ok {
		return demand <= p.MaxLoad
	}
	p, ok := e.registry.Load(shipID)
	if !ok {
		return false
	}
	e.cache[shipID] = p
	return demand <= p.MaxLoad
}
