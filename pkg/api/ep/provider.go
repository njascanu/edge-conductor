// Code generated by go-swagger; DO NOT EDIT.

//
//   Copyright (c) 2022 Intel Corporation.
//
//   SPDX-License-Identifier: Apache-2.0
//
//
//

package ep

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Provider provider
//
// swagger:model provider
type Provider []*ProviderItems0

// Validate validates this provider
func (m Provider) Validate(formats strfmt.Registry) error {
	var res []error

	for i := 0; i < len(m); i++ {
		if swag.IsZero(m[i]) { // not required
			continue
		}

		if m[i] != nil {
			if err := m[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName(strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName(strconv.Itoa(i))
				}
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this provider based on the context it is used
func (m Provider) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	for i := 0; i < len(m); i++ {

		if m[i] != nil {
			if err := m[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName(strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName(strconv.Itoa(i))
				}
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ProviderItems0 provider items0
//
// swagger:model ProviderItems0
type ProviderItems0 struct {

	// name
	// Enum: [cluster-api kubeadm byoh metal3]
	Name string `json:"name,omitempty"`

	// parameters
	Parameters *ProviderItems0Parameters `json:"parameters,omitempty"`

	// provider type
	// Enum: [CoreProvider BootstrapProvider ControlPlaneProvider InfrastructureProvider]
	ProviderType string `json:"provider_type,omitempty"`

	// url
	// Pattern: (?:(?:https?|http|ftp|file|oci)://|www.|ftp.)(?:([-A-Z0-9+&@#/%=~_|$?!:,.]*)|[-A-Z0-9+&@#/%=~_|$?!:,.])*(?:([-A-Z0-9+&@#/%=~_|$?!:,.]*)|[A-Z0-9+&@#/%=~_|$])
	URL string `json:"url,omitempty"`
}

// Validate validates this provider items0
func (m *ProviderItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParameters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProviderType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var providerItems0TypeNamePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["cluster-api","kubeadm","byoh","metal3"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		providerItems0TypeNamePropEnum = append(providerItems0TypeNamePropEnum, v)
	}
}

const (

	// ProviderItems0NameClusterDashAPI captures enum value "cluster-api"
	ProviderItems0NameClusterDashAPI string = "cluster-api"

	// ProviderItems0NameKubeadm captures enum value "kubeadm"
	ProviderItems0NameKubeadm string = "kubeadm"

	// ProviderItems0NameByoh captures enum value "byoh"
	ProviderItems0NameByoh string = "byoh"

	// ProviderItems0NameMetal3 captures enum value "metal3"
	ProviderItems0NameMetal3 string = "metal3"
)

// prop value enum
func (m *ProviderItems0) validateNameEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, providerItems0TypeNamePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ProviderItems0) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	// value enum
	if err := m.validateNameEnum("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *ProviderItems0) validateParameters(formats strfmt.Registry) error {
	if swag.IsZero(m.Parameters) { // not required
		return nil
	}

	if m.Parameters != nil {
		if err := m.Parameters.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("parameters")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("parameters")
			}
			return err
		}
	}

	return nil
}

var providerItems0TypeProviderTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CoreProvider","BootstrapProvider","ControlPlaneProvider","InfrastructureProvider"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		providerItems0TypeProviderTypePropEnum = append(providerItems0TypeProviderTypePropEnum, v)
	}
}

const (

	// ProviderItems0ProviderTypeCoreProvider captures enum value "CoreProvider"
	ProviderItems0ProviderTypeCoreProvider string = "CoreProvider"

	// ProviderItems0ProviderTypeBootstrapProvider captures enum value "BootstrapProvider"
	ProviderItems0ProviderTypeBootstrapProvider string = "BootstrapProvider"

	// ProviderItems0ProviderTypeControlPlaneProvider captures enum value "ControlPlaneProvider"
	ProviderItems0ProviderTypeControlPlaneProvider string = "ControlPlaneProvider"

	// ProviderItems0ProviderTypeInfrastructureProvider captures enum value "InfrastructureProvider"
	ProviderItems0ProviderTypeInfrastructureProvider string = "InfrastructureProvider"
)

// prop value enum
func (m *ProviderItems0) validateProviderTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, providerItems0TypeProviderTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ProviderItems0) validateProviderType(formats strfmt.Registry) error {
	if swag.IsZero(m.ProviderType) { // not required
		return nil
	}

	// value enum
	if err := m.validateProviderTypeEnum("provider_type", "body", m.ProviderType); err != nil {
		return err
	}

	return nil
}

func (m *ProviderItems0) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.Pattern("url", "body", m.URL, `(?:(?:https?|http|ftp|file|oci)://|www.|ftp.)(?:([-A-Z0-9+&@#/%=~_|$?!:,.]*)|[-A-Z0-9+&@#/%=~_|$?!:,.])*(?:([-A-Z0-9+&@#/%=~_|$?!:,.]*)|[A-Z0-9+&@#/%=~_|$])`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this provider items0 based on the context it is used
func (m *ProviderItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateParameters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProviderItems0) contextValidateParameters(ctx context.Context, formats strfmt.Registry) error {

	if m.Parameters != nil {
		if err := m.Parameters.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("parameters")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("parameters")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProviderItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProviderItems0) UnmarshalBinary(b []byte) error {
	var res ProviderItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ProviderItems0Parameters provider items0 parameters
//
// swagger:model ProviderItems0Parameters
type ProviderItems0Parameters struct {

	// metadata
	// Pattern: (?:(?:https?|http|ftp|file|oci)://|www.|ftp.)(?:([-A-Z0-9+&@#/%=~_|$?!:,.]*)|[-A-Z0-9+&@#/%=~_|$?!:,.])*(?:([-A-Z0-9+&@#/%=~_|$?!:,.]*)|[A-Z0-9+&@#/%=~_|$])
	Metadata string `json:"metadata,omitempty"`

	// provider label
	ProviderLabel string `json:"provider_label,omitempty"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this provider items0 parameters
func (m *ProviderItems0Parameters) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProviderItems0Parameters) validateMetadata(formats strfmt.Registry) error {
	if swag.IsZero(m.Metadata) { // not required
		return nil
	}

	if err := validate.Pattern("parameters"+"."+"metadata", "body", m.Metadata, `(?:(?:https?|http|ftp|file|oci)://|www.|ftp.)(?:([-A-Z0-9+&@#/%=~_|$?!:,.]*)|[-A-Z0-9+&@#/%=~_|$?!:,.])*(?:([-A-Z0-9+&@#/%=~_|$?!:,.]*)|[A-Z0-9+&@#/%=~_|$])`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this provider items0 parameters based on context it is used
func (m *ProviderItems0Parameters) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ProviderItems0Parameters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProviderItems0Parameters) UnmarshalBinary(b []byte) error {
	var res ProviderItems0Parameters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
