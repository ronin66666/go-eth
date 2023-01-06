package service

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"testing"
)

func TestFilterTransfer(t *testing.T) {

	s, err := NewService(context.Background(), "https://test.lixb.io")

	if err != nil {
		log.Fatal(err)
	}
	//chain.GetTransactionReceipt(service, common.HexToHash("0x14e4966b2480c682f1cde158e1e582692cf0fd182141e270f45740fc7959e6df"))
	const start uint64 = 447933
	var end uint64 = 449414 //nil为查询到最新区块
	//查询交易日志
	var from = []common.Address{common.HexToAddress("0x7EA470A705C67257cbe9bE45e08824706c2a490A"), common.HexToAddress("0xAADBD70f7d25710632A7669281902202533A9FB7")}

	s.FilterTransfer(context.Background(), "0x1218836BAFE6d71160AD0275007ffF1Ea71A0348", start, &end, from, nil)
}

// WS监听交易日志
func TestWatchTransfer(t *testing.T) {

	s, err := NewService(context.Background(), "https://test.lixb.io")

	if err != nil {
		log.Fatal(err)
	}

	var start uint64 = 447933

	s.WatchTransfer(context.Background(), "0x127B29D77fE7FA99cc9ADb50ecE0BF0C0c721cf0", &start)
}
