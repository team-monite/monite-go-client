// This file was auto-generated by Fern from our API Definition.

package entities

import (
	monitegoclient "github.com/team-monite/monite-go-client"
)

type PersonRequest struct {
	// The person's address
	Address *monitegoclient.PersonAddressRequest `json:"address,omitempty" url:"-"`
	// The person's date of birth
	DateOfBirth *string `json:"date_of_birth,omitempty" url:"-"`
	// The person's first name
	FirstName string `json:"first_name" url:"-"`
	// The person's last name
	LastName string `json:"last_name" url:"-"`
	// The person's email address
	Email string `json:"email" url:"-"`
	// The person's phone number
	Phone *string `json:"phone,omitempty" url:"-"`
	// Describes the person's relationship to the entity
	Relationship *monitegoclient.PersonRelationshipRequest `json:"relationship,omitempty" url:"-"`
	// The person's ID number, as appropriate for their country
	IdNumber *string `json:"id_number,omitempty" url:"-"`
	// The last four digits of the person's Social Security number
	SsnLast4 *string `json:"ssn_last_4,omitempty" url:"-"`
	// Required for persons of US entities. The country of the person's citizenship, as a two-letter country code (ISO 3166-1 alpha-2). In case of dual or multiple citizenship, specify any.
	Citizenship *monitegoclient.AllowedCountries `json:"citizenship,omitempty" url:"-"`
}

type OptionalPersonRequest struct {
	// The person's address
	Address *monitegoclient.OptionalPersonAddressRequest `json:"address,omitempty" url:"-"`
	// The person's date of birth
	DateOfBirth *string `json:"date_of_birth,omitempty" url:"-"`
	// The person's first name
	FirstName *string `json:"first_name,omitempty" url:"-"`
	// The person's last name
	LastName *string `json:"last_name,omitempty" url:"-"`
	// The person's email address
	Email *string `json:"email,omitempty" url:"-"`
	// The person's phone number
	Phone *string `json:"phone,omitempty" url:"-"`
	// Describes the person's relationship to the entity
	Relationship *monitegoclient.OptionalPersonRelationship `json:"relationship,omitempty" url:"-"`
	// The person's ID number, as appropriate for their country
	IdNumber *string `json:"id_number,omitempty" url:"-"`
	// The last four digits of the person's Social Security number
	SsnLast4 *string `json:"ssn_last_4,omitempty" url:"-"`
	// Required for persons of US entities. The country of the person's citizenship, as a two-letter country code (ISO 3166-1 alpha-2). In case of dual or multiple citizenship, specify any.
	Citizenship *monitegoclient.AllowedCountries `json:"citizenship,omitempty" url:"-"`
}

type PersonOnboardingDocumentsPayload struct {
	AdditionalVerificationDocumentBack  *string `json:"additional_verification_document_back,omitempty" url:"-"`
	AdditionalVerificationDocumentFront *string `json:"additional_verification_document_front,omitempty" url:"-"`
	VerificationDocumentBack            *string `json:"verification_document_back,omitempty" url:"-"`
	VerificationDocumentFront           *string `json:"verification_document_front,omitempty" url:"-"`
}
