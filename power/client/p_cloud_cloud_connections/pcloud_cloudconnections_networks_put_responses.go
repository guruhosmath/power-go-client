// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_cloud_connections

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudCloudconnectionsNetworksPutReader is a Reader for the PcloudCloudconnectionsNetworksPut structure.
type PcloudCloudconnectionsNetworksPutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudCloudconnectionsNetworksPutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPcloudCloudconnectionsNetworksPutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 202:
		result := NewPcloudCloudconnectionsNetworksPutAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPcloudCloudconnectionsNetworksPutBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewPcloudCloudconnectionsNetworksPutUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewPcloudCloudconnectionsNetworksPutUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPcloudCloudconnectionsNetworksPutInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPcloudCloudconnectionsNetworksPutOK creates a PcloudCloudconnectionsNetworksPutOK with default headers values
func NewPcloudCloudconnectionsNetworksPutOK() *PcloudCloudconnectionsNetworksPutOK {
	return &PcloudCloudconnectionsNetworksPutOK{}
}

/*PcloudCloudconnectionsNetworksPutOK handles this case with default header values.

OK
*/
type PcloudCloudconnectionsNetworksPutOK struct {
	Payload *models.CloudConnection
}

func (o *PcloudCloudconnectionsNetworksPutOK) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}/networks/{network_id}][%d] pcloudCloudconnectionsNetworksPutOK  %+v", 200, o.Payload)
}

func (o *PcloudCloudconnectionsNetworksPutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CloudConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsNetworksPutAccepted creates a PcloudCloudconnectionsNetworksPutAccepted with default headers values
func NewPcloudCloudconnectionsNetworksPutAccepted() *PcloudCloudconnectionsNetworksPutAccepted {
	return &PcloudCloudconnectionsNetworksPutAccepted{}
}

/*PcloudCloudconnectionsNetworksPutAccepted handles this case with default header values.

Accepted
*/
type PcloudCloudconnectionsNetworksPutAccepted struct {
	Payload *models.JobReference
}

func (o *PcloudCloudconnectionsNetworksPutAccepted) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}/networks/{network_id}][%d] pcloudCloudconnectionsNetworksPutAccepted  %+v", 202, o.Payload)
}

func (o *PcloudCloudconnectionsNetworksPutAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobReference)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsNetworksPutBadRequest creates a PcloudCloudconnectionsNetworksPutBadRequest with default headers values
func NewPcloudCloudconnectionsNetworksPutBadRequest() *PcloudCloudconnectionsNetworksPutBadRequest {
	return &PcloudCloudconnectionsNetworksPutBadRequest{}
}

/*PcloudCloudconnectionsNetworksPutBadRequest handles this case with default header values.

Bad Request
*/
type PcloudCloudconnectionsNetworksPutBadRequest struct {
	Payload *models.Error
}

func (o *PcloudCloudconnectionsNetworksPutBadRequest) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}/networks/{network_id}][%d] pcloudCloudconnectionsNetworksPutBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudCloudconnectionsNetworksPutBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsNetworksPutUnauthorized creates a PcloudCloudconnectionsNetworksPutUnauthorized with default headers values
func NewPcloudCloudconnectionsNetworksPutUnauthorized() *PcloudCloudconnectionsNetworksPutUnauthorized {
	return &PcloudCloudconnectionsNetworksPutUnauthorized{}
}

/*PcloudCloudconnectionsNetworksPutUnauthorized handles this case with default header values.

Unauthorized
*/
type PcloudCloudconnectionsNetworksPutUnauthorized struct {
	Payload *models.Error
}

func (o *PcloudCloudconnectionsNetworksPutUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}/networks/{network_id}][%d] pcloudCloudconnectionsNetworksPutUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudCloudconnectionsNetworksPutUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsNetworksPutUnprocessableEntity creates a PcloudCloudconnectionsNetworksPutUnprocessableEntity with default headers values
func NewPcloudCloudconnectionsNetworksPutUnprocessableEntity() *PcloudCloudconnectionsNetworksPutUnprocessableEntity {
	return &PcloudCloudconnectionsNetworksPutUnprocessableEntity{}
}

/*PcloudCloudconnectionsNetworksPutUnprocessableEntity handles this case with default header values.

Unprocessable Entity
*/
type PcloudCloudconnectionsNetworksPutUnprocessableEntity struct {
	Payload *models.Error
}

func (o *PcloudCloudconnectionsNetworksPutUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}/networks/{network_id}][%d] pcloudCloudconnectionsNetworksPutUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PcloudCloudconnectionsNetworksPutUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsNetworksPutInternalServerError creates a PcloudCloudconnectionsNetworksPutInternalServerError with default headers values
func NewPcloudCloudconnectionsNetworksPutInternalServerError() *PcloudCloudconnectionsNetworksPutInternalServerError {
	return &PcloudCloudconnectionsNetworksPutInternalServerError{}
}

/*PcloudCloudconnectionsNetworksPutInternalServerError handles this case with default header values.

Internal Server Error
*/
type PcloudCloudconnectionsNetworksPutInternalServerError struct {
	Payload *models.Error
}

func (o *PcloudCloudconnectionsNetworksPutInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}/networks/{network_id}][%d] pcloudCloudconnectionsNetworksPutInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudCloudconnectionsNetworksPutInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
