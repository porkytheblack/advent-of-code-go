package parsers

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	id int
	rounds []map[string]int
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



func DayTwo1(filepath string, red int, green int, blue int){

	games := []Game{}

	file_content_bytes, err := os.ReadFile(filepath); if err != nil {
		fmt.Println("Unable to read file", err)
	}

	file_content := string(file_content_bytes)

	lines := getLines(file_content)

	for _, line := range lines {
		game_id := getGameID(line)
		rounds := getRounds(line)

		games = append(games, Game{
			id: game_id,
			rounds: rounds,
		})
	}

	valid_games := []Game{}

	for _, game := range games {
		is_valid := isGameValid(game, red, blue, green)

		if(is_valid) {
			fmt.Printf("\n\n Valid Game %v \n\n", game)
			valid_games = append(valid_games, game)
		}
	}

	id_sum := 0;

	for _, game := range valid_games {
		id_sum = id_sum + game.id
	}

	fmt.Println(id_sum)
	
}