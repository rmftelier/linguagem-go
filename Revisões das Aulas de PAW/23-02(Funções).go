// ---> Funções (Parte 1 da Aula)
package main 

import(
	"fmt"
	"log"
)

// Variável global 
var rendimentoTinta float64 

/* Informações Diversas sobre Funções -> 

1. Pode estar em qualquer parte do código. 

2. O tipo de retorno fica após os parâmetros. 

3. Caso não retorne nada (void) basta não declarar
um tipo de retorno após os parâmetros da função.

*/

func tintaNecessaria(largura, altura float64) (float64, error){
     //Tratamento de erro 
		 if largura <= 0 || altura <= 0 {
			 return 0, fmt.Errorf("A largura e altura devem ser positivos.")
		 }
		 area := largura * altura
		 return area/rendimentoTinta, nil
}

// Função principal 
func main(){
     rendimentoTinta = 12.1
		 //Passagem de valores pelo parâmetro
		 resultado, err := tintaNecessaria(4.2, 5.11)
     if err != nil{
			 log.Fatal(err)
		 }
     
		 fmt.Printf("%5.2f litros de tinta.\n", resultado)
}