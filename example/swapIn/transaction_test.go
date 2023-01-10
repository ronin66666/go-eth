package swapIn

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"math/big"
	"testing"
)

//test address 0x294e3E70Aa2078f62ca32291d169537A9e6d10a3

func TestService_SwapInfo(t *testing.T) {
	s, err := NewService(context.Background(), "https://mainnet.infura.io/v3/92983deb8689407bb1736bdf82bf9c9c")

	if err != nil {
		log.Fatal(err)
	}

	//swapInfo, err := s.SwapInfo("0xD799EbFf3a2d3AE9dca57e8EcFD844B3268Bd3Dc", big.NewInt(39093))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Printf("%v", swapInfo)

	swapInfo2, err := s.SwapInfo("0xD799EbFf3a2d3AE9dca57e8EcFD844B3268Bd3Dc", big.NewInt(39094))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v", swapInfo2)

}

func TestService_Transfer(t *testing.T) {

	s, err := NewService(context.Background(), "https://test.lixb.io")

	if err != nil {
		log.Fatal(err)
	}

	//privateKey: string
	//只负责发送交易， 不会等到成功的回调
	tx, err := s.Transfer("", "0x1218836BAFE6d71160AD0275007ffF1Ea71A0348", "0x294e3E70Aa2078f62ca32291d169537A9e6d10a3", big.NewInt(1000))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("txHash = ", tx.Hash())
	txData, err := tx.MarshalJSON()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("txData = ", string(txData))

}

func TestService_TransactionReceiptJson(t *testing.T) {
	s, err := NewService(context.Background(), "https://test.lixb.io")

	if err != nil {
		log.Fatal(err)
	}
	data, err := s.TransactionReceiptJson(common.HexToHash("0x3e6bfab767a580c59e90d151df6224ffd9061c2b29f62f03ab69c864d1d1a59f"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))

}
