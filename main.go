package main

import (
	"fmt"
	"math"
)

type precoAcao struct {
	acao, data string
	precoret   float64
}

type rendimentoAcao struct {
	dataLucro         string
	precoAtual, lucro float64
}

func main() {

	//COMPRA DA AÇÃO
	compraAcao := precoAcao{"AAAA4", "05/04/2020", 22.99}
	fmt.Println("Ação:", compraAcao.acao, "\nData da compra:", compraAcao.data, "\nPreço da Ação: R$", compraAcao.precoret)

	//CONTAGEM LUCRO PRIMEIRA DATA
	contLucro1 := rendimentoAcao{"05/10/2020", 24.19, 0}
	contLucro1.lucro = contLucro1.precoAtual - compraAcao.precoret
	fmt.Println("Lucro da data", contLucro1.dataLucro, "é de: R$", math.Round(contLucro1.lucro*100)/100)

	//CONTAGEM LUCRO SEGUNDA DATA
	contLucro2 := rendimentoAcao{"05/04/2021", 25.74, 0}
	contLucro2.lucro = contLucro2.precoAtual - contLucro1.lucro - compraAcao.precoret
	fmt.Println("Lucro da data", contLucro2.dataLucro, "é de: R$", math.Round(contLucro2.lucro*100)/100)

	//CONTADEM DO LUCRO ANUAL
	retornoAnual := contLucro2.lucro + contLucro1.lucro
	fmt.Println("Lucro anual é de: R$", retornoAnual)

}
