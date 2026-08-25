package ship

type Ship struct {
	ID    string
	Name  string
	Berth string
}

func NewShip(id, name string) *Ship {
	return &Ship{ID: id, Name: name}
}

func (s *Ship) Moor(berth string) {
	s.Berth = berth
}

func (s *Ship) MooredAt() string {
	return s.Berth
}
