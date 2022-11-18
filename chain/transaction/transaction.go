package transaction

import (
	"bytes"
	"context"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// getTransactionCount 根据区块hash获取该块中的交易数量
func getTransactionCount(client *ethclient.Client, blockHash common.Hash) (uint, error) {
	return client.TransactionCount(context.Background(), blockHash)
}

// 获取交易信息
func GetTransactions(client *ethclient.Client, block *types.Block) {
	chainId, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	for _, tx := range block.Transactions() {

		fmt.Println(tx.Hash().Hex())

		// 读取发送方的地址
		// AsMessage() 方法需要EIP155签名者
		msg, err := tx.AsMessage(types.NewEIP155Signer(chainId), nil)
		if err != nil {
			log.Fatal(err)
		}
		
		fmt.Println("msg.from = ", msg.From().Hex(), " to = ", tx.To().Hex())

		
		buf := new(bytes.Buffer)
		if err := tx.EncodeRLP(buf); err != nil {
			log.Fatal(err)
		}
		data := hex.EncodeToString(buf.Bytes())
		fmt.Println("data = ", data)

		GetTransactionRecipt(client, tx.Hash())
	}
}

func GetTransactionRecipt(client *ethclient.Client, txHash common.Hash) {
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

