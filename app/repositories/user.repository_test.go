package repositories

import (
	"api/interfaces"
	"testing"
)

// Checks if repositories.UserRepository
// is actually implementing the
// interfaces.UserRepositoryI interface
func TestInterfaceImplem(t *testing.T) {
	implem := UserRepository{}

	var _ interfaces.UserRepositoryI = &implem

	if _, ok := interface{}(&implem).(interfaces.UserRepositoryI); !ok {
		t.Errorf("UserRepository does not implement interfaces.UserRepositoryI interface")
	}
}
