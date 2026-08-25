package phase

func ParseSequence(value string) (Sequence, bool) {
	switch value {
	case "ABC":
		return SequenceForward, true
	case "CBA":
		return SequenceReverse, true
	default:
		return "", false
	}
}
