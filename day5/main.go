package day5

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
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

func getSeeds( seeds_with_ranges []int )[]int{

	seeds := []int{}

	for i := 0; i < len(seeds_with_ranges); i++ {
		
		if (i % 2 == 0) {
			for j := 0; j < seeds_with_ranges[i + 1]; j++ {
				fmt.Println("CURRENTLY ON SEED::", seeds_with_ranges[i] + j)
				seeds = append(seeds, seeds_with_ranges[i] + j)

			}

		}


	}


	return seeds
}



func main(input string){

	seedData := SD{}

	seed_mapping, almanac := ParseInput(input)

	seeds := getSeeds(seed_mapping)

	for _, s := range seeds {

		soilValue := getValue("seed-to-soil",s, almanac)
		
		fertilizeValue := getValue("soil-to-fertilizer", soilValue, almanac)

		waterValue := getValue("fertilizer-to-water", fertilizeValue, almanac)

		lightValue := getValue("water-to-light", waterValue, almanac)

		temperatureValue := getValue("light-to-temperature", lightValue, almanac)

		humidityValue := getValue("temperature-to-humidity", temperatureValue, almanac)

		locationValue := getValue("humidity-to-location", humidityValue, almanac)

		_seed := SeedDescription{
			Soil: soilValue,
			Fertilizer: fertilizeValue,
			Water: waterValue,
			Light: lightValue,
			Temparature: temperatureValue,
			Humidity: humidityValue,
			Location: locationValue,
			Seed: s,
		}

		seedData = append(seedData, _seed)
	}

	sort.Sort(seedData)

	

	data, err := json.Marshal(seedData); if err != nil {
		fmt.Println("Error:", err)
		panic("Unable to marshal data")
	}


	fmt.Printf("\n\n\nData %v \n\n\n", string(data))


	fmt.Println("Answe:: ", seedData[0].Location)

}

func Run (input string) {
	main(input)
}