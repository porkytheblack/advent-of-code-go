package parsers

import (
	"fmt"
	"os"
)



func DayTwo2(filepath string, red int, green int, blue int){

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