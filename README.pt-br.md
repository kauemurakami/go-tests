[![pt-br](https://img.shields.io/badge/language-pt--br-green.svg)](https://github.com/kauemurakami/go-tests/blob/main/README.pt-br.md)
[![en](https://img.shields.io/badge/language-en-orange.svg)](https://github.com/kauemurakami/go-tests/blob/main/README.md)  
go version 1.22.1  

## Go tests
Testes automazidos em GO, resumindo, ele vai ser uma função que vai testar outra função sua e ver se o resultado dela é o que você está esperando realmente. Uma pratica muito comum, para que você possa garantir o comportamento das coisas.<br/><br/>
Imagine que você tenha uma função que recebe dois parâmetros e ele deve retornar um valor ou tipo específico. Os testes existem pra garantir que sua função, recebendo esses parâmetros vai de fato retornar o resultado que você está esperando.<br/><br/>
É um jeito de você garantir que o que você implementou está certo, e que vai continuar certo ao longo do tempo, os testes te dão uma segurança muito grande no código, imagine que você tem uma grande função que esta funcionando e retorne o que espero que retorne, e amanha fazemos uma alteração, se ela parar de devolver o resultado por conta da alteração, nosso teste irá acusar pra gente, fazendo você reavaliar o novo comportamente ou algum efeito colateral não previsto.<br/><br/>

### Introdução aos tests
Fazendo um teste bem simples, vamos criar um diretório ```go-tests```, crie o arquivo ```main.go``` preencha com o código básico ou use a extensão [go-fast-snippets](https://marketplace.visualstudio.com/items?itemName=go-snippets.go-fast-snippets) disponível para vscode, com ela você precisaria apenas começar a escrever ```gomain``` e o código será gerado, caso não possua este é o código básico:  
```go
package main

func main() {
}
```  
Agora crie um diretório interno chamado ```addresses``` e nele crie um arquivo ```addresses.go```, caso possua nossa extensão [go-fast-snippets](https://marketplace.visualstudio.com/items?itemName=go-snippets.go-fast-snippets) basta começar a digitar ```gofile``` no arquivo em branco e o código será gerado, caso não veja:  
```go
package addresses

func addressType(address string) string {
  validTypes := []string{
    "street", "avenue", "road", "highway",
  }
}
```
Nele vamos criar essa função que irá verificar se o ```address``` passado contém no início dela algum dos ```validTypes``` pré definidos.   

Agora vamos terminar a função:  
```go
package addresses

import "strings"

func addressType(address string) string {
	validTypes := []string{
		"street", "avenue", "road", "highway",
	}
	// address in lowercase
	lowercaseAddress := strings.ToLower(address)
	// Split text in array separing peer empty spaces
	// ex split with empty space result 0-RUA 1-ABC 2-DEF
	// and set in firstWordAddress recovering position 0
	// of the created array
	firstWordAddress := strings.Split(lowercaseAddress, " ")[0]

	isValid := false //first word is valid or not

	for _, t := range validTypes { //iterate with validTypes and check is valid
		if t == firstWordAddress { // if compatible 
			isValid = true //isvalid is true
		}
	}

	if isValid {
		return firstWordAddress // return type "first word of address"
	}
	return "Invalid type" // case not match return message
}
```  
Aqui inicialmente alteramos nossos valores de ```validTypes``` para todas as palavras serem minúsculas, logo depois criamos a variável ```lowercaseAddress``` convertendo a palavre do nosso parâmetro ```address``` para que todas as letras sejam minúsculas, logo depois criamos a variável ```firstWordAddress``` que recebe a primeira palavra do nosso ```lowercaseAddress```, após fazendo um ```split``` nele, separando cada palavra separada por espaço, logo o indice ```[0]``` é aprimeira palavra.  
Agora criamos a variável ```isValid``` com valor inicial ```false```, logo em seguida iteramos ```validTypes``` e verificamos se um dos valores dele é compativel com ```firstWordAddress```, caso seja ```isValid``` recebe true, e é verificado ao sair do loop, e se for válido retornamos ```firstWordAddress```, caso contrário retornamos uma mensagem. E é nossa função.<br/><br/>

Agora no nível do diretório raiz em nosso terminal, vamos criar um módulo:  
```shell
go mod init go-tests
```
E você tera um arquivo ```go.mod``` com este conteúdo inicialmente:  
```go
module go-tests

go 1.22.1

```
Lembrando que como não estamos usando pacotes externos não há problemas em criar o ```go.mod``` depois de criar nosso package interno, no caso ```addresses.go```, mas poderia ocorrer problemas caso usassemos pacotes externos à aplicação.  
Uma correção, como nossa função ```addressType``` vai ser importada, temos que iniciar o nome dela com a primeira letra MAIÚSCULA, ou seja nosso A de address:  
```go
package addresses
....
// Verify if address contains a valid type in first word
func AddressType(address string) string {
...
}
```  
Agora voltando à nossa ```main.go``` e vamos chamar no pacote:  
```go
package main

import (
	"fmt"
	"go-tests/addresses"
)

func main() {
	typeAddress := addresses.AddressType("Street dos bobos")
	// typeAddress := addresses.AddressType("abc dos bobos") // output Invalid Type

	fmt.Println(typeAddress)
}
```
Também possuimos algumas alterações na nossa função ```AddressType```:  
Primeiro em seu terminal instale o seguinte pacote externo:  
```shell
go get golang.org/x/text/cases
```
Agora dentro da função procure por estas linhas e edite:  
```go
...
import (
	"strings"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)
...
if isValid {
		caser := cases.Title(language.BrazilianPortuguese) //set language
		return caser.String(firstWordAddress) // To uppercase first letter
	}
...
```
Usamos este pacote para deixar a primeira letra de ```firstWordAddress``` maiúscula e para nos ambientarmos com os pacotes externos *Agora vamos para os testes*.

### Teste unitário básico
Vamos usar um pacote do go chamado ```test```, e apra esse pacote funcionar corretamente com nossas funções existem algumas regras a serem seguidas.  
1 - O teste de uma função nunca fica no mesmo arquivo da função em sí.  
2 - Os arquivos de teste para serem reconhecidos pelo go eles devem ter um nome específico, o nome do ```arquivo.go``` deve ser alterado, quando estamos criando uma função de teste para ```arquivo_test.go```, isso se deve pois para rodar todos os testes, vamos rodar pela linha de comando, e esse comando vai entrar nos arquivo que possuem ```arquivos_test.go``` e começar a executar as funções de teste dentro dele, portanto essa nomenclatura é obrigatório.<br/><br/>

### Criando arquivo de teste unitário
Dentro do próprio diretório de ```address```, que possui o arquivo ```addresses.go```, crie um arquivo chamado ```addresses_test.go```.  
Um *teste unitário* é um teste que vai testar a menor unidade do seu código, no nosso caso  nossa função ```AddressType```, também existem teste de integração, que cobrem um escopo um pouco maior, varias funções, fluxos completos, veremos adiante.  
*Assinatura de um código de teste*  
```go
package addresses_test

import "testing"

func TestAddressType(t *testing.T) {
  address_to_test := "Avenue Paulista" // address used to testing
	expected_address_type := "Avenue" // expected type
  // run tested function
	receivedAddressType := addresses.AddressType(address_to_test) 

	if receivedAddressType != expected_address_type {
		// método do param t, ele chama um erro no seu test
		// o erro sera logado no terminal e será considerado que
		// quebrou ou não está fazendo o que esperamos
		t.Error("Received type invalid")
	}
}
```
Reparem que o nome do package é igual a do arquivo ```addresses.go```, neste diretório, o go dá essa exceção que você pode ter dois pacotes diferentes dentro da mesma pasta.  
Outro detalhe é o uso do pacote do go ```testing```, ao criar nossa função ```TestAddressType``` ela recebe um parâmetro, comumente de t, e seu tipo é um ponteiro de ```(t *testing.T)```.<br/><br/>

Também a função deve começar com a plavra ```Test``` com T maiúsculo, em inglês, e o nome da função que vamos testar, começando com a letra maiúscula, no nosso caso ```TestAddressType```, após ```Test```, obrigatóriamente a próxima letra deve ser maiúscula.  
Usando essa nomenclatura de arquivo juntamente com a sintax da função, o go vai identificar essa função pra ser testada.<br/><br/> 

Nesta função adicionamos um valor à variável que será testada ```address_to_test```, também definimos um tipo esperado em ```expected_address_type```, nossa variável ```receivedAddressType``` recebe o resultado da nossa função ```AddressType(address_to_test)```.  
Após obter o resultado fazemos a verificação no nosso if se ```receivedAddressType``` é diferente do nosso resultado esperado ```expected_address_type``` e caso seja diferente, sinal de retorno não esperado da função, chamamos ```t.Error()``` que irá logar o erro no terminal e o go vai considerar que seu teste quebrou, se não cair no erro, vamos considerar que o teste passou.  
Agora uma pequena alteração no no if de verificação do nosso teste para exibir o que esperava e oq recebeu:  
```go
...
func TestAddressType(t *testing.T) {
  if receivedAddressType != expected_address_type {
		// método do param t, ele chama um erro no seu test
		// o erro sera logado no terminal e será considerado que
		// quebrou ou não está fazendo o que esperamos
		t.Errorf("Received type invalid, wait %s and receive %s",
			expected_address_type,
			receivedAddressType,
		)
	}
}
...
```  
Agora abra o terminal dentro do diretório do package ```/addresses``` rode o comando ```go test```  

### Teste unitário com mais de um cenário
Vamos fazer isso agora refatorando o test passado:  
```go
...
type test_scenario struct {
	address_inserted string
	expected_return  string
}

func TestAddressType(t *testing.T) {
	scenarios_of_Test := []test_scenario{
		{"Street Abc", "Street"},
		{"Avenue xyz", "Avenue"},
		{"Road 138", "Road"},
		{"Square park", "Invalid type"},
		{"HiGhway dbo", "Highway"},
		{"", "Invalid type"},
	}

	for _, scenario := range scenarios_of_Test {
		receivedAddressType := AddressType(scenario.address_inserted)

		if receivedAddressType != scenario.expected_return {
			// método do param t, ele chama um erro no seu test
			// o erro sera logado no terminal e será considerado que
			// quebrou ou não está fazendo o que esperamos
			t.Errorf("Received type invalid, wait %s and receive %s",
				scenario.expected_return,
				receivedAddressType,
			)
		}
	}
}
```
Rode no terminal dentro de ```/addresses``` ```go test``` e veja o resultado.  

### Dicas
```go test ./...``` entra em todas os packages verificando os aquivos test  
```go test -v``` modo verboso do test  
Para rodar em paralelo podemos adicionar no inicio da função, caso  haja mais de uma função de test no seu arquivo teste ```t.Parallel()```  e deve ser adicionado em todas as funções que você deseja que rode em paralelo.   

```go test --cover``` Mostra se seu cenário está sendo coberto 100%, mostra a % da cobertura dos seus estados / execuções, ou seja todas as linhas da função que estamos testando estão cobertas.  

```go test --coverprofile doc.txt``` contem um relatorio das linhas que estao cobertas e que nao estao  

```go tool cover --func=doc.txt``` vai ler o arquivo txt entender e jogar no terminal

```go tool cover --html=doc.txt``` mostra um arquivo html que vai ter um visual bonitinho de todas as linhas nao cobertas  

### Subtests
Create a new directory called forms with two ```form.go``` files,```form_test.go```:  
```form.go```  
```go
package form

import (
	"math"
)

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}
func (c Circle) Area() float64 {
	return math.Pi * (c.Rad * c.Rad)
}

type Circle struct {
	Rad float64
}

type Rectangle struct {
	Height float64
	Width  float64
}

type Form interface {
	Area() float64
}
```  
```form_test.go```  
```go
package form

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	t.Run("Rectangle area", func(t *testing.T) {
		r := Rectangle{10, 12}
		expectedArea := float64(120)
		receivedArea := r.Area()

		if expectedArea != receivedArea {
			//fatal vai parar os testes
			t.Fatalf("Received area %f, expected is %f", receivedArea, expectedArea)
		}
	})
	t.Run("Circle area", func(t *testing.T) {
		c := Circle{10}

		expectedArea := float64(math.Pi * 100)
		receivedArea := c.Area()
		if expectedArea != receivedArea {
			//fatal vai parar os testes
			t.Fatalf("Received area %f, expected is %f", receivedArea, expectedArea)
		}
	})
}
```  
Here we make a group of tests, running one after the other in sequence, if you don't want your application to stop the tests when giving an error, just change ```t.Fatalf``` to ```t.Errof```










