// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/aptible/go-deploy/models"
)

// GetStacksStackIDVpcPeersReader is a Reader for the GetStacksStackIDVpcPeers structure.
type GetStacksStackIDVpcPeersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStacksStackIDVpcPeersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetStacksStackIDVpcPeersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetStacksStackIDVpcPeersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetStacksStackIDVpcPeersOK creates a GetStacksStackIDVpcPeersOK with default headers values
func NewGetStacksStackIDVpcPeersOK() *GetStacksStackIDVpcPeersOK {
	return &GetStacksStackIDVpcPeersOK{}
}

/*GetStacksStackIDVpcPeersOK handles this case with default header values.

successful
*/
type GetStacksStackIDVpcPeersOK struct {
	Payload *models.InlineResponse20041
}

func (o *GetStacksStackIDVpcPeersOK) Error() string {
	return fmt.Sprintf("[GET /stacks/{stack_id}/vpc_peers][%d] getStacksStackIdVpcPeersOK  %+v", 200, o.Payload)
}

func (o *GetStacksStackIDVpcPeersOK) GetPayload() *models.InlineResponse20041 {
	return o.Payload
}

func (o *GetStacksStackIDVpcPeersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponse20041)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStacksStackIDVpcPeersDefault creates a GetStacksStackIDVpcPeersDefault with default headers values
func NewGetStacksStackIDVpcPeersDefault(code int) *GetStacksStackIDVpcPeersDefault {
	return &GetStacksStackIDVpcPeersDefault{
		_statusCode: code,
	}
}

/*GetStacksStackIDVpcPeersDefault handles this case with default header values.

Error response. Often a 4xx or 5xx status code
*/
type GetStacksStackIDVpcPeersDefault struct {
	_statusCode int

	Payload *models.InlineResponseDefault
}

// Code gets the status code for the get stacks stack ID vpc peers default response
func (o *GetStacksStackIDVpcPeersDefault) Code() int {
	return o._statusCode
}

func (o *GetStacksStackIDVpcPeersDefault) Error() string {
	return fmt.Sprintf("[GET /stacks/{stack_id}/vpc_peers][%d] GetStacksStackIDVpcPeers default  %+v", o._statusCode, o.Payload)
}

func (o *GetStacksStackIDVpcPeersDefault) GetPayload() *models.InlineResponseDefault {
	return o.Payload
}

func (o *GetStacksStackIDVpcPeersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponseDefault)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
