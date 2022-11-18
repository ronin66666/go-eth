package wallet

import (
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/crypto"
)

type Wallet struct {
	Address    string `json:"address"`
	PrivateKey string `json:"privateKey"`
	PublicKey  string `json:"publicKey"`
}

type PartnerWallet struct {
	Partner *Wallet   `json:"partner"`
	Levels  []*Wallet `json:"levels"`
}

// GeneratePartnerWalletToFile 生成partner， 文件追加
func GeneratePartnerWalletToFile(num int, fileName string) {
	partners, err := PartnerJsonFileToStruct(fileName)
	if err != nil {
		panic(err)
	}
	for i := 0; i < num; i++ {
		partnerWallet := new(PartnerWallet)
		partner := GenerateWallet()
		partnerWallet.Partner = partner

		for i := 0; i < 21; i++ {
			level := GenerateWallet()
			partnerWallet.Levels = append(partnerWallet.Levels, level)
		}
		partners = append(partners, partnerWallet)
	}
	jsonBytes, err := json.Marshal(&partners)
	if err != nil {
		panic("json marshal err ")
	}
	//jsonString := string(jsonBytes)
	err = ioutil.WriteFile(fileName, jsonBytes, 0666)
	if err != nil {
		panic("写入文件失败")
	}
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

// PartnerJsonFileToStruct 读取partner Jons 数据
func PartnerJsonFileToStruct(path string) ([]*PartnerWallet, error) {
	jsonFile, err := os.Open(path)
	if err != nil {
		fmt.Println("error opening json file")
		return nil, err
	}
	defer jsonFile.Close()
	decoder := json.NewDecoder(jsonFile)
	for {
		var partners []*PartnerWallet
		err := decoder.Decode(&partners)
		if err == io.EOF { //当所有的Json文件的数据被解码完毕以后，Decode方法会返回一个EOF，这个时候退出循环
			return partners, nil
		}
		if err != nil {
			fmt.Println("解码失败")
			return nil, err
		}
		fmt.Println("wallet count = ", len(partners))
		return partners, nil

	}

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
