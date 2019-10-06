package models

import (
	"log"
	"sync"

	gonanoid "github.com/matoous/go-nanoid"
)

// https://medium.com/mindorks/https-medium-com-yashishdua-synchronizing-states-using-mutex-vs-channel-in-go-25e646c83567
// https://gobyexample.com/mutexes

// Token is used for handshaking a receiver and a supplier
type Token struct {
	uuid string
}

const (
	tokenlenght int = 4
)

var tokensmap = make(map[string]Token)
var mutex = &sync.Mutex{}

// NewToken creates a token object and returns it to the receiver
func NewToken() *Token {
	t := new(Token)
	uid := t.generateToken()
	for !tokenExsists(uid) {
		t.uuid = uid
		break
	}
	return t
}

// generates the unique token
func (t Token) generateToken() string {
	utoken, err := gonanoid.Nanoid(tokenlenght)
	if err != nil {
		log.Fatal(err)
	}
	return utoken
}

// adds the token to the tokensmap
func (t Token) registerToken() {
	defer mutex.Unlock()
	mutex.Lock()
	tokensmap[t.uuid] = t
}

// removes the token from the tokensmap
func (t Token) removeToken() {
	defer mutex.Unlock()
	mutex.Lock()
	delete(tokensmap, t.uuid)
	return
}

// see if the token is unique
func tokenExsists(t string) bool {
	defer mutex.Unlock()
	mutex.Lock()
	_, exsists := tokensmap[t]
	return exsists
}
