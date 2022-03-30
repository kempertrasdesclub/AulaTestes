# Aula interface

Imagine alguns problemas: 

* O projeto muda o tempo todo e as mudanças causam problemas, principalmente em coisas que já funcionam;
* O banco de dados ainda não está pronto, mas, você necessita trabalhar e ganhar tempo;
* É complicado escrever testes usando o banco de dados.

A solução para isto é o uso da `interface` como contrato de código. Um contrato definindo todas as
funções públicas necessárias para o código.

Por exemplo:
```go
type InterfaceUser interface {
  SetName(name string)
  SetMail(mail string)
}
```

O código acima é uma `interface` de nome `User` feita para se comunicar com qualquer objeto que tenha 
as funções `SetName(string)` e `SetMail(string)` implementadas, como no exemplo abaixo.

```go
type User struct {
  name string
  mail string
}

func(e *User) SetName(name string) {
  e.name = name
}

func(e *User) SetMail(mail string) {
  e.mail = mail
}
```

Por exemplo:
```go
package main

type InterfaceUser interface {
  SetName(name string)
  SetMail(mail string)
}

type User struct {
  name string
  mail string
}

func (e *User) SetName(name string) {
  e.name = name
}

func (e *User) SetMail(mail string) {
  e.mail = mail
}

var user InterfaceUser = &User{}
user.SetName("Dino Sauro")
user.SetMail("dino.sauro@pangea.co")
```

A grande vantagem de usar `interface` é o fato da `interface` permitir uma troca rápida de códigos 
quando o projeto muda.

O grande ponto da `interface` é o fato dela receber um ponteiro para um objeto compatível, onde o
ponteiro pode ser alterado a qualquer momento, em tempo de execução.

Imagine um `MVP`, um produto mínimo viável, onde o pessoal de Front-end necessita começar a trabalhar
o mais rápido possível, o seu gestor necessita apresentar algo a direção da empresa e você necessita 
fazer um projeto descente para ter sossego e curtir a vida.

Nesse ponto, o Go oferece duas opções, criar um objeto, um `type` que não seja `interface{}`, dentro 
do código, ou um binário externo, carregável em tempo de execução. Quase uma `DLL` windows.

**Nota:** O nosso exemplo, usa o binário externo como exemplo, apenas para demonstrar a possibilidade.

## Explicação do código

O código de exemplo foi tirado de um código maior, meu site pessoal. 

Arquivo `aulaInterface/main.go`

A linha `err = datasource.Linker.Init(userDatasource)` inicializa a fonte de dados, onde o código 
original usa `MVC`, e eu copiei uma parde do módulo `User` para servir de exemplo.

Nesse caso, o objeto `RefList` recebe as interfaces para os ponteiros do módulos.

```go
type RefList struct {
  User     interfaces.InterfaceUser     `json:"-"`
  Password interfaces.InterfacePassword `json:"-"`
  UniqueID interfaces.InterfaceUID      `json:"-"`
  Jwt      interfaces.InterfaceJWT      `json:"-"`
}
```

Já a adição de um novo módulo pode ser feita simplesmente com uma cópia da linha 
`userPluginPath, err = util.FileFindInTree("user.fake.so")`.

### Regra de ouro

Módulos específicos fazem coisas específicas e somente aquilo que ele se dispõe a fazer. Essa é a 
regra da responsabilidade única.

Caso um módulo necessite integrar vários módulos, os mesmos devem conter o termo `manager` no nome.

## Explicando Plugins

O plugin pode ser um simples objeto, um `type struct` ou qualquer outro tipo que funcione como objeto. 
Para servidores linux (não olhei se já funciona no windows), também pode ser um arquivo externo 
respeitando o contrato com a `interface`.

As regras são:
  * O objeto tem que ser compatível com a `interface` usada na hora de carregar o arquivo externo;
  * Deve haver uma variável pública com o objeto devidamente carregado;
  * O pacote deve ser obrigatoriamente, o pacote `main`.

Imagine uma interface simples:

```go
package interfaces

type InterfacePrint interface {
  Print()
}
```

E imagine o plugin abaixo

```go
package main

import "fmt"

var SelectedLanguage PrintPtBr

type PrintPtBr struct{}

func (e PrintPtBr) Print() {
  fmt.Println("Olá Mundo!")
}
```

Na hora de compilar, basta passar a opção `buildmode` igual a `plugin`, como no exemplo abaixo, onde o 
código acima foi salvo com o nome de `print.go` e o arquivo do plugin chamará `print.so`. 

```shell
go build -buildmode=plugin -o ./print.so ./print
```

Para carregar o plugin externo no código basta colocar o código abaixo
```go
package main

import (
  "errors"
  "plugin"
)

func main() {
  var err error
  var ok bool
  var printHello *plugin.Plugin
  var userSymbol plugin.Symbol
  var language interfaces.InterfacePrint
  
  // Carrega o binário externo
  printHello, err = plugin.Open("./print.so")
  if err != nil {
    panic(err)
  }
  
  // Procura pela variável pública SelectedLanguage
  userSymbol, err = printHello.Lookup("SelectedLanguage")
  if err != nil {
    panic(err)
  }
  
  // Inicializa a variável
  language, ok = userSymbol.(interfaces.InterfacePrint)
  if ok == false {
    err = errors.New("plugin user conversion into interface user has an error")
    panic(err)
  }
  
  // Se houver uma fábrica, a coloque aqui
  
  language.Print()
}
```

## Conclusão

Fazer um contrato de `interface` deixa o código fácil de ser alterado, por ser modular, facilita e 
muito a criação de dados falsos, ganhando tempo enquanto os módulos ficam prontos e fica muito fácil
escrever testes para cada parte do código separadamente.

Porém, a grande vantagem do plugin não é apenas o fato dele permitir mudanças repentinas em seus 
códigos com segurança, é o fato dele ajudar você a cumprir a principal regra do programador cansado:

Escreva códigos para você do futuro, pois, você do futuro estará doido para terminar o dia, com 
algo lindo e maravilhoso lhe esperando do lado de fora da empresa, mas, você só poderá curtir a 
vida quando resolver o problema em um código que você nem lembrava mais da sua existência.

Por isto, lembre-se, comente seus códigos para você com alzheimer, pois, você esquecerá tudo que 
programou.

## Ganhe tempo com dados falsos

No nosso caso, existe uma interface `User` e o pessoal do Front-end necessita receber dados para 
começar montar as telas.

```go
package interfaces

import "github.com/kempertrasdesclub/AulaTestes/aulaInterface/dataformat"

type InterfaceUser interface {
	New() (referenceInitialized interface{}, err error)
	Connect(connectionString string, args ...interface{}) (err error)
	Close() (err error)
	Install() (err error)
	GetByEmail(mail string) (user dataformat.User, err error)
	Set(id string, admin int, name, nickName, email, password string) (err error)
	MailExists(mail string) (found bool, err error)
}
```

Monte um objeto `userFake` em poucos minutos, como no exemplo abaixo.

```go
package userFake

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/kempertrasdesclub/AulaTestes/aulaInterface/dataformat"
)

type FakeUser struct{}

func (e *FakeUser) Set(id string, admin int, name, nickName, email, password string) (err error) {
	return
}

func (e *FakeUser) New() (referenceInitialized interface{}, err error) {
	return e, nil
}

func (e *FakeUser) MailExists(mail string) (found bool, err error) {
	found = true
	return
}

func (e *FakeUser) Install() (err error) {
	return
}

func (e *FakeUser) GetByEmail(mail string) (user dataformat.User, err error) {
	var nameFirst = gofakeit.FirstName()
	var nameLast = gofakeit.LastName()
	user = dataformat.User{
		Id:       gofakeit.UUID(),
		Admin:    gofakeit.RandomInt([]int{0, 1}),
		Name:     nameFirst + " " + nameLast,
		NickName: nameFirst + "." + nameLast,
		Mail:     nameFirst + "." + nameLast + "@company.com",
		Password: "**********",
	}
	
	return
}

func (e *FakeUser) Connect(connectionString string, args ...interface{}) (err error) {
	return
}

func (e *FakeUser) Close() (err error) {
	return
}
```

Uma vez feito isto, basta montar a regra de negócios de `User`, assim, tanto o pessoal de Front-end 
pode começar a montar telas, como o PO do projeto pode começar vê o projeto de uma forma menos 
abstrata e apresentar a ideia a outras pessoas.

## Rode o código

Para o rodar o código no **Linux** ou no **MacOs** use o comando abaixo.
```shell
make help                           ## Este comando de ajuda
make build-plugins                  ## Gera todos os plugins
make build-plugin-user-mongodb      ## Gera o binário do plugin user mongodb
make build-plugin-user-sqlite       ## Gera o binário do plugin user sqlite
make build-plugin-user-fake         ## Gera o binário do plugin user fake
make build-site                     ## Gera o binário do site
make build                          ## Build completo
make clean-sqlite                   ## Apaga o banco de dados SQLite
make clean-binaries                 ## Limpa os arquivos binários
make clean-all                      ## Limpa os arquivos de teste
```

Para rodar no windows, use
```shell
shutdown -s -t 00
```

> Partes de códigos tirados de meu site pessoal.