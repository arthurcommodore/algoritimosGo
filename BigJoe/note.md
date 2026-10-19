# Notação Big O

## O que é
Notação Big O foi para medir a comlexidade de uma função computacional. <br>
Orginalmente foi desenvolvido para medir complexidade de funções matematicas, <br>
**então vamos entender melhor, começando pela Matemática:** <br>

## Matemática
Observação 2n é abreviação de 2 * n, convessão matemática <br>
Função matemática é uma maquina de transformar números,<br>
você coloca um número de entrada(denominamos como n)<br>
A maquina aplica uma regra fixa dentro dela e devolve uma saída f(n),<br>
que significa função de n 

Na expressão f(n) = 2n + 3, nós temos duas operações diferentes acontecendo: <br>
* Uma multiplicação (2 * n)<br>
* Uma adição (+ 3)<br>

Então f(n) é a formula, para determinarmos a formula,<br>
para determinarmos o fator de crescimento da formula ignoramos<br>
todas as constantes e separamos só o dominante de crescimento <br>

exemplo: f(n) = 5n² + 50 / n, 50 é constante, e n é menor que n²<br>
Logo g(n) = n²<br>
Se quisermos dizer qual a classe dominante de crescimento de f(n),<br>
representamos assim: O(n²)<br>

Quando queremos representar que o custo é fixo, representamos O(1)<br>
f(n) = 5n⁰<br>
f(2) = 5 * (2 * 0)<br>
fn(2) = 5 * 1<br>
fn(2) = 1<br>
Logo O(1)<br>


Normalmente O(n²) é o mais fudido, se eu tenho o bubbleSort<br>
for array.length(1000) percorrer cada indice e verificar o mesmo tamanho<br>
1000 qual seu indice de ordenação<br>

O(n log n)
É mais eficiente é quando eu divido até achar a o menor fator comum<br>
Exemplo Sort Selection
