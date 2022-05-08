package main 

import(
	"fmt"
)

/* ---> Estrutura de Dados: Slice

Exemplificação da estrutura de dados "Slice" com o 
auxílio de um algoritmo que recebe as notas do 
usuário até que ele digite um número negativo e 
após isso mostre o slice criado e a média calculada. 

*/

func main(){
    //Declaração de um slice vazio
		notas := make([]float64, 0)

		/* Declaração de uma variável que será lida pelo
    usuário. */
		var nota float64

		//Leitura de dados
		for nota >= 0{
			  fmt.Printf("Digite uma nota: ")
				fmt.Scanf("%f",  &nota)

				/* Adicionando a nota digitada pelo usuário 
			  no slice já declarado */
				notas = append(notas, nota)
		}

    //Para obtermos uma faixa específica de elementos
		//Declaração: slice[de : até]

		/* Temos  que diminuir um do 
		tamanho, pois o último valor é o valor negativo que serve como ponto de 
		parada e não deve ser utilizado no cálculo da média. */
		subSlice := notas[0 : len(notas) - 1] 
    
		var soma float64
		for _, p := range notas{
			  soma += p 
		}
		
		//Saída de dados 
		fmt.Printf("%f \n", subSlice)
		fmt.Printf("Media: %.2f \n", soma/float64(len(notas)))
}