package main

import (
	"fmt"
	"github.com/t-okkn/go-enjaxel/crypto"
)

var (
	Version string
	Revision string
)

// summary => main関数
/////////////////////////////////////////
func main() {
	fmt.Println(crypto.PasswordHash("Hoge"))
}
