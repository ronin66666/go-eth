package main

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ronin66666/go-eth/global"
)

const (
	TestERC20Token = ""
)

func main() {

	client, err := ethclient.Dial(global.GOERLI_URL)
	if err != nil {

	}

}
