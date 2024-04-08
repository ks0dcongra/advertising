package helpers

import "time"

var set time.Time

func Now() time.Time {
	if set.IsZero() {
		return time.Now()
	}
	return set
}

// SetNow only used by verifyDateTime function in unit testing
func SetNow(n time.Time) {
    set = n.In(time.FixedZone("CST", 8*3600)) 
}