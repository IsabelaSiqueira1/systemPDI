## Diagrama de classe

5 classes principais foram identificadas para o sistema de gerenciamento de filas de atendimento:
 - Serviço
 - Ficha
 - Profissional
 - Atendimento
 - ConfiguraçãoWebhook



## 1. Service
Representa um guichê/serviço de atendimento.
Gerencia sua fila de fichas, profissionais ativos e a operação de chamar próximo cliente.

Atributos
 - id: UUID
 - name: String
 - fila: lista de fichas (Armazena fichas em ordem de prioridade)
 - profissionais ativos: lista de profissionais (Profissionais disponíveis para chamar fichas)

metodos
 - adicionarFicha(f: Ficha)  - Adiciona ficha na fila, respeitando prioridade
 - chamarProximo(): Ficha  - Seleciona a próxima ficha conforme regras de prioridade
 - ativarProfissional(p: Profissional) - Define profissional como ativo
 - desativarProfissional(p: Profissional) - Remove profissional da lista de ativos / marca como inativo

Regras de Negócio associadas
 - Um serviço possui sua própria fila.
 - Só é permitido chamar próximo cliente se houver profissional ativo.
 - A fila deve respeitar ordem: Imediato > Média > Baixa.
 - Ao chamar um cliente, o serviço registra o início de um atendimento.


## 2. Ficha
Representa a solicitação de atendimento feita por um cliente.

Atributos
 - id: UUID
 - clientName: String
 - category: CategoriaAtendimento (Prioridade Imediato, Média, Baixa)
 - emitidaEm: DateTime (Momento de emissão)
 - idService: UUID (Serviço ao qual a ficha pertence)

Regras
 - Cada ficha sempre pertence a um único serviço.
 - Uma ficha pode gerar no máximo um atendimento.
 - Fichas são consumidas ao serem chamadas.


## 3. Profissional
Representa o atendente que chama fichas e realiza os atendimentos.

Atributos
 - id: UUID
 - name: String
 - idService: UUID (Serviço ao qual está vinculado)
 - status: StatusProfissional (ex: AVAILABLE, BUSY, OFFLINE)

Metodos
 - iniciarAtendimento() - Marca profissional como BUSY
 - encerrarAtendimento() - Marca profissional como AVAILABLE
 - ativar() - Fica ativo no serviço
 - desativar() - Fica inativo no serviço

Regras
 - Só pode chamar ficha se estiver AVAILABLE.
 - Só pode chamar ficha do serviço ao qual está vinculado.
 - Pode ter vários atendimentos ao longo do dia.


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


## WebhookConfig

Atributos
 - url: string

metodos
 - enviar(atendimento: Atendimento) - Envia notificação do atendimento via POST



## Relacionamentos

Service —— Fichas
 1 Service possui muitas Fichas
 cada ficha pertence a exatamente um unico Service

multiplicidade:
 Service 1 —— 0..* Ficha

----------

Service ——   Profissionais
 1 Service possui muitos Profissionais
 1 Profissional pertence a exatamente 1 Serviço

multiplicidade:
 Service 1 —— 0..* Profissional

----------

Profissional —— Atendimento
 1 Profissional pode ter zero, um ou muitos Atendimentos
 Cada Atendimento está ligado a 1 Profissional

multiplicidade:
 Profissional 1 —— 0..* Atendimento

----------

Ficha —— Atendimento
 1 Ficha pode gerar zero ou um Atendimento
 Cada Atendimento está ligado a 1 Ficha

multiplicidade:
 Ficha 1 —— 0..1 Atendimento

----------

serviço —— Atendimento
 1 Serviço pode ter zero, um ou muitos Atendimentos
 1 Atendimento ocorre em exatamente 1 Serviço

multiplicidade:
 Service 1 —— 0..* Atendimento

----------