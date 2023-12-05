package day4

import (
	"fmt"
	"testing"
)

func TestParese2(t *testing.T){
	compiled := Parse2("../input/day4/test.txt")
	fmt.Println(compiled)
}