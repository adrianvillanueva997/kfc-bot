package randomgeneration

import (
	"crypto/rand"
	"log"
	"math/big"
)

func RandomInteger(max, min int) int {
	number, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		log.Fatalln(err)
	}
	tmp := *number

	return int(tmp.Int64())
}
