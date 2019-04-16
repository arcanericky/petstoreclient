package main

import (
	"fmt"

	apiclient "github.com/arcanericky/petstoreclient/client"
	"github.com/arcanericky/petstoreclient/client/pet"
	"github.com/arcanericky/petstoreclient/models"
)

func addPet(id int64, petName string, urls []string) {
	newPet := models.Pet{ID: id, Name: &petName, PhotoUrls: urls}
	newPetParams := pet.NewAddPetParams()
	newPetParams.Body = &newPet

	fmt.Println("Adding pet")
	apiclient.Default.Pet.AddPet(newPetParams, nil)
}

func getPet(id int64) {
	getPetParams := pet.NewGetPetByIDParams()
	getPetParams.PetID = id

	fmt.Println("Fetching pet")
	// HELP HERE
	//
	//   The endpoint is specified at:
	//     https://petstore.swagger.io/#/pet/getPetById
	//
	//   Question 1
	//     This endpoint allows response types of "application/xml"
	//     and "application/json". How do you specify these?
	//
	//   Question 2
	//     How can the raw response body be retrieved? This response
	//     body can be in XML or JSON depending on #1. In both cases
	//     I need the raw response.
	if authOk, err := apiclient.Default.Pet.GetPetByID(getPetParams, nil); err == nil {
		fmt.Println(authOk)
		fmt.Printf("ID: %d\nName: %s\nPhotoUrls: %s\n",
			authOk.Payload.ID,
			*authOk.Payload.Name,
			authOk.Payload.PhotoUrls)
	} else {
		fmt.Println(err)
	}
}

func main() {
	const petID = 13371337

	addPet(petID, "geek", []string{"https://en.wikipedia.org/wiki/Geek"})
	getPet(petID)
}
