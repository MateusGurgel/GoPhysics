# Contexto

A relação entre objeto e collider representa um cenário de dependência mútua.
Microsoft Learn

Abordei as seguintes soluções:
Mediadores

Vantagens:

    Segregação mais definida entre os objetos.
    GitHub

Desvantagens:

    Desempenho problemático, tanto em termos de memória quanto de tempo de acesso entre as classes.
    Amazon Web Services, Inc.

    Complexidade adicional.
    Medium

# Solução

Propor interfaces para a criação de colliders, que ocultem atributos como dimensões dos objetos, facilitando o uso e eliminando dependências cíclicas.
Architectural Decision Records

# Consequências

Essa solução impede que os objetos acessem diretamente os colliders de maneira simples; ou seja, não é possível escrever um script dentro de um objeto para interagir com um collider sem verificar sua classe antes.
GitHub