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

// InlineResponse20025 inline response 200 25
// swagger:model inline_response_200_25
type InlineResponse20025 struct {

	// resource type
	// Required: true
	ResourceType *string `json:"_type"`

	// links
	// Required: true
	Links *InlineResponse2003EmbeddedEmbeddedCurrentImageLinks `json:"_links"`

	// created at
	// Required: true
	CreatedAt *string `json:"created_at"`

	// docker ref
	// Required: true
	DockerRef *string `json:"docker_ref"`

	// docker repo
	// Required: true
	DockerRepo *string `json:"docker_repo"`

	// dualstack hint
	// Required: true
	DualstackHint *int64 `json:"dualstack_hint"`

	// exposed ports
	// Required: true
	ExposedPorts []int64 `json:"exposed_ports"`

	// git ref
	// Required: true
	GitRef *string `json:"git_ref"`

	// git repo
	// Required: true
	GitRepo *string `json:"git_repo"`

	// id
	// Required: true
	ID *int64 `json:"id"`

	// platform
	// Required: true
	Platform *string `json:"platform"`

	// release
	// Required: true
	Release *string `json:"release"`

	// scan
	// Required: true
	Scan *string `json:"scan"`

	// updated at
	// Required: true
	UpdatedAt *string `json:"updated_at"`
}

// Validate validates this inline response 200 25
func (m *InlineResponse20025) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResourceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDockerRef(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDockerRepo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDualstackHint(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExposedPorts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGitRef(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGitRepo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlatform(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRelease(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScan(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InlineResponse20025) validateResourceType(formats strfmt.Registry) error {

	if err := validate.Required("_type", "body", m.ResourceType); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse20025) validateLinks(formats strfmt.Registry) error {

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

func (m *InlineResponse20025) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("created_at", "body", m.CreatedAt); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse20025) validateDockerRef(formats strfmt.Registry) error {

	if err := validate.Required("docker_ref", "body", m.DockerRef); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse20025) validateDockerRepo(formats strfmt.Registry) error {

	if err := validate.Required("docker_repo", "body", m.DockerRepo); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse20025) validateDualstackHint(formats strfmt.Registry) error {

	if err := validate.Required("dualstack_hint", "body", m.DualstackHint); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse20025) validateExposedPorts(formats strfmt.Registry) error {

	if err := validate.Required("exposed_ports", "body", m.ExposedPorts); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse20025) validateGitRef(formats strfmt.Registry) error {

	if err := validate.Required("git_ref", "body", m.GitRef); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse20025) validateGitRepo(formats strfmt.Registry) error {

	if err := validate.Required("git_repo", "body", m.GitRepo); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse20025) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse20025) validatePlatform(formats strfmt.Registry) error {

	if err := validate.Required("platform", "body", m.Platform); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse20025) validateRelease(formats strfmt.Registry) error {

	if err := validate.Required("release", "body", m.Release); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse20025) validateScan(formats strfmt.Registry) error {

	if err := validate.Required("scan", "body", m.Scan); err != nil {
		return err
	}

	return nil
}

func (m *InlineResponse20025) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updated_at", "body", m.UpdatedAt); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InlineResponse20025) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InlineResponse20025) UnmarshalBinary(b []byte) error {
	var res InlineResponse20025
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
