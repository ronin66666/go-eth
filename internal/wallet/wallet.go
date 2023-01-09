package wallet

import (
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"io/ioutil"
	"log"
)

type Wallet struct {
	Address    string `json:"address"`
	PrivateKey string `json:"privateKey"`
	PublicKey  string `json:"publicKey"`
}

// GenerateWallet 生成钱包
func GenerateWallet() *Wallet {

	//生成随机私钥
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)
	fmt.Println("privateKey = ", hexutil.Encode(privateKeyBytes)[2:])

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println("publicKey = ", hexutil.Encode(publicKeyBytes)[4:])

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println("address = ", address) // 0x96216849c49358B10257cb55b28eA603c874b05E

	wallet := &Wallet{
		Address:    address,
		PublicKey:  hexutil.Encode(publicKeyBytes)[4:],
		PrivateKey: hexutil.Encode(privateKeyBytes)[2:],
	}
	return wallet
}

func GenerateWalletToFile(num int, fileName string) {
	var wallets []*Wallet

	for i := 0; i < num; i++ {
		fmt.Printf("**************生成第%d个账户*************\n", i)
		wallet := GenerateWallet()
		wallets = append(wallets, wallet)
	}

	jsonBytes, err := json.Marshal(&wallets)
	if err != nil {
		panic("json marshal err ")
	}
	//jsonString := string(jsonBytes)
	err = ioutil.WriteFile(fileName, jsonBytes, 0666)
	if err != nil {
		panic("写入文件失败")
	}
	//fmt.Printf("%v", jsonString)
}
