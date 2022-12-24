package main

import (
	"fmt"
	"github.com/ronin66666/go-eth/wallet"
)

func main() {
	w := wallet.GenerateWallet()
	fmt.Println(w)
}
