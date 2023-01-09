package subscribe

import (
	"context"
	"log"
	"testing"
)

func TestSubscribe_FilterReserved(t *testing.T) {
	s, err := NewService(context.Background(), "https://test.lixb.io")
	if err != nil {
		log.Fatal(err)
	}
	const start uint64 = 450238
	s.FilterReserved(context.Background(), "0xEf4aD3f4a7E1f59Dc6D17B7f87e984f349f1828D", start, nil)

}
