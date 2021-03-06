// Code generated by go-swagger; DO NOT EDIT.

package pet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/arcanericky/petstoreclient/models"
)

// FindPetsByStatusReader is a Reader for the FindPetsByStatus structure.
type FindPetsByStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindPetsByStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewFindPetsByStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewFindPetsByStatusBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewFindPetsByStatusOK creates a FindPetsByStatusOK with default headers values
func NewFindPetsByStatusOK() *FindPetsByStatusOK {
	return &FindPetsByStatusOK{}
}

/*FindPetsByStatusOK handles this case with default header values.

successful operation
*/
type FindPetsByStatusOK struct {
	Payload []*models.Pet
}

func (o *FindPetsByStatusOK) Error() string {
	return fmt.Sprintf("[GET /pet/findByStatus][%d] findPetsByStatusOK  %+v", 200, o.Payload)
}

func (o *FindPetsByStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindPetsByStatusBadRequest creates a FindPetsByStatusBadRequest with default headers values
func NewFindPetsByStatusBadRequest() *FindPetsByStatusBadRequest {
	return &FindPetsByStatusBadRequest{}
}

/*FindPetsByStatusBadRequest handles this case with default header values.

Invalid status value
*/
type FindPetsByStatusBadRequest struct {
}

func (o *FindPetsByStatusBadRequest) Error() string {
	return fmt.Sprintf("[GET /pet/findByStatus][%d] findPetsByStatusBadRequest ", 400)
}

func (o *FindPetsByStatusBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
