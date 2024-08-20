package panicIf

import (
	"fmt"
	"github.com/VegieDoggie/go-sdkx/check"
	"reflect"
)

// Err panic if the last argument is a non-nil error
func Err(arguments ...any) {
	if a := arguments[len(arguments)-1]; a != nil {
		switch a.(type) {
		case error:
			panic(fmt.Errorf("panic Err: %v", a))
		}
	}
}

// True panic if the argument is true
func True(argument bool, info ...string) {
	if argument {
		panic(fmt.Errorf("panic True: %+v", info))
	}
}

// False panic if the argument is false
func False(argument bool, info ...string) {
	if !argument {
		panic(fmt.Errorf("panic False: %+v", info))
	}
}

// Nil panic if the argument is nil
func Nil(argument any, info ...string) {
	if argument == nil || reflect.ValueOf(argument).IsNil() {
		panic(fmt.Errorf("panic Nil: %+v", info))
	}
}

// Empty panic if the argument is nil or empty
func Empty(argument any, info ...string) {
	if check.IsEmpty(argument) {
		panic(fmt.Errorf("panic Empty: %+v", info))
	}
}
