// ---> Ponteiros (Parte 2 da Aula)
package main 

import(
	"fmt"
)

/* Uso de Paramêtros por Referência */

/* Explicação do programa:

&qtd possuí o endereço da variável, ou seja, 
passamos o endereço e consequentemente o valor
localizado nesse endereço para o ponteiro num 
(que agora apontará para o mesmo local na 
memória que a variável qtd) declarado como 
parâmetro da função double, dessa maneira ele 
é dobrado e quando chamamos a variável qtd no
main o seu valor foi mudado também, pois as 
duas variáveis apontam para o mesmo local de
memória. 

*/

func main(){
	qtd := 6
	//Chamando a função no main
	double(&qtd)
	fmt.Printf("%d \n", qtd)
}

/* Passagem de parâmetro por referência */
func double(num *int){
	*num *= 2 //*num = num * 2
	//Output: 12 
}
