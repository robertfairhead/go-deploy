// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// InlineResponse2003EmbeddedLinks inline response 200 3 embedded links
// swagger:model inline_response_200_3__embedded__links
type InlineResponse2003EmbeddedLinks struct {

	// account
	Account *InlineResponse200EmbeddedEmbeddedLinksAccount `json:"account,omitempty"`

	// configurations
	Configurations *InlineResponse200EmbeddedEmbeddedLinksAccount `json:"configurations,omitempty"`

	// current configuration
	CurrentConfiguration *InlineResponse200EmbeddedEmbeddedLinksAccount `json:"current_configuration,omitempty"`

	// current image
	CurrentImage *InlineResponse200EmbeddedEmbeddedLinksAccount `json:"current_image,omitempty"`

	// ephemeral sessions
	EphemeralSessions *InlineResponse200EmbeddedEmbeddedLinksAccount `json:"ephemeral_sessions,omitempty"`

	// images
	Images *InlineResponse200EmbeddedEmbeddedLinksAccount `json:"images,omitempty"`

	// operations
	Operations *InlineResponse200EmbeddedEmbeddedLinksAccount `json:"operations,omitempty"`

	// self
	Self *InlineResponse200EmbeddedEmbeddedLinksAccount `json:"self,omitempty"`

	// services
	Services *InlineResponse200EmbeddedEmbeddedLinksAccount `json:"services,omitempty"`

	// vhosts
	Vhosts *InlineResponse200EmbeddedEmbeddedLinksAccount `json:"vhosts,omitempty"`
}

// Validate validates this inline response 200 3 embedded links
func (m *InlineResponse2003EmbeddedLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConfigurations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCurrentConfiguration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCurrentImage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEphemeralSessions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImages(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServices(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVhosts(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InlineResponse2003EmbeddedLinks) validateAccount(formats strfmt.Registry) error {

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

func (m *InlineResponse2003EmbeddedLinks) validateConfigurations(formats strfmt.Registry) error {

	if swag.IsZero(m.Configurations) { // not required
		return nil
	}

	if m.Configurations != nil {
		if err := m.Configurations.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("configurations")
			}
			return err
		}
	}

	return nil
}

func (m *InlineResponse2003EmbeddedLinks) validateCurrentConfiguration(formats strfmt.Registry) error {

	if swag.IsZero(m.CurrentConfiguration) { // not required
		return nil
	}

	if m.CurrentConfiguration != nil {
		if err := m.CurrentConfiguration.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("current_configuration")
			}
			return err
		}
	}

	return nil
}

func (m *InlineResponse2003EmbeddedLinks) validateCurrentImage(formats strfmt.Registry) error {

	if swag.IsZero(m.CurrentImage) { // not required
		return nil
	}

	if m.CurrentImage != nil {
		if err := m.CurrentImage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("current_image")
			}
			return err
		}
	}

	return nil
}

func (m *InlineResponse2003EmbeddedLinks) validateEphemeralSessions(formats strfmt.Registry) error {

	if swag.IsZero(m.EphemeralSessions) { // not required
		return nil
	}

	if m.EphemeralSessions != nil {
		if err := m.EphemeralSessions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ephemeral_sessions")
			}
			return err
		}
	}

	return nil
}

func (m *InlineResponse2003EmbeddedLinks) validateImages(formats strfmt.Registry) error {

	if swag.IsZero(m.Images) { // not required
		return nil
	}

	if m.Images != nil {
		if err := m.Images.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("images")
			}
			return err
		}
	}

	return nil
}

func (m *InlineResponse2003EmbeddedLinks) validateOperations(formats strfmt.Registry) error {

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

func (m *InlineResponse2003EmbeddedLinks) validateSelf(formats strfmt.Registry) error {

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

func (m *InlineResponse2003EmbeddedLinks) validateServices(formats strfmt.Registry) error {

	if swag.IsZero(m.Services) { // not required
		return nil
	}

	if m.Services != nil {
		if err := m.Services.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("services")
			}
			return err
		}
	}

	return nil
}

func (m *InlineResponse2003EmbeddedLinks) validateVhosts(formats strfmt.Registry) error {

	if swag.IsZero(m.Vhosts) { // not required
		return nil
	}

	if m.Vhosts != nil {
		if err := m.Vhosts.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vhosts")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InlineResponse2003EmbeddedLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InlineResponse2003EmbeddedLinks) UnmarshalBinary(b []byte) error {
	var res InlineResponse2003EmbeddedLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
