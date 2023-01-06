package service

import (
	"context"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Service struct {
	*ethclient.Client
}

func NewService(ctx context.Context, url string) (*Service, error) {
	client, err := ethclient.DialContext(ctx, url)
	return &Service{
		client,
	}, err
}
