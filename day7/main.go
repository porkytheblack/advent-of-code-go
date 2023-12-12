package day7

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// A  K  Q  J  T  9 8 7 6 5 4 3 2
// 14 13 12 11 10 9 8 7 6 5 4 3 2

type Hand struct {
	Kind int
	Value int
	Bid int
	Representation string
}

type ByHand []Hand

func (arr ByHand) Len() int {return len(arr)}

func (arr ByHand) Swap(i, j int) {arr[i], arr[j] = arr[j], arr[i]}

func (arr ByHand) Less (i, j int) bool {
	if(arr[i].Kind != arr[j].Kind) {
		return arr[i].Kind < arr[j].Kind
	}
	
	if(arr[i].Representation == arr[j].Representation){
		fmt.Println("Representation is the same::", arr[i].Representation , " ", arr[j].Representation)
		return false
	}
	for index, c := range arr[i].Representation {
		cValue := GetValue(string(c))
		dValue := GetValue(string(arr[j].Representation[index]))

		if(cValue != dValue){
			return cValue < dValue
		} else {
			continue
		}
	}
	return false
}

func RemoveDuplicates(card string)[]string{
	seen := make(map[rune]bool)
	result := ""

	for _, char := range card {
		if !seen[char] {
			result = result + string(char)
			seen[char] = true
		}
	}

	return strings.Split(result, "")
}

func GetValue(value string)int{
	if(value == "A") {
		return 14
	}else if(value == "K"){
		return 13
	}else if(value == "Q"){
		return 12
	}else if(value == "J"){
		return 0 // J is now joker
	}else if(value == "T"){
		return 10
	}
	num, err := strconv.Atoi(value); if err != nil {
		fmt.Printf("Unable to convert number ")
		panic(err)
	}

	return num
}

func GetCardKind(cards string) int {

	/**
	7 - Five of a kind
	6 - Four of a kind
	5 - Full house
	4 - Three of a kind
	3 - Two Pair
	2 - One Pair
	1 - High Card
	0 - No kind 
	**/

	cardNumMap := make(map[string]int)

	for _, card := range cards {

		mapValue := cardNumMap[string(card)]

		newValue := mapValue + 1

		cardNumMap[string(card)] = newValue
	}

	uniqueCards := RemoveDuplicates(cards)

	if(len(uniqueCards) == 1){
		return 7 // five of a kind
	}else if(len(uniqueCards) == 2){
		for _, c := range uniqueCards {
			// four of a kind
			if(cardNumMap[c] == 4) {
				return 6
			}
			// full house
			if(cardNumMap[c] == 3){
				// get the other card
				other := func()string{
					for _, _c := range uniqueCards {
						if(_c != c) {
							return _c
						}
					}
					return "none"
				}()

				if(cardNumMap[other] == 2){
					return 5
				}
			}


		}
	}else if(len(uniqueCards) == 3){

		// 3 of a kind
		for _, c := range uniqueCards {
			if(cardNumMap[c] == 3){
				others := func()[]string{
					arr := []string{}

					for _, _c := range uniqueCards {
						if(_c != c) {
							arr  = append(arr, _c)
						}
					}

					return arr
				}()
				// 3 of a kind
				if(len(others) == 2) {
					return 4
				}
			}

		}

		// 2 pair
		pairs := 0 

		for _, c := range uniqueCards {
			if(cardNumMap[c] == 2){
				pairs += 1
			}
		}

		if(pairs == 2) {
			return 3
		}

	} else if len(uniqueCards) == 4{
			return 2
	} else if(len(uniqueCards) == 5){
			return 1
	}
	fmt.Println("No Card Kind::", uniqueCards)
	return 0
}

func ParseFile(file string)[]string{
	buff, err := os.ReadFile(file); if err != nil {
		panic("Something went wrong")
	}
	file_content := string(buff)

	lines := strings.Split(file_content, "\n")

	return lines
}


func ParseLine(line string)Hand {
	parts := strings.Split(line, " ")
	firstPart := parts[0]
	lastPart := parts[1]

	values := []int{}


	for _, cardValue := range strings.Split(firstPart, "") {
		value := GetValue(cardValue)
		values = append(values, value)
	}

	accumulated_value := 0

	for _, v := range values {
		accumulated_value += v
	}

	bidAmount, err := strconv.Atoi(lastPart); if err != nil {
		fmt.Println("Unable to convert bid amount to number")
		panic(err)
	}

	return Hand {
		Kind: GetCardKind(firstPart),
		Value: accumulated_value,
		Bid: bidAmount,
		Representation: firstPart,
	}
}


func main(input string) {

	hands := ByHand{}
	lines := ParseFile(input)

	for _, line := range lines {
		hand := ParseLine(line)

		hands = append(hands, hand)
	}

	sort.Sort(hands)

	sumOfProduct := 0

	for index, hand := range hands {
		fmt.Printf("\n\n Hand %v times %d ", hand, index + 1)
		sumOfProduct += (index + 1) * hand.Bid
	}

	fmt.Println("\n\nResult:: ", sumOfProduct)
}