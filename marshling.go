package marshaling

import (
	"reflect"
)

type Marshaler interface {
	GoMarshal() ([]byte, error)
}

type Unmarshaler interface {
	GoUnmarshal([]byte) error
}

var (
	marshalerType   = reflect.TypeOf((*Marshaler)(nil)).Elem()
	unmarshalerType = reflect.TypeOf((*Unmarshaler)(nil)).Elem()
)

func MarshalerType() reflect.Type {
	return marshalerType
}
func UnmarshalerType() reflect.Type {
	return unmarshalerType
}
