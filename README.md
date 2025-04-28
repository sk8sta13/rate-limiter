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

   ```bash
   curl http://localhost:8080 -H "API_KEY: ffa1754d27f18f1a95b5f9a76adfb62b64982691"
   ```  
O sistema verificará a quantidade de acessos e aplicará as regras de limitação conforme configurado.

Outro ponto é que ao editar uma configuração não é necessário restartar o programa, basta salvar o .env e o programa passara a considerar as novas configurações.

## Detalhes do desenvolvimento

### Como é aplicado as regras para o ip e ou token?

Ao receber uma request verificamos se o dado "API_KEY" está presente no head, se sim então seguindo a regra de precedencia o middleware de token é aplicado, caso esse dado não sejá enviado no header da request então aplicamos o middleware de ip.

Os dados salvos no redis tem a seuginte estrutura:

```bash
{"FirstMoment": 1745773976, "LastMoment": 1745773980, "Qtd": 10}
```

- FirstMoment: Esse dado representa em timestamp a primeira request realizada;
- LastMoment: Esse dado representa em timestamp a última request realizada;
- Qtd: Representa o número de requests realizadas;

Tanto para o middleware de IP quanto o Token seguem essa estrutura o que diferencia eles dentro do redis é o key, para o IP o key é simplismente o próprio IP, e para o Token o key do redis é o Token informado no API_KEY com o IP.

Com esses dados é possivel aplicar as regras definidas, veja no video abaixo:

https://github.com/user-attachments/assets/c597b8f2-4ddf-4683-9a9d-b71d40610320

## Teste e Estresse

Utilizando o programa ab, eu executei 1000 request, sendo 100 paralelas e mandei escrever um arquivo teste com a saida do programa:

![image](https://github.com/user-attachments/assets/70240ef2-73bc-4a59-83d3-8b0e20866950)

Verifique que no arquivo teste tenho 765 request foram diferentes de status code 200, e sabemdo que o meu programa retorna apenas três status code possíveis 200, 401 e 429.  
Já removendo a possibilidade de erro no API_KEY informado pois eu setei um válido, sendo assim sabemos que nenhuma das requests retornaram 401 unauthorized.  
Eu fiz uma busca por uma string "HTTP/1.0 429" no arquivo que tem os status code de todas as requests realizadas e temos agora certeza de que o programa está funcionando corretamente pois o número de ocorrências encontrado bate exatamente com o número de requests diferentes de 200 do relatório do programa utilizado para o teste "ab", podemos ver isso nos dados destacados em um retangulo vermelho no print.

## Observações

Se o token exceder o número de requisições permitido, o IP ou token será bloqueado por um tempo configurável.

O código é modularizado para facilitar substituições (ex: usar outro banco de dados que não o Redis).
