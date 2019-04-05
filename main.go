package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	subdomain := os.Getenv("DUCKDNS_SUBDOMAIN")
	token := os.Getenv("DUCKDNS_TOKEN")
	interval, _ := strconv.Atoi(getEnv("DUCKDNS_REFRESH_INTERVAL", "0"))

	if interval >= 60 {
		for {
			updateDuckDNS(subdomain, token)
			log.Printf("Sleeping for %d seconds", interval)
			time.Sleep(time.Duration(interval) * time.Second)
		}
	} else {
		updateDuckDNS(subdomain, token)
	}
}

func updateDuckDNS(subdomain string, token string) {
	requestURL := fmt.Sprintf(
		"https://www.duckdns.org/update?domains=%s&token=%s",
		subdomain,
		token)

	log.Printf("Updating %s.duckdns.org\n", subdomain)
	resp, err := http.Get(requestURL)

	if err != nil {
		log.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if strings.Contains(string(body), "KO") {
		log.Println("An error occurred, exiting...")
		os.Exit(1)
	} else if strings.Contains(string(body), "OK") {
		log.Printf("Successfully updated %s.duckdns.org", subdomain)
	}
}

// Simple helper function to read an environment or return a default value
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
