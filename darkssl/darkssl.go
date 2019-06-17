package darkssl

import (
	"fmt"
	"net/http"
)

type DarkSSL struct {
	DomainName string
}

func (ssl DarkSSL) CheckSSL() {
	fmt.Println("Checking domain: " + ssl.DomainName)
	_, err := http.Get("https://" + ssl.DomainName)
	if err != nil {
		fmt.Println(err)
	} else {

	}
}
