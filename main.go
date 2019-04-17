package main

import (
	"fmt"

	apiclient "github.com/arcanericky/petstoreclient/client"
	"github.com/arcanericky/petstoreclient/client/pet"
	"github.com/arcanericky/petstoreclient/models"
	"github.com/go-openapi/strfmt"
)

func addPet(id int64, petName string, urls []string) {
	newPet := models.Pet{ID: id, Name: &petName, PhotoUrls: urls}
	newPetParams := pet.NewAddPetParams()
	newPetParams.SetBody(&newPet)

	fmt.Println("Adding pet")
	apiclient.Default.Pet.AddPet(newPetParams, nil)
}

func getPet(id int64) {
	getPetParams := pet.NewGetPetByIDParams()
	getPetParams.PetID = id

	fmt.Println("Fetching pet")
	// See functions below to see my attempt at resolving question #2
	//
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

// My attempt to retrieve the response bodies by using a custom
// RoundTripper inside a custom Transport. The custom Roundtripper
// reads and stores the response, then replaces the original response
// with a new Reader. It works, but it is correct?
func addPetWithBody(id int64, petName string, urls []string) {
	newPet := models.Pet{ID: id, Name: &petName, PhotoUrls: urls}
	newPetParams := pet.NewAddPetParams()
	newPetParams.Body = &newPet

	runtime := newCustomRuntime()
	customClient := apiclient.New(runtime.runtime, strfmt.Default)

	fmt.Println("Adding pet")
	customClient.Pet.AddPet(newPetParams, nil)
	fmt.Println("Body:", runtime.body)
}

func getPetWithBody(id int64) {
	getPetParams := pet.NewGetPetByIDParams()
	getPetParams.PetID = id

	runtime := newCustomRuntime()
	customClient := apiclient.New(runtime.runtime, strfmt.Default)

	fmt.Println("Fetching pet")
	if _, err := customClient.Pet.GetPetByID(getPetParams, nil); err == nil {
		fmt.Println("Body:", runtime.body)
	} else {
		fmt.Println(err)
	}
}

func main() {
	const petID = 13371337

	addPetWithBody(petID, "geek", []string{"https://en.wikipedia.org/wiki/Geek"})
	getPetWithBody(petID)
}
