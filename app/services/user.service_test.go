package services

import (
	"api/interfaces"
	"testing"
)

func TestInterfaceImpl(t *testing.T) {
	implem := UserService{}

	if _, ok := interface{}(&implem).(interfaces.UserServiceI); !ok {
		t.Errorf("UserService does not implement interfaces.UserServiceI interface")
	}
}
