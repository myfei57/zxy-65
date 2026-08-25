package meter

import "fmt"

type Reading struct {
	Berth     string
	Timestamp string
	Value     float64
}

func NewReading(berth, timestamp string, value float64) (Reading, error) {
	if berth == "" {
		return Reading{}, fmt.Errorf("berth is required")
	}
	if timestamp == "" {
		return Reading{}, fmt.Errorf("timestamp is required")
	}
	if value < 0 {
		return Reading{}, fmt.Errorf("value must not be negative")
	}
	return Reading{Berth: berth, Timestamp: timestamp, Value: value}, nil
}

func (r Reading) String() string {
	return fmt.Sprintf("%s %s %.2f", r.Berth, r.Timestamp, r.Value)
}
