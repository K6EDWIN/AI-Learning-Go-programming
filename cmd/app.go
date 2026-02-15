package main

import (
	"fmt"
	"net/http"
	"time"
)

func check(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		c <- link + " might be down😪"
		return
	}
	c <- link + " is up and running!✅"
}

func main() {
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
	fmt.Printf("Done🙂took %v\n", time.Now().Format("15:04:05"))
}