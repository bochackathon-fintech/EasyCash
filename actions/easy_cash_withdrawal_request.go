package actions

import (
	"database/sql"
	"fmt"
	"strconv"

	"log"

	"github.com/gobuffalo/buffalo"
	"github.com/markbates/pop"
	"github.com/matteo107/humanevolutionapi/integration"
	"github.com/matteo107/humanevolutionapi/models"
	"github.com/pkg/errors"
)

// EasyCashWithdrawalRequestMake default implementation.
func EasyCashWithdrawalRequestMake(c buffalo.Context) error {
	log.SetFlags(log.Lshortfile)
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
	fromUser := c.Request().URL.Query().Get("user")
	if fromUser == "" {
		c.Set("authresponse", "Error.user not in input")
		return c.Render(500, r.HTML("easy_cash_withdrawal_request/make.html"))
	}

	tx := models.DB
	query := pop.Q(tx)
	query = tx.Where("name = ?", fromUser)
	user := models.User{}
	err = query.First(&user)

	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			msg := fmt.Sprintf("User %s not found. Not possible to authorize", fromUser)
			c.Set("authresponse", msg)
			return c.Render(500, r.HTML("easy_cash_withdrawal_request/make.html"))
		}
		c.Set("authresponse", "Error querying line 46")
		return errors.WithStack(err)
		//return c.Render(500, r.HTML("easy_cash_withdrawal_request/make.html"))
	}
	log.Println("user.Authprovidername", user.Authprovidername)

	customerID := integration.GetUser(user.Authprovidername)

	if customerID == "" {
		c.Set("authresponse", "Withdrawal declined. Human ATM not found")
		return c.Render(500, r.HTML("easy_cash_withdrawal_request/make.html"))
	} else {
		// ask Bob to authorized with websocket?
		//to Bob -> WithdrawRequest(ForAlice,Amount)
		//
		ToAccountID := integration.GetAccounts()
		if ToAccountID != "" {
			status := integration.PostMakeTransaction(requestedAmount, ToAccountID)

			if status == "COMPLETED" {

				c.Set("authresponse", status+". Your Transaction ID:"+integration.TransactionIDs[0])
				return c.Render(200, r.HTML("easy_cash_withdrawal_request/make.html"))
			} else {
				c.Set("authresponse", "Withdrawal declined. Internal error")
				return c.Render(500, r.HTML("easy_cash_withdrawal_request/make.html"))
			}
		} else {
			c.Set("authresponse", "Withdrawal declined. ToAccountID not found")
			return c.Render(500, r.HTML("easy_cash_withdrawal_request/make.html"))
		}

	}
}
