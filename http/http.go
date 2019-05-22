package http

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type Http struct {
	Name string // Domain name
}

func (httpConnect Http) Run(scheme string, domainName string) {
	name := scheme + domainName
	resp, err := http.Get(name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Kod HTTP:", resp.StatusCode, http.StatusText(resp.StatusCode))
	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		fmt.Println("Strona powinna śmigać")
	} else if resp.StatusCode >= 500 && resp.StatusCode <= 599 {
		fmt.Println("Nie działa strona")
	} else if resp.StatusCode >= 300 && resp.StatusCode <= 399 {
		if scheme == "https://" {
			fmt.Println("Powstała pętla przekierowań - wina klienta")
			os.Exit(0)
		}
		tmp := resp.Header.Get("Location")
		fmt.Printf("Przekierowanie na %s\n", tmp)

		//httpConnect.Run("https://", domainName)
	} else {
		fmt.Println("Strona nie działa, bo zapewne jej nie ma. Sprawdź wskazania, blokadę oraz czy jego pliki istnieją")
	}
}