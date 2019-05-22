// Copyright 2019 Paweł Otlewski. All rights reserved.
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/fatih/color"
	"os"
	"os/signal"
	"regexp"
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

	name := os.Args[1]

	domainRegExp := `^(([a-zA-Z]{1})|([a-zA-Z]{1}[a-zA-Z]{1})|([a-zA-Z]{1}[0-9]{1})|([0-9]{1}[a-zA-Z]{1})|([a-zA-Z0-9][a-zA-Z0-9-_]{1,61}[a-zA-Z0-9]))\.([a-zA-Z]{2,6}|[a-zA-Z0-9-]{2,30}\.[a-zA-Z
 ]{2,3})$`
	RegExp := regexp.MustCompile(domainRegExp)

	validated := RegExp.MatchString(name)

	if !validated {
		color.Red("Podaj domenę, a nie jakieś: %s", name)
		os.Exit(1)
	}

	dns := domain.Domain{
		Name: name,
	}

	color.Red("### Weryfikacja DNS ###")
	dns.CheckRecords(name)

	strona := http.Http{}

	fmt.Println()
	color.Red("### Weryfikacja Strony ###")
	strona.Run("http://", name)
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
