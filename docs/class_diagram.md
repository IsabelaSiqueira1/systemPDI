## Diagrama de classe

Draw.io: https://drive.google.com/file/d/1UR5ViAbHcCuxr4jYljAVzmnJzwghlo9q/view?usp=sharing

Um diagrama de classe mostra as estruturas principais do sistema:
 - classes(entidades, objetos, modelos)
 - atributos
 - metodos relevantes
 - relacionamento entre classes

## 1. Service
Atributos
 - id: UUID
 - name: String
 - queue: Queue FIFO
 - webhookUrl: String

Regras
 - Gera tickets
 - Mantém ordem FIFO
 - Notifica via webhook quando necessário

## 2. Ticket
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
Profissional que chama a próxima pessoa da fila.

Atributos
 - id: UUID
 - name: String
 - status: ex: AVAILABLE, BUSY, OFFLINE

Regras
 - Apenas chama tickets quando estiver AVAILABLE
 - Após chamar, vira BUSY até concluir o atendimento


## 4. Attendance
Registro do atendimento de um ticket por um profissional.

Atributos
 - id: UUID
 - professionalId: UUID
 - ticketId: UUID
 - startedAt: DateTime
 - finishedAt: DateTime

Regras
 - Um ticket gera no máximo 1 atendimento
 - Um profissional pode ter nenhum ou vários atendimentos ao longo do dia


## Relacionamentos

Service → Ticket
 1 Service possui muitos Tickets

Multiplicidade:
 Service 1 —— 0..* Ticket

---------------
Ticket → Attendance
 1 Ticket pode ter 0 ou 1 Attendance

Multiplicidade:
 Ticket 1 —— 0..1 Attendance 

--------------
Professional → Attendance
 1 Professional pode ter 0..* Attendances

Multiplicidade:
 Professional 1 —— 0..* Attendance 

