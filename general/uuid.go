package general

import (
	"fmt"

	"github.com/google/uuid"
)

func MakeUUID() (string, error) {
	u1, err := uuid.NewUUID()
	if err != nil {
		return "", fmt.Errorf("make uuid: falied: %w", err)
	}
	return u1.String(), err
}
