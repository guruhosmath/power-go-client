// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_v_p_n_policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudIpsecpoliciesGetReader is a Reader for the PcloudIpsecpoliciesGet structure.
type PcloudIpsecpoliciesGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudIpsecpoliciesGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPcloudIpsecpoliciesGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPcloudIpsecpoliciesGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewPcloudIpsecpoliciesGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewPcloudIpsecpoliciesGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewPcloudIpsecpoliciesGetUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPcloudIpsecpoliciesGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPcloudIpsecpoliciesGetOK creates a PcloudIpsecpoliciesGetOK with default headers values
func NewPcloudIpsecpoliciesGetOK() *PcloudIpsecpoliciesGetOK {
	return &PcloudIpsecpoliciesGetOK{}
}

/*PcloudIpsecpoliciesGetOK handles this case with default header values.

OK
*/
type PcloudIpsecpoliciesGetOK struct {
	Payload *models.IPSecPolicy
}

func (o *PcloudIpsecpoliciesGetOK) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn-connections/ipsec-policies/{ipsec_policy_id}][%d] pcloudIpsecpoliciesGetOK  %+v", 200, o.Payload)
}

func (o *PcloudIpsecpoliciesGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IPSecPolicy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudIpsecpoliciesGetBadRequest creates a PcloudIpsecpoliciesGetBadRequest with default headers values
func NewPcloudIpsecpoliciesGetBadRequest() *PcloudIpsecpoliciesGetBadRequest {
	return &PcloudIpsecpoliciesGetBadRequest{}
}

/*PcloudIpsecpoliciesGetBadRequest handles this case with default header values.

Bad Request
*/
type PcloudIpsecpoliciesGetBadRequest struct {
	Payload *models.Error
}

func (o *PcloudIpsecpoliciesGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn-connections/ipsec-policies/{ipsec_policy_id}][%d] pcloudIpsecpoliciesGetBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudIpsecpoliciesGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudIpsecpoliciesGetUnauthorized creates a PcloudIpsecpoliciesGetUnauthorized with default headers values
func NewPcloudIpsecpoliciesGetUnauthorized() *PcloudIpsecpoliciesGetUnauthorized {
	return &PcloudIpsecpoliciesGetUnauthorized{}
}

/*PcloudIpsecpoliciesGetUnauthorized handles this case with default header values.

Unauthorized
*/
type PcloudIpsecpoliciesGetUnauthorized struct {
	Payload *models.Error
}

func (o *PcloudIpsecpoliciesGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn-connections/ipsec-policies/{ipsec_policy_id}][%d] pcloudIpsecpoliciesGetUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudIpsecpoliciesGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudIpsecpoliciesGetForbidden creates a PcloudIpsecpoliciesGetForbidden with default headers values
func NewPcloudIpsecpoliciesGetForbidden() *PcloudIpsecpoliciesGetForbidden {
	return &PcloudIpsecpoliciesGetForbidden{}
}

/*PcloudIpsecpoliciesGetForbidden handles this case with default header values.

Forbidden
*/
type PcloudIpsecpoliciesGetForbidden struct {
	Payload *models.Error
}

func (o *PcloudIpsecpoliciesGetForbidden) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn-connections/ipsec-policies/{ipsec_policy_id}][%d] pcloudIpsecpoliciesGetForbidden  %+v", 403, o.Payload)
}

func (o *PcloudIpsecpoliciesGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudIpsecpoliciesGetUnprocessableEntity creates a PcloudIpsecpoliciesGetUnprocessableEntity with default headers values
func NewPcloudIpsecpoliciesGetUnprocessableEntity() *PcloudIpsecpoliciesGetUnprocessableEntity {
	return &PcloudIpsecpoliciesGetUnprocessableEntity{}
}

/*PcloudIpsecpoliciesGetUnprocessableEntity handles this case with default header values.

Unprocessable Entity
*/
type PcloudIpsecpoliciesGetUnprocessableEntity struct {
	Payload *models.Error
}

func (o *PcloudIpsecpoliciesGetUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn-connections/ipsec-policies/{ipsec_policy_id}][%d] pcloudIpsecpoliciesGetUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PcloudIpsecpoliciesGetUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudIpsecpoliciesGetInternalServerError creates a PcloudIpsecpoliciesGetInternalServerError with default headers values
func NewPcloudIpsecpoliciesGetInternalServerError() *PcloudIpsecpoliciesGetInternalServerError {
	return &PcloudIpsecpoliciesGetInternalServerError{}
}

/*PcloudIpsecpoliciesGetInternalServerError handles this case with default header values.

Internal Server Error
*/
type PcloudIpsecpoliciesGetInternalServerError struct {
	Payload *models.Error
}

func (o *PcloudIpsecpoliciesGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn-connections/ipsec-policies/{ipsec_policy_id}][%d] pcloudIpsecpoliciesGetInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudIpsecpoliciesGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
