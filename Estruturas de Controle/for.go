package main

import (
	"fmt"
)

/* ---> Estrutura de código do For:

1. Tradicional
for < inicializador >; <controlador>; <incremento> { }

2. Repetição com inicializador e incremento desassociado (While):

<inicializador> := 1

for <controlador> <= 10 {
	  //ação
		<incremento>
}
*/

/* ---> Problema 1065 do Beecrowd, solucionado anteriormente
com a adição da indicação da quantidade de valores ímpares lidos: */

func main() {

	//Declaração de variáveis
	var num int64
	pares := 0   //Quantidade de números pares
	impares := 0 //Quantidade de números ímpares

	fmt.Printf("Digite 5 números: ")
	/* 1ª forma de resolução: For tradicional */
	for contador := 1; contador <= 5; contador++ {
		fmt.Scanf("%d", &num)

		if num%2 == 0 {
			pares++
		}
	}

	/* 2ª forma de resolução: While
	   - Agora com os números ímpares
	*/
	contador := 1
	for contador <= 5 {
		fmt.Scanf("%d", &num)

		if num%2 != 0 {
			impares++
		}
		contador++
	}

	//Saída de dados
	fmt.Printf("%d valores pares\n", pares)
	fmt.Printf("%d valores ímpares\n", impares)
}
