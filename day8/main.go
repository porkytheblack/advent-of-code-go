package day8

import (
	"fmt"
	"os"
	"strings"
	// "time"
)

func RotateDirections(directions []int)[]int{
	start := directions[0]
	endPopped := directions[1:]

	return append(endPopped, start)
}

func ParseFile (file string) ([]int,map[string][]string) {
	buff, err := os.ReadFile(file); if err != nil {
		fmt.Println("Unable to read file")
		panic(err)
	}

	file_content := string(buff)

	parts := strings.Split(file_content, "\n\n")

	firstPart := parts[0]

	secondPart := parts[1]

	journeyMap := make(map[string][]string) 

	directions := []int{}

	for _, direction := range strings.TrimSpace(firstPart) {
		if(string(direction) == "R") {
			directions = append(directions, 1)
		}
		if(string(direction) == "L"){
			directions = append(directions, 0)
		}
	}

	for _, line := range strings.Split(secondPart, "\n") {
		children := []string{}
		parts = strings.Split(line, " = ")
		firstPart := parts[0]
		secondPart := parts[1]
		nodeKey := strings.TrimSpace(firstPart)

		unwrapped := strings.Replace(strings.Replace(secondPart, "(", "", 1), ")", "", 1)

		for _, child := range strings.Split(unwrapped, ","){
			children = append(children, strings.TrimSpace(child))
		}

		journeyMap[nodeKey] = children
	}

	return directions, journeyMap
}

func CountSteps(goal string, journeyMap map[string][]string, path []string, directions []int)int{
	fmt.Println("CURRENT STEP:: ", len(path) -1)
	// fmt.Println("Current Path length:: ", len(path))
	// fmt.Println("PATH::", path)
	// time.Sleep(time.Second * 10)
	top := path[0]
	if(top == goal) {
		return len(path)
	}

	currentlyAt := path[0]
	goTo := directions[0]

	if(goTo == 0) {
		fmt.Println("Going Left")
	}else if(goTo == 1) {
		fmt.Println("Going Right")
	}

	newLocation := journeyMap[currentlyAt][goTo]

	path = append([]string{newLocation}, path...)
	directions = RotateDirections(directions)

	return CountSteps(goal, journeyMap, path, directions)

}


func main(input string, goal string, start string){

	directions, journeyMap := ParseFile(input)

	steps := CountSteps(goal, journeyMap, []string{start}, directions)

	fmt.Println("Steps ::", steps - 1)
}


func Day8(input string, goal string, start string){
	main(input, goal, start)
}