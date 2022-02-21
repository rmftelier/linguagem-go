# 📁 Gerando números aleatórios e strings em movimento
https://flaviocopes.com/go-random/

- Original em inglês: [Flavio Copes](https://flaviocopes.com/go-random/)
- Tradução pt-br por: @rmftelier

> Tradução até a parte da geração de um número inteiro randômico

## Números Pseudoaleatórios 
<p> O pacote **math/rand** fornecido pela Go Standard Library nos fornece geradores de números pseudo-aleatórios (PRNG), também chamados de geradores de bits aleatórios determinísticos. </p>

<p> Como em todos os pseudo geradores de números, qualquer número gerado por math/rand não é realmente aleatório por padrão, pois sendo **determinístico**, sempre imprimirá o mesmo valor a cada vez. </p>

<p> Como exemplo, tente executar este código que introduz rand.Intn(n), que retorna um número aleatório entre 0 e n-1. \n

[playground](https://go.dev/play/p/-HsFj0sqWD)
 </p>


 <p>  Você sempre verá a mesma sequência toda vez que executar o programa. O número aleatório muda dentro do programa, mas toda vez que você o executa, você obtém a mesma saída: </p>

<img src="" >

<p> Isso ocorre porque, por padrão, a semente (seed) é sempre a mesma, o número 1. Para realmente obter um número aleatório, você precisa fornecer uma semente exclusiva para seu programa. Faremos isso da seguinte maneira: </p>

<p> Use rand.Seed() antes de chamar qualquer método math/rand, passando um valor int64. Você só precisa semear uma vez em seu programa, não sempre que precisar de um número aleatório. A semente mais usada é a hora atual, convertida para int64 pelo UnixNano com rand.Seed(time.Now().UnixNano()): </p>


## Gerando um número inteiro 

