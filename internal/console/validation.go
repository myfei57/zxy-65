package console

import "strings"

func normalizeCode(value string) string {
	return strings.TrimSpace(value)
}

func validBerthCode(value string) bool {
	code := normalizeCode(value)
	return code != "" && len(code) <= 16
}

func validShipID(value string) bool {
	id := normalizeCode(value)
	return len(id) >= 2 && len(id) <= 32
}

func validFrequency(value float64) bool {
	return value >= 45 && value <= 65
}

func validLoad(value float64) bool {
	return value > 0
}

func validTemperature(value float64) bool {
	return value >= -40 && value <= 200
}
