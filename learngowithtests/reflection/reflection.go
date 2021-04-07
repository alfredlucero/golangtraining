package reflection

import (
	"reflect"
)

// Reflection is the ability of a program to examine its own structure, particularly through types. It is a form of metaprogramming.
// For scenarios where you want to write a function where you don't know the type at compile time, interfaces come in handy and can think of
// as the any type. We lose type safety though.
// i.e. walk(x interface{}, fn func(string)) will accept any value for x.
func walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	walkValue := func(value reflect.Value) {
		walk(value.Interface(), fn)
	}

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walkValue(val.Field(i))
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			walkValue(val.Index(i))
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walkValue(val.MapIndex(key))
		}
	case reflect.Chan:
		for v, ok := val.Recv(); ok; v, ok = val.Recv() {
			walk(v.Interface(), fn)
		}
	case reflect.Func:
		valFnResult := val.Call(nil)
		// This allows for arbitrary number of return values.
		for _, res := range valFnResult {
			walk(res.Interface(), fn)
		}
	}
}

func getValue(x interface{}) reflect.Value {
	// reflect package has ValueOf which returns us a Value of a given variable.
	// We can inspect a value, including its fields.
	val := reflect.ValueOf(x)

	// We cannot use NumField on a pointer Value, so we extract underlying value with Elem().
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	return val
}
