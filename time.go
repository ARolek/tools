package tools

import (
	"time"

	"github.com/dustin/go-humanize"
)

func FmtTs(ts int64) string {
	t := time.Unix(0, ts*int64(time.Millisecond))
	if ts == 0 {
		return "not set"
	}
	return humanize.Time(t)
}
