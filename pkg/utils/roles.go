package utils

import (
	"fmt"

	"github.com/jaram2024/jaram-auth/pkg/repository"
)

// VerifyRole func for verifying a given role.
func VerifyRole(role string) (string, error) {
	// Switch given role.
	switch role {
	case repository.AdminRoleName:
		// Nothing to do, verified successfully.
	case repository.ModeratorRoleName:
		// Nothing to do, verified successfully.
	case repository.UserRoleName:
		// Nothing to do, verified successfully.
	default:
		// Return error message.
		return "", fmt.Errorf("role '%v' does not exist", role)
	}

	return role, nil
}
