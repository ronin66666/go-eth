package client

import "github.com/ethereum/go-ethereum/ethclient"

func GetClient(url string) (*ethclient.Client, error) {
	return ethclient.Dial(url)
}
