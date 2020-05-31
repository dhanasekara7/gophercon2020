package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	fmt.Println(reflect.ValueOf(math.Sqrt(64)).Interface())
}
