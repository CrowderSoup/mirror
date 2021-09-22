package services

// Service a Dynamic DNS Service
type Service interface {
	Update(map[string]interface{}) error
}
