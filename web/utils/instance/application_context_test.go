package instance

import (
	"fmt"
	"reflect"
	"testing"
)

type TestService struct {
}

func TestApplicationContext_GetInstanceByName(t *testing.T) {
	testService := &TestService{}

	applicationContext := NewApplicationContext()
	applicationContext.RegisterInstance(testService)

	instance := applicationContext.MustGetInstance(&TestService{}).(*TestService)
	fmt.Println(reflect.TypeOf(instance))
}
