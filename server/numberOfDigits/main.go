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

}
