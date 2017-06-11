package actions

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"log"

	"github.com/gobuffalo/buffalo"
	"github.com/markbates/pop"
	"github.com/matteo107/easycash/integration"
	"github.com/matteo107/easycash/models"
	"github.com/pkg/errors"
)

// EasyCashWithdrawalRequestMake default implementation.
func EasyCashWithdrawalRequestShow(c buffalo.Context) error {
	log.SetFlags(log.Lshortfile)
	log.Println("EasyCashWithdrawalRequestShow")
	log.Println(c.Session().Session.Values)
	//BoCAPI -> GetBalance(AuthProviderName,AuthID)
	viewid := integration.GetView()
	balance := integration.GetBalance(viewid)
	balString := fmt.Sprintf("%.2f", balance)

	log.Println("Setting the bal in context")
	c.Set("balance", balString)
	log.Println("The bal is set in context")
	// requestedAmount := c.Value("amount").(string)
	requestedAmount := c.Session().Session.Values["amount"].(string)
	log.Println("Requested amount", requestedAmount)
	c.Set("requested_amount", requestedAmount)
	value, err := strconv.ParseFloat(requestedAmount, 32)
	if err != nil {
		return errors.New("Your requested amount is incorrect")
	}
	var reqamount = float32(value)
	log.Println("Check balance")
	if balance < reqamount {
		log.Println("balance < reqamount")
		return errors.New("Withdrawal declined. Not Enough Balance")
	}
	log.Println("Get user")
	// fromUser := c.Request().URL.Query().Get("user")
	fromUser := c.Session().Session.Values["user"].(string)
	log.Println("From user", fromUser)
	if fromUser == "" {
		return errors.New("Error.user not in input")
		// c.Set("authresponse", "Error.user not in input")
		// return c.Render(500, r.HTML("easy_cash_withdrawal_request/make.html"))
	}

	tx := models.DB
	query := pop.Q(tx)
	query = tx.Where("name = ?", fromUser)
	user := models.User{}
	err = query.First(&user)

	log.Println("User is returned from db")
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			msg := fmt.Sprintf("User %s not found. Not possible to authorize", fromUser)
			return errors.New(msg)

		}
		return errors.New("Error querying line 46")

		//return c.Render(500, r.HTML("easy_cash_withdrawal_request/make.html"))
	}
	log.Println("user.Authprovidername", user.Authprovidername)

	customerID := integration.GetUser(user.Authprovidername)
	// customerID := integration.GetUser(integration.AuthProviderNameBob)
	// customerID := integration.GetUser("12345")
	if customerID == "" {
		return errors.New("Withdrawal declined. Human ATM not found")

	} else {
		// ask Bob to authorized with websocket?
		//to Bob -> WithdrawRequest(ForAlice,Amount)
		//
		ToAccountID := integration.GetAccounts()
		if ToAccountID != "" {
			time.Sleep(time.Second * 3)
			status := integration.PostMakeTransaction(requestedAmount, ToAccountID)

			if status == "COMPLETED" {
				return errors.New(status + ". Your Transaction ID:" + integration.TransactionIDs[0])

			} else {
				return errors.New("Withdrawal declined. Internal error")
				// return c.Render(500, r.HTML("easy_cash_withdrawal_request/make.html"))
			}
		} else {
			return errors.New("Withdrawal declined. ToAccountID not found")
			// return c.Render(500, r.HTML("easy_cash_withdrawal_request/make.html"))
		}

	}
}

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
	// customerID := integration.GetUser(integration.AuthProviderNameBob)
	// customerID := integration.GetUser("12345")
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
				c.Set("authresponse", status)
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
