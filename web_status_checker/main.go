//channels exercise

package main

import (
	"fmt"
	"net/http"
	"time"
)


func main () {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)


	for _, link := range links {
		go checkLink(link, c)
	}

	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-c)
	// }

	// for { // infinite loop
	// 	go checkLink(<-c, c)
	// }

	// more readable version
	for l := range c {
		// OLD: not proper implementation
		// it's blocking the main routine for 5 seconds every time
		// time.Sleep(5 * time.Second) // 5 * time.Second for pausing 5 seconds
		// go checkLink(l, c)

		//function literal:
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
			fmt.Println("------")
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err	:= http.Get(link) 

	if err != nil {
		fmt.Println(link, "might be down!")
		// c <- "Might be down"
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	// c <- "It's up"
	c <- link
}
