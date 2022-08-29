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

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Certificate certificate
//
// swagger:model certificate
type Certificate struct {

	// ca
	Ca *CertificateCa `json:"ca,omitempty"`

	// client
	Client *CertificateClient `json:"client,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// server
	Server *CertificateServer `json:"server,omitempty"`
}

// Validate validates this certificate
func (m *Certificate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCa(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClient(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServer(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Certificate) validateCa(formats strfmt.Registry) error {
	if swag.IsZero(m.Ca) { // not required
		return nil
	}

	if m.Ca != nil {
		if err := m.Ca.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ca")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ca")
			}
			return err
		}
	}

	return nil
}

func (m *Certificate) validateClient(formats strfmt.Registry) error {
	if swag.IsZero(m.Client) { // not required
		return nil
	}

	if m.Client != nil {
		if err := m.Client.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("client")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("client")
			}
			return err
		}
	}

	return nil
}

func (m *Certificate) validateServer(formats strfmt.Registry) error {
	if swag.IsZero(m.Server) { // not required
		return nil
	}

	if m.Server != nil {
		if err := m.Server.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("server")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("server")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this certificate based on the context it is used
func (m *Certificate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCa(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateClient(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateServer(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Certificate) contextValidateCa(ctx context.Context, formats strfmt.Registry) error {

	if m.Ca != nil {
		if err := m.Ca.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ca")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ca")
			}
			return err
		}
	}

	return nil
}

func (m *Certificate) contextValidateClient(ctx context.Context, formats strfmt.Registry) error {

	if m.Client != nil {
		if err := m.Client.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("client")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("client")
			}
			return err
		}
	}

	return nil
}

func (m *Certificate) contextValidateServer(ctx context.Context, formats strfmt.Registry) error {

	if m.Server != nil {
		if err := m.Server.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("server")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("server")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Certificate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Certificate) UnmarshalBinary(b []byte) error {
	var res Certificate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CertificateCa certificate ca
//
// swagger:model CertificateCa
type CertificateCa struct {

	// cert
	// Pattern: ^[a-zA-Z.\/][a-zA-Z0-9-_.\/]*$
	Cert string `json:"cert,omitempty"`

	// csr
	// Pattern: ^[a-zA-Z.\/][a-zA-Z0-9-_.\/]*$
	Csr string `json:"csr,omitempty"`

	// key
	// Pattern: ^[a-zA-Z.\/][a-zA-Z0-9-_.\/]*$
	Key string `json:"key,omitempty"`
}

// Validate validates this certificate ca
func (m *CertificateCa) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCert(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCsr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CertificateCa) validateCert(formats strfmt.Registry) error {
	if swag.IsZero(m.Cert) { // not required
		return nil
	}

	if err := validate.Pattern("ca"+"."+"cert", "body", m.Cert, `^[a-zA-Z.\/][a-zA-Z0-9-_.\/]*$`); err != nil {
		return err
	}

	return nil
}

func (m *CertificateCa) validateCsr(formats strfmt.Registry) error {
	if swag.IsZero(m.Csr) { // not required
		return nil
	}

	if err := validate.Pattern("ca"+"."+"csr", "body", m.Csr, `^[a-zA-Z.\/][a-zA-Z0-9-_.\/]*$`); err != nil {
		return err
	}

	return nil
}

func (m *CertificateCa) validateKey(formats strfmt.Registry) error {
	if swag.IsZero(m.Key) { // not required
		return nil
	}

	if err := validate.Pattern("ca"+"."+"key", "body", m.Key, `^[a-zA-Z.\/][a-zA-Z0-9-_.\/]*$`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this certificate ca based on context it is used
func (m *CertificateCa) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CertificateCa) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CertificateCa) UnmarshalBinary(b []byte) error {
	var res CertificateCa
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CertificateClient certificate client
//
// swagger:model CertificateClient
type CertificateClient struct {

	// cert
	// Pattern: ^[a-zA-Z.\/][a-zA-Z0-9-_.\/]*$
	Cert string `json:"cert,omitempty"`

	// csr
	// Pattern: ^[a-zA-Z.\/][a-zA-Z0-9-_.\/]*$
	Csr string `json:"csr,omitempty"`

	// key
	// Pattern: ^[a-zA-Z.\/][a-zA-Z0-9-_.\/]*$
	Key string `json:"key,omitempty"`
}

// Validate validates this certificate client
func (m *CertificateClient) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCert(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCsr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CertificateClient) validateCert(formats strfmt.Registry) error {
	if swag.IsZero(m.Cert) { // not required
		return nil
	}

	if err := validate.Pattern("client"+"."+"cert", "body", m.Cert, `^[a-zA-Z.\/][a-zA-Z0-9-_.\/]*$`); err != nil {
		return err
	}

	return nil
}

func (m *CertificateClient) validateCsr(formats strfmt.Registry) error {
	if swag.IsZero(m.Csr) { // not required
		return nil
	}

	if err := validate.Pattern("client"+"."+"csr", "body", m.Csr, `^[a-zA-Z.\/][a-zA-Z0-9-_.\/]*$`); err != nil {
		return err
	}

	return nil
}

func (m *CertificateClient) validateKey(formats strfmt.Registry) error {
	if swag.IsZero(m.Key) { // not required
		return nil
	}

	if err := validate.Pattern("client"+"."+"key", "body", m.Key, `^[a-zA-Z.\/][a-zA-Z0-9-_.\/]*$`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this certificate client based on context it is used
func (m *CertificateClient) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CertificateClient) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CertificateClient) UnmarshalBinary(b []byte) error {
	var res CertificateClient
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CertificateServer certificate server
//
// swagger:model CertificateServer
type CertificateServer struct {

	// cert
	// Pattern: ^[a-zA-Z.\/][a-zA-Z0-9-_.\/]*$
	Cert string `json:"cert,omitempty"`

	// csr
	// Pattern: ^[a-zA-Z.\/][a-zA-Z0-9-_.\/]*$
	Csr string `json:"csr,omitempty"`

	// key
	// Pattern: ^[a-zA-Z.\/][a-zA-Z0-9-_.\/]*$
	Key string `json:"key,omitempty"`
}

// Validate validates this certificate server
func (m *CertificateServer) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCert(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCsr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CertificateServer) validateCert(formats strfmt.Registry) error {
	if swag.IsZero(m.Cert) { // not required
		return nil
	}

	if err := validate.Pattern("server"+"."+"cert", "body", m.Cert, `^[a-zA-Z.\/][a-zA-Z0-9-_.\/]*$`); err != nil {
		return err
	}

	return nil
}

func (m *CertificateServer) validateCsr(formats strfmt.Registry) error {
	if swag.IsZero(m.Csr) { // not required
		return nil
	}

	if err := validate.Pattern("server"+"."+"csr", "body", m.Csr, `^[a-zA-Z.\/][a-zA-Z0-9-_.\/]*$`); err != nil {
		return err
	}

	return nil
}

func (m *CertificateServer) validateKey(formats strfmt.Registry) error {
	if swag.IsZero(m.Key) { // not required
		return nil
	}

	if err := validate.Pattern("server"+"."+"key", "body", m.Key, `^[a-zA-Z.\/][a-zA-Z0-9-_.\/]*$`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this certificate server based on context it is used
func (m *CertificateServer) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CertificateServer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CertificateServer) UnmarshalBinary(b []byte) error {
	var res CertificateServer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
