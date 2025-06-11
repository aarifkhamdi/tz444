package challenge

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"sync"
	"time"

	"golang.org/x/crypto/argon2"
)

type PoWService struct {
	challenges map[string]*Challenge
	mutex      sync.RWMutex
}

func NewPoWService() *PoWService {
	powService := &PoWService{
		challenges: make(map[string]*Challenge),
	}

	go powService.cleanupExpired()

	return powService
}

func (p *PoWService) GenerateChallenge() (*Challenge, error) {
	id, err := generateRandomID()
	if err != nil {
		return nil, fmt.Errorf("failed to generate challenge id: %w", err)
	}

	const difficulty = 6
	ch := &Challenge{
		ID:         id,
		Difficulty: difficulty,
		ExpiresAt:  time.Now().Add(30 * time.Second),
	}

	p.mutex.Lock()
	p.challenges[id] = ch
	p.mutex.Unlock()

	return ch, nil
}

func (p *PoWService) VerifySolution(challengeID string, nonce int) (bool, error) {
	ch, exists := p.getChallenge(challengeID)
	if !exists {
		return false, fmt.Errorf("challenge not found or expired")
	}

	input := ch.ID + fmt.Sprintf("%d", nonce)
	hash := argon2.IDKey([]byte(input), []byte(ch.ID), 1, 64*1024, 1, 32)

	if hasLeadingZeros(hash, ch.Difficulty) {
		p.mutex.Lock()
		delete(p.challenges, challengeID)
		p.mutex.Unlock()
		return true, nil
	}

	p.mutex.Lock()
	delete(p.challenges, challengeID)
	p.mutex.Unlock()

	return false, nil
}

func (p *PoWService) getChallenge(id string) (*Challenge, bool) {
	p.mutex.RLock()
	ch, exists := p.challenges[id]
	p.mutex.RUnlock()
	if !exists {
		return nil, false
	}

	if time.Now().After(ch.ExpiresAt) {
		return nil, false
	}

	return ch, true
}

func (p *PoWService) cleanupExpired() {
	ticker := time.NewTicker(30 * time.Second)
	for range ticker.C {
		p.mutex.Lock()
		for id, ch := range p.challenges {
			if time.Now().After(ch.ExpiresAt) {
				delete(p.challenges, id)
			}
		}
		p.mutex.Unlock()
	}
}

func generateRandomID() (string, error) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func hasLeadingZeros(hash []byte, difficulty int) bool {
	for i := 0; i < difficulty/8; i++ {
		if hash[i] != 0 {
			return false
		}
	}
	if difficulty%8 > 0 {
		mask := byte(0xFF >> (8 - difficulty%8))
		return (hash[difficulty/8] & mask) == 0
	}
	return true
}
