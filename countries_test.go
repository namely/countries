package main

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type testSuite struct {
	suite.Suite
}

func TestTestSuite(t *testing.T) {
	suite.Run(t, new(testSuite))
}

func (s *testSuite) SetupTest() {
}

func (s *testSuite) Test_GetCountry() {
	tests := map[string]struct {
		countryCode    string
		expectedResult interface{}
	}{
		"country exists": {
			countryCode:    "ES",
			expectedResult: map[string]interface{}{"address_format": "{{recipient}}\n{{street}}\n{{postalcode}} {{city}} {{region}}\n{{country}}", "alpha2": "ES", "alpha3": "ESP", "continent": "Europe", "country_code": "34", "currency": "EUR", "international_prefix": "00", "ioc": "ESP", "languages": []interface{}{"es"}, "latitude": "40 00 N", "longitude": "4 00 W", "name": "Spain", "names": []interface{}{"Spain", "Spanien", "Espagne", "España"}, "national_destination_code_lengths": []interface{}{2}, "national_number_lengths": []interface{}{9}, "national_prefix": "None", "nationality": "Spanish", "number": "724", "region": "Europe", "subdivision_type": "Province", "subregion": "Southern Europe", "un_locode": "ES"},
		},
		"country does not exist": {
			countryCode:    "LALALAND",
			expectedResult: nil,
		},
	}
	for name, t := range tests {
		s.Run(name, func() {
			res := GetCountry(t.countryCode)
			s.Assert().Equal(t.expectedResult, res)
		})
	}
}

func (s *testSuite) Test_HasSubdivisions() {
	tests := map[string]struct {
		countryCode    string
		expectedResult bool
	}{
		"has subdivisions": {
			countryCode:    "ES",
			expectedResult: true,
		},
		"no subdivisions": {
			countryCode:    "SG",
			expectedResult: false,
		},
	}
	for name, t := range tests {
		s.Run(name, func() {
			res := HasSubdivisions(t.countryCode)
			s.Assert().Equal(t.expectedResult, res)
		})
	}
}

func (s *testSuite) Test_GetSubdivisionsName() {
	tests := map[string]struct {
		countryCode    string
		expectedResult []string
	}{
		"has subdivisions - ES": {
			countryCode:    "ES",
			expectedResult: []string{"A Coruña", "Albacete", "Alicante", "Almería", "Asturias", "Badajoz", "Baleares", "Barcelona", "Burgos", "Cantabria", "Castellón", "Ceuta", "Ciudad Real", "Cuenca", "Cáceres", "Cádiz", "Córdoba", "Girona", "Granada", "Guadalajara", "Guipúzcoa", "Huelva", "Huesca", "Jaén", "La Rioja", "Las Palmas", "León", "Lleida", "Lugo", "Madrid", "Melilla", "Murcia", "Málaga", "Navarra", "Ourense", "Palencia", "Pontevedra", "Salamanca", "Santa Cruz de Tenerife", "Segovia", "Sevilla", "Soria", "Tarragona", "Teruel", "Toledo", "Valencia", "Valladolid", "Vizcaya", "Zamora", "Zaragoza", "Álava", "Ávila"},
		},
		"has subdivisions - BB": {
			countryCode:    "BB",
			expectedResult: []string{"Christ Church", "Saint Andrew", "Saint George", "Saint James", "Saint John", "Saint Joseph", "Saint Lucy", "Saint Michael", "Saint Peter", "Saint Philip", "Saint Thomas"},
		},
		"several names": {
			countryCode:    "BE",
			expectedResult: []string{"Antwerpen (nl)", "Brabant Wallon (fr)", "Brussels", "Hainaut (fr)", "Limburg (nl)", "Liège (fr)", "Luxembourg (fr)", "Namur (fr)", "Oost-Vlaanderen (nl)", "Vlaams Brabant (nl)", "West-Vlaanderen (nl)"},
		},
		"no subdivisions": {
			countryCode:    "SG",
			expectedResult: nil,
		},
	}
	for name, t := range tests {
		s.Run(name, func() {
			res := GetSubdivisionsName(t.countryCode)
			s.Assert().Equal(t.expectedResult, res)
		})
	}
}
