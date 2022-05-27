// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"

	"github.com/project-flotta/flotta-operator/models"
)

// NewEnrolDeviceParams creates a new EnrolDeviceParams object
// no default values defined in spec.
func NewEnrolDeviceParams() EnrolDeviceParams {

	return EnrolDeviceParams{}
}

// EnrolDeviceParams contains all the bound params for the enrol device operation
// typically these are obtained from a http.Request
//
// swagger:parameters EnrolDevice
type EnrolDeviceParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Device ID
	  Required: true
	  In: path
	*/
	DeviceID string
	/*
	  Required: true
	  In: body
	*/
	EnrolmentInfo *models.EnrolmentInfo
	/*Namespace where the device resides
	  Required: true
	  In: path
	*/
	Namespace string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewEnrolDeviceParams() beforehand.
func (o *EnrolDeviceParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rDeviceID, rhkDeviceID, _ := route.Params.GetOK("device-id")
	if err := o.bindDeviceID(rDeviceID, rhkDeviceID, route.Formats); err != nil {
		res = append(res, err)
	}

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.EnrolmentInfo
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("enrolmentInfo", "body", ""))
			} else {
				res = append(res, errors.NewParseError("enrolmentInfo", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.EnrolmentInfo = &body
			}
		}
	} else {
		res = append(res, errors.Required("enrolmentInfo", "body", ""))
	}
	rNamespace, rhkNamespace, _ := route.Params.GetOK("namespace")
	if err := o.bindNamespace(rNamespace, rhkNamespace, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindDeviceID binds and validates parameter DeviceID from path.
func (o *EnrolDeviceParams) bindDeviceID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.DeviceID = raw

	return nil
}

// bindNamespace binds and validates parameter Namespace from path.
func (o *EnrolDeviceParams) bindNamespace(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.Namespace = raw

	return nil
}
