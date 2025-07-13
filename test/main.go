package main

import (
	"fmt"
	"time"
)

func main() {

	latestArchiveStartTime := time.Date(2025, time.February, 24, 07, 05, 00, 0, time.UTC)

	latestArchiveMigratedTime := time.Date(2025, time.February, 24, 7, 0, 0, 0, time.UTC)

	offset := time.Minute * 10
	//windowLength := time.Minute * 10

	//jobWindowEnd := lamt.Add(windowLength)
	jobWindowEnd := latestArchiveMigratedTime

	updateUpcomingMigrationTimeMetric(offset, latestArchiveStartTime, jobWindowEnd)
}

// wl = 10
// off = 10m
// lamt = 7
func updateUpcomingMigrationTimeMetric(offset time.Duration, latestArchiveStartTime, jobWindowEnd time.Time) {
	t := latestArchiveStartTime.Add(-offset).Sub(jobWindowEnd).Hours()

	fmt.Printf("Upcoming migration time metric for ISID %v in signed hours %f\n", 128, t)

	//var t1 float64
	//if t > 0 {
	//	t1 = t.Hours()
	//} else {
	//	t1 = 0
	//}

	//fmt.Printf("Upcoming migration time metric for ISID %v in hours %f", 128, t1)
}
