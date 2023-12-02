package parsers2

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type ByRed []map[string]int

func (a ByRed) Len() int {
	return len(a)
}


type Game struct {
	id int
	rounds []map[string]int
	power int
}

func getLines(file_content string) []string {
	lines := strings.Split(file_content, "\n")
	return lines
}

func getGameID(line string) int {
	_split := strings.Split(line, ":")
	first := _split[0] //Game 1

	new_split := strings.Split(first, " ");

	num := new_split[1] // "1"
	// fmt.Println("Game ::",num)
	_int, err := strconv.Atoi(num); if(err != nil){
		panic(err)
	}

	return _int // 1
}

func getRounds(line string) []map[string]int {
	
	_split := strings.Split(line, ":")
	second := _split[1] //  3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
	rounds := strings.Split(second, "; ")

	all_rounds := make([]map[string]int, 0)

	for round_index, round := range rounds {
		// 3 blue, 4 red
		balls := strings.Split(round, ", ") // ["3 blue", "4 red"]

		round_map := make(map[string]int)

		for _, ball := range balls {
			// fmt.Println("Ball ", ball)
			sp := strings.Split(strings.TrimSpace(ball), " ")
			_int, err := strconv.Atoi(sp[0]); if err != nil {
				fmt.Println("Something went wrong")
				panic(err)
 			}
			round_map[sp[1]] = _int
		}

		round_map["round_index"] = round_index

		all_rounds = append(all_rounds, round_map)
	}	

	return all_rounds
}

func isGameValid(game Game, red int, blue int, green int) bool {

	highest_red := 0;
	highest_blue := 0; 
	highest_green := 0; 

	for _, round := range game.rounds {

		if(round["red"] > highest_red){
			highest_red = round["red"]
		}
		if(round["blue"] > highest_blue){
			highest_blue = round["blue"]
		}
		if(round["green"] > highest_green){
			highest_green = round["green"]
		}
	}

	

	return red >= highest_red && blue >= highest_blue && green >= highest_green
}

func SortByField(slice []map[string]int, field string) {
	sort.SliceStable(slice, func(i, j int) bool {
		// if(slice[i][field] == 0){
		// 	fmt.Println(slice[i][field])
		// 	return false
		// }
		return (slice[i][field] > slice[j][field])
	})
}

func getLowest(rounds []map[string]int, clr string)int{

	value := 1;

	for _, round := range rounds {
		val := round[clr]
		if(val != 0){
			value = val
			break
		}
	}

	return value

}

func getPower(game Game)int{
	lowest_red := 1;
	lowest_blue := 1; 
	lowest_green := 1; 

	SortByField(game.rounds, "red")
	lowest_red = getLowest(game.rounds, "red")
	SortByField(game.rounds, "green")
	lowest_green = getLowest(game.rounds, "green")
	SortByField(game.rounds, "blue")
	lowest_blue = getLowest(game.rounds, "blue")

	fmt.Printf("\n\n GAME: %d LOWESTS %d %d %d \n\n", game.id, lowest_red, lowest_green, lowest_blue)

	power :=  lowest_red * lowest_green * lowest_blue

	fmt.Println("POWER:: ", power)

	return power
}

func getTotalPowers(games []Game)int{
	total := 0;

	for _, game := range games {
		total = total + game.power
	}

	return total
}



func DayTwo2(filepath string){

	games := []Game{}

	file_content_bytes, err := os.ReadFile(filepath); if err != nil {
		fmt.Println("Unable to read file", err)
	}

	file_content := string(file_content_bytes)

	lines := getLines(file_content)

	for _, line := range lines {
		game_id := getGameID(line)
		rounds := getRounds(line)
		game := Game{
			id: game_id,
			rounds: rounds,
		}
		game.power = getPower(game)
		games = append(games, game)
	}

	total := getTotalPowers(games)

	fmt.Println(total)
	
}