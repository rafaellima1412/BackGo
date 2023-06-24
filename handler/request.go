package handler

import "fmt"

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

type CreateOpeningRequest struct {
	Role     string `json:"role,omitempty"`
	Company  string `json:"company,omitempty"`
	Location string `json:"location,omitempty"`
	Remote   *bool  `json:"remote,omitempty"`
	Link     string `json:"link,omitempty"`
	Salary   int64  `json:"salary,omitempty"`
}

func (r *CreateOpeningRequest) Validade() error {

	if r.Role == "" && r.Company == "" && r.Location == "" && r.Remote == nil && r.Salary <= 0 {
		return fmt.Errorf("request body is empty")
	}
	if r.Role == "" {
		return errParamIsRequired("role", "string")
	}
	if r.Company == "" {
		return errParamIsRequired("company", "string")
	}
	if r.Location == "" {
		return errParamIsRequired("location", "string")
	}
	if r.Link == "" {
		return errParamIsRequired("link", "string")
	}
	if r.Remote == nil {
		return errParamIsRequired("remote", "bool")
	}
	if r.Salary <= 0 {
		return errParamIsRequired("salary", "int64")
	}
	return nil
}

type UpdateOpeningRequest struct {
	Role     string `json:"role,omitempty"`
	Company  string `json:"company,omitempty"`
	Location string `json:"location,omitempty"`
	Remote   *bool  `json:"remote,omitempty"`
	Link     string `json:"link,omitempty"`
	Salary   int64  `json:"salary,omitempty"`
}

func (r *UpdateOpeningRequest) Validade() error {

	if r.Role != "" || r.Company != "" || r.Location != "" || r.Remote != nil || r.Salary < 0 {
		return nil
	}
	return fmt.Errorf("at least one valid field must be provided")

}
