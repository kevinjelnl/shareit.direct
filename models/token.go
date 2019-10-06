package models

// https://medium.com/mindorks/https-medium-com-yashishdua-synchronizing-states-using-mutex-vs-channel-in-go-25e646c83567

// Token is used for handshaking a receiver and a supplier
type Token struct {
	uuid string
}

// Token creates a token object and returns it to the receiver
func newToken() *Token {
	t := new(Token)
	return t
}
