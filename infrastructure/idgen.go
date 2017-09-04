package infrastructure

import (
	"github.com/satori/go.uuid"
)

func GenID() string {
	return uuid.NewV4().String()
}
