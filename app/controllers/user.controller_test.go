package controllers

import (
	"api/interfaces"
	"testing"
)

// Checks if controllers.UserController
// is actually implementing the
// interfaces.UserControllerI interface
func TestInterfaceImplem(t *testing.T) {
	implem := UserController{}

	var _ interfaces.UserControllerI = &implem

	if _, ok := interface{}(&implem).(interfaces.UserControllerI); !ok {
		t.Errorf("UserController does not implement interfaces.UserControllerI interface")
	}
}
