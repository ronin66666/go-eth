package service

import (
	"context"
	"fmt"
	"log"
	"testing"
)

func TestBalanceOf(t *testing.T) {
	s, err := NewService(context.Background(), "https://test.lixb.io")

	if err != nil {
		log.Fatal(err)
	}

	balance, err := s.BalanceOf("0x1218836BAFE6d71160AD0275007ffF1Ea71A0348", "0xFcd64fd5CDe2A7c60BeADE304A08c9Ab4d610751")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(balance.String())

}
