// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// InlineResponse2007Embedded inline response 200 7 embedded
// swagger:model inline_response_200_7__embedded
type InlineResponse2007Embedded struct {

	// configurations
	Configurations []*InlineResponse2007EmbeddedConfigurations `json:"configurations"`
}

// Validate validates this inline response 200 7 embedded
func (m *InlineResponse2007Embedded) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfigurations(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InlineResponse2007Embedded) validateConfigurations(formats strfmt.Registry) error {

	if swag.IsZero(m.Configurations) { // not required
		return nil
	}

	for i := 0; i < len(m.Configurations); i++ {
		if swag.IsZero(m.Configurations[i]) { // not required
			continue
		}

		if m.Configurations[i] != nil {
			if err := m.Configurations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("configurations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *InlineResponse2007Embedded) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InlineResponse2007Embedded) UnmarshalBinary(b []byte) error {
	var res InlineResponse2007Embedded
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
