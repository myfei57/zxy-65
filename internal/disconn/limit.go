package disconn

import "portpower/internal/ship"

type LimitEnforcer struct {
	registry *ship.Registry
}

func NewLimitEnforcer(reg *ship.Registry) *LimitEnforcer {
	return &LimitEnforcer{registry: reg}
}

func (e *LimitEnforcer) Allow(shipID string, demand float64) bool {
	p, ok := e.registry.Load(shipID)
	if !ok {
		return false
	}
	return demand <= p.MaxLoad
}
