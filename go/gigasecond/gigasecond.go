package gigasecond

import "time"

const GIGA_SECOND = 1_000_000_000

func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * GIGA_SECOND)
}
