package meter

func (m *Meter) Freeze() {
	m.frozen = true
	if m.store != nil {
		_ = m.store.Put("meter.frozen."+m.Berth, []byte("1"))
	}
}
