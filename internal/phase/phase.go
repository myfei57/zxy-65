package phase

type Sequence string

const (
	SequenceForward Sequence = "ABC"
	SequenceReverse Sequence = "CBA"
)

type Verifier struct {
	sequence Sequence
}

func NewVerifier(seq Sequence) *Verifier {
	return &Verifier{sequence: seq}
}

func (v *Verifier) IsForward() bool {
	return v.sequence == SequenceForward
}
