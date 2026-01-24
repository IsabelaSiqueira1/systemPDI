## UML - Unified Modeling Language

A UML foi desenvolvida por Grady Booch, Ivar Jacobson e James Rumbaugh na Rational Software em meados dos anos 90.

A Linguagem de Modelagem Unificada, é uma linguagem de modelagem de propósito geral. Ela é apresentada na forma de diagramas e habilita os profissionais de tecnologia a modelar e documentar aplicações de software.A UML permite visualizar o projeto ou design de um sistema por meio de uma representação gráfica.

A linguagem não é um método ou uma linguagem de desenvolvimento, ou seja, você não programa em UML. Ela serve como auxílio para projetar o sistema de forma visual, sendo independente de plataforma ou linguagem.

O elemento central da UML é o diagrama, que é uma representação gráfica do modelo de um sistema. Os diagramas permitem representar o sistema de duas maneiras distintas: estática (estrutural) ou dinâmica (comportamental).

### 1. Diagramas Estruturais: Diagramas que mostram a estrutura estática do sistema, incluindo objetos, operações, atributos e métodos. Exemplos incluem o diagrama de classes, objeto, componente e pacote.

### 2. Diagramas Comportamentais: Mostram o comportamento dinâmico e o funcionamento do sistema, por meio da colaboração dos objetos e da mudança de estados internos. Exemplos incluem o diagrama de casos de uso, sequência, comunicação e atividade.

## Caso de uso (Estrutura comportamental)

O **diagrama de caso de uso** é um diagrama comportamental da UML que descreve **quais funcionalidades** um sistema oferece e **quem** (atores) interage com essas funcionalidades. Ele é usado para entender o sistema do ponto de vista do usuário e para alinhar requisitos sem entrar em detalhes de implementação.

### Para que serve

- Delimitar o **escopo** do sistema (o que está dentro e fora).
- Identificar **funcionalidades** (casos de uso).
- Identificar **atores** (tipos de usuários e sistemas externos).
- Apoiar a escrita de **requisitos** e **cenários**.
- Servir de base para outros diagramas (ex.: sequência).

### Elementos principais

- **Ator (Actor)**: entidade externa que interage com o sistema. Pode ser:
  - Pessoa (usuário)
  - Outro sistema (integração)
  - Dispositivo
- **Caso de uso (Use case)**: funcionalidade observável do sistema. Normalmente escrito como **verbo + objeto**, por exemplo: “Cadastrar X”, “Consultar Y”.
- **Boundary (Sistema)**: caixa que representa a fronteira do sistema; tudo dentro dela é responsabilidade do software.
- **Associação**: linha entre ator e caso de uso indicando participação.
- **Relacionamentos entre casos de uso**:
  - `<<include>>`: um caso de uso **sempre** reutiliza outro como parte do fluxo (obrigatório).
  - `<<extend>>`: um caso de uso **pode** estender outro em uma condição específica (opcional/condicional).
  - **Generalização**: ator ou caso de uso pode especializar outro (herança).

### Como escrever um caso de uso (texto)

Um caso de uso textual normalmente contém:

- **Nome** (verbo + objeto)
- **Ator principal**
- **Objetivo**
- **Pré-condições** (o que precisa ser verdade antes do fluxo)
- **Fluxo principal** (passo a passo de sucesso)
- **Fluxos alternativos / exceções** (erros, validações, situações alternativas)
- _(Opcional)_ **Pós-condições** (estado final após sucesso)

### Boas práticas

- Descreva **o comportamento**, não a implementação (evite falar de banco, classes, tabelas, etc.).
- Mantenha casos de uso focados (evite “casos gigantes”).
- Use `include` para partes comuns repetidas.
- Use `extend` quando algo acontece só em certos cenários.
- Atores representam **papéis**, não pessoas específicas.

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

- - : Publico (acessivel por qualquer classe)
- - : Privado (acessivel apenas pela propria classe)
- (#): Protegido (acessivel pela propria classe e suas subclasses)
- ~ : Pacote (acessivel apenas por classes do mesmo pacote)

## Relacionamentos entre classes:

Um relacionamento é uma conexão entre itens. Existem varios tipos de relacionamento possiveis entre classes:

- dependencia
- associação
- agregação
- composição
- generalização

# Relacionamento de dependencia:

Dependencia fraca, usualmente transiente, que ilustra que uma classe usa informações e serviços de outra classe em algum momento, dependendo dela.

Do tipo "A classe A depende da classe B".

Exemplo:
Carro A --------------> Roda B

## Multiplicidade:

A multiplicidade é usada para determinar o numero minimo e o numero maximo de objetos envolvidos na associação, de cada lado, e tambem pode especificar o nivel de dependencia entre os objetos.

Exemplos comuns de multiplicidade:

- 0..1 (No minimo zero, no maximo um. Indica não obrigatoriedade do relacionamento)
- 1..1 (Um e somente um. Um objeto de classe se relaciona com um objeto da outra classe)
- 0..\* (Minimo nenhuma e no maximo muitos. Indica que um objeto pode se relacionar com varios objetos da outra classe)
- 1..\* (Minimo um e no maximo muitos. Indica que um objeto deve se relacionar com varios objetos da outra classe)
- (\*) (Muitos)

## Relacionamento de Associação:

Relacionamento mais forte do que a independencia, indica que uma classe mantem uma referencia a outra classe ao longo do tempo. As associações podem conectar mais de duas classes.

Do tipo "Classe A tem uma Classe B"

A associação pode ter um nome e possui multiplicidade.

Exemplo:
_ assina _
Pessoa A ————————————> Revista B

A seta representa navegabilidade, que identifica o sentido em que as informações são transmitidas entre os objetos das classes relacionadas.

## Relacionamenot Agregação:

Relacionamento mais especifico do que a associação, indica que uma classe é um container ou uma coleção de outras classes. As classes contidas não dependem do container - assim, quando o container é destruido, as classes continuam existindo.

Do tipo "Classe A possui uma Classe B"

Exemplo:
1..\* possui 1
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

## Diagrama de Sequencia

O **diagrama de sequência** é um diagrama comportamental da UML que mostra, ao longo do tempo, **a ordem das interações** (mensagens) entre um ator e os componentes do sistema para executar um cenário (geralmente um caso de uso ou parte dele).

Ele detalha “como” um comportamento ocorre de forma temporal, sem precisar entrar no nível de código, mas já aproximando da implementação.

### Para que serve

- Entender a **ordem das chamadas**.
- Explicitar **responsabilidades** entre camadas/objetos.
- Identificar **validações**, retornos e ramificações (erros/alternativas).
- Ajudar a planejar endpoints, serviços, módulos e integrações.

### Elementos principais

- **Participantes (lifelines)**: ator, objetos ou componentes que participam do fluxo.
- **Mensagens**: setas representando chamadas:
  - Chamada síncrona (espera resposta)
  - Chamada assíncrona (não espera)
- **Ativações**: períodos em que um participante está executando algo.
- **Retornos**: respostas para mensagens (opcional desenhar, mas útil).
- **Fragmentos combinados**:
  - `alt`: caminhos alternativos (if/else)
  - `opt`: bloco opcional (if simples)
  - `loop`: repetição
  - `par`: paralelo

### Como montar um diagrama de sequência a partir de um caso de uso

1. Escolha o **cenário**: normalmente o **fluxo principal** do caso de uso.
2. Liste os **participantes** envolvidos (ator + componentes).
3. Transforme cada passo do fluxo em **mensagens** entre participantes.
4. Para cada exceção/validação, use `alt/opt` com o resultado esperado.
5. Finalize com o retorno de sucesso e as possíveis respostas de erro.

### Boas práticas

- Não desenhe “tudo do sistema” em um diagrama: escolha um cenário.
- Mantenha nomes de mensagens claros (ações/verbos).
- Valide e trate erros nos pontos onde realmente acontecem.
- Use `alt` para exceções importantes (não precisa de todas as micro-validações).
- O diagrama deve ser legível: se ficou gigante, divida em mais diagramas.
