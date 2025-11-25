Documento de requisitos funcionais e não funcionais

O que é: um "contrato" do sistema que descreve o que ele precisa fazer(requisitos funcionais) e como ele deve se comportar em termos de qualiade e limitações (requisitos não funcionais)

Ele serve para alinhar o escopo, guiar o codigo e virar base de testes e checkpoints

"uma condição ou capacidade necessaria para um usuario resolver um problema ou alcançar um objetivo"

"Durante o desenvolvimento do projeto, quanto mais tarde um defeito nos requisitos for corrigido, mais  altos serão os custos associados com a sua correção. Por exemplo, o esforço necessario para corrigiir um defeito durante a programação é ate 20 vezes maior do que realizar a correção durante a engenharia de requisitos. Se o defeito for corrigido durante os testes de aceitação, o esforço exigido pode ser ate 100 vezes maior" [BOEHM, 1981]


Objetivo do Sistema
Criar o back-end de uma aplicação para gerenciamento de filas de atendimento (guichês), permitindo:
 - cadastro e listagem de serviços
 - registro de profissionais vinculados a um serviço
 - controle de profissionais ativos/inativos
 - emissão de fichas por clientes com categoria de prioridade
 - chamada do próximo cliente
 - envio de webhook quando um cliente for chamado
 - controle de atendimentos inteiramente em memória volátil


Funcionais(RF) -- o que o sistema faz
 - O sistema deve disponibilizar endpoint para Listar serviços cadastrados.
 - O sistema deve disponibilizar endpoint para Criar um serviço.
 - O sistema deve permitir que o profissional inicie ou encerre seu atendimento em um serviço (status ativo/inativo)
 - O sistema deve permitir emitir ficha de serviço informando serviço, cliente, categoria e grau de prioridade
 - O sistema deve ordenar o atendimento por categoria de prioridade (Imediato > prioridade media > prioridade baixa) e, dentro de cada categoria, por ordem de chegada (FIFO).
 - O sistema deve fornecer um endpoint para que um profissional chame o próximo cliente da fila.
 - O sistema deve notificar o proximo da fila com os dados do atendimento
 - O sistema deve encerrar atendimento caso não haja profissionais ativos
 - Ao chamar o proximo, o sistema deve registrar o fim do atendimento anterior e marcar o inicio de um novo



Não Funcionais(RNF) -- como o sistema deve ser
 - O sistema deve ser desenvolvido utilizando Go.
 - O sistema deve persistir dados na memoria volatil (não usar banco de dados).
 - O sistema deve utilizar estruturas adequadas de filas e listas em memória.
 - Toda comunicação deve acontecer via API REST.
 - Todos os endpoints devem ser documentados utilizando Swagger/OpenAPI.
 - O sistema deve ser executável localmente com Docker/Docker Compose.
 - O endpoint externo para o webhook deve ser configurável.


Regra de negócios
 - Só é possível chamar o próximo cliente se existir pelo menos um profissional ativo no serviço.
 - categorias validas (Imediato, prioridade media, prioridade baixa)
 - Cada ficha deve possuir um ID único por serviço.
 - Se não houver clientes na fila, o sistema não deve permitir chamada.
 - O profissional só pode realizar chamadas no serviço ao qual está vinculado.
