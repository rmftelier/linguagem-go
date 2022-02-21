package main

/* Os nomes no import não são os pacotes propriamente ditos e
sim os caminhos para eles.
Exemplo: math/rand é o caminho do pacote math onde está localizado
rand */
import (
  "fmt"
	"math/rand"
	"time"
	"bufio"
	"os"
	"log"
	"strings"
	"strconv"
)
/* Programa de um jogo onde o usuário terá que acertar um 
número em até 10 tentativas 

Conteúdo: Random, Leitura de Dados, Estruturas de Controle */
func main(){
    
	   /* Gerando um número aleatório para o jogo */

		 seed := time.Now().Unix()
		 rand.Seed(seed)

		 numero := rand.Intn(100) + 1 
		 fmt.Println("Pensei em um número, de 1 a 100.")
		 fmt.Println("Consegue advinhar?")

		 // Leitura de dados 
		 reader := bufio.NewReader(os.Stdin)
		 sucesso := false

		 // Processamento de dados
		 for i := 0 ; i < 10 ; i++{
					fmt.Println("Você tem:", 10 - i, "chances sobrando")
					fmt.Println("Qual o seu chute?")

					input, err := reader.ReadString('\n')
					//Tratamento de erro
					if err != nil{
						log.Fatal(err)
					}

					//Conversão da string lida para int 
					input = strings.TrimSpace(input) //Tiramos os espaços em branco da string
          chute, err := strconv.Atoi(input)		 
			    
					if err != nil {
						log.Fatal(err)
					}

					if chute == numero {
						 fmt.Println("Você acertou!")
						 sucesso = true 
						 break //O break para a execução do loop atual. 
					} else if chute < numero {
						   fmt.Println("Opa, abaixo do que pensei.")
							 continue /* O continue é usado para pular o restante do loop
               e retornar ao topo do loop */
					} else {
						    fmt.Println("Eita, essa foi acima do que pensei")
					}
	  	}

			if !sucesso {
				  fmt.Println("Você perdeu. O número era:", numero)
			}
}