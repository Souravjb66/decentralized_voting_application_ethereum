package main

import (
	// "context"
	"context"
	"io/ioutil"
	"math/big"

	"fmt"
	"log"

	abi "votede/abgen"
	// "votede/mywallet"

	"net/http"
	method "votede/contractmethod"
	"votede/routes"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"

	// "github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	// "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/mux"
)

type Alldetail struct {
	Nonce    uint64
	GasPrice *big.Int
	Auth     *bind.TransactOpts
}

var Details Alldetail

func main() {
	fmt.Println("server start")
	dest, _ := ioutil.ReadFile("./mywallet/UTC--2024-03-11T10-27-08.250407292Z--98298abf352d714d193697501789c52ea3920cec")
	router := mux.NewRouter()
	var sepolaUrl = "https://eth-sepolia.g.alchemy.com/v2/lh4oR7wrbDIA2c2qBJMo5g5NDk_YtuKB"
	client, err := ethclient.Dial(sepolaUrl)
	if err != nil {
		fmt.Println(err)
	}
	defer client.Close()
	// claddr := "0x39601Da937225a60FaFA1BAF8b4586E993aBbFA1"

	// decFile,_:= keystore.DecryptKey(dest, "password")         //decrypt the encrypt file using assign password

	// ac1:=crypto.PubkeyToAddress(decFile.PrivateKey.PublicKey)
	// nonce, _:= client.PendingNonceAt(context.TODO(), ac1)
	// gasprice, _ := client.SuggestGasPrice(context.TODO())
	// chainId, _ := client.NetworkID(context.TODO())
	// auth, _ := bind.NewKeyedTransactorWithChainID(decFile.PrivateKey,chainId)
	// Details.Auth=auth
	// Details.GasPrice=gasprice
	// Details.Nonce=nonce
	decFile, _ := keystore.DecryptKey(dest, "password")
	fmt.Println(decFile)
	chainId, cerr := client.NetworkID(context.TODO())
	if cerr != nil {
		fmt.Println(cerr)
	}
	auth, auerror := bind.NewKeyedTransactorWithChainID(decFile.PrivateKey, chainId)
	if auerror != nil {
		fmt.Println(auerror)
		return
	}

	myadd := "0xfa4c1fb70bddf67cc1f5029a6272adfe4f0baa85" //to it is contract address
	add := common.HexToAddress(myadd)
	contract, conerror := abi.NewVote(add, client) //making a connction to smart contract using contract address and object of sepola test network where contract is deployed
	if conerror != nil {
		fmt.Println(conerror)
	}
	method.Solcontract.Contract = contract
	method.Solcontract.Client = client
	method.Solcontract.Auth = auth

	router.HandleFunc("/setmember", routes.ForSetMember).Methods("POST")
	router.HandleFunc("/give-vote", routes.ForGiveVote).Methods("POST")
	router.HandleFunc("/get-candidate", routes.ToGetCandidate).Methods("GET")
	router.HandleFunc("/get-members", routes.ToGetMembers).Methods("GET")
	router.HandleFunc("/get-totalvote", routes.ToGetAllVote).Methods("GET")
	router.HandleFunc("get-votemember", routes.ToGetVotedMemmber).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))

}
