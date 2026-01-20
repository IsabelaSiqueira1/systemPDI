## Caso de uso

draw.io: https://app.diagrams.net/#G1UR5ViAbHcCuxr4jYljAVzmnJzwghlo9q#%7B%22pageId%22%3A%227V9Oo_qkh-ZA_1hTpyhE%22%7D

## Listar Serviços

Ator: Administrador/cliente
Objetivo: Retornar todos os serviços cadastrados.

Pré-condição: Nenhuma.

Fluxo Principal:

1.  O ator solicita a listagem dos serviços.
2.  O sistema recupera todos os serviços da memória.
3.  O sistema retorna a lista.

exceções:

- Não ha serviços disponivel: retornar lista vazia
- Erro de comunicação/requisição inválida: retorna erro.
- Falha interna ao recuperar serviços: o sistema retorna erro e loga.
- tempo de resposta alta: o sistema retorna timeout e loga.

---

## Criar Serviço

Ator: Administrador
Objetivo: Cadastrar um novo serviço que poderá ser atendido pelos profissionais.

pré-condição:

- O nome do serviço deve ser informado.

Fluxo Principal:

1.  O usuário solicita a criação de um novo serviço.
2.  O sistema valida se o nome do serviço foi informado.
3.  O sistema cria o serviço na memória.
4.  O sistema retorna o serviço criado com seu ID.

exceções:

- Nome não informado: O sistema rejeita a operação e retorna erro.
- Nome invalido(vazio, com espaços): o sistema retorna erro "nome invalido".
- serviço existente: o sistema deve identificar caso haja um serviço duplicado e informar e retornar o problema.
- Falha interna ao criar o serviço: o sistema retorna erro generico e registra o log.

---

## Encerrar Atendimento

Ator: Profissional
Objetivo: Encerrar o período de atendimento do profissional.

pré-condição:

- O serviço deve existir.
- O profissional deve estar vinculado a um serviço existente.
- O profissional deve estar com status OCUPADO.

Fluxo Principal:

1. o Profissional solicita encerrar o atendimento atual.
2. O sistema valida se o serviço existe.
3. O sistema valida se o profissional pertence ao serviço.
4. O sistema valida se o profissional está OCUPADO.
5. O sistema registra o fim do atendimento atual.
6. O sistema altera o status do profissional para DISPONIVEL.
7. O sistema confirma a operação.

exceções:

- Serviço inexistente: o sistema retorna erro.
- Profissional não vinculado ao serviço: o sistema bloqueia a operação.
- Profissional não está OCUPADO: o sistema retorna erro.
- Profissional tentando encerrar atendimento de outro profissional (ID incorreto enviado): o sistema rejeita a operação.
- Erro interno ao alterar o estado: o sistema registra log e retorna erro.

---

## Emitir Ficha

Ator: Cliente
Objetivo: Gerar uma ficha de atendimento com prioridade.

Pré-condições:

- O serviço deve existir.
- A categoria deve ser válida (imediato, média, baixa).
- O nome do cliente deve ser informado.

Fluxo Principal:

1.  O cliente solicita a emissão da ficha informando:
    – nome do cliente
    – serviço
    – categoria
2.  O sistema valida os dados.
3.  O sistema gera uma nova ficha com ID único por serviço.
4.  O sistema registra o momento de emissão.
5.  O sistema adiciona a ficha na fila correta do serviço.
6.  O sistema retorna a ficha emitida.

exceções:

- Nome do cliente não informado: o sistema retorna erro.
- Categoria inválida: O sistema retorna erro.
- Serviço inexistente: O sistema retorna erro.
- Serviço existe mas não há profissional DISPONIVEL: O sistema emite a ficha normalmente e pode retornar um aviso informativo.
- Fila do serviço corrompida(erro na memoria):o sistema registra o erro e loga.

---

## Chamar Próximo Cliente

Ator: Profissional
Objetivo: Chama o próximo cliente da fila, respeitando prioridade e ordem.

Pré-condições:

- O serviço deve existir.
- O profissional deve estar vinculado ao serviço existente.
- O profissional deve estar com status DISPONIVEL.
- Deve existir cliente na fila.

Fluxo Principal:

1.  O profissional solicita chamar o próximo cliente.
2.  O sistema valida se o serviço existe.
3.  O sistema valida se o profissional pertence ao serviço.
4.  O sistema valida se o profissional está DISPONIVEL.
5.  O sistema valida que a fila não está vazia.
6.  O sistema seleciona a próxima ficha respeitando a prioridade.
7.  O sistema cria um Atendimento (registra inicioEm, associa serviço/profissional/ficha).
8.  O sistema muda o status do profissional para OCUPADO.
9.  O sistema envia o webhook.
10. O sistema retorna os dados do atendimento.

exceções:

- Serviço inexistente: o sistema retorna erro.
- Profissional não pertence ao serviço: o sistema retorna erro/bloqueia.
- Profissional está INDISPONIVEL: o sistema retorna erro.
- Profissional está OCUPADO: o sistema retorna erro (“encerre o atendimento atual antes de chamar outro”).
- Fila vazia: o sistema retorna erro.
- Falha na geração do registro de atendimento: retorna erro e loga.
- Webhook falhou: o sistema registra a falha, mas prossegue com a chamada.
- URL do webhook não configurada: o sistema registra erro e não envia (mas prossegue).
- Dois profissionais chamam ao mesmo tempo: o sistema deve garantir controle.

---

## Enviar Webhook

Ator: Sistema interno
Objetivo: Enviar os dados da chamada para um endpoint configurado.

pre-condição: nenhuma

Fluxo Principal:

1.  Quando um cliente é chamado, o sistema prepara o payload do webhook.
2.  O sistema envia uma requisição HTTP para o endpoint configurado.
3.  O sistema registra que o webhook foi disparado.

exceções:

- Endpoint fora do ar: O sistema registra o erro no log.
- Endpoint recusou: o sistema registra erro.
- URL do webhook não configurada: o sistema registra erro e não envia nada.
- Payload invalido: o sistema registra erro.
- falha na rede: o sistema registra erro.
- a resposta do endpoint demorou demais: o sistema registra timeout e loga.

## Configurar a URL do Weebhook

Ator:Administrador
Objetivo: Configurar o endpoint para onde os webhooks serão enviados.

Fluxo principal:

1.  O administrador acessa a configuração do sistema.
2.  O administrador informa a URL do webhook.
3.  O sistema valida o formato da URL.
4.  O sistema salva a configuração.
