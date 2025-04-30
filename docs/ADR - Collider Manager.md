# Contexto

O sistema de colisões é complexo, precisamos ter um loop verificando a cada update se os objetos colidiram, caso colidiram, temos que iniciar ações personalizadas.

# Solução

Aceitar um observer dentro do collider, que notificará todas as funções presentes dentro do collider
# Consequências

Essa solução irá gerar grande complexidade na criação de instancias de novos objetos.