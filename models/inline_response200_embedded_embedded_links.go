// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// InlineResponse200EmbeddedEmbeddedLinks inline response 200 embedded embedded links
// swagger:model inline_response_200__embedded__embedded__links
type InlineResponse200EmbeddedEmbeddedLinks struct {

	// account
	Account *InlineResponse200EmbeddedEmbeddedLinksAccount `json:"account,omitempty"`

	// containers
	Containers *InlineResponse200EmbeddedEmbeddedLinksAccount `json:"containers,omitempty"`

	// database
	Database *InlineResponse200EmbeddedEmbeddedLinksAccount `json:"database,omitempty"`

	// ephemeral containers
	EphemeralContainers *InlineResponse200EmbeddedEmbeddedLinksAccount `json:"ephemeral_containers,omitempty"`

	// operations
	Operations *InlineResponse200EmbeddedEmbeddedLinksAccount `json:"operations,omitempty"`

	// self
	Self *InlineResponse200EmbeddedEmbeddedLinksAccount `json:"self,omitempty"`
}

// Validate validates this inline response 200 embedded embedded links
func (m *InlineResponse200EmbeddedEmbeddedLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContainers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDatabase(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEphemeralContainers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InlineResponse200EmbeddedEmbeddedLinks) validateAccount(formats strfmt.Registry) error {

	if swag.IsZero(m.Account) { // not required
		return nil
	}

	if m.Account != nil {
		if err := m.Account.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("account")
			}
			return err
		}
	}

	return nil
}

func (m *InlineResponse200EmbeddedEmbeddedLinks) validateContainers(formats strfmt.Registry) error {

	if swag.IsZero(m.Containers) { // not required
		return nil
	}

	if m.Containers != nil {
		if err := m.Containers.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("containers")
			}
			return err
		}
	}

	return nil
}

func (m *InlineResponse200EmbeddedEmbeddedLinks) validateDatabase(formats strfmt.Registry) error {

	if swag.IsZero(m.Database) { // not required
		return nil
	}

	if m.Database != nil {
		if err := m.Database.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("database")
			}
			return err
		}
	}

	return nil
}

func (m *InlineResponse200EmbeddedEmbeddedLinks) validateEphemeralContainers(formats strfmt.Registry) error {

	if swag.IsZero(m.EphemeralContainers) { // not required
		return nil
	}

	if m.EphemeralContainers != nil {
		if err := m.EphemeralContainers.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ephemeral_containers")
			}
			return err
		}
	}

	return nil
}

func (m *InlineResponse200EmbeddedEmbeddedLinks) validateOperations(formats strfmt.Registry) error {

	if swag.IsZero(m.Operations) { // not required
		return nil
	}

	if m.Operations != nil {
		if err := m.Operations.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("operations")
			}
			return err
		}
	}

	return nil
}

func (m *InlineResponse200EmbeddedEmbeddedLinks) validateSelf(formats strfmt.Registry) error {

	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InlineResponse200EmbeddedEmbeddedLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InlineResponse200EmbeddedEmbeddedLinks) UnmarshalBinary(b []byte) error {
	var res InlineResponse200EmbeddedEmbeddedLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
