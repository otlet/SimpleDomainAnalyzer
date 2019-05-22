// Copyright 2019 Paweł Otlewski. All rights reserved.
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/fatih/color"
	"os"
	"os/signal"
	"superbok/domain"
	"superbok/http"
	"syscall"
)

func main() {

	SetupCloseHandler()

	if len(os.Args) < 2 {
		color.Red("Halo, podaj domenę!")
		os.Exit(1)
	}

	hello := os.Args[1]

	dns := domain.Domain{
		Name: hello,
	}

	color.Red("### Weryfikacja DNS ###")
	dns.CheckRecords(hello)

	strona := http.Http{
		hello,
	}

	fmt.Println()
	color.Red("### Weryfikacja Strony ###")
	strona.Run("http://", hello)
}

func SetupCloseHandler() {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\r- Ctrl+C <- Ale czemu mnie naciskasz?")
		os.Exit(0)
	}()
}
