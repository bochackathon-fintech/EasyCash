package actions

import (
	"fmt"
	"strconv"

	"github.com/gobuffalo/buffalo"
	"github.com/matteo107/humanevolutionapi/integration"
)

// EasyCashWithdrawalRequestMake default implementation.
func EasyCashWithdrawalRequestMake(c buffalo.Context) error {
	//BoCAPI -> GetBalance(AuthProviderName,AuthID)
	viewid := integration.GetView()
	balance := integration.GetBalance(viewid)
	balString := fmt.Sprintf("%.2f", balance)

	c.Set("balance", balString)
	requestedAmount := c.Request().URL.Query().Get("amount")
	c.Set("requested_amount", requestedAmount)
	value, err := strconv.ParseFloat(requestedAmount, 32)
	if err != nil {
		// do something sensible
	}
	var reqamount = float32(value)
	if balance < reqamount {
		c.Set("authresponse", "Withdrawal declined. Not Enough Balance")
		return c.Render(200, r.HTML("easy_cash_withdrawal_request/make.html"))
	} else {
		c.Set("authresponse", "Withdrawal authorized")
	}
	customerID := integration.GetUser(integration.AuthProviderNameBob)
	// customerID := integration.GetUser("12345")
	if customerID == "" {
		c.Set("authresponse", "Withdrawal declined. Human ATM not found")
		return c.Render(200, r.HTML("easy_cash_withdrawal_request/make.html"))
	} else {
		// ask Bob to authorized with websocket?
		//to Bob -> WithdrawRequest(ForAlice,Amount)
		//
		// BoCAPI -> Make Payment
	}

	return c.Render(200, r.HTML("easy_cash_withdrawal_request/make.html"))
}
