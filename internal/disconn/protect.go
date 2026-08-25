package disconn

import "portpower/internal/cable"

type Protector struct {
	threshold float64
	filter    *cable.Filter
	tripped   bool
}

func NewProtector(threshold float64, filter *cable.Filter) *Protector {
	return &Protector{threshold: threshold, filter: filter}
}

func (p *Protector) ResetFilter() {
	p.filter.Reset()
}

func (p *Protector) Verdict(sample float64) bool {
	filtered := p.filter.Add(sample)
	if filtered > p.threshold {
		p.tripped = true
		return true
	}
	return false
}

func (p *Protector) Tripped() bool {
	return p.tripped
}

func (p *Protector) OpenAndReset(breaker *Breaker) error {
	if err := breaker.Open(); err != nil {
		return err
	}
	p.ResetFilter()
	p.tripped = false
	return nil
}
