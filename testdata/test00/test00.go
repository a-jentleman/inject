package test00

import (
	i2 "github.com/a-jentleman/inject"
)

func ProvideInt() int {
	return i2.Inject[int](ReturnInt)
}

func ReturnInt() int {
	return 1
}