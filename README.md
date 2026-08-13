# OpsCenter API

API em Go para estudo e evolução de uma plataforma de gestão de incidentes e operações.

> **Status:** projeto em desenvolvimento. O repositório contém a estrutura inicial da aplicação e ainda não representa um sistema pronto para produção.

## Objetivo

Explorar, de forma progressiva, fundamentos de backend aplicados a um domínio operacional:

- organização em camadas;
- modelagem de incidentes e fluxos;
- separação de responsabilidades;
- persistência de dados;
- validação e tratamento de erros;
- testes automatizados;
- observabilidade e auditoria.

## Estrutura atual

O projeto está organizado nos seguintes pacotes:

- `cmd`: inicialização da aplicação;
- `config`: configuração;
- `controller`: entrada e saída HTTP;
- `service`: regras de negócio;
- `repository`: acesso a dados;
- `entity`: entidades do domínio;
- `DTO`: objetos de transferência de dados.

## Tecnologias

- Go
- Go modules

## Próximas etapas

- consolidar os casos de uso do domínio;
- definir e documentar os contratos da API;
- implementar persistência e migrations;
- ampliar testes unitários e de integração;
- adicionar autenticação e autorização quando o fluxo exigir;
- documentar decisões técnicas e limitações.

## Observação

Itens como SLA, auditoria, processamento assíncrono e sistemas distribuídos fazem parte da direção de estudo do projeto. Eles só serão apresentados como funcionalidades quando estiverem implementados e validados.
