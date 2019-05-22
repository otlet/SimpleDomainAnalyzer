package domain

import (
	"fmt"
	"net"
)

type Domain struct {
	Name string
}

func (dns Domain) CheckNameServer() {
	fmt.Println("Checking domain: " + dns.Name)
}

func (dns Domain) checkARecord(domainName string) {
	names, err := net.LookupHost(domainName)
	if err != nil {
		panic(err)
	}
	if len(names) == 0 {
		fmt.Printf("Records not found")
	}

	for _, name := range names {
		rev, err := net.LookupAddr(string(name))
		if err != nil {
			//panic(err)
		}
		fmt.Printf("%s , revDNS: %s\n", name, rev)
	}
}

func (dns Domain) checkMXRecord(domainName string) {
	mxs, err := net.LookupMX(domainName)
	if err != nil {
		fmt.Println("Domain with MX record")
	}
	for _, mx := range mxs {
		names, err := net.LookupHost(mx.Host)
		if err != nil {
			panic(err)
		}
		if len(names) == 0 {
			fmt.Printf("Records not found")
		}
		for _, name := range names {
			rev, err := net.LookupAddr(string(name))
			if err != nil {
				//panic(err)
			}
			fmt.Printf("MX: %s, PRIO: %v, IP: %s, revDNS: %s\n", mx.Host, mx.Pref, name, rev)
		}
	}
}

func (dns Domain) checkNSRecord(domainName string) {
	nss, err := net.LookupNS(domainName)
	if err != nil {
		fmt.Println("Domain without NS record?")
	}
	for _, ns := range nss {
		names, err := net.LookupHost(ns.Host)
		if err != nil {
			panic(err)
		}
		if len(names) == 0 {
			fmt.Printf("Records not found")
		}
		for _, name := range names {
			rev, err := net.LookupAddr(string(name))
			if err != nil {
				//panic(err)
			}
			fmt.Printf("NS: %s, IP: %s, revDNS: %s\n", ns.Host, name, rev)
		}
	}
}

func (dns Domain) checkTXTRecord(domainName string) {
	names, err := net.LookupTXT(domainName)
	if err != nil {
		panic(err)
	}
	if len(names) == 0 {
		fmt.Printf("Records not found")
	}

	for _, name := range names {
		fmt.Println("TXT: " + name)
	}
}

func (dns Domain) CheckRecords(domainName string) {
	fmt.Println("=== Rekord A ===")
	dns.checkARecord(domainName)

	fmt.Println("=== Rekord MX ===")
	dns.checkMXRecord(domainName)

	fmt.Println("=== Rekord NS ===")
	dns.checkNSRecord(domainName)

	fmt.Println("=== Rekord TXT ===")
	dns.checkTXTRecord(domainName)
}
