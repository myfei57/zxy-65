package plug

func (p *Plug) Release() error {
	if !p.engaged {
		return errNotEngaged
	}
	p.engaged = false
	return nil
}
