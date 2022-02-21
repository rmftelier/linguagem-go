package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/* ---> Estrutura de código do Switch:

switch <variável a ser analisada>{

case <critério a ser analisado>:
	   //código caso ele seja aceito.

default:
     //código caso todos os 'cases' sejam recusados.
}

Palavras-chaves: switch, case e default
*/

/* ---> Problema 1038 do Beecrowd, solucionado anteriormente com
adição da leitura utilizada pelo professor Ricardo: */

func main() {

	var total float64

	//Leitura de dados:
	fmt.Print("Digite o código do seu lanche: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	input = strings.TrimSpace(input)
	codigo, _ := strconv.ParseInt(input, 10, 64)

	fmt.Print("Digite a quantidade que deseja: ")
	reader = bufio.NewReader(os.Stdin)
	input, err = reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	input = strings.TrimSpace(input)
	qtnd, _ := strconv.ParseFloat(input, 64)

	//Utilização do switch:
	switch codigo {
	case 1:
		total = 4.00 * qtnd
	case 2:
		total = 4.50 * qtnd
	case 3:
		total = 5.00 * qtnd
	case 4:
		total = 2.00 * qtnd
	case 5:
		total = 1.50 * qtnd
	default:
		total = 0.00
	}

	//Saída de dados
	fmt.Printf("Total: R$ %.2f\n", total)
}
