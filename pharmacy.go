package fake

import (
	"fmt"
	"strconv"
)

type pharmacyGen struct{}

var Pharmacy = &pharmacyGen{}

func (pg *pharmacyGen) Drug() string {
	return drugNames[random.Intn(len(drugNames))]
}

func (pg *pharmacyGen) Strength() string {
	str := strconv.Itoa(random.Intn(9999) * 5 % 500)
	units := [2]string{"mg", "mL"}
	return fmt.Sprintf("%s %s", str, units[random.Intn(2)])
}

func (pg *pharmacyGen) FullDrug() string {
	return fmt.Sprintf("%s %s", pg.Drug(), pg.Strength())
}

func (pg *pharmacyGen) Form() string {
	return drugForms[random.Intn(len(drugForms))]
}
