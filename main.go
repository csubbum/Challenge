package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var SPELL_CHECK_SUBSCRIPTION_KEY = flag.String("spellCheckKey", "", "Used for spell check bing api")

func main() {
	flag.Parse()
	if *SPELL_CHECK_SUBSCRIPTION_KEY == "" {
		log.Println("Please provide spell check subscription key.")
		log.Panic("Syntax: challenge -spellCheckKey=<key>")
	}

	srv := startServer()
	var gracefulStop = make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)
	<-gracefulStop
	wait, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	err := srv.Shutdown(wait)
	if err != nil {
		panic(err)
	}
}
