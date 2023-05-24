package utils

import (
	"math/rand"
	"strconv"
	"time"
)

func GetRandSixCode() string {
	rand.Seed(time.Now().UnixNano())

	code := ""
	for i := 0; i < 6; i++ {
		code = code + strconv.Itoa(rand.Intn(10))
	}
	return code
}
