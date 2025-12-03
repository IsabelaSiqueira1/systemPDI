## Diagrama de classe

Draw.io: https://drive.google.com/file/d/1UR5ViAbHcCuxr4jYljAVzmnJzwghlo9q/view?usp=sharing

Um diagrama de classe mostra as estruturas principais do sistema:
 - classes(entidades, objetos, modelos)
 - atributos
 - metodos relevantes
 - relacionamento entre classes

## 1. Service
Representa cada serviço disponível no sistema (ex.: “Atendimento Geral”, “Financeiro”, etc.)

Contém sua própria fila
Contém os profissionais ativos naquele serviço

Atributos
 - id: UUID
 - name: String
 - fila: lista de fichas
 - profissionais ativos: lista de profissionais

metodos
 - adicionarFicha(f: Ficha)
 - chamarProximo(): Ficha
 - ativarProfissional(p: Profissional)
 - desativarProfissional(p: Profissional)

## 2. Ficha
Representa a ficha retirada pelo cliente.

Atributos
 - id: UUID
 - clientName: String
 - category: Category
 - issuedAt: DateTime

Regras
 - Sempre pertence a um único Service
 - Só pode estar em uma fila por vez
 - É consumido no atendimento


## 3. Professional
Profissional que atende um único serviço por vez.

Atributos
 - id: UUID
 - name: String
 - idService: UUID
 - status: ex: AVAILABLE, BUSY, OFFLINE

Metodos
 - iniciarAtendimento()
 - encerrarAtendimento()

Regras
 - Apenas chama tickets quando estiver AVAILABLE
 - Após chamar, vira BUSY até concluir o atendimento


## 4. Attendance
Representa uma chamada efetivada.

Atributos
 - id: UUID
 - professionalId: UUID
 - fichaID: UUID
 - serviceId: UUID
 - startedAt: DateTime
 - finishedAt: DateTime

Regras
 - Um ticket gera no máximo 1 atendimento
 - Um profissional pode ter nenhum ou vários atendimentos ao longo do dia


## WebhookConfig

Atributos
 - url: string

metodos
 - enviar(atendimento: Atendimento)



## Relacionamentos

Service 1 → N Fichas
 1 Service possui muitas Fichas

Multiplicidade:
 Service 1 —— 0..* Ficha

A fila é uma lista de fichas dentro do serviço.

----------

Service 1 → N Profissionais
 1 Service possui muitos Profissionais

Multiplicidade:
 Service 1 —— 0..* Profissional

Mas apenas profissionais ativos podem chamar clientes.

----------

Profissional 1 → 1 Serviço

multiplicidade:
 Profissional 1 —— 1 Service

Um profissional só atende um único serviço por vez.

----------

Atendimento 1 → 1 Ficha

----------

Atendimento 1 → 1 Profissional

----------

Atendimento 1 → 1 Serviço

----------