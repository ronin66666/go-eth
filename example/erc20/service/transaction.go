package service

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ronin66666/go-eth/contract"
	"log"
	"math/big"
)

func (s *Service) BalanceOf(ctAddr string, account string) (*big.Int, error) {

	ct, err := contract.NewERC20(common.HexToAddress(ctAddr), s)
	if err != nil {
		log.Fatal(err)
	}
	return ct.BalanceOf(nil, common.HexToAddress(account))
}