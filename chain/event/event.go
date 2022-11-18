package event

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"utils/abi/libra_auction"
	"utils/abi/usdt_deposit"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const ethUrl = "https://mainnet.infura.io/v3/92983deb8689407bb1736bdf82bf9c9c"

//busd token
const deposit = "0xe069aE4B336Ca73142cDc5206ed4a4d3A3ff39f6 "

// FilterEvent
//filtering for evens: 主要用于抓取历史事件
func FilterEvent() {
	client, err := ethclient.Dial(ethUrl)
	if err != nil {
		log.Fatalln(err)
	}

	addr := common.HexToAddress(deposit)

	ct, err := usdt_deposit.NewUsdtDeposit(addr, client)
	if err != nil {
		log.Fatalln(err)
	}
	const start uint64 = 15889067
	//var end uint64 = start + 2

	filterOpts := &bind.FilterOpts{
		Context: context.Background(), Start: start, End: nil,
	}

	itr, err := ct.FilterDeposited(filterOpts)
	if err != nil {
		log.Fatalln(err)
	}
	for itr.Next() {
		event := itr.Event
		data, err := json.Marshal(event)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(string(data))
	}
}

// 只支持ws连接
func WatchEvent() {
	client, err := ethclient.Dial(ethUrl)
	if err != nil {
		log.Fatalln(err)
	}

	addr := common.HexToAddress(deposit)

	ct, err := libra_auction.NewLibraAuction(addr, client)
	if err != nil {
		log.Fatalln(err)
	}

	var start uint64 = 341957
	watchOpts := &bind.WatchOpts{
		Start:   &start,
		Context: context.Background(),
	}

	chanel := make(chan *libra_auction.LibraAuctionRebate)

	go func() {
		sub, err := ct.WatchRebate(watchOpts, chanel, nil)
		if err != nil {
			log.Fatalln("watch rebate error = ", err)
		}
		defer sub.Unsubscribe()

	}()

	event := <-chanel
	data, err := json.Marshal(event)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(data))
}
