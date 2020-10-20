package main

import (
	"flag"

	pit "github.com/typester/go-pit"
)

func main() {
	var timeoutMin int
	var user string
	var ipaddr string
	flag.IntVar(&timeoutMin, "timeout", 180, "timeout [min]")
	flag.StringVar(&user, "user", "testuser", "username")
	flag.StringVar(&ipaddr, "ipaddr", "0.0.0.0", "IP Address")
	isLogout := flag.Bool("logout", false, "Set if logout")
	flag.Parse()

	profs, _ := pit.Get("pa-xml-api", pit.Requires{"key": "", "host": ""})
	key, _ := (*profs)["key"]
	host, _ := (*profs)["host"]

	DisableTLSVerification(true)
	if *isLogout == false {
		LoginUserForIP(user, ipaddr, timeoutMin, host, key)
	} else {
		LogoutUserForIP(user, ipaddr, host, key)
	}
}
