package disconn

import (
	"errors"

	"portpower/internal/freq"
)

func (b *Breaker) ParallelClose(sync *freq.Syncer, target float64, tolerance float64) error {
	current := sync.Current()
	if !freq.InBand(current, target, tolerance) {
		return errors.New("frequency out of sync")
	}
	b.close()
	return nil
}
