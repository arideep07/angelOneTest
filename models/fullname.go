package models

import "errors"

// FullNameRequest is the request body for the full name api
type FullNameRequest struct {
	FirstName string `json:"firstName" example:"Nitesh"`
	LastName  string `json:"lastName" example:"Jain"`
}

// FullNameResponse is the response body for the full name api
type FullNameResponse struct {
	FullName string `json:"name" example:"Nitesh Jain"`
}

// Validate is used to validate the request body
func (r FullNameRequest) Validate() error {
	if r.FirstName == "" {
		return errors.New("first name cannot be empty")
	}
	return nil
}
