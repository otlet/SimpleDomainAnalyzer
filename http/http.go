package http

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type Http struct{}

func (httpConnect Http) Run(scheme string, domainName string) {
	name := scheme + domainName
	resp, err := http.Get(name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("HTTP Code:", resp.StatusCode, http.StatusText(resp.StatusCode))
	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		fmt.Println("Website work properly")
	} else if resp.StatusCode >= 500 && resp.StatusCode <= 599 {
		fmt.Println("Server Error")
	} else if resp.StatusCode >= 300 && resp.StatusCode <= 399 {
		if scheme == "https://" {
			fmt.Println("Redirect loop error")
			os.Exit(0)
		}
		tmp := resp.Header.Get("Location")
		fmt.Printf("Redirect: %s\n", tmp)

		//httpConnect.Run("https://", domainName)
	} else {
		fmt.Println("Website not found. Check ban and website existing")
	}
}
