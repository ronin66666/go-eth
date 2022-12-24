package subscribe

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ronin66666/go-eth/internal/abi/subscribe"
	"log"
)

// 2888
const lixbTestURl = "https://test.lixb.io"
const subscribeAddr = "0x8EB0C7072a6e258cC73228C46d6C7e5683E7b3CA"

func FilterDeposited(client *ethclient.Client, subscribeAddr string) {

	addr := common.HexToAddress(subscribeAddr)

	ct, err := subscribe.NewSubscribe(addr, client)
	if err != nil {
		log.Fatalln(err)
	}
	const start uint64 = 411899
	filterOpts := &bind.FilterOpts{
		Start:   start,
		End:     nil,
		Context: context.Background(),
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
