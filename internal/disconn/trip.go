package disconn

import "errors"

func (b *Breaker) Trip() {
	b.tripped = true
	b.open()
}

func (b *Breaker) Reset() error {
	if b.closed {
		return errors.New("breaker already closed")
	}
	b.tripped = false
	b.close()
	return nil
}
