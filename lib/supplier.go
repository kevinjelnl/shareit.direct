package lib

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
)

// Secret value
type Secret struct {
	Token     string
	Value     string
	Keyphrase bool
}

// tokenExists finds the receiver based on the token
func (s Secret) tokenExists() (r *Receiver, err error) {
	tokenUpper := strings.ToUpper(s.Token)
	stored, exists := GDB.Get(tokenUpper)
	if !exists {
		err := fmt.Errorf("Token not found")
		return nil, err
	}
	r = *stored.(**Receiver)
	return r, nil
}

// SupplierHandler provides the secret to the user
func SupplierHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	header := w.Header()
	header.Set("Content-Type", "application/json")
	header.Set("Access-Control-Allow-Methods", header.Get("Allow"))
	header.Set("Access-Control-Allow-Origin", "*")

	var s Secret
	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		panic(err)
	}

	receiver, err := s.tokenExists()
	if err != nil {
		resp, _ := json.Marshal("Supplied token invalid!")
		w.Write(resp)
		return
	}
	// ship the secret
	go receiver.SendSecret(s)
	resp, _ := json.Marshal("Secret shared!")
	w.Write(resp)
	return
}
