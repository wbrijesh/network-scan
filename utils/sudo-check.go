package utils

import (
	"fmt"
	"os/user"
)

func CheckSudoPrivileges() (bool, error) {
	currentUser, err := user.Current()
	if err != nil {
		return false, fmt.Errorf("error getting current user: %w", err)
	}
	return currentUser.Uid == "0", nil
}
