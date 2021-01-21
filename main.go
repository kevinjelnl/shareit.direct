package main

import (
	"fmt"
	"log"
	"net/http"

	"./lib"

	"github.com/julienschmidt/httprouter"
)

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	header := w.Header()
	header.Set("Access-Control-Allow-Methods", header.Get("Allow"))
	header.Set("Access-Control-Allow-Origin", "*")
	// header := w.Header()
	fmt.Println("...index hit")
	amount := lib.GDB.ItemCount()
	log.Printf("%+v", amount)
	return
}

func main() {
	log.Println(("...server started."))
	router := httprouter.New()
	router.GET("/", index)
	router.GET("/ws", lib.WebsocketHandler)
	router.POST("/supply", lib.SupplierHandler)
	// gdb.Set("foo", "bar", cache.DefaultExpiration)
	// fmt.Println(gdb.Get("None"))
	// fmt.Println(gdb.Get("foo"))
	log.Fatal(http.ListenAndServe(":8080", router))
}
