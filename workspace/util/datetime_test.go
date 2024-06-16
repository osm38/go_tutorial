package util

import (
	"fmt"
	"testing"
)

func TestRandomDateTime(t *testing.T) {
	actual := RandomDateTime()
	fmt.Println(actual)
}
