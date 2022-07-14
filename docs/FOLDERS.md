
## Organização (Inicial)

Abaixo, o modelo inicial de organização do projeto:

```
project
...
├── cmd
├── docs
├── internal
├── pkg
├── src
│   ├─── handlers
│   ├─── interfaces
│   ├─── services
│   └─── middlewares
...
```

Abaixo serão descritas as características e definições de cada nível apresentado anteriormente

### Cmd

A utilização da pasta `/cmd` garante que o ponto central de inicialização da aplicação seja feito de maneira simplificada utilizando somente referencias internas do projeto, presentes em `/internal` e `/src`.

A separação de responsabilidades garante, de maneira geral, que detalhes como conexão de banco de dados, router frameworks entre outros tenha uma única fonte de verdade, dando ao time de desenvolvimento maior conforto e segurança na manutenção e implementação de novos recursos.

---
### Docs

Utilizada para qualquer categoria de documentação referente as interfaces externas do serviço, sejam elas arquivos `.json` de plataforma de API como postman ou insominia, arquivos auto gerados utilizando swagger, godocs etc.

A ideia é centralizar de maneira simplificada o "onboarding" e/ou consulta dos desenvolvedores, fazendo com que não percam tempo buscando referencias para modificações e utilização dos recursos, serviços e interfaces.

---
## Internal
Responsável por organizar helpers, métodos internos entre outros recursos presentes somente para o projeto no qual ele estará inserido. Dessa forma, sua exportação para outros serviços e recursos não é recomendada (mesmo que em alguns casos seja possível). Funções e pacotes auxiliares podem ser aplicados e criados dentro da pasta /internal à fim de auxiliar recursos presentes em `/cmd`, `/src/services`, `/src/middlewares` e `src/handlers`.
### Como se da interação dos arquivos Pkg com a aplicação?
- Ideia aqui é com cetrar todos os utils da aplicação, funções que recebem dados e retornam dados tendo dependência nenhuma com nada.
---
## Pkg
Diferente da pasta `/internal`, o `/pkg` visa implementar recursos externos (dependências de terceiros e ou globais) através de interfaces para o contexto da aplicação.
 `pkg = packges`.

Implementações de banco de dados, pacotes de ORM, recursos de CLI, pacotes de serviços de gateway entre outros podem ser implementados nesta pasta, além de serem passiveis de exportação para outros projetos, visto que não possuem nenhuma dependência com recursos internos do projeto.
### Como se da interação dos arquivos Pkg com a aplicação?
- Você irá criar uma interface na aplicação para algum fim, por exemplo você instalou um pacote de validação de email. Com isso, nós vamos criar uma interface que será injetada na camada serviço que você deseja utilizar e agora você se pergunta a implementação? Exatamente meu caro a implementação dessa bibloteca fica aqui na Pkg. Muito provavelmente você verá varios `adpaters` por aqui.
---

## Src

Utilizada como pilar centra do projeto, contendo interfaces de acesso, regras de negócio, interfaces e comunicação com a camada de dados e declaração das estrutura gerais do projeto.


---

