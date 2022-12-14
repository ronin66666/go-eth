package utils

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

// import (
//
//	"crypto/ecdsa"
//	"math/big"
//	"reflect"
//	"regexp"
//	"strconv"
//
//	"github.com/ethereum/go-ethereum/common"
//	"github.com/ethereum/go-ethereum/common/hexutil"
//	"github.com/ethereum/go-ethereum/crypto"
//	"github.com/shopspring/decimal"
//
// )
//
//	func PrivateKeyToAddress(iprivateKey interface{}) common.Address {
//		var address common.Address
//		switch p := iprivateKey.(type) {
//		case []byte:
//
//		default:
//			return address
//		}
//	}

func GetAccountAuth(client *ethclient.Client, accountAddress string) *bind.TransactOpts {

	privateKey, err := crypto.HexToECDSA(accountAddress)
	if err != nil {
		panic(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic("invalid key")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	//fetch the last use nonce of account
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		panic(err)
	}
	fmt.Println("nounce=", nonce)
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		panic(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		panic(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // in wei
	//auth.GasLimit = uint64(3000000) // in units
	//auth.GasPrice = big.NewInt(1000000)

	return auth
}

//
//
//// IsValidAddress validate hex address
//func IsValidAddress(iaddress interface{}) bool {
//	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
//	switch v := iaddress.(type) {
//	case string:
//		return re.MatchString(v)
//	case common.Address:
//		return re.MatchString(v.Hex())
//	default:
//		return false
//	}
//}
//
//func ZeroAddress() common.Address {
//	return common.HexToAddress("0x0000000000000000000000000000000000000000")
//}
//
//// IsZeroAddress validate if it's a 0 address
//func IsZeroAddress(iaddress interface{}) bool {
//	var address common.Address
//	switch v := iaddress.(type) {
//	case string:
//		address = common.HexToAddress(v)
//	case common.Address:
//		address = v
//	default:
//		return false
//	}
//
//	zeroAddressBytes := common.FromHex("0x0000000000000000000000000000000000000000")
//	addressBytes := address.Bytes()
//	return reflect.DeepEqual(addressBytes, zeroAddressBytes)
//}
//
//// ToDecimal wei to decimals
//func ToDecimal(ivalue interface{}, decimals int) decimal.Decimal {
//	value := new(big.Int)
//	switch v := ivalue.(type) {
//	case string:
//		value.SetString(v, 10)
//	case *big.Int:
//		value = v
//	}
//
//	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
//	num, _ := decimal.NewFromString(value.String())
//	result := num.Div(mul)
//
//	return result
//}
//
//// ToWei decimals to wei
//func ToWei(iamount interface{}, decimals int) *big.Int {
//	amount := decimal.NewFromFloat(0)
//	switch v := iamount.(type) {
//	case string:
//		amount, _ = decimal.NewFromString(v)
//	case float64:
//		amount = decimal.NewFromFloat(v)
//	case int64:
//		amount = decimal.NewFromFloat(float64(v))
//	case int:
//		amount = decimal.NewFromFloat(float64(v))
//	case decimal.Decimal:
//		amount = v
//	case *decimal.Decimal:
//		amount = *v
//	}
//
//	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
//	result := amount.Mul(mul)
//
//	wei := new(big.Int)
//	wei.SetString(result.String(), 10)
//
//	return wei
//}
//
//// CalcGasCost calculate gas cost given gas limit (units) and gas price (wei)
//func CalcGasCost(gasLimit uint64, gasPrice *big.Int) *big.Int {
//	gasLimitBig := big.NewInt(int64(gasLimit))
//	return gasLimitBig.Mul(gasLimitBig, gasPrice)
//}
//
//// SigRSV signatures R S V returned as arrays
//func SigRSV(isig interface{}) ([32]byte, [32]byte, uint8) {
//	var sig []byte
//	switch v := isig.(type) {
//	case []byte:
//		sig = v
//	case string:
//		sig, _ = hexutil.Decode(v)
//	}
//
//	sigstr := common.Bytes2Hex(sig)
//	rS := sigstr[0:64]
//	sS := sigstr[64:128]
//	R := [32]byte{}
//	S := [32]byte{}
//	copy(R[:], common.FromHex(rS))
//	copy(S[:], common.FromHex(sS))
//	vStr := sigstr[128:130]
//	vI, _ := strconv.Atoi(vStr)
//	V := uint8(vI + 27)
//
//	return R, S, V
//}
//
//// LogTopicHash generates the log topic hash given the event function declaration signature
//func LogTopicHash(fnsig string) common.Hash {
//	eventSignature := []byte(fnsig)
//	hash := crypto.Keccak256Hash(eventSignature)
//	return hash
//}
