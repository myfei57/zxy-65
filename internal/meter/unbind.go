package meter

import (
	"encoding/json"
	"errors"
)

func (b *Binder) Unbind(berth string) error {
	return b.store.Delete(bindingKey(berth))
}

func (b *Binder) RecordFinal(berth string, value float64) error {
	if _, ok := b.Load(berth); !ok {
		return errors.New("binding missing")
	}
	raw, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return b.store.Put("meter.final."+berth, raw)
}

func (b *Binder) FinalValue(berth string) (float64, bool) {
	raw, ok := b.store.Get("meter.final." + berth)
	if !ok {
		return 0, false
	}
	var value float64
	if err := json.Unmarshal(raw, &value); err != nil {
		return 0, false
	}
	return value, true
}
