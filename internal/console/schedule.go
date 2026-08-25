package console

import (
	"encoding/json"
	"net/http"

	"portpower/internal/berth"
	"portpower/internal/meter"
	"portpower/internal/phase"
	"portpower/internal/ship"
)

func (s *Server) handleReservationBook(w http.ResponseWriter, r *http.Request) {
	var req berth.Reservation
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	if err := s.runtime.sched.Book(req); err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, map[string]bool{"booked": true})
}

func (s *Server) handleReservationList(w http.ResponseWriter, r *http.Request) {
	if berthCode := r.URL.Query().Get("berth"); berthCode != "" {
		reservation, ok := s.runtime.sched.Get(berthCode)
		writeJSON(w, http.StatusOK, map[string]any{"reservation": reservation, "found": ok})
		return
	}
	writeJSON(w, http.StatusOK, s.runtime.sched.List())
}

func (s *Server) handleReservationCancel(w http.ResponseWriter, r *http.Request) {
	berthCode := r.URL.Query().Get("berth")
	if err := s.runtime.sched.Cancel(berthCode); err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, map[string]bool{"cancelled": true})
}

func (s *Server) handleShipHistory(w http.ResponseWriter, r *http.Request) {
	shipID := r.URL.Query().Get("ship")
	records := s.runtime.history.List(shipID)
	latest, ok := s.runtime.history.Latest(shipID)
	writeJSON(w, http.StatusOK, map[string]any{
		"records": records,
		"latest":  latest,
		"found":   ok,
	})
}

func (s *Server) handleMeterPeriod(w http.ResponseWriter, r *http.Request) {
	start := r.URL.Query().Get("start")
	end := r.URL.Query().Get("end")
	writeJSON(w, http.StatusOK, meter.Aggregate(s.runtime.store, start, end))
}

func (s *Server) handleCableMonitor(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]any{
		"peak":    s.runtime.monitor.Peak(),
		"count":   s.runtime.monitor.Count(),
		"average": s.runtime.monitor.Average(),
	})
}

func (s *Server) handleBreakerInterlock(w http.ResponseWriter, r *http.Request) {
	arm := ship.NewBreakerArm()
	arm.Arm()
	verifier := phase.NewVerifier(s.runtime.phase.Current())
	writeJSON(w, http.StatusOK, map[string]bool{"ready": s.runtime.shore.InterlockReady(arm, verifier)})
}
