package main

import (
	"context"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ronin66666/go-eth/events/subscribe"
	"github.com/ronin66666/go-eth/global"
	"github.com/sirupsen/logrus"
	"io"
	"io/fs"
	"log"
	"os"
)

func init() {
	//setupLogger()
}

func main() {
	//w := wallet.GenerateWallet()
	//fmt.Println(w)

	client, err := ethclient.DialContext(context.Background(), "https://lixb.io")
	if err != nil {
		log.Fatal(err)
	}

	//chain.GetTransactionReceipt(client, common.HexToHash("0x14e4966b2480c682f1cde158e1e582692cf0fd182141e270f45740fc7959e6df"))

	//wallet.GenerateWallet()

	subscribe.FilterDeposited(client, "0x127B29D77fE7FA99cc9ADb50ecE0BF0C0c721cf0")
}

func setupLogger() {
	logger := logrus.New()
	logger.SetLevel(logrus.TraceLevel)
	logger.SetReportCaller(true) //是否输出文件名

	fileName := "storage/logs/log.log"
	fwriter, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, fs.ModePerm)
	if err != nil {
		logger.Error(err)
	}

	logger.SetOutput(io.MultiWriter(os.Stdout, fwriter))
	global.Logger = logger
}
