# Story API

Esta API permite a criação e gerenciamento de histórias interativas, com integração para geração de conteúdo usando a API do Gemini.

## Instalação

Clone o repositório:

```sh
git clone https://github.com/seu-usuario/story-api.git
```

Entre no diretório do projeto:

```sh
cd story-api
```

Instale as dependências:

```sh
go mod tidy
```

## Execução

Para iniciar o servidor, execute:

```sh
make run 
```

ou 

```sh
go run cmd/main.go
```

A API estará disponível em `http://localhost:8080`.

## Documentação

A documentação da API pode ser acessada em `http://localhost:8080/swagger/index.html`.

## Endpoints

### Autenticação

- `POST /register`: Registrar um novo usuário.
- `POST /login`: Fazer login e obter um token JWT.

### Histórias

- `GET /stories`: Listar todas as histórias.
- `POST /stories`: Adicionar uma nova história.
- `GET /stories/:id`: Obter uma história pelo ID.
- `PUT /stories/:id`: Atualizar uma história pelo ID.
- `DELETE /stories/:id`: Remover uma história pelo ID.
- `POST /stories/:id/generate`: Gerar conteúdo para uma história usando a API do Gemini.

### Usuários

- `GET /users`: Listar todos os usuários (requer autenticação).
- `GET /users/:id`: Obter um usuário pelo ID (requer autenticação).
- `PUT /users/:id`: Atualizar um usuário pelo ID (requer autenticação).
- `DELETE /users/:id`: Remover um usuário pelo ID (requer autenticação).

## Testes

Para executar os testes, use:

```sh
make test
```


