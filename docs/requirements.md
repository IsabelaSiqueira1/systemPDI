## Documento de requisitos funcionais e não funcionais

O que é: um "contrato" do sistema que descreve o que ele precisa fazer(requisitos funcionais) e como ele deve se comportar em termos de qualiade e limitações (requisitos não funcionais)

Ele serve para alinhar o escopo, guiar o codigo e virar base de testes e checkpoints

"uma condição ou capacidade necessaria para um usuario resolver um problema ou alcançar um objetivo"

"Durante o desenvolvimento do projeto, quanto mais tarde um defeito nos requisitos for corrigido, mais altos serão os custos associados com a sua correção. Por exemplo, o esforço necessario para corrigiir um defeito durante a programação é ate 20 vezes maior do que realizar a correção durante a engenharia de requisitos. Se o defeito for corrigido durante os testes de aceitação, o esforço exigido pode ser ate 100 vezes maior" [BOEHM, 1981]

## Objetivo do Sistema

Criar o back-end de uma aplicação para gerenciamento de filas de atendimento (guichês), permitindo:
- cadastro e listagem de serviços (catálogo de serviços)
- registro de profissionais
- vínculo do profissional a um serviço (seleção de serviço ao iniciar expediente)
- controle de status do profissional (DISPONIVEL, OCUPADO, INDISPONIVEL)
- emissão de fichas por clientes com categoria de prioridade
- chamada do próximo cliente
- encerramento de atendimento
- envio de webhook quando um cliente for chamado
- controle de atendimentos inteiramente em memória volátil

## Funcionais -- o que o sistema faz

#### Serviços

- O sistema deve disponibilizar endpoint para listar serviços cadastrados.
- O sistema deve disponibilizar endpoint para criar um serviço.

#### Profissionais

- O sistema deve disponibilizar endpoint para cadastrar profissional
informando nome, email e senha.
- O email deve ser único no sistema e utilizado como identificador do profissional.
- O sistema deve permitir que um profissional inicie expediente selecionando um serviço existente (ativando vínculo/serviço atual).
- O sistema deve registrar o histórico de atuação do profissional,
armazenando quando iniciou e quando encerrou o expediente em determinado serviço,
independentemente da existência de atendimentos realizados.
- O sistema deve permitir que um profissional encerre expediente (ficando INDISPONIVEL).
- O sistema deve permitir que um profissional se vincule a um serviço existente (selecionar serviço ao iniciar expediente).
- O sistema deve permitir que o profissional altere seu status entre:
  - DISPONIVEL (no expediente e apto a atender)
  - INDISPONIVEL (fora do expediente / não participa da fila)
  - OCUPADO (durante um atendimento)
- O sistema deve colocar o profissional em OCUPADO ao chamar o próximo cliente e voltar para DISPONIVEL ao encerrar atendimento.
-O sistema deve impedir que um profissional altere de serviço (ou inicie expediente em outro serviço) enquanto estiver OCUPADO.
- O sistema deve impedir cadastro de profissionais com email já existente.

#### Fichas e Fila

- O sistema deve permitir emitir ficha informando serviço, cliente e categoria de prioridade.
- O sistema deve ordenar o atendimento por categoria de prioridade (Priority > Medium > Low) e, dentro da mesma categoria, por ordem de chegada (FIFO).

#### Atendimento

- O sistema deve fornecer endpoint para que um profissional chame o próximo cliente da fila.
- O sistema deve fornecer endpoint para que um profissional encerre o atendimento atual.
- O sistema deve impedir chamada de próximo cliente quando:
  - o profissional não estiver vinculado ao serviço,
  - o profissional não estiver DISPONIVEL,
  - a fila estiver vazia.

#### Webhook

- O sistema deve notificar via webhook quando uma ficha for chamada, enviando os dados do atendimento.
- O endpoint externo para o webhook deve ser configurável via API.
- A falha no envio do webhook não deve impedir a criação do atendimento (deve apenas registrar log/erro).

## Não Funcionais -- como o sistema deve ser

- O sistema deve ser desenvolvido utilizando Go.
- O sistema deve persistir dados na memoria volatil (não usar banco de dados).
- O sistema deve utilizar estruturas adequadas de filas e listas em memória.
- Toda comunicação deve acontecer via API REST.
- Todos os endpoints devem ser documentados utilizando Swagger/OpenAPI.
- O sistema deve ser executável localmente com Docker/Docker Compose.
- O endpoint externo para o webhook deve ser configurável.
- O sistema deve ter um tempo limite de resposta da requisição.
- O sistema deve garantir consistência do estado em operações concorrentes (ex.: duas chamadas simultâneas para chamar próximo cliente).

## Regra de negócios

#### Sobre vínculo do profissional ao serviço

- Um profissional deve selecionar/vincular-se a um serviço para poder atender.
- Um profissional só pode atuar no serviço ao qual está vinculado.
- Um profissional só pode estar vinculado a no máximo 1 serviço por vez.
- Para trocar de serviço, o profissional deve encerrar o expediente (ficar INDISPONIVEL) e então iniciar expediente novamente escolhendo outro serviço.

#### Sobre status do profissional

- Só é possível chamar o próximo cliente se o profissional estiver DISPONIVEL.
- Um profissional OCUPADO não pode chamar o próximo cliente.
- Um profissional INDISPONIVEL não participa do atendimento e não pode chamar o próximo cliente.
- Um profissional só pode ficar INDISPONIVEL se não estiver OCUPADO.
- Ao chamar o próximo cliente:
  - o sistema cria um atendimento com inicioEm,
  - e altera o status do profissional para OCUPADO.
- Ao encerrar atendimento:
  - o sistema registra fimEm,
  - e altera o status do profissional para DISPONIVEL.
- Um profissional não pode encerrar expediente enquanto existirem clientes aguardando atendimento na fila do serviço.

#### Sobre fila e fichas

- Categorias válidas: Priority, Medium, Low.
- A fila respeita prioridade (Priority > Medium > Low) sem furar, e FIFO dentro da categoria.
- Cada ficha deve possuir um ID único por serviço.
- Se não houver clientes na fila, o sistema não deve permitir chamada.
