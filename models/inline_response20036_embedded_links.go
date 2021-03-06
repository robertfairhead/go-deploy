// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// InlineResponse20036EmbeddedLinks inline response 200 36 embedded links
// swagger:model inline_response_200_36__embedded__links
type InlineResponse20036EmbeddedLinks struct {

	// accounts
	Accounts *InlineResponse200EmbeddedEmbeddedLinksSelf `json:"accounts,omitempty"`

	// intrusion detection reports
	IntrusionDetectionReports *InlineResponse200EmbeddedEmbeddedLinksSelf `json:"intrusion_detection_reports,omitempty"`

	// organization
	Organization *InlineResponse200EmbeddedEmbeddedLinksSelf `json:"organization,omitempty"`

	// self
	Self *InlineResponse200EmbeddedEmbeddedLinksSelf `json:"self,omitempty"`

	// vpc peers
	VpcPeers *InlineResponse200EmbeddedEmbeddedLinksSelf `json:"vpc_peers,omitempty"`

	// vpn tunnels
	VpnTunnels *InlineResponse200EmbeddedEmbeddedLinksSelf `json:"vpn_tunnels,omitempty"`
}

// Validate validates this inline response 200 36 embedded links
func (m *InlineResponse20036EmbeddedLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccounts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIntrusionDetectionReports(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganization(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVpcPeers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVpnTunnels(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InlineResponse20036EmbeddedLinks) validateAccounts(formats strfmt.Registry) error {

	if swag.IsZero(m.Accounts) { // not required
		return nil
	}

	if m.Accounts != nil {
		if err := m.Accounts.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("accounts")
			}
			return err
		}
	}

	return nil
}

func (m *InlineResponse20036EmbeddedLinks) validateIntrusionDetectionReports(formats strfmt.Registry) error {

	if swag.IsZero(m.IntrusionDetectionReports) { // not required
		return nil
	}

	if m.IntrusionDetectionReports != nil {
		if err := m.IntrusionDetectionReports.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("intrusion_detection_reports")
			}
			return err
		}
	}

	return nil
}

func (m *InlineResponse20036EmbeddedLinks) validateOrganization(formats strfmt.Registry) error {

	if swag.IsZero(m.Organization) { // not required
		return nil
	}

	if m.Organization != nil {
		if err := m.Organization.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("organization")
			}
			return err
		}
	}

	return nil
}

func (m *InlineResponse20036EmbeddedLinks) validateSelf(formats strfmt.Registry) error {

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

func (m *InlineResponse20036EmbeddedLinks) validateVpcPeers(formats strfmt.Registry) error {

	if swag.IsZero(m.VpcPeers) { // not required
		return nil
	}

	if m.VpcPeers != nil {
		if err := m.VpcPeers.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vpc_peers")
			}
			return err
		}
	}

	return nil
}

func (m *InlineResponse20036EmbeddedLinks) validateVpnTunnels(formats strfmt.Registry) error {

	if swag.IsZero(m.VpnTunnels) { // not required
		return nil
	}

	if m.VpnTunnels != nil {
		if err := m.VpnTunnels.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vpn_tunnels")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InlineResponse20036EmbeddedLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InlineResponse20036EmbeddedLinks) UnmarshalBinary(b []byte) error {
	var res InlineResponse20036EmbeddedLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
