package store

import (
	"errors"
	"fmt" // Add this import
	"sync"
	"time"

	"ethglobal-nd-backend/internal/httpdtos"

	"github.com/google/uuid"
)

var (
	activeSession *httpdtos.SessionDTO
	mu            sync.Mutex
)

// CreateOrJoinSession creates new session if none exists, or joins existing one
func CreateOrJoinSession(address string) (*httpdtos.SessionDTO, string, bool) {
	mu.Lock()
	defer mu.Unlock()

	// Case 1: No active session → create new
	if activeSession == nil {
		activeSession = &httpdtos.SessionDTO{
			SessionId:     uuid.New(),
			SenderAddress: address,
			Currency:      "PYUSD",
			Status:        "waiting",
			ExpiresAt:     time.Now().Add(5 * time.Minute),
		}
		fmt.Printf("Session created: %+v\n", activeSession)
		return activeSession, "created", true
	}

	// Case 2: Session exists but waiting for receiver
	if activeSession.ReceiverAddress == "" {
		// Prevent sender re-joining as receiver
		if activeSession.SenderAddress == address {
			fmt.Printf("Join attempt with same address: %+v\n", activeSession)
			return nil, "cannot join with same address", false
		}
		activeSession.ReceiverAddress = address
		activeSession.Status = "active"
		fmt.Printf("Session joined: %+v\n", activeSession)
		return activeSession, "joined", true
	}

	// Case 3: Already full
	fmt.Printf("Join attempt, but session already full: %+v\n", activeSession)
	return nil, "session already full", false
}

func GetActiveSession() (*httpdtos.SessionDTO, bool) {
	mu.Lock()
	defer mu.Unlock()

	if activeSession == nil {
		return nil, false
	}
	return activeSession, true
}

func ApproveTransaction(address string, jwt string) (*httpdtos.SessionDTO, error) {
	mu.Lock()
	defer mu.Unlock()

	if activeSession == nil {
		return nil, errors.New("Session is not active")
	}

	if activeSession.SenderAddress == "" || activeSession.ReceiverAddress == "" {
		return nil, errors.New("Session is not fully joined")
	}
	if activeSession.SenderAddress == address && !activeSession.SenderApproval {
		activeSession.SenderApproval = true
		activeSession.SenderJwt = jwt
	}

	if activeSession.ReceiverAddress == address && !activeSession.ReceiverApproval {
		activeSession.ReceiverApproval = true
		activeSession.ReceiverJwt = jwt
	}

	return activeSession, nil
}
