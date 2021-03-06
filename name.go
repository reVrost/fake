package fake

import (
	"fmt"
)

type nameGen struct{}

// Gender specifies a gender of a person
type Gender int

const (
	Any = iota
	Male
	Female
)

// Name generator
var Name = &nameGen{}

func (ng *nameGen) LastName() string {
	return lastNames[random.Intn(len(lastNames))]
}

func (ng *nameGen) FirstName(g Gender) string {
	var firstName string
	switch g {
	case Male:
		firstName = maleFirstNames[random.Intn(len(maleFirstNames))]
	case Female:
		firstName = femaleFirstNames[random.Intn(len(femaleFirstNames))]
	default:
		rn := random.Intn(100)
		if rn > 49 {
			firstName = femaleFirstNames[random.Intn(len(femaleFirstNames))]
		}
		firstName = maleFirstNames[random.Intn(len(maleFirstNames))]
	}
	return firstName
}

func (ng *nameGen) FullName(g Gender) string {
	return fmt.Sprintf(
		"%s %s",
		ng.FirstName(g),
		ng.LastName(),
	)
}
