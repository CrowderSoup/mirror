package services

import "errors"

// NameCheap the NameCheap service
type NameCheap struct{}

// NewNameCheap returns a new NameCheap
func NewNameCheap() Service {
	return &NameCheap{}
}

// Update tries to update the DNS in NameCheap
func (n *NameCheap) Update(params map[string]interface{}) error {
	return errors.New("Not Implemented")
}
