// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Workload workload
//
// swagger:model workload
type Workload struct {

	// configmaps
	Configmaps ConfigmapList `json:"configmaps,omitempty"`

	// data
	Data *DataConfiguration `json:"data,omitempty"`

	// image registries
	ImageRegistries *ImageRegistries `json:"imageRegistries,omitempty"`

	// Log collection target for this workload
	LogCollection string `json:"log_collection,omitempty"`

	// metrics
	Metrics *Metrics `json:"metrics,omitempty"`

	// Name of the workload
	Name string `json:"name,omitempty"`

	// Namespace of the workload
	Namespace string `json:"namespace,omitempty"`

	// specification
	Specification string `json:"specification,omitempty"`
}

// Validate validates this workload
func (m *Workload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfigmaps(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImageRegistries(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetrics(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Workload) validateConfigmaps(formats strfmt.Registry) error {

	if swag.IsZero(m.Configmaps) { // not required
		return nil
	}

	if err := m.Configmaps.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("configmaps")
		}
		return err
	}

	return nil
}

func (m *Workload) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
	}

	if m.Data != nil {
		if err := m.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data")
			}
			return err
		}
	}

	return nil
}

func (m *Workload) validateImageRegistries(formats strfmt.Registry) error {

	if swag.IsZero(m.ImageRegistries) { // not required
		return nil
	}

	if m.ImageRegistries != nil {
		if err := m.ImageRegistries.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("imageRegistries")
			}
			return err
		}
	}

	return nil
}

func (m *Workload) validateMetrics(formats strfmt.Registry) error {

	if swag.IsZero(m.Metrics) { // not required
		return nil
	}

	if m.Metrics != nil {
		if err := m.Metrics.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metrics")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Workload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Workload) UnmarshalBinary(b []byte) error {
	var res Workload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
