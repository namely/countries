package countries

import (
	"fmt"
	"io/ioutil"
	"sort"

	"gopkg.in/yaml.v3"
)

// GetCountry returns the country data
func GetCountry(countryCode string) interface{} {
	path := "lib/data/countries.yaml"
	yfile, err := ioutil.ReadFile(path)
	if err != nil {
		return false
	}

	data := make(map[string]interface{})

	err = yaml.Unmarshal(yfile, &data)
	if err != nil {
		return false
	}

	return data[countryCode]
}

// HasSubdivisions returns true or false for a country
func HasSubdivisions(countryCode string) bool {
	path := fmt.Sprintf("lib/data/subdivisions/%s.yaml", countryCode)
	yfile, err := ioutil.ReadFile(path)
	if err != nil {
		return false
	}

	data := make(map[interface{}]interface{})

	err = yaml.Unmarshal(yfile, &data)
	return err == nil
}

// GetSubdivisionsName gets subdivisions name property for a specific country
func GetSubdivisionsName(countryCode string) []string {
	path := fmt.Sprintf("lib/data/subdivisions/%s.yaml", countryCode)
	yfile, err := ioutil.ReadFile(path)
	if err != nil {
		return nil
	}

	data := make(map[string]map[string]interface{})

	err = yaml.Unmarshal(yfile, &data)
	if err != nil {
		return nil
	}
	result := make([]string, 0)
	for _, v := range data {
		result = append(result, v["name"].(string))
	}

	// sort result to be consistent
	sort.Strings(result)

	return result
}
