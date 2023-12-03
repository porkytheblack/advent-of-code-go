package day3_2

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Num struct {
	value int
	row int 
	column int
	is_start bool 
	is_end bool
	is_middle bool
}

type ConstructedNumber struct {
	value int
	str string 
	nums []Num
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
	"~",
	"`",
	"-",
	"=",
	"+",
	"_",
	"\\",
	"|",
	"/",
	"?",
	">",
	"<",
	",",
	":",
	";",
}

func isNumber(value byte) bool {
	_, err := strconv.ParseInt(string(value), 10, 64)
	return err == nil
}

func isDot(value byte) bool {
	return string(value) == "."
}

func isSymbol(value byte) bool {
	exists := false;

	for _, sym := range symbols {
		if(sym == string(value)) {
			exists = true
			break
		}
	}

	return exists
}

func Runner(maze []string, row_size, col_size, row int, column int, nums []Num, symbols []Symbol)([]Num, []Symbol) {

	if(row >= row_size) {
		// fmt.Printf("done")
		return nums, symbols
	}

	value := maze[row][column]

	
	if(isNumber(value)) {
		_v, err := strconv.Atoi(string(value)); if err != nil {
			panic("SOmething is wrong")
		}

		left := func()bool{
			if(column == 0){
				return false
			}
			return isNumber(maze[row][column - 1])
		}()

		right := func()bool{
			if(column + 1 >= col_size) {
				return false
			}
			return isNumber(maze[row][column + 1])
		}()

		is_start := func()bool{
			if(column == 0){
				return true
			}else if(!left){
				return true
			}	
			return false
		}()

		is_middle := left && right

		is_end := func()bool{
			if(!left && !right){
				return true
			}
			if(left && !right){
				return true
			}
			return false
		}()

		nums = append(nums, Num{
			value: _v,
			row: row,
			column: column,
			is_start: is_start,
			is_middle: is_middle,
			is_end: is_end,
		})
	}

	if(isSymbol(value)){
		symbols = append(symbols, Symbol{
			value: string(value),
			row: row,
			column: column,
		})
	}

	new_row := func ()int{
		if(column + 1 >= col_size){
			return row + 1
		}
		return row
	}()

	new_column := func ()int{
		if(column + 1 >= col_size){
			return 0
		}
		return column + 1
	}()
	
	return Runner(maze, row_size, col_size, new_row, new_column, nums, symbols)
}

func lookup_home(constructed map[string]ConstructedNumber, num Num)string{
	for index, value := range constructed {

		found := false

		for _, _n := range value.nums {

			if(num.is_middle){ // if its a middle digit
				if(_n.row == num.row && _n.column + 1 == num.column) { // if _n comes before this digit 
					found = true
					break
				}
				if(_n.row == num.row && _n.column - 1 == num.column) { // if _n comes after this digit 
					found = true
					break
				}
			}
			if(num.is_end){// if its an end digit
				if(_n.row == num.row && _n.column + 1 == num.column){//if _n comes before this digit
					found = true 
					break
				}
			}


		}

		if(found){
			return index
		}

	}

	return "-1" // not in the current list
}

func NumberCollector(all_numbers []Num)map[string]ConstructedNumber{
	constructed_numbers := make(map[string]ConstructedNumber)


	for _, num := range all_numbers {

		if(num.is_start){
			constructed_numbers[fmt.Sprintf("%d-%d",num.row,num.column)] = ConstructedNumber{
				str: fmt.Sprint(num.value),
				nums: []Num{
					num,
				},
			}
		}

	}

	// fmt.Printf("\n\n\n Constructed Numbers %v \n\n\n", constructed_numbers)

	all_other_numbers := func()[]Num{
		arr := []Num{}

		for _, n := range all_numbers {
			if(!n.is_start){
				arr = append(arr, n)
			}
		}

		return arr
	}()

	for _, num := range all_other_numbers {

		home := lookup_home(constructed_numbers, num)
		
		parent := constructed_numbers[home]
		parent.nums = append(parent.nums, num)
		parent.str = parent.str + fmt.Sprint(num.value)
		constructed_numbers[home] = parent
		// constructed_numbers[home].str = constructed_numbers[home].str + fmt.Sprintf(num.value) 

	}

	// fmt.Printf("\n\n\n\n Constructed \n %v", constructed_numbers)

	return constructed_numbers
}

func hasSymbol(num Num, symbols []Symbol) bool {
	has_sym := false
	for _, sym := range symbols {

		if(num.row == sym.row){ // same row test
			
			if(num.column + 1 == sym.column) {
				has_sym = true
				break
			} else if (num.column - 1 == sym.column) {
				has_sym = true
				break
			}
		}

		if(num.column == sym.column) { // same column
			if(num.row + 1 == sym.row){
				has_sym = true
				break
			}else if(num.row - 1 == sym.row){
				has_sym = true
				break
			}
		}

		if(num.column + 1 == sym.column){ // first diagonal check
			if(num.row + 1 == sym.row){
				has_sym = true
				break
			}else if(num.row - 1 == sym.row){
				has_sym = true
				break
			}
		}

		if(num.column - 1 == sym.column){ // second diagonal check
			if(num.row + 1 == sym.row){
				has_sym = true
				break
			}else if(num.row - 1 == sym.row){
				has_sym = true
				break
			}
		}

	}

	return has_sym
}

func getGearNumbers(gear Symbol, constructed map[string]ConstructedNumber)int{

	product := 1;
	number_of_times := 0

	for _, _construction := range constructed {

		is_valid := false

		for _, num := range _construction.nums {

			if(num.row == gear.row) { // same row

				if(gear.column + 1 == num.column){
					is_valid = true 
					break
				}else if (gear.column - 1 == num.column){
					is_valid = true 
					break
				}

			}

			if(gear.column + 1 == num.column){ // gear is b4
				if(gear.row - 1 == num.row) {//below
					is_valid = true 
					break
				}else if (gear.row + 1 == num.row) {//above
					is_valid = true 
					break
				}
			}

			if(gear.column - 1 == num.column) { // gear is after
				if(gear.row - 1 == num.row) {// below 
					is_valid = true 
					break
				}else if (gear.row + 1 == num.row){// above
					is_valid = true 
					break
				}
			}

			if(gear.column == num.column){
				if(gear.row - 1 == num.row){
					is_valid = true
					break
				}else if (gear.row + 1 == num.row){
					is_valid = true 
					break
				}
			}

			

		}

		if(is_valid){
			const_num, err := strconv.Atoi(_construction.str); if err != nil {
				panic("Unable to convert construction string into a number")
			}

			product = product * const_num
			number_of_times = number_of_times + 1
		}

	}

	if (number_of_times == 0 || number_of_times == 1) {
		return 0
	}

	return product

}

func Parse(file string){
	file_content_bytes, err := os.ReadFile(file); if err != nil {
		panic("Unable to read file")
	}

	file_content := string(file_content_bytes)

	maze := strings.Split(file_content, "\n") 
	nums := []Num{}
	symbols := []Symbol{}

	all_numbers, all_symbols := Runner(maze, len(maze), len(maze[0]), 0, 0, nums, symbols)

	gears := func ()[]Symbol{
		arr := []Symbol{}

		for _, sym := range all_symbols {
			if(sym.value == "*"){
				arr = append(arr, sym)
			}
		}

		return arr
	}()

	fmt.Print(gears)

	constructed := NumberCollector(all_numbers)

	
	all_gears_product_total := 0 

	for _, gear := range gears {
		product := getGearNumbers(gear, constructed)

		all_gears_product_total = all_gears_product_total + product
	}

	fmt.Printf("\n\n\n\n FINAL ANSWER:: %d \n\n", all_gears_product_total)
	

}