package models

import (
	"testing"
)

func TestTokenModel(t *testing.T) {
	token := NewToken()
	// test if token length is being handled correct
	genTokenLen := len([]rune(token.Uuid))
	if genTokenLen != tokenlenght {
		t.Errorf("Tokenlenght: %d, tokenlenght should be: %d", genTokenLen, tokenlenght)
	}
}

func TestTokenExsists(t *testing.T) {
	token := NewToken()
	tokenExsists := tokenExsists(token.Uuid)
	if !tokenExsists {
		t.Log(tokensmap)
		t.Errorf("just created token: %s is not in the tokenmap", token.Uuid)
	}
}

func TestTokenRemoved(t *testing.T) {
	token := NewToken()
	token.removeToken()
	tokenExsists := tokenExsists(token.Uuid)
	if tokenExsists {
		t.Error("token not removed")
	}
}
