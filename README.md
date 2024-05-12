[![pt-br](https://img.shields.io/badge/language-pt--br-green.svg)](https://github.com/kauemurakami/go-tests/blob/main/README.pt-br.md)
[![en](https://img.shields.io/badge/language-en-orange.svg)](https://github.com/kauemurakami/go-tests/blob/main/README.md)  
go version 1.22.1  

## Go tests
Automated tests in GO, in short, it will be a function that will test another function of yours and see if its result is what you are really expecting. A very common practice, so you can guarantee the behavior of things.<br/><br/>
Imagine you have a function that takes two parameters and it must return a specific value or type. The tests exist to ensure that your function, receiving these parameters, will actually return the result you are expecting.<br/><br/>
It's a way for you to ensure that what you implemented is right, and that it will continue to be right over time, tests give you great security in the code, imagine that you have a great function that is working and returns what I expect that returns, and tomorrow we make a change, if it stops returning the result due to the change, our test will report it to us, making you reevaluate the new behavior or any unforeseen side effect.<br/><br/>

### Introduction to testing
Doing a very simple test, let's create a ```go-tests``` directory, create the ```main.go``` file, fill it with the basic code or use the extension [go-fast-snippets](https://marketplace.visualstudio.com/items?itemName=go-snippets.go-fast-snippets) available for vscode, with it you would just need to start writing ```gomain``` and the code will be generated, otherwise This is the basic code:  
```go
package main

func main() {
}
```  
Now create an internal directory called ```addresses``` and create a file ```addresses.go``` in it, if it has our extension [go-fast-snippets](https://marketplace.visualstudio.com/items?itemName=go-snippets.go-fast-snippets) just start typing ```gofile``` in the blank file and the code will be generated, if you don't see:   
```go
package addresses

func addressType(address string) string {
  validTypes := []string{
    "street", "avenue", "road", "highway",
  }
}
```
In it we will create this function that will check if the passed ```address``` contains any of the pre-defined ```validTypes``` at the beginning.  

Now let's finish the function:  
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
Here we initially change our ```validTypes``` values ​​so that all words are lowercase, then we create the variable ```lowercaseAddress``` converting the word in our ```address``` parameter so that all letters are lowercase, then we create the variable ```firstWordAddress``` which receives the first word of our ```lowercaseAddress```, after doing a ```split``` on it, separating each word separated by space, then the index ```[0]``` is the first word.
Now we create the variable ```isValid``` with an initial value ```false```, then we iterate ```validTypes``` and check if one of its values ​​is compatible with ```firstWordAddress``` , if it is ```isValid``` it receives true, and is checked when exiting the loop, and if it is valid we return ```firstWordAddress```, otherwise we return a message. And it's our job.<br/><br/>

Now at the root directory level in our terminal, let's create a module:   
```shell
go mod init go-tests
```
And you will have a ```go.mod``` file with this content initially:  
```go
module go-tests

go 1.22.1

```
Remembering that as we are not using external packages there is no problem creating ```go.mod``` after creating our internal package, in this case ```addresses.go```, but problems could occur if we used external packages the app.  
One correction, as our ```addressType``` function will be imported, we have to start its name with the first letter UPPERCASE, that is, our A for address:  
```go
package addresses
....
// Verify if address contains a valid type in first word
func AddressType(address string) string {
...
}
```  
Now going back to our ```main.go``` and let's call the package:   
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
We also have some changes to our ```AddressType``` function:  
First in your terminal install the following external package:  
```shell
go get golang.org/x/text/cases
```
Now inside the function look for these lines and edit:  
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
We use this package to make the first letter of ```firstWordAddress``` capitalized and to get used to external packages *Now let's move on to testing*.  

### Basic unit testing
We are going to use a Go package called ```test```, and for this package to work correctly with our functions there are some rules to be followed.  
1 - The test of a function is never in the same file as the function itself.  
2 - To be recognized by go, the test files must have a specific name, the name of the ```arquivo.go``` must be changed, when we are creating a test function for ```arquivo_test.go`` `, this is because to run all the tests, we will run it through the command line, and this command will enter the files that have ```arquivos_test.go``` and start executing the test functions within it, so this nomenclature is mandatory.<br/><br/>

### Creating unit test file
Inside the ```address``` directory itself, which has the ```addresses.go``` file, create a file called ```addresses_test.go```.  
A *unit test* is a test that will test the smallest unit of your code, in our case our ```AddressType``` function, there are also integration tests, which cover a slightly larger scope, several functions, complete flows , we will see later.  
*Signing a test code*  
```go
package addresses_test

import "testing"

func TestAddressType(t *testing.T) {
  address_to_test := "Avenue Paulista" // address used to testing
	expected_address_type := "Avenue" // expected type
  // run tested function
	receivedAddressType := addresses.AddressType(address_to_test) 

	if receivedAddressType != expected_address_type {
		// param t method, it calls an error in your test
    // the error will be logged in the terminal and it will be considered that
    // broken or not doing what we expect
		t.Error("Received type invalid")
	}
}
```
Note that the package name is the same as the file ```addresses.go```, in this directory, go gives this exception that you can have two different packages within the same folder.
Another detail is the use of the go package ```testing```, when creating our function ```TestAddressType``` it receives a parameter, commonly t, and its type is a pointer of ```(t * testing.T)```.<br/><br/>

The function must also start with the word ```Test``` with a capital T, in English, and the name of the function we are going to test, starting with the capital letter, in our case ```TestAddressType```, after ` ``Test```, the next letter must be capitalized.
Using this file naming along with the function syntax, go will identify this function to be tested.<br/><br/>

In this function we add a value to the variable that will be tested ```address_to_test```, we also define an expected type in ```expected_address_type```, our variable ```receivedAddressType``` receives the result of our function ``` AddressType(address_to_test)```.  
After obtaining the result, we check in our if if ```receivedAddressType``` is different from our expected result ```expected_address_type``` and if it is different, an unexpected return signal from the function, we call ```t. Error()``` which will log the error in the terminal and go will consider that your test broke, if it doesn't show an error, we will consider that the test passed.  
Now a small change in the if check of our test to display what you expected and what you received:  
```go
...
func TestAddressType(t *testing.T) {
  if receivedAddressType != expected_address_type {
		// param t method, it calls an error in your test
    // the error will be logged in the terminal and it will be considered that
    // broken or not doing what we expect
		t.Errorf("Received type invalid, wait %s and receive %s",
			expected_address_type,
			receivedAddressType,
		)
	}
}
...
```  
Now open the terminal inside the package directory ```/addresses``` and run the command ```go test```    

### Unit testing with more than one scenario
Let's do this now by refactoring the past test:   
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
			// param t method, it calls an error in your test
      // the error will be logged in the terminal and it will be considered that
      // broken or not doing what we expect
			t.Errorf("Received type invalid, wait %s and receive %s",
				scenario.expected_return,
				receivedAddressType,
			)
		}
	}
}
```
Run in the terminal inside ```/addresses``` ```go test``` and see the result. 











