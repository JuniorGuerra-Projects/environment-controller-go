package envcontroller

import (
	"fmt"
	"os"
	"testing"
)

type ControllerTest struct {
	AccNo      string `env:"ACC_NO"`
	CustomerId string `env:"CUSTOMER_ID"`
	ClientId   string `env:"CLIENT_ID"`
}

func initEnv() {
	os.Setenv("ACC_NO", "1234567890")
	os.Setenv("CUSTOMER_ID", "10")
	os.Setenv("CLIENT_ID", "5")
}

func TestGetData(t *testing.T) {
	initEnv()
	var ctlTest ControllerTest
	result := New(ctlTest)

	if result == nil {
		t.Errorf("Do error data")
	}

	InterfaceToStruct(result, &ctlTest)

	fmt.Printf("%+v\n", ctlTest)
}
