package actions

import "github.com/gobuffalo/buffalo"

import "github.com/matteo107/humanevolutionapi/integration"

// EasyCashWithdrawalRequestMake default implementation.
func EasyCashWithdrawalRequestMake(c buffalo.Context) error {
	//BoCAPI -> GetBalance(AuthProviderName,AuthID)
	integration.GetBalance()
	//BoCAPI -> GetUser(Bob)

	//to Bob -> WithdrawRequest(ForAlice,Amount)

	return c.Render(200, r.HTML("easy_cash_withdrawal_request/make.html"))
}
