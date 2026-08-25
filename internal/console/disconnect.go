package console

import (
	"errors"

	"portpower/internal/disconn"
	"portpower/internal/meter"
	"portpower/internal/plug"
	"portpower/internal/ship"
)

type DisconnectRequest struct {
	Berth      string  `json:"berth"`
	Ship       string  `json:"ship"`
	FinalValue float64 `json:"final_value"`
	Connection string  `json:"connection"`
}

type DisconnectResult struct {
	FinalRecorded float64 `json:"final_recorded"`
	Released      bool    `json:"released"`
}

func (rt *Runtime) Disconnect(req DisconnectRequest) (DisconnectResult, error) {
	if req.Berth == "" {
		return DisconnectResult{}, errors.New("berth is required")
	}
	arm := ship.NewBreakerArm()
	arm.Arm()
	if err := rt.shore.CloseAfterArm(arm); err != nil {
		return DisconnectResult{}, err
	}
	connector := plug.NewPlug(req.Connection)
	if err := connector.Engage(req.Berth, rt.binder); err != nil {
		return DisconnectResult{}, err
	}
	connection := disconn.Connection{Breaker: rt.shore, Binder: rt.binder, Berth: req.Berth}
	if err := disconn.BatchDisconnect([]disconn.Connection{connection}, map[string]float64{req.Berth: req.FinalValue}); err != nil {
		return DisconnectResult{}, err
	}
	m := meter.NewMeter("meter-"+req.Berth, req.Berth)
	state := rt.berthReg.Add(req.Berth)
	if !state.Occupied() {
		if err := state.Assign(req.Ship); err != nil {
			return DisconnectResult{}, err
		}
	}
	if err := state.Depart(m); err != nil {
		return DisconnectResult{}, err
	}
	rt.fleet.Add(req.Ship, "vessel").Moor("")
	if err := connector.Release(); err != nil {
		return DisconnectResult{}, err
	}
	if err := rt.protect.OpenAndReset(rt.shore); err != nil {
		return DisconnectResult{}, err
	}
	final, _ := rt.binder.FinalValue(req.Berth)
	if err := rt.ledger.Append("disconnect", req.Berth); err != nil {
		return DisconnectResult{}, err
	}
	return DisconnectResult{FinalRecorded: final, Released: !connector.Engaged()}, nil
}
