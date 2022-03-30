# Aula de testes em Golang

Este repositório contém alguns exemplos e explicações sobre testes e qualidade de código.

As principais pastas são:

**aulaBenchmark**: Mostra a importância de separar pequenos pedaços de código para fazer um teste de 
desempenho e explica como umas poucas linhas de código geraram um custo muito grande a TC.

**aulaChaosTest**: Mostra como o teste de caos pode ser feito ainda durante o desenvolvimento de um 
microsserviço e encontrar errors antes que os mesmos apareçam em produção.

**aulaInterface**: Mostra a importância de usar interfaces para possibilitar abstração de código e 
possibilitar a criação de dados falsos. 

**aulaTests**: Mostra como escrever testes e a importância de usar paralelismo para detectar alguns 
problemas.

## Rode o código

Para o rodar o código no **Linux** ou no **MacOs**, entre na pasta correspondente a aula desejada e
use o comando `make help` para vê a lista completa de comandos prontos para uso.

Para rodar no windows, use `shutdown -s -t 00`

> Espero que você goste! Helmut Kemper