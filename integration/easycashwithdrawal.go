package integration

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var BankID = "bda8eb884efcef7082792d45"
var AuthProviderName = "01440902243000"
var AuthID = "123456789"
var OcpApimSubscriptionKey = "05d6874316504138959f5e9cd9d3c7d0"

const (
	Host       string = "api.bocapi.net"
	ApiVersion string = "/v1"
)

type ViewStruct struct {
	ID string `json:"id"`
}
type ViewsStruct struct {
	Views []ViewStruct `json:"views"`
}

func GetBalance() string {

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

	var record ViewsStruct
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	bodyString := string(bodyBytes)
	log.Println(bodyString)

	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
	}
	fmt.Println("Record", record)
	fmt.Println("View ID. = ", record.Views[0].ID)
	return record.Views[0].ID
}
