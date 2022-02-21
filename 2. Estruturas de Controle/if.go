package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/* ---> Estrutura de código do If:

<Palavra chave > [Critério] { \\Ação }

Palavras-chaves: if, else e else if
*/

/* ---> Problema 1037 do Beecrowd, solucionado anteriormente com
adição da leitura utilizada pelo professor Ricardo: */

func main() {

	//Leitura de dados
	fmt.Printf("Digite um número: ")
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	input = strings.TrimSpace(input)
	num, _ := strconv.ParseFloat(input, 64)

	//Utilização do if:
	if num >= 0 && num <= 25 {
		fmt.Printf("Intervalo [0,25]\n")
	} else if num > 25 && num <= 50 {
		fmt.Printf("Intervalo (25,50]\n")
	} else if num > 50 && num <= 75 {
		fmt.Printf("Intervalo (50,75]\n")
	} else if num > 75 && num <= 100 {
		fmt.Printf("Intervalo (75,100]\n")
	} else {
		fmt.Printf("Fora de intervalo\n")
	}
}
