# AnderBank

Projeto para desenovlver um banco digital, apenas para estudos. 

## Sobre as funções

- Cadastro de usuário
- Cadastro de conta 
- Transferência de dinheiro
- Extrato 
- Extrato por data 
- Cadastro de cartão de crédito
- Extrato 
- Compras no cartão de crédito
- Extrato listando as compras no débido e crédito

## O que foi utilizado no projeto

- Golang 
- Postgres 
- Gin
- SQL 

Neste projeto foi optado por não utilizar nenhum ORM será utilizado sempre libs nativas. 

## TODO
<ul>
<li>🚧 melhorias</li>
<li>🚫 bugs</li>
<li>🧪 Teste </li>
</ul>

### Lista de tarefas do teste
| Status | Tipo | Descricao |
| :---   | :---   | ------------- |
| ✅ | 🧪 | Rota home get/home |
| ❌ | 🧪 | Rota user get/user |
| ✅ | 🧪 | Rota user get/:id | 
| ✅ | 🧪 | Rota user get/delete/ |
| ✅ | 🧪 | Rota user post/user |
| ✅ | 🧪 | Rota user delete/:id |
| ❌ | 🧪 | Rota account get/account |
| ✅ | 🧪 | Rota account get/account/todos |
| ❌ | 🧪 | Rota account post/account/ |
| | 🧪 | Rota account get/account/balance/:id |
| | 🧪 | Rota account get/account/balance/delete/:id |

### Lista de tarefas do bugs
| Status | Tipo | Descricao |
| :---   | :---   | ------------- |
|  |🚫| A rota user get/ deveria retornar apenas o usuarios ativos |
|  |🚫| A rota account/ esta retornando todos os dados da conta, so deve retornar os seguintes campos (numberAccount, balance, dateCreate, debit, credit) as demais informacoes sao desnecessarias |
| ✅ |🚫| Dentro da rota user/delete/:id o verbo está incorreto, deveria ser um post ou put |
|  |🚫| O campo "numberAccount" deve ser gerado automaticamente, verificar a rota post account/|

### Lista de tarefas das melhorias
| Status | Tipo | Descricao |
| :---   | :---   | ------------- |
| | 🚧 | Criar uma nova rota para conseguir ver todos os usuários, ativos e inativos |
| | 🚧 | Organizar as configuracoes do Isominia dentro da pasta do projeto, para importacao |
