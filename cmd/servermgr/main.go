package main

import (
	"math/rand"
	"time"

	"github.com/logeable/serverguard/internal/servermgr"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	servermgr.NewApp().Run()
}
