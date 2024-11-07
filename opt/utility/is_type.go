package utility

import "reflect"

func IsArrayOrSlice(input interface{}) bool {
	kind := reflect.TypeOf(input).Kind()
	return kind == reflect.Array || kind == reflect.Slice
}

func IsMap(input interface{}) bool {
	return reflect.TypeOf(input).Kind() == reflect.Map
}

func IsStruct(input interface{}) bool {
	return reflect.TypeOf(input).Kind() == reflect.Struct
}

func IsPointer(input interface{}) bool {
	return reflect.TypeOf(input).Kind() == reflect.Ptr
}

func IsInterface(input interface{}) bool {
	return reflect.TypeOf(input).Kind() == reflect.Interface
}

func IsFunction(input interface{}) bool {
	return reflect.TypeOf(input).Kind() == reflect.Func
}

func IsChan(input interface{}) bool {
	return reflect.TypeOf(input).Kind() == reflect.Chan
}

func IsString(input interface{}) bool {
	return reflect.TypeOf(input).Kind() == reflect.String
}

func IsBool(input interface{}) bool {
	return reflect.TypeOf(input).Kind() == reflect.Bool
}

func IsInt(input interface{}) bool {
	return reflect.TypeOf(input).Kind() == reflect.Int || reflect.TypeOf(input).Kind() == reflect.Int8 || reflect.TypeOf(input).Kind() == reflect.Int16 || reflect.TypeOf(input).Kind() == reflect.Int32 || reflect.TypeOf(input).Kind() == reflect.Int64
}
