package main

import(
	"fmt"
)

/* ---> Estrutura de Dados: Vetor

Exemplificação da estrutura de dados "Array" com o 
auxílio de um algoritmo que calcula a média dos 
números armazenados dentro de um array declarado.

*/

func main(){
    
	//Declaração de um vetor já inicializado:
	pesos := [5]float64{15.4, 14.3, 30.0, 22.1, 10}
	
	/* Declaração de uma variável que calculará a 
	   soma dos números para o cálculo da média que 
		 ocorrerá posteriormente:
	*/
	var soma float64

	//Utilizando foreach para percorrer o vetor
	//for {indice}, {valor}
	for _, p := range pesos{
      soma += p 
	}

	//Cálculo da média 
	media := soma / float64(len(pesos)) //len == tamanho do vetor pesos

	//Saída de dados 
	fmt.Printf("Média de pesos = %.2f \n", media)
}
