## Diagrama de classe

draw.io: https://app.diagrams.net/#G1UR5ViAbHcCuxr4jYljAVzmnJzwghlo9q#%7B%22pageId%22%3A%22SPg6ARsj7m4EUYAmXlnh%22%7D

´´´mermaid

classDiagram

class Servico {
+UUID id
+String name
+List~Profissional~ profissionais
}

class Fila {
+Queue~Ficha~ Priority
+Queue~Ficha~ Medium
+Queue~Ficha~ Low
+adicionaFicha(f: Ficha) void
+chamarProximo() Ficha
}

class Ficha {
+UUID id
+String clientName
+CategoriaAtendimento category
+DateTime emitidaEm
+UUID idService
}

class Profissional {
+UUID id
+String name
+UUID idService
+StatusProfissional status
+MarcarOcupado() void
+MarcarDisponivel() void
+MarcarIndisponivel() void
}

class Atendimento {
+UUID id
+UUID idProfissional
+UUID idFicha
+UUID idServico
+DateTime inicioEm
+DateTime fimEm
}

class Webhook {
+StatusWebhook status
+enviar(payload) ResultadoWebhook
}

class WebhookConfig {
+String url
}

class CategoriaAtendimento {
<<enumeration>>
Priority
Medium
Low
}

class StatusProfissional {
<<enumeration>>
DISPONIVEL
OCUPADO
INDISPONIVEL
}

class StatusWebhook {
<<enumeration>>
SUCESSO
ERRO
NAO_CONFIGURADO
}

class ResultadoWebhook {
<<enumeration>>
SUCESSO
FALHA
}

Servico "1" _-- "1" Fila : possui
Servico "1" o-- "0.._" Ficha : emite
Servico "1" o-- "0.._" Profissional : possui
Servico "1" o-- "0.._" Atendimento : registra

Profissional "1" o-- "0..\*" Atendimento : realiza
Ficha "1" o-- "0..1" Atendimento : gera

Webhook ..> WebhookConfig : usa
´´´

## Classes

- Serviço
- Fila
- Ficha
- Profissional
- Atendimento
- Webhook
- WebhookConfig

## 1. Serviço

Gerencia sua fila de fichas, profissionais vinculados e a operação de chamar o próximo cliente.

Atributos

- id: UUID
- name: String
- profissionais: List<Profissional>

Regras de Negócio associadas

- Um serviço possui exatamente 1 fila prioritária.
- As fichas emitidas para um serviço entram apenas na fila daquele serviço.
- O atendimento respeita prioridade: IMEDIATO > MÉDIA > BAIXA (sem furar), e FIFO dentro da mesma categoria.
- O profissional só pode chamar próximo cliente do serviço ao qual está vinculado.

## 2. Fila

Representa a estrutura que mantém as fichas em ordem de prioridade.

Atributos

- Priority: Queue<Ficha>
- Medium: Queue<Ficha>
- Low: Queue<Ficha>

Metodos

- adicionaFicha(f: Ficha) - Adiciona a ficha na fila correspondente à categoria.
- chamarProximo(): Ficha - Retorna e remove a próxima ficha respeitando:
  - imediata se não estiver vazia
  - senão media se não estiver vazia
  - senão baixa se não estiver vazia
  - se todas vazias, retorna erro/sem ficha

## 2. Ficha

Representa a solicitação de atendimento feita por um cliente.

Atributos

- id: UUID
- clientName: String
- category: CategoriaAtendimento (Priority, Medium, Low)
- emitidaEm: DateTime (Momento de emissão)
- idService: UUID (Serviço ao qual a ficha pertence)

Regras

- Cada ficha sempre pertence a um único serviço.
- Uma ficha pode gerar no máximo um atendimento.
- Ao ser chamada, a ficha é removida da fila.

## 3. Profissional

Representa o atendente que chama fichas e realiza os atendimentos.

Atributos

- id: UUID
- name: String
- idService: UUID (Serviço ao qual está vinculado)
- status: StatusProfissional (DISPONIVEL | OCUPADO | INDISPONIVEL)

Metodos

- MarcarOcupado() - Marca profissional como OCUPADO
- MarcarDisponivel() - marcar profissional como DISPONIVEL
- MarcarIndisponivel() - Marca profissional como INDISPONIVEL

Regras

- Só pode chamar próximo cliente se estiver DISPONIVEL.
- Se estiver OCUPADO, deve encerrar o atendimento antes de chamar outro.
- Se estiver INDISPONIVEL, não pode chamar próximo nem encerrar atendimento.
- Só atua no serviço ao qual está vinculado.

## 4. Atendimento

Representa o registro oficial de um atendimento realizado.

Atributos

- id: UUID (Identificador do atendimento)
- idProfissional: UUID (Profissional que realizou a chamada)
- idFicha: UUID (Ficha atendida)
- idServico: UUID (Serviço onde ocorreu)
- inicioEm: DateTime (Início da chamada)
- fimEm: DateTime (Fim da chamada)

Regras

- Um atendimento está sempre relacionado a exatamente 1 ficha, 1 profissional e 1 serviço.
- Uma ficha deve gerar, no máximo, 1 atendimento.

## Webhook

atributos

- status: StatusWebhook (SUCESSO | ERRO | NAO_CONFIGURADO)

metodos

- enviar(payload): ResultadoWebhook

## WebhookConfig

Atributos

- url: string

## Relacionamentos

Serviço —— Fila

Cada serviço possui exatamente 1 fila prioritária
Service 1 —— 1 Fila

---

Serviço —— Ficha

Um serviço possui muitas fichas
Service 1 —— 0..\* Ficha

---

Serviço —— Profissional

Um serviço possui muitos profissionais
Service 1 —— 0..\* Profissional

---

Profissional —— Atendimento
Um profissional pode realizar muitos atendimentos

---

Profissional 1 —— 0..\* Atendimento

---

Ficha —— Atendimento

---

Ficha 1 —— 0..1 Atendimento

---

Serviço —— Atendimento

---

Service 1 —— 0..\* Atendimento

---
