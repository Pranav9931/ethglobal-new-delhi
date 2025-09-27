package httpdtos

type UserDTO struct {
	ID               string          `json:"id"`
	CreatedAt        int64           `json:"created_at"`
	LinkedAccounts   []LinkedAccount `json:"linked_accounts"`
	MfaMethods       []interface{}   `json:"mfa_methods"`
	HasAcceptedTerms bool            `json:"has_accepted_terms"`
	IsGuest          bool            `json:"is_guest"`
}

type LinkedAccount struct {
	// some fields are optional, so pointers or omitempty can be used if needed
	ID               string `json:"id,omitempty"`
	Type             string `json:"type"`
	Address          string `json:"address"`
	VerifiedAt       int64  `json:"verified_at"`
	FirstVerifiedAt  int64  `json:"first_verified_at"`
	LatestVerifiedAt int64  `json:"latest_verified_at"`

	WalletIndex      *int   `json:"wallet_index,omitempty"`
	ChainID          string `json:"chain_id,omitempty"`
	ChainType        string `json:"chain_type,omitempty"`
	Delegated        *bool  `json:"delegated,omitempty"`
	WalletClient     string `json:"wallet_client,omitempty"`
	WalletClientType string `json:"wallet_client_type,omitempty"`
	ConnectorType    string `json:"connector_type,omitempty"`
	Imported         *bool  `json:"imported,omitempty"`
	RecoveryMethod   string `json:"recovery_method,omitempty"`
}
