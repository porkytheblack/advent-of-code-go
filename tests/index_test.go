package tests

import (
	"fmt"
	"strings"
	"testing"
)


func Test(t *testing.T){

	split := strings.Split("s: dd", "s:")

	fmt.Printf("%v",split)

	fmt.Println("Length", len(split))

}