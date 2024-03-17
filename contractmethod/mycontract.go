package contractmethod

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	// "reflect"
	abi "votede/abgen"
)

type Mycontract struct {
	Contract *abi.Vote
}

var Solcontract Mycontract

func TransactionSetMember(address common.Address, name string, age uint16) uint {
	pvk := "79e0e280d7a984023e686ca6e40a08aeaa3010baa2406984bd64692405d1204b"
	password := "code@hyper256###"
	auth, _ := bind.NewTransactor(strings.NewReader(pvk), password)

	// fmt.Println(address)
	// fmt.Println("type of :", reflect.TypeOf(address))
	opts := &bind.TransactOpts{
		From: auth.From,
	}
	// &bind.TransactOpts{From: address}
	ss, err := Solcontract.Contract.SetMember(opts, name, age)
	if err != nil {
		fmt.Println(err)
		return 404
	}
	fmt.Println(ss)
	return 200

}
func TransactionGiveVote(from common.Address, vote common.Address) uint {
	vo, err := Solcontract.Contract.GiveVote(&bind.TransactOpts{From: from}, vote) //transactopts is use if function modify something in blockchain
	if err != nil {
		fmt.Println(err)
		return 404
	}
	fmt.Println(vo)
	return 200
}
func TransactionGetCandidate(user common.Address) uint16 {
	total, err := Solcontract.Contract.GetCandidate(&bind.CallOpts{From: user}) //callopts is use if function doesnt modify anything in blockchain
	if err != nil {
		fmt.Println(err)
		return 404
	}
	return total
}
func TransactionGetMembers(user common.Address) []abi.VotingMember { //
	value, err := Solcontract.Contract.GetMembers(&bind.CallOpts{From: user})
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return value
}
func TransactionGetAllVote(user common.Address) (*big.Int, *big.Int, *big.Int, error) {
	first, sec, third, err := Solcontract.Contract.GetAllVote(&bind.CallOpts{From: user})
	if err != nil {
		fmt.Println(err)
		return nil, nil, nil, err
	}
	return first, sec, third, nil

}
func TransactionGetVotedMember(user common.Address) (string, error) {
	value, err := Solcontract.Contract.VoterVote(&bind.CallOpts{From: user})

	if err != nil {
		fmt.Println(err)
		return "not found", err

	}
	vl := value.Hex()
	return vl, nil
}
