// https://adventofcode.com/2023/day/5

package main

import (
	"2023/input"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
	"sync"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
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
	fmt.Printf("Part2: %d\n", part2(parsed))
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
		newSeedGuide.soil = getAlmanacValue(gardenAlmanac.seedToSoil, newSeedGuide.seed)
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

// perform calculations for seed ranges in place with goroutines instead of saving to memory because this balloons to 62 gigs
func part2(gardenAlmanac gardenAlmanac) int {
	// fmt.Printf("%#v\n", gardenAlmanac)
	p := message.NewPrinter(language.English) // printer for adding commas for large numbers
	var wg sync.WaitGroup
	// fix seeds to read as ranges
	var seedStart int
	// slice to hold return values of all goroutines
	var lowLocations []int = make([]int, len(gardenAlmanac.seeds)/2)
	fmt.Println("Starting Calculations for", len(gardenAlmanac.seeds)/2, "entries...")
	for i, seed := range gardenAlmanac.seeds {
		// even entries are seed starting positions, odd entries are ranges
		if i%2 == 0 {
			seedStart = seed
		} else {
			p.Printf("%d: %d seeds.\n", i/2, seed)
			// use goroutines to parallelize calculations
			wg.Add(1)
			go func(id, seedStart, seed int, lowLocations *[]int) {
				defer wg.Done()
				var localLowLocation int = int(math.Inf(1))
				for j := 0; j < seed; j++ {
					newSeedGuide := seedGuide{seed: seedStart + j}
					newSeedGuide.soil = getAlmanacValue(gardenAlmanac.seedToSoil, newSeedGuide.seed)
					newSeedGuide.fertilizer = getAlmanacValue(gardenAlmanac.soilToFertilizer, newSeedGuide.soil)
					newSeedGuide.water = getAlmanacValue(gardenAlmanac.fertilizerToWater, newSeedGuide.fertilizer)
					newSeedGuide.light = getAlmanacValue(gardenAlmanac.waterToLight, newSeedGuide.water)
					newSeedGuide.temperature = getAlmanacValue(gardenAlmanac.lightToTemperature, newSeedGuide.light)
					newSeedGuide.humidity = getAlmanacValue(gardenAlmanac.temperatureToHumidity, newSeedGuide.temperature)
					newSeedGuide.location = getAlmanacValue(gardenAlmanac.humidityToLocation, newSeedGuide.humidity)
					if newSeedGuide.location < localLowLocation {
						localLowLocation = newSeedGuide.location
					}
				}
				(*lowLocations)[id] = localLowLocation
				fmt.Print(id, " ")
			}(i/2, seedStart, seed, &lowLocations)
		}
	}
	fmt.Print("done: ")
	wg.Wait()
	fmt.Println()

	// calculate lowest location from goroutines
	var lowLocation int = int(math.Inf(1))
	// fmt.Println(lowLocations)
	for _, loc := range lowLocations {
		if lowLocation > loc {
			lowLocation = loc
		}
	}

	return lowLocation
}
