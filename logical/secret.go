package logical

import (
	"fmt"
	"time"
)

// Secret represents the secret part of a response.
type Secret struct {
	// InternalData is JSON-encodable data that is stored with the secret.
	// This will be sent back during a Renew/Revoke for storing internal data
	// used for those operations.
	InternalData map[string]interface{} `json:"internal_data"`

	// Lease is the duration that this secret is valid for. Vault
	// will automatically revoke it after the duration + grace period.
	Lease            time.Duration `json:"lease"`
	LeaseGracePeriod time.Duration `json:"lease_grace_period"`

	// Renewable, if true, means that this secret can be renewed.
	Renewable bool `json:"renewable"`

	// LeaseIncrement will be the lease increment that the user requested.
	// This is only available on a Renew operation and has no effect
	// when returning a response.
	LeaseIncrement time.Duration `json:"-"`

	// VaultID is the ID returned to the user to represent this secret.
	// This is generated by Vault core. Any set value will be ignored.
	// For requests, this will always be blank.
	VaultID string
}

func (s *Secret) Validate() error {
	if s.Lease <= 0 {
		return fmt.Errorf("lease duration must not be less than zero")
	}
	if s.LeaseGracePeriod < 0 {
		return fmt.Errorf("lease grace period must not be less than zero")
	}

	return nil
}

func (s *Secret) GoString() string {
	return fmt.Sprintf("*%#v", *s)
}
