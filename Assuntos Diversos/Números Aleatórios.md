#  📁 Gerando números aleatórios

- Original (eng): [Flavio Copes](https://flaviocopes.com/go-random/)
- Tradução (pt-br): Google Tradutor 

> Tradução até a parte da geração de um número inteiro randômico

## 🔢 Números Pseudoaleatórios 
O pacote **math/rand** fornecido pela Go Standard Library nos fornece [geradores de números pseudo-aleatórios (PRNG)](https://en.wikipedia.org/wiki/Pseudorandom_number_generator), também chamados de geradores de bits aleatórios determinísticos. 

Como em todos os pseudo geradores de números, qualquer número gerado por math/rand não é realmente aleatório por padrão, pois sendo **determinístico**, sempre imprimirá o mesmo valor a cada vez.

Como exemplo, tente executar este código que introduz **rand.Intn(n)**, que retorna um número aleatório entre 0 e n-1. 

<img src="https://user-images.githubusercontent.com/63109114/155031120-3cfd5246-62d3-4ecc-94e7-4d8ef685b6a2.png" align="center" width="400">

[playground](https://go.dev/play/p/-HsFj0sqWD)

Você sempre verá a mesma sequência toda vez que executar o programa. O número aleatório muda dentro do programa, mas toda vez que você o executa, você obtém a mesma saída:

<img src="https://user-images.githubusercontent.com/63109114/155031297-4586dca1-18ae-4562-9c4a-096f1729d110.png" align="center" width="400">

Isso ocorre porque, por padrão, a semente (seed) é sempre a mesma, o número 1. Para realmente obter um número aleatório, você precisa fornecer uma semente exclusiva para seu programa. Faremos isso da seguinte maneira:

Use **rand.Seed()** antes de chamar qualquer método math/rand, passando um valor int64. Você só precisa semear uma vez em seu programa, não sempre que precisar de um número aleatório. A semente mais usada é a hora atual, convertida para int64 pelo UnixNano com **rand.Seed(time.Now().UnixNano())**:

<img src="https://user-images.githubusercontent.com/63109114/155031412-29d62614-d799-4d43-a926-ba4853c49ed1.png" align="center" width="400">


## 🔁 Gerando um número inteiro 

<img src="https://user-images.githubusercontent.com/63109114/155031557-cace8220-68a2-48ca-b701-481a1d91bfed.png" align="center" width="400">

