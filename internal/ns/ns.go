package ns

type Harbor struct {
	ID   string
	Name string
}

func NewHarbor(id, name string) *Harbor {
	return &Harbor{ID: id, Name: name}
}

func (h *Harbor) Identity() string {
	return h.Name + ":" + h.ID
}
