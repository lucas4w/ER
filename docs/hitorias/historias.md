# História de Usuário 1 – Usar Habilidade Noturna

## Requisito original
O sistema deve permitir que o jogador utilize sua habilidade secreta durante a fase da Noite.

## História de usuário
Como **Jogador**,  
quero **utilizar minha habilidade noturna escolhendo um alvo**,  
para **influenciar o resultado da partida de acordo com meu papel**.

## Critérios de aceitação
1. Dado que estou vivo e é a fase da Noite,  
quando eu seleciono um alvo e confirmo o uso da habilidade,  
então o sistema registra minha ação de forma secreta e me informa que ela foi recebida.

2. Dado que já utilizei minha habilidade na noite atual,  
quando tento utilizá-la novamente,  
então o sistema impede a ação e informa que ela já foi realizada.

---

# História de Usuário 2 – Votar para Eliminação

## Requisito original
O sistema deve permitir que os jogadores vivos votem para eliminar um suspeito durante a fase do Dia.

## História de usuário
Como **Jogador**,  
quero **votar em um jogador durante a fase de votação**,  
para **contribuir com a eliminação de um suspeito e ajudar meu time a vencer**.

## Critérios de aceitação
1. Dado que estou vivo e a fase de Votação está ativa,  
quando eu seleciono um jogador vivo e confirmo meu voto,  
então o sistema registra meu voto e atualiza a contagem de forma visível.

2. Dado que já votei na rodada atual,  
quando tento alterar ou votar novamente,  
então o sistema permite a alteração do voto até o fim do tempo de votação.

---

# História de Usuário 3 – Enviar Mensagem no Chat Geral

## Requisito original
O sistema deve permitir que os jogadores vivos enviem mensagens no chat geral somente durante a fase do Dia.

## História de usuário
Como **Jogador**,  
quero **enviar mensagens no chat geral durante a fase do Dia**,  
para **discutir com os outros jogadores.**

## Critérios de aceitação
1. Dado que estou vivo e a partida está na fase do Dia,  
quando eu digito uma mensagem e confirmo o envio,  
então a mensagem é exibida no chat geral para todos os jogadores vivos.

2. Dado que a partida está na fase da Noite,  
quando eu tento enviar uma mensagem no chat geral,  
então o sistema impede o envio e informa que o chat geral só está disponível durante o Dia.