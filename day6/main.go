package day6

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Race struct {
	duration int
	record_distance int
}

func solveQuad(race Race) (float64, float64) {

	b := float64(race.duration)
	c := float64(race.record_distance)

	ans1 := (b + math.Sqrt(math.Pow(b, 2) - 4 * 1 * c)) / 2
	ans2 := (b - math.Sqrt(math.Pow(b, 2) - 4 * 1 * c)) / 2


	fmt.Println("SOLUTION FOR S FOR DISTANCE 1: ", c, " ", ans1)
	fmt.Println("SOLUTION FOR S FOR DISTANCE 2: ", c, " ", ans2)

	return math.Ceil(ans1),math.Floor(ans2)
}


func ParseInput(file string) []Race {
	races := []Race{}
	buff, err := os.ReadFile(file); if err != nil {
		fmt.Println("Unable to read file")
		panic(err)
	}

	file_content := string(buff)

	lines := strings.Split(file_content, "\n")

	times := lines[0]
	distances := lines[1]

	timesParts := strings.Split(times, "Time:")
	timeNumbers := func ()[]int{
		agg  := ""
		for _, value := range strings.Split(strings.TrimSpace(timesParts[1]), " "){
			if(value != ""){
				agg += strings.TrimSpace(value)
			}
		}
		num, err := strconv.Atoi(agg); if err != nil {
			panic("Unable to resove time value")
		}
		return []int{num}
	}() 


	distanceParts := strings.Split(distances, "Distance:")
	distanceNumbers := func ()[]int{
		agg  := ""
		for _, value := range strings.Split(strings.TrimSpace(distanceParts[1]), " "){
			if(value != ""){
				agg += strings.TrimSpace(value)
			}
		}
		num, err := strconv.Atoi(agg); if err != nil {
			panic("Unable to resove time value")
		}
		return []int{num}
	}()

	for i, duration := range timeNumbers {
		races = append(races, Race{
			duration: duration,
			record_distance: distanceNumbers[i],
		})
	}

	return races
}


func main(file string){
	races := ParseInput(file)


	product := 1
	for _,race := range races {
		higher_end, lower_end := solveQuad(race)
		ans := int(higher_end - lower_end -1)

		product = product * ans
	}

	fmt.Println("Answer:: ", product)
}