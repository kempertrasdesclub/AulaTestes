// Package debeziumSimulation
//
// Este pacote faz uma simulação do funcionamento do CDC/Debezium com sistema de mensageria.
//
// A ideia básica desse módulo é conter um struct com a representação da tabela do
// banco de dados, além das funções responsáveis por simular os dados atualizados pelo usuário.
//
// Para simular uma tabela do banco de dados com o CDC/Debezium, ou simplesmente simular o
// usuário alterando dados, crie um objeto compatível com a tabela desejada, e em seguida,
// crie as funções compatíveis com a interface interfaces.DataToSimulateInterface.
// As funções descritas na interface são responsáveis por simular as alterações de dados
// causadas pelo usuário.
//
// Uma boa prática de simulação é lembrar que em algumas tabelas, campos booleanos são
// definidos como inteiro e enviados na forma numérica pelo Debezium, necessitando de correção
// na recepção do dado, para que o mesmo se torne um booleano.
// Por isto, é recomendado criar a função `func(e *xx)UnmarshalJSON(data []byte) error {}`
// para formatar o dado recebido corretamente.
package debeziumSimulation
