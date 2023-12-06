// https://adventofcode.com/2023/day/5

package main

import (
	"2023/input"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

func main() {
	input := input.FileStringInput("input.txt")

	// input := []string{
	// 	"seeds: 79 14 55 13",
	// 	"",
	// 	"seed-to-soil map:",
	// 	"50 98 2",
	// 	"52 50 48",
	// 	"",
	// 	"soil-to-fertilizer map:",
	// 	"0 15 37",
	// 	"37 52 2",
	// 	"39 0 15",
	// 	"",
	// 	"fertilizer-to-water map:",
	// 	"49 53 8",
	// 	"0 11 42",
	// 	"42 0 7",
	// 	"57 7 4",
	// 	"",
	// 	"water-to-light map:",
	// 	"88 18 7",
	// 	"18 25 70",
	// 	"",
	// 	"light-to-temperature map:",
	// 	"45 77 23",
	// 	"81 45 19",
	// 	"68 64 13",
	// 	"",
	// 	"temperature-to-humidity map:",
	// 	"0 69 1",
	// 	"1 0 69",
	// 	"",
	// 	"humidity-to-location map:",
	// 	"60 56 37",
	// 	"56 93 4",
	// }

	parsed := parseInput(input)

	fmt.Printf("Part1: %d\n", part1(parsed))
	// fmt.Printf("Part2: %d\n", part2(parsed))
}

// almanacRange holds range of numbers that have offsets
type almanacRange struct {
	startSource      int
	endSource        int
	startDestination int
	length           int
}

// gardenAlmanac holds all garden info
type gardenAlmanac struct {
	seeds                 []int
	seedToSoil            []almanacRange
	soilToFertilizer      []almanacRange
	fertilizerToWater     []almanacRange
	waterToLight          []almanacRange
	lightToTemperature    []almanacRange
	temperatureToHumidity []almanacRange
	humidityToLocation    []almanacRange
}

func (g gardenAlmanac) GoString() string {
	var ret string
	ret += fmt.Sprintf("seeds: %v\n", g.seeds)
	ret += fmt.Sprintf("seedToSoil: %v\n", g.seedToSoil)
	ret += fmt.Sprintf("soilToFertilizer: %v\n", g.soilToFertilizer)
	ret += fmt.Sprintf("fertilizerToWater: %v\n", g.fertilizerToWater)
	ret += fmt.Sprintf("waterToLight: %v\n", g.waterToLight)
	ret += fmt.Sprintf("lightToTemperature: %v\n", g.lightToTemperature)
	ret += fmt.Sprintf("temperatureToHumidity: %v\n", g.temperatureToHumidity)
	ret += fmt.Sprintf("humidityToLocation: %v\n", g.humidityToLocation)
	return ret
}

func newAlmanacRange(destination, source, length int) almanacRange {
	var almanacRange almanacRange
	almanacRange.startSource = source
	almanacRange.startDestination = destination
	almanacRange.endSource = source + length - 1
	almanacRange.length = length
	return almanacRange
}

// strItInt converts a string representation of an integer to an integer, fails on error
func strToInt(num string) int {
	intNum, err := strconv.Atoi(num)
	if err != nil {
		log.Fatalln(err)
	}
	return intNum
}

func parseInput(input []string) gardenAlmanac {
	var gardenAlmanac gardenAlmanac
	var mode string
	for i, line := range input {
		// read seeds
		if i == 0 {
			seedStr := strings.Split(line, ": ")[1]
			seedInput := strings.Split(seedStr, " ")
			seeds := make([]int, len(seedInput))
			for j, seed := range seedInput {
				seedInt, _ := strconv.Atoi(seed)
				seeds[j] = seedInt
			}
			gardenAlmanac.seeds = seeds
			continue
		}
		// skip empty lines
		if line == "" {
			mode = ""
			continue
		}
		// set the current mode for scanning
		// fmt.Println(line)
		split := strings.Split(line, " ")
		if split[1] == "map:" {
			mode = line
			continue
		}
		// process incoming text
		var source, destination, length int = strToInt(split[0]), strToInt(split[1]), strToInt(split[2])
		switch mode {
		case "seed-to-soil map:":
			gardenAlmanac.seedToSoil = append(gardenAlmanac.seedToSoil, newAlmanacRange(source, destination, length))
		case "soil-to-fertilizer map:":
			gardenAlmanac.soilToFertilizer = append(gardenAlmanac.soilToFertilizer, newAlmanacRange(source, destination, length))
		case "fertilizer-to-water map:":
			gardenAlmanac.fertilizerToWater = append(gardenAlmanac.fertilizerToWater, newAlmanacRange(source, destination, length))
		case "water-to-light map:":
			gardenAlmanac.waterToLight = append(gardenAlmanac.waterToLight, newAlmanacRange(source, destination, length))
		case "light-to-temperature map:":
			gardenAlmanac.lightToTemperature = append(gardenAlmanac.lightToTemperature, newAlmanacRange(source, destination, length))
		case "temperature-to-humidity map:":
			gardenAlmanac.temperatureToHumidity = append(gardenAlmanac.temperatureToHumidity, newAlmanacRange(source, destination, length))
		case "humidity-to-location map:":
			gardenAlmanac.humidityToLocation = append(gardenAlmanac.humidityToLocation, newAlmanacRange(source, destination, length))
		}
	}
	return gardenAlmanac
}

type seedGuide struct {
	seed        int
	soil        int
	fertilizer  int
	water       int
	light       int
	temperature int
	humidity    int
	location    int
}

func getAlmanacValue(r []almanacRange, target int) int {
	// if target is within an almanacRange
	for _, ran := range r {
		if target >= ran.startSource && target < ran.startSource+ran.length {
			// calculate value and return
			offset := target - ran.startSource
			return ran.startDestination + offset
		}
	}
	// otherwise just return same value
	return target
}

func part1(gardenAlmanac gardenAlmanac) int {
	// fmt.Printf("%#v\n", gardenAlmanac)
	// get values for each seed
	var seedGuides []seedGuide
	for _, seedId := range gardenAlmanac.seeds {
		newSeedGuide := seedGuide{seed: seedId}
		newSeedGuide.soil = getAlmanacValue(gardenAlmanac.seedToSoil, seedId)
		newSeedGuide.fertilizer = getAlmanacValue(gardenAlmanac.soilToFertilizer, newSeedGuide.soil)
		newSeedGuide.water = getAlmanacValue(gardenAlmanac.fertilizerToWater, newSeedGuide.fertilizer)
		newSeedGuide.light = getAlmanacValue(gardenAlmanac.waterToLight, newSeedGuide.water)
		newSeedGuide.temperature = getAlmanacValue(gardenAlmanac.lightToTemperature, newSeedGuide.light)
		newSeedGuide.humidity = getAlmanacValue(gardenAlmanac.temperatureToHumidity, newSeedGuide.temperature)
		newSeedGuide.location = getAlmanacValue(gardenAlmanac.humidityToLocation, newSeedGuide.humidity)
		seedGuides = append(seedGuides, newSeedGuide)
	}
	// fmt.Println(seedGuides)
	// get lowest location value and return that
	var lowLocation int = int(math.Inf(1))
	for _, seed := range seedGuides {
		if seed.location < lowLocation {
			lowLocation = seed.location
		}
	}
	return lowLocation
}

// ==============================================================================================================
// ==============================================================================================================
// ==============================================================================================================
// ==============================================================================================================
// ==============================================================================================================
// ==============================================================================================================
// ==============================================================================================================
// ==============================================================================================================

// type gardenMap struct {
// 	seeds                  []int
// 	seedSoilMap            map[int]int
// 	soilFertilizerMap      map[int]int
// 	fertilizerWaterMap     map[int]int
// 	waterLightMap          map[int]int
// 	lightTemperatureMap    map[int]int
// 	temperatureHumidityMap map[int]int
// 	humidityLocationMap    map[int]int
// }

// func (g gardenMap) GoString() string {
// 	var ret string
// 	ret += fmt.Sprintf("seeds: %v\n", g.seeds)
// 	ret += fmt.Sprintf("seedSoilMap: %v\n", g.seedSoilMap)
// 	ret += fmt.Sprintf("soilFertilizerMap: %v\n", g.soilFertilizerMap)
// 	ret += fmt.Sprintf("fertilizerWaterMap: %v\n", g.fertilizerWaterMap)
// 	ret += fmt.Sprintf("waterLightMap: %v\n", g.waterLightMap)
// 	ret += fmt.Sprintf("lightTemperatureMap: %v\n", g.lightTemperatureMap)
// 	ret += fmt.Sprintf("temperatureHumidityMap: %v\n", g.temperatureHumidityMap)
// 	ret += fmt.Sprintf("humidityLocationMap: %v\n", g.humidityLocationMap)
// 	return ret
// }

// func newGardenMap() *gardenMap {
// 	var g gardenMap
// 	g.seedSoilMap = map[int]int{}
// 	g.soilFertilizerMap = map[int]int{}
// 	g.fertilizerWaterMap = map[int]int{}
// 	g.waterLightMap = map[int]int{}
// 	g.lightTemperatureMap = map[int]int{}
// 	g.temperatureHumidityMap = map[int]int{}
// 	g.humidityLocationMap = map[int]int{}
// 	return &g
// }

// // strItInt converts a string representation of an integer to an integer, fails on error
// func strToInt(num string) int {
// 	intNum, err := strconv.Atoi(num)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	return intNum
// }

// // updateMap adds input offsets to target map
// func updateMap(targetMap map[int]int, inputLine string) {
// 	lineItems := strings.Split(inputLine, " ")
// 	var source, destination, length int = strToInt(lineItems[0]), strToInt(lineItems[1]), strToInt(lineItems[2])
// 	for i := 0; i < length; i++ {
// 		targetMap[destination+i] = source + i
// 	}
// }

// func parseInput(input []string) gardenMap {
// 	// build garden maps based on input
// 	gardenMap := *newGardenMap()
// 	var mode string
// 	for i, line := range input {
// 		// read seeds
// 		if i == 0 {
// 			seedStr := strings.Split(line, ": ")[1]
// 			seedInput := strings.Split(seedStr, " ")
// 			seeds := make([]int, len(seedInput))
// 			for j, seed := range seedInput {
// 				seedInt, _ := strconv.Atoi(seed)
// 				seeds[j] = seedInt
// 			}
// 			gardenMap.seeds = seeds
// 		}
// 		// skip empty lines
// 		if line == "" {
// 			mode = ""
// 			continue
// 		}
// 		// set the current mode for scanning
// 		fmt.Println(line)
// 		split := strings.Split(line, " ")
// 		if split[1] == "map:" {
// 			mode = line
// 			continue
// 		}
// 		// process incoming text
// 		switch mode {
// 		case "seed-to-soil map:":
// 			updateMap(gardenMap.seedSoilMap, line)
// 		case "soil-to-fertilizer map:":
// 			updateMap(gardenMap.soilFertilizerMap, line)
// 		case "fertilizer-to-water map:":
// 			updateMap(gardenMap.fertilizerWaterMap, line)
// 		case "water-to-light map:":
// 			updateMap(gardenMap.waterLightMap, line)
// 		case "light-to-temperature map:":
// 			updateMap(gardenMap.lightTemperatureMap, line)
// 		case "temperature-to-humidity map:":
// 			updateMap(gardenMap.temperatureHumidityMap, line)
// 		case "humidity-to-location map:":
// 			updateMap(gardenMap.humidityLocationMap, line)
// 		}
// 	}
// 	return gardenMap
// }

// // getGardenMapValue returns the value from the map, or the input number if the value is not present in the map
// func getGardenMapValue(m map[int]int, num int) int {
// 	value, exists := m[num]
// 	if exists {
// 		return value
// 	} else {
// 		return num
// 	}
// }

// type seedGuide struct {
// 	seed        int
// 	soil        int
// 	fertilizer  int
// 	water       int
// 	light       int
// 	temperature int
// 	humidity    int
// 	location    int
// }

// func part1(gardenMap gardenMap) int {
// 	fmt.Printf("%#v\n", gardenMap)
// 	// get values for each seed
// 	var seedGuides []seedGuide
// 	for _, seedId := range gardenMap.seeds {
// 		newSeedGuide := seedGuide{seed: seedId}
// 		newSeedGuide.soil = getGardenMapValue(gardenMap.seedSoilMap, seedId)
// 		newSeedGuide.fertilizer = getGardenMapValue(gardenMap.soilFertilizerMap, newSeedGuide.soil)
// 		newSeedGuide.water = getGardenMapValue(gardenMap.fertilizerWaterMap, newSeedGuide.fertilizer)
// 		newSeedGuide.light = getGardenMapValue(gardenMap.waterLightMap, newSeedGuide.water)
// 		newSeedGuide.temperature = getGardenMapValue(gardenMap.lightTemperatureMap, newSeedGuide.light)
// 		newSeedGuide.humidity = getGardenMapValue(gardenMap.temperatureHumidityMap, newSeedGuide.temperature)
// 		newSeedGuide.location = getGardenMapValue(gardenMap.humidityLocationMap, newSeedGuide.humidity)
// 		seedGuides = append(seedGuides, newSeedGuide)
// 	}
// 	// fmt.Println(seedGuides)
// 	// get lowest location value and return that
// 	var lowLocation int = int(math.Inf(1))
// 	for _, seed := range seedGuides {
// 		if seed.location < lowLocation {
// 			lowLocation = seed.location
// 		}
// 	}
// 	return lowLocation
// }
