package phase

type State struct {
	current Sequence
	history []Sequence
}

func NewState(initial Sequence) *State {
	return &State{current: initial, history: []Sequence{initial}}
}

func (s *State) Set(seq Sequence) {
	s.current = seq
	s.history = append(s.history, seq)
}

func (s *State) Current() Sequence {
	return s.current
}

func (s *State) History() []Sequence {
	out := make([]Sequence, len(s.history))
	copy(out, s.history)
	return out
}

func (s *State) Reversed() bool {
	return s.current == SequenceReverse
}
