package collection

import (
	"fmt"
	"testing"
)

func TestNil(t *testing.T) {
	var map1 map[string]string

	map2 := map[string]string{}

	fmt.Printf("map1[\"test\"]: %v\n", map1["test"])

	fmt.Printf("map2[\"test\"]: %v\n", map2["test"])

	fmt.Println(map1["test"] == "")

	fmt.Println(map2["test"] == "")
}

func TestArray(t *testing.T) {
	str := []string{"alipay", "quickpay"}

	str2 := append(str, "test")

	fmt.Printf("str: %+v, str2: %+v\n", str, str2)

	change(str)

	fmt.Printf("str: %+v\n", str)
}

func change(list []string) {
	l := []string{"1"}
	list = l
}
