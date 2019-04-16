package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	apiclient "github.com/arcanericky/petstoreclient/client"
	"github.com/arcanericky/petstoreclient/client/pet"
	"github.com/arcanericky/petstoreclient/models"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

type customRoundTripper struct {
	defaultRoundTripper http.RoundTripper
	body                string
}

func (crt *customRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	resp, err := crt.defaultRoundTripper.RoundTrip(r)
	if err != nil {
		return nil, err
	}

	responseBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	crt.body = string(responseBytes)

	err = resp.Body.Close()
	if err != nil {
		return nil, err
	}

	resp.Body = ioutil.NopCloser(bytes.NewReader(responseBytes))

	return resp, nil
}

func newCustomTransport(crt *customRoundTripper) *httptransport.Runtime {
	transport := httptransport.New(apiclient.DefaultTransportConfig().Host,
		apiclient.DefaultTransportConfig().BasePath,
		apiclient.DefaultTransportConfig().Schemes)

	crt.defaultRoundTripper = transport.Transport
	transport.Transport = crt

	return transport
}

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

// My attempt to retrieve the response body by using a custom
// RoundTripper inside a custom Transport. The custom Roundtripper
// reads and stores the response, then replaces the original response
// with a new Reader. It works, but it is correct?
func getPetWithBody(id int64) {
	getPetParams := pet.NewGetPetByIDParams()
	getPetParams.PetID = id

	crt := new(customRoundTripper)
	client := apiclient.New(newCustomTransport(crt), strfmt.Default)

	fmt.Println("Fetching pet")
	if authOk, err := client.Pet.GetPetByID(getPetParams, nil); err == nil {
		fmt.Println(authOk)
		fmt.Printf("ID: %d\nName: %s\nPhotoUrls: %s\n",
			authOk.Payload.ID,
			*authOk.Payload.Name,
			authOk.Payload.PhotoUrls)
		fmt.Println("Body:", crt.body)
	} else {
		fmt.Println(err)
	}
}

func main() {
	const petID = 13371337

	addPet(petID, "geek", []string{"https://en.wikipedia.org/wiki/Geek"})
	getPetWithBody(petID)
}
