package n_american_cities

import "testing"

func TestUniqueCities(t *testing.T) {
	prevCities := make(map[string]bool)

	for i, city := range cities {
		key := city.name + city.provState + city.country
		if prevCities[key] {
			t.Errorf("city %d, %v, already found", i+1, city)
		}
		prevCities[key] = true
	}
}

func TestCityAbbrevs(t *testing.T) {
	canProvAbbrevs := make(map[string]bool)
	for _, x := range canadianProvs {
		canProvAbbrevs[x.abbrev] = true
	}

	usStateAbbrevs := make(map[string]bool)
	for _, x := range usStates {
		usStateAbbrevs[x.abbrev] = true
	}

	var found bool
	for i, city := range cities {
		found = false

		switch city.country {
		case "CAN":
			found = canProvAbbrevs[city.provState]
		case "USA":
			found = usStateAbbrevs[city.provState]

		case "MEX": // Don't have abbreviations for Mexican states
			found = true

		default:
			found = false
		}

		if !found {
			t.Errorf("city %d, did not find stateProv for %s in %v", i+1, city.provState, city)
		}
	}
}
