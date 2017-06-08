package main

import (
	"github.com/animenotifier/arn"
	"github.com/fatih/color"
)

func main() {
	color.Yellow("Updating user references")

	arn.Truncate("NickToUser")
	arn.Truncate("EmailToUser")

	// Get a stream of all anime
	allUsers, err := arn.AllUsers()

	if err != nil {
		panic(err)
	}

	// Iterate over the stream
	count := 0
	for user := range allUsers {
		count++
		println(count, user.Nick)

		user.SetNick(user.Nick)

		if user.Email != "" {
			user.SetEmail(user.Email)
		}
	}

	color.Green("Finished.")
}
