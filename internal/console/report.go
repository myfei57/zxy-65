package console

import (
	"encoding/json"
	"net/http"

	"portpower/internal/meter"
)

func (s *Server) handleCables(w http.ResponseWriter, r *http.Request) {
	ids := s.runtime.cables.IDs()
	cables := map[string]float64{}
	for _, id := range ids {
		if c, ok := s.runtime.cables.Get(id); ok {
			cables[id] = c.RatedTemperature()
		}
	}
	selected, hasSelection := s.runtime.cables.Select(90)
	writeJSON(w, http.StatusOK, map[string]any{
		"count":   s.runtime.cables.Count(),
		"ids":     ids,
		"cables":  cables,
		"selected": map[string]any{"id": selected, "found": hasSelection},
	})
}

func (s *Server) handleAllocate(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Ship string `json:"ship"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	code, ok := s.runtime.berthReg.Allocate(req.Ship)
	if !ok {
		writeJSON(w, http.StatusConflict, map[string]string{"error": "no berth available"})
		return
	}
	s.runtime.fleet.Add(req.Ship, "vessel").Moor(code)
	writeJSON(w, http.StatusOK, map[string]string{"berth": code})
}

func (s *Server) handleVacate(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Berth string `json:"berth"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	shipID, ok := s.runtime.berthReg.ShipAt(req.Berth)
	if ok {
		s.runtime.fleet.Add(shipID, "vessel").Moor("")
	}
	writeJSON(w, http.StatusOK, map[string]bool{"vacated": s.runtime.berthReg.Vacate(req.Berth)})
}

func (s *Server) handleBilling(w http.ResponseWriter, r *http.Request) {
	rate := s.runtime.tariff.Rate
	total := meter.TotalUsage(s.runtime.store)
	writeJSON(w, http.StatusOK, map[string]float64{
		"total_usage": total,
		"rate":        rate,
		"charge":      s.runtime.tariff.Charge(total),
		"bill":        meter.Bill(s.runtime.store, rate),
	})
}

func (s *Server) handleQuotaReserve(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Ship   string  `json:"ship"`
		Demand float64 `json:"demand"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, map[string]bool{"reserved": s.runtime.quotas.Reserve(req.Ship, req.Demand)})
}

func (s *Server) handleQuotaRelease(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Ship   string  `json:"ship"`
		Amount float64 `json:"amount"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, map[string]bool{"released": s.runtime.quotas.Release(req.Ship, req.Amount)})
}

func (s *Server) handleAuditSummary(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, s.runtime.ledger.Summary())
}

func (s *Server) handleBreakerTrip(w http.ResponseWriter, r *http.Request) {
	s.runtime.shore.Trip()
	writeJSON(w, http.StatusOK, map[string]bool{"tripped": s.runtime.shore.Tripped()})
}

func (s *Server) handleBreakerReset(w http.ResponseWriter, r *http.Request) {
	err := s.runtime.shore.Reset()
	if err != nil {
		writeJSON(w, http.StatusConflict, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, map[string]bool{"closed": s.runtime.shore.Closed()})
}
