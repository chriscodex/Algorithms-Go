package main

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
