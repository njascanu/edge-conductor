// Code generated by go-swagger; DO NOT EDIT.

//
//   Copyright (c) 2022 Intel Corporation.
//
//   SPDX-License-Identifier: Apache-2.0
//
//
//

package plugins

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

// Node node
//
// swagger:model node
type Node struct {

	// bmc endpoint
	// Pattern: (?:(?:https?|http|ftp|file|oci)://|www.|ftp.)(?:([-A-Z0-9+&@#/%=~_|$?!:,.]*)|[-A-Z0-9+&@#/%=~_|$?!:,.])*(?:([-A-Z0-9+&@#/%=~_|$?!:,.]*)|[A-Z0-9+&@#/%=~_|$])
	BmcEndpoint string `json:"bmc_endpoint,omitempty"`

	// bmc password
	BmcPassword string `json:"bmc_password,omitempty"`

	// bmc protocol
	// Enum: [redfish ipmi]
	BmcProtocol string `json:"bmc_protocol,omitempty"`

	// bmc user
	BmcUser string `json:"bmc_user,omitempty"`

	// critype
	// Enum: [containerd dockerd]
	Critype string `json:"critype,omitempty"`

	// ip
	// Pattern: (\b25[0-5]|\b2[0-4][0-9]|\b[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}\b
	IP string `json:"ip,omitempty"`

	// labels
	Labels []*NodeLabelsItems0 `json:"labels"`

	// mac
	// Pattern: ^[a-fA-F0-9]{2}(:[a-fA-F0-9]{2}){5}$
	Mac string `json:"mac,omitempty"`

	// name
	// Pattern: ^[a-zA-Z_$][a-zA-Z_.\-$0-9]*$
	Name string `json:"name,omitempty"`

	// role
	Role []string `json:"role"`

	// ssh key
	SSHKey string `json:"ssh_key,omitempty"`

	// ssh key path
	SSHKeyPath string `json:"ssh_key_path,omitempty"`

	// ssh passwd
	SSHPasswd string `json:"ssh_passwd,omitempty"`

	// ssh port
	SSHPort int64 `json:"ssh_port,omitempty"`

	// user
	User string `json:"user,omitempty"`
}

// Validate validates this node
func (m *Node) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBmcEndpoint(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBmcProtocol(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCritype(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIP(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabels(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMac(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Node) validateBmcEndpoint(formats strfmt.Registry) error {
	if swag.IsZero(m.BmcEndpoint) { // not required
		return nil
	}

	if err := validate.Pattern("bmc_endpoint", "body", m.BmcEndpoint, `(?:(?:https?|http|ftp|file|oci)://|www.|ftp.)(?:([-A-Z0-9+&@#/%=~_|$?!:,.]*)|[-A-Z0-9+&@#/%=~_|$?!:,.])*(?:([-A-Z0-9+&@#/%=~_|$?!:,.]*)|[A-Z0-9+&@#/%=~_|$])`); err != nil {
		return err
	}

	return nil
}

var nodeTypeBmcProtocolPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["redfish","ipmi"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		nodeTypeBmcProtocolPropEnum = append(nodeTypeBmcProtocolPropEnum, v)
	}
}

const (

	// NodeBmcProtocolRedfish captures enum value "redfish"
	NodeBmcProtocolRedfish string = "redfish"

	// NodeBmcProtocolIpmi captures enum value "ipmi"
	NodeBmcProtocolIpmi string = "ipmi"
)

// prop value enum
func (m *Node) validateBmcProtocolEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, nodeTypeBmcProtocolPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Node) validateBmcProtocol(formats strfmt.Registry) error {
	if swag.IsZero(m.BmcProtocol) { // not required
		return nil
	}

	// value enum
	if err := m.validateBmcProtocolEnum("bmc_protocol", "body", m.BmcProtocol); err != nil {
		return err
	}

	return nil
}

var nodeTypeCritypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["containerd","dockerd"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		nodeTypeCritypePropEnum = append(nodeTypeCritypePropEnum, v)
	}
}

const (

	// NodeCritypeContainerd captures enum value "containerd"
	NodeCritypeContainerd string = "containerd"

	// NodeCritypeDockerd captures enum value "dockerd"
	NodeCritypeDockerd string = "dockerd"
)

// prop value enum
func (m *Node) validateCritypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, nodeTypeCritypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Node) validateCritype(formats strfmt.Registry) error {
	if swag.IsZero(m.Critype) { // not required
		return nil
	}

	// value enum
	if err := m.validateCritypeEnum("critype", "body", m.Critype); err != nil {
		return err
	}

	return nil
}

func (m *Node) validateIP(formats strfmt.Registry) error {
	if swag.IsZero(m.IP) { // not required
		return nil
	}

	if err := validate.Pattern("ip", "body", m.IP, `(\b25[0-5]|\b2[0-4][0-9]|\b[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}\b`); err != nil {
		return err
	}

	return nil
}

func (m *Node) validateLabels(formats strfmt.Registry) error {
	if swag.IsZero(m.Labels) { // not required
		return nil
	}

	for i := 0; i < len(m.Labels); i++ {
		if swag.IsZero(m.Labels[i]) { // not required
			continue
		}

		if m.Labels[i] != nil {
			if err := m.Labels[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("labels" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("labels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Node) validateMac(formats strfmt.Registry) error {
	if swag.IsZero(m.Mac) { // not required
		return nil
	}

	if err := validate.Pattern("mac", "body", m.Mac, `^[a-fA-F0-9]{2}(:[a-fA-F0-9]{2}){5}$`); err != nil {
		return err
	}

	return nil
}

func (m *Node) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.Pattern("name", "body", m.Name, `^[a-zA-Z_$][a-zA-Z_.\-$0-9]*$`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this node based on the context it is used
func (m *Node) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLabels(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Node) contextValidateLabels(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Labels); i++ {

		if m.Labels[i] != nil {
			if err := m.Labels[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("labels" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("labels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Node) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Node) UnmarshalBinary(b []byte) error {
	var res Node
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// NodeLabelsItems0 node labels items0
//
// swagger:model NodeLabelsItems0
type NodeLabelsItems0 struct {

	// name
	Name string `json:"name,omitempty"`

	// value
	Value string `json:"value,omitempty"`
}

// Validate validates this node labels items0
func (m *NodeLabelsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this node labels items0 based on context it is used
func (m *NodeLabelsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NodeLabelsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NodeLabelsItems0) UnmarshalBinary(b []byte) error {
	var res NodeLabelsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
