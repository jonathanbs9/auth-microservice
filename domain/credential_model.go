package domain

import "fmt"

type Credential struct {
	ID           CredentialID `json:"id"`
	Name         string       `json:"name"`
	PasswordHash string       `json:"-"`
	AuthToken    AuthToken    `json:"-"`
}

func NewCredential(params SaveParams) (Credential, error) {
	if err := validate(params); err != nil {
		return Credential{}, err
	}

	return Credential{
		ID:   params.ID,
		Name: params.Name,
	}, nil
}

func validate(params SaveParams) error {
	if params.Name == "" {
		return fmt.Errorf("name is required")
	}

	if params.ID == 0 {
		if params.Password == "" {
			return fmt.Errorf("password is required")
		}
	}

	if params.Password != "" && len(params.Password) < 8 {
		return fmt.Errorf("password must be at least 8 characters")
	}
	return nil
}

//* func (a *Credential) HashPassword(hasher Hasher, password string) error
