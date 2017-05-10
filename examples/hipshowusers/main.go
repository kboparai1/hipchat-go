package main

import (
	"flag"
	"fmt"

	"github.com/tbruyelle/hipchat-go/hipchat"
)

var (
	token           = flag.String("token", "", "The HipChat AuthToken")
	maxResults      = flag.Int("maxResults", 1000, "Max results per request")
	includeDeleted = flag.Bool("includeDeleted", false, "Include deleted users?")
	includeGuests = flag.Bool("includeGuests", false, "Include deleted users?")
)

func main() {
	flag.Parse()
	if *token == "" {
		flag.PrintDefaults()
		return
	}
	c := hipchat.NewClient(*token)
	startIndex := 0
	totalRequests := 0
	var allUsers []hipchat.User

	for {
		opt := &hipchat.UserListOptions{
			ListOptions:     hipchat.ListOptions{StartIndex: startIndex, MaxResults: *maxResults},
			IncludeDeleted:  *includeDeleted,
			IncludeGuests: *includeGuests}

		users, resp, err := c.User.List(opt)

		if err != nil {
			fmt.Printf("Error during user list req %q\n", err)
			fmt.Printf("Server returns %+v\n", resp)
			return
		}

		totalRequests++

		allUsers = append(allUsers, users.Items...)
		if users.Links.Next != "" {
			startIndex += *maxResults
		} else {
			break
		}
	}

	fmt.Printf("Your group has %d users, it took %d requests to retrieve all of them:\n",
		len(allUsers), totalRequests)
	for _, r := range allUsers {
		fmt.Printf("%d %s \n", r.ID, r.Name)
	}
}
