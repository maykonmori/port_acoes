package main

import (
	"fmt"
	"math"
)

type precoAcao struct {
	acao     string
	data     string
	precoret float64
}

type rendimentoAcao struct {
	dataLucro  string
	rendimento float64
}

func main() {
	var a precoAcao
	a.acao = "AAAA0"
	a.data = "14/04/2020"
	a.precoret = 23.49

	fmt.Println(a.acao, a.data, a.precoret)

	var r1 rendimentoAcao
	r1.dataLucro = "05/10/2020"
	r1.rendimento = 24.79

	lucro1 := r1.rendimento - a.precoret
	fmt.Println(math.Round(lucro1*100) / 100)

	var r2 rendimentoAcao
	r2.dataLucro = "05/04/2021"
	r2.rendimento = 25.74

	lucro2 := r2.rendimento - a.precoret
	fmt.Println(math.Round(lucro2*100) / 100)

	retornoAnual := lucro2 + lucro1
	fmt.Println(math.Round(retornoAnual*100) / 100)

}
