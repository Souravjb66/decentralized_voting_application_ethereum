package alladdress

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type Candidate struct {
	First  common.Address
	Second common.Address
	Third  common.Address
}
type Users struct {
	One   common.Address
	Two   common.Address
	Three common.Address
}

var Allcandidate Candidate
var Allusers Users

func CandidateAddress() {
	pvkMemOne, oneerr := crypto.GenerateKey()
	if oneerr != nil {
		fmt.Println(oneerr)

	}
	firstCanAddr := crypto.PubkeyToAddress(pvkMemOne.PublicKey)
	fmt.Println("first candi : ", firstCanAddr)

	pvkMemTwo, twoerr := crypto.GenerateKey()
	if twoerr != nil {
		fmt.Println(twoerr)
	}
	secondCanAddr := crypto.PubkeyToAddress(pvkMemTwo.PublicKey)
	fmt.Println("second candi : ", secondCanAddr)

	pvkMemThird, thirderr := crypto.GenerateKey()
	if thirderr != nil {
		fmt.Println(thirderr)
	}
	thirdCanAddr := crypto.PubkeyToAddress(pvkMemThird.PublicKey)
	fmt.Println("third candi : ", thirdCanAddr)

	Allcandidate = Candidate{
		First:  firstCanAddr,
		Second: secondCanAddr,
		Third:  thirdCanAddr,
	}

}
func VoterAddress() {
	pvkVone, _ := crypto.GenerateKey()
	voneAddr := crypto.PubkeyToAddress(pvkVone.PublicKey)
	fmt.Println("one voter : ", voneAddr)

	pvkVtwo, _ := crypto.GenerateKey()
	vtwoAddr := crypto.PubkeyToAddress(pvkVtwo.PublicKey)
	fmt.Println("two voter : ", vtwoAddr)

	pvkVthree, _ := crypto.GenerateKey()
	vthreeAddr := crypto.PubkeyToAddress(pvkVthree.PublicKey)
	fmt.Println("three voter : ", vthreeAddr)

	Allusers = Users{
		One:   voneAddr,
		Two:   vtwoAddr,
		Three: vthreeAddr,
	}
}
