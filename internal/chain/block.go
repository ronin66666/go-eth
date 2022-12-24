package chain

import (
	"context"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

//GetLastHeaderByNumber 您可以调用客户端的HeaderByNumber 来返回有关一个区块的头信息。若您传入nil，它将返回最新的区块头。
func (client *Client) GetLastHeaderByNumber(number *big.Int) (*types.Header, error) {

	return client.HeaderByNumber(context.Background(), number)
}

//GetBlockByNumber 调用客户端的BlockByNumber方法来获得完整区块。您可以读取该区块的所有内容和元数据，例如，区块号，区块时间戳，区块摘要，区块难度以及交易列表等等。
func (client *Client) GetBlockByNumber(blockNumber *big.Int) (*types.Block, error) {
	return client.BlockByNumber(context.Background(), blockNumber)
}
