package main

import (
	"fmt"
	"log"

	abi "votede/abgen"

	"net/http"
	method "votede/contractmethod"
	"votede/routes"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("server start")
	router := mux.NewRouter()
	var sepolaUrl = "https://eth-sepolia.g.alchemy.com/v2/lh4oR7wrbDIA2c2qBJMo5g5NDk_YtuKB"
	client, err := ethclient.Dial(sepolaUrl)
	if err != nil {
		fmt.Println(err)
	}
	defer client.Close()

	myadd := "0xfa4c1fb70bddf67cc1f5029a6272adfe4f0baa85" //to it is contract address
	add := common.HexToAddress(myadd)
	contract, conerror := abi.NewVote(add, client) //making a connction to smart contract using contract address and object of sepola test network where contract is deployed
	if conerror != nil {
		fmt.Println(conerror)
	}
	method.Solcontract.Contract = contract

	router.HandleFunc("/setmember", routes.ForSetMember).Methods("POST")
	router.HandleFunc("/give-vote", routes.ForGiveVote).Methods("POST")
	router.HandleFunc("/get-candidate", routes.ToGetCandidate).Methods("GET")
	router.HandleFunc("/get-members", routes.ToGetMembers).Methods("GET")
	router.HandleFunc("/get-totalvote", routes.ToGetAllVote).Methods("GET")
	router.HandleFunc("get-votemember", routes.ToGetVotedMemmber).Methods("GET")
	log.Fatal(http.ListenAndServe(":8081", router))

}
