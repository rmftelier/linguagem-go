package main

import (
	"fmt"
	"time"
)

func main() {
	/* Utilizamos o pacote time para
	manipular o tempo na linguagem Go
	*/

	//Time é o tipo dentro do pacote time
	//var now time.Time = time.Now()

	/* Essa declaração de variável é a mais
	utilizada
	*/
	now := time.Now()

	//Mostraremos o ano atual da seguinte maneira:
	ano := now.Year()
	fmt.Println("Ano atual: ", ano)
	fmt.Println("Ano atual: ", ano)

}
