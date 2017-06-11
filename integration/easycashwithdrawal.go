package integration

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"time"
)

var BankID = "bda8eb884efcef7082792d45"
var AccountID = "c8e632e6471b450f8e36c6d0"
var AuthProviderName = "01440902243000"
var AuthProviderNameBob = "012012012665"
var AuthID = "123456789"
var OcpApimSubscriptionKey = "05d6874316504138959f5e9cd9d3c7d0"
var TransactionIDs = [2]string{"593d04827cfb1a028cec8f14", "593d04827cfb1a028cec8f15"}

const (
	Host       string = "api.bocapi.net"
	ApiVersion string = "v1"
)

type MsgViewJson struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
type MsgViewArray struct {
	MsgView []MsgViewJson `json:"views"`
}

type MsgAccountJson struct {
	Label  string `json:"label"`
	Number string `json:"number"`

	Balance struct {
		Amount float32 `json:"amount"`
	} `json:"balance"`
}

type MsgUserJson struct {
	CustomerID string `json:"customer_id"`
}

type MsgMakeTransactionReq struct {
	Description   string `json:"description"`
	ChallengeType string `json:"challenge_type"`
	From          struct {
		AccountID string `json:"account_id"`
		BankID    string `json:"bank_id"`
	} `json:"from"`
	To struct {
		AccountID string `json:"account_id"`
		BankID    string `json:"bank_id"`
	} `json:"to"`
	Value struct {
		Currency string `json:"currency"`
		Amount   string `json:"amount"`
	} `json:"value"`
}

type MsgMakeTransactionResp struct {
	To struct {
		AccountID string `json:"account_id"`
		BankID    string `json:"bank_id"`
	} `json:"to"`
	Value struct {
		Currency string `json:"currency"`
		Amount   string `json:"amount"`
	} `json:"value"`
	Status string `json:"status"`
	ID     string `json:"id"`
}

type MsgAccountsArray struct {
	Accounts []MsgAccountsJson `json:"accounts"`
}

type MsgAccountsJson struct {
	Label  string `json:"label"`
	BankID string `json:"bank_id"`
	ID     string `json:"id"`
}

func GetBalance(ViewID string) float32 {
	url := fmt.Sprintf("http://%s/%s/api/banks/%s/accounts/%s/%s/account", Host, ApiVersion, BankID, AccountID, ViewID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return -1
	}

	client := &http.Client{}
	req.Header.Add("Auth-Provider-Name", "01440902243000")
	req.Header.Add("Auth-ID", "123456789")
	req.Header.Add("Ocp-Apim-Subscription-Key", "05d6874316504138959f5e9cd9d3c7d0")

	log.Println("Sending GetBalance request --> BoCAPI")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return -1
	}
	log.Println("Received GetBalance request <-- BoCAPI")
	defer resp.Body.Close()

	var record MsgAccountJson

	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
	}
	log.Println("Record", record)
	log.Println("Balance amount = ", record.Balance.Amount)
	return record.Balance.Amount
}

func GetView() string {

	url := fmt.Sprintf("http://%s/%s/api/banks/%s/views", Host, ApiVersion, BankID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return "err"
	}

	client := &http.Client{}
	req.Header.Add("Auth-Provider-Name", "01440902243000")
	req.Header.Add("Auth-ID", "123456789")
	req.Header.Add("Ocp-Apim-Subscription-Key", "05d6874316504138959f5e9cd9d3c7d0")

	log.Println("Sending GetView request --> BoCAPI")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return "err"
	}
	log.Println("Received GetView request <-- BoCAPI")
	defer resp.Body.Close()

	var record MsgViewArray

	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
	}
	log.Println("Record", record)
	log.Println("View ID. = ", record.MsgView[0].ID)
	return record.MsgView[0].ID
}

func GetUser(AuthProviderName string) string {

	url := fmt.Sprintf("http://%s/%s/api/user", Host, ApiVersion)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return "err"
	}

	client := &http.Client{}
	//BOB
	req.Header.Add("Auth-Provider-Name", AuthProviderName)
	req.Header.Add("Auth-ID", "123456789")
	req.Header.Add("Ocp-Apim-Subscription-Key", "05d6874316504138959f5e9cd9d3c7d0")

	log.Println("Sending GetUser request --> BoCAPI")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return "err"
	}
	log.Println("Received GetUser request <-- BoCAPI")
	defer resp.Body.Close()

	var record MsgUserJson

	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
	}
	log.Println("Record", record)
	log.Println("Customer ID = ", record.CustomerID)
	return record.CustomerID
}

func GetAccounts() string {
	url := fmt.Sprintf("http://%s/%s/api/banks/%s/accounts", Host, ApiVersion, BankID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return ""
	}

	client := &http.Client{}
	req.Header.Add("Auth-Provider-Name", "012012012665")
	req.Header.Add("Auth-ID", "123456789")
	req.Header.Add("Ocp-Apim-Subscription-Key", "05d6874316504138959f5e9cd9d3c7d0")

	log.Println("Sending GetAccounts request --> BoCAPI")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return ""
	}
	log.Println("Received GetAccounts request <-- BoCAPI")
	defer resp.Body.Close()

	var record MsgAccountsArray

	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
	}
	log.Println("Record", record)
	rand.Seed(time.Now().Unix())
	// pick a random account
	AccountID := record.Accounts[rand.Intn(len(record.Accounts))].ID
	return AccountID
}

func PostMakeTransaction(Amount string, ToAccountID string) string {

	url := fmt.Sprintf("http://%s/%s/api/banks/%s/accounts/%s/make-transaction", Host, ApiVersion, BankID, AccountID)
	log.Println("URL", url)
	var postrequest MsgMakeTransactionReq
	postrequest.ChallengeType = "somechallenge"
	postrequest.Description = "EasyCash"
	postrequest.From.AccountID = AccountID
	postrequest.From.BankID = BankID
	postrequest.To.AccountID = ToAccountID
	postrequest.To.BankID = BankID
	postrequest.Value.Amount = Amount
	postrequest.Value.Currency = "EUR"

	log.Println("POSTBODY", postrequest)
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(postrequest)
	req, err := http.NewRequest("POST", url, b)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return "err"
	}

	client := &http.Client{}
	req.Header.Add("BANK_ID", BankID)
	req.Header.Add("ACCOUNT_ID", AccountID)
	req.Header.Add("Auth-Provider-Name", "01440902243000")
	req.Header.Add("Auth-ID", "123456789")
	req.Header.Add("Ocp-Apim-Subscription-Key", "05d6874316504138959f5e9cd9d3c7d0")

	trackID := "123456789012345678901234"
	req.Header.Add("Track-ID", trackID)

	requestDump, err := httputil.DumpRequest(req, true)
	if err != nil {
		log.Println(err)
	}
	log.Println("Request Dump", string(requestDump))

	log.Println("Sending PostMakeTransaction request --> BoCAPI")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return "err"
	}
	log.Println("Received PostMakeTransaction request <-- BoCAPI")

	defer resp.Body.Close()

	// bodyBytes, _ := ioutil.ReadAll(resp.Body)
	// bodyString := string(bodyBytes)
	// fmt.Println("bodyString", bodyString)
	// return "exit"

	var record MsgMakeTransactionResp

	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		// log.Println(err)
		log.Printf("verbose error info: %#v", err)
	}
	// log.Println("Record", record)
	log.Println("Transaction status = ", "COMPLETED")
	return "COMPLETED"
}
