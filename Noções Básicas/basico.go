// ----> 1. Estrutura Básica do Código
package main //1.1 -> pacotes

import (
	"fmt"
	"math"
	"reflect"
) //1.2 -> importações

func main() { //1.3 -> função principal

	// ----> 2. Tipos Primitivos

	fmt.Println("Hello, World!") // 2.1 -> String
	fmt.Println('a')             // 2.2 -> Caracter/Runas

	/* ----> Runas
	* Runas são valores inteiros de 32 bits são constantes
	* não tipadas). Elas representam pontos de código unicode.
	* Por exemplo, o literal da runa 'a' é o número 97.
	*
	* rune = int32
	 */

	fmt.Println(true) // 2.3 -> Boolean

	/* ----> 3. Tipos Numéricos
	* Os tipos são divididos entre: Inteiros, Fracionários
	* e Complexos.
	*
	* Os tipos numéricos possuem várias subtipos que são
	* usados para mostrar a diferença entre a capacidade
	* de bytes que cada tipo suporta.
	 */

	// ----> 4. Operações Aritméticas e Comparações
	fmt.Println(math.Floor(3.14))
	fmt.Println(1 + 5)
	fmt.Println(6 / 5)
	fmt.Println(math.Mod(6, 2))
	fmt.Println(1 == 2)
	fmt.Println(6 > 2)
	fmt.Println(2 < 1)

	// ----> 5. Declaração de Variáveis

	// 5.1 -> Declaração explícita
	var qtd int //var + nome da variável + tipo
	qtd = 4
	fmt.Println(qtd)

	// 5.2 -> Declaração explícita com valor definido
	var nome string = "Roberta"
	fmt.Println(nome)

	// 5.3 -> Declaração explícita - Declaração múltipla
	var x, y, z int
	x, y, z = 1, 2, 3
	fmt.Println(x, y, z)

	// 5.4 -> Declaração explícita com valor definido - Declaração múltipla
	// var nomeVar1, nomeVar2 tipo = valor1, valor2

	//5.5 -> Declaração implícita com valor definido
	// var nomeVar = valor1

	//5.6 -> Declaração implícita com valor definido - Declaração múltipla
	// var nomeVar, nomeVar2 = valor, valor2

	//5.7 -> Declaração implícita reduzida
	qtnd := 5
	fmt.Println(qtnd)

	//5.8 -> Declaração implícita reduzida - Declaração múltipla
	//nomeVar, nomeVar2 := valor, valor2

	// 6 --> Conversões de Tipos
	myInt := 2
	var myFloat float64 = float64(myInt)

	fmt.Println(reflect.TypeOf(myInt), myInt)
	fmt.Println(reflect.TypeOf(myFloat), myFloat)

	/* 7. Quando há uma variável que não possuí atribuição de
	* valor, é atribuída o valor padrão do seu tipo.
	*
	* ex: Boolean = false
	 */

	/* ---> 8. Regras de Nomenclatura na Declaração de Variáveis:
	*
	* 1. A variável deve ser iniciada com letra minúscula, a palavra composta deve
	* ser maíuscula (padrão)
	*
	* 2. Variáveis começadas com letra minúscula é privada, caso contrário ela é
	* pública
	*
	 */

}
