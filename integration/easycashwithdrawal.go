package integration

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var BankID = "bda8eb884efcef7082792d45"
var AccountID = "bda8eb884efcea209b2a6240"
var AuthProviderName = "01440902243000"
var AuthProviderNameBob = "012012012665"
var AuthID = "123456789"
var OcpApimSubscriptionKey = "05d6874316504138959f5e9cd9d3c7d0"

const (
	Host       string = "api.bocapi.net"
	ApiVersion string = "/v1"
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

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return -1
	}

	defer resp.Body.Close()

	var record MsgAccountJson

	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
	}
	fmt.Println("Record", record)
	fmt.Println("Balance amount = ", record.Balance.Amount)
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

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return "err"
	}

	defer resp.Body.Close()

	var record MsgViewArray

	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
	}
	fmt.Println("Record", record)
	fmt.Println("View ID. = ", record.MsgView[0].ID)
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

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return "err"
	}

	defer resp.Body.Close()

	var record MsgUserJson

	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
	}
	fmt.Println("Record", record)
	fmt.Println("Customer ID = ", record.CustomerID)
	return record.CustomerID
}
