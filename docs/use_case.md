Caso de uso

Listar Serviços 
Ator: Administrador/cliente
Objetivo: Retornar todos os serviços cadastrados.

Pré-condição: Nenhuma.

Fluxo Principal:3
 1. O ator solicita a listagem dos serviços.
 2. O sistema recupera todos os serviços da memória.
 3. O sistema retorna a lista.

exceções:
 - Não ha serviços disponivel: retornar a mensagem que não ha serviço disponivel no momento
 - falha na requisição: o sistema retorna erro.
 - Falha interna ao recuperar serviços: o sistema retorna erro e loga.
 - tempo de resposta alta: o sistema retorna timeout e loga.

------------------
Criar Serviço
Ator: Administrador
Objetivo: Cadastrar um novo serviço que poderá ser atendido pelos profissionais.    

pré-condição:
– O nome do serviço deve ser informado.

Fluxo Principal:
 1. O usuário solicita a criação de um novo serviço.
 2. O sistema valida se o nome do serviço foi informado.
 3. O sistema cria o serviço na memória.
 4. O sistema retorna o serviço criado com seu ID.

exceções:
 - Nome não informado: O sistema rejeita a operação e retorna erro.
 - Nome invalido(vazio, com espaços): o sistema retorna erro "nome invalido".
 - serviço existente: o sistema deve identificar caso haja um serviço duplicado e informar e retornar o problema.
 - Falha interna ao criar o serviço: o sistema retorna erro generico e registra o log.

--------------------
Ativar Atendimento
Ator: Profissional
Objetivo: O profissional informa que está disponível para atender um serviço.

pré-condição:
– O profissional deve estar vinculado a um serviço existente.

Fluxo Principal:
 1. O profissional solicita ativação no serviço escolhido.
 2. O sistema valida se o serviço existe.
 3. O sistema marca o profissional como ativo naquele serviço.
 4. O sistema confirma a ativação.

exceções:
 - Serviço inexistente: O sistema retorna erro.
 - Profissional não vinculado ao serviço: o sistema bloqueia a operação.
 - Profissional ja está ativo: O sistema retorna aviso.
 - Erro interno ao alterar o estado: o sistema registra log e retorna erro. 

--------------------
Desativar Atendimento
Ator: Profissional
Objetivo: Encerrar o período de atendimento do profissional.

pré-condição: nenhuma

Fluxo Principal:
 1. O profissional solicita desativação do serviço.
 2. O sistema valida que o profissional está ativo.
 3. O sistema marca o profissional como inativo.
 4. O sistema confirma a operação.

exceções:
 - Profissional não ativo: O sistema retorna erro.
 - Profissional não vinculado ao serviço: o sistema bloqueia a operação.
 - Profissional tentando desativar outro profissional(ID errado enviado): o sistema rejeita a operação. 
 - Erro interno ao alterar o estado: o sistema registra log e retorna erro. 

-------------------
Emitir Ficha
Ator: Cliente
Objetivo: Gerar uma ficha de atendimento com prioridade.

Pré-condições:
– O serviço deve existir.
– A categoria deve ser válida (imediato, média, baixa).

Fluxo Principal:
 1. O cliente solicita a emissão da ficha informando:
    – nome do cliente
    – serviço
    – categoria
2. O sistema valida os dados.
3. O sistema gera uma nova ficha com ID único por serviço.
4. O sistema registra o momento de emissão.
5. O sistema adiciona a ficha na fila correta do serviço.
6. O sistema retorna a ficha emitida.

exceções:
 - Nome do cliente não informado: o sistema retorna erro.
 - Categoria inválida: O sistema retorna erro.
 - Serviço inexistente: O sistema retorna erro.
 - Serviço existe mas ainda nao tem nenhum profissional ativo: o sistema retorna que não ha profissionais disponiveis.
 - Fila do serviço corrompida(erro na memoria):o sistema registra o erro e loga. 

------------------
Chamar Próximo Cliente
Ator: Profissional
Objetivo: Chama o próximo cliente da fila, respeitando prioridade e ordem.

Pré-condições:
 - Deve existir pelo menos 1 profissional ativo no serviço.
 - Deve existir cliente na fila.

Fluxo Principal:
 1. O profissional solicita chamar o próximo cliente.
 2. O sistema verifica se há profissionais ativos no serviço.
 3. O sistema verifica se existe cliente na fila.
 4. O sistema seleciona a ficha com maior prioridade.
 5. O sistema registra o início do novo atendimento.
 6. O sistema registra o fim do atendimento anterior (se houver).
 7. O sistema envia um webhook com os dados da chamada.
 8. O sistema retorna o cliente chamado.

exceções:
 - nenhum profissionais ativos: O sistema retorna erro.
 - Fila vazia: O sistema retorna erro.
 - Profissional não pertence ao serviço: o sistema retorna erro.
 - Profissional tenta chamar a ficha  de outro serviço(ID incorreto): o sistema bloqueia a operação.
 - Falha na geração do registro de atendimento: o sistema retorna erro e registra o log.
 - webhook falhou: O sistema registra a falha, mas prossegue com a chamada.
 - payload do webhook invalido: O sistema registra erro interno.
 - Dois profissionais chamam ao mesmo tempo: o sistema deve garantir controle.

-----------------
Enviar Webhook
Ator: Sistema externo
Objetivo: Enviar os dados da chamada para um endpoint configurado.

pre-condição: nenhuma

Fluxo Principal:
 1. Quando um cliente é chamado, o sistema prepara o payload do webhook.
 2. O sistema envia uma requisição HTTP para o endpoint configurado.
 3. O sistema registra que o webhook foi disparado.

exceções:
 - Endpoint fora do ar: O sistema registra o erro no log.
 - Endpoint recusou: o sistema registra erro.
 - URL do webhook não configurada: o sistema registra erro e não envia nada.
 - Payload invalido: o sistema registra erro.
 - falha na rede: o sistema registra erro.
 - a resposta do endpoint demorou demais: o sistema registra timeout e loga.