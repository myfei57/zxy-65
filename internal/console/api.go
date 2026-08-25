package console

import (
	"encoding/json"
	"io"
	"net/http"

	"portpower/internal/audit"
	"portpower/internal/quota"
	"portpower/internal/ship"
)

func (s *Server) handleBerthStatus(w http.ResponseWriter, r *http.Request) {
	status := s.runtime.berthReg.StatusAll()
	writeJSON(w, http.StatusOK, map[string]any{
		"occupied_count": s.runtime.berthReg.OccupiedCount(),
		"berths":         status,
	})
}

func (s *Server) handleShips(w http.ResponseWriter, r *http.Request) {
	fleet := s.runtime.fleet
	ids := fleet.IDs()
	moored := map[string]bool{}
	for _, id := range fleet.Moored() {
		moored[id] = true
	}
	profiles := map[string]ship.Profile{}
	for _, id := range ids {
		if p, ok := s.runtime.ships.Load(id); ok {
			profiles[id] = p
		}
	}
	writeJSON(w, http.StatusOK, map[string]any{
		"count":    fleet.Count(),
		"ship_ids": ids,
		"moored":   fleet.Moored(),
		"profiles": profiles,
	})
}

func (s *Server) handleAudit(w http.ResponseWriter, r *http.Request) {
	event := r.URL.Query().Get("event")
	var entries []audit.Entry
	if event == "" {
		entries = s.runtime.ledger.Entries()
	} else {
		entries = s.runtime.ledger.ByEvent(event)
	}
	writeJSON(w, http.StatusOK, map[string]any{
		"count":   len(entries),
		"events":  s.runtime.ledger.Events(),
		"entries": entries,
	})
}

func (s *Server) handleAuditLatest(w http.ResponseWriter, r *http.Request) {
	entry, ok := s.runtime.ledger.Latest()
	if !ok {
		writeJSON(w, http.StatusNotFound, map[string]string{"error": "no audit entries"})
		return
	}
	writeJSON(w, http.StatusOK, entry)
}

func (s *Server) handleQuotaList(w http.ResponseWriter, r *http.Request) {
	quotas := map[string]quota.Quota{}
	for _, key := range s.runtime.store.Keys("quota.") {
		shipID := key[len("quota."):]
		if q, ok := s.runtime.quotas.Get(shipID); ok {
			quotas[shipID] = q
		}
	}
	writeJSON(w, http.StatusOK, quotas)
}

func (s *Server) handleQuotaSet(w http.ResponseWriter, r *http.Request) {
	var req quota.Quota
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	if err := s.runtime.quotas.Set(req); err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{
		"set":       true,
		"remaining": s.runtime.quotas.Remaining(req.ShipID),
	})
}

func (s *Server) handleSnapshot(w http.ResponseWriter, r *http.Request) {
	raw, err := s.runtime.store.ExportJSON()
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(raw)
}

func (s *Server) handleRestore(w http.ResponseWriter, r *http.Request) {
	raw, err := io.ReadAll(r.Body)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	if err := s.runtime.store.ImportJSON(raw); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, map[string]bool{"restored": true})
}

func (s *Server) handleFreqGet(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]float64{"current": s.runtime.freq.Current()})
}

func (s *Server) handleFreqSet(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Value float64 `json:"value"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	if err := s.runtime.freq.Refresh(req.Value); err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	if err := s.runtime.freq.RecordSample(req.Value); err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, map[string]float64{"current": s.runtime.freq.Current()})
}

func (s *Server) handleFreqSamples(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]any{
		"samples": s.runtime.freq.Samples(),
		"current": s.runtime.freq.Current(),
		"within":  s.runtime.freq.Within(50, 0.5),
	})
}
