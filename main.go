package main

import (
	"fmt"

	apiclient "github.com/arcanericky/petstoreclient/client"
)

func main() {
	fmt.Println("Fetching Inventory")

	authOk, err := apiclient.Default.Store.GetInventory(nil, nil)

	if err == nil {
		fmt.Println(authOk)
	} else {
		fmt.Println(err)
	}
}
