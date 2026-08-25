package disconn

import "errors"

var errNotClosed = errors.New("breaker is not closed")

type Breaker struct {
	ID      string
	closed  bool
	tripped bool
}

func NewBreaker(id string) *Breaker {
	return &Breaker{ID: id}
}

func (b *Breaker) Closed() bool {
	return b.closed
}

func (b *Breaker) Tripped() bool {
	return b.tripped
}

func (b *Breaker) close() {
	b.closed = true
}

func (b *Breaker) open() {
	b.closed = false
}
