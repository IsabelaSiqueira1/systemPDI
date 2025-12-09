## UML - Unified Modeling Language

A UML foi desenvolvida por Grady Booch, Ivar Jacobson e James Rumbaugh na Rational Software em meados dos anos 90.

A Linguagem de Modelagem Unificada, é uma linguagem de modelagem de propósito geral. Ela é apresentada na forma de diagramas e habilita os profissionais de tecnologia a modelar e documentar aplicações de software.A UML permite visualizar o projeto ou design de um sistema por meio de uma representação gráfica.

A linguagem não é um método ou uma linguagem de desenvolvimento, ou seja, você não programa em UML. Ela serve como auxílio para projetar o sistema de forma visual, sendo independente de plataforma ou linguagem.

O elemento central da UML é o diagrama, que é uma representação gráfica do modelo de um sistema. Os diagramas permitem representar o sistema de duas maneiras distintas: estática (estrutural) ou dinâmica (comportamental).

## 1. Diagramas Estruturais: Diagramas que mostram a estrutura estática do sistema, incluindo objetos, operações, atributos e métodos. Exemplos incluem o diagrama de classes, objeto, componente e pacote.

## 2. Diagramas Comportamentais: Mostram o comportamento dinâmico e o funcionamento do sistema, por meio da colaboração dos objetos e da mudança de estados internos. Exemplos incluem o diagrama de casos de uso, sequência, comunicação e atividade.


## Caso de uso (Estrutura comportamental)
(em andamento)
...

## Diagrama de classes (Estrutura estática)
O diagrama de classes é usado para descrever a estrutura estática de um sistema, permitindo definir os atributos, operações (métodos) e relacionamentos entre as classes.

Apresenta uma visão estatica da organização das classes, definindo sua estrutura logica.

É um dos diagramas mais populares, e serve como base para a construção de outros diagramas UML. Basicamentre, ele descreve o que deve estar presente no sistema modelado.


## Classes, atributos e metodos:
Uma classe é uma representação de um item do mundo real, fisico e abstrato, na forma de que um tipo de dados personalizado.

As Classes possuem estruturas internas chamadas de Atributos e de Métodos.

Atributos são usados para armazenar os dados dos objetos de uma classe.

Metodos são as operações, funções que a instancia da classe pode realizar.

Uma instancia de uma classe é chamada de Objeto.

## Exemplo: 
Classe: Pessoa
Atributos:
 - Altura
 - Nome
 - Idade
 - Peso

Metodos:
 - Andar()
 - Correr()
 - Falar()
 - Dormir()
 - Comer()
 - Trabalhar()
 - Estudar()

Objetos da classe (Instancia):
Atrubitos:
 - Altura: 1,75m
 - Nome: João
 - Idade: 30
 - Peso: 80kg

 
## Representação de uma classe:
Representamos uma classe usando o diagrama dividido em tres partes:
 - Nome da classe: inclui o nome da classe (informações sobre a classe)
 - Atributos: Lista de atributos da classe no formato nome:tipo ou nome:tipo = valor
 - Metodos: Lista de metodos da classe no formato nome(parametros):tipo_retorno

## Visibilidade dos membros (Atributos e Metodos):
Representamos a visibilidade dos membros (atributos e metodos) usando os seguintes simbolos:
 - + : Publico (acessivel por qualquer classe)
 - - : Privado (acessivel apenas pela propria classe)
 - # : Protegido (acessivel pela propria classe e suas subclasses)
 - ~ : Pacote (acessivel apenas por classes do mesmo pacote) 


## Relacionamentos entre classes: 
Um relacionamento é uma conexão entre itens. Existem varios tipos de relacionamento possiveis entre classes:
 - dependencia
 - associação
 - agregação
 - composição
 - generalização


# Relacionamento de dependencia:
Dependencia fraca, usualmente transiente, que ilustra que uma classe usa informações e serviços  de outra classe em algum momento, dependendo dela.

Do tipo "A classe A depende da classe B".

Exemplo:
Carro A --------------> Roda B


## Multiplicidade:
A multiplicidade é usada para determinar o numero minimo e o numero maximo de objetos envolvidos na associação, de cada lado, e tambem pode especificar o nivel de dependencia entre os objetos.

Exemplos comuns de multiplicidade:
  - 0..1 (No minimo zero, no maximo um. Indica não obrigatoriedade do relacionamento)
  - 1..1 (Um e somente um. Um objeto de classe se relaciona com um objeto da outra classe)
  - 0..* (Minimo nenhuma e no maximo muitos. Indica que um objeto pode se relacionar com varios objetos da outra classe)
  - 1..* (Minimo um e no maximo muitos. Indica que um objeto deve se relacionar com varios objetos da outra classe)
  - * (Muitos)


## Relacionamento de Associação:
Relacionamento mais forte do que a independencia, indica que uma classe mantem uma referencia a outra classe ao longo do tempo. As associações podem conectar mais de duas classes.

Do tipo "Classe A tem uma Classe B"

A associação pode ter um nome e possui multiplicidade.

Exemplo:
          *  assina  *
Pessoa A ————————————> Revista B

A seta representa navegabilidade, que identifica o sentido em que as informações são transmitidas entre os objetos das classes relacionadas.


## Relacionamenot Agregação:
Relacionamento mais especifico do que a associação, indica que uma classe é um container ou uma coleção de outras classes.  As classes contidas não dependem do container - assim, quando o container é destruido, as classes continuam existindo.

Do tipo "Classe A possui uma Classe B"

Exemplo:
              1..*  possui    1
Departamento A <>———————————— Instrutor B


## Relacionamento de Composição:
Variação mais especifico da agregação, este relacionamento indica uma dependencia de ciclo de vida forte entre as classes, de modo que quando um container é destruido, seu conteudo tambem é destruido.

Do tipo "Classe A é parte da Classe B"


Exemplo:

Janela A ◀▶———————————— BarraMenus B


## Relacionamento de Generalização/Especialização:
Relacionamento entre iteins gerais(superclasses / classes-mãe) e tipos mais especificos deste itens (subclasses / classes-filhas). Representa a Herança entre as classes.

A classe filha herda propriedades da classe mãe. Principalmente atributos e metodos, e pode possuir seus proprios atributos e metodos adicionais.

Do tipo "Classe B é um tipo de Classe A"


Exemplo:

Animal A <———————————— Peixe B
