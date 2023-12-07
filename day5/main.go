package day5

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

type VariableDescription struct {
	source_range_start int
	destination_range_start int
	mapping_range int
}

type SeedDescription struct {
	Soil int `json:"soil"`
	Fertilizer int `json:"fertilizer"`
	Water int	`json:"water"`
	Light int	`json:"light"`
	Temparature int `json:"temparature"`
	Humidity int `json:"humidity"`
	Location int `json:"location"`
	Seed int `json:"seed"`
}

type SD []SeedDescription

func (a SD) Len() int {
	return len(a)
}

func (a SD) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a SD) Less(i, j int) bool {
	return a[i].Location < a[j].Location
}

func deriveValues(l string)[]int{
	arr := []int{}
				
	for _, n := range strings.Split( strings.TrimSpace(l), " ") {
		num, err := strconv.Atoi(strings.TrimSpace(n)); if err != nil {
			fmt.Println("Error", err)
			panic("Unable to convert string to number")
		}
		arr = append(arr, num)
	}

	return arr
}

func ParseInput(file_path string) ([]int, map[string]VariableDescription) {

	buff, err := os.ReadFile(file_path); if err != nil {
		fmt.Println("Something went wrong ::" ,err)
	}

	file_content := string(buff)

	lines := strings.Split(file_content, "\n\n")

	seeds := []int{}

	almanac := make(map[string]VariableDescription)

	for _, line := range lines {

		if(strings.HasPrefix(line, "seeds:")){

			split := strings.Split(line, "seeds:")

			_seeds := deriveValues(split[1])

			for _, seed := range _seeds {
				seeds = append(seeds, seed)
			}		
		}

		if(strings.HasPrefix(line, "seed-to-soil map:")){
			
			for i, l := range strings.Split(line,"\n"){
				if(i != 0) {
					arr := deriveValues(l)
					curr := VariableDescription{
						source_range_start: arr[1],
						destination_range_start: arr[0],
						mapping_range: arr[2],
					}
					almanac[fmt.Sprintf("seed-to-soil-%d", i)] = curr
				}
			}

		}

		if(strings.HasPrefix(line, "soil-to-fertilizer map:")){
			
			for i, l := range strings.Split(line,"\n"){
				if(i != 0) {
					arr := deriveValues(l)
					curr := VariableDescription{
						source_range_start: arr[1],
						destination_range_start: arr[0],
						mapping_range: arr[2],
					}
					almanac[fmt.Sprintf("soil-to-fertilizer-%d", i)] = curr
				}
			}

		}

		if(strings.HasPrefix(line, "fertilizer-to-water map:")){
			
			for i, l := range strings.Split(line,"\n"){
				if(i != 0) {
					arr := deriveValues(l)
					curr := VariableDescription{
						source_range_start: arr[1],
						destination_range_start: arr[0],
						mapping_range: arr[2],
					}
					almanac[fmt.Sprintf("fertilizer-to-water-%d", i)] = curr
				}
			}

		}

		if(strings.HasPrefix(line, "water-to-light map:")){
			
			for i, l := range strings.Split(line,"\n"){
				if(i != 0) {
					arr := deriveValues(l)
					curr := VariableDescription{
						source_range_start: arr[1],
						destination_range_start: arr[0],
						mapping_range: arr[2],
					}
					almanac[fmt.Sprintf("water-to-light-%d", i)] = curr
				}
			}
		}

		if(strings.HasPrefix(line, "light-to-temperature map:")){
			
			for i, l := range strings.Split(line,"\n"){
				if(i != 0) {
					arr := deriveValues(l)
					curr := VariableDescription{
						source_range_start: arr[1],
						destination_range_start: arr[0],
						mapping_range: arr[2],
					}
					almanac[fmt.Sprintf("light-to-temperature-%d", i)] = curr
				}
			}
		}


		if(strings.HasPrefix(line, "temperature-to-humidity map:")){
			
			for i, l := range strings.Split(line,"\n"){
				if(i != 0) {
					arr := deriveValues(l)
					curr := VariableDescription{
						source_range_start: arr[1],
						destination_range_start: arr[0],
						mapping_range: arr[2],
					}
					almanac[fmt.Sprintf("temperature-to-humidity-%d", i)] = curr
				}
			}
		}

		if(strings.HasPrefix(line, "humidity-to-location map:")){
			
			for i, l := range strings.Split(line,"\n"){
				if(i != 0) {
					arr := deriveValues(l)
					curr := VariableDescription{
						source_range_start: arr[1],
						destination_range_start: arr[0],
						mapping_range: arr[2],
					}
					almanac[fmt.Sprintf("humidity-to-location-%d", i)] = curr
				}
			}
		}

	}

	return seeds, almanac
}

func getValue( map_name string, initial int, almanac map[string]VariableDescription)int{
	value := initial
	for key, description := range almanac {
		if(!strings.HasPrefix(strings.TrimSpace(key), map_name)){
			continue
		}

		if(description.source_range_start <= initial && (description.source_range_start + description.mapping_range) >= initial ){

			difference := description.destination_range_start - description.source_range_start // negative or positive

			value = value + difference

			break
		}
		continue
	}

	return value
}

func getLowestLocationValue( seed int, seed_range int, almanac map[string]VariableDescription)int{
			BATCH_SIZE := 1000
			leastLocation :=  make(chan int)
			seedValues := make(chan int)
			var wg sync.WaitGroup

			seedWorker := func(){
				defer wg.Done()

				for seedValue := range seedValues {
					soilValue := getValue("seed-to-soil",seedValue, almanac)
				
					fertilizeValue := getValue("soil-to-fertilizer", soilValue, almanac)
	
					waterValue := getValue("fertilizer-to-water", fertilizeValue, almanac)
	
					lightValue := getValue("water-to-light", waterValue, almanac)
	
					temperatureValue := getValue("light-to-temperature", lightValue, almanac)
	
					humidityValue := getValue("temperature-to-humidity", temperatureValue, almanac)
	
					locationValue := getValue("humidity-to-location", humidityValue, almanac)
					
					leastLocation <- locationValue
				}

			}

			for i := 0; i < 1000; i++ {
				wg.Add(1)
				go seedWorker()
			}

			go func(){
				defer close(seedValues)
				for i := 0; i < seed_range; i += BATCH_SIZE {
					batchEnd := i + BATCH_SIZE
	
					if(batchEnd > seed_range) {
						batchEnd = seed_range
					}
	
					for j := i; j < batchEnd; j++ {
						seedValues <- seed + j
						
					}
	
				}
			}()

			go func(){
				wg.Wait()
				close(leastLocation)
			}()

			leastLocationValue := -1

			for value := range leastLocation {
				if(leastLocationValue == -1 || value < leastLocationValue) {
					leastLocationValue = value
				}
			}

			fmt.Println("Least for ", seed, " IS ", leastLocationValue)

			return leastLocationValue
}



func main(input string){

	seed_mapping, almanac := ParseInput(input)

	leastLocationValue := -1

	for i := 0; i < len(seed_mapping) / 2; i = i + 2 {


		func(seed int, seed_range int, almanac map[string]VariableDescription) {

			locationValue := getLowestLocationValue(seed, seed_range, almanac)

			if(leastLocationValue == -1 || locationValue < leastLocationValue) {
				leastLocationValue = locationValue
			}
		}(seed_mapping[i], seed_mapping[i + 1], almanac)


	}

	
	

	fmt.Println("Leas Location is::", leastLocationValue)

}

func Run (input string) {
	main(input)
}