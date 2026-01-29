## Diagrama Entidade-Relacionamento (ER)

Se eu fosse salvar isso num banco de dados, quais seriam as tabelas, quais colunas elas teriam e como elas se relacionam?

entidade vira tabela
atibutos vira colunas
relacionamentos vira chaves estrangeiras(FK)

PK (Primary Key): identificador único da tabela
FK (Foreign Key): campo que aponta para outra tabela

### Tabelas

#### services

- id (PK)
- name (unique)

#### profissionais

- id (PK)
- name
- service_id (FK → services.id)
- status (DISPONIVEL | OCUPADO | INDISPONIVEL)

#### fichas

- id (PK)
- service_id (FK → services.id)
- client_name
- category (Priority | Medium | Low)
- emitido_em

#### atendimentos

- id (PK)
- service_id (FK → services.id)
- professional_id (FK → profissionais.id)
- ficha_id (FK → fichas.id)
- começou_em
- terminou_em

#### webhook_config

- id (PK)
- url
- updated_at

Relacionamentos
services 1 — N profissionais
services 1 — N tickets
services 1 — N atendimentos
profissionais 1 — N atendimentos
fichas 1 — 0..1 atendimentos
