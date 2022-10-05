package data

import (
	"github.com/goombaio/namegenerator"
	"sync"
	"time"
)

// ClientStore acts as a simple in-memory client_id datastore
type ClientStore struct {
	IDs       []string
	Generator namegenerator.Generator
	mu        sync.RWMutex
}

func NewClientStore() ClientStore {
	seed := time.Now().UTC().UnixNano()
	return ClientStore{
		IDs:       make([]string, 0),
		Generator: namegenerator.NewNameGenerator(seed),
	}
}

func (c *ClientStore) Create() string {
	name := c.Generator.Generate()

	c.mu.Lock()
	defer c.mu.Unlock()

	c.IDs = append(c.IDs, name)

	return name
}

func (c *ClientStore) Delete(id string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	for i, v := range c.IDs {
		if v == id {
			c.IDs = append(c.IDs[:i], c.IDs[i+1:]...)
			break
		}
	}
}

func (c *ClientStore) Contains(id string) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()

	for _, v := range c.IDs {
		if v == id {
			return true
		}
	}

	return false
}
