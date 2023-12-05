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

func CompileCard(card Card, cards []Card, visited []string)[]string{

	if(len(card.matching) == 0) {

		return visited
	}
	
	for current_prize_card := card.card_num + 1; current_prize_card <= card.card_num + len(card.matching); current_prize_card ++ {
		
		_card := func()Card{
			for _, c := range cards {
				if(c.card_num == current_prize_card){
					return c
				}
			}
			panic("Unable to find card")
		}()

		visit_code_prefix := func ()string{
			if(card.visit_code == ""){
				return fmt.Sprintf("%d", card.card_num)
			}
			return card.visit_code
		}()

		visit_code := fmt.Sprintf("%s-%d",visit_code_prefix, _card.card_num)
		_card.visit_code = visit_code

		
		visits := CompileCard(_card, cards, visited)

			

			for _, visit := range visits {
				// if(visit == "1-2-4-5 "){
					// 	fmt.Println("Exists::", visit)
					// }
				already_done := func()bool {
					fmt.Println("VISIT:: ", visit)
					for _, code := range visited {
						if(code == visit) {
							return true
						}
					}
					return false
				}()

				if(!already_done){
					visited = append(visited, visit)
				}
			}

			visited = append(visited, visit_code)
		


	}

	// fmt.Printf("\n\nVisits Count: %d Visits: %v\n\n", len(visited), visited)

	return visited

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

	total := 0 
	
	for _, card := range cards {

		visited := []string{fmt.Sprintf("%d", card.card_num)}

		visits := CompileCard(card, cards, visited)

		fmt.Printf("\n\nCard Visits :: :: %v \n\n", visits)
		total = total + len(visits)
		// all_compiled = append(all_compiled, card)

		// for _, compilation := range card_compilations {
		// 	all_compiled = append(all_compiled, compilation)
		// }
	}


	fmt.Println("THE STACK HAS::", total)
	

	return total

}