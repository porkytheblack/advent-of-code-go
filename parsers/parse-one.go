package parsers

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func DayOne1(file_path string){
	data, err := os.ReadFile(file_path); if err != nil {
		log.Fatal(err)
	}

	as_string := string(data)

	lines := strings.Split(as_string, "\n")

	line_numbers := []string{}

	for _, line := range lines {
		
		first_int := '0';
		last_int := '0';

		for i := 0; i < len(line); i++ {
			character := line[i]

			if(unicode.IsDigit(rune(character))){
				first_int = rune(character)
				break
			}
		}

		for j := len(line) - 1; j >= 0; j-- {
			character := line[j]

			if(unicode.IsDigit(rune(character))){
				last_int = rune(character)
				break
			}
		}

		combined := string(first_int) + string(last_int)

		line_numbers = append(line_numbers, combined)
	}

	// fmt.Printf("%v",line_numbers)

	total := 0;

	for i := 0; i < len(line_numbers); i++ {
		num, err := strconv.Atoi(line_numbers[i]); if err != nil {
			log.Fatal(err)
		}
		total = total + num
	}

	fmt.Printf("%d", total)
 
}