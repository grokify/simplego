package timeutil

import (
	"fmt"
	"time"
)

// TimeSeriesSlice builds a time series based on supplied interval.
func TimeSeriesSlice(interval Interval, times []time.Time) []time.Time {
	if len(times) == 0 {
		return times
	}
	min, max := SliceMinMax(times)
	return TimeSeriesMinMax(interval, min, max)
}

// TimeSeriesMinMax builds a time series based on supplied interval.
func TimeSeriesMinMax(interval Interval, min, max time.Time) []time.Time {
	min, max = MinMax(min, max)
	series := []time.Time{}
	switch interval {
	case Month:
		min = MonthStart(min)
		max = MonthStart(max)
		series = append(series, min)
		cur := min
		for {
			cur = cur.AddDate(0, 1, 0)
			if cur.After(max) {
				break
			}
			series = append(series, cur)
		}
	default:
		panic(fmt.Sprintf("E_INTERVAL_NOT_SUPPORTED [%v]", interval))
	}
	return series
}
