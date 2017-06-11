package actions

import "github.com/gobuffalo/buffalo"

// Bob's page
func ServeBob(c buffalo.Context) error {
	return c.Render(200, r.HTML("bob.html"))
}
