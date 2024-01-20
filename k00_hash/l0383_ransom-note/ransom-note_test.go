package l0383_ransom_note

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGoChar(t *testing.T) {
	s := "abcdefg"
	fmt.Println(reflect.TypeOf('a'))
	for i := range s {
		fmt.Println(s[i] - 'a')
		fmt.Println(reflect.TypeOf(s[i] - 'a'))

	}
}
