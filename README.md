# Loans

Esse repositório faz parte de um conjunto de desafios que estão disponíveis no perfil <a href="https://github.com/backend-br/desafios/" target="_blank">backend-br</a>.

## Desafio

Implementar um serviço que determine quais modalidades de empréstimo uma pessoa tem acesso.

### Modalidades de Empréstimos
- Pessoal: taxa de juros de 4%.
- Consignado: taxa de juros de 2%.
- Com garantia: taxa de juros de 3%.

### Requisitos para definir quais tipos de empréstimos um usuário tem acesso
- Conceder o empréstimo pessoal se o salário do cliente for igual ou inferior a R$ 3000.
- Conceder o empréstimo pessoal se o salário do cliente estiver entre R$ 3000 e R$ 5000, se o cliente tiver menos de 30 anos e residir em São Paulo (SP).
- Conceder o empréstimo consignado se o salário do cliente for igual ou superior a R$ 5000.
- Conceder o empréstimo com garantia se o salário do cliente for igual ou inferior a R$ 3000.
- Conceder o empréstimo com garantia se o salário do cliente estiver entre R$ 3000 e R$ 5000, se o cliente tiver menos de 30 anos e residir em São Paulo (SP).