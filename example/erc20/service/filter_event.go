package service

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ronin66666/go-eth/contract"
	"log"
)

// FilterTransfer 根据开始和结束区块查询日志信息
func (s *Service) FilterTransfer(ctx context.Context, subscribeAddr string, Start uint64, End *uint64, from, to []common.Address) {

	addr := common.HexToAddress(subscribeAddr)
	ct, err := contract.NewERC20(addr, s)

	if err != nil {
		log.Fatalln(err)
	}

	filterOpts := &bind.FilterOpts{
		Start:   Start,
		End:     End,
		Context: ctx,
	}

	//var from = []common.Address{common.HexToAddress("0x7EA470A705C67257cbe9bE45e08824706c2a490A"), common.HexToAddress("0xAADBD70f7d25710632A7669281902202533A9FB7")}

	itr, err := ct.FilterTransfer(filterOpts, from, to)
	if err != nil {
		log.Fatalln(err)
	}

	for itr.Next() {
		raw := itr.Event.Raw
		txLog, err := ct.ParseTransfer(raw) //将日志解析为对象
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Printf(" %v \n", txLog)
	}
}

// WatchTransfer ws 来监听事件日志
func (s *Service) WatchTransfer(ctx context.Context, contractAddr string, start *uint64) {
	addr := common.HexToAddress(contractAddr)
	ct, err := contract.NewERC20(addr, s)
	if err != nil {
		log.Fatalln(err)
	}
	watchOpts := &bind.WatchOpts{
		Start:   start,
		Context: ctx,
	}
	//初始化通道
	sink := make(chan *contract.ERC20Transfer)

	subscription, err := ct.WatchTransfer(watchOpts, sink, nil, nil)

	if err != nil {
		log.Fatalln("watch ", err)
	}

	for {
		select {
		case deposited := <-sink:
			fmt.Printf("deposited %v ", deposited)
		case err := <-subscription.Err():
			subscription.Unsubscribe()
			log.Fatalln(err)
		}
	}
}
