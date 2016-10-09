package main

import (
	"fmt"
	"strconv"
)

type PharmacyGen struct{}

var Pharmacy = &PharmacyGen{}

func (pg *PharmacyGen) Drug() string {
	return drugNames[random.Intn(len(drugNames))]
}

func (pg *PharmacyGen) Strength() string {
	str := strconv.Itoa(random.Intn(9999) * 5 % 500)
	units := [2]string{"mg", "mL"}
	return fmt.Sprintf("%s %s", str, units[random.Intn(2)])
}

func (pg *PharmacyGen) FullDrug() string {
	return fmt.Sprintf("%s %s", pg.Drug(), pg.Strength())
}

func (pg *PharmacyGen) Form() string {
	return drugForms[random.Intn(len(drugForms))]
}
