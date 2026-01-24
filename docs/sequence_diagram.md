## Diagrama de sequencia

draw.io: https://app.diagrams.net/#G1UR5ViAbHcCuxr4jYljAVzmnJzwghlo9q#%7B%22pageId%22%3A%22ZbWJfetMCKG_KI7-XJCB%22%7D

### Diagrama de sequencia - Chamar proximo Cliente()

```mermaid

sequenceDiagram
    actor Prof as Profissional
    participant API as "API (AtendimentosController)"
    participant SVC as "AtendimentoService"
    participant Repo as "ServicoRepo (memoria)"
    participant Serv as "Servico"
    participant Fila as "Fila (do servico)"
    participant ARepo as "AtendimentoRepo (memoria)"
    participant WH as "Webhook"
    participant WHC as "WebhookConfig"
    participant Ext as "Sistema Externo"

    Prof->>API: POST /v1/servicos/{id}/chamar-proximo {idProfissional}
    API->>SVC: chamarProximo(idServico, idProfissional)

    SVC->>Repo: buscarServico(idServico)
    alt servico nao existe
        Repo-->>SVC: null
        SVC-->>API: erro 404 (servico nao encontrado)
        API-->>Prof: 404
    else servico existe
        Repo-->>SVC: Servico

        SVC->>Serv: buscarProfissional(idProfissional)
        alt profissional nao pertence ao servico
            Serv-->>SVC: erro
            SVC-->>API: erro 403 (nao permitido)
            API-->>Prof: 403
        else profissional ok
            Serv-->>SVC: Profissional (modelo)

            SVC->>Serv: statusDoProfissional(idProfissional)
            alt status != DISPONIVEL (OCUPADO ou INDISPONIVEL)
                Serv-->>SVC: status
                SVC-->>API: erro 409 (profissional indisponivel)
                API-->>Prof: 409
            else status == DISPONIVEL
                Serv-->>SVC: DISPONIVEL

                SVC->>Fila: chamarProximo() (Priority -> Medium -> Low)
                alt fila vazia
                    Fila-->>SVC: sem ficha
                    SVC-->>API: erro 404/409 (fila vazia)
                    API-->>Prof: 404/409
                else tem ficha
                    Fila-->>SVC: ficha

                    SVC->>ARepo: criarAtendimento(idServico,idProf,idFicha,inicioEm=agora)
                    ARepo-->>SVC: atendimento

                    SVC->>Serv: marcarProfissionalOcupado(idProfissional)

                    SVC->>WH: enviar(payload)
                    WH->>WHC: getURL()

                    alt URL nao configurada
                        WHC-->>WH: vazio
                        WH-->>SVC: nao envia (loga e segue)
                    else URL configurada
                        WHC-->>WH: url
                        WH->>Ext: POST url {payload}
                        alt webhook falhou (timeout/erro)
                            Ext-->>WH: erro
                            WH-->>SVC: falha (loga e segue)
                        else webhook ok
                            Ext-->>WH: 200 OK
                            WH-->>SVC: ok
                        end
                    end

                    SVC-->>API: atendimentoDTO
                    API-->>Prof: 200 OK + JSON
                end
            end
        end
    end
```

### Diagrama de sequencia - Encerrar Atendimento()

```mermaid

sequenceDiagram
actor Prof as Profissional
participant API as API (AtendimentosController)
participant SVC as AtendimentoService
participant Repo as ServicoRepo (memoria)
participant Serv as Servico
participant ARepo as AtendimentoRepo (memoria)

    Prof->>API: PUT /v1/servicos/{id}/encerrar-atendimento {idProfissional}
    API->>SVC: encerrarAtendimento(idServico, idProfissional)

    SVC->>Repo: buscarServico(idServico)
    alt servico nao existe
        Repo-->>SVC: null
        SVC-->>API: erro 404 (servico nao encontrado)
        API-->>Prof: 404
    else servico existe
        Repo-->>SVC: Servico

        SVC->>Serv: buscarProfissional(idProfissional)
        alt profissional nao pertence ao servico
            Serv-->>SVC: erro
            SVC-->>API: erro 403 (nao permitido)
            API-->>Prof: 403
        else profissional ok
            Serv-->>SVC: Profissional

            SVC->>Serv: statusDoProfissional(idProfissional)
            alt status != OCUPADO (DISPONIVEL ou INDISPONIVEL)
                Serv-->>SVC: status
                SVC-->>API: erro 409 (nao ha atendimento em andamento)
                API-->>Prof: 409
            else status == OCUPADO
                Serv-->>SVC: OCUPADO

                SVC->>ARepo: buscarAtendimentoEmAberto(idServico, idProfissional)
                alt nao encontrou atendimento em aberto
                    ARepo-->>SVC: null
                    SVC-->>API: erro 500/409 (inconsistencia)
                    API-->>Prof: 500/409
                else encontrou atendimento
                    ARepo-->>SVC: atendimento

                    SVC->>ARepo: encerrarAtendimento(atendimento.id, fimEm=agora)
                    ARepo-->>SVC: ok

                    SVC->>Serv: marcarProfissionalDisponivel(idProfissional)

                    SVC-->>API: 200 OK
                    API-->>Prof: 200 OK
                end
            end
        end
    end
```

### Diagrama de sequencia - Emitir Ficha()

```mermaid

sequenceDiagram
actor Cli as Cliente
participant API as API (FichasController)
participant SVC as FichaService
participant Repo as ServicoRepo (memoria)
participant Serv as Servico
participant Fila as Fila (do servico)

    Cli->>API: POST /v1/servicos/{id}/fichas {nomeCliente, categoria}
    API->>SVC: emitirFicha(idServico, nomeCliente, categoria)

    alt nomeCliente vazio
        SVC-->>API: erro 400 (nome obrigatorio)
        API-->>Cli: 400
    else nome ok
        alt categoria invalida
            SVC-->>API: erro 400 (categoria invalida)
            API-->>Cli: 400
        else categoria ok
            SVC->>Repo: buscarServico(idServico)
            alt servico nao existe
                Repo-->>SVC: null
                SVC-->>API: erro 404 (servico nao encontrado)
                API-->>Cli: 404
            else servico existe
                Repo-->>SVC: Servico

                SVC->>Serv: gerarFicha(nomeCliente, categoria, agora)
                Serv-->>SVC: ficha

                SVC->>Fila: adicionarFicha(ficha)
                Fila-->>SVC: ok

                SVC-->>API: fichaDTO
                API-->>Cli: 201 Created + JSON
            end
        end
    end

```

### Diagrama de sequencia - Listar Servicos()

```mermaid

sequenceDiagram
    actor User as Administrador/Cliente
    participant API as API (ServicosController)
    participant SVC as ServicoService
    participant Repo as ServicoRepo (memoria)

    User->>API: GET /v1/servicos
    API->>SVC: listarServicos()
    SVC->>Repo: buscarTodos()
    Repo-->>SVC: listaServicos
    SVC-->>API: listaServicosDTO
    API-->>User: 200 OK + JSON (pode ser lista vazia)
```

### Diagrama de sequencia - Criar Servico()

```mermaid

sequenceDiagram
    actor Admin as Administrador
    participant API as API (ServicosController)
    participant SVC as ServicoService
    participant Repo as ServicoRepo (memoria)

    Admin->>API: POST /v1/servicos {nome}
    API->>SVC: criarServico(nome)

    alt nome vazio ou so espacos
        SVC-->>API: erro 400 (nome invalido)
        API-->>Admin: 400
    else nome ok
        SVC->>Repo: existePorNome(nome)?
        alt ja existe
            Repo-->>SVC: true
            SVC-->>API: erro 409 (servico ja existe)
            API-->>Admin: 409
        else nao existe
            Repo-->>SVC: false
            SVC->>Repo: salvar(servicoNovo + fila)
            Repo-->>SVC: servicoCriado
            SVC-->>API: servicoDTO
            API-->>Admin: 201 Created + JSON
        end
    end
```

### Diagrama de sequencia - Configurar Webhook()

```mermaid

sequenceDiagram
    actor Admin as Administrador
    participant API as API (WebhookController)
    participant SVC as WebhookConfigService
    participant Repo as WebhookConfigRepo (memoria)

    Admin->>API: PUT /v1/webhook/config {url}
    API->>SVC: configurarURL(url)

    alt url vazia ou formato invalido
        SVC-->>API: erro 400 (url invalida)
        API-->>Admin: 400
    else url ok
        SVC->>Repo: salvarOuAtualizar(url)
        Repo-->>SVC: ok
        SVC-->>API: ok
        API-->>Admin: 200 OK
    end
```

### Diagrama de sequencia - Enviar Webhook()

```mermaid

sequenceDiagram
    participant Caller as AtendimentoService
    participant WH as Webhook
    participant ConfRepo as WebhookConfigRepo (memoria)
    participant Ext as Sistema Externo

    Caller->>WH: enviar(payload)
    WH->>ConfRepo: getURL()

    alt url nao configurada
        ConfRepo-->>WH: vazio/null
        WH-->>Caller: ok (nao envia, loga)
    else url configurada
        ConfRepo-->>WH: url
        WH->>Ext: POST url {payload}
        alt falhou (timeout/erro)
            Ext-->>WH: erro
            WH-->>Caller: falha (loga e segue)
        else sucesso
            Ext-->>WH: 200 OK
            WH-->>Caller: ok
        end
    end
```
