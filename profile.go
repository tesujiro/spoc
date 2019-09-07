package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func (spoc *Spoc) profile(endpoint string) {
	b, err := spoc.get(endpoint, nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	var user User
	err = json.Unmarshal(b, &user)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	if !flagOnlyIDs {
		fmt.Printf("User:\t%v\tType: %v\tDisplayName: %v\n", user.Id, user.Type, user.DisplayName)
	} else {
		fmt.Printf("%v\n", user.Id)
	}
}
