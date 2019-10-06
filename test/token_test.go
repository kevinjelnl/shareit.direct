package test

import (
	"testing"

	"github.com/kevinjelnl/shareit.direct/models"
)

func TestNewToken(t *testing.T) {
	token := models.NewToken()
	t.Log(token)
}
