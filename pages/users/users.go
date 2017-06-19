package users

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
)

// Get ...
func Get(ctx *aero.Context) string {
	users, err := arn.GetActiveUsersCached()

	if err != nil {
		return ctx.Error(http.StatusInternalServerError, "Could not fetch user data", err)
	}

	return ctx.HTML(components.Users(users))
}