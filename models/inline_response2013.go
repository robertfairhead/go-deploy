// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// InlineResponse2013 inline response 201 3
// swagger:model inline_response_201_3
type InlineResponse2013 struct {

	// resource type
	// Required: true
	ResourceType *string `json:"_type"`

	// links
	// Required: true
	Links *InlineResponse2007EmbeddedLinks `json:"_links"`

	// env
	// Required: true
	Env interface{} `json:"env"`

	// id
	// Required: true
	ID *int64 `json:"id"`
}

// Validate validates this inline response 201 3
func (m *InlineResponse2013) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResourceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnv(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InlineResponse2013) validateResourceType(formats strfmt.Registry) error {

	if err := validate.Required("_type", "body", m.ResourceType); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse2013) validateLinks(formats strfmt.Registry) error {

	if err := validate.Required("_links", "body", m.Links); err != nil {
		return err
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (m *InlineResponse2013) validateEnv(formats strfmt.Registry) error {

	if err := validate.Required("env", "body", m.Env); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse2013) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InlineResponse2013) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InlineResponse2013) UnmarshalBinary(b []byte) error {
	var res InlineResponse2013
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}