package freq

func InBand(current, target, tolerance float64) bool {
	diff := current - target
	if diff < 0 {
		diff = -diff
	}
	return diff <= tolerance
}
