package fake

import (
	"strconv"
)

type NumberGen struct{}

var Number = &NumberGen{}

func (ng *NumberGen) Number(max int) int {
	return random.Intn(max)
}

func (ng *NumberGen) Code(maxChar int) string {
	code := strconv.Itoa((random.Intn(10) % 9) + 1)
	for i := 1; i < maxChar; i++ {
		code += strconv.Itoa(random.Intn(10))
	}
	return code
}
