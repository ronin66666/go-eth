package chain

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ronin66666/go-eth/global"
	"log"
	"math/big"
	"testing"
)

// 获取区块信息
func TestGetBlockByNumber(t *testing.T) {
	client, err := NewClient(global.ETH_RUL)
	if err != nil {
		log.Fatal(err)
	}

	block, err := client.GetBlockByNumber(big.NewInt(16001987))
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
	fmt.Println("transaction count = ", block.Size())

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

func getTransferReceipt(client *Client, txHash common.Hash) {
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
