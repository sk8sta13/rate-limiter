# rate-limiter
Atividade Pós-GoExpert Rate Limiter

## Descrição

Este projeto é um **Rate Limiter** escrito em **Go**, que permite limitar o número de requisições feitas a uma API baseado em configurações dinâmicas via **arquivo `.env`**.

A persistência dos dados é feita utilizando **Redis**, permitindo o controle de tentativas por IP ou Token.

## Funcionalidades

- **Limitação de Requisições** por token de API.
- **Configuração Dinâmica** através de variáveis de ambiente.
- **Bloqueio Temporário** após excesso de requisições.
- **Persistência de Dados** usando Redis.
- **Extensível** via camada de repositórios para facilitar futuras integrações com outros bancos.

## Como Configurar

1. Clone o repositório e copie o arquivo de exemplo .env-example para .env:

   ```bash
   git clone https://github.com/sk8sta13/rate-limiter.git
   cd rate-limiter
   cp .env-example .env
   ```

2. Edite o .env com seus parâmetros:

   ```bash
   DB_HOST=172.27.0.2
   DB_PORT=6379
   DB_PASS=

   IP_MAX_REQUESTS=10
   IP_MAX_REQUESTS_IN_SECONDS=20
   IP_BLOCKED_FOR_SECONDS=20

   TONEK_P=ffa1754d27f18f1a95b5f9a76adfb62b64982691
   TONEK_P_MAX_REQUESTS=20
   TOKEN_P_MAX_REQUESTS_IN_SECONDS=40
   TOKEN_P_BLOCKED_FOR_SECONDS=40

   TONEK_M=34a4a72e6e1d39e762951871166404cf64778bba
   TONEK_M_MAX_REQUESTS=30
   TOKEN_M_MAX_REQUESTS_IN_SECONDS=60
   TOKEN_M_BLOCKED_FOR_SECONDS=60

   TONEK_G=1a5c575d0d7c3978a5c77e7bab984ce12e433cc5
   TONEK_G_MAX_REQUESTS=40
   TOKEN_G_MAX_REQUESTS_IN_SECONDS=80
   TOKEN_G_BLOCKED_FOR_SECONDS=80
   ```  
A idéia é que tenhamos três tokens com quantidades progressivas de acesso.
Outro ponto é que a aplicação não precisa ser reiniciada para que a alteração de configuração passe a valer.

3. Suba os containers:

   ```bash
   docker-compose up -d
   ```

## Como Testar

Faça as requests:

   ```bash
   curl http://localhost:8080
   ```  

    OU

   ```bash
   curl http://localhost:8080 -H "API_KEY: ffa1754d27f18f1a95b5f9a76adfb62b64982691"
   ```  
O sistema verificará a quantidade de acessos e aplicará as regras de limitação conforme configurado.

## Observações

Se o token exceder o número de requisições permitido, o IP ou token será bloqueado por um tempo configurável.

O código é modularizado para facilitar substituições (ex: usar outro banco de dados que não o Redis).