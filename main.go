package main

import "fmt"

func main() {
	client := NewClient()
	token := client.BearerToken()

	fmt.Printf("\n\n\n=======\n\n\n This is your BearerToken: %s \n\n\n======\n\n\n", token.AccessToken)
}
