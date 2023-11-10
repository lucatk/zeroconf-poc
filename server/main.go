package main

import (
	"github.com/grandcat/zeroconf"
	"log"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
)

func main() {
	advertiseStr := os.Getenv("ADVERTISE")
	if advertiseStr == "" {
		log.Fatal("ADVERTISE not set")
	}

	advertise := strings.Split(advertiseStr, ",")

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	for _, svc := range advertise {
		svcSplit := strings.Split(svc, ":")
		svcName := svcSplit[0]
		svcPort, err := strconv.Atoi(svcSplit[1])
		if err != nil {
			panic(err)
		}

		log.Println("Advertising " + svcName + ":" + svcSplit[1])
		server, err := zeroconf.Register(svcName, "_theben._tcp", "local.", svcPort, nil, nil)
		if err != nil {
			panic(err)
		}
		defer server.Shutdown()
	}

	<-sigChan
	log.Println("Shutting down.")
}
