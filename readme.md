# Go Project - Multithreading API

## Projeto é um desafio do curso <b>Go Expert</b> da Full Cycle.

A aplicacão deve buscar o resultado mais rápido entre duas APIs distintas.

As duas requisições devem ser feitas simultaneamente para as seguintes APIs:

https://brasilapi.com.br/api/cep/v1/01153000 + cep

http://viacep.com.br/ws/" + cep + "/json/

Os requisitos para este desafio são:

- Acatar a API que entregar a resposta mais rápida e descartar a resposta mais lenta.

- O resultado da request deverá ser exibido no command line com os dados do endereço, bem como qual API a enviou.

- Limitar o tempo de resposta em 1 segundo. Caso contrário, o erro de timeout deve ser exibido.


### Este repositório contém o projeto `api` escrito em Go.
 
## Configuração do Ambiente

Certifique-se de ter o Go instalado no seu sistema. Você pode baixá-lo em [golang.org](https://golang.org/dl/).

Após a instalação, baixe o projeto:
  
Para executar o projeto api, siga os passos abaixo:

1. Abra o terminal
2. `cd Multithreading `
3. `go mod tidy`
2. Navegue até a pasta 'server' (`cd api/server/`)
3. Execute o arquivo `main.go` com o comando `go run main.go`
4. A aplicação estará disponível em http://localhost:8000



Para testar a aplicacão

`curl http://localhost:8000/cep/cepvalido`


## Observações

Certifique-se de ter o Go instalado em sua máquina antes de tentar executar estes projetos.
