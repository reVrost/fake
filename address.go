package main

import (
	"fmt"
	"strconv"
)

type AddressGen struct{}

var Address = &AddressGen{}

func (ag *AddressGen) ZipCode(char int) string {
	return Number.Code(char)
}

func (ag *AddressGen) StreetSuffix() string {
	return streetSuffix[random.Intn(len(streetSuffix))]
}

func (ag *AddressGen) StreetSuffixAbv() string {
	return streetSuffixAbv[random.Intn(len(streetSuffixAbv))]
}

func (ag *AddressGen) StreetName() string {
	return streets[random.Intn(len(streets))]
}

func (ag *AddressGen) Australian() string {
	// "12/213 Gazilas Rd, Surry Hills NSW 2020"

	return fmt.Sprintf(
		"%s/%s %s %s, %s %s %s",
		strconv.Itoa(Number.Number(999)),
		strconv.Itoa(Number.Number(999)),
		ag.StreetName(),
		ag.StreetSuffix(),
		australianSuburbs[random.Intn(len(australianSuburbs))],
		australianState[random.Intn(len(australianState))],
		ag.ZipCode(4),
	)
}
