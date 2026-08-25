package console

import (
	"errors"

	"github.com/google/uuid"

	"portpower/internal/meter"
	"portpower/internal/phase"
	"portpower/internal/plug"
	"portpower/internal/quota"
	"portpower/internal/ship"
)

type ConnectRequest struct {
	Berth       string  `json:"berth"`
	Ship        string  `json:"ship"`
	Phase       string  `json:"phase"`
	Frequency   float64 `json:"frequency"`
	Load        float64 `json:"load"`
	Temperature float64 `json:"temperature"`
}

type ConnectResult struct {
	ConnectionID string `json:"connection_id"`
	Engaged      bool   `json:"engaged"`
	PlugState    string `json:"plug_state"`
	Demand       float64 `json:"demand"`
	Headroom     float64 `json:"headroom"`
}

func (rt *Runtime) Connect(req ConnectRequest) (ConnectResult, error) {
	if !validBerthCode(req.Berth) || !rt.catalog.Has(req.Berth) {
		return ConnectResult{}, errors.New("invalid or unknown berth code")
	}
	if !validShipID(req.Ship) {
		return ConnectResult{}, errors.New("invalid ship id")
	}
	if !validFrequency(req.Frequency) {
		return ConnectResult{}, errors.New("invalid frequency")
	}
	if !validLoad(req.Load) {
		return ConnectResult{}, errors.New("invalid load")
	}
	if !validTemperature(req.Temperature) {
		return ConnectResult{}, errors.New("invalid temperature")
	}
	connID := uuid.NewString()
	berthState := rt.berthReg.Add(req.Berth)
	if err := berthState.Assign(req.Ship); err != nil {
		return ConnectResult{}, err
	}
	vessel := rt.fleet.Add(req.Ship, "vessel")
	vessel.Moor(req.Berth)
	profile := ship.Profile{ShipID: req.Ship, MaxLoad: req.Load, PhaseSeq: req.Phase, Frequency: req.Frequency}
	if !profile.Valid() || !profile.PhaseValid() {
		return ConnectResult{}, errors.New("invalid ship power profile")
	}
	if err := rt.ships.Save(profile); err != nil {
		return ConnectResult{}, err
	}
	if err := rt.history.Append(profile); err != nil {
		return ConnectResult{}, err
	}
	if err := rt.berths.Bind(req.Berth, req.Ship); err != nil {
		return ConnectResult{}, err
	}
	connector := plug.NewPlug(connID)
	if err := rt.binder.Bind(meter.Binding{Berth: req.Berth, Ship: req.Ship}); err != nil {
		return ConnectResult{}, err
	}
	if err := connector.Engage(req.Berth, rt.binder); err != nil {
		return ConnectResult{}, err
	}
	sequence, ok := phase.ParseSequence(req.Phase)
	if !ok {
		return ConnectResult{}, errors.New("invalid phase sequence")
	}
	verifier := phase.NewVerifier(sequence)
	rt.phase.Set(sequence)
	if err := rt.shore.CloseAfterPhase(verifier); err != nil {
		return ConnectResult{}, err
	}
	if err := rt.freq.Refresh(req.Frequency); err != nil {
		return ConnectResult{}, err
	}
	if err := rt.shore.ParallelClose(rt.freq, req.Frequency, 0.5); err != nil {
		return ConnectResult{}, err
	}
	rt.sensor.Set(req.Temperature)
	rt.monitor.Observe(req.Temperature)
	_ = rt.protect.Verdict(rt.sensor.Read())
	_ = rt.quotas.Set(quota.Quota{ShipID: req.Ship, Limit: req.Load})
	_ = rt.quotas.Allow(req.Ship, req.Load)
	if err := rt.ledger.Append("connect", connID); err != nil {
		return ConnectResult{}, err
	}
	return ConnectResult{
		ConnectionID: connID,
		Engaged:      connector.Engaged(),
		PlugState:    string(connector.State()),
		Demand:       rt.ships.Demand(req.Ship),
		Headroom:     rt.ships.Headroom(req.Ship),
	}, nil
}
