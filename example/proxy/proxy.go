package main

import (
	"fmt"
	"github.com/aerotical/mehdi-client"
)

func main() {
	session := azuretls.NewSession()

	session.SetProxy("http://username:password@ip:port")

	response, err := session.Get("https://api.ipify.org")

	if err != nil {
		panic(err)
	}

	fmt.Println(response.StatusCode)
	fmt.Println(string(response.Body))

	session.Close()
}
