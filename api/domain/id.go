package domain

import "crypto/rand"

func GenID() int64 {
	p, _ := rand.Prime(rand.Reader, 32)
	return p.Int64()
}
