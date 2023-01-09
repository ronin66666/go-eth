package subscribe

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"math/big"
	"testing"
)

// 获取最新区块
func TestLastBlockNumber(t *testing.T) {

	s, err := NewService(context.Background(), "https://test.lixb.io")
	if err != nil {
		log.Fatal(err)
	}

	n, err := s.BlockNumber(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	println("lastBlockNumber = ", n)
}

// 获取区块信息
func TestGetBlockByNumber(t *testing.T) {
	client, err := NewService(context.Background(), "https://test.lixb.io")
	if err != nil {
		log.Fatal(err)
	}

	block, err := client.BlockByNumber(context.Background(), big.NewInt(449414))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(block.Hash())

	// 获取区块头信息
	header, err := json.Marshal(block.Header())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("header = ", string(header))

	//获取交易数量
	count := block.Transactions().Len()
	fmt.Println("transaction count = ", count)
	fmt.Println("transaction size = ", block.Size())

	//遍历获取每笔交易的信息
	for _, tx := range block.Transactions() {
		fmt.Println("transaction hash = ", tx.Hash())
		//buf := new(bytes.Buffer) //tx RLP编码信息
		//if err := tx.EncodeRLP(buf); err != nil {
		//	log.Fatal(err)
		//}
		data, err := tx.MarshalJSON()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("json data = ", string(data))
		getTransferReceipt(client, tx.Hash())
	}
}

func getTransferReceipt(client *Service, txHash common.Hash) {
	receipt, err := client.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("receipt.Status = ", receipt.Status)
	bytes, err := receipt.MarshalJSON()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("receipt = ", string(bytes))
}
