package fake

import (
	"math/rand"
	"time"
)

var random *rand.Rand

func init() {
	random = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// Want something like this to reduce code duplicate but..
//func Raffle(list []interface{}) interface{} {
//return list[random.Intn(len(list))]
//}
