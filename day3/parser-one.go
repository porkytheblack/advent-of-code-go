package day3

import "strconv"

type Num struct {
	value int
	row int 
	column int
}

type Symbol struct {
	value string 
	row int 
	column int
}

var symbols = []string{
	"!",
	"@",
	"#",
	"$",
	"%",
	"^",
	"&",
	"*",
	"(",
	")",
}

func isNumber(value byte) bool {
	_, err := strconv.ParseInt(string(value), 10, 64)
	return err == nil
}

func isDot(value byte) bool {
	return string(value) == "."
}

func isSymbol(value byte) bool {

}

func Runner(maze []string,row int, column int, nums []Num, symbols []Symbol){

	value := maze[row][column]

	
	if(isNumber(value)) {
		let 
		nums = append(nums, Num{
			value: ,
		})
	}

}