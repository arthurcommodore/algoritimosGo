# Progressão Aritmética (PA) e sua relação com algoritmos

## 1. O que é uma Progressão Aritmética?

Uma **Progressão Aritmética (PA)** é uma sequência de números em que a
diferença entre um termo e o próximo é sempre a mesma.

Essa diferença constante é chamada de **razão**, representada pela letra
`r`.

Exemplos:

-   `1, 2, 3, 4, 5` tem razão `r = 1`.
-   `2, 4, 6, 8, 10` tem razão `r = 2`.
-   `10, 8, 6, 4, 2` tem razão `r = -2`.

A sequência pode ser crescente ou decrescente. O importante é que a
diferença entre termos consecutivos seja constante.

Uma sequência como `1, 3, 4, 8, 10` não é uma PA, porque suas diferenças
(`2, 1, 4, 2`) não são constantes.

## 2. Como somar os termos de uma PA?

Considere:

\[ S = 1 + 2 + 3 + 4 + 5 \]

Podemos escrever a mesma soma na ordem inversa:

\[ S = 5 + 4 + 3 + 2 + 1 \]

Somando as duas linhas, coluna por coluna:

    Primeira linha   Segunda linha   Soma da coluna
  ---------------- --------------- ----------------
                 1               5                6
                 2               4                6
                 3               3                6
                 4               2                6
                 5               1                6

Todas as colunas têm o mesmo resultado. Como existem cinco colunas:

\[ S + S = 6 + 6 + 6 + 6 + 6 \]

Como `S + S = 2S`:

\[ 2S = 5 `\times 6`{=tex} = 30 \]

Mas 30 representa **duas cópias da soma original**. Queremos apenas uma
cópia, então dividimos por 2:

\[ S = `\frac{30}{2}`{=tex} = 15 \]

A divisão por 2 desfaz o fato de termos contado a soma duas vezes.

## 3. Fórmula da soma de uma PA

O raciocínio anterior funciona para qualquer PA finita. A fórmula é:

\[ `\boxed{S_n = \frac{n(a_1 + a_n)}{2}}`{=tex} \]

Onde:

-   `S_n`: soma de todos os termos;
-   `n`: quantidade de termos;
-   `a_1`: primeiro termo;
-   `a_n`: último termo.

Em palavras:

**soma = quantidade de termos × (primeiro termo + último termo) ÷ 2**

### Exemplo

Some os números de 1 até 7:

-   Primeiro termo: `1`
-   Último termo: `7`
-   Quantidade de termos: `7`

\[ S_7 = `\frac{7(1 + 7)}{2}`{=tex} \]

\[ S_7 = `\frac{7 \times 8}{2}`{=tex} = 28 \]

Conferindo:

\[ 1 + 2 + 3 + 4 + 5 + 6 + 7 = 28 \]

## 4. Os números precisam estar ordenados?

Não precisam estar originalmente em ordem crescente. A fórmula também
funciona para uma PA decrescente.

O requisito é que os termos formem uma progressão com diferença
constante. A fórmula usa a quantidade de termos, o primeiro e o último
termo da PA.

Por exemplo:

\[ 10 + 8 + 6 + 4 + 2 = 30 \]

É uma PA decrescente, com razão `-2`:

\[ S_5 = `\frac{5(10 + 2)}{2}`{=tex} = 30 \]

Não podemos aplicar diretamente a fórmula a qualquer conjunto de
números. Por exemplo, `1, 3, 4, 8, 10` não forma uma PA.

## 5. Exemplo em Go: somando uma PA

A função abaixo calcula a soma de uma PA finita usando a fórmula
matemática:

``` go
package main

import "fmt"

func SomaPA(quantidade, primeiro, ultimo int) int {
    return quantidade * (primeiro + ultimo) / 2
}

func main() {
    resultado := SomaPA(7, 1, 7)

    fmt.Println(resultado) // 28
}
```

A função recebe:

-   `quantidade`: número de termos;
-   `primeiro`: primeiro termo;
-   `ultimo`: último termo.

**Atenção:** esse exemplo usa `int` e é adequado quando o resultado
intermediário e final cabem nesse tipo. Em uma implementação genérica,
também seria necessário validar se os valores realmente formam uma PA e
considerar possíveis limites numéricos.

## 6. Relação com o Selection Sort

O Selection Sort organiza um vetor procurando o menor elemento da parte
ainda não ordenada e colocando-o na posição correta.

Considere um vetor com cinco elementos:

``` go
nums := []int{8, 3, 5, 1, 9}
```

Na primeira passagem, o algoritmo procura o menor entre os cinco
elementos. Para isso, faz quatro comparações.

Depois, o primeiro elemento já está na posição correta. Na segunda
passagem, procura o menor entre os quatro restantes: faz três
comparações.

As passagens seguintes fazem:

    Passagem   Elementos ainda não ordenados   Comparações
  ---------- ------------------------------- -------------
           1                               5             4
           2                               4             3
           3                               3             2
           4                               2             1

Total:

\[ C(5) = 4 + 3 + 2 + 1 = 10 \]

Essa sequência de comparações é uma PA decrescente, com primeiro termo
`4`, último termo `1` e quantidade de termos `4`.

Aplicando a fórmula:

\[ C(5) = `\frac{4(4 + 1)}{2}`{=tex} = 10 \]

### Fórmula geral de comparações

Para um vetor com `n` elementos, o Selection Sort faz:

\[ (n-1) + (n-2) + `\cdots `{=tex}+ 2 + 1 \]

Essa é uma PA com:

-   primeiro termo: `n - 1`;
-   último termo: `1`;
-   quantidade de termos: `n - 1`.

Aplicando a fórmula da soma:

\[ C(n) = `\frac{(n-1)((n-1)+1)}{2}`{=tex} \]

Simplificando:

\[ `\boxed{C(n) = \frac{n(n-1)}{2}}`{=tex} \]

Essa expressão conta exatamente as comparações entre elementos feitas
pela implementação tradicional do Selection Sort.

### Alguns tamanhos de vetor

    Elementos (`n`)   Comparações
  ----------------- -------------
                  5            10
                 10            45
                100         4.950
              1.000       499.500
             10.000    49.995.000

O crescimento fica rápido: quando dobramos aproximadamente o tamanho do
vetor, o número de comparações fica próximo de quatro vezes maior.

## 7. Implementação do Selection Sort em Go

``` go
package algoritimo

func SelectionSort(nums []int) []int {
    for i := 0; i < len(nums)-1; i++ {
        sort(nums, i)
    }

    return nums
}

func sort(nums []int, indice int) {
    var menorIndice int = indice

    for i := indice + 1; i < len(nums); i++ {
        if nums[menorIndice] > nums[i] {
            menorIndice = i
        }
    }

    nums[indice], nums[menorIndice] = nums[menorIndice], nums[indice]
}
```

A função `sort` procura o índice do menor elemento na parte não ordenada
e depois troca esse elemento com o elemento da posição atual.

A cada passagem, a região não ordenada diminui em um elemento. Por isso,
a quantidade de comparações segue a sequência:

``` text
n-1, n-2, n-3, ..., 2, 1
```

## 8. E o que isso tem a ver com Big O?

A expressão exata para o número de comparações é:

\[ C(n) = `\frac{n(n-1)}{2}`{=tex} \]

Abrindo a multiplicação:

\[ C(n) = `\frac{n^2-n}{2}`{=tex} \]

Separando os termos:

\[ C(n) = `\frac{n^2}{2}`{=tex} - `\frac{n}{2}`{=tex} \]

Na análise Big O, observamos como o crescimento se comporta quando `n`
aumenta muito. Por isso, ignoramos constantes multiplicativas e termos
de ordem menor.

O termo que domina o crescimento é `n²`. Portanto:

\[ `\boxed{C(n) \in O(n^2)}`{=tex} \]

Isso não significa que a fórmula exata e o Big O sejam a mesma coisa:

-   **Fórmula exata:** informa quantas comparações são realizadas.
-   **Big O:** descreve a classe de crescimento do trabalho conforme a
    entrada aumenta.

No Selection Sort tradicional, a quantidade de comparações é quadrática
mesmo quando o vetor já está ordenado, porque o algoritmo continua
procurando o menor elemento em cada parte não ordenada.

## 9. O que guardar deste estudo?

1.  Uma PA é uma sequência cuja diferença entre termos consecutivos é
    constante.
2.  A soma de uma PA pode ser calculada multiplicando a quantidade de
    termos pela soma do primeiro e do último, e dividindo por 2.
3.  Dividimos por 2 porque, ao somar a sequência original com a
    invertida, contamos a soma duas vezes.
4.  O Selection Sort faz uma quantidade decrescente de comparações a
    cada passagem.
5.  Essa quantidade forma uma PA: `n-1, n-2, ..., 1`.
6.  A soma dessa PA resulta em `n(n-1)/2`, que cresce como `O(n²)`.

A ideia principal é aprender a construir a fórmula a partir do
comportamento do algoritmo, em vez de apenas decorar o resultado.
