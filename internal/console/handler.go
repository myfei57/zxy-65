package console

import (
	"encoding/json"
	"net/http"

	"portpower/internal/meter"
)

func writeJSON(w http.ResponseWriter, code int, value any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(value)
}

func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"identity": s.runtime.harbor.Identity()})
}

func (s *Server) handleConnect(w http.ResponseWriter, r *http.Request) {
	var req ConnectRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	result, err := s.runtime.Connect(req)
	if err != nil {
		writeJSON(w, http.StatusConflict, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, result)
}

func (s *Server) handleDisconnect(w http.ResponseWriter, r *http.Request) {
	var req DisconnectRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	result, err := s.runtime.Disconnect(req)
	if err != nil {
		writeJSON(w, http.StatusConflict, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, result)
}

func (s *Server) handleMeter(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Berth     string  `json:"berth"`
		Timestamp string  `json:"timestamp"`
		Value     float64 `json:"value"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	reading, err := meter.NewReading(req.Berth, req.Timestamp, req.Value)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	m := meter.NewMeter("meter-"+reading.Berth, reading.Berth)
	if !m.CanRecord() {
		writeJSON(w, http.StatusConflict, map[string]string{"error": "meter frozen"})
		return
	}
	if !s.runtime.store.Has("meter.binding." + reading.Berth) {
		writeJSON(w, http.StatusConflict, map[string]string{"error": "no meter binding for berth"})
		return
	}
	collector := meter.NewCollector(s.runtime.store)
	if err := collector.Record(reading.Timestamp, reading.Value); err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{"recorded": true, "reading": reading.String()})
}

func (s *Server) handleUsage(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]float64{"total": meter.TotalUsage(s.runtime.store)})
}

func (s *Server) handleBerths(w http.ResponseWriter, r *http.Request) {
	berths := map[string]string{}
	for key := range s.runtime.store.Values("berth.mapping.") {
		code := key[len("berth.mapping."):]
		if shipID, ok := s.runtime.berths.ShipFor(code); ok {
			berths[code] = shipID
		}
	}
	writeJSON(w, http.StatusOK, map[string]any{"count": s.runtime.store.Count("berth.mapping."), "berths": berths})
}

func (s *Server) handleRetrofit(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Ship      string  `json:"ship"`
		AddedLoad float64 `json:"added_load"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	if err := s.runtime.ships.Retrofit(req.Ship, req.AddedLoad); err != nil {
		writeJSON(w, http.StatusNotFound, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, map[string]bool{"retrofitted": true})
}
