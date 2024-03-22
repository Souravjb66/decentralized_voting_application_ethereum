package contractmethod

import (
	"context"
	"crypto/ecdsa"
	"fmt"

	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	// "reflect"
	abi "votede/abgen"
)

type Mycontract struct {
	Contract *abi.Vote
	Client   *ethclient.Client
	Auth     *bind.TransactOpts
}

var Solcontract Mycontract

func TransactionSetMember(pvk *ecdsa.PrivateKey, name string, age uint16) uint {
	add := crypto.PubkeyToAddress(pvk.PublicKey)

	Solcontract.Auth.GasLimit = 300000
	a, aerr := Solcontract.Client.SuggestGasPrice(context.TODO())
	if aerr != nil {
		fmt.Println(aerr)
		return 404
	}
	b, berror := Solcontract.Client.PendingNonceAt(context.TODO(), add)
	if berror != nil {
		fmt.Println(berror)
		return 404
	}
	Solcontract.Auth.GasPrice = a
	Solcontract.Auth.Nonce = big.NewInt(int64(b))

	ss, err := Solcontract.Contract.SetMember(Solcontract.Auth, name, age)
	if err != nil {
		fmt.Println(err)
		return 404
	}
	fmt.Println(ss)
	return 200

}
func TransactionGiveVote(from *ecdsa.PrivateKey, vote common.Address) uint {
	a, _ := Solcontract.Client.SuggestGasPrice(context.TODO())
	b, _ := Solcontract.Client.PendingNonceAt(context.TODO(), crypto.PubkeyToAddress(from.PublicKey))
	Solcontract.Auth.GasLimit = 300000
	Solcontract.Auth.GasPrice = a
	Solcontract.Auth.Nonce = big.NewInt(int64(b))

	vo, err := Solcontract.Contract.GiveVote(Solcontract.Auth, vote) //transactopts is use if function modify something in blockchain
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
