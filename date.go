package fake

import (
	"time"
)

type dateGen struct{}

var Date = &dateGen{}

func (dg *dateGen) Past(years int, startDate time.Time) time.Time {
	var maxSeconds = years * 365 * 24 * 60 * 60
	return startDate.Add(time.Duration(random.Intn(maxSeconds)) * time.Second * -1)
}

func (dg *dateGen) Future(years int, startDate time.Time) time.Time {
	var maxSeconds = years * 365 * 24 * 60 * 60
	return startDate.Add(time.Duration(random.Intn(maxSeconds)) * time.Second)
}

func (dg *dateGen) Between(startDate time.Time, endDate time.Time) time.Time {
	var diffDuration = endDate.Sub(startDate)
	var maxSeconds = int(diffDuration.Seconds())
	return startDate.Add(time.Duration(random.Intn(maxSeconds)) * time.Second)
}
