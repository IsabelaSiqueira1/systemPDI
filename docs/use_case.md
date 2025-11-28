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
- serviço existente: o sistema deve identificar caso haja um serviço duplicado e informar e retornar o problema

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
- Categoria inválida: O sistema retorna erro.
- Serviço inexistente: O sistema retorna erro.

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
 - Sem profissionais ativos: O sistema retorna erro.
 - Fila vazia: O sistema retorna erro.
 - webhook falhou: O sistema registra a falha, mas prossegue com a chamada.

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