package routes

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"math/big"
	"net/http"
	method "votede/contractmethod"
)

func ForSetMember(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("type of :", reflect.TypeOf(address))
	type Member struct {
		Name string
		Age  uint16
		User common.Address
	}
	var member Member
	json.NewDecoder(r.Body).Decode(&member)
	result := method.TransactionSetMember(member.User, member.Name, member.Age)
	if result == 404 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("not found")
		return
	}
	w.WriteHeader(200)
	json.NewEncoder(w).Encode("submited")

}
func ForGiveVote(w http.ResponseWriter, r *http.Request) {
	type Voter struct {
		From      common.Address
		Candidate common.Address
	}
	var myvote Voter
	json.NewDecoder(r.Body).Decode(&myvote)
	result := method.TransactionGiveVote(myvote.From, myvote.Candidate)
	if result == 404 {
		w.WriteHeader(404)
		json.NewEncoder(w).Encode("not found")
		return
	}
	w.WriteHeader(200)
	json.NewEncoder(w).Encode("succesfuly voting")

}
func ToGetCandidate(w http.ResponseWriter, r *http.Request) {
	var user common.Address
	json.NewDecoder(r.Body).Decode(&user)
	result := method.TransactionGetCandidate(user)
	if result == 404 {
		w.WriteHeader(404)
		json.NewEncoder(w).Encode("not found")
		return
	}
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(result)

}
func ToGetMembers(w http.ResponseWriter, r *http.Request) {
	var user common.Address
	json.NewDecoder(r.Body).Decode(&user)
	result := method.TransactionGetMembers(user)
	if result == nil {
		w.WriteHeader(404)
		json.NewEncoder(w).Encode("not found")
		return
	}
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(result)

}
func ToGetAllVote(w http.ResponseWriter, r *http.Request) {
	type vote struct {
		First  *big.Int
		Second *big.Int
		Third  *big.Int
	}

	var user common.Address
	json.NewDecoder(r.Body).Decode(&user)
	one, two, three, err := method.TransactionGetAllVote(user)
	if err != nil {
		w.WriteHeader(404)
		json.NewEncoder(w).Encode("not found")
		return
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(vote{First: one, Second: two, Third: three})

}
func ToGetVotedMemmber(w http.ResponseWriter, r *http.Request) {
	var user common.Address
	json.NewDecoder(r.Body).Decode(&user)
	result, err := method.TransactionGetVotedMember(user)
	if err != nil {
		log.Println(err)
		w.WriteHeader(404)
		json.NewEncoder(w).Encode("not found")
		return
	}
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(result)

}
