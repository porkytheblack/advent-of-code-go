package parsers

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type NI struct {
	number string 
	index int
}

type ByIndex []NI 

func (a ByIndex) Len() int {
	return len(a)
}

func (a ByIndex) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByIndex) Less(i, j int) bool {
	return a[i].index < a[j].index
}

func findAllSubstringIndices(s, substr string) []int {
	var indexes []int
	startIndex := 0

	for {
		index := strings.Index(s[startIndex:], substr)
		if index == -1 {
			break
		}
		indexes = append(indexes, startIndex+index)
		startIndex += index + len(substr)
	}

	return indexes
}

var valid_numbers = [20]string{
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
	"zero",
	"1",
	"2",
	"3",
	"4",
	"5",
	"6",
	"7",
	"8",
	"9",
	"0",
}


func getIntString(num string) string {
	if(num == "zero") {
		return "0"
	}else if(num == "one"){
		return "1"
	}else if(num == "two"){
		return "2"
	}else if(num == "three"){
		return "3"
	}else if(num == "four"){
		return "4"
	}else if(num == "five"){
		return "5"
	}else if(num == "six"){
		return "6"
	}else if(num == "seven"){
		return "7"
	}else if(num == "eight"){
		return "8"
	}else if(num == "nine"){
		return "9"
	}else if(num == "1"){
		return "1"
	}else if(num == "2"){
		return "2"
	}else if(num == "3"){
		return "3"
	}else if(num == "4"){
		return "4"
	}else if(num == "5"){
		return "5"
	}else if(num == "6"){
		return "6"
	}else if(num == "7"){
		return "7"
	}else if(num == "8"){
		return "8"
	}else if(num == "9"){
		return "9"
	}else if(num == "0"){
		return "0"
	}
	return ""
}

func getFirst (st string) string {

	valid := []NI{}

	indexMap := make(map[string][]int)

	for _, v := range valid_numbers {
		indexMap[v] = findAllSubstringIndices(st, v)
	}

	for element, indices := range indexMap {
		if(len(indices) > 0){
			for _, index := range indices {
				valid = append(valid, NI{
					number: element,
					index: index,
				})

			}
		}
	}

	sort.Sort(ByIndex(valid))


	if(len(valid) > 0) {
		return getIntString(valid[0].number)
	}


	return "0"
}


func getLast (st string) string {
	valid := []NI{}

	indexMap := make(map[string][]int)
	for _, v := range valid_numbers {
		indexMap[v] = findAllSubstringIndices(st, v)
	}

	for element, indices := range indexMap {
		if(len(indices) > 0){
			for _, index := range indices {
				valid = append(valid, NI{
					number: element,
					index: index,
				})

			}
		}
	}
	sort.Sort(ByIndex(valid))

	fmt.Printf("%v",valid)

	if(len(valid) > 0) {
		last_index := len(valid) - 1
		return getIntString(valid[last_index].number)
	}

	return "0"
}

func DayOne2(file_path string){
	data, err := os.ReadFile(file_path); if err != nil {
		log.Fatal(err)
	}

	as_string := string(data)

	lines := strings.Split(as_string, "\n")

	line_numbers := []string{}

	for _, line := range lines {
		
		first_int := getFirst(line);
		last_int := getLast(line);

		combined := string(first_int) + string(last_int)

		line_numbers = append(line_numbers, combined)
	}

	fmt.Printf("\n%v\n",line_numbers)

	total := 0;

	for i := 0; i < len(line_numbers); i++ {
		num, err := strconv.Atoi(line_numbers[i]); if err != nil {
			log.Fatal(err)
		}
		total = total + num
	}

	fmt.Printf("%d", total)
 
}