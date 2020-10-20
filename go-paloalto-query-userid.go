package main

import (
	"log"

	pit "github.com/typester/go-pit"
)

func main() {
	profs, err := pit.Get("pa-xml-api", pit.Requires{"key": "", "host": ""})
	key, _ := (*profs)["key"]
	host, _ := (*profs)["host"]

	DisableTLSVerification(true)
	entries, err := GetIPUserMapping(host, key)
	DisableTLSVerification(false)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("entries: %d\n", len(entries))
	for idx, entry := range entries {
		log.Printf("  [%d] ip=%s\n", idx, entry.IpString)
		log.Printf("  [%d] user=%s\n", idx, entry.User)
		log.Printf("  [%d] timeout=%s\n", idx, entry.TimeoutString)
	}
}
