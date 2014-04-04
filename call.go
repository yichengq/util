package util

import (
	"fmt"
	"reflect"
)

func Call(fptr interface{}, args ...interface{}) (rets []interface{}) {
	fn := reflect.ValueOf(fptr)
	fnType := fn.Type()
	if fnType.IsVariadic() {
		panic("unsupport for variadic parameter")
	}
	if len(args) != fnType.NumIn() {
		panic(fmt.Sprintf("unmatched parameter number %d and %d", len(args), fnType.NumIn()))
	}
	argValues := make([]reflect.Value, fnType.NumIn())
	for i, arg := range args {
		if reflect.TypeOf(arg) != fnType.In(i) {
			panic("unmatched argument "+reflect.TypeOf(arg).Name()+" and "+fnType.In(i).Name())
		}
		argValues[i] = reflect.ValueOf(arg)
	}
	retValues := fn.Call(argValues)
	rets = make([]interface{}, fnType.NumOut())
	for i, retValue := range retValues {
		rets[i] = retValue.Interface()
	}
	return
}
