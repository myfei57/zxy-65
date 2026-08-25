package disconn

func (b *Breaker) Open() error {
	if !b.closed {
		return errNotClosed
	}
	b.open()
	return nil
}
