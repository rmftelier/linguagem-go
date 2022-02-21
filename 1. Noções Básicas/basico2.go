package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
) 

func main() {
 
	/* -----> 1. Manipulando o tempo: */

				/* Utilizamos o pacote time para
				manipular o tempo na linguagem Go

				Obs: Time é o tipo dentro do pacote time
				*/

				/* Essa declaração de variável é a mais
				utilizada: */
				now := time.Now()

				// Mostraremos o ano atual da seguinte maneira:
				ano := now.Year()
				fmt.Println("Ano atual: ", ano)

	/* -----> 2. Pequeno programa que: lê uma nota do
	aluno e analisa sua situação. */

				//Leitura de dados 
				fmt.Println("Entre com uma nota: ")

				/* O pacote bufio implementa E/S (I/O) bufferizados.
				A E/S com buffer tem um desempenho muito melhor do que
				sem buffer. 
				Esse pacote então envolve um objeto io.Reader, criando 
				outro objeto que também implementa a interface, mas 
				fornece buffer e alguma ajuda para E/S textual.
				*/

				//O programa lê um nome da entrada usando bufio.NewReader: 
				reader := bufio.NewReader(os.Stdin)

				/* Passaramos a entrada padrão para o bufio.NewReader ou
				seja leremos o que for digitado até que o Enter (\n) 
				seja acionado */
				input, err := reader.ReadString('\n')

				/* input, _ := reader.ReadString('\n')  -> o ( _ ) 
				ignora o erro, não é recomendável utilizar dessa maneira */

				/* Go não possuí exceções mas podemos por agora mostrar o 
        erro para o usuário da seguinte maneira: */
				if err != nil{
					//nil == nulo 
					 log.Fatal(err)
				}

				/* Passo a passo da conversão da string que foi lida e 
				vinculada a variável reader para o tipo float: */

				/* 1º: Removemos espaços em brancos à esquerda e à 
				direita da string lida */
				input = strings.TrimSpace(input)

				// 2º: A conversão é feita
				nota, _ := strconv.ParseFloat(input, 64)

				// Declaração de uma variável
				var status string 

				//Processamento de dados 
			   if nota >= 7 {
						   status = "Aprovado"
					 } else if nota < 3 {
						   status = "Reprovado"
					 } else {
						   status = "De final"
					}
			  

				 //Saída de dados:
				 fmt.Println("Sua nota foi:", nota, ", Situação:", status)			 
}
