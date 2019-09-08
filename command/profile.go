package command

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/tesujiro/spoc/global"
)

func (cmd *Command) getProfile(endpoint string) {
	b, err := cmd.Api.Get(endpoint, nil)
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
	if !global.FlagOnlyIDs {
		fmt.Printf("User:\t%v\tType: %v\tDisplayName: %v\n", user.Id, user.Type, user.DisplayName)
	} else {
		fmt.Printf("%v\n", user.Id)
	}
}

func (cmd *Command) GetMyProfile() {
	endpoint := cmd.endpoint("profile/me")
	cmd.getProfile(endpoint)
	return
}

func (cmd *Command) GetUserProfile(id string) {
	endpoint := strings.ReplaceAll(cmd.endpoint("profile"), "{user_id}", id)
	cmd.getProfile(endpoint)
	return
}
