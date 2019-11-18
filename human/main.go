package human

import (
	"github.com/dustin/go-humanize"
	"time"
)

func Unix_Time(unix int64) string {
	unix_time := time.Unix(unix, 0)
	return humanize.Time(unix_time)
}
