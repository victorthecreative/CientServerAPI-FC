
# CientServerAPI-FC

## Visão Geral

CientServerAPI-FC é um projeto em Go que consiste em uma API para consultar a cotação do dólar (USD-BRL) e salvar essa cotação em um banco de dados MySQL. Além disso, o projeto inclui um cliente que faz uma requisição para essa API e grava a cotação em um arquivo de texto.

## Funcionalidades

- **API de Cotação do Dólar**: Exposição de uma rota HTTP (`/cotacao`) que retorna a cotação atual do dólar em relação ao real.
- **Salvamento em Banco de Dados**: Armazenamento da cotação obtida no banco de dados MySQL.
- **Cliente HTTP**: Um cliente que consome a API de cotação e salva o resultado em um arquivo de texto (`cotacao.txt`).

## Estrutura do Projeto

- **/internal/models**: Contém as definições das estruturas de dados usadas para a cotação.
- **/internal/database**: Contém as funções para manipulação do banco de dados, incluindo a criação de tabelas e inserção de dados.
- **/internal/server**: Contém o código do servidor que expõe a API e realiza a lógica de salvar a cotação no banco.
- **/internal/main.go**: Contém o código que inicia um  novo servidor.
- **/internal/Client**: Contem o código do client que consome essa api e escreve em um arquivo txt a cotação do dolar, e temos o arquivo main.go que executa o client

## Pré-requisitos

- Go 1.18 ou superior
- Docker e Docker Compose

## Instalação

1. **Clone o repositório**:
   ```bash
   git clone https://github.com/victorthecreative/CientServerAPI-FC.git
   cd CientServerAPI-FC
   ```

2. **Configuração do Banco de Dados usando Docker**:
    - Execute o comando abaixo para iniciar o contêiner MySQL usando Docker Compose:
   ```bash
   docker-compose up -d
   ```
    - O banco de dados MySQL estará disponível na porta `3306` com as seguintes credenciais:
        - **Usuário:** root
        - **Senha:** root
        - **Banco de dados:** RateExchange

3. **Compilação do Projeto**:
   ```bash
   go build -o cientserverapi-fc
   ```

## Uso

1. **Executar o Servidor**:
    - Inicie o servidor que expõe a API na porta 8080.
   ```bash
   ./cientserverapi-fc server
   ```

2. **Executar o Cliente**:
    - O cliente realiza uma requisição para o servidor, obtém a cotação do dólar e grava o valor em um arquivo `cotacao.txt`.
   ```bash
   ./cientserverapi-fc client
   ```

3. **Verifique o Arquivo de Saída**:
    - O arquivo `cotacao.txt` será gerado na raiz do projeto com o valor da cotação atual do dólar.

## Estrutura de Dados

### JsonInternalResponse

```json
{
  "code": "USD",
  "codein": "BRL",
  "name": "Dólar Americano/Real Brasileiro",
  "high": "5.50",
  "low": "5.40",
  "varBid": "0.10",
  "pctChange": "1.82",
  "bid": "5.45",
  "ask": "5.50",
  "timestamp": "1638475600",
  "create_date": "2024-08-29"
}
```

### JsonExternalResponse

```json
{
  "USDBRL": {
    "code": "USD",
    "codein": "BRL",
    "name": "Dólar Americano/Real Brasileiro",
    "high": "5.50",
    "low": "5.40",
    "varBid": "0.10",
    "pctChange": "1.82",
    "bid": "5.45",
    "ask": "5.50",
    "timestamp": "1638475600",
    "create_date": "2024-08-29"
  }
}
```

## Contribuição

Contribuições são bem-vindas! Sinta-se à vontade para abrir issues e pull requests para sugerir melhorias ou corrigir problemas.

## Licença

Este projeto está licenciado sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.

## Contato

- Autor: Victor Coelho
- Email: victor_developer@live.com
- GitHub: [https://github.com/victorthecreative](https://github.com/victorthecreative)

