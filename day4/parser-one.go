package day4

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)


type Card struct {
	card_number string 
	value int
	winning []int
	matching []int
	card_num int
	visit_code string
}

func ParseLine(line string)Card{

	winning_nums := []int{}
	matching_nums := []int{}

	first_split := strings.Split(line, ":")
	card_string := first_split[0]

	second_split := strings.Split(strings.TrimSpace(first_split[1]), "|")
	winning := strings.Split(strings.TrimSpace(second_split[0]), " ")
	user_has := strings.Split(strings.TrimSpace(second_split[1]), " ")

	// fmt.Printf("\n\n WINNING:: %v \n\n USER:: %v \n\n", winning, user_has)

	for _, win := range winning {
		if(win == " " || win == "") {
			continue
		}
		num, err := strconv.Atoi(win); if err != nil {
			fmt.Print("WT::", win)
			panic("Unable to convert string to int")
		}

		winning_nums = append(winning_nums, num)
	}

	for _, user_h := range user_has {
		if(user_h == " " || user_h == ""){
			continue
		}
		num, err := strconv.Atoi(user_h); if err != nil {
			fmt.Print("UsT::", user_h)
			panic("Unable to convert string to int")
		}

		exists := func()bool{
			match := false
			for _, win := range winning_nums {
				if(num == win){
					match = true
					break
				}
			}

			return match
		}()

		if(exists){
			matching_nums = append(matching_nums, num)
		}
	}

	value := func () int {

		if(len(matching_nums)> 0){
			return int(math.Pow(2,float64(len(matching_nums)) - 1))
		}

		return 0

	}()

	return Card {
		card_number: card_string,
		value: value,
		winning: winning_nums,
		matching: matching_nums,
	}

}


func Parse(file string){
	file_content_bytes, err := os.ReadFile(file); if err != nil {
		panic("Unable to read file")
	}
	file_content := string(file_content_bytes)

	lines := strings.Split(file_content, "\n")

	cards := []Card{}

	for _, line := range lines {
		card := ParseLine(line)
		fmt.Printf("\n\n %v \n\n", card)
		cards =append(cards, card)
	}


	total := 0 

	for _, card := range cards {
		total += card.value
	}

	fmt.Println("Total:: ", total)


}