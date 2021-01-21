package lib

import (
	"log"
	"time"

	"github.com/patrickmn/go-cache"
)

// GDB handles the db conn and interactions
var GDB *cache.Cache

func init() {
	GDB = cache.New(5*time.Minute, 10*time.Minute)
	log.Println("...db initiated!")
}
