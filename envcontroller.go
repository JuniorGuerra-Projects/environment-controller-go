package envcontroller

import (
	"encoding/json"
	"os"
	"reflect"
)

func InterfaceToStruct(in interface{}, out interface{}) error {
	var err error
	var data []byte
	if data, err = json.Marshal(in); err != nil {
		return err
	}

	err = json.Unmarshal(data, &out)
	return err
}

func New(in interface{}) interface{} {
	return getData(in)
}

func getData(in interface{}) interface{} {
	refVal := reflect.ValueOf(in)
	refType := reflect.TypeOf(in)
	out := make(map[string]interface{})

	for x := 0; x < refVal.NumField(); x++ {
		// field := refVal.Field(x)
		// type1 := field.Type().String()
		tag := refType.Field(x).Tag.Get("env")
		out[refType.Field(x).Name] = os.Getenv(tag)

	}

	return out
}
