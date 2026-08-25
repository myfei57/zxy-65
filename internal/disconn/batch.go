package disconn

import "portpower/internal/meter"

type Connection struct {
	Breaker *Breaker
	Binder  *meter.Binder
	Berth   string
}

func BatchDisconnect(conns []Connection, finalValues map[string]float64) error {
	for _, conn := range conns {
		if err := conn.Binder.Unbind(conn.Berth); err != nil {
			return err
		}
		if err := conn.Breaker.Open(); err != nil {
			return err
		}
		if err := conn.Binder.RecordFinal(conn.Berth, finalValues[conn.Berth]); err != nil {
			return err
		}
	}
	return nil
}
