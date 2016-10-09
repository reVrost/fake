package main

import (
	"time"
)

type DateGen struct{}

var Date = &DateGen{}

func (dg *DateGen) Past(years int, startDate time.Time) time.Time {
	var maxSeconds = years * 365 * 24 * 60 * 60
	return startDate.Add(time.Duration(random.Intn(maxSeconds)) * time.Second * -1)
}

func (dg *DateGen) Future(years int, startDate time.Time) time.Time {
	var maxSeconds = years * 365 * 24 * 60 * 60
	return startDate.Add(time.Duration(random.Intn(maxSeconds)) * time.Second)
}

func (dg *DateGen) Between(startDate time.Time, endDate time.Time) time.Time {
	var diffDuration = endDate.Sub(startDate)
	var maxSeconds = int(diffDuration.Seconds())
	return startDate.Add(time.Duration(random.Intn(maxSeconds)) * time.Second)
}
