package mywallet

import (
	"fmt"
	// "io"
	"io/ioutil"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	// "os"
	"github.com/ethereum/go-ethereum/accounts/keystore"
)

type AllKeys struct {
	Publick  common.Address
	PrivateK common.Address
	PAddress common.Address
}

var Acaddress AllKeys

func Wallet() {
	password := "password"
	fileF1 := "./mywallet/UTC--2024-03-11T10-27-08.250407292Z--98298abf352d714d193697501789c52ea3920cec"
	fileF2 := "./mywallet/UTC--2024-03-11T10-27-09.837561219Z--f5c9135f06f42bf6a37895aa91f9aa53d4e5cbbf"
	acOne, onerr := ioutil.ReadFile(fileF1)
	if onerr != nil {
		log.Println(onerr)
		return
	}
	acTwo, twoerr := ioutil.ReadFile(fileF2)
	if twoerr != nil {
		log.Println(twoerr)
		return
	}
	forOne, derrone := keystore.DecryptKey(acOne, password)
	if derrone != nil {
		log.Println(derrone)
		return
	}
	pvkAcOne := crypto.FromECDSA(forOne.PrivateKey)
	pvkHex := hexutil.Encode(pvkAcOne)

	pubkeyOne := crypto.FromECDSAPub(&forOne.PrivateKey.PublicKey)
	pubHex := hexutil.Encode(pubkeyOne)

	AddOne := crypto.PubkeyToAddress(forOne.PrivateKey.PublicKey)

	forTwo, derrtwo := keystore.DecryptKey(acTwo, password)
	if derrtwo != nil {
		log.Println()
	}

	pvkAcTwo := crypto.FromECDSA(forTwo.PrivateKey)
	fmt.Println(hexutil.Encode(pvkAcTwo))
	Acaddress = AllKeys{
		Publick:  common.HexToAddress(pubHex),
		PrivateK: common.HexToAddress(pvkHex),
		PAddress: AddOne,
	}

}
