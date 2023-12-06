package day4

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)


func ParseLine2(line string)Card{

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

	fmt.Println(card_string)
	fmt.Println("Split", strings.Split(card_string, "Card")[1])
	c_num, err := strconv.Atoi(strings.TrimSpace(strings.Split(card_string, "Card")[1])); if err != nil {
		fmt.Println(strings.Split(card_string, "")[1])
		panic("Unable to get card number")
	}

	return Card {
		card_number: card_string,
		value: value,
		winning: winning_nums,
		matching: matching_nums,
		card_num: c_num,
	}
}

func CompileCard(card Card, cards []Card, count int)int{

	if len(card.matching) == 0 {
		return 0
	} 

	children := func ()[]Card{
		arr := []Card{}

		for i := card.card_num + 1; i <= card.card_num + len(card.matching); i++ {

			for _, _card := range cards {
				if(_card.card_num == i){
					arr = append(arr, _card)
					break;
				}
			}

		}

		return arr
	}()

	child_total := len(children)
	for _, child := range children {
		 sub_ch := CompileCard(child, cards,  child_total)
		 child_total = child_total + sub_ch
	}

	fmt.Printf("\n\n Parent %d Children %d \n\n", card.card_num, child_total)

	return child_total

}


func Parse2(file string)int{
	file_content_bytes, err := os.ReadFile(file); if err != nil {
		panic("Unable to read file")
	}
	file_content := string(file_content_bytes)

	lines := strings.Split(file_content, "\n")

	cards := []Card{}

	for _, line := range lines {
		card := ParseLine2(line)
		// fmt.Printf("\n\n %v \n\n", card)
		cards =append(cards, card)
	}

	total := len(cards)
	
	for _, card := range cards {

		// visited := []string{fmt.Sprintf("%d", card.card_num)}

		total += CompileCard(card, cards, 1)
		// fmt.Println("Total:: ", total)

		// fmt.Printf("\n\nCard Visits :: :: %v \n\n", visits)
		// total = total + len(visits)
		// all_compiled = append(all_compiled, card)

		// for _, compilation := range card_compilations {
		// 	all_compiled = append(all_compiled, compilation)
		// }
	}


	fmt.Println("THE STACK HAS::", total)
	

	return total

}