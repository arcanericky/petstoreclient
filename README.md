# Swagger Pet Store Client

### Purpose

Illustrate a request for help using [go-swagger](https://github.com/go-swagger/go-swagger) to:
- Change the header to `accept: application/json` or `accept: application/xml` to receive different responses
- Get the response body output (unparsed, raw XML or JSON)

The original example showed calls to `addPet()` and `getPet()` where the response body cannot be accessed.

This updated example shows calls to `addPetWithBody()` and `getPetWithBody()` using a custom [`RoundTripper`](https://golang.org/pkg/net/http/#RoundTripper) embedded in a custom go-swagger/go-openapi [`Runtime`](https://godoc.org/github.com/go-openapi/runtime/client#Runtime) to resolve the second item. This custom `RoundTripper`:
- calls the original `RoundTripper`
- extracts the body from the response
- replaces the body with a new reader containing original response contents

Are these 3 steps the ordained way to get at the body of a response with go-swagger? It seems to be a lot of code for something so simple.

[See `main.go`](https://github.com/arcanericky/petstoreclient/blob/master/main.go) for the code that uses the custom `RoundTripper` and `Runtime`. See the implementation of the custom code [in `customruntime.go`](https://github.com/arcanericky/petstoreclient/blob/master/customruntime.go)

The code works. Run it with `go run .`

### Reference

The API endpoint is live and can be referenced at https://petstore.swagger.io