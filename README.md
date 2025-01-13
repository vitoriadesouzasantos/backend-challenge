# Projeto

Este projeto foi desenvolvido para realizar a validação de senhas com os requisitos exigidos.
A liguagem escolhida foi Golang.

## Critérios usados para validação

A implementação foi realizada utilizando expressóes regulares para que tenha maior facilidade
de manutenção e entendimento.
Para a verificação da caracteres duplicados foi utilizado um map para armazenar os caracteres já
encontrados durante a iteração da senha, pois em Go não é possível utilizar backreferences ("\1") nas
expressões regulares.

## Como executar?

1. Identifique se o Go está instalado em sua máquina
2. Clone o repositório usando: 
 ````
 git clone https://github.com/vitoriadesouzasantos/backend-challenge
 ````
3. Acesse a pasta do projeto
4. Execute o comando: 
```
go run cmd/api/main.go
```

