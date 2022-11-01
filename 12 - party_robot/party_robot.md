Instruções
Era uma vez um programador excêntrico morando em uma casa estranha com janelas gradeadas. Um dia, ele aceitou um emprego em um quadro de empregos online para construir um robô de festa. O robô deve cumprimentar as pessoas e ajudá-las a se sentarem. A primeira edição foi muito técnica e mostrou a falta de interação humana do programador. Alguns dos quais também chegaram à próxima edição.

1. Dê as boas-vindas a um novo convidado para a festa
Implemente a Welcomefunção para retornar uma mensagem de boas-vindas usando o nome fornecido:

Welcome("Christiane")
// => Welcome to my party, Christiane!

2. Dê as boas-vindas a um novo convidado para a festa cujo aniversário é hoje
Implemente a HappyBirthdayfunção para retornar uma mensagem de aniversário usando o nome e a idade da pessoa. Infelizmente, o programador é um pouco exibido, então o robô tem que demonstrar seu conhecimento sobre o aniversário de cada convidado.

HappyBirthday("Frank", 58)
// => Happy birthday Frank! You are now 58 years old!

3. Dê instruções
Implemente a AssignTablefunção para dar direções. Deve aceitar 5 parâmetros.

O nome do novo convidado a ser recebido ( string)
O número da tabela ( int)
O nome do companheiro de assento ( string)
A direção onde encontrar a mesa ( string)
A distância até a mesa ( float64)
O formato de resultado exato pode ser visto no exemplo abaixo.

O robô fornece o número da mesa em formato de 3 dígitos. Se o número tiver menos de 3 dígitos, obtém zeros à esquerda extras para se tornarem 3 dígitos (por exemplo, 3 torna-se 003). O robô também menciona a distância da mesa. Felizmente, apenas com uma precisão limitada a 1 dígito.

AssignTable("Christiane", 27, "Frank", "on the left", 23.7834298)
// =>
// Welcome to my party, Christiane!
// You have been assigned to table 027. Your table is on the left, exactly 23.8 meters from here.
// You will be sitting next to Frank.
