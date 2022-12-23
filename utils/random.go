package utils

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)


const alphabet = "abcdefghijklmnopqrstuvwxyz" 

func init(){
	rand.Seed(time.Now().UnixNano())
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomName() string {
	return RandomString(6)
}

func RandomEmail() string {
	return fmt.Sprintf("%s@%s.com", RandomString(5), RandomString(3))
}
