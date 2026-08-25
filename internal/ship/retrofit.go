package ship

import "errors"

func (r *Registry) Retrofit(ship string, addedLoad float64) error {
	p, ok := r.Load(ship)
	if !ok {
		return errors.New("ship profile not found")
	}
	p.MaxLoad += addedLoad
	return r.Save(p)
}
