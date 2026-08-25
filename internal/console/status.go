package console

import (
	"net/http"

	"portpower/internal/cable"
)

func (s *Server) handleStatus(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]any{
		"harbor":        s.runtime.harbor.Identity(),
		"breaker_closed": s.runtime.shore.Closed(),
		"breaker_tripped": s.runtime.shore.Tripped(),
		"breaker_state": s.runtime.shore.Describe(),
		"breaker_state_raw": s.runtime.shore.State(),
		"frequency":     s.runtime.freq.Current(),
		"temperature":   s.runtime.sensor.Read(),
		"cable_rating":  cable.Classify(s.runtime.sensor.Read(), s.runtime.cable.RatedTemperature()),
		"audit_count":   s.runtime.ledger.Count(),
		"berth_occupied": s.runtime.berthReg.OccupiedCount(),
		"fleet_count":   s.runtime.fleet.Count(),
		"cable_count":   s.runtime.cables.Count(),
		"phase_current": s.runtime.phase.Current(),
		"phase_history": s.runtime.phase.History(),
		"phase_reversed": s.runtime.phase.Reversed(),
		"freq_within":   s.runtime.freq.Within(50, 0.5),
		"store_keys":    s.runtime.store.Size(),
		"store_bytes":   s.runtime.store.TotalBytes(),
		"berth_catalog": s.runtime.catalog.Codes(),
		"catalog_count": s.runtime.catalog.Count(),
	})
}
