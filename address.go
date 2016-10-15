package fake

import (
	"fmt"
	"strconv"
)

type addressGen struct{}

var Address = &addressGen{}

func (ag *addressGen) ZipCode(char int) string {
	return Number.Code(char)
}

func (ag *addressGen) StreetSuffix() string {
	return streetSuffix[random.Intn(len(streetSuffix))]
}

func (ag *addressGen) StreetSuffixAbv() string {
	return streetSuffixAbv[random.Intn(len(streetSuffixAbv))]
}

func (ag *addressGen) StreetName() string {
	return streets[random.Intn(len(streets))]
}

func (ag *addressGen) Australian() string {
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
