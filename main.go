package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func main() {
	target := "https://notify-api.line.me/api/notify"
	token := "token" //<> invalid

	var msg string

	firstUserMsg := "user1 약드셈"
	secondUserMsg := "user2 약드셈"

	arg := strings.Join(os.Args[1:2], "")

	fmt.Println(arg)

	if arg == "user1" {
		msg = firstUserMsg
	} else if arg == "user2" {
		msg = secondUserMsg
	} else {
		panic("invalid bald's name")
	}

	u, err := url.ParseRequestURI(target)
	if err != nil {
		log.Fatal(err)
	}

	c := &http.Client{}

	form := url.Values{}
	form.Add("message", msg)

	body := strings.NewReader(form.Encode())

	req, err := http.NewRequest("POST", u.String(), body)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := c.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	bytes, _ := ioutil.ReadAll(resp.Body)
	str := string(bytes)
	fmt.Println(str)
}
