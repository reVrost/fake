package fake

import (
	"strconv"
)

type numberGen struct{}

var Number = &numberGen{}

func (ng *numberGen) Number(max int) int {
	return random.Intn(max)
}

func (ng *numberGen) Code(maxChar int) string {
	code := strconv.Itoa((random.Intn(10) % 9) + 1)
	for i := 1; i < maxChar; i++ {
		code += strconv.Itoa(random.Intn(10))
	}
	return code
}
