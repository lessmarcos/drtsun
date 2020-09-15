// Dr. Tsun checks if a given mailbox exists.
//
// Doing `go run dr_tsun.go example.com sales hr` checks
// if sales@example.com and hr@example.com exist.
// Achtung! Dr. Tsun may give false negatives (e.g., hr@example.com exists
// but Dr. Tsun cannot confirm).
package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/smtp"
	"os"
	"strings"
	"time"
)

func main() {
	log.SetFlags(0)

	args := os.Args[1:]

	if len(args) < 2 {
		log.Fatal("usage: dr_tsun username1 username2 ...")
	}

	host := args[0]
	usernames := args[1:]

	rand.Seed(time.Now().Unix())
	mxs, err := net.LookupMX(host); if err != nil {
		panic(err)
	}

	mx := mxs[rand.Intn(len(mxs))]

	log.Printf("Using MX host %s", mx.Host)

	c, err := smtp.Dial(fmt.Sprintf("%s:%d", mx.Host, 25)); if err != nil {
		panic(err)
	}

	err = c.Hello("somename.com"); if err != nil {
		panic(err)
	}

	if err := c.Mail("drtsun@somename.com"); err != nil {
		log.Fatal(err)
	}

	// Discard catch-all alias mailboxes
	if err := c.Rcpt(fmt.Sprintf("nOnExisTenTmAilBox3@%s", host)); err == nil {
		log.Fatalf("%s is using catch-all email alias", host)
	}

	if err := c.Rcpt(fmt.Sprintf("AnoThErnOnExisTenTmAilBox3@%s", host)); err != nil {
		// Prevent greylisting by omitting the "first" error
	}

	for _, username := range usernames {
		u := strings.ToLower(strings.ReplaceAll(username, " ", "."))
		log.Printf("Trying %s@%s", u, host)
		if err = c.Rcpt(fmt.Sprintf("%s@%s", u, host)); err != nil {
			continue
		}

		log.Printf("\t%s@%s exists\n", u, host)
	}
}
