# Portuguese version

### Lingagem escolhida

#### Go

Motivos da escolha:

- Tipada fortemente:
 - - Strings
- -  bools
- - Intxxxxx
- - Floats
- Naturalmente a linguagem lida com concorrência.
- Leve.
- Fácil de aprender.


### Objetivo do projeto.

API de usuários onde esses usuários serão atualizados via webhook. Ou seja, um sistema de terceiro vai fazer trigger (disparar) numa URL do nosso sistema existente e atualizar esse usuário(s) para nós.

Regras básicas:
- Tem que lidar com concorrência.
- Validações simples:
- - existe o usuário.
- - foi deletado?
- - deu errado e precisamos descobrir o porque deu errado.

## Regras para desenvolvedores

- Orientado a testes
--  TDD, BDD, IT 
- Ter artillery -> https://www.artillery.io/

## Estrutura

Estrutura básica da nossa pasta de projetos:

- MVC


# Enlgish version

### Chosen Language

#### Go

Reasons for the choice:

- Strongly typed:

- Strings
- booles
- Intxxxxx
- Floats
- Naturally handles concurrency.

- Lightweight.

- Easy to learn.

### Project Objective.

User API where these users will be updated via webhook. That is, a third-party system will trigger a URL in our existing system and update this user(s) for us.

Basic rules:

- Must handle concurrency.

- Simple validations:

- The user exists.

- Was it deleted?

- Something went wrong and we need to find out why it went wrong.

## Rules for Developers

- Test-driven
- TDD, BDD, IT

- Have Artillery -> https://www.artillery.io/

## Structure

Basic structure of our project folder:

- MVC