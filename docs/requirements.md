## Documento de requisitos funcionais e não funcionais

O que é: um "contrato" do sistema que descreve o que ele precisa fazer(requisitos funcionais) e como ele deve se comportar em termos de qualiade e limitações (requisitos não funcionais)

Ele serve para alinhar o escopo, guiar o codigo e virar base de testes e checkpoints

"uma condição ou capacidade necessaria para um usuario resolver um problema ou alcançar um objetivo"

"Durante o desenvolvimento do projeto, quanto mais tarde um defeito nos requisitos for corrigido, mais altos serão os custos associados com a sua correção. Por exemplo, o esforço necessario para corrigiir um defeito durante a programação é ate 20 vezes maior do que realizar a correção durante a engenharia de requisitos. Se o defeito for corrigido durante os testes de aceitação, o esforço exigido pode ser ate 100 vezes maior" [BOEHM, 1981]

## Objetivo do Sistema

Criar o back-end de uma aplicação para gerenciamento de filas de atendimento (guichês), permitindo:

- cadastro e listagem de serviços
- registro de profissionais vinculados a um serviço
- vínculo do profissional a um serviço (seleção de serviço ao iniciar expediente)
- controle de status do profissional (DISPONIVEL, OCUPADO, INDISPONIVEL)
- emissão de fichas por clientes com categoria de prioridade
- chamada do próximo cliente
- envio de webhook quando um cliente for chamado
- controle de atendimentos inteiramente em memória volátil

## Funcionais -- o que o sistema faz

#### Serviços

- O sistema deve disponibilizar endpoint para listar serviços cadastrados.
- O sistema deve disponibilizar endpoint para criar um serviço.

#### Profissionais

- O sistema deve disponibilizar endpoint para cadastrar profissional.
- O sistema deve permitir que um profissional se vincule a um serviço existente (selecionar serviço ao iniciar expediente).
- O sistema deve permitir que o profissional altere seu status entre:
  - DISPONIVEL (no expediente e apto a atender)
  - INDISPONIVEL (fora do expediente / não participa da fila)
  - OCUPADO (durante um atendimento)
- O sistema deve colocar o profissional em OCUPADO ao chamar o próximo cliente e voltar para DISPONIVEL ao encerrar atendimento.

#### Fichas e Fila

- O sistema deve permitir emitir ficha informando serviço, cliente, categoria e grau de prioridade.
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

## Não Funcionais -- como o sistema deve ser

- O sistema deve ser desenvolvido utilizando Go.
- O sistema deve persistir dados na memoria volatil (não usar banco de dados).
- O sistema deve utilizar estruturas adequadas de filas e listas em memória.
- Toda comunicação deve acontecer via API REST.
- Todos os endpoints devem ser documentados utilizando Swagger/OpenAPI.
- O sistema deve ser executável localmente com Docker/Docker Compose.
- O endpoint externo para o webhook deve ser configurável.
- O sistema deve ter um tempo limite de resposta da requisição.

## Regra de negócios

#### Sobre vínculo do profissional ao serviço

- Um profissional deve selecionar/vincular-se a um serviço para poder atender (e para poder chamar próximo cliente).
- Um profissional só pode atuar no serviço ao qual está vinculado.

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

#### Sobre fila e fichas

- Categorias válidas: Priority, Medium, Low.
- A fila respeita prioridade (Priority > Medium > Low) sem furar, e FIFO dentro da categoria.
- Cada ficha deve possuir um ID único por serviço.
- Se não houver clientes na fila, o sistema não deve permitir chamada.
