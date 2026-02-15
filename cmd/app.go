package main

import (
	"fmt"
	"net/http"
	"time"
)

func check(link string, c chan string) {
	client := http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(link)
	if err != nil {
		c <- "🔴 " + link + " might be down"
		return
	}
	defer resp.Body.Close() 
	c <- "✅ " + link + " up and running"
}

func main() {
	start := time.Now() 
	links := []string{
		"https://google.com",
		"https://golang.org",
		"https://github.com",
	}

	c := make(chan string)

	for _, link := range links {
		go check(link, c)
	}

	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}

	fmt.Printf("\nDone🙂 Total time: %v\n", time.Since(start))
}