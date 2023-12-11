package days

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/thomas-marquis/advent-of-code-2023/utils"
)

type MappingRange struct {
	SourceRangeStart int
	DestRangeStart int
	RangeLength int
}

func (m *MappingRange) GetDestValue(sourceValue int) (int, bool) {
	if sourceValue < m.SourceRangeStart || sourceValue > m.SourceRangeStart + m.RangeLength {
		return 0, false
	}
	return m.DestRangeStart + sourceValue - m.SourceRangeStart, true
}

func (m *MappingRange) GetSourceValue(destValue int) (int, bool) {
	if destValue < m.DestRangeStart || destValue > m.DestRangeStart + m.RangeLength {
		return 0, false
	}
	return m.SourceRangeStart + destValue - m.DestRangeStart, true
}

func (m *MappingRange) GetDestRange() []int {
	var destRange []int
	for i := m.DestRangeStart; i < m.DestRangeStart + m.RangeLength; i++ {
		destRange = append(destRange, i)
	}
	return destRange
}

func (m *MappingRange) GetSourceRange() []int {
	var sourceRange []int
	for i := m.SourceRangeStart; i < m.SourceRangeStart + m.RangeLength; i++ {
		sourceRange = append(sourceRange, i)
	}
	return sourceRange
}

type Mapping struct {
	SourceName string
	DestName string
	MappingRanges []MappingRange
	PrevMapping *Mapping
	NextMapping *Mapping
	HasPrevMapping bool
	HasNextMapping bool
}

func (m *Mapping) GetDestValue(sourceValue int) int {
	for _, mappingRange := range m.MappingRanges {
		destValue, ok := mappingRange.GetDestValue(sourceValue)
		if ok {
			return destValue
		}
	}
	return sourceValue
}

func (m *Mapping) GetSourceValue(destValue int) int {
	for _, mappingRange := range m.MappingRanges {
		sourceValue, ok := mappingRange.GetSourceValue(destValue)
		if ok {
			return sourceValue
		}
	}
	return destValue
}

func (m *Mapping) GetSourceValueRecursively(destValue int, sourceName string) (int, bool) {
	if m.SourceName == sourceName {
		return m.GetSourceValue(destValue), true
	}
	if m.HasPrevMapping {
		return m.PrevMapping.GetSourceValueRecursively(destValue, sourceName)
	}
	return 0, false
}


func (m *Mapping) GetMaxDestValue() int {
	var max int
	for _, mappingRange := range m.MappingRanges {
		if mappingRange.DestRangeStart + mappingRange.RangeLength > max {
			max = mappingRange.DestRangeStart + mappingRange.RangeLength
		}
	}
	return max
}

func Day5() {
	
	scanner, _ := utils.ReadFileLines("resources/day5_input")

	var seeds []int
	var seedToSoilMaps []MappingRange
	var soilToFertilizerMaps []MappingRange
	var fertilizerToWaterMaps []MappingRange
	var waterToLightMaps []MappingRange
	var lightToTemperatureMaps []MappingRange
	var temperatureToHumidityMaps []MappingRange
	var humidityToLocationMaps []MappingRange

	var stepName string
	for scanner.Scan() {
		lineContent := scanner.Text()

		seedPrefix := "seeds: "
		if strings.HasPrefix(lineContent, seedPrefix) {
			seedsLinePart := strings.Split(lineContent, seedPrefix)[1]
			seedsAsStr := strings.Split(seedsLinePart, " ")
			for _, seedAsStr := range seedsAsStr {
				s, _ := strconv.Atoi(seedAsStr)
				seeds = append(seeds, s)
			}
			continue
		}

		// seed-to-soil map:
		if lineContent == "seed-to-soil map:" {
			stepName = "seed-to-soil"
			continue
		}
		if stepName == "seed-to-soil" {
			if lineContent == "" { 
				stepName = ""
				continue
			}
			seedToSoilMaps = buildMappingRange(lineContent, seedToSoilMaps)
		}

		// soil-to-fertilizer map:
		if lineContent == "soil-to-fertilizer map:" {
			stepName = "soil-to-fertilizer"
			continue
		}
		if stepName == "soil-to-fertilizer" {
			if lineContent == "" { 
				stepName = ""
				continue
			}
			soilToFertilizerMaps = buildMappingRange(lineContent, soilToFertilizerMaps)
		}

		// fertilizer-to-water map:
		if lineContent == "fertilizer-to-water map:" {
			stepName = "fertilizer-to-water"
			continue
		}
		if stepName == "fertilizer-to-water" {
			if lineContent == "" { 
				stepName = ""
				continue
			}
			fertilizerToWaterMaps = buildMappingRange(lineContent, fertilizerToWaterMaps)
		}

		// water-to-light map:
		if lineContent == "water-to-light map:" {
			stepName = "water-to-light"
			continue
		}
		if stepName == "water-to-light" {
			if lineContent == "" { 
				stepName = ""
				continue
			}
			waterToLightMaps = buildMappingRange(lineContent, waterToLightMaps)
		}

		// light-to-temperature map:
		if lineContent == "light-to-temperature map:" {
			stepName = "light-to-temperature"
			continue
		}
		if stepName == "light-to-temperature" {
			if lineContent == "" { 
				stepName = ""
				continue
			}
			lightToTemperatureMaps = buildMappingRange(lineContent, lightToTemperatureMaps)
		}

		// temperature-to-humidity map:
		if lineContent == "temperature-to-humidity map:" {
			stepName = "temperature-to-humidity"
			continue
		}
		if stepName == "temperature-to-humidity" {
			if lineContent == "" { 
				stepName = ""
				continue
			}
			temperatureToHumidityMaps = buildMappingRange(lineContent, temperatureToHumidityMaps)
		}

		// humidity-to-location map:
		if lineContent == "humidity-to-location map:" {
			stepName = "humidity-to-location"
			continue
		}
		if stepName == "humidity-to-location" {
			if lineContent == "" { 
				stepName = ""
				continue
			}
			humidityToLocationMaps = buildMappingRange(lineContent, humidityToLocationMaps)
		}
	}

	seedToSoil := Mapping{
		SourceName: "seeds",
		DestName: "soil",
		MappingRanges: seedToSoilMaps,
		HasPrevMapping: false,
		HasNextMapping: true,
	}
	soilToFertilizer := Mapping{
		SourceName: "soil",
		DestName: "fertilizer",
		MappingRanges: soilToFertilizerMaps,
		HasPrevMapping: true,
		HasNextMapping: true,
	}
	fertilizerToWater := Mapping{
		SourceName: "fertilizer",
		DestName: "water",
		MappingRanges: fertilizerToWaterMaps,
		HasPrevMapping: true,
		HasNextMapping: true,
	}
	waterToLight := Mapping{
		SourceName: "water",
		DestName: "light",
		MappingRanges: waterToLightMaps,
		HasPrevMapping: true,
		HasNextMapping: true,
	}
	lightToTemperature := Mapping{
		SourceName: "light",
		DestName: "temperature",
		MappingRanges: lightToTemperatureMaps,
		HasPrevMapping: true,
		HasNextMapping: true,
	}
	temperatureToHumidity := Mapping{
		SourceName: "temperature",
		DestName: "humidity",
		MappingRanges: temperatureToHumidityMaps,
		HasPrevMapping: true,
		HasNextMapping: true,
	}
	humidityToLocation := Mapping{
		SourceName: "humidity",
		DestName: "location",
		MappingRanges: humidityToLocationMaps,
		HasPrevMapping: true,
		HasNextMapping: false,
	}

	seedToSoil.NextMapping = &soilToFertilizer
	soilToFertilizer.PrevMapping = &seedToSoil
	soilToFertilizer.NextMapping = &fertilizerToWater
	fertilizerToWater.PrevMapping = &soilToFertilizer
	fertilizerToWater.NextMapping = &waterToLight
	waterToLight.PrevMapping = &fertilizerToWater
	waterToLight.NextMapping = &lightToTemperature
	lightToTemperature.PrevMapping = &waterToLight
	lightToTemperature.NextMapping = &temperatureToHumidity
	temperatureToHumidity.PrevMapping = &lightToTemperature
	temperatureToHumidity.NextMapping = &humidityToLocation
	humidityToLocation.PrevMapping = &temperatureToHumidity

	for i := 0; i < humidityToLocation.GetMaxDestValue(); i++ {
		res, ok := humidityToLocation.GetSourceValueRecursively(i, "seeds")
		if ok && utils.IsIn(res, seeds) {
			fmt.Printf("destination %d => seed %d\n", i, res)
			break
		}
	}

}

func buildMappingRange(lineContent string, mappingRanges []MappingRange) []MappingRange {
	valuesStr := strings.Split(lineContent, " ")
	values := utils.ToIntSlice(valuesStr)

	return append(mappingRanges, MappingRange{
		SourceRangeStart: values[1],
		DestRangeStart: values[0],
		RangeLength: values[2],
	})
}