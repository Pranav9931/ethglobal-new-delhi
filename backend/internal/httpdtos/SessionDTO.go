package httpdtos

import (
	"time"

	"github.com/google/uuid"
)

type SessionDTO struct {
	SessionId        uuid.UUID `json:"sessionId"`
	SenderAddress    string    `json:"sendorAddress"`
	SenderApproval   bool      `json:"senderApproval"`
	SenderJwt        string    `json:"senderJwt"`
	ReceiverAddress  string    `json:"receiverAddress"`
	ReceiverApproval bool      `json:"receiverApproval"`
	ReceiverJwt      string    `json:"receiverJwt"`
	Currency         string    `json:"currency"`
	Status           string    `json:"status"`
	ExpiresAt        time.Time `json:"expiresAt"`
}
