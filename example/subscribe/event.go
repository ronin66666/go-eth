package subscribe

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ronin66666/go-eth/contract"
	"log"
)

// FilterReserved 根据开始和结束区块查询日志信息
func (s *Service) FilterReserved(ctx context.Context, subscribeAddr string, start uint64, end *uint64) {

	addr := common.HexToAddress(subscribeAddr)
	ct, err := contract.NewSubscribeV2(addr, s)

	if err != nil {
		log.Fatalln(err)
	}

	filterOpts := &bind.FilterOpts{
		Start:   start,
		End:     end,
		Context: ctx,
	}

	itr, err := ct.FilterReserved(filterOpts)
	if err != nil {
		log.Fatalln(err)
	}

	for itr.Next() {
		raw := itr.Event.Raw
		txLog, err := ct.ParseReserved(raw) //将日志解析为对象
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Printf(" %v \n", txLog)
	}
}

func (s *Service) FilterRefund(ctx context.Context, contractAddr string, start uint64, end *uint64) {
	addr := common.HexToAddress(contractAddr)
	ct, err := contract.NewSubscribeV2(addr, s)

	if err != nil {
		log.Fatalln(err)
	}

	filterOpts := &bind.FilterOpts{
		Start:   start,
		End:     end,
		Context: ctx,
	}

	//var from = []common.Address{common.HexToAddress("0x7EA470A705C67257cbe9bE45e08824706c2a490A"), common.HexToAddress("0xAADBD70f7d25710632A7669281902202533A9FB7")}

	itr, err := ct.FilterRefunded(filterOpts)
	if err != nil {
		log.Fatalln(err)
	}

	for itr.Next() {
		raw := itr.Event.Raw
		txLog, err := ct.ParseReserved(raw) //将日志解析为对象
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Printf(" %v \n", txLog)
	}
}
