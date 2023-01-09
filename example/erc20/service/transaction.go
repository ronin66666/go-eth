package service

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ronin66666/go-eth/contract"
	"github.com/ronin66666/go-eth/pkg/eth/utils"
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

func (s *Service) Transfer(privateKey string, ctAddr string, to string, amount *big.Int) (*types.Transaction, error) {
	ct, err := contract.NewERC20(common.HexToAddress(ctAddr), s)
	if err != nil {
		log.Fatal(err)
	}
	auth := utils.GetAccountAuth(s.Client, privateKey)
	return ct.Transfer(auth, common.HexToAddress(to), amount)
}

func (s *Service) TransactionReceiptJson(txHash common.Hash) ([]byte, error) {
	receipt, err := s.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("receipt.Status = ", receipt.Status)
	return receipt.MarshalJSON()

}
