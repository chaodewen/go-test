package reflect

import (
	"reflect"
	"fmt"
	"testing"
)

func TestValueOf(t *testing.T) {
	testMap := map[string]int {
		"a": 1,
		"b": 2,
	}

	testMapValue := reflect.ValueOf(testMap)

	fmt.Printf("testMapValue: %v, %v, %v\n", testMapValue, testMapValue.Kind(), testMapValue.Type())

	for _, key := range testMapValue.MapKeys() {
		value := testMapValue.MapIndex(key)
		fmt.Println(value)
	}

	testMapValueValue := reflect.ValueOf(testMapValue)

	fmt.Printf("testMapValueValue: %v, %v, %v\n", testMapValueValue, testMapValueValue.Kind(),
		testMapValueValue.Kind())
}
