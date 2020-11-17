// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CloudConnectionEndpointGRE cloud connection endpoint g r e
// swagger:model CloudConnectionEndpointGRE
type CloudConnectionEndpointGRE struct {

	// enable gre for this cloud connection (default=false)
	Enabled bool `json:"enabled,omitempty"`

	// gre tunnels configured
	Tunnels []*CloudConnectionEndpointGRETunnelsItems0 `json:"tunnels"`
}

// Validate validates this cloud connection endpoint g r e
func (m *CloudConnectionEndpointGRE) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTunnels(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudConnectionEndpointGRE) validateTunnels(formats strfmt.Registry) error {

	if swag.IsZero(m.Tunnels) { // not required
		return nil
	}

	for i := 0; i < len(m.Tunnels); i++ {
		if swag.IsZero(m.Tunnels[i]) { // not required
			continue
		}

		if m.Tunnels[i] != nil {
			if err := m.Tunnels[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tunnels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CloudConnectionEndpointGRE) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloudConnectionEndpointGRE) UnmarshalBinary(b []byte) error {
	var res CloudConnectionEndpointGRE
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CloudConnectionEndpointGRETunnelsItems0 cloud connection endpoint g r e tunnels items0
// swagger:model CloudConnectionEndpointGRETunnelsItems0
type CloudConnectionEndpointGRETunnelsItems0 struct {

	// gre network in CIDR notation (192.168.0.0/24)
	// Required: true
	Cidr *string `json:"cidr"`

	// gre destination IP address
	// Required: true
	DestIPAddress *string `json:"destIPAddress"`

	// gre auto-assigned source IP address
	SourceIPAddress string `json:"sourceIPAddress,omitempty"`
}

// Validate validates this cloud connection endpoint g r e tunnels items0
func (m *CloudConnectionEndpointGRETunnelsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCidr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestIPAddress(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudConnectionEndpointGRETunnelsItems0) validateCidr(formats strfmt.Registry) error {

	if err := validate.Required("cidr", "body", m.Cidr); err != nil {
		return err
	}

	return nil
}

func (m *CloudConnectionEndpointGRETunnelsItems0) validateDestIPAddress(formats strfmt.Registry) error {

	if err := validate.Required("destIPAddress", "body", m.DestIPAddress); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CloudConnectionEndpointGRETunnelsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloudConnectionEndpointGRETunnelsItems0) UnmarshalBinary(b []byte) error {
	var res CloudConnectionEndpointGRETunnelsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}