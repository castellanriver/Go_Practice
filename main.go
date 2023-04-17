package main

import (
	"fmt"
	"main/accounts"
)

func main() {
	account := accounts.NewAccount("nico")
	fmt.Println(account)
}
