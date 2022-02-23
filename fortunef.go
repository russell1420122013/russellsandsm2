package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
	"time"
)

func fortune(msg chan bool) {
	byteArray, err := ioutil.ReadFile("Fortunes.txt")
	if err != nil {
		log.Fatal(err)
	}
	str := string(byteArray)
	slcOStr := strings.Split(str, "%%")
	for true {
		<-msg
		rand.Seed(time.Now().UnixNano())
		rndnum := rand.Intn(len(slcOStr) - 1)
		fmt.Println(slcOStr[rndnum])
	}
}

func main() {
	msg := make(chan bool)
	go fortune(msg)
	for true {
		var response string
		fmt.Println("Would you like another fortune? ")
		fmt.Scan(&response)
		if response == "yes" || response == "YES" || response == "Yes" {
			msg <- true
			time.Sleep(time.Second * 1)
		} else if response == "no" || response == "NO" || response == "No" {
			break
		} else {
			fmt.Println()
		}
	}
}
