## Desafio Client Server API 

Criar dois sistemas em Go:
- client.go
- server.go

Os requisitos para cumprir este desafio são:
 
- client.go
  - deverá realizar uma requisição HTTP no server.go solicitando a cotação do dólar.
  - precisará receber do server.go apenas o valor atual do câmbio (campo "bid" do JSON)
  - Utilizando o package "context" terá um timeout máximo de 300ms para receber o resultado do server.go
  - terá que salvar a cotação atual em um arquivo "cotacao.txt" no formato: Dólar: {valor}

- server.go 
  - deverá consumir a API contendo o câmbio de Dólar e Real no endereço: https://economia.awesomeapi.com.br/json/last/USD-BRL e em seguida deverá retornar no formato JSON o resultado para o cliente.
  - Usando o package "context", deverá:
    - registrar no banco de dados SQLite cada cotação recebida, com timeout de:
      - 200ms para chamar a API de cotação do dólar
      - 10ms para conseguir persistir os dados no banco
  - endpoint necessário gerado pelo server.go para este desafio será: /cotacao e a porta a ser utilizada pelo servidor HTTP será a 8080.
 
