package data

import (
	"sync"
	"time"
)

type Interaction struct {
	ClientID string
	DeviceID string
	UserCode string
	Expires  time.Time
}

func (i *Interaction) Create(clientID string, deviceID string, userCode string, ttl time.Duration) Interaction {
	expires := time.Unix(time.Now().Unix()+int64(ttl.Seconds()), 0)

	return Interaction{
		ClientID: clientID,
		DeviceID: deviceID,
		UserCode: userCode,
		Expires:  expires,
	}
}

func (i *Interaction) IsExpired() bool {
	if time.Now().After(i.Expires) {
		return true
	}
	return false
}

type InteractionStore struct {
	Interactions []Interaction
	mu           sync.RWMutex
}

func (s *InteractionStore) Add(i Interaction) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.Interactions = append(s.Interactions, i)
}

func (s *InteractionStore) Delete(i Interaction) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for x, v := range s.Interactions {
		if v == i {
			s.Interactions = append(s.Interactions[:x], s.Interactions[x+1:]...)
			break
		}
	}
}

// Retrieve attempts to return an unexpired interaction given a user_code
func (s *InteractionStore) Retrieve(userCode string) *Interaction {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, v := range s.Interactions {
		if v.UserCode == userCode && !v.IsExpired() {
			return &v
		}
	}

	return nil
}
