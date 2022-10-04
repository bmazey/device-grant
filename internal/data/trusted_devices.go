package data

import (
	"sync"
)

type TrustedDevice struct {
	ID    string
	Owner string
}

type TrustedDeviceStore struct {
	Devices []TrustedDevice
	mu      sync.RWMutex
}

func NewDeviceStore() TrustedDeviceStore {
	return TrustedDeviceStore{
		Devices: make([]TrustedDevice, 0),
	}
}

func (s *TrustedDeviceStore) AddDevice(d TrustedDevice) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.Devices = append(s.Devices, d)
}

// Contains is a little different - we need to know if there's a device with matching id / client_id
func (s *TrustedDeviceStore) Contains(deviceID string, clientID string) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, device := range s.Devices {
		if device.ID == deviceID && device.Owner == clientID {
			return true
		}
	}

	return false
}
