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
A lista de tarefas esta sendo feita conforme e montado o projeto
🚜 melhorias 
🚫 bugs
🔰 coisas para fazer
🧪 teste
<br>
<hr>
<br>

🔰 Verificar se precisa de mais algum campo no schema de user
🔰 Por hora adicionar apenas recursos do debito 
🔰 Rota para consulta de saldo, o que deve ser mostrado para o usuario
🔰 Rota para deletar a conta, neste caso nao deve deletar apenas inativar 
🔰 Rota para fazer um update no saldo do usuario, repassando apenas o ID do usuario de origem 
🔰 Adicionar recurso para transferir dinheiro entre os usuarios
🔰 Criar uma nova tabela para extrato da conta de debito, conforme e feito no nubank


🚫 Ao utilizar a rota "account/balance" e repassar um numero de conta que nao existe, esta quebrando a aplicacao

🚜 Adicionar recurso para o pix
🚜 Numero da conta, melhorar a funcao para gerar esse numero

### Testes

| ✅ || Metodo | Rota | Descricao |
| ❌ | get | user/ | Consulta todos os usuarios ativos |
| ❌ | get | user/id | Consulta uma determinado usuario | 
| ❌ | get | user/delete/ | Consulta todos os usuarios inativos |
| ❌ | post | user/ | Cria um novo Usuario | 
| ❌ | delete | user/delete/id | Deleta um usuario | 
| ❌ | post | login/ | Gera um numero de token | 
| ❌ | get | account/ | Lista todas as contas ativas | 
| ❌ | post | account/ | Cria uma nova conta | 
| ❌ | get | account/balance/id | Consulta uma saldo, atraves do numero da conta | 