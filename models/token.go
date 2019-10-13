package models

import (
	"log"
	"sync"

	gonanoid "github.com/matoous/go-nanoid"
)

// Token is used for handshaking a receiver and a supplier
type Token struct {
	Uuid string
}

const (
	tokenlenght int = 4
)

var tokensmap = make(map[string]Receiver)
var mutex = &sync.Mutex{}

// NewToken creates a token object and returns it to the receiver
func NewToken() *Token {
	t := new(Token)
	uid := t.generateToken()
	for !tokenExsists(uid) {
		t.Uuid = uid
		// t.registerToken()
		break
	}
	return t
}

// generates the unique token
func (t Token) generateToken() string {
	utoken, err := gonanoid.Nanoid(tokenlenght)
	if err != nil {
		log.Print(err)
	}
	return utoken
}

// adds the token to the tokensmap
func (t Token) registerToken(r Receiver) {
	defer mutex.Unlock()
	mutex.Lock()
	tokensmap[t.Uuid] = r
}

// RemoveToken removes the token from the tokensmap
func RemoveToken(t string) {
	defer mutex.Unlock()
	mutex.Lock()
	delete(tokensmap, t)
	return
}

// see if the token is unique
func tokenExsists(t string) bool {
	defer mutex.Unlock()
	mutex.Lock()
	_, exsists := tokensmap[t]
	return exsists
}
