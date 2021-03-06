// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Participant participant
// swagger:model Participant
type Participant struct {

	// the api endpoint for compliance service
	ComplianceServiceURL string `json:"compliance_service_url,omitempty"`

	// the api endpoint for federation service
	FederationServiceURL string `json:"federation_service_url,omitempty"`

	// collection of currency pairs
	QuoteCurrencySet []*CurrencyPair `json:"quote_currency_set"`

	// the api endpoint for providing quotes
	QuoteServiceURL string `json:"quote_service_url,omitempty"`

	// the stellar address for the participant
	// Required: true
	StellarAddress *string `json:"stellar_address"`

	// the stellar domain for the participant
	// Required: true
	StellarDomain *string `json:"stellar_domain"`
}

// Validate validates this participant
func (m *Participant) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateQuoteCurrencySet(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStellarAddress(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStellarDomain(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Participant) validateQuoteCurrencySet(formats strfmt.Registry) error {

	if swag.IsZero(m.QuoteCurrencySet) { // not required
		return nil
	}

	for i := 0; i < len(m.QuoteCurrencySet); i++ {

		if swag.IsZero(m.QuoteCurrencySet[i]) { // not required
			continue
		}

		if m.QuoteCurrencySet[i] != nil {

			if err := m.QuoteCurrencySet[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("quote_currency_set" + "." + strconv.Itoa(i))
				}
				return err
			}

		}

	}

	return nil
}

func (m *Participant) validateStellarAddress(formats strfmt.Registry) error {

	if err := validate.Required("stellar_address", "body", m.StellarAddress); err != nil {
		return err
	}

	return nil
}

func (m *Participant) validateStellarDomain(formats strfmt.Registry) error {

	if err := validate.Required("stellar_domain", "body", m.StellarDomain); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Participant) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Participant) UnmarshalBinary(b []byte) error {
	var res Participant
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
