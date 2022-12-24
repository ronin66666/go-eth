package main

import (
	"fmt"
	"github.com/ronin66666/go-eth/global"
	"github.com/ronin66666/go-eth/internal/wallet"
	"github.com/sirupsen/logrus"
	"io"
	"io/fs"
	"os"
)

func init() {
	setupLogger()
}

func main() {
	w := wallet.GenerateWallet()
	fmt.Println(w)
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
