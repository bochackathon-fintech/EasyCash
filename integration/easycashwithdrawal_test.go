package integration

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestView(t *testing.T) {
	if GetView() != "5710bba5d42604e4072d1e92" {
		t.Error("Expected 5710bba5d42604e4072d1e92")
	}
}

func TestBalance(t *testing.T) {
	view := "5710bba5d42604e4072d1e92"
	if GetBalance(view) != 607.34 {
		t.Error("Expected 607.34")
	}
}

func TestGetUser(t *testing.T) {
	user := "d9395cc00979c72735b51715"
	AuthProviderNameBob := "012012012665"
	if GetUser(AuthProviderNameBob) != user {
		t.Error("Expected d9395cc00979c72735b51715")
	}
	// user2 := "d9395cc00979c72735b51715"
	if GetUser("09102") != "" {
		t.Error("Expected empty customerID")
	}
}

func TestParseView(t *testing.T) {
	file, e := ioutil.ReadFile("C:\\Users\\user\\Go\\src\\github.com\\matteo107\\humanevolutionapi\\samples\\GetViewsResponse.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	fmt.Printf("%s\n", string(file))

	//m := new(Dispatch)
	//var m interface{}
	var jsontype MsgViewArray
	json.Unmarshal(file, &jsontype)
	fmt.Printf("Results: %v\n", jsontype)
}

func TestParseAccount(t *testing.T) {
	file, e := ioutil.ReadFile("C:\\Users\\user\\Go\\src\\github.com\\matteo107\\humanevolutionapi\\samples\\GetBalanceAccount.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	fmt.Printf("%s\n", string(file))

	//m := new(Dispatch)
	//var m interface{}
	var jsontype MsgAccountJson
	json.Unmarshal(file, &jsontype)
	fmt.Printf("Results: %v\n", jsontype)
}

func TestParseUser(t *testing.T) {
	file, e := ioutil.ReadFile("C:\\Users\\user\\Go\\src\\github.com\\matteo107\\humanevolutionapi\\samples\\GetUser.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	fmt.Printf("%s\n", string(file))

	//m := new(Dispatch)
	//var m interface{}
	var jsontype MsgUserJson
	json.Unmarshal(file, &jsontype)
	fmt.Printf("Results: %v\n", jsontype)
}

func TestParseMakeTransactionReq(t *testing.T) {
	file, e := ioutil.ReadFile("C:\\Users\\user\\Go\\src\\github.com\\matteo107\\humanevolutionapi\\samples\\PostMakeTransactionReq.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	fmt.Printf("%s\n", string(file))

	//m := new(Dispatch)
	//var m interface{}
	var jsontype MsgMakeTransactionReq
	json.Unmarshal(file, &jsontype)
	fmt.Printf("Results: %v\n", jsontype)

}
func TestMsgMakeTransactionResp(t *testing.T) {
	file, e := ioutil.ReadFile("C:\\Users\\user\\Go\\src\\github.com\\matteo107\\humanevolutionapi\\samples\\PostMakeTransactionResp.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	fmt.Printf("%s\n", string(file))

	//m := new(Dispatch)
	//var m interface{}
	var jsontype MsgMakeTransactionResp
	json.Unmarshal(file, &jsontype)
	fmt.Printf("Results: %v\n", jsontype)
}

func TestParseAccounts(t *testing.T) {
	file, e := ioutil.ReadFile("C:\\Users\\user\\Go\\src\\github.com\\matteo107\\humanevolutionapi\\samples\\GetAccountsResponse.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	fmt.Printf("%s\n", string(file))

	//m := new(Dispatch)
	//var m interface{}
	var jsontype MsgAccountsArray
	json.Unmarshal(file, &jsontype)
	fmt.Printf("Results: %v\n", jsontype)
}
