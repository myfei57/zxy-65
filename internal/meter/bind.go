package meter

import (
	"encoding/json"

	"portpower/internal/store"
)

type Binding struct {
	Berth string `json:"berth"`
	Ship  string `json:"ship"`
}

type Binder struct {
	store   *store.Store
	current map[string]Binding
}

func NewBinder(st *store.Store) *Binder {
	return &Binder{store: st, current: map[string]Binding{}}
}

func bindingKey(berth string) string {
	return "meter.binding." + berth
}

func (b *Binder) Bind(bind Binding) error {
	raw, err := json.Marshal(bind)
	if err != nil {
		return err
	}
	if err := b.store.Put(bindingKey(bind.Berth), raw); err != nil {
		return err
	}
	b.current[bind.Berth] = bind
	return nil
}

func (b *Binder) Load(berth string) (Binding, bool) {
	raw, ok := b.store.Get(bindingKey(berth))
	if !ok {
		return Binding{}, false
	}
	var bind Binding
	if err := json.Unmarshal(raw, &bind); err != nil {
		return Binding{}, false
	}
	return bind, true
}

func (b *Binder) Current(berth string) (Binding, bool) {
	bind, ok := b.current[berth]
	return bind, ok
}
