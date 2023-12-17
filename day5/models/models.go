package models

import "sort"

type Range struct {
	DestStart   int `json:"destStart"`
	SourceStart int `json:"sourceStart"`
	RangeLen    int `json:"rangeLen"`
}

type Map struct {
	SeedToSoil            []Range `json:"seed_to_soil"`
	SoilToFertilizer      []Range `json:"soil_to_fertilizer"`
	FertilizerToWater     []Range `json:"fertilizer_to_water"`
	WaterToLight          []Range `json:"water_to_light"`
	LightToTemperature    []Range `json:"light_to_temperature"`
	TemperatureToHumidity []Range `json:"temperature_to_humidity"`
	HumidityToLocation    []Range `json:"humidity_to_location"`
}

func (m *Map) SortRanges() {
	sort.Slice(m.SeedToSoil, func(i, j int) bool {
		return m.SeedToSoil[i].SourceStart < m.SeedToSoil[j].SourceStart
	})
	sort.Slice(m.SoilToFertilizer, func(i, j int) bool {
		return m.SoilToFertilizer[i].SourceStart < m.SoilToFertilizer[j].SourceStart
	})
	sort.Slice(m.FertilizerToWater, func(i, j int) bool {
		return m.FertilizerToWater[i].SourceStart < m.FertilizerToWater[j].SourceStart
	})
	sort.Slice(m.WaterToLight, func(i, j int) bool {
		return m.WaterToLight[i].SourceStart < m.WaterToLight[j].SourceStart
	})
	sort.Slice(m.LightToTemperature, func(i, j int) bool {
		return m.LightToTemperature[i].SourceStart < m.LightToTemperature[j].SourceStart
	})
	sort.Slice(m.TemperatureToHumidity, func(i, j int) bool {
		return m.TemperatureToHumidity[i].SourceStart < m.TemperatureToHumidity[j].SourceStart
	})
	sort.Slice(m.HumidityToLocation, func(i, j int) bool {
		return m.HumidityToLocation[i].SourceStart < m.HumidityToLocation[j].SourceStart
	})
}
