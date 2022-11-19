package chain

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ronin66666/go-eth/global"
	"log"
)

// GetTransactionCount 根据区块hash获取该块中的交易数量
func (client *Client) GetTransactionCount(blockHash common.Hash) (uint, error) {
	return client.TransactionCount(context.Background(), blockHash)
}

//GetTransactions 获取交易信息
func (client *Client) GetTransactions(block *types.Block) {
	chainId, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	for _, tx := range block.Transactions() {
		fmt.Println(tx.Hash().Hex())

		// 读取发送方的地址
		// AsMessage() 方法需要EIP155签名者
		msg, err := tx.AsMessage(types.NewEIP155Signer(chainId), nil)
		if err != nil {
			global.Logger.Error(err)
		}
		fmt.Println("msg.from = ", msg.From().Hex(), " to = ", tx.To().Hex())

		data, err := tx.MarshalJSON()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(data))
	}
}

func (client *Client) GetTransactionReceipt(txHash common.Hash) {
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
