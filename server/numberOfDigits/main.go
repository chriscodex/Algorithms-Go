package main

import "strconv"

type reqStruct struct {
	Number int `json:"number"`
}

type resStruct struct {
	CantidadDigitos int    `json:"cantidad_de_digitos"`
	Digitos         string `json:"digitos"`
	SumaDigitos     int    `json:"suma_de_digitos"`
	DigitoMayor     int    `json:"digito_mayor"`
	DigitoMenor     int    `json:"digito_menor"`
	NumeroInvertido int    `json:"numero_invertido"`
	Capicua         bool   `json:"capicua"`
}

func showDigits(number string) string {
	var cont int
	var str string
	for cont < len(number)-1 {
		str += string(number[cont])
		str += ", "
		cont++
	}
	str += string(number[cont])
	return str
}

func digitsSum(number string) int {
	var sum int
	for _, e := range number {
		val, _ := strconv.Atoi(string(e))
		sum += val
	}
	return sum
}

func biggerAndMinorDigit(number string) (int, int) {
	var bigger int
	minor, _ := strconv.Atoi(string(number[0]))
	var cont int
	for cont < len(number) {
		n, _ := strconv.Atoi(string(number[cont]))
		if n > bigger {
			bigger = n
		}
		if n < minor {
			minor = n
		}
		cont++
	}
	return bigger, minor
}

func reverseNumber(s string) int {
	var str string
}
